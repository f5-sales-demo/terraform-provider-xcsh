---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1023311022121202-2100322213130030-0000113020103113-3123001221123122-3013211203111212-3212211322322301-0023203201122112-1011021231131202"></a>

## `protected_cookies.name` property

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200): complete subsection reference.

<a id="canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_httponly

<a id="canonical-3211233200132230-1320201203133210-3222121231102200-3101310200102331-1131033211103203-2210123131100133-0230202031023310-2213013200001220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_secure

<a id="canonical-0020002302213013-0321120233322011-2303112122210301-0220003000230021-0212331113021030-2020021121211131-3020302100222213-3111320232030233"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.disable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.disable_tampering_protection

<a id="canonical-1311001213222122-0311202222022223-0120102331233220-3130333230233133-1110013230032203-3223310200332130-2130333212112111-3111321302113321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.enable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.enable_tampering_protection

<a id="canonical-0002202213011112-0311010233013220-1302123011033320-0012132113121123-2210221023223110-2311311303013232-2010300133122211-3032030323010300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_httponly

<a id="canonical-2021211012122313-0031031302100232-0001013320130133-0131033130312033-3231021213132201-1230300332232332-0233220101310323-3103110023013223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_max_age

<a id="canonical-3102101331001031-1020221120320220-3023120133211031-3232311303122110-2220012122230101-1131002101220123-1201320113013011-3003112300031121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_samesite

<a id="canonical-1003303130031103-3220313222211303-1313122113333312-0130130022102001-0312331133011002-2011220003023222-2033310012230110-3132322330103132"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_secure

<a id="canonical-3021010300301100-3321220031113313-3033233222300202-0223123021330123-1120113120322110-0132203013000112-1011020203231313-0203311310020133"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_lax

<a id="canonical-2231000213132033-1031012230003310-1001222221110022-3000311232220212-2310010100220213-3022011033303122-3020122001101120-0322321103313103"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_none

<a id="canonical-3012022311323223-2331300201331233-0103310031123132-0031221331122013-0203200322000332-2033113110210131-1111332221232013-1130113033101123"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_strict

<a id="canonical-2121330213211000-2200123230023111-1120021233023311-2222101302323212-2122012322123302-0010323322232231-1011201201130112-0332321022200302"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312120313213323-2103003010102231-1132113202300010-0103032212233002-0220031323220302-3231002030321130-0010020301002111-1312011221320111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `random` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- random

<a id="canonical-1021313003120223-3101211023223000-1333330100323102-3033202113022313-2111113220003123-1112002011233323-1312030112001012-1111020111100321"></a>

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
random = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- rate_limit

<a id="canonical-1212231232113102-0000210123322230-2321331212231011-1223210321030302-2020032100331231-0213102112123132-3222021323133333-1010333102020321"></a>

Type: `"object"`. single nested block, Optional.

Load-balancer-wide per-client rate limiting. The counter applies across every path; use
api\_rate\_limit rules when only selected paths such as /login should be limited.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("no_policies",
    "policies")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321223131033012-2032331230200132-1302111022313203-1000133020113002-2231301330103221-0222223212320201-2322012230101001-3002000100302201"></a>

### Direct properties for `rate_limit`

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002): complete subsection reference.

