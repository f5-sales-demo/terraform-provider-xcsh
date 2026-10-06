---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1011302001123301-2313002203212101-0233321031103031-0132102102022202-3111110221001123-1311000231023013-2310321131201012-3120312323220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- active_service_policies

<a id="canonical-1300321000130123-3230103322123302-0120211103233131-0213012012033012-3212122102230002-1022103201121131-3220202003131132-2010123231230312"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

List of service policies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

OneOf alternatives in this subsection:

- [active_service_policies](resources--http_loadbalancer--reference--group-004.md#canonical-1300321000130123-3230103322123302-0120211103233131-0213012012033012-3212122102230002-1022103201121131-3220202003131132-2010123231230312)
- [no_service_policies](resources--http_loadbalancer--reference--group-022.md#canonical-1101022221112332-2030130211022031-1133103101333213-1101003230013311-3103302223310030-3213132312012201-0000201122030113-1223213310313002)
- [service_policies_from_namespace](resources--http_loadbalancer--reference--group-027.md#canonical-3123232231202213-3113013113132233-2130200300221332-2101202232201032-3213020100120021-3201232030111013-1022231010320313-1100333020311221)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102200020223302-3300112002102221-0220212211302130-3320200220120130-1213023211011030-2321221320020033-2032323330300231-3023212131313220"></a>

### Direct properties for `active_service_policies`

- [policies](resources--http_loadbalancer--reference--group-004.md#canonical-1102120130022111-2203300021320112-2313312112103102-0201100323223120-1301033002223032-1303221322202122-3222201100233131-0111223211111011): complete subsection reference.

<a id="canonical-1102120130022111-2203300021320112-2313312112103102-0201100323223120-1301033002223032-1303221322202122-3222201100233131-0111223211111011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [active_service_policies](resources--http_loadbalancer--reference--group-004.md#canonical-1011302001123301-2313002203212101-0233321031103031-0132102102022202-3111110221001123-1311000231023013-2310321131201012-3120312323220332)
- active_service_policies.policies

<a id="canonical-0111012221132030-1211003011123033-1030023300110322-3110121000112023-3132131003332220-0132102202233020-1332113100210210-1201223000132110"></a>

Type: `"object"`. list nested block, Optional.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023220212202031-3033102300333202-3120012312022330-1322002001102111-2032210211211321-2311012133321011-0233232110233303-1320212203311110"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-1032201003110210-0131111013003202-2212003200213213-1322220213330011-3012312120311113-3023321022031310-0130232100131322-1031301213201103"></a>

#### `active_service_policies.policies.name` property

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

<a id="canonical-1202301202231012-3112123223133103-1001210120110213-1121023030310211-3011120020113300-2112323330331111-3320221032003233-0102000122011011"></a>

<a id="canonical-1010020112202322-1300320030111220-3131332210021012-3322031130223303-0121131102102300-3111232031232230-3321311111033222-0132013331333312"></a>

#### `active_service_policies.policies.namespace` property

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

<a id="canonical-3330302330301100-1120220112133112-1120321210010002-3311221202231213-0220233313222210-3201210122132311-0030221330300112-3122120000031321"></a>

<a id="canonical-0130220120002310-1012203003222031-2213323133012013-1003011302010333-0011203331110330-2010320312003231-2213111202010112-2013232201023310"></a>

#### `active_service_policies.policies.tenant` property

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

<a id="canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- advertise_custom

<a id="canonical-1323310123021132-2201333113300112-0000332012011101-1332211332121220-2231020301111111-3200321010321031-1032013200002320-2013131100312020"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_dualstack\_on\_public, advertise\_on\_public,
advertise\_on\_public\_default\_vip, advertise\_v6\_on\_public, do\_not\_advertise; Default:
advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP on specific sites.

Additional upstream details:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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

OneOf alternatives in this subsection:

- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-1323310123021132-2201333113300112-0000332012011101-1332211332121220-2231020301111111-3200321010321031-1032013200002320-2013131100312020)
- [advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-3133202320110300-1221202032022320-2123233221311111-0220213003213322-3131002302103012-2313211330110110-2023230303302220-1231232220321012)
- [advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-3120301223332031-2313032220213200-3110101203310333-0021212203000211-1301233323010213-0332001033311132-2001200210310230-3201203203013220)
- [advertise_on_public_default_vip](resources--http_loadbalancer--reference--group-005.md#canonical-3201330121022122-3323220020131223-3101123310000020-1020300323303132-3021003123202123-3011313330020133-1101012321201313-2230032300223013)
- [advertise_v6_on_public](resources--http_loadbalancer--reference--group-005.md#canonical-1121221002220212-0021120313311013-2132132013231303-1020103320301022-0302011322201310-2312232122212000-1221031020022111-3101323212220030)
- [do_not_advertise](resources--http_loadbalancer--reference--group-017.md#canonical-3021311322031332-3001130120301333-0101131021301003-2310200320010222-0133031223311213-1001132011030231-1000103013131203-3110003031031311)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203113133200033-3013130223200323-0303233112323031-3110121332031200-3300000000303121-2323022313221110-2330012310032300-2211022102020021"></a>

### Direct properties for `advertise_custom`

- [advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120): complete subsection reference.

<a id="canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- advertise_custom.advertise_where

<a id="canonical-2121302303323311-1113133302000031-1203110321332223-1023310023020223-2000223333233322-2100013100023023-2023221011132302-1111032131130130"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
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

<a id="canonical-0013322313320312-3101201131131303-3313323322211203-1111212111011032-1332000311001210-1121013020033301-1221230203111132-2023010201231032"></a>

### Direct properties for `advertise_custom.advertise_where`

- [advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-2003111123102233-1131302111100233-2023010012232333-2121210322103331-0122202113321211-3021200113111200-2313211233100321-3222321303233022): complete subsection reference.

- [advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-0332101201303102-0213100222310231-2101131011223002-1313320022230023-1132232031032112-0233122212110322-0020033033221220-3133131002302221): complete subsection reference.

- [advertise_v6_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-3333213333033232-0001311112101310-3233010022313232-2111001103233101-3132222301100013-3211003320332323-0112303220303212-0012222131303323): complete subsection reference.

<a id="canonical-3001123313003323-3232331030321301-2221121032031132-3003202133330302-1210332110223232-1332201032322033-3233303220103330-3023230032111020"></a>

<a id="canonical-1111212110003002-1220211112221313-0212111232103113-0222202013213230-0221311322010032-2111332012212233-0311122201101033-1112112112123120"></a>

#### `advertise_custom.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2000113111023001-1332133021133320-1301003112300110-0303211210113002-1231000322003011-2222233301123123-1322131103321212-2122321332201210"></a>

<a id="canonical-1023122132120231-2301011212030320-0021300200200112-3302121131033313-1323031030333021-0130131102133122-2110021210100221-0323021222202301"></a>

#### `advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

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

- [site](resources--http_loadbalancer--reference--group-004.md#canonical-3223332101012121-0323331232121320-3221110230330223-3300020230301000-3201301301131220-3113002032102023-2211122312123100-2020322231030010): complete subsection reference.

- [use_default_port](resources--http_loadbalancer--reference--group-004.md#canonical-2120112203220322-0333003120010022-3103323201200001-0120131312031010-0012133133121033-2222303231110323-0303321132023013-0320122113323221): complete subsection reference.

- [virtual_network](resources--http_loadbalancer--reference--group-004.md#canonical-3202101122301212-3221002001221321-3133320132032033-0302031321101001-2303223203133313-3102303303030211-0112100012031122-3213303020032121): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-004.md#canonical-3221131013011321-1113030320010203-2222330303332000-2333011323121303-3131032130013322-0013102201003333-1302202230321222-0313331030211201): complete subsection reference.

- [virtual_site_with_vip](resources--http_loadbalancer--reference--group-004.md#canonical-2000120200111101-0202102122122023-0013210323130103-0231321302021202-1120122100012003-0121310031222133-3230222202210023-3111232102001222): complete subsection reference.

- [vk8s_service](resources--http_loadbalancer--reference--group-004.md#canonical-0001302012000011-0331023310222121-0133300122120212-2223133302102112-2312030133033230-2002012200300221-1202011102112310-0311321032012131): complete subsection reference.

<a id="canonical-2003111123102233-1131302111100233-2023010012232333-2121210322103331-0122202113321211-3021200113111200-2313211233100321-3222321303233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2331021200212122-2330221202011303-2320102203310303-0032310302303231-1031212333013111-2313203313000021-0323012311003302-2231210201020221"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122023231302123-1321311103133213-1303001011213130-0003232113330221-0132122031012303-2223331120310301-2203323103311013-2201033333011211"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-004.md#canonical-0122113122322032-0100010222301323-3300011213212211-2122023032301132-2023233320230010-2220213102300032-2213230331013200-0332011223021312): complete subsection reference.

<a id="canonical-0122113122322032-0100010222301323-3300011213212211-2122023032301132-2023233320230010-2220213102300032-2213230331013200-0332011223021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-2003111123102233-1131302111100233-2023010012232333-2121210322103331-0122202113321211-3021200113111200-2313211233100321-3222321303233022)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-2311331222301101-3200013223001033-1202212001011010-2202321020232220-1101012320201033-1023021011223212-2233032222332102-3300331202111120"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302203301001220-1322020033030331-2303210221331312-1103310031312001-1121023002132310-3122322103010231-3303303203202013-3101323133232003"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-1133223222133230-2122010322030201-1030112221310300-0012330213312302-3203210020121001-0313102220112133-3122222213310213-3133123300301120"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-3030313233011123-0010213301203010-2010302021020323-0102212222122023-3013110030221023-0233303111013001-0112011312310111-0112033100011311"></a>

<a id="canonical-1332013233322020-2003130320301022-1213110121311103-2210013230310213-2132211232100131-3312031000333230-0203331033100233-2212333111330322"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-3303002012233232-3010301330131323-1221100221123333-0013222112203323-1310022220121221-1333010331220112-2202002300303312-2120003312212232"></a>

<a id="canonical-0300322201023203-0113230321100300-0111233103310111-1210122233310021-2100121110212313-3123100320133100-3121222231103010-3330210211333023"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-0332101201303102-0213100222310231-2101131011223002-1313320022230023-1132232031032112-0233122212110322-0020033033221220-3133131002302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-1112013220231233-2031002032321200-1301312210110001-3011123111013333-2323313210010020-3112232123000202-0230321101021000-3231131020032222"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112030313223331-0132320022310003-0311310320100121-2200320201213223-1212130123122101-1011100002221123-0330110202323221-1020121033133133"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-004.md#canonical-3332232020030122-2033132303020032-2311133020122331-2111213210313103-3332132230032131-3302103302112333-0310323213030233-2002001001231003): complete subsection reference.

<a id="canonical-3332232020030122-2033132303020032-2311133020122331-2111213210313103-3332132230032131-3302103302112333-0310323213030233-2002001001231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-0332101201303102-0213100222310231-2101131011223002-1313320022230023-1132232031032112-0233122212110322-0020033033221220-3133131002302221)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-2232322110020132-3130001112301201-1130220331232130-0021231032201032-3021230301331102-0021321332332232-2323230100121013-3313120211222100"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300223131331223-0020023033222131-1102021300022331-1002112230020331-2122213130012330-0322032201333111-2211003301202022-3020033112031222"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-3333212033131322-2202330303002010-3313223013311030-0012200311201100-3023231232131131-3301203311022023-0022022221121202-0111332023120230"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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

<a id="canonical-0300302323311101-3320121330200121-1001113033012121-0222310221332100-3100221121310211-2231121313013312-1200312112203333-1031303313312111"></a>

<a id="canonical-3113032100110223-2101021311211122-0202100213310101-1032223322310020-0320023020313331-2110033203120003-1320231300220001-1321331320210222"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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

<a id="canonical-1101332102132212-0212101203322110-1113313231121321-1221323310100211-3211300110122223-2121212332210003-1020312220022222-3022132331221230"></a>

<a id="canonical-1301313210030211-0112201302220231-3311132012331000-0011013232321222-1132102130302032-2332010302013302-0330102322123121-0330121121332100"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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

<a id="canonical-3333213333033232-0001311112101310-3233010022313232-2111001103233101-3132222301100013-3211003320332323-0112303220303212-0012222131303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-3221023221102230-1101203301023213-3302322211311230-2133221311233010-1133211032221210-3331030011330203-2220120031131133-2130021113220300"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302023012002330-2232133311331312-0221121311231310-0323302132310112-1220031122003232-2312201100230223-1011020311222220-2212211211030001"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-004.md#canonical-3201010302231321-0300220231210210-2320323132303010-3103300001313133-2233302210300122-2233200212131133-3101310133210310-2011020023132200): complete subsection reference.

<a id="canonical-3201010302231321-0300220231210210-2320323132303010-3103300001313133-2233302210300122-2233200212131133-3101310133210310-2011020023132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-3333213333033232-0001311112101310-3233010022313232-2111001103233101-3132222301100013-3211003320332323-0112303220303212-0012222131303323)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-0101202213221311-3322303200310310-0231120033332302-3021211001332031-0130233313200211-2131222222102001-1120120123013232-2033120331202302"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223103220332010-0011012000311012-1033020033000220-3331031022120032-2213313321131221-1000032230130100-3213013232200331-1223103100030011"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-0201310300112130-1102110030223110-3013111031302030-1001212003122110-2003031133000201-2110331102003103-3023000000032031-0022331313312211"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-2330221130322111-0021032131212011-3323121133311201-2201001102301133-3332213310311220-0100230201303322-3101020320031232-1110200321212332"></a>

<a id="canonical-2211012032111233-2020103132213320-1203110112221013-1111023031032103-3112023223030213-2012110023221001-2301332320133223-3112222232320001"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-1130313311133323-1200001001112001-3233212120220323-0122210022022311-2233023031203313-0313103221031332-2111122020133310-2230121300332200"></a>

<a id="canonical-1032113310323301-0100021232112212-1230300013103212-1123320100123310-3321333112120013-2030321031020133-2120133320211033-1101103312313220"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-3223332101012121-0323331232121320-3221110230330223-3300020230301000-3201301301131220-3113002032102023-2211122312123100-2020322231030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.site

<a id="canonical-3203121201100003-3222201121030021-0302321110232311-1312312202212030-0322331120013111-1131210221311100-3010301012030130-3223121203212113"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2332301310123200-2321223222030131-3322012003303033-1103330133023011-0111121003310101-0132122223000023-3321130331022101-1323211020021322"></a>

### Direct properties for `advertise_custom.advertise_where.site`

<a id="canonical-2333310312121213-1323322200320103-0112220012121222-0121001320301221-2211133122303031-2211202200101123-1203123231032002-3332323110320301"></a>

#### `advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2232012113032003-0020003113120212-0331123302121313-2100011000020021-3231030132202113-0301031210001023-2200332313133331-0013332303033003"></a>

<a id="canonical-0023110313121222-2100130332110333-0121330330010021-1311223001320030-1002211211022101-2132032123120120-0023121032113313-3313130103021033"></a>

#### `advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [site](resources--http_loadbalancer--reference--group-004.md#canonical-0221022220003300-3300020021203001-0001013211213331-3100312221031002-2203301001120313-0010212030210032-2101132232030113-3323012223211223): complete subsection reference.

<a id="canonical-0221022220003300-3300020021203001-0001013211213331-3100312221031002-2203301001120313-0010212030210032-2101132232030113-3323012223211223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.site](resources--http_loadbalancer--reference--group-004.md#canonical-3223332101012121-0323331232121320-3221110230330223-3300020230301000-3201301301131220-3113002032102023-2211122312123100-2020322231030010)
- advertise_custom.advertise_where.site.site

<a id="canonical-0230222102101031-2333103022132303-0210013313003212-2131330130031232-3003321031203102-2330230113333231-1013220300212131-2122221130100310"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130303320232310-1333301232023112-3131310223320221-3001022000321300-1102131123100201-2300222031122330-0010131013001301-0021002333313333"></a>

### Direct properties for `advertise_custom.advertise_where.site.site`

<a id="canonical-2232011022111223-3012331121333022-2213123131030110-3010003321121322-1201112302010320-0130001323300000-3333101331130233-2333230130010230"></a>

#### `advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-3002122201011122-0310230022112123-1222131202022012-2110320013223103-0001101013330022-3022003220302130-0311201210230122-0312233210130231"></a>

<a id="canonical-0122120333032302-0322111301300111-0310332331330323-0103121000022022-3230310222213100-2331323113022222-0001103323112330-3301103233103203"></a>

#### `advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-2010110032302003-0233211230201311-0202322201112220-3131013100313101-2103310100101330-2121011001311213-0330323000033223-1313213012012303"></a>

<a id="canonical-0203022203301331-1111131203213012-0021030302232332-3301310033120113-1132001033331331-1133112133022131-3230112031112131-2120113032202130"></a>

#### `advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-2120112203220322-0333003120010022-3103323201200001-0120131312031010-0012133133121033-2222303231110323-0303321132023013-0320122113323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-0211122133113212-3212133133011222-3111033010320300-2330303122232010-3330002220112110-1231311331103120-3100112131023010-2301330012211102"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202101122301212-3221002001221321-3133320132032033-0302031321101001-2303223203133313-3102303303030211-0112100012031122-3213303020032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-1213002323231330-0232032110132100-2120133223303131-3212212120132001-2232002013211102-3200220232013321-2023031202030333-1102132131132301"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211333130122300-2113320033000333-1012122102200012-0112220333333133-2320023302031103-2310322213130301-3111312033122112-0303231032232032"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](resources--http_loadbalancer--reference--group-004.md#canonical-2022001103010323-0202003202010220-2101320300331222-1113330300220020-1321102210210320-0033102322303303-2221202331202110-2230213101133101): complete subsection reference.

- [default_vip](resources--http_loadbalancer--reference--group-004.md#canonical-1202321101130232-1311323212122321-0221211231222123-3212203200113332-1000300123321223-3211222303122031-3120301110031202-0320131222121002): complete subsection reference.

<a id="canonical-3030232022020021-2201031301312300-2003302222033003-2313033002232131-0111200322331102-1022322002203033-2003111122020302-3301202102000322"></a>

<a id="canonical-0312310300131332-2023321321313123-3313110020303100-0032101110231130-1103220021331030-3000030122032301-0010131233110212-3213030010213111"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3211001301031213-3011223011020220-2201222001323231-2300333222210320-3201112112303110-3130202233222010-0301313320031310-0311223302221231"></a>

<a id="canonical-3030032011233131-0200220012311330-3230302202130012-1021311321200020-2222103021202332-2130102221021111-2203000321312103-1300201201111201"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [virtual_network](resources--http_loadbalancer--reference--group-004.md#canonical-0003302322230302-3133103212211012-0133231122303323-1212010330312130-2031122223133102-0220013000221222-3133110310330000-0130312323221022): complete subsection reference.

<a id="canonical-2022001103010323-0202003202010220-2101320300331222-1113330300220020-1321102210210320-0033102322303303-2221202331202110-2230213101133101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.virtual_network](resources--http_loadbalancer--reference--group-004.md#canonical-3202101122301212-3221002001221321-3133320132032033-0302031321101001-2303223203133313-3102303303030211-0112100012031122-3213303020032121)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-1021220121120321-1313231331020130-2001122201112320-0102123131221230-2201111310020333-3110221033333223-3032131310330031-2230123122111330"></a>

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
default_v6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202321101130232-1311323212122321-0221211231222123-3212203200113332-1000300123321223-3211222303122031-3120301110031202-0320131222121002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.virtual_network](resources--http_loadbalancer--reference--group-004.md#canonical-3202101122301212-3221002001221321-3133320132032033-0302031321101001-2303223203133313-3102303303030211-0112100012031122-3213303020032121)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-0130103330313221-1012211033112112-2112012210222210-0032303210233122-2211330031311120-1011323231031011-3101110303201310-1113031333102210"></a>

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
default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003302322230302-3133103212211012-0133231122303323-1212010330312130-2031122223133102-0220013000221222-3133110310330000-0130312323221022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.virtual_network](resources--http_loadbalancer--reference--group-004.md#canonical-3202101122301212-3221002001221321-3133320132032033-0302031321101001-2303223203133313-3102303303030211-0112100012031122-3213303020032121)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-3133211021302212-1212001211333321-1230103210103103-0001013012022203-0133001102233303-2331130312010112-1203020200003033-2020032323112103"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303213002332021-3330032120002131-0030322002130113-1130223123300200-3130312320330213-2131101020002212-3311122201020101-1012322331110122"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-2101322022321230-1002201100032230-3330112311200120-1012332123033320-2301110211211320-1202130101220010-3202230302031303-0311332213100002"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.name` property

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

<a id="canonical-2032033210321122-0212123110012100-3101112021303300-2022212022103030-2010131331133131-3311130023003001-1131013122330132-2120103133323220"></a>

<a id="canonical-3203013301022333-1113200101100230-0021311032010023-1131130220313030-3121232332223010-1110313321133132-3213211201200100-0200122201102012"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

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

<a id="canonical-2101111123212100-1202311203021111-0033313032130103-2020332302320211-2020011103001111-2200111332211113-0001122111022203-0232331002120301"></a>

<a id="canonical-1011311233113001-1030013020210321-2122131320130321-3010330230301310-1110223323102213-1013102012033033-0222213002312323-3330000231112232"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

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

<a id="canonical-3221131013011321-1113030320010203-2222330303332000-2333011323121303-3131032130013322-0013102201003333-1302202230321222-0313331030211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-1131121110322001-0200333311030200-0303320121012013-2133023330131313-0100323230000020-1203201202110331-2002130311112002-3131000013023110"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2332012000210302-1132001010102102-0212113312011322-1012313333223011-1210030120131013-3202310131002310-3223222113201012-1313301200210132"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site`

<a id="canonical-1320203313133203-1031120023002001-3310110322131031-2223221231100330-0121303110212121-1213123322223112-1221133111202312-0320110033332102"></a>

#### `advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [virtual_site](resources--http_loadbalancer--reference--group-004.md#canonical-3023301212322111-1020210200112301-0121222221302310-3203201210031301-2102000310220300-3303133320233202-2130102301010302-3022220002332220): complete subsection reference.

<a id="canonical-3023301212322111-1020210200112301-0121222221302310-3203201210031301-2102000310220300-3303133320233202-2130102301010302-3022220002332220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.virtual_site](resources--http_loadbalancer--reference--group-004.md#canonical-3221131013011321-1113030320010203-2222330303332000-2333011323121303-3131032130013322-0013102201003333-1302202230321222-0313331030211201)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-2323310301110210-1103211301311220-3132013100302121-0320201012103303-3032231213020321-2321233210221231-2310232230233001-0330233221120112"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322011210202223-1111300303211311-3020230110300303-3132111111123023-3220232311330233-0000312301103010-0120321011002002-1221031221322211"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2101123202131312-3012311330101101-0002221031333320-2121302300333112-2133000033002022-2132210313100232-1212032312021010-1102232300200322"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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

<a id="canonical-1330020011221111-2102022320303222-3000311103121313-0312131130230311-2322231000310303-0313101223231210-1021322001000011-2311230113223213"></a>

<a id="canonical-2033210310303213-0321221101322021-1010003012023323-1112210323112111-0231012130332321-0323222110102111-0321112121331210-1203013012212003"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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

<a id="canonical-1210301230323321-2111103232011123-1303231130233202-2001311202233223-1121110013030302-3303000133320221-0300103310322130-2033021031120101"></a>

<a id="canonical-3022233320030100-2012331032303321-3331111333323033-2300021022000021-0222202111113332-3023130302111300-0033011110132331-3200013103222310"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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

<a id="canonical-2000120200111101-0202102122122023-0013210323130103-0231321302021202-1120122100012003-0121310031222133-3230222202210023-3111232102001222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-3001130333013303-0002302320232311-3123202231132021-1122001120302213-3002012322020230-1102031020120102-2003331230223130-3022320103130123"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type and IP where a load
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
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103123300000113-3301231122321013-1023010032333331-0213113033331112-3122113302002220-1010220010303010-2011121000311030-1321220033003233"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-0021022300303033-3102013212313311-1322331323202321-3131210331323333-0021121330210123-1233203113330313-2110323102021120-0130233313333123"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3331131202321000-2331223013332222-1120320302022210-3012213323020132-3321310002112011-0220012121110210-0131012000233011-2133131321121113"></a>

<a id="canonical-3223013130200313-2112210023213210-2323110020331320-3111301032301202-1031011301310012-3002110003311000-2323201031012131-2313021113203202"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_SPECIFIED_VIP_INSIDE","SITE_NETWORK_SPECIFIED_VIP_OUTSIDE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--http_loadbalancer--reference--group-004.md#canonical-0031313201333202-1020132110321122-2333130030110200-2023203332312020-3212023310330333-0222112003313223-1011213113100322-0222112132120201): complete subsection reference.

<a id="canonical-0031313201333202-1020132110321122-2333130030110200-2023203332312020-3212023310330333-0222112003313223-1011213113100322-0222112132120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--http_loadbalancer--reference--group-004.md#canonical-2000120200111101-0202102122122023-0013210323130103-0231321302021202-1120122100012003-0121310031222133-3230222202210023-3111232102001222)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-2211102210223313-2113110020101331-3220321303021033-3023022300220021-1233100211030032-1233231010212322-2123121031113330-3321310033031030"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331013202122031-0230131222312321-1211002103030212-2030110122131323-1103320130013332-3021113132333231-3230030102031323-3203113120001232"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-0012100333002220-0023102121220003-0112020321100200-3131102112112022-1200131211021121-2000233021313011-0223231232111030-2201033330112002"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

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

<a id="canonical-3220032022310120-3003301303122332-2001133303223220-0331011132113301-0022022101211101-2003112002232110-0333000332222322-3332003311222001"></a>

<a id="canonical-2023121031333303-3111323321333320-1132003331222023-3220320011331321-1030301302212002-0322110113300112-1321211221011212-2233111300011110"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

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

<a id="canonical-2333300130320030-0023130331023123-3231311023033021-0003132133020032-0000020002302123-0110121001101213-0133310220221112-0212133301303220"></a>

<a id="canonical-1210213202120013-1223000103132221-0200100333000111-0322220021303232-1320210203221202-0312011133202130-1133303200203233-0023212200203210"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

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

<a id="canonical-0001302012000011-0331023310222121-0133300122120212-2223133302102112-2312030133033230-2002012200300221-1202011102112310-0311321032012131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-2301320010300133-2110030221010103-0032033303100311-2312201312321332-2011323322123321-1222012303221033-2212322011003210-2331302231133132"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1221200230213133-2102220003031031-3221032113233323-0031330123302303-0312003323002011-1123020133021213-2220200111132321-2210221033200201"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service`

- [site](resources--http_loadbalancer--reference--group-004.md#canonical-3110320300102300-1233333120232223-1313110001221223-1003231013221310-0213121031031011-2210232020131231-1303230122100110-1003131123110102): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-004.md#canonical-3111201121203320-0032103212010231-3200112303132321-2332300213111230-3231210031222131-3102213303112212-0310330210001102-3331021303303323): complete subsection reference.

<a id="canonical-3110320300102300-1233333120232223-1313110001221223-1003231013221310-0213121031031011-2210232020131231-1303230122100110-1003131123110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.vk8s_service](resources--http_loadbalancer--reference--group-004.md#canonical-0001302012000011-0331023310222121-0133300122120212-2223133302102112-2312030133033230-2002012200300221-1202011102112310-0311321032012131)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-1300133102011002-2013213231231222-2010003323032100-0122232123123033-1332213121002031-1130032111033101-2333232130003032-1200021203031302"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310313202113320-1011012303101021-2020231322010003-3321011002013332-1323012012322021-3323101223223312-1323302112232000-3301030211303013"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-0232030231331133-0013130211103033-3330202133103001-2231321330111231-3312232023112021-2111203210220200-1113100320331113-2213132200010022"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.name` property

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

<a id="canonical-1012300101302210-0332200330013323-2011332310223211-2101121013010321-0131113011032323-0100322001321303-0303011100030123-2001323032220123"></a>

<a id="canonical-3010121310311012-2121311111100022-0002213213103013-1012323323002233-3033321301023312-1003002013200000-1203200333312310-0200230203311000"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.namespace` property

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

<a id="canonical-0221013213023013-1022302223300100-1203033212013213-3113232121123330-1000110201311311-1220102112332331-1103301103202023-2210020003031300"></a>

<a id="canonical-1123033300013330-0102132213332120-0010002233231103-1330333302003332-3321320033023203-3110203300312221-3131300023311011-3321033031301131"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.tenant` property

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

<a id="canonical-3111201121203320-0032103212010231-3200112303132321-2332300213111230-3231210031222131-3102213303112212-0310330210001102-3331021303303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011)
- [advertise_custom.advertise_where](resources--http_loadbalancer--reference--group-004.md#canonical-0321102301003223-0020323003211310-0002300130032100-2010111231011120-1220112331113230-3113012232321203-2320302322101231-1102320020302120)
- [advertise_custom.advertise_where.vk8s_service](resources--http_loadbalancer--reference--group-004.md#canonical-0001302012000011-0331023310222121-0133300122120212-2223133302102112-2312030133033230-2002012200300221-1202011102112310-0311321032012131)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1031212313221232-2110210032311300-1010021002003023-1203131101001123-1202301130312221-2120103102123101-3133220301323002-1323233220002330"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220322303312112-3013213330320002-1130210312233333-1221302210320320-1130112323023213-1302230222003211-3322013231133213-0121121223123233"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-2023130132300122-3101330023103213-3122133331313331-2301132010203000-0113120331000102-2111130001232032-1011303310200022-3312112032030110"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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

<a id="canonical-0130301313030302-0131201132213010-1023132011220121-0111113320320101-0023201110230101-2300110212321112-2330201313323013-2003022212110033"></a>

<a id="canonical-3021101301232220-2002102103213303-2223332233112301-3023002012001232-3323230331002003-0033223112203133-1011101231131102-3213200211212310"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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

<a id="canonical-0321311001130100-0011231321333230-3111332030123003-1302012022320030-1110200220033311-2003233211001123-2032010121310202-1233002131100122"></a>

<a id="canonical-1322310022012121-3110131012230302-0332201323010211-3223200011223133-0201333201302131-0233310331233331-0113332011320020-2133001201212230"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

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

<a id="canonical-0223122302110001-1233213021312213-1303312231122322-3210130310233332-1323202101300301-0012222323223300-1111202232330131-0221301031130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- advertise_dualstack_on_public

<a id="canonical-3133202320110300-1221202032022320-2123233221311111-0220213003213322-3131002302103012-2313211330110110-2023230303302220-1231232220321012"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211213111023300-2011010022301321-2303221213231330-3232112010331012-0022223220003020-2211123131320013-2302100130100313-3300311311332311"></a>

### Direct properties for `advertise_dualstack_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-004.md#canonical-2133303022303203-0021001102210203-3232231031003301-1100020300202023-0211221102321120-2101221032110013-2122212013032220-0223220333310202): complete subsection reference.

<a id="canonical-2133303022303203-0021001102210203-3232231031003301-1100020300202023-0211221102321120-2101221032110013-2122212013032220-0223220333310202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-0223122302110001-1233213021312213-1303312231122322-3210130310233332-1323202101300301-0012222323223300-1111202232330131-0221301031130130)
- advertise_dualstack_on_public.public_ip

<a id="canonical-0000003031331131-1212231102030030-3011322133200302-1012033220330200-3330211203232132-1231122000020013-0112013222300020-2323333021020200"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133110101103221-0322032330230021-0010220302022101-0001111112030212-0332033330212210-3003010100303203-2231212311103002-3330100211113123"></a>

### Direct properties for `advertise_dualstack_on_public.public_ip`

<a id="canonical-1223033220303211-0210210000120201-2210001311312223-0022121302130011-2023312100231003-1233211223303310-1103223032013331-3031222101121231"></a>

#### `advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-2223321231321123-2013330323221230-3031230120131222-1100203031130032-1210132121323300-2020013233003231-1133102130311101-0102103302331302"></a>

<a id="canonical-3312332121031013-2200300202220223-3320122203032213-3312302112331313-2022002321101232-0310222301001032-1121310122322332-2302302321213133"></a>

#### `advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-0222002301203220-2230331032201003-3022211222001303-0031200130232322-1100101211112220-2313101101222100-0011112211020122-0310132223331123"></a>

<a id="canonical-0113303021132011-3120023331131300-0101223133223231-0023230031001113-3012333302132103-3230131201212033-3221023213132120-2110230220011321"></a>

#### `advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-2032030213022320-1201110321010130-2223122210020021-2330313311130213-2003011213011220-2003321021213303-2303033002202100-1300231010012311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- advertise_on_public

<a id="canonical-3120301223332031-2313032220213200-3110101203310333-0021212203000211-1301233323010213-0332001033311132-2001200210310230-3201203203013220"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000101323013200-3333302103233012-2303301221312021-2211312021323223-3012001222223103-1210022002211303-1230213123002302-1321020013131312"></a>

### Direct properties for `advertise_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-004.md#canonical-3323230100212101-2133010102211321-2003302122211303-2113310230303032-0011201313000113-1222220322101112-2111121121102012-3103223300121301): complete subsection reference.

<a id="canonical-3323230100212101-2133010102211321-2003302122211303-2113310230303032-0011201313000113-1222220322101112-2111121121102012-3103223300121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-2032030213022320-1201110321010130-2223122210020021-2330313311130213-2003011213011220-2003321021213303-2303033002202100-1300231010012311)
- advertise_on_public.public_ip

<a id="canonical-1002121123301011-2011132222312303-2102322001010213-1212101220033031-3313213030233301-1011033023100303-1322112203033003-3133313312130231"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220030223010302-3312013332321120-2330113000100002-1231002011322203-3133300311131003-2023111331001101-2233223022321131-3203223002320331"></a>

### Direct properties for `advertise_on_public.public_ip`

<a id="canonical-1010112000022021-1323310310332222-2023213211333020-3003213322030003-2201213022123212-0002232113000130-3203102112101231-2120021133202333"></a>

#### `advertise_on_public.public_ip.name` property

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

<a id="canonical-3121001321100121-3330111221121221-2102323331020312-0131003023021013-3301102001330023-3032320332020133-1302303032033120-1003332310212231"></a>
