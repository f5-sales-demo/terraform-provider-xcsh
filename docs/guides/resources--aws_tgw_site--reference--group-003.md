---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-2120200123021003-0222130113030311-3211020100113121-1002331131210312-1133022212032011-0222132303112331-1021202120213023-1333120022310211"></a>

### Direct properties for `tgw_security.active_network_policies`

- [network_policies](resources--aws_tgw_site--reference--group-003.md#canonical-2303210302020230-0330032302331132-3033322230220123-3331232211012013-3201201131301032-3223333130233302-3210330021122201-0013111113323011): complete subsection reference.

<a id="canonical-2303210302020230-0330032302331132-3033322230220123-3331232211012013-3201201131301032-3223333130233302-3210330021122201-0013111113323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_network_policies.network_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120)
- tgw_security.active_network_policies.network_policies

<a id="canonical-3322031111021233-2111122223311232-1123330223011130-1113203303012103-0101212220101211-1310231030032023-0030213132100231-3030310320013233"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231321012022330-2121313100102211-0222021210120321-3321203003201210-0101101100121331-0101230023010001-0331212132012201-2231101023313231"></a>

### Direct properties for `tgw_security.active_network_policies.network_policies`

<a id="canonical-2001010001020323-1231020032133321-1022223202113001-1222220202031303-2331032202300020-3032031333202302-3313111013102201-1013301303130000"></a>

#### `tgw_security.active_network_policies.network_policies.name` property

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

<a id="canonical-3231103000203331-1032023231010333-1221201311023113-0232101103212323-2230311131320313-0121222330212100-0222223220231330-1121203023213010"></a>

<a id="canonical-3223021122232133-3113211110121332-3331303303003212-2302300110021312-2123322130032332-2002323321332311-0302312201222220-3330200312013011"></a>

#### `tgw_security.active_network_policies.network_policies.namespace` property

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

<a id="canonical-0023021330101130-3110202032110331-3321013023003222-3012322110130030-3212033111010211-0032032303213112-3023330021223122-0033030101312023"></a>

<a id="canonical-3001121232122313-2102002221011232-2023120032211231-0213313232322213-1102313031122331-3211230021233202-0113121101002101-3123112231010333"></a>

#### `tgw_security.active_network_policies.network_policies.tenant` property

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

<a id="canonical-1212231303220233-2010113303122033-3011320331223131-0321033320220030-1220333031302100-2013330003103020-3132213201030302-3013202322330332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.east_west_service_policy_allow_all` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.east_west_service_policy_allow_all

<a id="canonical-2230001032112130-1221022223012331-1201031222002003-3213100210232013-3121010322112012-2203012101131333-1022333201302303-0001223010310023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for east west service policy allow all.

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
east_west_service_policy_allow_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201120002231301-0121123222011213-0222322100210132-2331203031331310-1021212323122132-2103312223133222-1200301331133103-3000320022330201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.forward_proxy_allow_all` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.forward_proxy_allow_all

<a id="canonical-2333003131113210-2222122020212330-3010313002130002-0013111320230030-0110201233302222-0313233211302130-1300131313312001-0332311230012302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023022012132131-2202121200230011-1120311221213111-2232023300131020-3021323023111122-3133100110121313-0221121132110210-2010132230010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_east_west_policy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.no_east_west_policy

<a id="canonical-0323023213030011-0012122230003031-3012313112310022-0233333100101312-2100201110023311-3003213002211131-3322021222131210-0123102200133322"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_east_west_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321021113032211-1112301111210302-2122022013230311-2310330302301131-1212323230212112-0001210222132001-2222022202023102-3331013003033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_forward_proxy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.no_forward_proxy

<a id="canonical-1001003212130100-3100301002123233-0301200133020200-3013032223303110-2003331220331303-0221130200210100-0201001200023223-3220230201321302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203031221233023-1200122111122211-2020320201211200-1312131320023211-3002302120132300-3203030331100323-3021220313112001-0230023220113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_network_policy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.no_network_policy

<a id="canonical-0332311301232102-1212012333323331-3110001311320100-0022223012133333-1130203331132211-1212322332103010-2003131031010210-2110111302021300"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311110330231022-3020003023311003-2213221230300113-1201311012233213-1111213120311303-1312303200031100-2033333131211220-2303001123013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- timeouts

<a id="canonical-1001330203023232-3222112201313300-2201220002003002-2031102013333200-0003203021120232-2121313033133330-0223323010302310-0320102032032131"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000013211233333-3102222312231121-0030330331312021-1001110001003300-0112321012212332-0103222023010302-2013211131223012-1110000102033011"></a>

### Direct properties for `timeouts`

<a id="canonical-2232310220122010-0111003333030323-3210311313211222-1321310321130222-0231112021112232-0330032323223330-3031031220300200-3000122012001020"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3033132130331010-2221120110131321-1113130133312200-1102122301300031-2111313121210311-2320303230120031-3211231121233212-3011320223123311"></a>

<a id="canonical-2121301232001220-0033232322320331-0220202331212203-1201021303310021-1320031130000121-2300023111211020-1232003221202220-1110030222310123"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0312201010021120-2010013212322031-2232301210030230-2030012131322313-3231312020103020-3133220302120113-0030003301220223-3321110002220020"></a>

<a id="canonical-0012012120333203-3001012320002132-0202302010021210-3231302202123032-2322230310111020-0103320232032220-1203130103122221-1003210032230322"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3211321230201002-1001303133330031-0212203131332103-2312020103030121-1320223003301333-3210312221102220-1111113022111301-3330121012111323"></a>

<a id="canonical-1232000110231333-1131001212222211-3232312210332021-3121012020200003-1203130232332102-0222130130022322-2132031121013210-3201110010221120"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- vn_config

<a id="canonical-0231213112103323-1313232113301233-0102211023023121-0112211110031022-1021302323123233-2322022103231322-2020110003210002-1321013101021322"></a>

Type: `"object"`. single nested block, Optional.

Virtual Network Configuration. Virtual Network Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
vn_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310223103213311-1232023030222300-3300311211333110-0210302312020330-0131201213022012-1312223132323202-0200202310320023-2300112322132323"></a>

### Direct properties for `vn_config`

- [allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213): complete subsection reference.

- [allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-3323120021313013-2101230331010010-2030321121100002-0302200130312330-3300220200133301-1211123132130100-3222211202010113-3231320002022332): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-1203200220100202-2301330300100333-1311122103213221-1131012011133010-0002123312102220-1112222133112321-0032111011031301-1201312233203130): complete subsection reference.

- [global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030): complete subsection reference.

- [inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130): complete subsection reference.

- [no_dc_cluster_group](resources--aws_tgw_site--reference--group-004.md#canonical-2011122312231230-3013232300312312-2202102000033103-3030221030323032-1020222321312033-1302222331002000-1132203010312230-3310010210310332): complete subsection reference.

- [no_global_network](resources--aws_tgw_site--reference--group-004.md#canonical-1220211023313210-3121131032130332-1202222301132133-1313020021300200-3101221300210022-2010132200330031-1332002033100321-3002120301331100): complete subsection reference.

- [no_inside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-0300132311300213-2213102221212301-1032330303012220-1032121031020202-0231221133312300-0101032232312332-3230133033132033-2301132320222332): complete subsection reference.

- [no_outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-0121330320232133-2321110122302131-1013232000113230-1212100032312110-2013322033302110-1203131120200211-3202133212002123-0302122002202023): complete subsection reference.

- [outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133): complete subsection reference.

- [sm_connection_public_ip](resources--aws_tgw_site--reference--group-004.md#canonical-3303020033221031-0121001232113232-0331101030300303-3303101011322302-1303033333020021-2321013112020223-0312203301333130-3001033131210333): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_tgw_site--reference--group-004.md#canonical-0021012122100010-3002323221023231-2002011020202100-1001130223230002-1100232013102230-1031022221333310-1130233020000232-2303021330003011): complete subsection reference.

<a id="canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.allowed_vip_port

<a id="canonical-0322123320232130-2232021210332110-2103000301200010-0103232333132211-0211131130313121-3223302030132130-0131213131230001-3321330030210213"></a>

Type: `"object"`. single nested block, Optional.

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202310210012201-2221301231322101-0101332032300120-0222203022020112-1332131232232202-3230123230030032-2102123201111030-1101213023200200"></a>

### Direct properties for `vn_config.allowed_vip_port`

- [custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-3221131113321332-3102323112131130-3300330122232331-3230201332010100-3102320020231131-1000233201321320-3032230022020103-3112211302031212): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-1323131123122310-3300232331110300-1313101312331110-0020023223022322-3022000032300232-1030001222121010-3332311210031001-1100311200313332): complete subsection reference.

- [use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-2301122231102012-1203132101231033-3210202223021133-1210331312102300-2111113220330011-2300010123321020-3000211010113130-0301112203203231): complete subsection reference.

- [use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-3133311202220330-0233223323100310-2111323013002320-2102213311200200-3002231323100230-3011023133013131-3212302031330030-1330033223123200): complete subsection reference.

- [use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-2330231312003021-3322020113120002-2023012313303111-2131321221202322-1101210313302022-0320213003200010-0213031032300030-1121100230120122): complete subsection reference.

<a id="canonical-3221131113321332-3102323112131130-3300330122232331-3230201332010100-3102320020231131-1000233201321320-3032230022020103-3112211302031212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.custom_ports` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213)
- vn_config.allowed_vip_port.custom_ports

