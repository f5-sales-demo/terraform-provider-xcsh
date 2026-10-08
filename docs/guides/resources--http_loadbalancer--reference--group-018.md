---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2012231012033323-2313120330300030-3113220001001310-2310312201210333-0330011223232101-2212201331321002-1003002130002031-1100322103203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-1011023010200020-2131223312032031-3312120331311200-0223103020132233-0020111301122031-1031223303202331-3131221302030103-3021222112233230"></a>

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
all_repos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011020133023203-1020010031211020-3210220300221021-1030003032211012-3320321132101110-2200123112011310-3123013030010000-1112331223011302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-2213001021220100-0212002231200333-0332131021022322-1332330200121230-3321131330133223-1011011231001110-3230213201130031-3113130020130321"></a>

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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303103323203323-1122131111300220-3313013300302221-0320221313101301-2320211233012221-1013130202002131-2113123222101203-0123011120203033"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-3122103331103203-2300202111030123-2003311232220010-0212312010300300-1210110332131323-2031331313113332-0223200232031102-3232230012121323"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

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

<a id="canonical-3323211331222232-0332211032323102-0312000103032013-0030301203013222-0333230310110223-2300122212032132-1332223123300100-3010333232222230"></a>

<a id="canonical-1332133003010200-2233212322223331-2220323332203203-2032021013130030-1300130232232330-3202132311121311-0302130033302132-0021223011312320"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

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

<a id="canonical-2100023113223133-0112303332131022-0200022321013031-3220132300121300-0330131112201210-0222103310132233-2201300311022322-1330013221203010"></a>

<a id="canonical-3320300103212232-3022323233311031-2222000311322012-0313313132032003-0300201111030133-3321133021102300-3003300323101202-1203001020102320"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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

<a id="canonical-1330221102102213-0130023021302212-1210323220223000-0200101320031231-2211330313210013-3330033102122132-1011102002321103-3013033213013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-0203333112210003-0002301012221031-0133202233030300-1221222311202120-3220211111233323-0211132111331013-3112221100310211-3211022201332333"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
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
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132122020212320-0221132213322301-0323013031213001-0133232002103330-1123302210312102-2333312103223112-2102023033321220-2332030000123120"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-0112113212311002-0103123220203100-3212002132223220-0102201232122102-1302310110232113-2311013110130331-2103023002130010-1312232023310230"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-2032220023231101-3330012233232200-1020310222223220-3323122201013101-1001010023111330-0201300311001023-3100130003020302-1100112010122230"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222312331113303-3112102110130122-0111130301103100-0102222212033101-1031303202122231-3210123302221011-3301100131321120-2003100220033212"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery`

- [api_discovery_ref](resources--http_loadbalancer--reference--group-018.md#canonical-0321132002313203-0132103033123323-2110211300011031-1000230020003333-1332223113131322-2013102213033313-2201313013202233-1203301002312032): complete subsection reference.

<a id="canonical-0321132002313203-0132103033123323-2110211300011031-1000230020003333-1332223113131322-2013102213033313-2201313013202233-1203301002312032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-1211001021322222-0312111230200030-2123123311122302-1012001003001200-1323211032212222-2312100320103210-3103320111030223-1201031202201101"></a>

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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222030002203103-2011310013100132-0011100313302110-1312323100302232-0112110211020221-2033011010102321-3001103023323032-2123203201121112"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-0003311232123102-3121121302231102-0023321323131202-0211023201111000-1210222213201133-1212330313100300-2111010222012212-2232201332303201"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

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

<a id="canonical-2023200001310013-0201311001123122-0020202302133130-1230310111132232-2031202122313102-0331310122312031-0202222312321303-3120120000213110"></a>

<a id="canonical-1031103230021213-3100330320103022-1000120331303001-3010002301100321-3323300323003130-3221122333222233-1013200000122230-1121101212003032"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

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

<a id="canonical-2322321120132011-0132221233231211-3223321121023222-0133313023120000-1311131201132020-2223030323332132-3311331203230213-2330313112331221"></a>

<a id="canonical-3103321030111012-3333021003313231-1301330001202331-2032233212301101-0303310331213031-0320022022313133-3022201202110222-0212112020220111"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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

<a id="canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-0001322113322313-1122330202103100-3133110123212230-2212233203110032-3303333201112213-2110112211220211-1131210010202132-0202010131000321"></a>

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
default_api_auth_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-0123122122113211-3210100333022211-3130321003220032-2121121211313021-2120212321021310-3122003032232223-0100212200111132-1020332231232120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.discovered_api_settings

<a id="canonical-3230103320013130-1123210101310112-1031111033221200-1023322130132321-3010002131133020-3321300112101210-1211113323132201-2033020022220113"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323101310221100-2132032222120032-2233000300222011-3122002330022200-0210130131130202-2233120213311302-1301112233103101-3110300020221322"></a>

### Direct properties for `enable_api_discovery.discovered_api_settings`

<a id="canonical-0102021100111210-1323102312123302-0110110223210031-1333113123232032-3231002022002100-3202213311000311-2230310102301121-2230010021031312"></a>

#### `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-3333233011332032-2021222131121220-2323113103112113-3133112202313123-3033300230331233-2301333000122113-0002200121221333-3001130211233320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_challenge

