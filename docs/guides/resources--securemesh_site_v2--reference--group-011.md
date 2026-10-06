---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3031232102113112-3012030032303110-3113232311311203-0303121312100000-0102211302002122-3331132310131121-1113001003200322-2000321311333123"></a>

## `local_vrf.slo_config.vip` property

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-3010131300322320-1003302333332300-1020211101201223-1330230321111000-3302100331030033-2131211010320101-0113010111321222-3110323122220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.no_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.no_static_routes

<a id="canonical-3232121310133000-2320211000131013-2103112100330113-3010310300301101-3112112032121321-2233010112302102-3233300213031313-1011120332203030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012230222212130-0202203112130311-0010030110131100-0230201032312322-0332203222000101-2001310101203203-2120032012002122-3001033010123112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.no_v6_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-2102330223033123-2221322213230122-3003302300012111-3032203132221100-1222110300020133-0300013211230310-2322223322333232-1200101021331033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.static_routes

<a id="canonical-1321212023130110-0212000110101322-0310101102113012-0202301133200112-3313033200122013-0202323232333032-2231310201122331-1213031300030222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033320021122132-0301022221033300-2110330030131320-2310203312232213-3130233313012300-3111023220213310-0322011021100332-1311011000003123"></a>

### Direct properties for `local_vrf.slo_config.static_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200): complete subsection reference.

<a id="canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-1032020232102223-3221010302110022-2031021002310332-0000002012222123-2102233103231101-1100110212200002-2323103121102222-2111102200020112"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001200010221303-2330023313133202-1102300013312001-1032123313131233-1011322202022313-3003120112003033-2223122033330101-1133231002230102"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes`

<a id="canonical-0201022233020313-1102200132030322-2011003131301301-3101330132012001-2211120131323230-0330113330123120-0022020130233030-0103231121003132"></a>

#### `local_vrf.slo_config.static_routes.static_routes.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-1010313030211131-2011012032320033-2033001112122133-3112003321210322-1112323103311221-0313222030300213-1232231011011031-1132032100211323): complete subsection reference.

<a id="canonical-0201323221111302-0220033201002320-1102102031001000-3002203001330230-3321200213122112-2203323121020013-1220123233010320-2020310031202302"></a>

<a id="canonical-1310012302023122-3231233103132212-1110110331221130-1212302202013011-1021303130001113-2101011103221212-2002233021221310-3033311002011313"></a>

#### `local_vrf.slo_config.static_routes.static_routes.ip_address` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-3122213110000311-3100021030322203-0333331200111233-3002300202032223-2112320330233211-3002033232230212-3302002222103133-1023011012102221"></a>

<a id="canonical-1003113220211103-1031202022301131-0000132232330120-1023002101103133-3110030102021030-1312311301220333-1221331032221322-3123002121220223"></a>

#### `local_vrf.slo_config.static_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122): complete subsection reference.

<a id="canonical-1010313030211131-2011012032320033-2033001112122133-3112003321210322-1112323103311221-0313222030300213-1232231011011031-1132032100211323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-2103312212110233-1203202030300230-3003200013331223-3030203012212000-1303023023330100-3323113012033021-0002213211332221-1033333220133331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-0133333012103113-0321122203233210-3011010222020313-0120321220120212-1110331221213210-0203023222322132-1211033123321023-0021033312232103"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220032033303232-2120301320102322-3313230211230021-1313003103332310-2030302110133330-3012311001223132-1330330300001211-0020121133202200"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface`

- [list](resources--securemesh_site_v2--reference--group-011.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113): complete subsection reference.

<a id="canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-0211311020320032-1012301333133023-1122231122203202-0300332331012213-3200131120132331-0213001023233000-0212231223131332-3130100303033023"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221121303031300-3313302221123110-1232002233023311-2033322303123120-2000131021331112-3030310310211211-3313232321322110-3003102223011033"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-011.md#canonical-2302111313203322-0232000210330122-3013320321311122-0131330222331230-3202313302113233-3313213023223332-1102200201110123-2211200203333332): complete subsection reference.

<a id="canonical-3021320000312222-0001103332330300-0222203210210033-3333000220022132-1131233021302120-2321132102201203-1103111230301113-3023300321220200"></a>

<a id="canonical-0030333100030003-3002313203201303-1233321233211011-3212012113030102-0102211203003102-0323112210011121-3103230233311111-3331120031332020"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Optional.

Node. Node name on this site.

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