<a id="canonical-2210221201311230-0020103212311020-1201232031223033-2012213110222102-2100010302003231-2301320313230130-0033311322002012-2323311222101111"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330102001203311-3203303121003031-0310331020130233-2300021302311223-0301121212322130-0020312310010333-1303323311311311-3113332110333003"></a>

### Direct properties for `vn_config.allowed_vip_port.custom_ports`

<a id="canonical-0022311211011332-1101201323233122-3001203011212301-3023121231033102-1003012223301011-3310330233300331-0103122321330100-1110020323320022"></a>

#### `vn_config.allowed_vip_port.custom_ports.port_ranges` property

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Provider validators and defaults (from schema source):

```go
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-1323131123122310-3300232331110300-1313101312331110-0020023223022322-3022000032300232-1030001222121010-3332311210031001-1100311200313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.disable_allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213)
- vn_config.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-3022123030022230-1111210330310001-3030131321200131-0021233110033031-2000113010213130-2032303033303133-2232312323310201-2331102033131202"></a>

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
disable_allowed_vip_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301122231102012-1203132101231033-3210202223021133-1210331312102300-2111113220330011-2300010123321020-3000211010113130-0301112203203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_http_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213)
- vn_config.allowed_vip_port.use_http_https_port

<a id="canonical-3022023121132231-0200213231322320-3111110023133120-2113012111320300-3131031232332220-2321230213113310-0211222300132123-0033120132331033"></a>

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
use_http_https_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133311202220330-0233223323100310-2111323013002320-2102213311200200-3002231323100230-3011023133013131-3212302031330030-1330033223123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_http_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213)
- vn_config.allowed_vip_port.use_http_port

<a id="canonical-1201032212201131-2230120010032220-3310331222230100-1323032023111332-1003102031323213-2312021120231202-2002011130311222-2322032012310223"></a>

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
use_http_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330231312003021-3322020113120002-2023012313303111-2131321221202322-1101210313302022-0320213003200010-0213031032300030-1121100230120122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213)
- vn_config.allowed_vip_port.use_https_port

<a id="canonical-0230221320330033-3303110111330223-0322013311303111-3231022330300230-1033103210111002-2330010123001023-2021032322123203-0233121133223302"></a>

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
use_https_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.allowed_vip_port_sli

<a id="canonical-2231110102000131-3330313330010311-1320020303321130-3132133101103312-0003300210222323-1321131220303331-3302200102221003-0102302010212310"></a>

Type: `"object"`. single nested block, Optional.

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231023323002311-0021013033323313-3102102320031113-2323101133223133-1320133001100230-3230130212312102-3310221110332332-2202032102231320"></a>

### Direct properties for `vn_config.allowed_vip_port_sli`

- [custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-3321310231011111-1012213322010001-2302023112230311-3201122333313213-0311112312230121-1121310022231133-0332012002322201-1131002032331331): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-0011323313332101-2210232100033200-0101311313123131-0233320130323320-0123323010021302-1033323303232122-2031000300132123-0303233312131002): complete subsection reference.

- [use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-2132023231131031-0223102322132030-0031131211032120-1231020100130002-1200300220130303-3301233303201201-3032320230123032-2311001302312101): complete subsection reference.

- [use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-1031103201331203-3233032210211133-1011210320212121-1220003211033312-2033332323122023-1330112311010220-1212011013332203-1312322133302123): complete subsection reference.

- [use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-3312102310012322-0310013233121303-1201313311313332-1231133013303133-0212213221310010-3222220203112312-1300303200312303-2322120211132121): complete subsection reference.

<a id="canonical-3321310231011111-1012213322010001-2302023112230311-3201122333313213-0311112312230121-1121310022231133-0332012002322201-1131002032331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.custom_ports` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="canonical-1023013121212221-0221013002310132-3010021331211023-0012313200022010-1123310102212122-3132010120123100-3323021203100122-0300211230011030"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122102121030202-3310312020301222-3330000123233030-1213133230231010-3202310221300311-2121330222213332-3212203012022102-0211133003221033"></a>