<a id="canonical-0331332220112013-2112002310330311-0000103101120120-1022133131132112-2221133100010123-1020333222003102-1303303322000302-2321231220030231"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133311032301030-1200121220020301-3030132203023221-0113032100300022-0313322210233013-0033312133213003-0021210123211223-0300011231233222"></a>

### Direct properties for `enable_challenge`

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1122310300102302-3011210221032332-3112302333322111-1300102120103333-0003232101022013-2103123333023313-3310023100122213-0012112333303132): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1231220200222332-1313233111303023-1313331022022121-2223300012113101-1231101221013011-0011112023203321-1132333112033000-2330210301313022): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-2100200102230112-1122110310232323-3002121100330303-2121230122033003-1112132030310331-0322312302131010-2001023023031203-2323030032123210): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3033232023011212-0012221102201233-3001212312010102-2021222023021103-2301330123301302-3210331130331312-2323330111223110-3212300223230132): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-0333223322232310-3133033231010031-2113311012200323-2232211121233033-3132012023221102-1020333312103020-3313232202013023-0302133132031101): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-018.md#canonical-2200030000202100-0111223200112113-1213031201301203-3032003032021013-0333013223212111-2311103031002002-2032320123012330-1111020321212132): complete subsection reference.

<a id="canonical-1122310300102302-3011210221032332-3112302333322111-1300102120103333-0003232101022013-2103123333023313-3310023100122213-0012112333303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-2203231132010022-1323102113113232-0310112231120103-1101113312111301-1232021103033130-2330012302331102-1222121110233132-3001301032333111"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203231132022022-0310030013210123-0311021011011212-1213011311020102-3003113310222122-3132333223031002-3303303011222331-1130213221211210"></a>

### Direct properties for `enable_challenge.captcha_challenge_parameters`

<a id="canonical-3332312122302200-2012301333132230-3303211012120133-3310232312100233-0312222221213103-1003000303230300-3132210112332223-2110111320003213"></a>

#### `enable_challenge.captcha_challenge_parameters.cookie_expiry` property

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3130110122213303-1200110321223102-1231212200022301-1021303110203023-2010203000231130-1331003203333310-1300121312220330-0310210133302331"></a>

<a id="canonical-2121122301022123-3301320220301313-2002002000021211-3202321302230300-1221033220321210-2113121101321322-1122210330130131-1303233302120312"></a>

#### `enable_challenge.captcha_challenge_parameters.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1231220200222332-1313233111303023-1313331022022121-2223300012113101-1231101221013011-0011112023203321-1132333112033000-2330210301313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-2220302033032331-2113332231011223-1312311001123333-2332320321102123-0300220120000211-1300002021221103-2321130100213000-0002112202212210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100200102230112-1122110310232323-3002121100330303-2121230122033003-1112132030310331-0322312302131010-2001023023031203-2323030032123210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-3020300013001121-3231221220202200-2123000212311033-1102011313103101-1012012010320231-1323023132331120-3103330302222331-0123110311013002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033232023011212-0012221102201233-3001212312010102-2021222023021103-2301330123301302-3210331130331312-2323330111223110-3212300223230132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_mitigation_settings

<a id="canonical-0012213112132321-3300331121100211-2103232123310333-2231121032331102-3230231100302320-2232312233101232-2313312323101112-1331021130111122"></a>

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
default_mitigation_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333223322232310-3133033231010031-2113311012200323-2232211121233033-3132012023221102-1020333312103020-3313232202013023-0302133132031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.js_challenge_parameters