<a id="canonical-2302111313203322-0232000210330122-3013320321311122-0131330222331230-3202313302113233-3313213023223332-1102200201110123-2211200203333332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-0310231231001113-3201101320022232-3213101202000100-0011020220103002-3323000020320020-0231303020133001-3200111323001103-3002230332321322"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

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

<a id="canonical-2222100002021231-3030200113323211-2102233020002321-3313311333001002-2033322200111203-0321032122123233-0323031122331100-3133131111301033"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface`

<a id="canonical-1203000202031313-1232203120020023-1100233203110020-0223230011021332-2113202011310132-3322313100133100-1311312003220123-1133020232311120"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.kind` property

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

<a id="canonical-2020312332211033-2203202021230131-1120220232322332-3313223201301113-3302300130132223-1323022320123312-1300232211200332-3101100111313131"></a>

<a id="canonical-0021222132101233-3113213102031311-0110012012033033-0210101021112323-3003202112223223-1230003002112031-3022020210122002-3302233211231323"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.name` property

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

<a id="canonical-0331010300102110-1022210123202130-1021031320033220-2303021213220000-1012222133113111-1230033310202321-3111302322122202-0211300203222012"></a>

<a id="canonical-3132310101101003-2302203002031301-3002302122030133-2113212032113013-1222113210002030-3132222303012320-3111031222303022-1203321113213303"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2302222202212230-0213013321001303-1300111301113213-0023002031303212-0322311210230130-0200320203223110-3233223103103030-2113310313000213"></a>

<a id="canonical-0232030232123302-0112220232331312-2220013203110131-2332320310212223-0311132032113220-0220301323312012-0120223021101133-3132101101233123"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` property

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

<a id="canonical-2232232220121020-1130230021301032-0021101132113020-3110022201332323-1020013331301133-0321030223332111-1123312100232123-3331000303102031"></a>

<a id="canonical-2120033320032020-3202322212101012-2023003331310301-3022320132320122-3102321113001020-0032131121333132-1232000231112233-0231023322121302"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.uid` property

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

<a id="canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-0132323220130201-2200123312322222-0120201133210122-2333331103000323-3232100022230333-2001323220020303-3211101120310332-1311023200313210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Additional upstream details:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203222010102013-3323030302301130-2021010221011103-1133203321233101-1110311331101120-3030333201130000-2100123213311133-2103233103333121"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202): complete subsection reference.

<a id="canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-1103112031031313-3011221211032023-1112111203013212-1313313323120210-2210321301022012-0001331101103203-2102111311021223-0130222233231320"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321020332210122-0032111122112000-1000131232202302-2020302231212132-0102103122221303-2012111020211221-1033312123110232-3220111002302100"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes`

<a id="canonical-2320101101101002-0231223100022210-1011132330322112-1210130120233200-2231212331221211-1310233320213020-0122333221020102-0310031230231010"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-3212031133321311-1231003013302211-2112031202020320-3030232001033202-2313212312031333-3100321031123330-0023221313101222-1201201220312101): complete subsection reference.

<a id="canonical-0333212202020102-0223302303120301-1002112031200300-1031232122112300-1222120302103331-3031330100220322-1320331131110100-0022312123332013"></a>

<a id="canonical-2330211301211203-2001101222213220-1122221201231300-2220110132320202-1011203111032210-0303220030020102-0000321330311222-0101301112021322"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.ip_address` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-3222023231220230-3310100032000102-0030230110231113-1102103213131031-0222222200332101-1030312313021102-3301313301110330-1320221331332010"></a>

<a id="canonical-2002123312032013-1313003011211220-0023130130320313-3330113302003010-2112133210031220-0302030321001312-3213333320131311-2100022310211101"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130): complete subsection reference.

<a id="canonical-3212031133321311-1231003013302211-2112031202020320-3030232001033202-2313212312031333-3100321031123330-0023221313101222-1201201220312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1313001231132101-3303211333312331-2203032322300033-3120111122231313-1023303032202310-3233132200032002-0001130331122221-1310330100110332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-2313301111001102-2230112300320213-1011023102220030-3133113021122120-3023322120130303-3012111023232333-1001023230032232-0123001103000312"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322221302201030-0003233213100031-1012000133220312-2123011121233332-0100212001123312-0111332313330000-1133200300123232-3211122020031001"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface`

- [list](resources--securemesh_site_v2--reference--group-011.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330): complete subsection reference.

