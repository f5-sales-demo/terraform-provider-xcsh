---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0000200030102031-0302100111210302-2323113211322313-2133230131002012-0312030000211033-1300200221332213-3313210133020021-3201321213311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1003333201033021-0302001102323231-0012211203023021-3003110100310220-3203311303332013-3133000310222133-1112112213012132-2131030310011020"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010313222101213-2030222023120120-3033111100332312-3000213103211301-1002213101332020-1012332212021033-2333200203003133-1002230333101111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2320121312112011-0031310232210001-3110210310120223-0323101321133103-0111033130011101-3310112330111321-1301222231113020-0302000132233123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213320000020033-2120001112320123-2330332010013101-3101232320223210-0223322031210210-3112031102302331-0322121121003330-1000320310232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2302203231130002-1022211330223011-3312101311123222-2132331300322312-3030300300310111-1023321333213132-3020130320230010-3202202312010113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122121110210020-0112331300201220-2103232132122022-2101302110211121-3013232333012111-3310330202000103-2001120322233003-2030011312230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.non_default_loadbalancer

<a id="canonical-2322103023232120-0303332033010212-1010213033132213-3232003301130121-2322230231112232-3232202311023330-0332322233332133-3010110313023103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320021323133122-2131132203230011-0222102300201203-2133211311213200-0330211132231220-3302123120111202-3201010230103132-0121030000310101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.pass_through` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.pass_through

<a id="canonical-1233131321112112-2023201230013123-3111331122222310-1320121132002322-2001303103033322-0312232120032130-3311231213003322-0001211320110232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.tls_cert_params

<a id="canonical-3233301011213101-1000031112213101-0010332210132112-2303320013132133-0320331321232100-3331323331331002-1320220313310221-2033203310031223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101021200030223-1303021110003002-0103231321010202-3010302220110213-1212200103020302-3201223313203002-3130132130312021-1111200010333310"></a>

### Direct properties for `https.tls_cert_params`

- [certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2023212112103030-3233311000100230-1003201100002300-0233320011103333-1010103112010201-0010203030023201-3303221000120001-0300233313212031): complete subsection reference.

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2122111201321030-0111221001022223-0300313130002313-2221332311201220-3312021030230011-2132312132200123-0003123331003331-3221021123223023): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230): complete subsection reference.

<a id="canonical-2023212112103030-3233311000100230-1003201100002300-0233320011103333-1010103112010201-0010203030023201-3303221000120001-0300233313212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.certificates

<a id="canonical-1032321203121323-0030221302320223-2300110232011230-2111033221000320-3122321032020332-0110020132032110-2231303202301120-1033332113211033"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000030212020210-2313032303222103-0030231022011211-0311011300312111-2203011132313220-3030100311301321-1301332221021121-0222013023003003"></a>

### Direct properties for `https.tls_cert_params.certificates`

<a id="canonical-3003210313232232-0003212302213011-1320123112133003-1331201302133211-1331111202230012-2113013323202112-0321112103213222-0022032112123211"></a>

#### `https.tls_cert_params.certificates.name` property

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

<a id="canonical-1000102132201202-2212033332103002-3302230023111010-0030210223232202-0321211230322333-2321232130120130-3230311332303120-0200201021102303"></a>

<a id="canonical-2230010231013023-2123130222012233-1202332303303311-3211312300110013-2220001003111311-0301312303211213-3011213113212013-1022300313202003"></a>

#### `https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-3121220331101020-3032301300311323-1003312030302301-2010220122330121-3212133211323013-2123121121310320-3033300321021213-2302110033100023"></a>

<a id="canonical-3222211220021030-0211320130312031-1101211131120132-3121302100310032-0120130302210123-3020310023022112-2203130130123102-2110111020303021"></a>

#### `https.tls_cert_params.certificates.tenant` property

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

<a id="canonical-2122111201321030-0111221001022223-0300313130002313-2221332311201220-3312021030230011-2132312132200123-0003123331003331-3221021123223023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.no_mtls

<a id="canonical-1132213030312033-2222310202131003-2030022120130213-3113032122013002-2002331221202221-1232033023100201-3110110303212201-2332102001030130"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.tls_config

<a id="canonical-2222312213321130-2103021001100102-0322033023130312-0033331332210103-3110331311312231-2331213002030220-0323331203011123-0201100130130232"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331310213322130-2300122031002113-3302120330110011-0303021100033212-2112123313220232-2233232331122001-2223103221011112-1200213113200123"></a>

### Direct properties for `https.tls_cert_params.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-0330033022331213-2311312122223110-2312120231032020-3030213213300212-2303010203200323-2331201011131322-2110333211110310-0101332020332220): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-0213322030003310-2001113002220300-3012211332302111-0131203303321231-0300333223202100-0212011001331133-3303230010120130-0023111213123200): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-0001210013221011-1011323010230003-2221332100031113-1312303212000221-1223011221210003-3302203201110333-0233010233220032-1200330303230322): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-2321013122303030-1033102001121011-0031020310333012-0130231021202011-1012033131102332-0003111330102131-3022100210113001-1333013311203322): complete subsection reference.

<a id="canonical-0330033022331213-2311312122223110-2312120231032020-3030213213300212-2303010203200323-2331201011131322-2110333211110310-0101332020332220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-1222221210221121-1113023102011003-0013031122311030-2131032113032323-1210031111122202-0013031231201102-0220123202113101-3122222121311330"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131101023121210-3313202201213130-1232032011111013-0001013121021220-3102031311031000-0111131233303222-2030100030022002-2222333111213231"></a>