<a id="canonical-2213231032301111-2333221232102133-2222032222010233-0300212130123321-1233030123232310-3101100320232113-2023202003230020-0200331012200033"></a>

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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102033321321330-1323300123232023-2120010220112120-2223000131322103-0022213001301033-3221222300223332-1321332322012002-0101101133200132"></a>

### Direct properties for `enable_challenge.js_challenge_parameters`

<a id="canonical-2213102220302033-1003120202100032-3033031313311110-0001201020102013-3212112232122311-2120310332322022-2310210323320131-2000311022010331"></a>

#### `enable_challenge.js_challenge_parameters.cookie_expiry` property

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0220012332112001-0002112221012012-3200232002332100-0110123032313300-0022332120201310-3233010003011033-1221100302300011-0123023202223101"></a>

<a id="canonical-1013302321220331-1003130132111323-2112022330111121-2321203230210330-2221112323202221-1232120133113020-2033330021010011-0312122311303210"></a>

#### `enable_challenge.js_challenge_parameters.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0333330301111300-1113323201302100-3333113321222022-3031303303002301-3232221203232323-0003333011131202-2330123130033231-0302313003201013"></a>

<a id="canonical-1311033130201213-3311032212111310-1023131212310320-3330111131211002-2232101111223222-1101122310123031-3013112210332201-1331203222003113"></a>

#### `enable_challenge.js_challenge_parameters.js_script_delay` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2200030000202100-0111223200112113-1213031201301203-3032003032021013-0333013223212111-2311103031002002-2032320123012330-1111020321212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.malicious_user_mitigation

<a id="canonical-0112333113330120-2200102210201331-3100303112211000-0032313111103310-0103331213202313-3013101201233301-1323202301330003-1203330003210311"></a>

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303310311322203-2021231032013211-1232032102002103-2300203130222110-2001131302332310-2330113102223331-1120022210331330-2212012021102132"></a>

### Direct properties for `enable_challenge.malicious_user_mitigation`

<a id="canonical-2120212311322022-0303112102122211-2212231130112232-0020000112202003-2002310102232330-1332012200303000-3220232030200021-2133000321020332"></a>

#### `enable_challenge.malicious_user_mitigation.name` property

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

<a id="canonical-2030311013201023-1112102011003231-3010223033313301-2312220310313311-1031030030232132-1221013322121013-2300012031132031-3131311100322300"></a>

<a id="canonical-0200032113313022-2032103130332013-1231032123320110-2121011322302321-0003302312311101-2232323012022121-3122103210231313-0122023333323130"></a>

#### `enable_challenge.malicious_user_mitigation.namespace` property

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

<a id="canonical-1231113113213332-3222102323223100-0200213222121010-1010300312131013-2013302232312310-1001001120021002-3330323130220110-1220203110220323"></a>

<a id="canonical-2030111103013122-1312122122122122-1331001031221110-2311233202003111-3313201212233131-0300301111222010-1203312003032231-0322113122212312"></a>

#### `enable_challenge.malicious_user_mitigation.tenant` property

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

<a id="canonical-2231212223321013-0123213202300103-1002202103000112-0121311132023203-1003223131331322-3021301311230232-3302012302222311-0121321022022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ip_reputation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_ip_reputation

<a id="canonical-1331031122011033-2302131121012122-3313303233203321-2313300032330123-2331111321311122-1221001122320101-3033133302202020-0230023201131132"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
enable_ip_reputation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202322233101010-1011202323000323-2231313232232203-1002223100122333-1032123220233012-2213303100101323-2330100031033213-0112100230112112"></a>

### Direct properties for `enable_ip_reputation`

<a id="canonical-3121021201103213-0201320300320231-3333323201003323-2121102323223312-1213210312230303-3033210331212020-3312230032132311-2133313102230220"></a>

#### `enable_ip_reputation.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

<a id="canonical-2102301301321320-1331123320232110-0223023132023311-1300312321320102-3102112231332231-2233232100233210-3122230110022023-1302310210123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_malicious_user_detection

<a id="canonical-1200132303110020-2331223220030122-3300330122310330-2033201303222103-2311130312023020-2001011122320233-2033303101211233-0313110100132203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111322010113111-1002113001030101-3211022303110233-2133221032220321-2011101130012312-1232301001111000-1203020010201230-2331021131223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_threat_mesh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_threat_mesh