<a id="canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2030302101310233-2103300220122021-3112200122032011-0023110332323102-3202321201223003-3310213200023333-2301103132023132-0122312213321003"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032131131220103-1011230120101230-1323302031330333-2030202203320032-2213300022101230-1322312020123020-3031223032311333-0323111333022033"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-011.md#canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330): complete subsection reference.

<a id="canonical-2322332230210310-3322100110211320-2311103113101200-2121303022020233-1022132311213120-2130031013001211-1021303033222010-3000322232200020"></a>

<a id="canonical-1020113001133123-2222203220111032-1201313122033110-0330323211122132-2330003303022223-0130302023203102-2311223113133333-3131310311003321"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Optional.

Node. Node name on this site.

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

<a id="canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-0321003312313320-1123311233202322-2301001101032031-1010303212330122-3111310220031130-0211201022101310-1332011222121223-3300301112303332"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

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

<a id="canonical-0130013120102112-1221200311102011-0020012220023332-3020033131000233-0103310230033031-3201103000133331-0000300022320232-2331313221221011"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-1000231220311020-0213220001112223-0323221201110122-2232311102010323-0222330231020332-2331233033113030-1032212130333320-0232002331111023"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

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

<a id="canonical-1132332233123003-2212201132101212-1003102131323322-2120013223330323-0000332111112013-3131200012031002-2002031210010330-3230333130020110"></a>

<a id="canonical-2010133332012001-3121332101322310-0130210031301131-1312111213010010-1023012101021203-1202213222230332-1001102313230113-1000220333113333"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

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

<a id="canonical-0323320010003222-2202331203132013-0120113211230100-3033030212021122-2203210012232021-3322231102210202-0322133222112322-1030321003333222"></a>

<a id="canonical-0212230202322200-0221011231212001-1130102210110023-3103001031323302-3032111111312331-2333212300220013-3130012332103123-1201001210001033"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0022233200332313-2120302200002012-1200202233031002-0301220233303231-2230103001000200-3101023230311333-3022321103021111-0210201131012112"></a>

<a id="canonical-0301021030230322-2211333323200203-0301123130220323-2121300112312022-3233220033113313-1300332032320011-0211011231123010-3102011101033321"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

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

<a id="canonical-3201211231221112-0231321132012000-3003031201210220-1100020200331130-1233120220113123-1312100100311102-0322130300011021-2030231212123300"></a>

<a id="canonical-2223221002131113-0023130303301201-1232013032032330-1101231110013011-3021213211233310-0030023012010311-1223211310030221-3313122111203300"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

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

<a id="canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- log_receiver_with_net