### Direct properties for `https.tls_cert_params.tls_config.custom_security`

<a id="canonical-2022231313033211-3113221112313301-3100231101122212-3031221011103202-2010201300012223-2030302011303032-0002012311320100-0310312103021210"></a>

#### `https.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233223030231011-1313212130300210-3331331331222032-1213021221301123-3333121101030022-1322311002201211-0202100211111322-0101010201120210"></a>

<a id="canonical-0132321021230311-3313310030223100-1222323030000023-1030022031131223-2003323033331100-0231132310233221-2100330100230003-0120112103311133"></a>

#### `https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102301103033132-2021232310121220-3000023112230110-1123031233033101-2321303313230212-2233003120032122-3100311310201121-1310021221133333"></a>

<a id="canonical-0121131312131020-3210103311110102-0212113013131031-2221032002303211-2133322111300302-3002120331113131-0312123110223000-3023332031100100"></a>

#### `https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213322030003310-2001113002220300-3012211332302111-0131203303321231-0300333223202100-0212011001331133-3303230010120130-0023111213123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-3003231122102310-3332132322221321-3201111122100322-2212332303132201-3130103110310121-0112000022113110-1133332030013010-3103111121103000"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001210013221011-1011323010230003-2221332100031113-1312303212000221-1223011221210003-3302203201110333-0233010233220032-1200330303230322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-0212033313212011-0231033103121300-0302030111213210-1132232030332200-0210122303200121-0121021212333301-3221330301002111-0311120121213023"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321013122303030-1033102001121011-0031020310333012-0130231021202011-1012033131102332-0003111330102131-3022100210113001-1333013311203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-1210202102322020-2331320101222301-3032232132000320-0201320103213320-1322101000131111-2213012020203223-1202302113230331-1210002311011200"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.use_mtls

<a id="canonical-0120233131212230-3003131013010111-2010100333223103-0002023133011232-1322001121301002-1230111331101223-1132212110333002-0112201203133303"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113300110111311-1003310131102232-0231132121222022-1130320210100022-2200303003020312-0101132210212013-2230300203130030-1013132121201010"></a>

### Direct properties for `https.tls_cert_params.use_mtls`

<a id="canonical-0013220031312131-0232223031233312-1211113323013022-0301321200023000-0323112322030303-3303302231322331-2001302032003231-0020000303030101"></a>

#### `https.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-0213223013112213-3232101122221211-0302300301322233-0031120132110111-3031133310101232-3111303310100122-1123231222120130-3202010133333203): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1201110112013000-3320022111210132-2022311323333201-2210203323233221-1221331122000310-3102300130213213-2121321222333130-1001100110002213): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2300023113211202-1232310031212033-1103320000111301-2023123322032203-0023301230133132-3222032033210003-3122212131020030-0331232002312013): complete subsection reference.

<a id="canonical-3202122212331120-2101031012331230-3020130231223023-1132012101013212-2333201012013333-3330120320222233-0201120123332131-3111213202221010"></a>

<a id="canonical-2220130211031222-2330023231111203-1322210132320012-0213123131233133-1313310122111132-3032031220100331-2120031133120320-2020222101012100"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-1023120331133303-0110123333003303-3223301012320202-0022201331221332-2233231110010223-0003022131301110-1331332021011132-0010020012321330): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-2030032300301111-0230222233121102-1221202032221023-0022320332212033-3312010223022100-2112211333011132-3103013020023032-1101213202000203): complete subsection reference.

<a id="canonical-0213223013112213-3232101122221211-0302300301322233-0031120132110111-3031133310101232-3111303310100122-1123231222120130-3202010133333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-2012322203013010-3010331231321120-2300333231330332-3220230120202321-2233321211203132-2320211021000231-3203033120001300-2011311302311301"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011012022330111-3013311313102302-2000111132020023-2030203312111233-1330122120033323-2013211111031130-3222212033231313-2101122330122003"></a>

### Direct properties for `https.tls_cert_params.use_mtls.crl`

<a id="canonical-2323023203333030-3300320203111331-3101332120001301-1113330122033111-1303132332223132-1202133101110123-1110003112003101-0233033001011232"></a>

#### `https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-1332022311011031-2012031300021312-2200303101003003-2031012221120201-2030322220312103-0331110303302100-1311321020210213-1322003333303310"></a>

<a id="canonical-0120313232223033-3113223310300301-0122301312311201-1101213002012113-1132303221200301-3200221033331113-1102100100020333-2223033030331120"></a>

#### `https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-0130011110210333-1211203010313322-0123233313021320-2212211222231333-3231231301212320-1300133112313332-0332013231101122-1322231330001101"></a>

<a id="canonical-3023003030203021-1330230232030121-3100102311031130-0333203332220021-1212211303022201-0012203230311022-1303320110131230-1202130300110023"></a>

#### `https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-1201110112013000-3320022111210132-2022311323333201-2210203323233221-1221331122000310-3102300130213213-2121321222333130-1001100110002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-2311031021103201-2222320202010023-3120221323230223-2013113130330103-0003202112201030-2110022012130133-1333322121111313-0220202231131312"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300023113211202-1232310031212033-1103320000111301-2023123322032203-0023301230133132-3222032033210003-3122212131020030-0331232002312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0200202013332301-3011333303012332-3110310021230312-0030311033231133-0130122031202321-1221020333030122-3103221010320211-2302322101210131"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121322122313223-1213002112210202-0323103312333012-3212130321201202-2211210130213030-1212200200311113-1202323030001201-2233110330003030"></a>