<a id="canonical-2100022020020011-0200313113330200-3010002022110312-1310112023231102-1110222331121122-3303232030030220-1330000212210230-0030303210201210"></a>

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
enable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023200122231012-2002203313103003-2012123021321211-3110001220301030-2122312230032131-3200023120120031-3000230111323102-0321222120322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_trust_client_ip_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_trust_client_ip_headers

<a id="canonical-3331120100110120-1303333032232002-0301220313322023-3200013000131023-3322323222012333-3110100300213120-3312203133332021-1101320122300121"></a>

Type: `"object"`. single nested block, Optional.

Trust Client IP Headers List. List of Client IP Headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("client_ip_headers")}
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
enable_trust_client_ip_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312310321002320-2013201330333332-2220211332121223-3231002020003200-0131330022001130-0020211203030131-2111210030012122-0102030122230023"></a>

### Direct properties for `enable_trust_client_ip_headers`

<a id="canonical-3020322332123320-2303232013311313-2112011213213012-2312203301121201-2333220321121231-2032032311310112-2210132210320223-0111133212103331"></a>

#### `enable_trust_client_ip_headers.client_ip_headers` property

Type: `["list", "string"]`. Optional.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- graphql_rules

<a id="canonical-1032222112002010-1031000023012012-2011032032213002-0131002301011011-1011220113220032-2111332222101023-2301110210203132-2120100301232110"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("exact_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("method_get",
    "method_post")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230001211301013-0010020000303130-0201201031230011-3323010031333110-2200123202100131-1333202333031020-3113120332023020-1223301131202302"></a>

### Direct properties for `graphql_rules`