### Direct properties for `vn_config.allowed_vip_port_sli.custom_ports`

<a id="canonical-0031310213331010-2320003023310333-0001303110022113-2021202113310321-0301201320031203-3220202220233102-2302011132322032-2112210002203021"></a>

#### `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` property

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Provider validators and defaults (from schema source):

```go
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-0011323313332101-2210232100033200-0101311313123131-0233320130323320-0123323010021302-1033323303232122-2031000300132123-0303233312131002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221)
- vn_config.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-0222111103301311-0312013300223131-0001001031201200-0223211220332231-3103003223003231-0200100303210002-0203200320101330-3330011112232022"></a>

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
disable_allowed_vip_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132023231131031-0223102322132030-0031131211032120-1231020100130002-1200300220130303-3301233303201201-3032320230123032-2311001302312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_http_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221)
- vn_config.allowed_vip_port_sli.use_http_https_port

<a id="canonical-1123103322221201-0211100120100323-3321321031330102-1203332012122133-0102320210211323-2001113322313331-2302223200313212-3211232231123100"></a>

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
use_http_https_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031103201331203-3233032210211133-1011210320212121-1220003211033312-2033332323122023-1330112311010220-1212011013332203-1312322133302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_http_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221)
- vn_config.allowed_vip_port_sli.use_http_port

<a id="canonical-2030102232313113-0013232300302121-0323300101032333-1221211002230201-3201203220213221-2300120033221023-1003231020323320-1312031132013222"></a>

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
use_http_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312102310012322-0310013233121303-1201313311313332-1231133013303133-0212213221310010-3222220203112312-1300303200312303-2322120211132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221)
- vn_config.allowed_vip_port_sli.use_https_port

<a id="canonical-3023230202320031-1121030022120030-3023201311311132-1221111102021132-0000011022131321-3201301023123102-1012130213112301-0302011111022201"></a>

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
use_https_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323120021313013-2101230331010010-2030321121100002-0302200130312330-3300220200133301-1211123132130100-3222211202010113-3231320002022332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.dc_cluster_group_inside_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.dc_cluster_group_inside_vn

<a id="canonical-0302231131110100-3022103113313103-1000221313201011-0233001012322322-3201101310121333-0132223223110233-3020231002021101-2232103313211103"></a>

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313013012020033-2332122320313211-1222321300012332-2331123132110313-1001322201333002-1332322203320111-2302331020100103-3311112033123012"></a>

### Direct properties for `vn_config.dc_cluster_group_inside_vn`

<a id="canonical-0303221312332333-0323020113123133-3320201111221322-2332011100033031-0302223320021032-0231111003310132-2201101311232033-2130212021031233"></a>

#### `vn_config.dc_cluster_group_inside_vn.name` property

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

<a id="canonical-1122322202131232-3011131231101100-3230010311023232-0003210321102030-2301303312002211-1312211233333321-3133300021221311-3010030123121222"></a>

<a id="canonical-3010033032030331-0032113103113233-2133101212320132-2131122123030221-1101010231331332-3012103313101102-3200010130123221-2221203021333220"></a>

#### `vn_config.dc_cluster_group_inside_vn.namespace` property

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

<a id="canonical-0023130323012222-2323331220221033-0011230002310123-2222132120001132-3232122220232310-3210020303303010-1012231120012023-1122211021033301"></a>

<a id="canonical-0312121230123321-2212330233120120-1001102201311333-1132011231111120-1121033011300132-0000222231202131-1311212332133213-2131033202231213"></a>

#### `vn_config.dc_cluster_group_inside_vn.tenant` property

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

<a id="canonical-1203200220100202-2301330300100333-1311122103213221-1131012011133010-0002123312102220-1112222133112321-0032111011031301-1201312233203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.dc_cluster_group_outside_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.dc_cluster_group_outside_vn

<a id="canonical-0023211323100002-3332202102013120-0303203322221301-2130212010220330-0320321012323323-3101123001321320-1212023202000310-0033121230221100"></a>

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202020230223201-1320110012033303-1130033031132101-3202113220112232-2131330021302110-1003103002202031-3120223201120211-0212102021303110"></a>

### Direct properties for `vn_config.dc_cluster_group_outside_vn`

<a id="canonical-1203012212120023-2112013000220013-3002212012123012-1001222232313212-0223131121111323-1330133032330320-1102110230221030-3220013202120303"></a>

#### `vn_config.dc_cluster_group_outside_vn.name` property

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

<a id="canonical-0200200103332022-2230330011313300-2001322301231020-3311201012012132-3021200020021231-3101112313012211-2303030100213202-2000021131213030"></a>

<a id="canonical-3110103111023331-2003332210301231-2023233212032020-2133232300221021-2331232302320020-0113032311120203-2010301332213202-2012202333200123"></a>

#### `vn_config.dc_cluster_group_outside_vn.namespace` property

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

<a id="canonical-3122323220210220-1022320132213100-2031003120021131-2332133013111010-2110103111330313-1110102000000133-3011000023333301-0330230303001232"></a>

<a id="canonical-3003311310103211-1131033001001301-3013033031331131-2103223130022321-2333021330022212-2231132131001211-2101332332013020-0122323323010321"></a>

#### `vn_config.dc_cluster_group_outside_vn.tenant` property

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

<a id="canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.global_network_list

<a id="canonical-2123322011213233-2230010301230121-0221122302013301-1322333211123322-0232321210121111-2032312032112020-3212312201110210-0123212333122030"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311033201331031-3211313311011021-2113130020001003-3020202100230212-2312202132013202-3213013130301102-1302211220013220-2333222333222231"></a>

### Direct properties for `vn_config.global_network_list`

- [global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232): complete subsection reference.

<a id="canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030)
- vn_config.global_network_list.global_network_connections

<a id="canonical-0313322200021003-0111210213133302-0230112130003021-1322201030210320-3301010222222020-3313301302102013-3102121030233202-3011232020122222"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221323113010312-2032112233321020-2312231233230132-2032331012210223-1333020320211232-0331130000333210-0202122330130012-1133322333021231"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections`