### Direct properties for `https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3200201120220100-3223330113010000-1312013222111231-0101003002332103-3020011013000010-3001111030332103-3111230000102020-2130212133323022"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-0113001120032211-2031112100021103-3122110213031220-2033200231322322-0213200320012023-2321031233331021-2102231313223001-3203221310112210"></a>

<a id="canonical-0313231211212222-3233101131113300-2133000020003001-3213202003000032-1020331210121030-0331333000002002-0003021310311123-2101031032000320"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1121231212320221-1223130121131133-3320131000311312-0202331210313102-3103320203022131-1311321023032312-2121132332000133-3022113300030203"></a>

<a id="canonical-1320221311330332-1221013100211323-1000012303021121-3222230301010121-2222311030013210-1102131010103030-1312001010022003-3022121021010013"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-1023120331133303-0110123333003303-3223301012320202-0022201331221332-2233231110010223-0003022131301110-1331332021011132-0010020012321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1233213323033011-0102231200002012-1111301021122222-1122210222003210-1222222010232023-1310201213310131-0331023203020220-1221103101031112"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030032300301111-0230222233121102-1221202032221023-0022320332212033-3312010223022100-2112211333011132-3103013020023032-1101213202000203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0312131213031111-2232231320233001-3131132311033233-0222321311133113-1131001222122010-1232231022100332-1321031333120300-2000123300111022"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011023323211321-2030110300210213-2000211030210213-1100003031310300-2003203033323020-3320210303310121-0220020131231031-3303213303300032"></a>

### Direct properties for `https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-1203012023003231-3330333302122121-3123033131030022-1220101313111033-3100322222301210-3320002210023233-1112202223131221-3202231133212330"></a>

#### `https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.tls_parameters

<a id="canonical-2303112103210221-3112000301021210-3001130121323132-1023131031001322-2231111331031201-2200303303211200-0001320332231020-1131213312202302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122130011222301-2222133100100312-1232321012221031-1121300132111231-0023011203230102-1200102030003223-2322122121020120-0032332323330020"></a>

### Direct properties for `https.tls_parameters`

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-0310032130223213-0000030332120330-3331033231032110-0012003112113033-3121211021200000-3321212301033033-0111112121023132-1102003021003133): complete subsection reference.

- [tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230): complete subsection reference.

<a id="canonical-0310032130223213-0000030332120330-3331033231032110-0012003112113033-3121211021200000-3321212301033033-0111112121023132-1102003021003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.no_mtls

<a id="canonical-3112333032131022-0330032203030032-0321300033001222-3112310130230023-2333121000023320-3320220222232010-0022201101303002-3003333332131023"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.tls_certificates

<a id="canonical-2003103320220312-0010201010022112-1203212022022000-1013230000322033-3301303303232123-0011213131220222-0101212332232113-0122310330203310"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212133100310101-2320231220030010-3131130030322002-1223122330031020-3112220221200203-0011032302212100-3010311001032010-2231220033211202"></a>

### Direct properties for `https.tls_parameters.tls_certificates`

<a id="canonical-3302233212023301-2110131231321233-3312333133312012-0311110123023033-1331012303122203-0111101333303302-2222333133223301-3303102223232332"></a>

#### `https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-019.md#canonical-2013303333001103-3011311020010221-1123023321112203-2333031023203133-3321020113223312-0221233102012220-0130102030032221-1231011223320321): complete subsection reference.

<a id="canonical-1101011030321211-1333311312123032-3113323312021121-1010202213030322-1323330223313203-2212221213001312-2003320020000020-2331001021222132"></a>

<a id="canonical-1320200323030202-0212322011233300-2310300223133220-0002030310100321-3130220131113130-2003100121032321-3010321330122310-2032223230200311"></a>

#### `https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103): complete subsection reference.

<a id="canonical-2013303333001103-3011311020010221-1123023321112203-2333031023203133-3321020113223312-0221233102012220-0130102030032221-1231011223320321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2132020131303202-0112110030003012-0313110103312311-2123301033020232-3132033201212000-0223130230013212-0122002033310030-3312321022223231"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310120332111103-1210313020332313-1002213313302201-1232020103200011-1322322111002000-1300130131332332-0011303300100100-3011300001220221"></a>

### Direct properties for `https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-2233030010212333-2000232020112001-3111221333122220-2011101112300103-3133312203222101-1110310232210001-1333300033333120-0202032200330013"></a>

#### `https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3121222102322021-2120213120022200-1323300213021113-1302123231233001-0120011301230031-3101120130202101-1330022222101011-3130213121132303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-1322303213121132-1322202213231111-0011300132033000-0201213301131201-2121001010203001-1232220000311021-1231110000100022-0032333011030133"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131001013032222-1101023113003003-1301130313131211-2021301330013112-2033221303100231-1323200322132023-0310130220302223-1133211011222131"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-2101021213111133-3000323023020322-3313130212120212-0320323223112323-1331130021032013-0032110321320020-3200120311301200-0031111303303012): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-1211202021023200-3331130303131220-3300220332121300-0233010100332222-2012303203323012-0010212212001101-0120013221322220-2101311321110123): complete subsection reference.

<a id="canonical-2101021213111133-3000323023020322-3313130212120212-0320323223112323-1331130021032013-0032110321320020-3200120311301200-0031111303303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2222102202303310-3013113203331333-0330133331202020-0010122023130210-1320231131210132-0101013231230012-3111220322002202-3330111200013100"></a>

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

<a id="canonical-3000013222031032-2332100210011102-2310313211303310-2113300032221120-2202120020120023-0023122232320003-0322113312310132-1320101333011232"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1330321003211302-3233033103213131-2200311310003203-0131033203130001-3233313123112131-1300103031210123-1331330003313311-3332102132213130"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1233033121221113-1210103031333132-3013023320032200-1003201221013310-1220131023031320-0002313122031331-3203233023030311-1112122221211300"></a>

