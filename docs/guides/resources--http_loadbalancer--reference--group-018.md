---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2302123202312111-3301030131020023-3202233131312000-2222221101000120-1322230133012023-1002033133020302-1320031011320232-0301210013321231"></a>

## default_pool_list.pools.pool — pool / 320311312233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.pool

<a id="canonical-1321233303001203-2111130000110322-0211222220122211-2123023121132032-1032320012321220-0131303212011332-1101301023112013-0003013300110320"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133000313130031-0212102100203203-3120023023113011-2023121032133012-3022021222032220-3333020030002311-0120222221330223-1132232031010011"></a>

## Direct properties — pool / 320311312233 / 3

<a id="canonical-1103222010120231-0100200233110002-1220013113330203-3333210122000023-1221312202221212-3032031003201223-3302103010103000-3223001323303211"></a>

<a id="canonical-3130310303233111-1113032000331331-3331100000103113-1002121331300021-0123210033021010-3033100122310201-1303132121000203-0021211101322100"></a>

## name property — pool / 320311312233 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2100233002120313-3201030132212123-0133331233101330-3022112132212221-0231332211031121-2323002013110211-0113131301203220-0330011120102133"></a>

<a id="canonical-0022300332211312-0031221101221220-3033112330030201-2210302132032332-1002323011312221-3223301020210111-3310012322011232-3330332220131100"></a>

## namespace property — pool / 320311312233 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3310123313033112-2332030321211130-2022230010133023-2313000120231121-2332002201223320-0113030131130032-2310231032122020-3222230320203322"></a>

<a id="canonical-2132220200110113-1200233220321210-0110212232313302-2223322300030031-0122233030221003-0230020132222210-3213302210010002-3333031110003032"></a>

## tenant property — pool / 320311312233 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0030302123331121-2023230220003030-1003301022311313-0132213330200210-0300223230132100-2002001031202021-0011300113032101-1121203003111302"></a>

## Next pages — pool / 320311312233 / 7

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111211333211322-2011031300210312-2200120222113021-3221311301201202-1320123221120321-1300131112132100-1012310023303323-1021123303322101"></a>

## default_route_pools — default_route_pools / 220031001233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_route_pools

<a id="canonical-2030002000113130-1013213300122022-0332132120103002-3113212002131130-0011121312031211-0300001031000233-1333033103010131-3022331121021300"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools used when no route is specified (default route).

Upstream description:

Origin Pools used when no route is specified (default route)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_route_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323132220313230-2322132032303220-1033113033001003-3121011010333303-0231332223302213-1011112021203130-2132313321122023-1113221203202302"></a>

## Direct properties — default_route_pools / 220031001233 / 3