- [sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-1323030132301113-1001220033133222-0222213000121020-1102103131200123-0013200000333010-1310021020300320-3032100002321120-2030322100201310): complete subsection reference.

- [slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-0323020221200102-1102230310311102-0303102120310010-3322203302013233-2000023012211003-1300222222332302-1313301000003201-2020211013222303): complete subsection reference.

<a id="canonical-1323030132301113-1001220033133222-0222213000121020-1102103131200123-0013200000333010-1310021020300320-3032100002321120-2030322100201310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.sli_to_global_dr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-1011001112231030-1002323311310123-3231030021330013-3321021021303333-3201032232100321-1230013102100111-0031131030013311-0120021001221032"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132201301003300-1010130012102010-1212122111301011-3100103301010030-3121312031030321-2101031022313030-3033300121002122-1010013130310231"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.sli_to_global_dr`

- [global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-2201300321012221-0320302313302323-2031111122012203-2111230320020133-1002313301013301-2013233130313212-3002301032000002-3223110233103213): complete subsection reference.

<a id="canonical-2201300321012221-0320302313302323-2031111122012203-2111230320020133-1002313301013301-2013233130313212-3002301032000002-3223110233103213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232)
- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-1323030132301113-1001220033133222-0222213000121020-1102103131200123-0013200000333010-1310021020300320-3032100002321120-2030322100201310)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-3211023022032301-1020022312131313-0211001301330310-1211001111013333-2230002030312200-2230332100032000-3123132230313003-3300222301023100"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301211121322110-0303130111010132-0103302101011033-1212313002213230-2000311111320321-0020031301033333-2233111133000200-3222320010120033"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn`

<a id="canonical-0203003332332300-3222233201333110-3113231210003211-3222211110221222-0133031211300222-0130121223301203-2102313002312330-2033223310112103"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` property

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