<a id="canonical-2121100133023100-3110321220001003-3022031031013111-1122023001213110-3033002323200021-2002103210133322-0233301300330020-1300120321310313"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-1233123331103103-2001030332101203-3033230023301013-0100232030132100-3233303201210012-0313221112310010-3120000221000212-0302013231131013"></a>

<a id="canonical-2212112303232333-0131231222202031-0032011010311121-0131103223031001-3002002013021320-3033330231311001-2103122113333001-3001000012311122"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-1211202021023200-3331130303131220-3300220332121300-0233010100332222-2012303203323012-0010212212001101-0120013221322220-2101311321110123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3123133212313302-1230231030231300-1003231202333100-0012002313302231-2202303131131232-0133321303231133-1210123030030003-1301333011100022"></a>

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

<a id="canonical-1130003131310033-2003130330020021-3323300012210032-0130102123131002-1032002212332221-2231013331113312-0123022133211011-1010002213320222"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2222310232111002-1001113012230220-1223003300001212-2221303200302023-2033332013332311-3021230101021331-2000013131323133-0000021133021233"></a>

#### `https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2120020221330001-1101321231110023-1012133330133301-0103222123112300-2332103202220001-0203303333333333-3330102233102331-0133320210033303"></a>

<a id="canonical-3021120112102131-1302233323012222-0233001232231011-1323123323023120-0231111110322312-2100133031102120-0321231002220002-2013131201012000"></a>

#### `https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2303201310310112-0220203301130313-3130311303212213-2100213201132333-3131301020111101-0123031313311110-1330330220011222-0222223112132213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.tls_config

<a id="canonical-2213301003123031-0130321302323013-3311012000232203-2230100300311322-2033012301001033-0033121320120122-0120301330010130-0101302103012100"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210312201021122-3101123122330010-2312303002100032-1201013100120010-1103300202312300-0213110220231331-0102112321130021-0313330323120330"></a>

### Direct properties for `https.tls_parameters.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-2231332332000212-1131230233023220-0003301313231321-0020100131230112-3312121012300032-2131021102332211-0300110333110231-3202232032002221): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-1013310333012201-2000320121203322-2303222121322232-3013000302333330-3013132333000123-3030301213012232-2330301001000201-1011322002113213): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-1101201203013130-0213013122203320-3230310031103130-2222130112000203-0312012200013133-1200120013233113-0000122211323131-0113321330233213): complete subsection reference.

<a id="canonical-2231332332000212-1131230233023220-0003301313231321-0020100131230112-3312121012300032-2131021102332211-0300110333110231-3202232032002221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-0210323020112013-2101101000100013-2232013332331201-3010020000103012-1222231230320130-0102201010301131-3012303023011123-2230232323322201"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122321230322321-2123021313020232-0332130102201203-1301023323222213-1231300120003022-0323321030022132-2121022122210122-0023122300130300"></a>

### Direct properties for `https.tls_parameters.tls_config.custom_security`

<a id="canonical-3330312120100202-3302032321032313-2123111232112303-1002033221302200-2221212101010132-2030032310321212-2320233000211221-0032323022332312"></a>

#### `https.tls_parameters.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330311010211003-0033012023003220-3022130201001323-2031132033210212-0200313333101030-2123100133112111-1102221331033130-2323133323112301"></a>

<a id="canonical-1021010133210320-2131012022301011-0102333013301123-0311231120102303-3010003031120132-1003020321033312-1030101023002101-1032200111131113"></a>

#### `https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0113232120230013-2023122230132020-3330112120301021-3110323303332323-3221311013331011-3232113300012321-1002013101113033-0332031120213020"></a>

<a id="canonical-1331303320031232-0211102103222020-1322111322121201-3032330103121222-1303300212302211-2330133012012100-1221113311322001-2232323310230312"></a>

#### `https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013310333012201-2000320121203322-2303222121322232-3013000302333330-3013132333000123-3030301213012232-2330301001000201-1011322002113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.default_security

<a id="canonical-3002000320000211-2203212130321103-1132203001131123-3112131110030012-3100011202002200-0232033212021331-3220103323331113-1033311330303210"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.low_security

<a id="canonical-1101312020213220-3310231023310023-3201323301210000-1023120131201012-1302111113230103-0123211203023122-2122331310132221-3100032322213333"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101201203013130-0213013122203320-3230310031103130-2222130112000203-0312012200013133-1200120013233113-0000122211323131-0113321330233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-2001320201021031-0312310022032331-2310001011010332-3122300323131112-0223322130313023-3131321201021213-2021202111101230-1210302130001202"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.use_mtls

<a id="canonical-3022100120320223-3301233221233120-1330113123121322-0233233013022122-2111132030111133-2120321302200233-3020130302311230-0010000300000202"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030202023210321-1302011312202023-2231333100222220-0021333002231013-2033233120112203-3023033133330101-3232010302230301-2133230003131200"></a>

### Direct properties for `https.tls_parameters.use_mtls`

<a id="canonical-0031002303212110-3011330210123311-3233330331212103-3023333120003131-0112021033111132-2201320121131003-1303120320122221-1101220310203330"></a>

#### `https.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110): complete subsection reference.

<a id="canonical-2010320113301112-1322313133011222-2233233032232102-3110322013212011-3300113133332013-0100213300221332-2321100130213023-3212112123120313"></a>