- [any_domain](resources--http_loadbalancer--reference--group-018.md#canonical-3232221130212110-0220113031100232-3030220000211332-2100002113221202-3330301230023021-1120230023230321-2200300122300101-1230331300300101): complete subsection reference.

<a id="canonical-2233102220221121-2322110101022020-0203211210123231-2132031321023121-2202013010200302-1233023032322200-0323331022001323-2312302323021120"></a>

<a id="canonical-2331113310030331-0221032200201001-3331112010203332-0110322103101131-0332320100132201-3001321232012200-1210033211012103-0131312330200231"></a>

#### `graphql_rules.exact_path` property

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Additional upstream details:

Default value is /GraphQL.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0113212322011212-1101010330022101-3221202101112320-1202231122032023-0223331030220020-1013130331013022-3033200001110022-0101030111320031"></a>

<a id="canonical-2133320102033031-0213012103100301-0032332203331103-3122321332330333-0030123013012131-2200101121333003-3320113131031313-3333322300223220"></a>

#### `graphql_rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-018.md#canonical-1311122103230322-0310010320111232-3223011301201233-2200311131220033-2130132102313000-0102213320031221-3101130301333310-1002120311120200): complete subsection reference.

- [method_get](resources--http_loadbalancer--reference--group-018.md#canonical-1131111003001301-0312122332120102-2200110102303301-3210020023332231-3022222101013121-3201333300300020-2210321301111112-2000203111220332): complete subsection reference.

- [method_post](resources--http_loadbalancer--reference--group-018.md#canonical-2231300223101031-0233020132112220-0133203032231333-2221320230211002-2202013323123213-2220223001302313-2311121332013112-3312321121130313): complete subsection reference.

<a id="canonical-1120202001232200-3021302302111002-1102131233031031-1201332202203233-3101211323101303-0201202102223333-2333111223313303-3103221303222203"></a>

<a id="canonical-1101330232020330-0230121313022032-2110221330302011-1011232313023031-2311102200030312-3311032113133102-3120020211200231-1313003302222331"></a>

#### `graphql_rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3232221130212110-0220113031100232-3030220000211332-2100002113221202-3330301230023021-1120230023230321-2200300122300101-1230331300300101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.any_domain

<a id="canonical-1300223311121233-0012220011221313-3112102012221120-3320100333333002-0123033103231022-1003231022210321-3103032030111321-1120132101121201"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.graphql_settings

<a id="canonical-2210321022120320-2221310330323023-3200023310123211-3002123210022303-3031203212132311-2130302311331203-0011231233013312-1000013133102321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GraphQL settings.

Additional upstream details:

GraphQL configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_batched_queries",
    "max_depth",
    "max_total_length"),
  validators.ConflictingObjectAttributes("disable_introspection",
    "enable_introspection")}
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
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

Terraform syntax:

```terraform
graphql_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021233321023310-3233130213023232-3333002222100111-1322322010012323-1303020131310012-3132101133012031-0313120201302121-1002322013013220"></a>

### Direct properties for `graphql_rules.graphql_settings`

- [disable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323): complete subsection reference.

- [enable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-0123033102031013-1310001202232221-0222212331330210-3133001110031331-0200113100323031-1202123312311021-1301112030320213-3302300003110021): complete subsection reference.

<a id="canonical-2312300221033301-0332311020233322-1021312320001331-3012131230102022-3000313102030113-2321133121232221-1101233033010212-3021300223213111"></a>

<a id="canonical-3311112322332320-0010232213303012-3300222313002030-0031332113011102-3321213032033132-1100113030010022-1111200330233023-1301020312311313"></a>

#### `graphql_rules.graphql_settings.max_batched_queries` property

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3202213210332121-2312103303233303-3033031033032212-3232212023302303-1010120120113200-3022231331100313-2031023231303232-3300001230132300"></a>

<a id="canonical-1111221132023111-3001210200302301-1020222210311210-2001011012313300-2212222000111011-1233211121212300-2003032213210032-0113020223030010"></a>

#### `graphql_rules.graphql_settings.max_depth` property

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-2122211323123202-2023231020201333-0110312102200223-3121023202213333-3123323213123003-3311221230111232-2131031001200321-0310112231230233"></a>

<a id="canonical-2002233110010313-0300032112121312-1232302220300011-0312200323212122-2031320223030102-1310113302113320-2022130213302330-2312023000321200"></a>

#### `graphql_rules.graphql_settings.max_total_length` property

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 16386),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.disable_introspection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-2312120300302212-0121202010322020-2121121302133130-1103011022112312-1230311200332210-3321210101313012-1202113303301101-2310302202310332"></a>

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
disable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123033102031013-1310001202232221-0222212331330210-3133001110031331-0200113100323031-1202123312311021-1301112030320213-3302300003110021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.enable_introspection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-2121301001221311-0323230003331333-1232131203113212-1012313302023121-0233131221203222-2102123012220123-2201202300322002-1231130110322020"></a>

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
enable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311122103230322-0310010320111232-3223011301201233-2200311131220033-2130132102313000-0102213320031221-3101130301333310-1002120311120200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.metadata

<a id="canonical-1312330201312313-0013011131322010-2203331002122013-1122330112112000-2022130220030201-2111302123213113-2001032113222223-0233103103320001"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203023111210211-1203232322312012-3301333011032113-2210002300111212-1112021212032120-3221131032133132-3330032000002132-1001132100113013"></a>

### Direct properties for `graphql_rules.metadata`

<a id="canonical-2310022333112110-3213012011013021-2232210322033213-0021120011131131-1032303122222030-1313331103230232-2011100212233101-0012320331111322"></a>

#### `graphql_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0310210330220132-1032121210001031-2320001321000003-0002231220011012-3323021221332010-1323211212302330-2103232121221301-1123331303331123"></a>

<a id="canonical-2302323233300032-3333330010322102-2322110131121203-0002323220322330-3020132311022113-1213200323311310-0200312133120233-2212030002131010"></a>

#### `graphql_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1131111003001301-0312122332120102-2200110102303301-3210020023332231-3022222101013121-3201333300300020-2210321301111112-2000203111220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_get` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.method_get

<a id="canonical-1010232111012121-1111131302103310-2331221201102200-3220331130011021-3303101021123103-2131001332322101-1333202303120302-3322321303010203"></a>

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
method_get = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231300223101031-0233020132112220-0133203032231333-2221320230211002-2202013323123213-2220223001302313-2311121332013112-3312321121130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_post` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.method_post

<a id="canonical-1331230101032323-1311132132033223-1010331133302011-1110020130021133-0203133000231032-0030330322223231-2302313132231212-1323102003230321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

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
method_post = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201330110311330-2123203220332301-2010310002232302-2200333212132031-2112023132010213-3001201001222010-0222333320131100-0013020102021323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- http

<a id="canonical-1202011002202030-1301123223221210-0221133001100200-3113323103312320-1112211202323013-2102200302020220-1011302212010002-2131232013211102"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy. Changing this type selection requires recreation and
may interrupt service. Supported settings within the same selected type remain updatable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1331221222213010-1111020232113203-3133222203001030-0333323200133110-3321330011033022-1101022221321330-2003331022022201-0210303231013101"></a>

### Direct properties for `http`

<a id="canonical-2131122102321022-2330313211011031-3303033230012121-2133113011020220-1320303033100221-0312200102313112-0020201201231333-2201211131330222"></a>

#### `http.dns_volterra_managed` property

Type: `"bool"`. Optional.

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

<a id="canonical-0120011102032230-1003033001211202-2330312220333013-0130001032231131-2002023113112120-0111020111331300-2000231112330211-0001102000221211"></a>

<a id="canonical-1313001221113221-3102032101322223-3330133000213311-0211003003131123-0000013323003132-1130022002103202-1332220233000122-2102320220122202"></a>

#### `http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1010212311212230-3223032300310201-1211001003221013-3233132221220131-2002133303103032-0003321311111223-2102212120312231-1302323232203311"></a>

<a id="canonical-1222133212100220-2002203123023130-0022301002100132-1323000130110133-1332112002020033-1333130121010012-3231031331001031-1032231012020322"></a>

#### `http.port_ranges` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- https

<a id="canonical-1300211231013103-1001303130221323-0000301002110021-2210133210111010-1323100310232132-0101111330112313-3100110313000100-1120103110300213"></a>

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

<a id="canonical-3301011010210201-2311221033333333-0210103131302100-1011003030320220-0102301331310131-1222220220330321-3001111333100331-2321110313322022"></a>

### Direct properties for `https`

<a id="canonical-2100112122110321-2010303103220203-3232301112203331-3002220221212210-1211113223133110-0222230200022231-0100321033103332-3100303201111321"></a>

#### `https.add_hsts` property

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

<a id="canonical-0223011010203002-1202313223300011-2101003030311003-3232233332010130-2020301123123231-2022113312123232-0102312220302133-3000331033003222"></a>

<a id="canonical-2320200110221002-2012210031021230-1123213030020310-1201210131030230-3300120300003001-2023100221323220-2302331120020202-2020113202232000"></a>

#### `https.append_server_name` property

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130): complete subsection reference.