- [cluster](resources--http_loadbalancer--reference--group-018.md#canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-018.md#canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-018.md#canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000): complete subsection reference.

<a id="canonical-0330313201303120-1120121031310101-2210120002330330-1110232011202121-2330200231230010-1130003330010010-3120322130131301-3033000121220021"></a>

<a id="canonical-3200132213031302-1111031322331322-2013231313021023-0313312213302002-2210222033311202-1310000121302312-1130033202230230-3131313330010313"></a>

## priority property — default_route_pools / 220031001233 / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3133330301132011-3110200033223102-3003102020133323-0222120231200131-3033032131111222-3032200323331021-2100103232101121-0022232200012203"></a>

<a id="canonical-0021030021131312-1011232002132133-0023201132220132-0323332303231312-1303133321200032-3021001213012301-2110210123323223-0211123322132013"></a>

## weight property — default_route_pools / 220031001233 / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020303210211031-1010323302101001-2033212032030212-1310301220133331-1022001203212221-2010001010000222-3311213100200320-1303321021302222"></a>

## Next pages — default_route_pools / 220031001233 / 6

- [default_route_pools.cluster](resources--http_loadbalancer--reference--group-018.md#canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332)
- [default_route_pools.endpoint_subsets](resources--http_loadbalancer--reference--group-018.md#canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321)
- [default_route_pools.pool](resources--http_loadbalancer--reference--group-018.md#canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022133022132112-0332300221313020-3010002331322323-3200212330312322-0321223003220202-2021202131300321-2300022010221231-2322223101012130"></a>

## default_route_pools.cluster — cluster / 113123210133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.cluster

<a id="canonical-2002221002211212-1223131201203133-2220103133110301-1121220022322222-2023301320102130-2132122333112203-0113111212121101-2311200332220212"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001222232323231-2133210032020233-0023010132120010-3230333223102321-1023113322211122-3301201121012313-2311103023030331-0322311201031202"></a>

## Direct properties — cluster / 113123210133 / 3

<a id="canonical-2321130310233212-3011120101131101-0122311321202322-1021012232310321-1313330103231211-2212121312003322-1331100331302010-3111232302131313"></a>

<a id="canonical-3123223223301132-0120001023122331-0320112213011320-2231022223300100-0300323311311232-3022032202021312-2330312021311023-3130011331020102"></a>

## name property — cluster / 113123210133 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2132021200000231-0203102312233100-3021313333113112-2202332213211223-3101300232322222-0332332222011122-2000003003323323-0303033001301231"></a>

<a id="canonical-0020121101300003-2121033303330302-1210022111302112-3320012100311122-0211120203032003-0013011010030032-0103103131101021-1332331132023330"></a>

## namespace property — cluster / 113123210133 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2102211032310201-1130032233102003-1021203102210223-1112333023310123-1133012101230322-1323311102023132-1023102312213322-3133110103122021"></a>

<a id="canonical-3332122101312202-0101223232111133-3112023300302123-0223301301303312-3330302231233030-3320333212201213-1201322113120220-3300020210312200"></a>

## tenant property — cluster / 113123210133 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0310233132013002-3320020013103022-0303313332102132-1302021320111213-3111233032022230-3102300132022231-0302103221133320-0311032021222100"></a>

## Next pages — cluster / 113123210133 / 7

- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133020130313331-2302230310321202-3223222002302022-2112200101121201-1100223313212330-1011120221333120-3110321212122122-2123223201121002"></a>

## default_route_pools.endpoint_subsets — endpoint_subsets / 310202110012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.endpoint_subsets

<a id="canonical-0112120313012032-0001303222003221-0211201103201130-2122120211000011-0323110300131202-0213103210100230-1120231223120312-1102301031011133"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-0021311120132303-0232300230301130-1120331213033002-3231022302300323-2321120131122133-1113011100202021-3211212202100020-3131110213313131"></a>

## Direct properties — endpoint_subsets / 310202110012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302132020113100-0132321110001130-2122320003003012-0111012211132010-2013312013201331-1323331320221202-1212131222100103-2233212232321320"></a>

## Next pages — endpoint_subsets / 310202110012 / 4

- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103203221110330-3232222320030333-1032132220311011-1322222231230311-3110003113321132-3330333133133332-1201120132203133-2021103200002021"></a>

## default_route_pools.pool — pool / 112213331223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.pool

<a id="canonical-1203032021222111-2113022103331021-1300131313111112-1232202130003023-0110330312210131-0133132312322110-0333230121100002-2101201013033011"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210301233203132-3220201221301321-2002111111022233-0013313332123000-2013103312202312-2122231312023020-0133332110302232-0001320100003312"></a>

## Direct properties — pool / 112213331223 / 3

<a id="canonical-3122123130000212-3202202332332100-0012310000222010-0231000023303120-1110023001033323-3221023310100320-1222311003133300-3201200231001333"></a>

<a id="canonical-1333200201100110-3223023202112330-1011111322210113-0101130301102003-1203322112211122-2323323223213030-3222301132333102-1021333002331002"></a>

## name property — pool / 112213331223 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1311111200130213-0111212202112123-3003121023302121-1220103230113311-2121002110301013-1320222032032311-1123111310200212-0021023223231212"></a>

<a id="canonical-1310331021302120-2021111021311012-1320220213312100-2133113311200203-2013130003102023-2102221333022211-0200013003200210-0231310221320021"></a>

## namespace property — pool / 112213331223 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1322332320222033-3223230221233020-0112021211033100-2100203221010300-3103122211123110-3312121002123322-1222131123201132-2213003001211233"></a>

<a id="canonical-1110301221123030-1301022200201110-0112232211100320-1322111111333102-3332313100033021-3030022010000030-2210111202302220-0322030002300202"></a>

## tenant property — pool / 112213331223 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1212231110102033-2223103323202121-3033002103000301-1313120100230311-0223330332322001-1122120101331231-1232030113311313-1201301113100213"></a>

## Next pages — pool / 112213331223 / 7

- [default_route_pools](resources--http_loadbalancer--reference--group-018.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012113030000132-3012100211231221-3130003303031233-0223013133101100-2030232213320012-3201201203311111-0130303112131112-1022213320131303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022032031200303-1023113333021332-1231032200123323-2233113031313033-2321233110312230-1213331101021113-3020033132123132-0211330222120203"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / 112030303030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_sensitive_data_policy

<a id="canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature. Defaults to \`map\[\]\`.
Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-018.md#canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2331322013011313-3111202121321023-1113002312200000-0132100333303032-2311313300200313-3003330023023213-0133322032323223-3212012023320222)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

<a id="canonical-1322200100231232-0112220221023210-1033021101210010-2200221130300223-2321221013102133-2113310202330013-2320012000121100-2000312331310103"></a>

## Direct properties — default_sensitive_data_policy / 112030303030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120211321302332-0110011123103001-1232302231021000-0303330001022320-1233233133222232-2132120221303103-1202001023002012-0222230132020113"></a>

## Next pages — default_sensitive_data_policy / 112030303030 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123111133030011-2123003110002011-1031032130100110-1022210200133330-3303330101000301-0203011300213011-1020303111133311-2132323123000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231120213011001-3030110001132112-3003100320220112-2130311220111110-3031213123002200-0111222201322311-2232112121333110-2121231111033021"></a>

## disable_api_definition — disable_api_definition / 020013001302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_definition

<a id="canonical-0121212322222130-3103021101033321-1123000003031210-3103120000300132-1002112302102003-3020032021001211-1330203332020133-1300330021233313"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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
disable_api_definition = {}
```

<a id="canonical-0121021001212300-2103301132123012-0000223301201010-0201113221300113-3003031331130322-2333002203020330-3213223003220120-2130121302313311"></a>

## Direct properties — disable_api_definition / 020013001302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300032033110233-2211230133230323-3013220330121110-1230112013123121-3330110203010101-3012121300200121-1202121130021320-3312332130230132"></a>

## Next pages — disable_api_definition / 020013001302 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0232001202112132-3122220002210323-2123112210120123-2111123031002213-0132120133332022-2111330310202011-2220013001013012-0203213300003333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212222011122202-1231112111223303-3202232212023302-3122113101313032-1302210232030303-0210330223223333-0132223320203031-0321331323132320"></a>

## disable_api_discovery — disable_api_discovery / 100121213030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_discovery

<a id="canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

<a id="canonical-3103231022000303-3000032330123212-1310201300203011-2212321310222111-3300010320023231-0112133322111003-3210232322110330-3013203211313311"></a>

## Direct properties — disable_api_discovery / 100121213030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110321222233022-3133003131112213-3111312031331020-2313101023201332-1202000013133202-0021300323021031-0301223121200322-3132301102111231"></a>

## Next pages — disable_api_discovery / 100121213030 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3313202103132030-0013213023113203-2231222203010002-3132330201211211-1031002112330010-2320302022023023-1021321113132320-2221311103102213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212121323122310-0322033000033011-1021300320122223-1111233131101230-1012233220031002-3203133131121033-1331232030113102-2011312311203121"></a>

## disable_api_testing — disable_api_testing / 102320232111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_testing

<a id="canonical-2102333133332200-1121021111010220-1122302223121000-2100012203322001-3000221313331130-1113021300210303-2123021303330312-2201022000213330"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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
disable_api_testing = {}
```

<a id="canonical-3033011322221112-1222201031301131-1112313113232102-3331320320010221-0213232011332111-0213303200310030-1320303333011002-0021311012132032"></a>

## Direct properties — disable_api_testing / 102320232111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230010120130322-1233131233130132-2303210101303033-1320210123122023-3123023330322303-1113102332320330-3232201123331200-0310130112033203"></a>

## Next pages — disable_api_testing / 102320232111 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3123003300002010-0033323131033110-3101003330221223-1232122203203021-0032123103302012-3321133121110332-3110131110231100-2002011203003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331233000203202-2333201033232103-0112202301131103-2321130113322323-3212110131023120-2130300022313102-3212302212011010-0221113112320303"></a>

## disable_bot_defense — disable_bot_defense / 033021113332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_bot_defense

<a id="canonical-0031322221201123-0210020223332001-1212312133213033-3030023312222333-1332200121023202-1033110032330120-1010322121131033-3331133222011133"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable bot defense. Defaults to \`map\[\]\`. Server applies default
when omitted.

Upstream description:

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
disable_bot_defense = {}
```

<a id="canonical-0230312111102122-1111232012013110-3330112210102121-1300132303301011-3213320231323203-2101130133200313-0300103300322331-0111122200132111"></a>

## Direct properties — disable_bot_defense / 033021113332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301102321332322-1311032323312333-0223321331331323-1102003201031132-2221210112023202-3111013233212210-0312201302212121-3210001131233002"></a>

## Next pages — disable_bot_defense / 033021113332 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2321101032330312-2021001302001033-0303012101301113-3221010231232001-2031231033003121-0102030231321121-2103303220300112-2032023320022021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331331210300110-2331320030113032-2320301113313320-2320120223131031-3322330312102131-3201103213022322-1331302122320003-3322330003233223"></a>

## disable_caching — disable_caching / 312120130312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_caching

<a id="canonical-2310010330303022-3212102030230223-3311321133320300-3300013001030002-1102021010321320-0321310020220321-0011203222103312-1122212300121322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable caching.

Upstream description:

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
disable_caching = {}
```

<a id="canonical-1022322111231220-2203323213023333-0132100200332012-0302200121133303-2022110130321102-3220013310222033-0230001100300010-2221121020330122"></a>

## Direct properties — disable_caching / 312120130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231011311333113-1332002302112221-0113312001332302-3111132113221023-1122013203220102-3230013123310123-2030000230332331-0122030202323331"></a>

## Next pages — disable_caching / 312120130312 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2122321323123120-0131330211013033-0223101332002110-1133012200030110-0300010032111010-3133032130220300-0032001323100111-3300101020203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211101003223302-2211120020020230-2201212323021320-1030221231023020-3320311010201110-2311110003031302-2301312203330131-0020301333030320"></a>

## disable_client_side_defense — disable_client_side_defense / 331232331012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_client_side_defense

<a id="canonical-1313000133203032-2310331201223313-3311201202200022-2303113031201211-1010211132031221-2313112333213303-2001211021211300-1211030023323302"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
disable_client_side_defense = {}
```

<a id="canonical-1233212230201012-3323013120321221-1202210021120001-1010112200320332-3311132113300013-0233001013123032-3321112232010113-3020300233033102"></a>

## Direct properties — disable_client_side_defense / 331232331012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310311303103103-1023122230302130-2130102112020201-3101202323023112-0113321221310331-3002031023030102-0023102120112210-3322230233230201"></a>

## Next pages — disable_client_side_defense / 331232331012 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1202122330212221-3201331122203123-3033330230333010-1333110133101010-1102330300133111-1223133210101012-2101013113121223-0330312222221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300133131023123-0000232202010211-2123030003231013-0013232203132331-3201000220113003-3321011121131221-1303322012001312-3010000232113023"></a>

## disable_ip_reputation — disable_ip_reputation / 112002301233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_ip_reputation

<a id="canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300)
- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-1331031122011033-2302131121012122-3313303233203321-2313300032330123-2331111321311122-1221001122320101-3033133302202020-0230023201131132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

<a id="canonical-1021133213020123-0132011022202021-3113120213100222-0020033310030202-3110131002331011-0113221200221321-0321033200222013-2130323330003322"></a>

## Direct properties — disable_ip_reputation / 112002301233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203300213122011-1030331131031221-3212031000131023-3013013122322322-2210302010201021-1230323323323312-1201000212321310-3332111222132122"></a>

## Next pages — disable_ip_reputation / 112002301233 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3012231223303232-2102332203003323-3102212131101310-2012202210031310-2221201202003333-1132132220302101-0321000202112303-1321313331130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131110121111011-0002003323113202-0110321022212002-3032032110220001-1023010322232333-2332203023302232-3310201301102333-1232322202123122"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / 111333122331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malicious_user_detection

<a id="canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231)
- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-1200132303110020-2331223220030122-3300330122310330-2033201303222103-2311130312023020-2001011122320233-2033303101211233-0313110100132203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

<a id="canonical-1320212033321112-3100111322112113-3100012322102010-3012300322023111-1111020201010303-1010301210303220-0213003322011223-2322312333132131"></a>

## Direct properties — disable_malicious_user_detection / 111333122331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230020030202331-0113231003231301-1120302020102133-0202212112001222-1011323311121212-3300120230033203-0320222113320120-3131101002011322"></a>

## Next pages — disable_malicious_user_detection / 111333122331 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1303333110313121-3231031113100003-0101230011321330-0011203322101330-0102021301032021-1333303102122201-3123112221102102-3010103301030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203222232222302-2000032111332232-3203021030120002-1333120310212202-3003102010002022-3011212222321210-2032113122300131-1211321012312302"></a>

## disable_malware_protection — disable_malware_protection / 233032233322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malware_protection

<a id="canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malware\_protection, malware\_protection\_settings; Default:
disable\_malware\_protection\] Configuration parameter for disable malware protection. Defaults to
\`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_malware_protection](resources--http_loadbalancer--reference--group-018.md#canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2131132110220032-0231001231223002-1030113021221223-1220220323132330-0320302100310030-1022231121013021-1003021302003300-2223310300300221)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malware_protection = {}
```

<a id="canonical-0122230002202030-3123102232110133-0331301121231120-2020222320031002-2233131022213311-0122312332301302-0331001300303002-3101223210112313"></a>

## Direct properties — disable_malware_protection / 233032233322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113032302123213-1101232223110210-0131230111112010-2303022313113221-0002011323322332-3102113110032012-0220131213131011-1200232022223312"></a>

## Next pages — disable_malware_protection / 233032233322 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3010203201310323-0211110203000202-0211202211033123-0200222331221023-2010100321330113-1113001031021000-2100021313301232-0222321320030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221010010020000-1321313021221211-2310312113310012-0011010112020032-0323300120001211-1200231201223123-1033133013120102-0232311033310210"></a>

## disable_rate_limit — disable_rate_limit / 211010030132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_rate_limit

<a id="canonical-1013223133201022-1313332221032232-1331200200123112-3000301121301112-0033001302203101-2312333301123320-2111320003320121-2311310003323022"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable rate limit. Defaults to \`map\[\]\`. Server applies default when
omitted.

Upstream description:

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
disable_rate_limit = {}
```

<a id="canonical-1120033012331032-1301321201022022-1133020113123300-0020011023103012-3301132030021233-1112011002210332-3101231013203212-1031230201103112"></a>

## Direct properties — disable_rate_limit / 211010030132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110121133213011-3012232222230102-2110321330112130-3323323133111210-0210032232223203-0023323331001213-0302323211031321-2001020323230220"></a>

## Next pages — disable_rate_limit / 211010030132 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2201101022220023-1103210213332120-0012201103310011-2012220223033011-2311212133223002-2010030131033033-1301101200102210-0132000311201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203010213003323-0130011213113033-1023000320201221-3111103313303111-3303232310203303-2230112211023123-0000103231132313-0102301130202020"></a>

## disable_threat_mesh — disable_threat_mesh / 120031012301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_threat_mesh

<a id="canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332)
- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-2100022020020011-0200313113330200-3010002022110312-1310112023231102-1110222331121122-3303232030030220-1330000212210230-0030303210201210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

<a id="canonical-1200332121013113-1332011123102003-0030212331300012-0132132233223220-2132303202302302-2111032131001320-1013110312221222-3313212023000021"></a>

## Direct properties — disable_threat_mesh / 120031012301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301333230231221-1303311021030323-1030112130222113-2311100003213233-2123212221031313-2022210120122302-1302020023022202-3023232113302021"></a>

## Next pages — disable_threat_mesh / 120031012301 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2010130233211201-0102210200231211-2331021322013011-3201112123233301-1203313132231022-1130220320000222-0131212013022231-0330001200132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002211111211321-0310033132213221-1103232133202112-1301013210312010-2102233113132102-3303110032323220-0130132303100030-2202130123233301"></a>

## disable_trust_client_ip_headers — disable_trust_client_ip_headers / 011310203120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_trust_client_ip_headers

<a id="canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211)
- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-3331120100110120-1303333032232002-0301220313322023-3200013000131023-3322323222012333-3110100300213120-3312203133332021-1101320122300121)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_trust_client_ip_headers = {}
```

<a id="canonical-3023202331231002-1220330101311023-0102231100331000-3030012123133010-0110323033310202-0033123001133120-0211111330103022-3120230101103121"></a>

## Direct properties — disable_trust_client_ip_headers / 011310203120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112223222301200-1100331021333110-0232200133211101-1201121003323220-0030010002113101-2232321122103030-3010202230002200-2220131013230000"></a>

## Next pages — disable_trust_client_ip_headers / 011310203120 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0231033130322231-2203223313010330-1200320322113021-1333023133103333-3331123032011122-3320211321333012-0333310010101210-0312032310201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201211120210310-1231220203331130-2123312101333123-2321132103032011-1303230330001231-2211030333100310-1123323303323231-1313023222121031"></a>

## disable_waf — disable_waf / 212202223312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_waf

<a id="canonical-0021311132330031-3013103020211321-1223133002212101-2010122201101202-3101313331110313-2013021300302031-2213220033313203-3000113332031001"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable waf. Defaults to \`map\[\]\`. Server applies default when
omitted.

Upstream description:

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
disable_waf = {}
```

<a id="canonical-1221322323110301-2113322022201120-1032000332123020-2213321012322213-2113021100302332-1101112113331102-2231230223210012-3313203201120011"></a>

## Direct properties — disable_waf / 212202223312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122002100121303-0333013000300321-3131122023130201-1201203201202133-3131201133311110-0220201123202031-3211300131323300-0120221012012312"></a>

## Next pages — disable_waf / 212202223312 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2302213303120331-0322020022021133-2103031201103012-0013213333030332-3331330132313320-2112113133220233-3320130323210022-3120123311111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223121010200100-3221003131323001-2320031011121233-0030001102330130-1223033320121302-1300130003131312-2011301321320202-3111310333232133"></a>

## do_not_advertise — do_not_advertise / 332103333020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- do_not_advertise

<a id="canonical-3021311322031332-3001130120301333-0101131021301003-2310200320010222-0133031223311213-1001132011030231-1000103013131203-3110003031031311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

Upstream description:

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
do_not_advertise = {}
```

<a id="canonical-1011023322102020-0010210222211203-2323221002001112-3010112232023123-2123003323201212-3233231030023212-2101102201123021-3211122321320213"></a>

## Direct properties — do_not_advertise / 332103333020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331331312200300-2103030133220131-3302333320003230-3222332322213321-0103221032101000-0022100000310332-1233311010212313-2133200323330311"></a>

## Next pages — do_not_advertise / 332103333020 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030002131303001-0023303102231103-2212131100113301-3101321331002112-3303200131002010-3333200332003322-3013033322133100-2200231130120112"></a>

## enable_api_discovery — enable_api_discovery / 130133310333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_api_discovery

<a id="canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012031020220100-3332330131123133-1010102200022032-1232130330003003-0113110122332011-3302211311131230-0003223300333132-2312331013010010"></a>

## Direct properties — enable_api_discovery / 130133310333 / 3

- [api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030): complete subsection reference.

<a id="canonical-3302022201100031-0103031220110130-1203021112002313-1023233110213132-3333100303231113-1210302221001121-0122110001311230-0203010331132221"></a>

## Next pages — enable_api_discovery / 130133310333 / 4

- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220)
- [enable_api_discovery.default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221)
- [enable_api_discovery.disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313)
- [enable_api_discovery.discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023)
- [enable_api_discovery.enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210212201020102-0220031023312332-0122121322322100-2133122102321131-2031230221030112-3132233110231131-0313131333010023-0223331032322322"></a>

## enable_api_discovery.api_crawler — api_crawler / 032110330122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.api_crawler

<a id="canonical-3313213101132122-1202132221130300-1302030331330003-0311102333000311-0023322113031310-0303110020123220-2331222223222032-1200233000212320"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333030120231212-2002111212002333-3230123300111013-0010223221310310-1310120330012312-0011011331230200-1300010132213133-0322021202322230"></a>

## Direct properties — api_crawler / 032110330122 / 3

- [api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210): complete subsection reference.

<a id="canonical-0310131233233010-2022310112101332-2233311320101231-1223111022122300-3002120230300213-3012013130102011-3313201021321330-3230000010222301"></a>

## Next pages — api_crawler / 032110330122 / 4

- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.disable_api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103321121120320-0223203021001320-0032031100023110-1201102311323111-2201202003322010-2031010302303311-1222003120010312-1300002220023001"></a>

## enable_api_discovery.api_crawler.api_crawler_config — api_crawler_config / 002222013010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-0011013001010323-0102201332220003-3312322223210210-1000003121102120-2110103300130133-0222112132001333-3313031321030031-3123020010111021"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323301302031223-2020322232221030-0100000300003013-1211231112223213-1333013231310103-1120222100201101-0023211330003113-0212321231303223"></a>

## Direct properties — api_crawler_config / 002222013010 / 3

- [domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002): complete subsection reference.

<a id="canonical-2131332001300212-2112132230212132-2013330010013330-0101022233320003-3032211130332213-1002232230133333-3000132132010003-3112002213002102"></a>

## Next pages — api_crawler_config / 002222013010 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223212023130233-1030031111233223-1101320131330101-0222231323330033-3102113321031221-3303133333010110-2230211133001233-0332231233102301"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — domains / 121210233202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1000021102003321-3312202002202311-1232002300103200-0303230203313302-1222132301110311-3203011300032320-2232110133033022-3100220102020121"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103103332210010-0202212022002021-0312323021120003-2032111130100132-0200033312021221-1130313113202022-3133112233132033-1211332310100311"></a>

## Direct properties — domains / 121210233202 / 3

<a id="canonical-3231210223231332-2322303133210301-2113102310231310-1010033112023331-2310102001302003-0110102010333310-1011322123333203-3020211010113302"></a>

<a id="canonical-0122001221133302-1022223202320230-1102323333123020-2031132001212002-2203023331202220-3332313223100120-2131332020123033-2233330013131323"></a>

## domain property — domains / 121210233202 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302): complete subsection reference.

<a id="canonical-0122311221110202-3213030233201030-0031311210002331-3330110232312011-3230223120201231-3223331130203013-3122023133012211-2331213231313231"></a>

## Next pages — domains / 121210233202 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300103322200223-3300132020200032-0302323210320033-0301321231311112-3322001101311302-2003123300112101-3020320001310301-0203113112322312"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — simple_login / 032000013121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-2231030121113112-0333331210332332-2033232111011202-3221331312310032-0221133131211003-0201231321332311-2303130001132121-0331103222231223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303101311223013-1202320110213302-0311112123003011-3102323130000010-3021130311020111-0333220330103201-0113320332301321-0231333002211233"></a>

## Direct properties — simple_login / 032000013121 / 3

- [password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021): complete subsection reference.

<a id="canonical-2211221111221131-3022220213010230-2121111111123201-3310220122311230-1232121000100223-1320032010100013-1233021331203020-3002211202131223"></a>

<a id="canonical-3032020321221233-1012110231331302-1311212312033320-3111300003233300-3102200310010232-3133021323021213-0020112312312002-3032301322001322"></a>

## user property — simple_login / 032000013121 / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3030110301231002-0013111212012031-0213213030001232-1223230233122000-0313100223233332-1110021203231021-1322100300001332-0031223231131323"></a>

## Next pages — simple_login / 032000013121 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031233220310211-2321023021233100-3230123111310031-1000113030232231-0131321202132002-0003100301332133-3112010100331100-0213102102120313"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — password / 111202212222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-0013221312120121-0121102003001202-0310012321221120-3320101300202223-0030113212200113-0330321012013030-0331302223133203-1201101320032231"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033300113223010-3033123122311022-2032321313223003-2121012211230010-2001133133321132-0101312013232233-2323032033110103-0313032301003010"></a>

## Direct properties — password / 111202212222 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-018.md#canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-018.md#canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021): complete subsection reference.

<a id="canonical-2203311320113033-0203303011112313-3111032121001102-3022232103111000-0020111222022321-0023102211030311-2230000001002321-1110001321121132"></a>

## Next pages — password / 111202212222 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--http_loadbalancer--reference--group-018.md#canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--http_loadbalancer--reference--group-018.md#canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001321030313203-1011000331232220-0030323302103230-0122231123230220-1001001201313333-3121122230022311-2002211221112212-2322203203031321"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — blindfold_secret_info / 022032113210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-1131032213202212-3233101101203313-2201300100101322-3323113312122311-3021123003223220-2120222122121201-0012032131012321-0311200132222132"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3012302311322313-2120100011121200-0121330231213310-2021003033002333-1020232130133313-0010212213020120-2102133212020232-3002010023331023"></a>

## Direct properties — blindfold_secret_info / 022032113210 / 3

<a id="canonical-3222202303113312-2003211322313021-1213323210211310-2021211030312112-0031121313200000-3030313202313133-0013312130301313-3202013313001010"></a>

<a id="canonical-1101022212032211-1322000200200013-2332112213102331-2013010110113030-1220013123112131-3302122320312030-3021211203202223-2231133333121011"></a>

## decryption_provider property — blindfold_secret_info / 022032113210 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303130123333311-1122031200201332-2002211102311220-1321301032212100-2130112023032011-0200321112132021-2131010331232300-0001302311212020"></a>

<a id="canonical-1320303320031230-1110322031020122-1122332113220301-1302303113220023-0021020333222321-2331230103112231-2232210322332132-3122302220313212"></a>

## location property — blindfold_secret_info / 022032113210 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3300313000211221-1202020320112203-0022312132111110-0212012031332300-2013212102230320-0332300313311300-3011221231222311-1301210323230113"></a>

<a id="canonical-2210221201202111-2313130203203312-0112202111113300-3112202200310110-3131012130333210-3323131030210002-3032130111031331-3133012123110200"></a>

## store_provider property — blindfold_secret_info / 022032113210 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3112021012213013-3312300302302001-2002111231122003-3203023022211222-0013332302121120-0211320013230132-1100102213100203-2310322320331011"></a>

## Next pages — blindfold_secret_info / 022032113210 / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302332122000322-2310031023103322-1203213021222011-0213212300232212-2102200020210332-0000011312111120-2013302111302230-0133301011320323"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — clear_secret_info / 303020331013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-018.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-018.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-018.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-0312021033113131-3203323223210212-1222201323010321-0110230211033123-0122013132221032-0021301133103122-1003102232332121-2031232321102300"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0001121030133213-0122311233220110-2210110322023223-3031122000103021-3021013113231011-2201302021233200-0020131210321222-0210101022230310"></a>

## Direct properties — clear_secret_info / 303020331013 / 3

<a id="canonical-0313023203130033-1233203102211122-0333010232111310-3213100031102221-0021310022220220-3123111302331110-3221232230210113-1233202323202212"></a>

<a id="canonical-1200100020310010-0002313211330031-3022023031033113-3233233312032330-1331103003332120-3123203030120101-1033222011133110-2000201221102220"></a>

## provider_ref property — clear_secret_info / 303020331013 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0111233133113321-0302020023211100-3301231210121012-1120101331222113-3021122011313230-3212212102302301-0022321222100201-1023332200121103"></a>

<a id="canonical-3120200332100113-0103321112301122-0332221322120303-0220030032031332-1313311322213020-1110203223321010-0111013102231032-3301322300032100"></a>

## URL property — clear_secret_info / 303020331013 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3311320332112123-2202211330321202-1031203231030023-3230330300331022-0111331300003023-3323133120132032-0020002322320221-2231012003203313"></a>

## Next pages — clear_secret_info / 303020331013 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-018.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021231133333102-0300323323221223-0231223020323033-0000203201221312-3311020201331322-1130000200331212-2222031021133312-1203231013230130"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — disable_api_crawler / 122031220030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-2213321311201302-2201122132100103-1002220003122122-0100222311301301-2031111322103202-0111030333301223-3100132223221232-3302010023331113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
disable_api_crawler = {}
```

<a id="canonical-2100102220210210-3101232021312202-0101022103013232-0202332001210231-1020121321302010-1310032220020222-2212322030001222-0132310101110023"></a>

## Direct properties — disable_api_crawler / 122031220030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010323011132323-3233121230112311-2303103012023000-2300112001102123-3333323020012031-1202133223003002-1033111210210133-1013100003201312"></a>

## Next pages — disable_api_crawler / 122031220030 / 4

- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033221303322210-0332221131333123-3021023010001312-1132333220222002-2303101102312232-1031302012210201-1120021131013223-3323131301022003"></a>

## enable_api_discovery.api_discovery_from_code_scan — api_discovery_from_code_scan / 022133133122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-1202202222012133-2133031222200200-2002201220002100-0331102113323201-2230220211230012-2122231101330113-2000323332001112-3130200223213121"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102302130122002-3110032211320311-1333223103000111-3031313012222012-0010102303130311-1323003302222210-1332000300021110-2313000102332130"></a>

## Direct properties — api_discovery_from_code_scan / 022133133122 / 3

- [code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303): complete subsection reference.

<a id="canonical-3212010230021200-0310010123301111-0230333111122022-2312222231003120-2211012210330320-3213220012020032-1021100101332023-1331313001101131"></a>

## Next pages — api_discovery_from_code_scan / 022133133122 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220213323011302-1033203303300211-3023112323033131-0000313121221103-3232300322011121-2100222230100013-3111100131031212-3133012112023202"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — code_base_integrations / 032103133011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-1223202010302331-0233102300033320-2300200202321111-3232012201030330-0013133001323003-2330103023103010-2001311313030110-1031313010013103"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Upstream description:

Configuration parameter for codebase integrations

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331021133111330-2113111122120313-1312223123012220-2031210323210121-3322220332221010-1132020211303230-3110323013111322-3321031012321122"></a>

## Direct properties — code_base_integrations / 032103133011 / 3

- [all_repos](resources--http_loadbalancer--reference--group-018.md#canonical-2012231012033323-2313120330300030-3113220001001310-2310312201210333-0330011223232101-2212201331321002-1003002130002031-1100322103203101): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-018.md#canonical-3011020133023203-1020010031211020-3210220300221021-1030003032211012-3320321132101110-2200123112011310-3123013030010000-1112331223011302): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-018.md#canonical-1330221102102213-0130023021302212-1210323220223000-0200101320031231-2211330313210013-3330033102122132-1011102002321103-3013033213013332): complete subsection reference.

<a id="canonical-0213200210211323-3112132201131322-1010002030103130-3201211232211333-2002132033220320-3001103002211222-2002102101323313-2133121011031321"></a>

## Next pages — code_base_integrations / 032103133011 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--http_loadbalancer--reference--group-018.md#canonical-2012231012033323-2313120330300030-3113220001001310-2310312201210333-0330011223232101-2212201331321002-1003002130002031-1100322103203101)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--http_loadbalancer--reference--group-018.md#canonical-3011020133023203-1020010031211020-3210220300221021-1030003032211012-3320321132101110-2200123112011310-3123013030010000-1112331223011302)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--http_loadbalancer--reference--group-018.md#canonical-1330221102102213-0130023021302212-1210323220223000-0200101320031231-2211330313210013-3330033102122132-1011102002321103-3013033213013332)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012231012033323-2313120330300030-3113220001001310-2310312201210333-0330011223232101-2212201331321002-1003002130002031-1100322103203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333200002213123-3301110022122213-1232030031213111-0312231203232301-3020123100321032-0030132110010213-0002121003333212-2123012303301212"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — all_repos / 133122001110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-1011023010200020-2131223312032031-3312120331311200-0223103020132233-0020111301122031-1031223303202331-3131221302030103-3021222112233230"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-3323211312333201-0110331323100102-0022110201133013-0201331030202330-2213132310322221-3301001223121311-3220200121103233-2123211322222012"></a>

## Direct properties — all_repos / 133122001110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333033323100302-2223330022320110-0110032222321012-3103203313003332-2200120210022321-2031113101232331-1233232123323313-1301133011202303"></a>

## Next pages — all_repos / 133122001110 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011020133023203-1020010031211020-3210220300221021-1030003032211012-3320321132101110-2200123112011310-3123013030010000-1112331223011302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303103323203323-1122131111300220-3313013300302221-0320221313101301-2320211233012221-1013130202002131-2113123222101203-0123011120203033"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — code_base_integration / 033011032333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-2213001021220100-0212002231200333-0332131021022322-1332330200121230-3321131330133223-1011011231001110-3230213201130031-3113130020130321"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332133003010200-2233212322223331-2220323332203203-2032021013130030-1300130232232330-3202132311121311-0302130033302132-0021223011312320"></a>

## Direct properties — code_base_integration / 033011032333 / 3

<a id="canonical-3122103331103203-2300202111030123-2003311232220010-0212312010300300-1210110332131323-2031331313113332-0223200232031102-3232230012121323"></a>

<a id="canonical-3320300103212232-3022323233311031-2222000311322012-0313313132032003-0300201111030133-3321133021102300-3003300323101202-1203001020102320"></a>

## name property — code_base_integration / 033011032333 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0200303020111223-2333212331123223-3132020131232223-3312121133330111-1322030203012330-3223113333003011-0031120123111212-1210021303131111"></a>

## namespace property — code_base_integration / 033011032333 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1033003102100030-0311100331123212-0330113202033231-0201210210212320-0123133223001333-3121330112320300-2313121122201321-0233233313210223"></a>

## tenant property — code_base_integration / 033011032333 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1223133013111220-0120012133032312-2302112320132123-0300013001201021-3103311310131121-2000222311210200-0230011020313213-0033001133221133"></a>

## Next pages — code_base_integration / 033011032333 / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1330221102102213-0130023021302212-1210323220223000-0200101320031231-2211330313210013-3330033102122132-1011102002321103-3013033213013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132122020212320-0221132213322301-0323013031213001-0133232002103330-1123302210312102-2333312103223112-2102023033321220-2332030000123120"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — selected_repos / 313011300011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-0203333112210003-0002301012221031-0133202233030300-1221222311202120-3220211111233323-0211132111331013-3112221100310211-3211022201332333"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1010323311033303-1022202002233022-3022111332320132-3312213313210333-3030102201023002-1020111030013133-1103121032223332-0121302230330033"></a>

## Direct properties — selected_repos / 313011300011 / 3

<a id="canonical-0112113212311002-0103123220203100-3212002132223220-0102201232122102-1302310110232113-2311013110130331-2103023002130010-1312232023310230"></a>

<a id="canonical-0312101323301001-2302103020232010-2302003123001022-3203102222012303-2121022230111322-0020001221221320-1123320012032200-3232201030111002"></a>

## api_code_repo property — selected_repos / 313011300011 / 4

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

<a id="canonical-0210030200223201-2311123220123330-0300310300330213-2030021130010220-3231310131310120-1033111333221211-2130330230021112-3300220033230122"></a>

## Next pages — selected_repos / 313011300011 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-018.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222312331113303-3112102110130122-0111130301103100-0102222212033101-1031303202122231-3210123302221011-3301100131321120-2003100220033212"></a>

## enable_api_discovery.custom_api_auth_discovery — custom_api_auth_discovery / 221030022112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-2032220023231101-3330012233232200-1020310222223220-3323122201013101-1001010023111330-0201300311001023-3100130003020302-1100112010122230"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3031033022101303-0003320221113201-3213032200110323-1332223122301101-1212011031021201-2120103200220333-1032331010202030-0212221202303002"></a>

## Direct properties — custom_api_auth_discovery / 221030022112 / 3

- [api_discovery_ref](resources--http_loadbalancer--reference--group-018.md#canonical-0321132002313203-0132103033123323-2110211300011031-1000230020003333-1332223113131322-2013102213033313-2201313013202233-1203301002312032): complete subsection reference.

<a id="canonical-1321322012332121-0002032030232321-2302222223023122-2302110211010233-3013012321010132-2031303212100000-3021233112012211-3132012220301110"></a>

## Next pages — custom_api_auth_discovery / 221030022112 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--reference--group-018.md#canonical-0321132002313203-0132103033123323-2110211300011031-1000230020003333-1332223113131322-2013102213033313-2201313013202233-1203301002312032)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0321132002313203-0132103033123323-2110211300011031-1000230020003333-1332223113131322-2013102213033313-2201313013202233-1203301002312032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222030002203103-2011310013100132-0011100313302110-1312323100302232-0112110211020221-2033011010102321-3001103023323032-2123203201121112"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — api_discovery_ref / 332030000100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-1211001021322222-0312111230200030-2123123311122302-1012001003001200-1323211032212222-2312100320103210-3103320111030223-1201031202201101"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031103230021213-3100330320103022-1000120331303001-3010002301100321-3323300323003130-3221122333222233-1013200000122230-1121101212003032"></a>

## Direct properties — api_discovery_ref / 332030000100 / 3

<a id="canonical-0003311232123102-3121121302231102-0023321323131202-0211023201111000-1210222213201133-1212330313100300-2111010222012212-2232201332303201"></a>

<a id="canonical-3103321030111012-3333021003313231-1301330001202331-2032233212301101-0303310331213031-0320022022313133-3022201202110222-0212112020220111"></a>

## name property — api_discovery_ref / 332030000100 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3332102011323111-1003300022112133-1011210113032111-0211123333003002-0302101211331320-3013203333030000-3023002120310322-1032302211122132"></a>

## namespace property — api_discovery_ref / 332030000100 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1202301121333132-1310231301301232-1123010021232321-3203103312110130-1232010122201112-1102010132030021-3122311011031121-2131031030020003"></a>

## tenant property — api_discovery_ref / 332030000100 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0203013312230312-1210021213323301-2202311230011001-3030200301301223-1121002220022000-2131032230021033-0123011001322130-0033311203133310"></a>

## Next pages — api_discovery_ref / 332030000100 / 7

- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212022311031003-1113202120203231-0231111113222123-0310113102101230-1311312101313010-3110103232123220-3322101223232020-1023313221301032"></a>

## enable_api_discovery.default_api_auth_discovery — default_api_auth_discovery / 111002100000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-0001322113322313-1122330202103100-3133110123212230-2212233203110032-3303333201112213-2110112211220211-1131210010202132-0202010131000321"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1320230223300112-2321213301211113-3233122220120120-2020022333321021-2002232303231121-1220131031213030-2221211130200232-3213033033312311"></a>

## Direct properties — default_api_auth_discovery / 111002100000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332211001203200-1121130120003320-2123111230033033-1102230323322213-3203233122331312-0313200021130120-2201333223302211-0001012133103231"></a>

## Next pages — default_api_auth_discovery / 111002100000 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130122202012230-2323100032030220-0300001332032330-1013011000002112-2022230000212211-3131313031123113-0121211202313332-1002000110103313"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — disable_learn_from_redirect_traffic / 013222132303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-0123122122113211-3210100333022211-3130321003220032-2121121211313021-2120212321021310-3122003032232223-0100212200111132-1020332231232120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

Upstream description:

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

<a id="canonical-0011010302322032-0010102200222133-0012221121023320-0223002133131203-0233101200302222-0032020030103122-1032212210330033-2202130313233331"></a>

## Direct properties — disable_learn_from_redirect_traffic / 013222132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322323331030013-2330202112313000-3012133020022003-2021321320201321-0000000221000320-3230200323302301-3232002312022332-0033323310032130"></a>

## Next pages — disable_learn_from_redirect_traffic / 013222132303 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323101310221100-2132032222120032-2233000300222011-3122002330022200-0210130131130202-2233120213311302-1301112233103101-3110300020221322"></a>

## enable_api_discovery.discovered_api_settings — discovered_api_settings / 130120302122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.discovered_api_settings

<a id="canonical-3230103320013130-1123210101310112-1031111033221200-1023322130132321-3010002131133020-3321300112101210-1211113323132201-2033020022220113"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2130030132302232-1033122110033002-0101300013210031-1133220022121032-3130313121122331-3102213013102233-1231011011300231-3122312213102310"></a>

## Direct properties — discovered_api_settings / 130120302122 / 3

<a id="canonical-0102021100111210-1323102312123302-0110110223210031-1333113123232032-3231002022002100-3202213311000311-2230310102301121-2230010021031312"></a>

<a id="canonical-2112323003230102-2020310002222321-0311020030223332-1303013321003313-1020203120103301-2103222331001322-3122133103212211-3122102021301013"></a>

## purge_duration_for_inactive_discovered_apis property — discovered_api_settings / 130120302122 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0031030212311331-3012021110201021-3120130021310113-2133332013030303-2011211111022000-3010012223101100-3231002322013302-3111112303132313"></a>

## Next pages — discovered_api_settings / 130120302122 / 5

- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103333123101222-3313003100310022-3222120221332313-0002120023112312-3012122001311001-2123330033213223-1223203310020201-3003210330213131"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_learn_from_redirect_traffic / 331212100220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-3333233011332032-2021222131121220-2323113103112113-3133112202313123-3033300230331233-2301333000122113-0002200121221333-3001130211233320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

Upstream description:

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

<a id="canonical-3001313200332013-1032032310211233-1333030331111303-1130101221213001-2002013323120302-1102111301213200-1130103010323020-3302111232113121"></a>

## Direct properties — enable_learn_from_redirect_traffic / 331212100220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131232113122020-0230323022301000-1030303203300301-0223333222302313-0033333023013330-0022022100331022-0022210212021131-3013332311200100"></a>

## Next pages — enable_learn_from_redirect_traffic / 331212100220 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133311032301030-1200121220020301-3030132203023221-0113032100300022-0313322210233013-0033312133213003-0021210123211223-0300011231233222"></a>

## enable_challenge — enable_challenge / 320001233102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_challenge

<a id="canonical-0331332220112013-2112002310330311-0000103101120120-1022133131132112-2221133100010123-1020333222003102-1303303322000302-2321231220030231"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3121201011131202-2211331210233323-1332023120232200-2001321122120333-1102022310120130-0031113111313220-2213331312011210-3113322232221320"></a>

## Direct properties — enable_challenge / 320001233102 / 3

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1122310300102302-3011210221032332-3112302333322111-1300102120103333-0003232101022013-2103123333023313-3310023100122213-0012112333303132): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1231220200222332-1313233111303023-1313331022022121-2223300012113101-1231101221013011-0011112023203321-1132333112033000-2330210301313022): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-2100200102230112-1122110310232323-3002121100330303-2121230122033003-1112132030310331-0322312302131010-2001023023031203-2323030032123210): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3033232023011212-0012221102201233-3001212312010102-2021222023021103-2301330123301302-3210331130331312-2323330111223110-3212300223230132): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-0333223322232310-3133033231010031-2113311012200323-2232211121233033-3132012023221102-1020333312103020-3313232202013023-0302133132031101): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-018.md#canonical-2200030000202100-0111223200112113-1213031201301203-3032003032021013-0333013223212111-2311103031002002-2032320123012330-1111020321212132): complete subsection reference.