<a id="canonical-3103033200310011-1033003331331221-3020333022021221-3313123231002321-3101013113200102-1011131123213310-1033110113020100-0132300300330232"></a>

#### `https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333): complete subsection reference.

<a id="canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.crl

<a id="canonical-2331030020221031-1231123300213211-1012322130331331-0101122012120301-2003231113020203-3221230223233303-3332100111121312-3102300203320223"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302122203021023-0010322301021102-3303011001110013-2103011002312131-2210003120033302-2131113310230033-1303101023330123-1112002022301110"></a>

### Direct properties for `https.tls_parameters.use_mtls.crl`

<a id="canonical-2212122201022003-2210020203120233-0132023213233133-1313101202002320-3203031031321323-1130200013021203-1100322100022103-0123123031112023"></a>

#### `https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-2011322131222112-3321132200030332-1312322011110303-3032022232111231-0323213130102233-3130302131213123-3313300012332021-1001112223310022"></a>

<a id="canonical-2303011120122132-3330113223011031-2332333012130030-2221002111023031-0221220113213233-3231000133022101-0011303203011311-3100021223113212"></a>

#### `https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-3332010102322023-1130112031100211-3133131210210121-2032033211200310-3032212313032231-0201030223320303-3003023312331123-0002301120311201"></a>

<a id="canonical-1222312011220331-0130210021200112-1130013213322021-1120123101102213-3332030101131300-3302102001130233-3022201232202230-1011313023201021"></a>

#### `https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-1230002332301133-0313013212112020-0323120312100111-0310012323202031-1112132103230200-1321203313230310-0213113231131302-1231323322211312"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2122333012202002-3210202000122111-2111210002202102-1002320330333210-2200110010013233-2310133101231233-0221320203011313-2221020100310333"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121212130323023-0101122011111203-3321231021310302-0020303121102233-1021001202033022-0120331301101131-2330211002201122-3220320130202313"></a>

### Direct properties for `https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1223133011301133-0120233310020313-0122030003032231-1123032031101223-0100001212120012-1103311333101000-1310020032121123-1331131122303102"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-1211333230221201-2320332332322311-2331010033133022-3122020323223032-0213013030132032-0210122111303321-2120023203210002-2010120011331321"></a>

<a id="canonical-0320210211100310-0111100032003203-0331133020302300-3323012030223033-2230112121030323-2232031030030130-2310002100023013-1301110320321130"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3120200012201003-1112212113021130-3212303133303103-0333031112000223-2002212020300233-0102323200212233-0323212300300113-2310320122310302"></a>

<a id="canonical-1122323132322330-3013112213022010-2321111233122113-3210120223203223-0301311330121031-3022001310010323-2303001332100010-2023322230223011"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2001330313213112-2003121033010131-1233202331001132-1031030001011100-3332012002200003-0310121321203212-1121303033022021-1113231103133003"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2300200302102222-2312213000310030-2331231223103023-2013230302023320-0232001013322012-0121010030132101-2222100220202321-2100231300102313"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101123200330222-1033130332223212-3110001021211222-1213003121333011-1203121131123100-3013300211021023-2020331333032112-1233113110132022"></a>

### Direct properties for `https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-2103202323130012-0131122233030220-0321022231131220-2122013031002230-1220021330011210-2003111113210102-0111003120222223-3102120320233331"></a>

#### `https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- https_auto_cert

<a id="canonical-1223123300112032-0222133020221001-3201331012031023-0133020113303300-1120222301233102-1120232121100332-2220301031031210-0023330100002100"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212233130303102-0211032023003220-1330010201111021-2030111110113301-1120330113022112-1020032101331231-1112030202303213-3111303311000323"></a>

### Direct properties for `https_auto_cert`

<a id="canonical-3223321310121123-2301223021013011-2111300211130232-0122000300102301-3333333133201231-1010132220131130-1303010200233013-3010322133001333"></a>

#### `https_auto_cert.add_hsts` property

Type: `"bool"`. Optional, Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031310103320212-0110301112112033-2022100102210201-0222311200101133-0311111221202023-2132303131333320-3233213120130133-0323322103002022"></a>

<a id="canonical-1001100232130231-1001110203232212-1212101300330122-1121301331231033-1233310320200030-2212221012103010-3323200113132311-1002022122031230"></a>

#### `https_auto_cert.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103): complete subsection reference.

<a id="canonical-0102002212110200-1321022333101111-3313032013110313-0200122112001310-3012223110030220-0112302110111312-0211132011302230-3333313201301130"></a>

<a id="canonical-0111333230002111-2101133003323202-3320031233202301-2133313001213221-0211332112211300-3220122102021323-2000023111002210-1131331202132303"></a>

#### `https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Additional upstream details:

Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
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

- [default_header](resources--http_loadbalancer--reference--group-019.md#canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122): complete subsection reference.

<a id="canonical-0323010101202030-1232230010101110-0301331011022103-2100331221100330-0032022001012220-2222002120303332-1322322020331233-1112132333110203"></a>

<a id="canonical-3203333030333211-2003132223020311-0012330203101012-1002133001002122-0223210201300201-2031032303122031-0311220231032233-1321323100230303"></a>

#### `https_auto_cert.http_redirect` property

Type: `"bool"`. Optional, Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332): complete subsection reference.

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312): complete subsection reference.

<a id="canonical-3102122321001103-0203131322222310-0201221031133023-2212331332021130-3320031012033222-3323112322220220-2232021011310222-1232131121311112"></a>

<a id="canonical-1101201322323321-2300331220002202-3322312331221202-3000011301320220-3120021212122122-0120221201112032-0323211122212330-1012331321212230"></a>