<a id="canonical-1222131012023103-2300202303002133-3112212323102320-1230201032232133-2332123111310012-0333132022200221-2203011330112320-2223302000230020"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("use_management_network",
    "use_slo_sli")}
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
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-011.md#canonical-1222131012023103-2300202303002133-3112212323102320-1230201032232133-2332123111310012-0333132022200221-2203011330112320-2223302000230020)
- [logs_streaming_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-0111001211331333-3333100033323212-2110200312112313-2301000232032002-0331312312200211-1233122002101233-2323212221312020-1333023010232331)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver_with_net {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121310033000313-0022032301310220-2001332130111003-1320222223001333-2303212113011001-2000023200211002-0000332300121033-3323323133230211"></a>

### Direct properties for `log_receiver_with_net`

- [log_receiver](resources--securemesh_site_v2--reference--group-011.md#canonical-1120012013013123-3022211010330233-3002013101202231-2213130010033121-3321013301122032-1100011113202113-0010303022202231-1003123003220200): complete subsection reference.

- [use_management_network](resources--securemesh_site_v2--reference--group-011.md#canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220): complete subsection reference.

- [use_slo_sli](resources--securemesh_site_v2--reference--group-011.md#canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230): complete subsection reference.

<a id="canonical-1120012013013123-3022211010330233-3002013101202231-2213130010033121-3321013301122032-1100011113202113-0010303022202231-1003123003220200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.log_receiver` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-011.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.log_receiver

<a id="canonical-3330103200300113-2010313112323020-3011202121021303-2200223220233222-1030223101122202-0022302311103032-3033033011121113-1031111200131121"></a>

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
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301002330202021-3031103331103310-2211231023223113-1300313022000312-3001022020113203-0030201112213213-1322221133211101-3023230022321332"></a>

### Direct properties for `log_receiver_with_net.log_receiver`

<a id="canonical-3231032321122013-2000303131120300-2122300022211023-3022333131203011-0203212231302233-1001321210023302-3122111012103003-2323113110013210"></a>

#### `log_receiver_with_net.log_receiver.name` property

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

<a id="canonical-1213120300332313-2330131031330231-0121021231311001-2321011201010121-2330033130100222-3231300201032031-0031230330232301-2211002321321103"></a>

<a id="canonical-2323320101113011-3222310122332301-3202231333203221-1231323020133111-3133230001012211-1021200130013330-1332222222202301-2330022310333331"></a>

#### `log_receiver_with_net.log_receiver.namespace` property

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

<a id="canonical-3012310101132323-0323331133012022-3323330232210103-0123030230331022-1332011300213320-2110310112013210-3122201302112023-3111311200232110"></a>

<a id="canonical-0023310210111110-3131330032233120-0023132312101232-2201100001301031-1233302213101300-0030013001132330-3303103013001212-1013223101102212"></a>

#### `log_receiver_with_net.log_receiver.tenant` property

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

<a id="canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.use_management_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-011.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.use_management_network

<a id="canonical-2033320112213113-3132213100330323-2202012010033003-0322311102120110-3023220123020300-2303113010231131-0200232233003001-2133332201101102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use management network.

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
use_management_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.use_slo_sli` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-011.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.use_slo_sli

<a id="canonical-1223031322103113-0031222033011113-3130120030202213-0321001022302333-0302331001120321-2323201302112202-3033011101232022-1300022001203033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use slo sli.

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
use_slo_sli = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210233313020231-0310131332331201-3031332230130222-0133101033233310-0223230331130213-2103112112102321-1230002312101232-1211013133212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- logs_streaming_disabled

<a id="canonical-0111001211331333-3333100033323212-2110200312112313-2301000232032002-0331312312200211-1233122002101233-2323212221312020-1333023010232331"></a>

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
logs_streaming_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231001003211033-1113201231132322-0013121220032323-2031203003230221-1133223030313302-3131212323231233-1301021322010201-1030310211331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_forward_proxy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_forward_proxy

<a id="canonical-1322030320001001-0330100333213223-0220321230132233-2313313023123323-2030302332103211-2033330212203300-3012021231012103-3230222113113031"></a>

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

<a id="canonical-1111331220330301-2030301032232003-0321010121010110-1033012330011003-3321221320333313-2222101321233313-3201213303322110-3011030103300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_network_policy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_network_policy

<a id="canonical-2213010322031120-2120333011333213-0322320013333003-3120023210232013-0222210010203123-1032320102313113-1233002003220122-1232110111321031"></a>

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

<a id="canonical-3123331333122110-3122333032323313-1322031023033330-3022013201322102-1030322011123202-1032113110203010-1103213002302323-3011223220033322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_proxy_bypass` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_proxy_bypass

<a id="canonical-3022023002231032-3101230312322203-1210102303113020-3133121101220113-3332021101223101-0031033231310313-2221010131313222-1131033300021022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no proxy bypass.

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
no_proxy_bypass = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322012222010332-2230100331221211-0132203200323130-2003020232110130-0012210112022001-1220312313020321-1303222003211323-2020120133231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_s2s_connectivity_sli` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_s2s_connectivity_sli

<a id="canonical-0330303002222222-1213002233022103-0000311222000331-1230302330001022-0022102302021300-2031322212021101-0113011320113023-0212000011230312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity sli.

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
no_s2s_connectivity_sli = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032112221231302-2123032103000023-1212321030120112-1301100133021333-3122002203103112-2022222012000331-0333232203322311-3201133120130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_s2s_connectivity_slo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_s2s_connectivity_slo

<a id="canonical-1030303003312203-3112321311031221-2310332332000120-1233313121311313-2111010213232133-2302111301212233-2211333200111312-1310302310103121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity slo.

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
no_s2s_connectivity_slo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- nutanix

<a id="canonical-1210012123223033-3213322002010300-1021212010202320-1123300232210200-2332210120200203-3222221311223100-3020212001233310-2032320102120102"></a>

Type: `"object"`. single nested block, Optional.

Nutanix Provider Type. Nutanix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
nutanix {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200312212032012-3130021302132010-0300101013310221-0221032101211321-0033313002021002-1001201330131221-1020011223233302-0310112103131312"></a>

### Direct properties for `nutanix`

- [not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000): complete subsection reference.

<a id="canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- nutanix.not_managed

<a id="canonical-3332103223333133-1331303223121321-0101200130020110-0123001130303122-1221300203223210-0121012031121122-2220132330303212-2000101111011231"></a>

Type: `"object"`. single nested block, Optional.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233331103031300-2232110122122132-0001103231020001-3031223313300030-2020112103021133-3102311001332133-0222030203000323-2111120303230203"></a>

### Direct properties for `nutanix.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231): complete subsection reference.

<a id="canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- nutanix.not_managed.node_list

<a id="canonical-1210031322321012-3222110033003100-0212011023002023-1223010201003232-3023100121212113-0003211310331302-0320022220132331-1323310012111122"></a>

Type: `"object"`. list nested block, Optional.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220211213332313-0200322312313030-0131110201033211-3333023021112231-0223220121322232-1000110120312110-0002110320032333-3033211311101113"></a>

### Direct properties for `nutanix.not_managed.node_list`

<a id="canonical-3201222221111330-0132222103330002-0312230333203021-3101122003232112-3101020211112200-0121003003323111-1332001300303301-3233111121113132"></a>

#### `nutanix.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202): complete subsection reference.

<a id="canonical-3121223103200210-0212302310101201-2232102210013313-1223031203010322-1100302110001311-2232332320312323-1031032233223000-0322223132302112"></a>

<a id="canonical-3020010121101231-0331300322213213-1300201231202302-1333022111200230-2210212323021310-1202011021101211-1303002113033332-3030000101200123"></a>

#### `nutanix.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0000300030011020-1301120123112320-0311033212112210-3322123122233112-2133113230113331-3330030211221203-2321333121221311-3331130021230033"></a>

<a id="canonical-3131123331222033-3121121222002200-2231003210232212-2323310001111331-2132131103221013-1321203203113011-0003030221123030-0203232230012232"></a>

#### `nutanix.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- nutanix.not_managed.node_list.interface_list

<a id="canonical-2321302213113003-0000212013222013-1021312112322210-3000111201312312-0110101012210203-2121002310003332-2130312001200220-1121300212111112"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032113011120003-2120310100201323-3003101230020213-3230030220103220-3012223223022233-2121003023033302-3302303322333121-3322131201333011"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312): complete subsection reference.

<a id="canonical-0010312010130023-2323303130223323-2120301303230321-2323311000303202-2312000112002313-2103012212212300-2302131200203033-3210310300130102"></a>

<a id="canonical-0323303211311012-1021230221231113-1200122232303121-1003323232122232-0031223303302221-3301222130103003-3011023312230123-0331123200022022"></a>

#### `nutanix.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-012.md#canonical-0122301031222001-3012001030032031-1220230130013231-2113030122330033-3223230311003133-2330110130310303-1103123101201022-1002300311110013): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-1220303103223201-0002300332200330-0210220311211213-3220113033021232-1111312303002001-1030223331130300-3322301012231302-2331321131003323): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202): complete subsection reference.

<a id="canonical-3101313001320113-0130203231201300-2111300231102020-0232113013101313-2232013303121131-3022201100020333-1130112300212022-1020311332313332"></a>

<a id="canonical-2211301222010002-1123110301021203-1111333200003032-1302103320210213-2113001221112313-3021313023121110-2331131030113320-1122332032130131"></a>

#### `nutanix.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3130022230312230-1112220311321101-2031300230100311-1223031111000133-1021332001101030-2003133122202101-3221231230313211-0213012233331223"></a>

<a id="canonical-1131033113303020-0130003222120200-1133311232233130-3113310333300033-3202300303220011-2332023202130132-1320323132213310-3213232102223023"></a>

#### `nutanix.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1022332031010300-1100001212003330-1223231012002103-1220121111110222-2022301201200020-1203201023033020-1332131302223332-2221112132231222"></a>

<a id="canonical-2321020030130312-3002113120232131-0121300030223111-1230311210121020-0312113032323002-3022322202211202-0222010010303001-0033320331212102"></a>

#### `nutanix.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

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
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-012.md#canonical-3123320032100001-1031133131223231-0011130301203201-1123122201301021-0120221213231131-1110333302111203-0032120301303121-3232123031302301): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-012.md#canonical-3123322001322102-3101003102311001-3021021212103320-1211232031000131-3021021232200112-1233221230310332-3303221022112021-3131022012122013): complete subsection reference.

<a id="canonical-0133112101030022-1231321122311313-3010232200323131-2211231323100212-2323111320330003-1321331012322130-3012010123022221-3023220320300220"></a>