- [no_policies](resources--http_loadbalancer--reference--group-024.md#canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223): complete subsection reference.

- [policies](resources--http_loadbalancer--reference--group-024.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101): complete subsection reference.

- [rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031): complete subsection reference.

<a id="canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.custom_ip_allowed_list

<a id="canonical-3023030301000222-2020202210131000-2230300203230310-2010122011332100-0212133223233100-0322001102312102-3132333221333020-0012303220013220"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132300111322010-0211021003112333-1101000211023210-3133030310113020-3232110100200221-2213032332133220-0010223220111031-2313233311030332"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-024.md#canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033): complete subsection reference.

<a id="canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2000123031021033-3223132012102132-3231013132332332-0130332221003321-1122312113200300-0023001323113301-0133122032332333-1122230203103312"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133301330101023-0100300220130200-0200013020101221-1300013131303220-1131311221223023-1300210323030321-2133212222122201-0330012120013303"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-0130033032121322-2110201010312330-2031012013112332-3203310322133323-2320212210211101-3302112112033131-2331123100110300-3323113302320110"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3021320010033301-1203300230030301-3310320321130000-1232331113100101-3321311201212001-1232202002000323-2323223222211221-1321201300211120"></a>

<a id="canonical-3101023303220002-3212131000301113-0001023033100322-0131310233111010-0303322323102012-3222213132033323-2002212031330211-2232030222110013"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1230103032221011-1112302112121002-1032213222300133-3121303132302231-0230313112322331-3333332021221320-2110331131320322-2030203223301020"></a>

<a id="canonical-3230021331331230-3331012101002212-2011100000223331-0031131132300111-2202221030103221-0100003033033301-0213203231201331-0203011011321320"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.ip_allowed_list

<a id="canonical-3103033003311002-1200130211131302-2213201303103331-3031320123113330-2021311303122332-1231311010022303-3230112212230320-1002133121102033"></a>

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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232211132101301-1030100301020331-3233313300013210-2213300100120033-0211101203311312-1021222200303320-1031132103313101-3113023301212200"></a>

### Direct properties for `rate_limit.ip_allowed_list`

<a id="canonical-2131112302013213-0233210323130101-3310001232032331-3003322021222112-1333002311223101-0130300103011221-2200312331110100-2132302231213031"></a>

#### `rate_limit.ip_allowed_list.prefixes` property

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

<a id="canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_ip_allowed_list

<a id="canonical-0122110212212020-1100201100100210-0122123233122013-0113231212102312-3022132010230013-2210233332010021-2121231233211120-2201203212303220"></a>

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
no_ip_allowed_list = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_policies

<a id="canonical-2311111332132033-3203120133203120-3223003203201302-0303213122001132-0320202211030100-0110123003333110-1313212021323313-2201333130013200"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.policies

<a id="canonical-1311212123020300-1202120201123332-1120101331132123-1310031310211112-3323003212013030-3102221103030333-1321333122323331-2122022321132103"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101003312111121-0030323102020111-2223210103110010-2032320122322332-1013021110103233-1030221133320233-1001220333031121-0221222313111012"></a>

### Direct properties for `rate_limit.policies`

- [policies](resources--http_loadbalancer--reference--group-024.md#canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302): complete subsection reference.

<a id="canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.policies](resources--http_loadbalancer--reference--group-024.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101)
- rate_limit.policies.policies

<a id="canonical-1132311333313033-3331110210011121-0223123000200333-3131030320223011-0133301132213021-2311122120312201-2033301003302032-2133001010133022"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303122020201103-0123320031123230-3302222022311100-3213131002331312-1322223033101313-1120030203222033-0232120120023100-2232123221113032"></a>

### Direct properties for `rate_limit.policies.policies`

<a id="canonical-1013132032310330-2010231113122122-1201110122301100-1232130121033210-3122233031300013-1033112103031230-0320233120203001-1211002012202131"></a>

#### `rate_limit.policies.policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1033103333211331-3112201311023132-0210310112011103-2030202100012021-1023201033300231-3203213012210211-3230211221302121-3013023022131300"></a>

<a id="canonical-3130002201210121-0201200011033220-0100203302213103-1211303202001300-1220211122303201-1312313330311211-1123030110000203-0220332233033331"></a>

#### `rate_limit.policies.policies.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0001232312122333-0010233211302322-1132102320310121-3203011231120012-1003132111011220-0211112001302103-3020030111031212-1113020211033331"></a>

<a id="canonical-0123321323002021-2010033221301223-0323213311132313-1011122102201230-3311031020233322-2012201112123130-2313023230323301-1233020233000001"></a>

#### `rate_limit.policies.policies.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.rate_limiter

<a id="canonical-1103110010233130-2233133213021322-0202121221220302-2233033212231103-1112302201122210-0211200201021331-3210320212303013-3222333310221232"></a>

Type: `"object"`. single nested block, Optional.

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("total_number"),
  validators.ConflictingObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingObjectAttributes("leaky_bucket",
    "token_bucket")}
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
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122333212200233-1201111333110123-3133302130013032-0101303131130320-0301221113333230-1311113033020000-3030323202101020-3223211020313113"></a>

### Direct properties for `rate_limit.rate_limiter`

- [action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333): complete subsection reference.

<a id="canonical-2000021212311200-1321200012003130-1102300000323302-3313133120001200-0103013203021230-2320202103212221-0110200223122210-2232132310121300"></a>

<a id="canonical-0232210112202310-1211203313030203-1303310232231110-1203311111301313-1233320100222020-2202233332220311-2013112201103020-3020230322202031"></a>

#### `rate_limit.rate_limiter.burst_multiplier` property

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](resources--http_loadbalancer--reference--group-024.md#canonical-3232033201121333-3023321102300202-0203022310220000-3131212232330322-3113303131212013-0200123021011021-1233230122113213-3103023100232011): complete subsection reference.

- [leaky_bucket](resources--http_loadbalancer--reference--group-024.md#canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320): complete subsection reference.

<a id="canonical-0210223003020100-1021313110031303-3123303211222220-1211332132003220-3113111030021210-0120223103210311-2201002003013033-1033232302023021"></a>

<a id="canonical-0111222112300022-1000301010022023-3200021102231302-3211100213222303-3000031001130022-3202203113011111-0132031032031322-1133033011111121"></a>

#### `rate_limit.rate_limiter.period_multiplier` property

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](resources--http_loadbalancer--reference--group-024.md#canonical-0303221200330202-3123321021000110-2332300002133112-2220222003033013-2013103231112210-3132200122112030-0221300021323023-2013233123003032): complete subsection reference.

<a id="canonical-0112031132301223-3130323311131010-3122232123200131-3213003010331131-3213220202011223-3130210231121332-3112123001103201-1113232031202102"></a>

<a id="canonical-2333103312100110-1130302212221221-2033332132023121-3121313102103211-1202203103200113-0121223023210132-2230303000013012-1111213020232133"></a>

#### `rate_limit.rate_limiter.total_number` property

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2132010130102001-2232231331221223-0021200330212032-2001320020100323-2312113233211103-0000013320231002-0322302012021003-2022022102011331"></a>

<a id="canonical-3231120210311213-0023212231011322-0323313123123313-0103321010221203-1001010131110332-0021211233222330-1321133101102021-3211212313231201"></a>

#### `rate_limit.rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.action_block

<a id="canonical-0213300333112130-3323221222133220-2311120201200003-1210111132202002-2122332221303211-0221233220200102-1322213132231300-2010310331100321"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323233230232103-1013000311321002-0220201022212222-1003231331133232-1123202201222102-1022203001032212-2311322203013221-3003300222333331"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](resources--http_loadbalancer--reference--group-024.md#canonical-2121310033333031-0103220231231331-1032030013313221-3122010023000002-1312012020333332-3210031232232311-3022101233033212-2213012122220320): complete subsection reference.

- [minutes](resources--http_loadbalancer--reference--group-024.md#canonical-1123220322033011-2022131300031302-0331123002023223-0013333220111023-2103032022232203-1331212310031113-3222233201131302-1132332303222222): complete subsection reference.

- [seconds](resources--http_loadbalancer--reference--group-024.md#canonical-1100321133312232-1010213301300133-1221233020303230-0320102001333003-3133303321011103-3133132232103001-0303113322032210-0221022013233021): complete subsection reference.

<a id="canonical-2121310033333031-0103220231231331-1032030013313221-3122010023000002-1312012020333332-3210031232232311-3022101233033212-2213012122220320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-1212121023112330-3230313123332212-2033332012333202-2200111302031332-3232122222121330-2200323100032002-1231233131220101-1223221211301331"></a>

Type: `"object"`. single nested block, Optional.

Hours. Input Duration Hours.

Receipt-pinned upstream constraints:

```json
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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103213001130112-3002133223311123-0101200303322210-2233002313030101-1101220010120323-0010223202011113-0030003113303110-1030321120210033"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-0133313110003300-2131003100120020-1310103120111211-0200010220002100-0231120322013233-0131300200211130-0311211312323100-1001023233313333"></a>

#### `rate_limit.rate_limiter.action_block.hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-1123220322033011-2022131300031302-0331123002023223-0013333220111023-2103032022232203-1331212310031113-3222233201131302-1132332303222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-2203203013223101-2200113021201011-1100330032330331-0103212031203031-2232031111002220-1110121103321222-1131312301030131-3313132313113021"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

Receipt-pinned upstream constraints:

```json
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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303220032301032-3333133000213011-0102220231223113-3031331312101321-0331112202012123-1012202313320133-2320333103331331-0130332022132211"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-0101133120101020-1322213102101030-2303033022031013-3102021121101130-2230330302113210-3301003023203132-0230111310213212-2310230212012133"></a>

#### `rate_limit.rate_limiter.action_block.minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-1100321133312232-1010213301300133-1221233020303230-0320102001333003-3133303321011103-3133132232103001-0303113322032210-0221022013233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-2102331331230211-3312013003332002-0332203201212003-2310331130213023-1121022213321120-3223001011213330-0331312030203110-3032321301032322"></a>

Type: `"object"`. single nested block, Optional.

Seconds. Input Duration Seconds.

Receipt-pinned upstream constraints:

```json
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
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102012013100311-2103313203303011-0212002033022233-0033203223320022-3000121332032202-0101101000132100-0230210300311131-1312222133010011"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-2033332002002300-2312210033323220-1311320033133103-2012022131322322-3303032331103121-2230301133121322-2312302012033331-0310030120210221"></a>

#### `rate_limit.rate_limiter.action_block.seconds.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-3232033201121333-3023321102300202-0203022310220000-3131212232330322-3113303131212013-0200123021011021-1233230122113213-3103023100232011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.disabled

<a id="canonical-0010013332033122-0111110320233332-2112113320311311-1230012102202313-0201013312020323-3030120221312133-3120300303002213-2113132201032112"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.leaky_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-2131030311121003-1030133220332212-1011212013002131-3231113221020320-2321223231122300-3301330201310330-0202201110222010-3123222021301022"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

Receipt-pinned upstream constraints:

```json
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
leaky_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303221200330202-3123321021000110-2332300002133112-2220222003033013-2013103231112210-3132200122112030-0221300021323023-2013233123003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.token_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-0100203310120213-3200101121020211-2202323212003001-1011323011030332-1100301103012221-0112021202023321-3112223113023112-1111200023011110"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

Receipt-pinned upstream constraints:

```json
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
token_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- ring_hash

<a id="canonical-3013210122112121-0122221322331320-0013232112123331-2012132320331222-1132231302310032-2101233212003331-1003333001210203-3312200031323332"></a>

Type: `"object"`. single nested block, Optional.

Hash Policy List. List of hash policy rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_policy")}
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
ring_hash {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302212102131332-2013333122123320-0332022333133022-3311100202303201-1001113302300230-2322333313321133-0233300311131302-1003021301303210"></a>

### Direct properties for `ring_hash`

- [hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313): complete subsection reference.

<a id="canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- ring_hash.hash_policy

<a id="canonical-1001233103230101-1000300000232320-2033102230313003-1310230311210113-1012230230203331-3223231202203232-0011132203210031-0130133230011322"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "header_name"),
  validators.ConflictingListObjectAttributes("cookie",
    "source_ip"),
  validators.ConflictingListObjectAttributes("header_name",
    "source_ip")}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223223003010302-0230300113113120-0011000323022103-3213302113210330-1313302022021321-0211203100323110-0233032313220301-3100022003013133"></a>

### Direct properties for `ring_hash.hash_policy`

- [cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022): complete subsection reference.

<a id="canonical-3130221302202133-2230001003300020-0010021133302231-1232033202332221-1010101011032233-3013000312011331-0301123102131101-0221313130020111"></a>

<a id="canonical-0113203213122220-0010200033110212-1010212120100230-3221032333233012-0000000200202220-0210303302221111-0010332231203232-3012112021002231"></a>

#### `ring_hash.hash_policy.header_name` property

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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

<a id="canonical-3022230202300222-1103201130121133-2020232311113203-0010032322013233-2231112201012302-2020302313220222-3122111300011000-0033301111001310"></a>

<a id="canonical-3213233031012232-2102220323312001-3212332120223031-2200323312131331-1303101131100323-1323013020012100-1132222132322213-1213130211032333"></a>

#### `ring_hash.hash_policy.source_ip` property

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100100313301221-2003203103011131-0210322222110321-3310323303021000-0013331320110102-3021311233100030-3213102021100013-1211321323300001"></a>

<a id="canonical-0112010311012000-2302310021211301-0232012103313010-0322213033302202-0112321212010002-1122022202030023-3300101230032012-1000301013130213"></a>

#### `ring_hash.hash_policy.terminal` property

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- ring_hash.hash_policy.cookie

<a id="canonical-1123221300031301-0233102300131322-0231223220233120-3303232200331313-0212311210210200-3311321223022320-0020111333303303-1231013112310213"></a>

Type: `"object"`. single nested block, Optional.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
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
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101123332113122-2333210302230022-0321130102231013-2103030331101000-1121113031002210-3213031120132201-2301202031223233-3302011331031201"></a>

### Direct properties for `ring_hash.hash_policy.cookie`

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100): complete subsection reference.

<a id="canonical-2112003132221332-1102011332030131-3301333332120300-3001301302323033-1211010120332021-2323323021010222-3023033010103221-3222112232231223"></a>

<a id="canonical-1313000002211333-3032122211233312-2112223020120322-2100331102013131-3122231110012232-2013100230111220-2020013323333023-0012332030212322"></a>

#### `ring_hash.hash_policy.cookie.name` property

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

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

<a id="canonical-0312330232332012-0031000233000213-1120331101023313-2302110201213031-1313012103322132-3201222312320331-3012323110133301-1333332032130123"></a>

<a id="canonical-1200322331310201-0213111302331000-1123030330221101-0310222012032113-2032033121022012-3212200011133020-1031111023211120-0033201021121201"></a>

#### `ring_hash.hash_policy.cookie.path` property

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301): complete subsection reference.

<a id="canonical-1132130012331100-0310322232232230-2022212013030023-1232113320022300-0100303231132003-1322201301120203-1201022110222211-0023001001103103"></a>

<a id="canonical-0323003211220232-1330033011112133-2133121310131020-2102000331023031-3221021231132230-0023021202300300-0131300301323102-1102210220210221"></a>

#### `ring_hash.hash_policy.cookie.ttl` property

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_httponly

<a id="canonical-3133310313000111-2031021211113022-2333121132001200-0002000320113033-1032321122023211-1203222021320021-3030013320002321-0023332000013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_secure

<a id="canonical-2223202333002113-3121121013230322-3100033102002300-0303013132133101-2002130031223011-3223113000330122-2330300202320002-3302130321002233"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="canonical-0102320010131233-1133121101111033-2311002213233001-1103001232310101-0220101200020010-0220321110212200-3301121111230021-2323122121003102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_samesite

<a id="canonical-3331020003223201-0012120213232031-1131223211002332-0112201012023132-0310021003100100-3113013012100203-0132032311221020-2130322123013200"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_secure

<a id="canonical-0333113130311222-0210223021200112-1213131031223333-1020320212103020-2133133332201312-1032021330020310-3322100020311310-1223302032102331"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_lax

<a id="canonical-3322132213203201-1231030122102312-3313323023333302-1103130202122202-2321000011313322-1023332010301112-0003012211311132-2021101230311332"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-3202023323212002-3032033213100112-1123312131202001-2131312331303001-0113230000100210-3001210211011302-0111113312301003-2322003301000330"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-2221232320232210-3102000210122201-0013322231302022-1213121230230323-0210210232213131-3323121133012011-3022332100130213-3032122111300113"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120201302000120-1101032023332332-1203301030103211-2301112131023100-3323333313021213-2102103120123013-3211232030113002-1310102202033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `round_robin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- round_robin

<a id="canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
round_robin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- routes

<a id="canonical-1022112302210201-1010011312111323-0033100111313112-0212233102202102-1021022230231123-1121103333323120-1231212033301110-1123111133332122"></a>

Type: `"object"`. list nested block, Optional.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("route_state_disabled",
    "route_state_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302333330103332-3000231123012200-1223313032013033-3110312130233300-3201000202121113-1221011200232010-0313220022331010-2221323130111313"></a>

### Direct properties for `routes`

- [custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130): complete subsection reference.

- [direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311): complete subsection reference.

- [redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223): complete subsection reference.

- [route_state_disabled](resources--http_loadbalancer--reference--group-025.md#canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303): complete subsection reference.

- [route_state_enabled](resources--http_loadbalancer--reference--group-025.md#canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203): complete subsection reference.

- [simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232): complete subsection reference.

<a id="canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.custom_route_object

<a id="canonical-2132120132332033-0111230301202203-1332310120131330-3101230211330320-1302310110310033-2231032202200231-0302213013300201-0020203132023302"></a>

Type: `"object"`. single nested block, Optional.

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310333203312222-2032203322200333-3103123323230211-1320102023312023-0013322211322022-0302203320221221-1122002313223023-1000303311113320"></a>

### Direct properties for `routes.custom_route_object`

- [caching_disable](resources--http_loadbalancer--reference--group-024.md#canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-024.md#canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110): complete subsection reference.

- [route_ref](resources--http_loadbalancer--reference--group-024.md#canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223): complete subsection reference.

<a id="canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_disable

<a id="canonical-1300132233003321-2310012321003102-3313000230313012-0113022222000330-3212023132023010-0023031201321130-3322202333303202-1100023203132133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_inherit

<a id="canonical-0023203021031101-3102012330230032-2322300303011202-1102031110221332-2122231020331130-0013003023132321-3233331212131000-3111121233333103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.route_ref

<a id="canonical-1123021100312002-1020112010103001-2313033300010013-1303122122012223-0322203113312222-3001003121332000-0101330122123222-0033201313203300"></a>

Type: `"object"`. single nested block, Optional.

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102303303101220-2220313232201222-2311212111013010-1023203313233111-0333303322101003-3011203321211000-0022233012120313-3313232320032320"></a>

### Direct properties for `routes.custom_route_object.route_ref`

<a id="canonical-2323000200333123-0132121222212333-2210330112033001-0112112303232223-1213013021101003-1213132222113321-2330222001330012-2320213231123322"></a>

#### `routes.custom_route_object.route_ref.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2201320121001121-0103233133130321-2101212203012032-2033101021032220-1230300223120133-1111112032301030-3312332122301100-0220322033203122"></a>

<a id="canonical-1311321212322103-3101311301011101-2031320300213121-3003320022201310-1210130033302301-3202212201010300-2321101213220000-1222103300233331"></a>

#### `routes.custom_route_object.route_ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0221030111013133-1101230221300312-1320131302031110-3310103212313113-0130122232133121-3313211110000220-0120311133031233-3233210101331211"></a>

<a id="canonical-3103132333232223-0210101111301211-2313102222332013-0112132211230132-3103123031231303-2321313013123321-2002130113113013-0022123222000133"></a>

#### `routes.custom_route_object.route_ref.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.direct_response_route

<a id="canonical-2212332230001201-2200303013102020-2331330332220323-1212131312310033-0303100201123223-3332120313321120-3323111102000031-3120322323221310"></a>

Type: `"object"`. single nested block, Optional.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
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
direct_response_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231203213112001-1101032030331021-1020030011212321-1300000102330231-3010001302223211-0111013003103303-2211200202013321-3211210323332211"></a>

### Direct properties for `routes.direct_response_route`

- [headers](resources--http_loadbalancer--reference--group-024.md#canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133): complete subsection reference.

<a id="canonical-1122320301230101-1201300111222321-3200131200000200-0011323233232012-2001002210321013-1223010111113202-3302221330003001-2012320102223311"></a>

<a id="canonical-3303333213133003-2121001302000121-3330230011310322-1130010131201222-3120003332332031-0100023121321202-2320310321322300-2033202022113323"></a>

#### `routes.direct_response_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-025.md#canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233): complete subsection reference.

- [route_direct_response](resources--http_loadbalancer--reference--group-025.md#canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320): complete subsection reference.

<a id="canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.headers

<a id="canonical-0012332300232233-2300123121101322-0230222021131133-1331332110303131-2000331133321110-0120322233122301-0033120130031021-0332222011302133"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013322223213310-2021313020321210-2012032031210100-0113001031130131-3311010233312212-3303322213331130-3312313030321202-1311100301303133"></a>

### Direct properties for `routes.direct_response_route.headers`

<a id="canonical-2013030023011233-0003132110221320-0233300111103123-2213132100123213-0303220033321100-0322200212300331-1133002002203333-1202033030221101"></a>

#### `routes.direct_response_route.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0030210001022320-2130302220112322-1133000102110022-2231303320003321-3132310003201332-0023001033223110-2331233332121103-0303221031322323"></a>