#### `https_auto_cert.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2013121000300130-2202001203102032-1313012203022002-0302230002020312-3221021222033303-2202301323301033-0320111021020130-1123131122103321"></a>

<a id="canonical-1232133303302003-2013103311313102-3311202011100320-2311011220323231-3211223231323203-3122021132100011-2230323232233121-2033220230130301"></a>

#### `https_auto_cert.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0310020220232133-0223222101221100-1230232033110333-0333310311322030-0030221331331110-2200233121213321-0133212311133101-2230303123022223"></a>

<a id="canonical-0113300213233320-1030330110032300-3231102113312232-1103333030330320-0232020010021002-3030320213032111-2323132022102032-1122212002232333"></a>

#### `https_auto_cert.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131): complete subsection reference.

<a id="canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.coalescing_options

<a id="canonical-3001231212332201-0313132220012130-0030022212322000-3112123233103301-0003021133311302-3323021100230320-1022133233231222-0310213323221033"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0120322221033330-0132332223122213-2133022113102222-0233222200330113-0310231213200031-1302001100101323-1200032323031203-3220320202013031"></a>

### Direct properties for `https_auto_cert.coalescing_options`

- [default_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033): complete subsection reference.

<a id="canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1101100331131002-3203323132211211-1313320203333302-1133030220030012-2331011320023013-0313322303300333-3231113222300000-1013220231003331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3200023331202132-2301112331203032-1112102033020122-0211031100203321-3101130023332222-2300210313110212-2111302020301211-1202323210220222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_header

<a id="canonical-2100311123231131-0133311110202100-2332023030121232-0023223011001303-2031232033201002-3302130331132312-3030110102022311-2200021303333233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_loadbalancer

<a id="canonical-3001000020233233-3001321322312021-2322123012222021-0133201022333112-3323333211203103-0330302331102301-1232203033312201-0302322313113110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.disable_path_normalize

<a id="canonical-0333213122002122-2030101001233332-2123333323000121-1110122231213323-3030331031030011-1223332223313103-3230211212032010-1302311321022120"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.enable_path_normalize

<a id="canonical-2223013223201010-3303013330023003-3300321112131202-0222013313100103-1320203101220202-2130020300130020-0313322332022230-2101230131000111"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.http_protocol_options

<a id="canonical-2011113132131102-1110223121331230-3310333210200303-2221103221112122-2231100322130011-2202011222323313-3020013102332103-0013323210332230"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1201211232312221-0313321232012333-1321101010111011-1203230113200033-2002232122233211-2321021130321220-1010202121310210-3210003303021102"></a>

### Direct properties for `https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002): complete subsection reference.

<a id="canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2031312311322022-2323332220121203-1103310231121222-3201301221130333-0101003230132100-2012103000330332-2201203132212101-1121303022130303"></a>

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

<a id="canonical-1302013103010301-3000331031202230-2303132032222032-0210212012210320-1031122101112212-3233032003303323-1230131321331311-3200200020313232"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223): complete subsection reference.

<a id="canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3223012123130213-3232203130132033-3233200120000330-2301100133230320-1123122133303311-2121332123111230-1310121310032223-0120132231012113"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2020010330230120-1310110000020231-0010231003133211-0032103231320200-2001212200313222-1123220211231123-3303121200220020-3122213212303211"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013): complete subsection reference.

<a id="canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2222233231313110-0113330311010200-3323002031102213-3132300131203323-2201023000323030-2333310321111220-0130122301213330-2103003001331332"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2330011023030002-3100012200112020-2211222321100330-2103322013300303-2003000212023231-1211311132121102-1113331102012121-1231112132102123"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1333101202011333-0032333210231203-2110202312123120-1010110033030321-3300330311001321-0130002033221101-2320311333213133-0333102113130001"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2123230010031123-1020303333310032-3131112012111011-2333133011021320-2203123130002302-1312012022032231-1321203033101131-1323131121133302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3122203002332132-0123021301103122-3103030331002102-2300331311323121-2330030312032202-2120020210212101-2133132222031032-1220013010123022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.no_mtls

<a id="canonical-2220332020001020-2302321002301012-2311000132102222-1122221231100012-3312231011131012-2303331110220020-3002122100311031-1211320332103221"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-1001221102303011-2220021100220012-0230032100202011-1301332010210020-2030022302012303-3302231211303103-2211120212331103-2333133230111310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.pass_through

<a id="canonical-3311221012012030-3213230332021221-3212331303133020-0333212210303111-1122021210220031-2131202102133211-0213003033202211-0220210023210101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.tls_config

<a id="canonical-0233120200312332-1211103300220120-2332300211130200-2113122123122012-0231332031100303-1132222200303331-2330333330121112-0033311121331021"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013311133002230-1233330322211320-1012013032323222-2313301320102232-3000101021131210-3102331332322313-2111020210230230-0131012132300222"></a>

### Direct properties for `https_auto_cert.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320): complete subsection reference.

<a id="canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.custom_security

<a id="canonical-1113313033212323-1021103010103233-0322232311101112-3022122122120223-1311212203012101-1101220012302303-1313212022223100-2101310330031212"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310000212323301-3230210312213230-2011211322322332-1221001332030311-2232000032310022-0212112112012220-0331301220221133-0332123122121013"></a>

### Direct properties for `https_auto_cert.tls_config.custom_security`

<a id="canonical-2011223032011232-3122112313013030-2330002032002232-0332023111103022-2300103322120111-2110102321111022-0232020310121312-1102332032220030"></a>