<a id="canonical-0313111203311301-2130220112001312-0222031022111231-2200320132230031-0112211203201111-0020232001213001-2322022003021212-3113221312023212"></a>

## Next pages — enable_challenge / 320001233102 / 4

- [enable_challenge.captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1122310300102302-3011210221032332-3112302333322111-1300102120103333-0003232101022013-2103123333023313-3310023100122213-0012112333303132)
- [enable_challenge.default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-1231220200222332-1313233111303023-1313331022022121-2223300012113101-1231101221013011-0011112023203321-1132333112033000-2330210301313022)
- [enable_challenge.default_js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-2100200102230112-1122110310232323-3002121100330303-2121230122033003-1112132030310331-0322312302131010-2001023023031203-2323030032123210)
- [enable_challenge.default_mitigation_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3033232023011212-0012221102201233-3001212312010102-2021222023021103-2301330123301302-3210331130331312-2323330111223110-3212300223230132)
- [enable_challenge.js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-0333223322232310-3133033231010031-2113311012200323-2232211121233033-3132012023221102-1020333312103020-3313232202013023-0302133132031101)
- [enable_challenge.malicious_user_mitigation](resources--http_loadbalancer--reference--group-018.md#canonical-2200030000202100-0111223200112113-1213031201301203-3032003032021013-0333013223212111-2311103031002002-2032320123012330-1111020321212132)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1122310300102302-3011210221032332-3112302333322111-1300102120103333-0003232101022013-2103123333023313-3310023100122213-0012112333303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203231132022022-0310030013210123-0311021011011212-1213011311020102-3003113310222122-3132333223031002-3303303011222331-1130213221211210"></a>

## enable_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 113301310121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-2203231132010022-1323102113113232-0310112231120103-1101113312111301-1232021103033130-2330012302331102-1222121110233132-3001301032333111"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

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

<a id="canonical-2121122301022123-3301320220301313-2002002000021211-3202321302230300-1221033220321210-2113121101321322-1122210330130131-1303233302120312"></a>

## Direct properties — captcha_challenge_parameters / 113301310121 / 3

<a id="canonical-3332312122302200-2012301333132230-3303211012120133-3310232312100233-0312222221213103-1003000303230300-3132210112332223-2110111320003213"></a>

<a id="canonical-2222131113321121-1323211131112033-1102011021212320-2212012133103103-0022013331221101-3112330230021331-2303131023101020-2011132000110111"></a>

## cookie_expiry property — captcha_challenge_parameters / 113301310121 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3202010333302211-0302230010331332-2331312311332302-1220002310211130-2322010033001302-3332300130300313-3102220311322122-2112033101330210"></a>

## custom_page property — captcha_challenge_parameters / 113301310121 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1030110323102223-0020321022223320-3132030000120213-3122210310110202-2021222323312322-3302010113133323-2130013302221133-3033123311300231"></a>

## Next pages — captcha_challenge_parameters / 113301310121 / 6

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231220200222332-1313233111303023-1313331022022121-2223300012113101-1231101221013011-0011112023203321-1132333112033000-2330210301313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102021210131032-0213000210031023-0311232031022311-3031330230001230-0212022301211312-3000133012330022-3332110032031011-3330122130320300"></a>

## enable_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 203010002300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-2220302033032331-2113332231011223-1312311001123333-2332320321102123-0300220120000211-1300002021221103-2321130100213000-0002112202212210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

Upstream description:

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

<a id="canonical-2330101231132013-3332131202210203-2013212233021313-1020301113101233-0301230212103102-2032121200111023-0100223022021330-2230220000003121"></a>

## Direct properties — default_captcha_challenge_parameters / 203010002300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201310122213220-0111332132011120-3010330121212020-2110311333020113-0210033303000223-1200222222222101-2231303112001303-1122223301301002"></a>

## Next pages — default_captcha_challenge_parameters / 203010002300 / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2100200102230112-1122110310232323-3002121100330303-2121230122033003-1112132030310331-0322312302131010-2001023023031203-2323030032123210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112000301213121-3131300222233013-2102322022100300-0133311323220121-3210133230102233-0010123313010230-0123232120230122-3202003220001132"></a>

## enable_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 103013322103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-3020300013001121-3231221220202200-2123000212311033-1102011313103101-1012012010320231-1323023132331120-3103330302222331-0123110311013002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

Upstream description:

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

<a id="canonical-0233302213331031-0102100312032032-2020300031120000-3230002330231102-2222130211312312-1010033231220310-0221011302001013-0213321232122001"></a>

## Direct properties — default_js_challenge_parameters / 103013322103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222102010001231-1222031332312210-0301020000013120-3203331131211200-2003301103322321-1213111212012130-0103032032330111-1230300313011112"></a>

## Next pages — default_js_challenge_parameters / 103013322103 / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3033232023011212-0012221102201233-3001212312010102-2021222023021103-2301330123301302-3210331130331312-2323330111223110-3212300223230132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001101310331032-1132020302032032-1221320320210223-1100101032022331-1013110313120033-2232223122200230-2333001200121031-3011033231303000"></a>

## enable_challenge.default_mitigation_settings — default_mitigation_settings / 300311202230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.default_mitigation_settings

<a id="canonical-0012213112132321-3300331121100211-2103232123310333-2231121032331102-3230231100302320-2232312233101232-2313312323101112-1331021130111122"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-3300110120203120-1203331030300211-0011103333020221-2030323100330301-2210010023332030-3311122102211232-1323020313030330-0223203002332200"></a>

## Direct properties — default_mitigation_settings / 300311202230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003121002233111-1332011233301300-1022330212123321-3313010103220333-2210211131321032-3212323202310210-2202112321010000-1133232222203032"></a>

## Next pages — default_mitigation_settings / 300311202230 / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0333223322232310-3133033231010031-2113311012200323-2232211121233033-3132012023221102-1020333312103020-3313232202013023-0302133132031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102033321321330-1323300123232023-2120010220112120-2223000131322103-0022213001301033-3221222300223332-1321332322012002-0101101133200132"></a>

## enable_challenge.js_challenge_parameters — js_challenge_parameters / 133100212122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.js_challenge_parameters

<a id="canonical-2213231032301111-2333221232102133-2222032222010233-0300212130123321-1233030123232310-3101100320232113-2023202003230020-0200331012200033"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

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

<a id="canonical-1013302321220331-1003130132111323-2112022330111121-2321203230210330-2221112323202221-1232120133113020-2033330021010011-0312122311303210"></a>

## Direct properties — js_challenge_parameters / 133100212122 / 3

<a id="canonical-2213102220302033-1003120202100032-3033031313311110-0001201020102013-3212112232122311-2120310332322022-2310210323320131-2000311022010331"></a>

<a id="canonical-1311033130201213-3311032212111310-1023131212310320-3330111131211002-2232101111223222-1101122310123031-3013112210332201-1331203222003113"></a>

## cookie_expiry property — js_challenge_parameters / 133100212122 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212230212031201-0213331110102020-0332102112320022-1222311012131221-1211111201213130-1232111211303010-3022331311323203-0112032101120132"></a>

## custom_page property — js_challenge_parameters / 133100212122 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3221211332110303-1020330312221311-2020110123002323-1101303013320210-0322131000022012-2201121003032121-0322311310023200-1301102012122030"></a>

## js_script_delay property — js_challenge_parameters / 133100212122 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232212301322231-0322230222021331-1123022202330133-3032212231020123-0323203300103020-0233303120122202-3003120032331121-0011030023021103"></a>

## Next pages — js_challenge_parameters / 133100212122 / 7

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2200030000202100-0111223200112113-1213031201301203-3032003032021013-0333013223212111-2311103031002002-2032320123012330-1111020321212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303310311322203-2021231032013211-1232032102002103-2300203130222110-2001131302332310-2330113102223331-1120022210331330-2212012021102132"></a>

## enable_challenge.malicious_user_mitigation — malicious_user_mitigation / 101310012303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- enable_challenge.malicious_user_mitigation

<a id="canonical-0112333113330120-2200102210201331-3100303112211000-0032313111103310-0103331213202313-3013101201233301-1323202301330003-1203330003210311"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200032113313022-2032103130332013-1231032123320110-2121011322302321-0003302312311101-2232323012022121-3122103210231313-0122023333323130"></a>

## Direct properties — malicious_user_mitigation / 101310012303 / 3

<a id="canonical-2120212311322022-0303112102122211-2212231130112232-0020000112202003-2002310102232330-1332012200303000-3220232030200021-2133000321020332"></a>

<a id="canonical-2030111103013122-1312122122122122-1331001031221110-2311233202003111-3313201212233131-0300301111222010-1203312003032231-0322113122212312"></a>

## name property — malicious_user_mitigation / 101310012303 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2310122332003213-2023023131012222-0002320232322333-2110102033321121-1223200010112033-3312102000130013-0231322111110212-1303100003302131"></a>

## namespace property — malicious_user_mitigation / 101310012303 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0331231201221233-3203200223010233-0200113033233032-1331122122010102-1313323313111010-1021013312012112-3330100113011112-1311023103323200"></a>

## tenant property — malicious_user_mitigation / 101310012303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2113030312103110-2212301110123230-3022302313232210-0110012200103222-0201030001320311-3203231201332211-3210320220022222-2322021312120120"></a>

## Next pages — malicious_user_mitigation / 101310012303 / 7

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231212223321013-0123213202300103-1002202103000112-0121311132023203-1003223131331322-3021301311230232-3302012302222311-0121321022022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202322233101010-1011202323000323-2231313232232203-1002223100122333-1032123220233012-2213303100101323-2330100031033213-0112100230112112"></a>

## enable_ip_reputation — enable_ip_reputation / 201121323121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_ip_reputation

<a id="canonical-1331031122011033-2302131121012122-3313303233203321-2313300032330123-2331111321311122-1221001122320101-3033133302202020-0230023201131132"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0102200313122301-3122001313202131-0303131020103033-3223123011301111-0103300103231321-0120320032200010-1113332013131310-2101303011022000"></a>

## Direct properties — enable_ip_reputation / 201121323121 / 3

<a id="canonical-3121021201103213-0201320300320231-3333323201003323-2121102323223312-1213210312230303-3033210331212020-3312230032132311-2133313102230220"></a>

<a id="canonical-2123212131120130-2130132300003303-2102300311210231-2312032021321131-3001102002230303-1332202311002010-0313221111011012-1300130301230012"></a>

## ip_threat_categories property — enable_ip_reputation / 201121323121 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2221121302303122-2302033232313202-0110312033333302-0311333133301123-0021023311122132-0330022130102301-0120220210302220-1232321013201002"></a>

## Next pages — enable_ip_reputation / 201121323121 / 5

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2102301301321320-1331123320232110-0223023132023311-1300312321320102-3102112231332231-2233232100233210-3122230110022023-1302310210123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002231203002330-1112210201133100-3210230020012212-3131213022010033-2203221012203032-2221131032030020-0210032203312203-3323300012222033"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / 203230103102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_malicious_user_detection

<a id="canonical-1200132303110020-2331223220030122-3300330122310330-2033201303222103-2311130312023020-2001011122320233-2033303101211233-0313110100132203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

Upstream description:

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

<a id="canonical-1033030200331030-1313203001121220-1320130331120100-2222333302011102-3133210100203013-2022123200130212-2223013202230011-3012001210313331"></a>

## Direct properties — enable_malicious_user_detection / 203230103102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303023332022331-2003313231131013-1020333322022221-0233313310330033-1131233033102232-2131022213211030-3303113321010121-2200001202011201"></a>

## Next pages — enable_malicious_user_detection / 203230103102 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3111322010113111-1002113001030101-3211022303110233-2133221032220321-2011101130012312-1232301001111000-1203020010201230-2331021131223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332330010011213-1111110113302112-3211220202010322-2001111122030302-0322302320110020-2123131232310130-0222013102302000-2221322133122010"></a>

## enable_threat_mesh — enable_threat_mesh / 001113100331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_threat_mesh

<a id="canonical-2100022020020011-0200313113330200-3010002022110312-1310112023231102-1110222331121122-3303232030030220-1330000212210230-0030303210201210"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1113303321023300-1233020010203102-3000131321223312-1003300310132302-0101003230200121-1221301203110221-0020033313013202-3210130203202111"></a>

## Direct properties — enable_threat_mesh / 001113100331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231210232212203-1030312033020121-2312003203020130-2102330322033103-0132200312122211-3131313031322013-1223101113221312-1201322222133203"></a>

## Next pages — enable_threat_mesh / 001113100331 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023200122231012-2002203313103003-2012123021321211-3110001220301030-2122312230032131-3200023120120031-3000230111323102-0321222120322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312310321002320-2013201330333332-2220211332121223-3231002020003200-0131330022001130-0020211203030131-2111210030012122-0102030122230023"></a>

## enable_trust_client_ip_headers — enable_trust_client_ip_headers / 120332002111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_trust_client_ip_headers

<a id="canonical-3331120100110120-1303333032232002-0301220313322023-3200013000131023-3322323222012333-3110100300213120-3312203133332021-1101320122300121"></a>

Type: `"object"`. single nested block, Optional.

Trust Client IP Headers List. List of Client IP Headers.

Upstream description:

List of Client IP Headers.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1100013310013101-2312213030331010-2002123010300212-3020311313231020-3133322330220102-1001230020121223-0003332002203221-1031120120310011"></a>

## Direct properties — enable_trust_client_ip_headers / 120332002111 / 3

<a id="canonical-3020322332123320-2303232013311313-2112011213213012-2312203301121201-2333220321121231-2032032311310112-2210132210320223-0111133212103331"></a>

<a id="canonical-0210013123013102-2231302122331201-2230111323203203-3030213131332321-3101213032020111-1213313132312012-0310100202022333-3200003103213202"></a>

## client_ip_headers property — enable_trust_client_ip_headers / 120332002111 / 4

Type: `["list", "string"]`. Optional.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined..

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0013320221112303-3331223000002301-1100110302300323-2333130030021203-3321300302221301-1120323201112020-2121000230210230-3003013002331003"></a>

## Next pages — enable_trust_client_ip_headers / 120332002111 / 5

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230001211301013-0010020000303130-0201201031230011-3323010031333110-2200123202100131-1333202333031020-3113120332023020-1223301131202302"></a>

## graphql_rules — graphql_rules / 000103211232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- graphql_rules

<a id="canonical-1032222112002010-1031000023012012-2011032032213002-0131002301011011-1011220113220032-2111332222101023-2301110210203132-2120100301232110"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2331113310030331-0221032200201001-3331112010203332-0110322103101131-0332320100132201-3001321232012200-1210033211012103-0131312330200231"></a>

## Direct properties — graphql_rules / 000103211232 / 3

- [any_domain](resources--http_loadbalancer--reference--group-018.md#canonical-3232221130212110-0220113031100232-3030220000211332-2100002113221202-3330301230023021-1120230023230321-2200300122300101-1230331300300101): complete subsection reference.

<a id="canonical-2233102220221121-2322110101022020-0203211210123231-2132031321023121-2202013010200302-1233023032322200-0323331022001323-2312302323021120"></a>

<a id="canonical-2133320102033031-0213012103100301-0032332203331103-3122321332330333-0030123013012131-2200101121333003-3320113131031313-3333322300223220"></a>

## exact_path property — graphql_rules / 000103211232 / 4

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /GraphQL.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1101330232020330-0230121313022032-2110221330302011-1011232313023031-2311102200030312-3311032113133102-3120020211200231-1313003302222331"></a>

## exact_value property — graphql_rules / 000103211232 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1030330331212332-1132110301121331-2222303101101013-2222233113212121-3112102001131011-0012033330111010-0210201023122332-2023231002302000"></a>

## suffix_value property — graphql_rules / 000103211232 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2133221313312121-3023330022022121-1033213302201313-3003122330221310-3313231321020122-1221210030111123-0002130332312322-3100331231100322"></a>

## Next pages — graphql_rules / 000103211232 / 7

- [graphql_rules.any_domain](resources--http_loadbalancer--reference--group-018.md#canonical-3232221130212110-0220113031100232-3030220000211332-2100002113221202-3330301230023021-1120230023230321-2200300122300101-1230331300300101)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- [graphql_rules.metadata](resources--http_loadbalancer--reference--group-018.md#canonical-1311122103230322-0310010320111232-3223011301201233-2200311131220033-2130132102313000-0102213320031221-3101130301333310-1002120311120200)
- [graphql_rules.method_get](resources--http_loadbalancer--reference--group-018.md#canonical-1131111003001301-0312122332120102-2200110102303301-3210020023332231-3022222101013121-3201333300300020-2210321301111112-2000203111220332)
- [graphql_rules.method_post](resources--http_loadbalancer--reference--group-018.md#canonical-2231300223101031-0233020132112220-0133203032231333-2221320230211002-2202013323123213-2220223001302313-2311121332013112-3312321121130313)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3232221130212110-0220113031100232-3030220000211332-2100002113221202-3330301230023021-1120230023230321-2200300122300101-1230331300300101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200222231202311-3203102200013022-0012233111122100-1223211131323022-1233231003312301-2033010123232233-1121101011231120-1211021332120123"></a>

## graphql_rules.any_domain — any_domain / 123012213021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.any_domain

<a id="canonical-1300223311121233-0012220011221313-3112102012221120-3320100333333002-0123033103231022-1003231022210321-3103032030111321-1120132101121201"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1200222113012132-0021200221202300-0332003032130100-0130001233231110-2330113121103210-1011320030020101-0321101030332323-3110313232120302"></a>

## Direct properties — any_domain / 123012213021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212303111202200-3031133320333031-0112232302133022-2210302012320300-1211133103032021-1323230330133202-1311122001200303-2132032130300211"></a>

## Next pages — any_domain / 123012213021 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021233321023310-3233130213023232-3333002222100111-1322322010012323-1303020131310012-3132101133012031-0313120201302121-1002322013013220"></a>

## graphql_rules.graphql_settings — graphql_settings / 002020301011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.graphql_settings

<a id="canonical-2210321022120320-2221310330323023-3200023310123211-3002123210022303-3031203212132311-2130302311331203-0011231233013312-1000013133102321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GraphQL settings.

Upstream description:

GraphQL configuration.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3311112322332320-0010232213303012-3300222313002030-0031332113011102-3321213032033132-1100113030010022-1111200330233023-1301020312311313"></a>

## Direct properties — graphql_settings / 002020301011 / 3

- [disable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323): complete subsection reference.

- [enable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-0123033102031013-1310001202232221-0222212331330210-3133001110031331-0200113100323031-1202123312311021-1301112030320213-3302300003110021): complete subsection reference.

<a id="canonical-2312300221033301-0332311020233322-1021312320001331-3012131230102022-3000313102030113-2321133121232221-1101233033010212-3021300223213111"></a>

<a id="canonical-1111221132023111-3001210200302301-1020222210311210-2001011012313300-2212222000111011-1233211121212300-2003032213210032-0113020223030010"></a>

## max_batched_queries property — graphql_settings / 002020301011 / 4

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2002233110010313-0300032112121312-1232302220300011-0312200323212122-2031320223030102-1310113302113320-2022130213302330-2312023000321200"></a>

## max_depth property — graphql_settings / 002020301011 / 5

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313220100330300-2203002323220213-0221131333010123-1220133123132003-1301231211202300-1023333131330032-0000113333020023-1213203322322231"></a>

## max_total_length property — graphql_settings / 002020301011 / 6

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0121101000222323-0323103303201232-2130020333302223-2232112322111023-0301223233303122-3223311222211233-1321311011123121-0303222301202312"></a>

## Next pages — graphql_settings / 002020301011 / 7

- [graphql_rules.graphql_settings.disable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323)
- [graphql_rules.graphql_settings.enable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-0123033102031013-1310001202232221-0222212331330210-3133001110031331-0200113100323031-1202123312311021-1301112030320213-3302300003110021)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2212111211233221-2103032000023022-0332312001010222-2210310210033121-1103020013301201-2320012321221030-0203303131122022-1321133123332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221321220213311-0011333022010322-3311311332011231-0113021112212232-1101200233020322-2032211033200231-2303002102302222-0113110221330312"></a>

## graphql_rules.graphql_settings.disable_introspection — disable_introspection / 103013023111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-2312120300302212-0121202010322020-2121121302133130-1103011022112312-1230311200332210-3321210101313012-1202113303301101-2310302202310332"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-0023001011331113-0300102202213131-0212331333112023-3210300301122320-1222220110122212-3131103120202210-1001223211310320-1100212313333231"></a>

## Direct properties — disable_introspection / 103013023111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111021320112322-2221203132330121-2303012002023220-2213300012113103-3233111313300131-1100031332003030-2021011020013221-1100121232300303"></a>

## Next pages — disable_introspection / 103013023111 / 4

- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123033102031013-1310001202232221-0222212331330210-3133001110031331-0200113100323031-1202123312311021-1301112030320213-3302300003110021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211012321232303-0123330001202121-1200202121231222-0210020132131200-0122313231320002-0013132001230000-2011221232022221-0221331331030310"></a>

## graphql_rules.graphql_settings.enable_introspection — enable_introspection / 300331002032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-2121301001221311-0323230003331333-1232131203113212-1012313302023121-0233131221203222-2102123012220123-2201202300322002-1231130110322020"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1223323110111120-1321000000212213-3331333330122213-0120322322230101-3330131203233033-3021233211113112-3002310231012032-3021132223210222"></a>

## Direct properties — enable_introspection / 300331002032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132130310132221-1103311100021013-0010021030202302-2023021011133000-3301330300331020-2131110230201303-2302002330220233-3322323032133031"></a>

## Next pages — enable_introspection / 300331002032 / 4

- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1311122103230322-0310010320111232-3223011301201233-2200311131220033-2130132102313000-0102213320031221-3101130301333310-1002120311120200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203023111210211-1203232322312012-3301333011032113-2210002300111212-1112021212032120-3221131032133132-3330032000002132-1001132100113013"></a>

## graphql_rules.metadata — metadata / 212303030130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.metadata

<a id="canonical-1312330201312313-0013011131322010-2203331002122013-1122330112112000-2022130220030201-2111302123213113-2001032113222223-0233103103320001"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302323233300032-3333330010322102-2322110131121203-0002323220322330-3020132311022113-1213200323311310-0200312133120233-2212030002131010"></a>

## Direct properties — metadata / 212303030130 / 3

<a id="canonical-2310022333112110-3213012011013021-2232210322033213-0021120011131131-1032303122222030-1313331103230232-2011100212233101-0012320331111322"></a>

<a id="canonical-0123121111000311-2223132321203011-0221113131031123-2111102120321110-2201012322312200-1210313201001220-3322312103132122-1302230322023113"></a>

## description_spec property — metadata / 212303030130 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0310210330220132-1032121210001031-2320001321000003-0002231220011012-3323021221332010-1323211212302330-2103232121221301-1123331303331123"></a>

<a id="canonical-2323202020110301-1001322211120320-1001200322111100-3121320320120310-0311300303312223-3223213312023031-1131132332133120-3301001001131211"></a>

## name property — metadata / 212303030130 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0210310012133303-0333000230123322-2003131100100313-2123322202210121-3130111123332322-3012032132232322-0132023032032112-0210313120202030"></a>

## Next pages — metadata / 212303030130 / 6

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131111003001301-0312122332120102-2200110102303301-3210020023332231-3022222101013121-3201333300300020-2210321301111112-2000203111220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212221133332202-3110113201132001-2101000131223101-0023201222311131-3101221201210132-0310310233223131-0121012110230313-0001313113033200"></a>

## graphql_rules.method_get — method_get / 312113303310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.method_get

<a id="canonical-1010232111012121-1111131302103310-2331221201102200-3220331130011021-3303101021123103-2131001332322101-1333202303120302-3322321303010203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-2131212021112020-1101332221330201-1333221122002021-1220012331221323-3010330312233212-1030131122000102-2232202222111332-0203211303211303"></a>

## Direct properties — method_get / 312113303310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012020100122002-2320201300201233-0232300211332300-0322131330021113-0023310311220321-2013120221021330-3331332321231012-0320031333011330"></a>

## Next pages — method_get / 312113303310 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231300223101031-0233020132112220-0133203032231333-2221320230211002-2202013323123213-2220223001302313-2311121332013112-3312321121130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001221103203003-2310303133030133-0311210331101333-3100323011301230-1210211021011103-3001013001111113-1201000321310110-2002312010202231"></a>

## graphql_rules.method_post — method_post / 301000010011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- graphql_rules.method_post

<a id="canonical-1331230101032323-1311132132033223-1010331133302011-1110020130021133-0203133000231032-0030330322223231-2302313132231212-1323102003230321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

Upstream description:

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

<a id="canonical-2320023223330231-2231031130111310-2102332300213320-3323101202202122-1033103033003233-0103331032111321-1110211330232333-2122111103322132"></a>

## Direct properties — method_post / 301000010011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001111211020021-1222110211030301-2310133120333320-0313000003320102-1223301230301100-2013302000103021-3001111133213130-1230011120231133"></a>

## Next pages — method_post / 301000010011 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3201330110311330-2123203220332301-2010310002232302-2200333212132031-2112023132010213-3001201001222010-0222333320131100-0013020102021323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331221222213010-1111020232113203-3133222203001030-0333323200133110-3321330011033022-1101022221321330-2003331022022201-0210303231013101"></a>

## http — http / 302213323300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- http

<a id="canonical-1202011002202030-1301123223221210-0221133001100200-3113323103312320-1112211202323013-2102200302020220-1011302212010002-2131232013211102"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy. Changing this type selection requires recreation and
may interrupt service. Supported settings within the same selected type remain updatable.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1313001221113221-3102032101322223-3330133000213311-0211003003131123-0000013323003132-1130022002103202-1332220233000122-2102320220122202"></a>

## Direct properties — http / 302213323300 / 3

<a id="canonical-2131122102321022-2330313211011031-3303033230012121-2133113011020220-1320303033100221-0312200102313112-0020201201231333-2201211131330222"></a>

<a id="canonical-1222133212100220-2002203123023130-0022301002100132-1323000130110133-1332112002020033-1333130121010012-3231031331001031-1032231012020322"></a>

## dns_volterra_managed property — http / 302213323300 / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

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

<a id="canonical-3022030301200023-0213201001013300-3331103221312330-3033112000113032-0021331131110003-1320330132020200-0213212212310002-2202001111313023"></a>

## port property — http / 302213323300 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