<a id="canonical-2201332000210120-2330130010120311-0313013222311230-1303122313003002-0100012012111322-2003133003003032-0111321013213300-0222201133120032"></a>

<a id="canonical-3311312222033133-0020310131220301-3312301020201122-3212210333301321-0121303302102200-2112203310203313-0310211231321323-1323310322020111"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` property

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

<a id="canonical-1212120011110023-0110110002300133-0120231320010312-3232011230113013-2310302003012023-3123302102122020-2131101023303223-0001211123022002"></a>

<a id="canonical-2333331232001022-2203000210131330-1033012000030131-0001210130222002-2123312213010012-0002221103233211-3122030132100233-0332312323023023"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` property

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

<a id="canonical-0323020221200102-1102230310311102-0303102120310010-3322203302013233-2000023012211003-1300222222332302-1313301000003201-2020211013222303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.slo_to_global_dr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1023020200122112-3230113212012111-0130232200313200-2310102111211232-1212301311130011-3223112101003211-2312121210312223-3323230323110333"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032111030033132-0121013210112030-0023202220100002-3001300020033111-2102203120011201-0203312032313113-1333120012001312-1221322010212100"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.slo_to_global_dr`

- [global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0012021201101102-1112102133231121-2100103023110030-0032013130110022-0313103222321121-1100112213323203-2000022233320113-3230111123032002): complete subsection reference.