#### `https_auto_cert.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0000312322102231-1030221133131220-0133321201312000-2132002003332112-1001013233213100-2202201031131022-1311202020320021-1222333312301030"></a>

<a id="canonical-1020310330320330-1010210023230002-1310312023201223-1331032301201201-0321033003002223-2032010000101200-0203301221322132-0002310233322332"></a>

#### `https_auto_cert.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0101311222222313-3100120110113331-1313221201011301-2011123033301001-3311331222011023-3012311323321133-1230012122223321-2311222101031312"></a>

<a id="canonical-3303213202012220-1332013112221031-2003021112013323-3123131101110012-0123022012123202-0211211312220132-2202321320231003-2000213010020213"></a>

#### `https_auto_cert.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.default_security

<a id="canonical-2202001103333021-2002220110001221-2202203001012330-3023320211102210-0300012131312102-1113032320203000-0211230312211022-1101212230013110"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.low_security

<a id="canonical-0201122001030021-2100130002313122-2332212021333030-1110221320113022-0032232203231303-3113312031202010-3100323021020020-2213332210201103"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.medium_security

<a id="canonical-0201130300131301-0013202212221303-0131133200011200-2310022310002311-3212332223112012-2331311203120202-1022110313303203-2113110320323011"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.use_mtls

<a id="canonical-2311300312121330-0000303231320033-3211203320333001-0221203223003110-3103112313220101-1230001333223313-3202003200101201-1223332333221203"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322101310220102-2111000222222213-3113123330213121-0001111122300311-2033123330130023-1011203221220122-1313002232332121-0133101300121121"></a>

### Direct properties for `https_auto_cert.use_mtls`

<a id="canonical-0201023200032302-0302323332120111-0002222331301011-3122322300233201-0133311121022333-2303103230031033-1231312132021220-2220303103131101"></a>

#### `https_auto_cert.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303): complete subsection reference.

<a id="canonical-3120331123332031-3333303223310123-1322133112210000-3031331210032330-0120320102000102-0023231210331030-1302010310211210-0030113301111000"></a>

<a id="canonical-3311310303230321-0033000101132113-2002322120133130-1301033002212102-1300113301203132-3313222010231320-2303300332020001-0313132122330000"></a>

#### `https_auto_cert.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110): complete subsection reference.

<a id="canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.crl

<a id="canonical-2130220123003332-3011211033010333-2003321000303101-0013030121103030-2021033201303122-0312101300313202-3201003333232301-2320302030310201"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001031303332221-2130122210020131-2212232230223210-0310210013212303-1302312103033300-2001333100011100-3030203213200222-3003122130021310"></a>

### Direct properties for `https_auto_cert.use_mtls.crl`

<a id="canonical-1022302021023111-2023222221100000-3001102201131211-3100201331321132-0330302133310230-1200002220001001-1213021323110322-3033010021302233"></a>

#### `https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-2101313001311330-0330103313032121-3313323202313310-0233230013312022-3230132322130320-0002011033011132-2023201320002120-3111122313300312"></a>

<a id="canonical-0231032200013121-0033230211101132-3010012021230332-2020010110002002-2113020100232210-2233302210232312-0003222132330010-1223113021120230"></a>

#### `https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-1021332322113212-0331312301312011-0122133221301331-2213222201310321-2232301220111002-1113001332100120-1023313012323120-0230123233133011"></a>

<a id="canonical-0222332333120211-2032133330101112-2312323110032222-1221012302222032-3211213001022311-1031121100300113-1203201230303300-2330112200021131"></a>

#### `https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-0013033311012301-1100112101011321-0300033202102030-3110021030101001-3231333032001011-3112322023223121-0213331013311021-1231222110100223"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-3201313232021021-2033101212131332-0101000120333313-2323310301323211-1131013133031003-2011132131022322-3321302030002301-1312102230102233"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010111133022120-0302231230003323-0133331301301033-3230103101120310-1222012001131210-2303223220230323-2220313013233012-0323103313033032"></a>

### Direct properties for `https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-2230101211212131-3200321203012220-1222322021231300-3023000201023021-3221302322300131-0101320102303020-3200021311313111-0212221030300133"></a>

#### `https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-1232311330133202-2213203011330210-0132030311213101-3033001031023333-0002323013323330-3103332003213022-0200300203333203-1212032213103302"></a>

<a id="canonical-2001011110201013-2212213313300331-0012132323001313-1300030112333001-2333032000301301-1011200030120003-3321312022203021-3231222103232331"></a>

#### `https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3001021131220101-3303123013013031-2110231313220111-1000331101302103-3013213221021201-3033321221132032-1201321301332032-1103221312133301"></a>

<a id="canonical-0012322110022031-0232120313201103-3233202121130300-1201003221301230-0232223233233203-0233033130213021-2121013022221112-2230201210022301"></a>

#### `https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3232020310031110-1031223133130020-3020032100101310-3020231123323112-1113303103032311-0031013310030113-0331310230030122-2022003110102122"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3312303100000321-0213221220310013-3320132300000302-2002022002003013-0232201032331021-1122033120112010-0302200220013131-0213302010332032"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013311331211022-0122313121000221-1232012201200103-2002000012021130-2131012112222212-0131011330311321-2323211010133100-0111133300031322"></a>

### Direct properties for `https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-2303012132330203-2221121002213302-0223201011320131-2112012033300121-0130011223023300-3202101333220332-1221102202203103-2010102020100023"></a>

#### `https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-0122213100023133-3102301113333103-2221320301323033-3302232320022112-0012001332132331-1022120132002212-0201313110302323-3320131033130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- js_challenge