<a id="canonical-3330133120100102-3202120011033200-1133033320331101-0101113002131000-3321120023011002-3003203320301323-0003132233122221-1312300010022102"></a>

<a id="canonical-1210100020203220-2101013310331110-0222333233330000-1102112120120032-1230002232330313-2222123133200202-1002113031020312-0020100101220210"></a>

#### `https.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--http_loadbalancer--reference--group-018.md#canonical-0001322102003113-3211120130133321-0221301230112220-0231230323033033-1020122212100120-0231201121003012-2223220021021011-0312001323312113): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-018.md#canonical-1202321302320013-2012321213132330-1020322133333023-2033121332202310-2201003010201233-0330000122003010-2123113120113011-3332022131113310): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-3310310330000320-0012333332303330-0013011232220321-3101220020113220-1031033213202333-2022111013202310-1321200233213200-1132333020110331): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222): complete subsection reference.

<a id="canonical-1122213102031302-3000123320103002-2220321030122221-3300030332012212-2200311010020010-1203313121332311-0313022120033201-2022130310223313"></a>

<a id="canonical-1310330333221301-0332220011221100-2031133121001012-3312000023110012-3033211011013030-2231012312323312-3003320001010210-0210333002032032"></a>

#### `https.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-3122121110210020-0112331300201220-2103232132122022-2101302110211121-3013232333012111-3310330202000103-2001120322233003-2030011312230303): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-1320021323133122-2131132203230011-0222102300201203-2133211311213200-0330211132231220-3302123120111202-3201010230103132-0121030000310101): complete subsection reference.

<a id="canonical-1302033032112013-3322330303202112-3313300300310301-3003312303320013-3303031300311210-2200000003003103-1302300120211023-2200011200123212"></a>

<a id="canonical-1300300210300231-1123312313101201-0213302330310022-2312330320110211-0330011033333331-2101123220121121-2333001130320020-2332320130203102"></a>

#### `https.port` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3231112133022202-3300033010201033-2020210010200303-1202020102220022-3303313110232011-0213100130310301-2131103333203233-1202201220000210"></a>

<a id="canonical-0003203302200310-3112131211002331-2033303323132322-1132121323130333-2100012010222333-3010232121133220-1212221000101113-2100002233033330"></a>

#### `https.port_ranges` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3202003303131203-0030300311313211-2013303221213131-1201101012311131-2210131023202330-1332123303023312-3133302033232131-2023331212101320"></a>