<a id="canonical-0012021201101102-1112102133231121-2100103023110030-0032013130110022-0313103222321121-1100112213323203-2000022233320113-3230111123032002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-3020111000002321-0233001200031331-3203100211320223-1010113310201221-3113233011301211-1321112232003100-2320003022121100-2021100212310232)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-0323020221200102-1102230310311102-0303102120310010-3322203302013233-2000023012211003-1300222222332302-1313301000003201-2020211013222303)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0031233112311030-3010022211030331-1012202101332102-1123230312212320-3201001131012213-2333210020013020-0020033130123330-2121311330103102"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011310113210122-1001102030301102-2303032121302003-0110223103232311-3311321312310212-2323012310201332-2022323120210022-0111132132333201"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn`

<a id="canonical-0331003101300130-2132001111132313-0022012011123213-1210232000133200-3023010030020302-2112333113103333-3230000233303012-2011111002232021"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` property

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

<a id="canonical-3213031111202030-1023031333113110-1000301020303003-1100103010120212-1113011232123301-0302333333133301-3000221301001002-3330031213311121"></a>

<a id="canonical-1022101031123210-2020011121111130-1103211312001212-1110033112012010-3300121122303301-1021111330201323-0201130331021220-0031322322323112"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` property

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

<a id="canonical-0310121213312121-0030301312211122-1331111123232013-2100213121300311-1030233013022131-1110013232302202-1023333102221323-2000130113210023"></a>