<a id="canonical-1011111331300221-3213330020020313-1033022332302021-1111331213030331-0032003020022230-3000023321331133-2002032233231133-0233011032133100"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210212111030230-1212022231030101-2010302221133200-3031010311103230-3303100223102130-1313311032323231-1213233211210321-3322222130310321"></a>

### Direct properties for `js_challenge`

<a id="canonical-3032132210021332-3300311000232001-0232220012102222-2311003102301333-2102321332033110-0001302132210111-2113300310302130-3022121102323112"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0101103200132210-0120122330012132-3100330021112013-2002303232121031-1131003232003111-3302321333011231-3132101303002322-3201230010233003"></a>

<a id="canonical-1332123223022102-3310323103303022-2331323322102033-3230211210021003-0223112133121212-3213113332203212-3233321120130132-3332102332210322"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2112211211300000-2003223012330332-2013201023130220-1303111122312102-0330012020111010-1102212333200113-2303210213101022-2321120220022002"></a>

<a id="canonical-2221203130302023-0322031111201010-3200312303110001-2210123321120203-0020113011133220-0230233003001330-1310120132312311-2113130203013033"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- jwt_validation

<a id="canonical-3232231033300133-0123002110102230-1310030120201003-2321203312311032-2323123120301112-1300312222020221-1323200030221031-0110022312130022"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("authorization_server",
    "jwks_config")}
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
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031303210200111-2010021320231012-1233202322023100-2111230002121201-1120201011230022-3222301031111121-1120003330233033-2222221123222131"></a>

### Direct properties for `jwt_validation`

- [action](resources--http_loadbalancer--reference--group-019.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333): complete subsection reference.

- [authorization_server](resources--http_loadbalancer--reference--group-019.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310): complete subsection reference.

- [jwks_config](resources--http_loadbalancer--reference--group-019.md#canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221): complete subsection reference.

- [mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030): complete subsection reference.

- [reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222): complete subsection reference.

- [target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221): complete subsection reference.

- [token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120): complete subsection reference.

<a id="canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.action

<a id="canonical-1312230111022013-2110100222223100-1202112333132220-3101322110023303-0210131312132213-2300313002201200-0301133130031330-3301213310111220"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320133001300322-1303100311010220-3312310111322230-0021312332200103-0130133020100031-3311320001112003-1230033303003201-2211202111312313"></a>

### Direct properties for `jwt_validation.action`

- [block](resources--http_loadbalancer--reference--group-019.md#canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-019.md#canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301): complete subsection reference.

<a id="canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-019.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.block

<a id="canonical-1312210302123110-2323210031232032-2012000111011300-0231331012131101-2331220021201230-0300201313220333-1132322000112232-3300210102110320"></a>

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
block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-019.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.report

<a id="canonical-1110333320121232-2011200133213320-2132302111321203-1112111321311301-3123300231013311-1303220332210100-3311033300322002-0222333321303331"></a>

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
report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.authorization_server

<a id="canonical-2312131033133200-1002110221220223-0320033000332210-0103320210201311-1311323022322231-0331232102312121-3033110202310123-1313031201301020"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
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
authorization_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121112323010212-3110200100210030-1103301320022220-1310131030002010-1222313011211223-0303012233323302-0000023332300311-3120100001220330"></a>

### Direct properties for `jwt_validation.authorization_server`

- [authorization_servers](resources--http_loadbalancer--reference--group-019.md#canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122): complete subsection reference.

<a id="canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server.authorization_servers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-019.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-1131313331231012-1112000110030112-1320220301022233-0210023011220120-1332320203021203-0302231132031310-2333332232032113-3123310032000303"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
authorization_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202312003202113-1030120323233321-2331131013112022-2312203231020001-3201230131123333-1220200221212222-3223103223303001-2113202312120202"></a>

### Direct properties for `jwt_validation.authorization_server.authorization_servers`

<a id="canonical-0331312002231320-1212331102222000-0111100333110002-3212231223022122-2223030023232020-1231210032002132-3200220303202113-0032311233023011"></a>

#### `jwt_validation.authorization_server.authorization_servers.name` property

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

<a id="canonical-3300130003020321-2021000110200122-1123021313010321-3320112303112011-3012022300030310-2323323231123211-2011132001320211-2000331312220120"></a>

<a id="canonical-0111210302311002-0003130101220113-1101103031303312-1032211102003310-2012031022322230-2120200003011200-2121232222100002-3221220220202031"></a>

#### `jwt_validation.authorization_server.authorization_servers.namespace` property

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

<a id="canonical-0111133121303211-1201122020121010-3210030222311000-0232021211222023-0333032313333333-1320121200123302-0311102132233033-1121023013120232"></a>

<a id="canonical-3013321110202021-0032200310012130-1132230002203223-0103312303332301-3022022200223331-1113133330113110-3000032302231230-0023311012220213"></a>

#### `jwt_validation.authorization_server.authorization_servers.tenant` property

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

<a id="canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.jwks_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.jwks_config

<a id="canonical-0212201230302231-0210300010032010-3322300020112131-1313032133011021-1121111131300211-0003313032231310-0311011031132020-0203003131022003"></a>

Type: `"object"`. single nested block, Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
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
jwks_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003301120023220-1303231233112102-1323320022333322-0120112322031313-1112321202031223-1202201313220303-2120213300323120-2010333102133300"></a>

### Direct properties for `jwt_validation.jwks_config`

<a id="canonical-2323213101210100-2200301001300110-3111223023032003-1313330323001301-3303131012103233-2030021221202223-1312211030033201-1233302111232130"></a>

#### `jwt_validation.jwks_config.cleartext` property

Type: `"string"`. Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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