<a id="canonical-2210032012020322-2033301002231301-2203200033311233-1321312331010021-2220121133120030-2231120232223212-0011120230122031-1210032233332003"></a>

#### `https.server_name` property

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133): complete subsection reference.

- [tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212): complete subsection reference.

<a id="canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.coalescing_options

<a id="canonical-1122033020202321-1320331133312222-2022311310023123-2233110100233032-2021032102101311-0311223200032323-2031301023033002-0130312111330333"></a>

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

<a id="canonical-0303130310112310-0301312332231220-3232233113110213-2210012210233323-0010022233332331-3111233112113323-1210100203100222-2211312231130301"></a>

### Direct properties for `https.coalescing_options`

- [default_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-3320122023030300-1303303301013333-3300220311333112-0311030201220313-3212203012212033-3110331321232100-3102232103221123-1101122320230000): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-2303122331132210-0313232231022031-0320321031222022-3322112320123101-0222002321122103-2202230132200132-2201300330300120-3000003023202202): complete subsection reference.

<a id="canonical-3320122023030300-1303303301013333-3300220311333112-0311030201220313-3212203012212033-3110331321232100-3102232103221123-1101122320230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- https.coalescing_options.default_coalescing

<a id="canonical-1030032320202301-1302021210011012-3321023313002201-1010010232012032-1213120023321213-0032121211132230-2213123203001321-2131233210200211"></a>

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

<a id="canonical-2303122331132210-0313232231022031-0320321031222022-3322112320123101-0222002321122103-2202230132200132-2201300330300120-3000003023202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- https.coalescing_options.strict_coalescing

<a id="canonical-0232130330011310-0030113300302203-1211302133103000-0200322003332102-3000210322321220-0231331213232333-1130230130012100-0021332103112130"></a>

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

<a id="canonical-0001322102003113-3211120130133321-0221301230112220-0231230323033033-1020122212100120-0231201121003012-2223220021021011-0312001323312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.default_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.default_header

<a id="canonical-3230323130312231-0102321210323322-1211022021212320-3332001332001221-2020001211132121-0300123221011312-1023130131203111-2202030020333100"></a>

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

<a id="canonical-1202321302320013-2012321213132330-1020322133333023-2033121332202310-2201003010201233-0330000122003010-2123113120113011-3332022131113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.default_loadbalancer

<a id="canonical-1223202332122332-3331232123313010-1302023121201110-0301030031023213-1003213310022000-2222101333101102-1231313122130323-2321320032330220"></a>

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

<a id="canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.disable_path_normalize

<a id="canonical-1031230121302321-1111221210301321-2000330132301000-3223120331030321-1222001201331133-2011021030120230-1300312232230230-1022100223320023"></a>

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

<a id="canonical-3310310330000320-0012333332303330-0013011232220321-3101220020113220-1031033213202333-2022111013202310-1321200233213200-1132333020110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.enable_path_normalize

<a id="canonical-3100122030321033-1011002222030332-2102213021201023-1033232022211120-1202231030330103-1010200200103322-2213201221032331-3310130202302031"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.http_protocol_options

<a id="canonical-0220313022230110-3030031100202020-0210331332011031-3110320012030313-3033131313020133-0113330202331320-2023131312002230-1201202131331232"></a>

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

<a id="canonical-0323113220021300-3232330032100120-1212201030311233-0003022300302012-3121333212300030-0131302023320303-0032330200121322-3000320201013321"></a>

### Direct properties for `https.http_protocol_options`

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-0010313222101213-2030222023120120-3033111100332312-3000213103211301-1002213101332020-1012332212021033-2333200203003133-1002230333101111): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-0213320000020033-2120001112320123-2330332010013101-3101232320223210-0223322031210210-3112031102302331-0322121121003330-1000320310232102): complete subsection reference.

<a id="canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1021122103311221-0102212302120022-3033113002220000-1131110313333032-3200201221221333-3033000300212330-0023203130202121-2121012200222320"></a>

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

<a id="canonical-3231001221001003-1012302231210010-0222303210010002-1303000103323022-1131122030012031-3301302233133020-1103230133231120-0202132201203100"></a>

### Direct properties for `https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300): complete subsection reference.

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

<a id="canonical-1030000001231131-1113000002213120-1320230101120302-2303120012031123-2102021110213032-0033121101013113-3131331102330122-3230201331111330"></a>

### Direct properties for `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223): complete subsection reference.

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
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
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
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
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
