---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2033131030312023-2200112222221310-2222232223012011-0200121201203231-3033103333011032-0331022200133101-2230032301132121-2012311230312312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030)
- service.deploy_options.deploy_re_sites.site

<a id="canonical-3321022030130120-1002110101331302-0100203320122111-3111000102200000-3322221112301121-2113110111332210-2101301020111222-2030221002021221"></a>

Type: `"object"`. list nested block, Optional.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122230130031210-0202201133123223-0112102301332032-3300003022200103-3211000312112231-2122212032302203-0332103103000313-1122012131210103"></a>

### Direct properties for `service.deploy_options.deploy_re_sites.site`

<a id="canonical-0113311010313211-1013213202321231-2013023011230210-3232202122330300-1030302312320332-1021322213112303-3222110132133210-2111011033323121"></a>

#### `service.deploy_options.deploy_re_sites.site.name` property

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

<a id="canonical-0203321021210101-3203000310331111-2011311122130021-1130111322332111-3331103212230111-3212320213200222-2013203010231030-1303223120313310"></a>

<a id="canonical-1330312133301132-1131100121103122-3032332110333110-2230123212330231-1332233112310032-2132030220302303-3230313100302013-1000231023222101"></a>

#### `service.deploy_options.deploy_re_sites.site.namespace` property

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

<a id="canonical-1123200303200102-0203230322213110-1021100202020233-1001100222232112-1210213313022320-0033320323321222-1021320322232211-2223322303212031"></a>

<a id="canonical-2112031220312221-3300233012112203-3000210103101200-1133110032230102-0113102111011112-3131130033032230-0100121010321000-3202001300110302"></a>

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

<a id="canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_re_virtual_sites

<a id="canonical-0222222021111203-1231331110311132-3303120220302320-0011221210101320-0230110102300222-1122020302302120-2123131030030213-3320230000331223"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211212120232321-2012132112021001-3120233010032313-1301102120113013-1233002022111121-2000232133233113-0312323013003121-3211330110332323"></a>

### Direct properties for `service.deploy_options.deploy_re_virtual_sites`