<a id="canonical-3032112332012102-3200103032303112-1113123022201122-3220000322001221-1332222033223200-2122110101013300-1200112212131212-1333133110200113"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` property

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

<a id="canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.inside_static_routes

<a id="canonical-3221311030211011-3101022211012001-0313003223332012-3001012023313120-3210331333233013-3101030101012320-0032202023111133-3022213232321331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Additional upstream details:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112110323200000-3031322232333303-1100113211200331-2132031231203313-2231101233021232-1020130130212113-2222023302210302-0310120012112223"></a>

### Direct properties for `vn_config.inside_static_routes`

- [static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231): complete subsection reference.

<a id="canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- vn_config.inside_static_routes.static_route_list

<a id="canonical-2333012130203302-2023133122320133-1121312300121312-1022323101220013-1303010220222101-3312321010212210-1211111020010223-0012302200020132"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231110032330322-3030032202023222-3001031032000032-3312000000323031-0120002221013102-1302122220231311-1031002012023132-3030011112202200"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list`

- [custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220): complete subsection reference.

<a id="canonical-3321302013300302-0033030333200030-1101012320303102-2113301131010021-0032022331220031-1211121001203202-3213310232333132-3130111112321233"></a>

<a id="canonical-1330121031130131-1030100032220022-3020231330212310-2310011330200102-1203312300131132-2213020131203020-2023003031013311-0130201120311023"></a>

#### `vn_config.inside_static_routes.static_route_list.simple_static_route` property

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- vn_config.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-0110010120123213-3312121231010122-0320121130113033-3000023310103322-0010321302220223-1012232011020103-0012101000233132-0002210110330201"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103300221121003-2230122301220211-3130012233032101-2111100112203033-2233311222222011-1110232012121231-1321302131001102-1222031231000111"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route`

<a id="canonical-1113022302212132-2230120030303030-0210201103303131-0323122021022220-2322022211011130-2333123201211210-2310132201210132-2311323030313200"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_tgw_site--reference--group-003.md#canonical-0303223121101303-1111122111002301-3120322232223212-0303320111130120-2213003211100111-3110033103331313-2031022032313003-2003220232010022): complete subsection reference.

- [nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200): complete subsection reference.

- [subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2312102122133200-0120302223320132-0133013231100200-0201333332120223-3022330101331300-0311113223000130-0123020123002012-1012331331330010): complete subsection reference.

<a id="canonical-0303223121101303-1111122111002301-3120322232223212-0303320111130120-2213003211100111-3110033103331313-2031022032313003-2003220232010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- vn_config.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1201003121233210-3212202011021322-1013322223230221-2113120231021100-3210200320231001-2311201320312010-0222111033302032-2012123101021302"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0122110313201220-0223220301211310-1223321130332331-1211033012120212-2210301123221012-1322020233311302-1122131322113201-2133113120220220"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300121313222033-1021002133332112-3221120332220333-1022312212023132-2102220010302020-2021222301210200-2013131200223002-3312103321121032"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop`

