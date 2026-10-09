---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3011111120212220-0130201303013010-1120233303211003-2312201013201301-1101310211112331-0103213132212132-1102111023303323-2200013220312000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030000001231131-1113000002213120-1320230101120302-2303120012031123-2102021110213032-0033121101013113-3131331102330122-3230201331111330"></a>

### Direct properties for `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-0000200030102031-0302100111210302-2323113211322313-2133230131002012-0312030000211033-1300200221332213-3313210133020021-3201321213311331): complete subsection reference.

<a id="canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3022103032321211-0101203203202213-1212322322013033-0311313103002331-3322100203131032-1201302203002202-0111012332331120-2233233302033312"></a>

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

<a id="canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2002111002031003-1221032101222103-0011310033300301-0331121222133213-0311100013120003-2023311303203001-2101203020310203-1002202210203132"></a>

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
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [blindfold](resources--http_loadbalancer--reference--group-019.md#canonical-1332013021233133-3201103111130120-3012013132231032-1331322321101020-3212020112321300-2203103012201001-3302311031303331-2011002000201120): complete subsection reference.

<a id="canonical-3302233212023301-2110131231321233-3312333133312012-0311110123023033-1331012303122203-0111101333303302-2222333133223301-3303102223232332"></a>

<a id="canonical-1320200323030202-0212322011233300-2310300223133220-0002030310100321-3130220131113130-2003100121032321-3010321330122310-2032223230200311"></a>

#### `https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0101323223303113-1221013001102000-3202102300311203-3121230112302220-3212130010200333-0113112110230303-1212312013333032-0211110120020203"></a>

#### `https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103): complete subsection reference.

<a id="canonical-1332013021233133-3201103111130120-3012013132231032-1331322321101020-3212020112321300-2203103012201001-3302311031303331-2011002000201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.blindfold

<a id="canonical-1113230320003210-2122110121111133-1011012211311130-0112021122322332-1300221032321000-0212301230022301-3113223330123000-3223301321002113"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3022033110333030-1330223312332302-0222201321011333-1002302001111102-0300213001223221-1130120012011120-2101113301032123-1022003110300302"></a>

### Direct properties for `https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-2231001111311012-1210202030102202-2322331130312130-2133023301000220-2103012232211000-0100003013202002-1133313021122012-0123233333330210"></a>

#### `https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-3223333232321100-0311201310101121-1033121033133013-2212000112133132-1023131012001123-3002021000120110-3000303310133010-3121102233313212"></a>

<a id="canonical-1131003303010212-2330132311013233-2122330321000012-1232211313033202-0002113313033321-0223220030323221-3020233221111131-1211033100132031"></a>

#### `https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1320311211033301-2031031212110103-2110013311332001-0330213202002313-3122113010100130-0230323310200003-0322222022213311-1122011003203223"></a>

<a id="canonical-1222101102013221-3002311121110321-0033032011202033-1210322101023320-3023133210322322-3031213222033133-2231213120033101-3201010302222303"></a>

#### `https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0010233121012201-3000011121100020-3302002031333213-0101022010102110-1323310232201331-1330120123132323-0331130332022310-1003301100201202"></a>

<a id="canonical-1031001230011001-1030133103122102-1330233323222232-1200210033101330-3113032221213101-2113323111131031-1031203231132213-3201233220033102"></a>

#### `https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0313203133011131-1310032231011203-0320103230033321-0332322310021310-3032103332030033-2020221302003213-1221332302003220-3100113031222311"></a>

<a id="canonical-2210310100231311-0302302130320123-2030311211223232-0202202330213302-3021200202310001-2111202030013332-2312120002033031-3332220321313130"></a>

#### `https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-2333002100000131-2033220121321123-3300022223131210-0301120001210322-0133303310210231-0001231220333120-1031301102132201-1123302312031310"></a>

<a id="canonical-2133203012310011-2123203110122013-2233100233222230-0201310130100202-0331033131331111-0321133313131012-1012233113213220-0303130030333311"></a>

#### `https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0132300001313210-2010330021120130-2130012001311231-0120131031332000-0233220322000231-1202000221210131-1302000212012203-2013302230011110"></a>

<a id="canonical-3300033212221031-0303110232013022-1313002300332020-1122020212123300-0011302332001211-1230233331130023-2230112100021301-1130203231300011"></a>

#### `https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3031200031032223-3203310033201200-2303133302330321-1032030011222230-3232310120211200-2332020110101110-3101223022211212-1332333322023330"></a>

<a id="canonical-3012100002100200-2330220020033112-0201002013030130-2031011303230302-2230131100332033-0110002210223322-3103300211213012-0000231031300331"></a>

#### `https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-3203031112020102-2331331221322032-3301131301301022-1012333113031332-3303033003001003-2021220000020111-1320133232222003-0321311200202021"></a>

<a id="canonical-1211031201222223-1000102221313313-3201332322200131-0210233322301132-1132312123211312-0010322032003013-0032003310313231-1212023003012023"></a>

#### `https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-2322321302211310-1312231201212301-3111020312311212-2132331213300323-2001110112313300-0131332132232012-2201123030030212-2313200312103203"></a>

<a id="canonical-3032333133122232-2231132013223220-1200102003303120-1203112310101120-1101203013020102-1320023130020223-2203003030110133-0310333332021120"></a>

#### `https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-0302223102020011-1021233112210332-3031000233000020-2001123300220110-1210332312231100-1003133233231231-1200113201211100-0213012200000010"></a>

<a id="canonical-1110113323113023-2223213201210020-2313111000130011-1312302111013233-2300310033020011-0212333010130311-2211023321231100-2302311220222310"></a>

#### `https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3023332213022111-3311000133020001-3033133101023133-0010331122322323-3312131200231233-3300230223301210-1001312333003313-3331121033200033"></a>

<a id="canonical-1221202201312321-3130200220123103-0302210030311202-1230001121131110-2200302323330010-3213132012303110-0210323103120120-3223313113320031"></a>

#### `https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2000011212130131-0232130112131231-1233221100311120-1310021030121222-1021020101211001-2120132111211123-3020022311231200-3222301221130302"></a>

<a id="canonical-2210213110301210-1300023010123113-2321033110301300-1022323312123103-2301211030121303-2032122020303201-0330112330332302-1211322331333303"></a>

#### `https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-1000003331331331-1001132010111322-2003212132302302-0212030222330333-1130300112303030-2301210222111211-1000010330100131-1022011111101023"></a>

<a id="canonical-2330103231120210-1203221311012320-1300100312110310-3001121300023112-0131231030200210-1100020031120213-3220110221000330-3021111222331223"></a>

#### `https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1312001212323012-0332002102323132-3123002002220213-1100013001210210-3033221010301020-2320212220012132-1000100233202220-3210100103223112"></a>

<a id="canonical-0300130103133220-3222222012202330-3220203031230332-1312132221020233-3210002032011301-3332230223103221-0123122320232333-0331130230001323"></a>

#### `https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-2300200110111130-0020200010121220-2110312013110302-3101312112323110-3330232033202230-2320123021020223-2030030120211230-0102302222213331"></a>

<a id="canonical-0321110231011230-3331132033212021-0201333223220000-0200232021323223-3113213310030100-0112212032102122-2000212300312031-0021031132130112"></a>

#### `https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-1123100103122300-2322202213111032-1033320202230123-0003020303211020-2022300030323133-1103022310023221-2022221121233311-1230022332031122"></a>

<a id="canonical-1031012000301313-2233101122300112-2023131221332230-3122222221210201-2212211300320221-0032133133120101-2010003021010000-0201022103020012"></a>

#### `https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-0001001322111231-2320010022121231-1101312013232323-1202311330231220-1313230013023110-3230123213021332-2033113302132312-3331113220003130"></a>

<a id="canonical-3012102121032210-1332310210202201-2122300201110313-1013123013221110-2031302002033213-0323113122323123-3032013003132122-1122101123102111"></a>

#### `https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1323333133021001-2130333030103032-0012110222012001-3332101121231210-2112300133201031-0112131210311202-2231123303333011-2230023021210320"></a>

<a id="canonical-1220133123013000-1200023303123011-2202032332320111-0320133222311031-2103313320323310-0201102211023203-0102102000221331-2202121320303012"></a>

#### `https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1233033121221113-1210103031333132-3013023320032200-1003201221013310-1220131023031320-0002313122031331-3203233023030311-1112122221211300"></a>

<a id="canonical-2121100133023100-3110321220001003-3022031031013111-1122023001213110-3033002323200021-2002103210133322-0233301300330020-1300120321310313"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

Receipt-pinned upstream constraints:

```json
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

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [crl](resources--http_loadbalancer--reference--group-020.md#canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-020.md#canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-020.md#canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110): complete subsection reference.

<a id="canonical-2010320113301112-1322313133011222-2233233032232102-3110322013212011-3300113133332013-0100213300221332-2321100130213023-3212112123120313"></a>

<a id="canonical-3103033200310011-1033003331331221-3020333022021221-3313123231002321-3101013113200102-1011131123213310-1033110113020100-0132300300330232"></a>

#### `https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-020.md#canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-020.md#canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333): complete subsection reference.