- [virtual_site](resources--workload--reference--group-017.md#canonical-2322220331033200-1103303020112230-3013011001212213-3010301232101123-2013313020321222-3011032223303110-2121123030131323-0231033133133200): complete subsection reference.

<a id="canonical-2322220331033200-1103303020112230-3013011001212213-3010301232101123-2013313020321222-3011032223303110-2121123030131323-0231033133133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-017.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202)
- service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-3122121120213322-1321020313301213-2303020310132333-0102230231021201-3313011220013121-1200331201000310-3033131201030031-2313222022013023"></a>

Type: `"object"`. list nested block, Optional.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203230220022201-3320113020010010-2331133010031312-3111323132000023-0031333101133213-0032233221032313-0032313130030310-1123332012231230"></a>

### Direct properties for `service.deploy_options.deploy_re_virtual_sites.virtual_site`

<a id="canonical-3003323132012212-0132303323030003-0203200231311011-2001333132303030-3012310123013020-2022103010200311-0301010021211131-2122212022113323"></a>

#### `service.deploy_options.deploy_re_virtual_sites.virtual_site.name` property

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

<a id="canonical-2311203021302013-3203110322311321-1202022013303300-2311232121301331-1232313103303311-0202321330311113-1113230332200201-2123022230132112"></a>

<a id="canonical-3331000110330320-2212123201133303-3012303333331312-0201132000332012-2230233121103232-2120223023112020-0313303120223310-3321122301222023"></a>

#### `service.deploy_options.deploy_re_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-1222330022103120-2202313202211022-3203020023021123-3210330322210233-3121311001230131-3330021303023223-2212310300202110-2013131221031330"></a>

<a id="canonical-2323132011131313-2201013313320302-0023220030312301-3232133112233123-3211201313002202-1211232301323122-3323031310323020-3211002201212313"></a>

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

<a id="canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.scale_to_zero` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.scale_to_zero

<a id="canonical-3012300321212121-1132001221112031-2333022003222023-3111011130310012-3031202110320300-3011220322003230-3302003002322201-0011302201032002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
scale_to_zero = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.volumes

<a id="canonical-0323112231212112-2031033023310330-0111121103130221-0300012223232011-0303013233030011-1113121231003200-3012212103230331-1211323311213333"></a>

Type: `"object"`. list nested block, Optional.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020220300310201-3011211220000330-0310111033220300-3311111121313320-0122201122330103-3210212320301203-1111030223133023-3223210232120333"></a>

### Direct properties for `service.volumes`

- [empty_dir](resources--workload--reference--group-017.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000): complete subsection reference.

- [host_path](resources--workload--reference--group-017.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223): complete subsection reference.

<a id="canonical-0210321122110332-3332130230330223-0310321033222331-1120102210232313-1113320201201213-1132320030221313-0033130303003012-0301121201203300"></a>

<a id="canonical-0033310011033013-2301201203122103-0203310313313203-2331002122122031-3123201010031322-1122120031220211-0110130021101210-1210001310213012"></a>

#### `service.volumes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [persistent_volume](resources--workload--reference--group-017.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331): complete subsection reference.

<a id="canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.empty_dir` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.empty_dir

<a id="canonical-0013300133100100-1113110123121321-3111203122110000-3001230320302321-1033333323033011-2102333222220223-0221203100323332-1123020123001331"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
empty_dir {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001201103000103-0203100000113101-0212022211332012-2101001231333203-3011232200133102-0200323032331200-3301303002100313-0220200113001312"></a>

### Direct properties for `service.volumes.empty_dir`

- [mount](resources--workload--reference--group-017.md#canonical-3202131021131212-0101213023301012-0231220211323020-2331102001300001-1020233213033112-1212100032103200-1202110312132003-2211301123313232): complete subsection reference.

<a id="canonical-0132332000113213-1120211203132332-3200303132331322-0001321132202203-2231132001003331-1312030201123212-2213103232013103-1231211031322233"></a>

<a id="canonical-2203312030221323-3031332010313120-0133031110002023-3122212222300333-2031013112311132-0313100200130102-0322200332301220-2203330003331100"></a>

#### `service.volumes.empty_dir.size_limit` property

Type: `"number"`. Optional.

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

<a id="canonical-3202131021131212-0101213023301012-0231220211323020-2331102001300001-1020233213033112-1212100032103200-1202110312132003-2211301123313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.empty_dir.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.empty_dir](resources--workload--reference--group-017.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000)
- service.volumes.empty_dir.mount

<a id="canonical-3200110313320120-1301012232022302-3002200123222132-1113212133033223-2222012030122222-0331101330013120-3213301223312132-3123100222010012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112032131032123-1221321212011002-3131030320321002-0032330200210311-1333132023122320-2201322233102220-1033002202303300-1012130013222231"></a>

### Direct properties for `service.volumes.empty_dir.mount`

<a id="canonical-0011311212322030-3000213100333003-0303013013110321-1020212111032132-0300121132203120-2022133203301113-1303033232220120-3332212301120201"></a>

#### `service.volumes.empty_dir.mount.mode` property

Type: `"string"`. Optional.

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

<a id="canonical-2112223001233030-1032021033333131-0122001321010212-0230122121002230-3102031231203323-2310221123231022-0333031002333201-2120311101020022"></a>

<a id="canonical-2131321003220330-3002022202211121-2131221323301023-2112120013233102-3303302123111122-3232311011232010-3333022311101033-3300132030232021"></a>

#### `service.volumes.empty_dir.mount.mount_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1223032003000302-0200111301032101-3203103211120130-2211301313111220-3233012131033321-2013122331220133-0110102211321033-1213130323033000"></a>

<a id="canonical-2320133203113010-2221201103323120-3022232321232232-3121211222013030-3332211022302302-1321002033030303-1013310033030020-0022012030122201"></a>

#### `service.volumes.empty_dir.mount.sub_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.host_path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.host_path

<a id="canonical-0102110212031231-2121203200102230-0133203021023222-1310331100120120-1121032013330202-2123121310133130-0311210032303203-0023311221310001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331132311012233-2010132130010233-2032313023322010-3200220022111132-3221320013013301-0022013020301113-2202212120232113-0212313013330112"></a>

### Direct properties for `service.volumes.host_path`

- [mount](resources--workload--reference--group-017.md#canonical-1303131310232102-2011333101230102-1121333013101031-3102011101122303-0320111112321001-1102313022211132-2210303010130031-0203210032311333): complete subsection reference.

<a id="canonical-1210301313203011-3113123120033122-2223131110021130-3031310011100010-3131233113333112-2223210232113201-2302323113122312-1223321030231303"></a>

<a id="canonical-0200122132122033-1132333200123033-3303000203221221-1233222121113120-0022032323032110-2213100101021021-3301211212031131-3203001020021331"></a>

#### `service.volumes.host_path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1303131310232102-2011333101230102-1121333013101031-3102011101122303-0320111112321001-1102313022211132-2210303010130031-0203210032311333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.host_path.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.host_path](resources--workload--reference--group-017.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223)
- service.volumes.host_path.mount

<a id="canonical-0321101100023131-2010212020031132-3110311213201330-2010122221122013-1020333102322302-1102230121332220-2301312003323333-2013131022001013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303013000202133-3013313222002303-2303103323113012-1201130223130100-3303201200111020-3320300033210222-1212020231002002-3022132221012010"></a>

### Direct properties for `service.volumes.host_path.mount`

<a id="canonical-1110132323310221-2311103023310013-0120310210033123-1132303132123210-0111323323020233-2222122031330032-2101113302001303-2030123012011332"></a>

#### `service.volumes.host_path.mount.mode` property

Type: `"string"`. Optional.

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

<a id="canonical-0003210232113002-1020033232313322-0323303012133323-3100003101000032-1232211101323133-3300022330232013-1223122311330202-0233310110031311"></a>

<a id="canonical-2333330013130102-3013022102031020-1301330033221232-3033021313201020-2220301320001312-0321102100131011-0003202132203130-3333203310011223"></a>

#### `service.volumes.host_path.mount.mount_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3221101333020313-2033123201020102-0011022230030232-3120301332300202-3131220322003001-0022032201321010-3201230030132323-0301120331002100"></a>

<a id="canonical-3213111221100131-0103130322120131-3123010020303123-1103000220131101-3300202013230023-0010212003120211-1211031233033323-1133133022302221"></a>

#### `service.volumes.host_path.mount.sub_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.persistent_volume

<a id="canonical-0331030333110000-3311030122030203-3313001232120320-1301111312302322-2101030020330020-0103202221303122-0013110333330223-1110000022313101"></a>

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

<a id="canonical-2330012312210333-3203020323322131-2033112023212311-1101300233311032-2231030121210233-0313120202210013-3330123202223101-1012033121122103"></a>

### Direct properties for `service.volumes.persistent_volume`

- [mount](resources--workload--reference--group-017.md#canonical-0230211303132020-2130130313320130-3020231002113003-1021322300103102-0220013320231100-0021103211212131-2122030310213311-3112100123213012): complete subsection reference.

- [storage](resources--workload--reference--group-017.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312): complete subsection reference.

<a id="canonical-0230211303132020-2130130313320130-3020231002113003-1021322300103102-0220013320231100-0021103211212131-2122030310213311-3112100123213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-017.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- service.volumes.persistent_volume.mount

<a id="canonical-2201312230202120-2212303012330022-1312101133332121-3122030020131131-1202230120332130-1100331022032002-3120132023213113-1313210313131102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021001311013321-0220223232231302-2131220233103231-1032210100033013-2013333302121302-2131101310311302-0200013030121211-1202302030332100"></a>

### Direct properties for `service.volumes.persistent_volume.mount`

<a id="canonical-1203302302020000-3033220120113321-3133103221133201-3021122333031201-0231331333201031-1232322312311031-2032022003311011-3303211200333223"></a>

#### `service.volumes.persistent_volume.mount.mode` property

Type: `"string"`. Optional.

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

<a id="canonical-2112303321133312-1310333020010002-0221202303110032-0323311303322323-1021122223100133-2322203221131220-0220112023131300-3221121232002112"></a>

<a id="canonical-1323013230303312-2120021212131223-3210102113323220-2101220102323213-3213310102001321-3313123330030320-3020230303113121-1021013030232122"></a>

#### `service.volumes.persistent_volume.mount.mount_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0213021022230223-3203233101132120-0303222112130111-0133122020332301-0200103200200012-3210000122220203-2320323122302012-3013103010201220"></a>

<a id="canonical-2220001320000210-1121321320022303-0301110333300232-0322222022123203-1003032330203313-1131221330100010-3313313022330203-3330220310230220"></a>

#### `service.volumes.persistent_volume.mount.sub_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-017.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- service.volumes.persistent_volume.storage

<a id="canonical-1133211230011302-3231213011010102-1032132310232121-1213303303232213-1333211102102232-1002333020110033-3301332002031300-0233131221122002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220010331212122-1320100021313213-0201031031033021-2223331133101230-2030023101001002-0211011113233020-1210113121330101-1120332231010003"></a>

### Direct properties for `service.volumes.persistent_volume.storage`

<a id="canonical-0121011120331330-2111121210033102-3030110210222021-1203100330132200-0213220202303201-2002030323211311-2203211122213303-0201330330102102"></a>

#### `service.volumes.persistent_volume.storage.access_mode` property

Type: `"string"`. Optional.

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

<a id="canonical-2101130010203102-2302133001230232-2210321033131213-2300233012103302-2320013331130211-0032110001132103-1213033213223033-3303230333220201"></a>

<a id="canonical-2330002010122300-0030300022021223-0233000233131022-0230210231013100-2102221301002233-2111120321102001-2020220002321221-1013113313111332"></a>

#### `service.volumes.persistent_volume.storage.class_name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](resources--workload--reference--group-017.md#canonical-3030021232102132-3032123220212030-1212020002332312-1203311122221322-2301311133233201-1131300212232312-1331331120023212-0103123233033101): complete subsection reference.

<a id="canonical-1000011321211123-1330211120011013-0111212123123010-0330232300322032-2303101101322300-3323300203010203-1202302321122313-3311211001020000"></a>

<a id="canonical-2332010332200031-2002033020231300-2202022132221231-3200303032223211-3232021201111320-0220021220132022-3202132031332201-1210232102033210"></a>

#### `service.volumes.persistent_volume.storage.storage_size` property

Type: `"number"`. Optional.

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

<a id="canonical-3030021232102132-3032123220212030-1212020002332312-1203311122221322-2301311133233201-1131300212232312-1331331120023212-0103123233033101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-017.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-017.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- [service.volumes.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312)
- service.volumes.persistent_volume.storage.default

<a id="canonical-3232323111203121-0022330230132021-2122021001122201-1311212300230212-3203320002100233-2310222010011212-3122323001311033-2313002010023222"></a>

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- simple_service

<a id="canonical-3033330332311101-3210113002001102-0301111101113100-3023101200113133-0000233331000233-1000221300201111-0133220110233111-3012103111100322"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
simple_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203002220302032-3131301330311110-0003002013001012-2003000113203113-3300113031010323-0222312211302001-2012120212310300-3233011032333123"></a>

### Direct properties for `simple_service`

- [configuration](resources--workload--reference--group-017.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313): complete subsection reference.

- [container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213): complete subsection reference.

- [disabled](resources--workload--reference--group-018.md#canonical-1110023311313011-0202000212230320-0210021130232023-0323322102330130-3223121303212321-2033302130023021-3120001020331002-1031122330013020): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-018.md#canonical-0003300212123321-0030022111222023-1030022232221021-1132031212101132-3230032012113311-1303232022202331-3111103333310112-3310021021132230): complete subsection reference.

- [enabled](resources--workload--reference--group-018.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030): complete subsection reference.

<a id="canonical-1002131132203101-1131232101213330-0032322330310111-2211001301212333-2231300212010012-2332222020210322-3110330322123132-2002212121103302"></a>

<a id="canonical-0001302321230311-1303111333132020-1320030010001303-0301102031323123-2100322301212331-1102300003031111-1102300323321122-0110213123021020"></a>

#### `simple_service.scale_to_zero` property

Type: `"bool"`. Optional.

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

- [simple_advertise](resources--workload--reference--group-018.md#canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202): complete subsection reference.

<a id="canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.configuration

<a id="canonical-1232232331100102-2211301230213012-2322020311120000-3001332301233103-0320330130231120-3132333102232030-1333212310002333-3221332033223121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003232031000132-1112101012021122-2111301233011100-2303300313332331-0021111001102310-1312010033112310-3030303101013113-3130111230021310"></a>

### Direct properties for `simple_service.configuration`

- [parameters](resources--workload--reference--group-017.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130): complete subsection reference.

<a id="canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-017.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- simple_service.configuration.parameters

<a id="canonical-3023003010131101-3113213303201203-2310202100320233-3003231031023112-3030102021331012-3303003002220203-0301032312231012-2011223312103202"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121202202020123-0020123132313122-2122111101313102-0200221211223303-2112231013231131-1133200021111231-3112112220123011-1101333020131331"></a>

### Direct properties for `simple_service.configuration.parameters`

- [env_var](resources--workload--reference--group-017.md#canonical-0102122312313001-0221331223003110-2133110201320030-0012331100210230-0200113132332023-3121201230202020-1010310113323210-3103220100311122): complete subsection reference.

- [file](resources--workload--reference--group-017.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102): complete subsection reference.

<a id="canonical-0102122312313001-0221331223003110-2133110201320030-0012331100210230-0200113132332023-3121201230202020-1010310113323210-3103220100311122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.env_var` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-017.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-017.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- simple_service.configuration.parameters.env_var

<a id="canonical-2332232003231032-0100033022223300-3212132001013332-0320210000212210-1311111231210310-2121300330232222-0303321222031001-0111331333313011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313120121112031-3003113031110013-2320201301110200-1230302030220211-1331123100033223-0033213333332301-2023110322101101-2023330231103213"></a>

### Direct properties for `simple_service.configuration.parameters.env_var`

<a id="canonical-0120033022232303-2012303111030203-0200102330312023-3122330232331221-0131112303100020-0322333001223223-2213221100100031-2103011310010211"></a>

#### `simple_service.configuration.parameters.env_var.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3122002020102310-0313020230020302-2110032332020012-2023001302130333-3100213120310032-2120333322132232-2223023230232023-0300320122120122"></a>

<a id="canonical-1323101332311102-3201230233311122-2020310220210122-3200223310010223-0023010222232013-1223033122121302-2033020023112111-2221030332011123"></a>

#### `simple_service.configuration.parameters.env_var.value` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.file` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-017.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-017.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- simple_service.configuration.parameters.file

<a id="canonical-1001032302332122-0310100122133230-0212100223222003-3010230130313121-1001033000113132-1231020332102030-2220111103133300-2210213013300300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000311332220331-0001223031310113-0312222111230222-1222311100311012-1310200103102013-2201330002102210-3102210311133130-3301211300222331"></a>

### Direct properties for `simple_service.configuration.parameters.file`

<a id="canonical-3010210033130313-0312302212012231-2031002312110322-0320200011300130-1310233301223223-3133033103322301-1031033231001301-2003231320003121"></a>

#### `simple_service.configuration.parameters.file.data` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-017.md#canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102): complete subsection reference.

<a id="canonical-0213311303231321-3332331213102132-1301001212001022-1331231121223220-1232323021110110-2213123102322022-1220301002303212-0202110322023323"></a>

<a id="canonical-1023232121223022-1330003010200033-3321102311212101-2332322302030031-2323023323323211-2121111323301111-0132101021312333-1101323313022112"></a>

#### `simple_service.configuration.parameters.file.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1020102202211311-0022000202330231-3012011302303222-2000023221012032-1000113311001021-3003322001010120-1330222033303311-3010001300123210"></a>

<a id="canonical-0133000202230311-2013022120001313-0323300230123023-3321020231113221-3302322013112303-1012200021023310-3110020030212120-0021331132222122"></a>

#### `simple_service.configuration.parameters.file.volume_name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.file.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-017.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-017.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- [simple_service.configuration.parameters.file](resources--workload--reference--group-017.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102)
- simple_service.configuration.parameters.file.mount

<a id="canonical-3212302022331323-2211221201222003-0123112130020313-0021121131303213-3002223111210012-3321032111020210-1020030120103103-1111213123200211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130320330101330-2222133220210332-3320120333220033-1111300323222110-3322212013020132-0331312020121223-0200023113021032-1300231323203020"></a>

### Direct properties for `simple_service.configuration.parameters.file.mount`

<a id="canonical-0003012212033002-0103213210031101-0321200331020101-3113210023133330-0302233113032010-2231203212002302-0012001312110210-3231032021002200"></a>

#### `simple_service.configuration.parameters.file.mount.mode` property

Type: `"string"`. Optional.

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

<a id="canonical-3212330223330222-1333101133203312-1323121233110023-3230103101201201-1001213011133211-0030032332130231-3212133222021230-3323213112201131"></a>

<a id="canonical-3301200100203221-2111312030020000-1020233021102130-0120230000113033-0220001221111211-2002131120022212-2020213002230331-3113010321312013"></a>

#### `simple_service.configuration.parameters.file.mount.mount_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0313113033320011-0122021333002211-0021301013220210-3112230121303232-3123030012020321-3203213133021130-0133000310031130-2023220213000122"></a>

<a id="canonical-2313010302111030-2101010030021303-1031121133022220-0223321200133022-2032120232220313-3210321011133301-2213200323330313-3002333232123211"></a>

#### `simple_service.configuration.parameters.file.mount.sub_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.container

<a id="canonical-0000332130000122-3333320032231313-0322210332113231-0003302103123222-1330101203111233-3131323011322223-2330000001200103-1031033200220213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
container {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101130232232331-2300013222111331-3013102321321013-1223111320232103-1233202033303202-0112101032220333-2230312332023311-2102021102212020"></a>

### Direct properties for `simple_service.container`

<a id="canonical-1320122102200012-0000332123203102-2100330202122002-1331133230000220-1323231203132120-1113111032322020-3300332101320233-0110002002331332"></a>

#### `simple_service.container.args` property

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1223111301003112-2002220311333111-3002231330301122-3211022012210103-1032210300312001-2322112323021120-1123010111130211-0220201310012133"></a>

<a id="canonical-2212202301123302-1222020002301220-2223321211233103-0032213100300201-2123021232211211-3022320312002203-3030303330331333-2001223010333302"></a>

#### `simple_service.container.command` property

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-017.md#canonical-1101122001010203-0213022111133112-2022320121302331-0210332313220111-1122010030301033-1213212030220200-2311111310312301-2202210102321311): complete subsection reference.

- [default_flavor](resources--workload--reference--group-017.md#canonical-1021322231332022-0200122020303230-1103121113233003-0303121302202231-0113031111122210-3022123231231321-3331113133012000-1201233031220130): complete subsection reference.

<a id="canonical-0232133120123001-3000001100013103-3321303323200310-1113302213220102-0202212323033013-1312313131211112-3231031133132202-1333123301332020"></a>

<a id="canonical-3201102102112011-2131320300230310-1130102322010333-3202310323002030-1001310130312323-0103113002000202-1003003013021303-1121123022321303"></a>

#### `simple_service.container.flavor` property

Type: `"string"`. Optional.

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

- [image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013): complete subsection reference.

<a id="canonical-2202013202312211-3221210010311123-1320130103202201-1220032201122102-3020201220212220-0023200133230011-0323110323133312-1212022122101330"></a>

<a id="canonical-3020313233311232-3003032232302000-1213212013330321-1110123310111132-3021300022033010-3122201103113101-1131022310201232-1102223001002212"></a>

#### `simple_service.container.init_container` property

Type: `"bool"`. Optional.

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

- [liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112): complete subsection reference.

<a id="canonical-2030101133223013-2330003002033321-1023321232213223-0330131033202202-0200310111311312-3031111320011023-0131212102003130-0032033131013033"></a>

<a id="canonical-2331121112213100-2220322210101321-3010210131102030-3231101002003302-3132211231123221-3320332203221111-1310330123313322-1220213320122321"></a>

#### `simple_service.container.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323): complete subsection reference.

<a id="canonical-1101122001010203-0213022111133112-2022320121302331-0210332313220111-1122010030301033-1213212030220200-2311111310312301-2202210102321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.custom_flavor

<a id="canonical-0231332000323331-0232001121003230-0213101031233131-2012030002323113-1030023121210121-0110320332210231-2212030133303110-3013310020232113"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310223003001321-0020230331022332-0220333101213221-2220102020002301-1322302312223102-2100231321010331-1330021121121031-3132221010300110"></a>

### Direct properties for `simple_service.container.custom_flavor`

<a id="canonical-3302322202313131-2231021121010320-1220003221001332-3010313112322221-2001302123111132-2030231321221222-1313330132202133-1300021131203321"></a>

#### `simple_service.container.custom_flavor.name` property

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

<a id="canonical-2101112003302230-0300123230020103-1111322323103302-3203123013110230-3102303322332001-0202203233011232-2010100302013301-0030101323120100"></a>

<a id="canonical-2220102012100031-2221320210122223-2000123212103120-1000131130001133-3230032221220223-2131101001231222-1330331120130223-3102121032003201"></a>

#### `simple_service.container.custom_flavor.namespace` property

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

<a id="canonical-0020220123010030-3201311203200111-3200113033123313-0011201123312003-0323212312231000-0101030011003232-2013001321031012-0123021220123232"></a>

<a id="canonical-3200123313220312-2331011132313230-3111021301130020-3002200002122231-2112200031313312-0330030113002010-3012023003301031-0010002221311330"></a>

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

<a id="canonical-1021322231332022-0200122020303230-1103121113233003-0303121302202231-0113031111122210-3022123231231321-3331113133012000-1201233031220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.default_flavor

<a id="canonical-3021132000011123-3113222221321033-1121321300311302-0220331332233301-2123010013212010-1333213021222221-3120110033310302-3002132102130130"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_flavor = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.image

<a id="canonical-1112332221131201-2021312133010013-3003121030022100-2101300231213223-2110203022112330-0301010102222201-2113011322302120-1020022311110130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233101213131113-0121203023320003-1131231203213220-0002121133122000-2000313213112311-2103131013002310-0301223321213101-0103312313303000"></a>

### Direct properties for `simple_service.container.image`

- [container_registry](resources--workload--reference--group-017.md#canonical-1200331112133231-3310003010123012-1110023010303211-2311012310220121-3021122231012000-1211101013332022-1003223220212311-2101321331222021): complete subsection reference.

<a id="canonical-3312301013333322-2011003320302201-2202132121203111-1022211130322201-2213200101012301-0100322131020203-3021200023033001-2031302323212232"></a>

<a id="canonical-0123123320332333-0221102102211220-1113123321132101-0100112100113322-1310002230100233-2323130011001323-3112103221100331-0003031001230222"></a>

#### `simple_service.container.image.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](resources--workload--reference--group-017.md#canonical-3011103200100232-3221201210120303-2332112223300201-3020111020011211-1313113233020102-3330003212121100-1100202011112101-1131031023231022): complete subsection reference.

<a id="canonical-0312202002101221-0303231031211302-3111213233122332-0222331323111210-2320232220010000-3030000323202111-2233303221123312-3232031201100330"></a>

<a id="canonical-0300223112120320-3201023111231223-0113001132030023-3301223210113332-3010212210123311-3203132100320210-3013010312321301-1120123032310220"></a>

#### `simple_service.container.image.pull_policy` property

Type: `"string"`. Optional.

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

<a id="canonical-1200331112133231-3310003010123012-1110023010303211-2311012310220121-3021122231012000-1211101013332022-1003223220212311-2101321331222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- simple_service.container.image.container_registry

<a id="canonical-1220011200030120-2301130113121230-3030323311330210-2100331133303323-1123111012311002-0000301231001210-1033223332312010-3202103023133201"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220013320033321-0311300110131301-3113313133222310-0030221101323313-3023133311030220-3232322010210311-3130302210032111-3011010322120020"></a>

### Direct properties for `simple_service.container.image.container_registry`

<a id="canonical-2312213330311131-2123032220303000-1002121223013033-1130202111300013-1130102110111303-0102312223333231-2222110311133221-1110333033011233"></a>

#### `simple_service.container.image.container_registry.name` property

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

<a id="canonical-0132223300020233-2110212110023321-0031022101200313-3131200331111330-3202320232120020-2121112032212022-3110331300130211-2210210023233100"></a>

<a id="canonical-2321220200212011-2020032303010022-3330022113311313-0301330012233211-2330330132312130-1210120233120330-0002310001133133-1230330221231101"></a>

#### `simple_service.container.image.container_registry.namespace` property

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

<a id="canonical-0330330222010330-3013032201003222-0323210201223313-0130033133300303-1112320330100202-1123112212111330-1021303111222021-3220213233203132"></a>

<a id="canonical-0322203132211232-3310311322031130-0303200212203012-1201233032101123-2020010102001021-3122103220320331-1101013331030202-1023312002201132"></a>

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

<a id="canonical-3011103200100232-3221201210120303-2332112223300201-3020111020011211-1313113233020102-3330003212121100-1100202011112101-1131031023231022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image.public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- simple_service.container.image.public

<a id="canonical-1113223313210011-1320221031202012-2030301033121031-2020001320030322-3112001120303322-2332111331011030-3001103202033023-3102231333210111"></a>

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
public = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.liveness_check

<a id="canonical-1011300101022230-3003203220333333-0203030211212102-2213201331300233-3000011330103313-0122222001012300-3313122211120231-1132231303103200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020121122012002-1212121311130133-2011220120101123-0323131120213010-3230212002333302-3220330030101013-3312201132333331-2313100203203131"></a>

### Direct properties for `simple_service.container.liveness_check`

- [exec_health_check](resources--workload--reference--group-017.md#canonical-1000233212122031-3123100131313301-3002121002223331-3013313223310031-1103132102110303-1102110103010010-1001300002102322-1211131313112020): complete subsection reference.

<a id="canonical-1210303210210322-1210013131201033-2303010122220331-2321132003312303-3311222300000103-2233210021220113-1111113122022100-2332310120301303"></a>

<a id="canonical-1022212312322223-1130123100200322-3303130110010100-3313122303121302-2302102111231331-3330002301012301-3213300232203321-0221122333120112"></a>

#### `simple_service.container.liveness_check.healthy_threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301): complete subsection reference.

<a id="canonical-0212212103321330-0332111033203130-2020012210030310-0130221000313120-0030110233232201-1032230203330022-1020032010322312-2031012111002312"></a>

<a id="canonical-3201222121003111-0202121311130220-2310013023120133-3122010222221113-1032032122002113-0230102222012302-2221301201133003-0223221332010102"></a>

#### `simple_service.container.liveness_check.initial_delay` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1210202100003123-1210310101322012-2231303131200113-2331013111122130-1211201200232020-2211200310322312-0332113220310221-1222111232211011"></a>

<a id="canonical-0333011311022202-1330103031131001-1313132301223030-1112210201303010-3313311100220000-0210132313122212-2232302001133020-2220112210101122"></a>

#### `simple_service.container.liveness_check.interval` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [tcp_health_check](resources--workload--reference--group-018.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111): complete subsection reference.

<a id="canonical-1112203121121123-0000033231010203-1002022132333133-0222102200231323-0033103303131102-3010223111203320-0203022311001212-0303311232321330"></a>

<a id="canonical-1311221012202023-0212310122211323-1103332220200120-0201113302122002-1121122030100030-3110311123033230-0233333011121100-1212330322321033"></a>

#### `simple_service.container.liveness_check.timeout` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3122222111113102-3222330003302223-0121002121200333-2313021032233032-1120221123310202-0103330331203020-3302232012023331-1011220033211212"></a>

<a id="canonical-2020133331011030-0011201003321111-2110223212310020-0112110020332210-2332301222033301-1312210211312032-0130321002202131-2233022313033130"></a>

#### `simple_service.container.liveness_check.unhealthy_threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1000233212122031-3123100131313301-3002121002223331-3013313223310031-1103132102110303-1102110103010010-1001300002102322-1211131313112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.exec_health_check

<a id="canonical-0300111010113133-0300210221221301-1011202110132121-0301023133233323-3302132323221031-0223301220111300-2213303211221201-1023200301033330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223322321320102-2120120321023011-1310323013030103-3031200212032211-1133200311121202-0113211203011033-2332302333220333-1033020022313011"></a>

### Direct properties for `simple_service.container.liveness_check.exec_health_check`

<a id="canonical-2001330122313001-1120332211030330-0221301012230112-0330333203200332-3331223031100201-1023101321132232-1200300003230221-1323031021103000"></a>

#### `simple_service.container.liveness_check.exec_health_check.command` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.http_health_check

<a id="canonical-1221201200330113-2312001202323230-3010130032113001-3011313001310013-0332221221100000-2233123121313033-0011033311331330-1132300003122133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103323001211213-0312332031123030-1013103031110020-2033132112033132-0103111013301102-2130130100311333-1110012130131003-1023210202300310"></a>

### Direct properties for `simple_service.container.liveness_check.http_health_check`

<a id="canonical-0300331312003222-3021230020002323-2213321020331212-1323112330131020-3201000212021120-0332023300311003-2032332130112101-0222211000103000"></a>

#### `simple_service.container.liveness_check.http_health_check.headers` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-0021112000311311-1013231323200233-1032331023211203-1313133220220300-0130203311102303-1301122010202112-3331103311121221-0323201320113223"></a>