- [interface](resources--aws_tgw_site--reference--group-003.md#canonical-0232113203110221-0131020122132313-1113223131023030-3020111133223123-1233133220132223-2203333032333320-1133031130202222-1031000330303112): complete subsection reference.

- [nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301): complete subsection reference.

<a id="canonical-1121332020110101-3001211000302033-2213211133203212-0200033203120133-2122010301121310-2020330021123123-3113022333120102-1010330222113103"></a>

<a id="canonical-3112012030112321-0002332122232012-3130212111132002-1233331230301231-2332202311010321-1213213201120133-0100313231331223-0021223113001130"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232113203110221-0131020122132313-1113223131023030-3020111133223123-1233133220132223-2203333032333320-1133031130202222-1031000330303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3010100223133102-3001321222130321-2222002003212022-1101322312011221-2233203202030121-0133120012121222-2111323202300123-3113231231321011"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Additional upstream details:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302302023002010-3122222311332030-1223101332301301-0103121131012313-1223222200303201-3233112203222021-2233323033130330-0033203132033333"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface`

<a id="canonical-3030000113003331-3222000213101101-0200103021212133-1211101021031212-2222122221130200-3303033000320123-2313211323132203-2130113203332102"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-1000213032310221-3202221331002231-1130032200332033-0012003333310311-0310320232220310-2103113133331321-2313002223123230-2021132001223121"></a>

<a id="canonical-3031122210002133-1301030000110023-0331332120013201-1120321221233120-1010312321021121-3332133033013002-1330200213312111-0012120003233300"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0200022330112330-2223321001212030-3230203122312211-3201031103201021-1231200203310203-3320033130201313-2130220000221013-1113130210232213"></a>

<a id="canonical-3101003120323113-2221133003220213-2120012031302032-1213330323311010-2322201312023330-0022033122310021-2121212120202112-0323130100001011"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
  }
}
```

<a id="canonical-3031102320111211-2330322010113103-3322302302113000-2011322333031323-1131101023030212-2300012322102301-0332100131020133-2312311303003301"></a>

<a id="canonical-3233110102332333-1200010032032112-3103030313000231-3230030030231211-3023211002132331-2130100210313302-3322333132203120-2120333220322321"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3221011100310021-2203333030011200-0312303303222022-1232101311203002-1011012100031020-2220313011212003-1302232003212301-1230230213002032"></a>

<a id="canonical-3121131330201330-2332022322121131-1231100033333121-3230031001220210-3232332201323012-0313301122223022-1023103310333210-0231022211231012"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1233313132101102-0202022011132222-3320231111210133-2300221123301233-3220102201132032-3221121310200013-2130212330303202-2013330112020231"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310230111320010-3213032320220001-2022010212302103-3122012132031001-0023012233020120-1000001110302030-0031112302131011-1133010220313012"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address`

- [dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-2300033023220010-2013031003211302-2300123003311122-0230203310332020-0021221020023131-1103221232133030-3120222130332310-0201110332021132): complete subsection reference.

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-0333133020303332-1323011201230012-2311303101010232-3211213101223311-2201123212232313-3233030020211010-1233332002232021-1030310120101212): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-1310330231031000-2130012120001210-1123023320201133-2223030113212011-1030333232311221-1102112023213320-3222211112103002-1323032211012122): complete subsection reference.

<a id="canonical-2300033023220010-2013031003211302-2300123003311122-0230203310332020-0021221020023131-1103221232133030-3120222130332310-0201110332021132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2300123323101133-0333233121212133-1323121113332300-0200012221230221-2132203130021110-1131112322211333-1322223333330132-1312300113203211"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121102210220101-3031012010332202-3330201200331330-1103120033301330-0232321300023103-0033220321201311-0102100000201120-3330311001330013"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack`

- [IPv4](resources--aws_tgw_site--reference--group-003.md#canonical-2123032203221322-2330122200130232-1201213000121321-3233021133230122-3030021011223322-1201202311113310-0012322231330231-2223302010310320): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-003.md#canonical-0022301031001102-2210102010213333-2020221110302200-3130022322013323-0300122022013330-1322130223200001-0100021121203223-3223001133210002): complete subsection reference.

<a id="canonical-2123032203221322-2330122200130232-1201213000121321-3233021133230122-3030021011223322-1201202311113310-0012322231330231-2223302010310320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-2300033023220010-2013031003211302-2300123003311122-0230203310332020-0021221020023131-1103221232133030-3120222130332310-0201110332021132)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3322212033133201-1300120303002122-3331102013210213-2332102123212213-0223012011211222-2032331202213303-0312331020303102-1201230201312113"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023110322320010-2310001101113313-3101121313100223-1200323310331311-3331103312102232-2103210300122030-2233221313130201-3323320002121202"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-2111002213230113-1020003211321131-2221221222303322-3330032322322121-1312211311223023-2123033301330323-1020322210123211-1210010222331331"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0022301031001102-2210102010213333-2020221110302200-3130022322013323-0300122022013330-1322130223200001-0100021121203223-3223001133210002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-2300033023220010-2013031003211302-2300123003311122-0230203310332020-0021221020023131-1103221232133030-3120222130332310-0201110332021132)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-1212131333032231-1121011031221110-3322011330203330-1330103321201032-1231203022030110-1220111102221123-0012132303100323-3223230301213221"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133011232000010-3002210231313101-1021202331100100-1003100011100123-2013302220231021-2200231200003101-1003320020122212-3212000331121113"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-0203101322303211-1003112120331221-0000120110012032-0301011030221211-1222233102200311-2232330131100110-0110212032110312-3331320011111111"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
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
