---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222022001221300-2202032033313220-2321130331022200-1030001102101013-2021223013132321-3123130212003223-2231032202201331-3111020003111112"></a>

## Property reference — Property reference / 121231303212 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- Property reference

<a id="canonical-1121222021110000-1120211223112211-3130303011001033-3201323232001020-0221113002122330-0111123211021211-3210303330332110-1010331122233330"></a>

## Direct properties — Property reference / 121231303212 / 3

<a id="canonical-2033300133303132-2210211131221132-2132102000231103-3110100213123321-2303100320002111-2202003120133011-0032211112322300-0223030322023023"></a>

<a id="canonical-2112200001201011-1232031002110101-3332112333033220-0232012110232122-2322202102031301-2130311000313200-1033321320000130-2330331203333032"></a>

## annotations property — Property reference / 121231303212 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230): complete subsection reference.

- [circuit_breaker](resources--cluster--reference--group-001.md#canonical-2221020232120031-3001013211111132-1121322102221222-0312203310302010-0011203101220332-1323012230031023-2230330130032331-3233102132202313): complete subsection reference.

<a id="canonical-3022333300332021-1013101111331101-0210003030221031-0030220300123311-2233131202133333-3211021312302213-2013000210023213-3121133200200221"></a>

<a id="canonical-2112211320220203-3332320212133122-0202012332020321-3123321110302231-0101233212323021-0000000231220313-2332031200332131-3311131300221312"></a>

## connection_timeout property — Property reference / 121231303212 / 5

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_subset](resources--cluster--reference--group-001.md#canonical-2203130112112230-2033101230311313-1232010302031100-1220213120323211-1000330212310013-1123310312113230-2123012021101121-1101210100101030): complete subsection reference.

<a id="canonical-2231111221222012-0220120303013321-3102121310111323-3010131123030322-0133212230202303-2220130123112323-3332203030311320-2311223320210103"></a>

<a id="canonical-2030211210100022-2210101122023033-2022120112132332-3223021003312131-2220331220030131-2223330001201313-2231233323131202-0121203221223002"></a>

## description property — Property reference / 121231303212 / 6

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-1320100213132113-3013103303222103-3130202200223023-0032231210202010-0133331330232332-1020220113121232-3032311030333310-3222321113013322"></a>

<a id="canonical-0121333333020002-2333012220121311-2213002131300320-2101133113100131-0313222012232002-0303111122102120-2113200310001110-2133230020010110"></a>

## disable property — Property reference / 121231303212 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011): complete subsection reference.

<a id="canonical-0132312332330033-3321111211332010-1222023131002303-1110230321220222-1212323313000020-1301112103230223-0200200000132233-2013230202311001"></a>

<a id="canonical-1300032310032220-3322222121100112-1023020123030100-2123021232033130-3132102232331010-2312022330021001-0212300231133301-3310112221221132"></a>

## endpoint_selection property — Property reference / 121231303212 / 8

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-0130002222020001-0232133310310030-3303113101232230-1302101000000023-2102221011311020-3031020222211032-2220002000132123-3001122011123303): complete subsection reference.

- [endpoints](resources--cluster--reference--group-001.md#canonical-1122033200033300-2212233112332301-0113231231121220-3011100310230301-2211231323321310-2322020033013220-2001321233013220-1011003202200223): complete subsection reference.

<a id="canonical-0233231222121300-2010132212000322-2312113202303231-0011201220303300-1011311003022203-3220103103100220-1230001010021201-1010200000103232"></a>

<a id="canonical-1221012001210130-2313123312321011-0112302012030323-2320223130330112-0100003203021322-1302010012113132-0201013020213211-2322231123000231"></a>

## fallback_policy property — Property reference / 121231303212 / 9

Type: `"string"`. Optional, Computed.

\[Enum: NO\_FALLBACK|ANY\_ENDPOINT|DEFAULT\_SUBSET\] Enumeration for SubsetFallbackPolicy if subset
match is not found. The request fails as if the cluster had no endpoint matching the subset policy
Any cluster endpoint may be selected if the cluster had no endpoint matching the subset policy Load
balancing is done over endpoints matching.. Possible values are \`NO\_FALLBACK\`, \`ANY\_ENDPOINT\`,
\`DEFAULT\_SUBSET\`. Defaults to \`NO\_FALLBACK\`.

Upstream description:

Enumeration for SubsetFallbackPolicy if subset match is not found.

The request fails as if the cluster had no endpoint matching the subset policy Any cluster endpoint
may be selected if the cluster had no endpoint matching the subset policy Load balancing is done
over endpoints matching default\_subset if the cluster had no endpoint matching the subset policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NO_FALLBACK",
  "enum": [
    "NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [health_checks](resources--cluster--reference--group-001.md#canonical-3333112022010122-1123221201321130-3332220012202132-3212101010213330-3221010130203101-2233333323313123-3210103231011213-1002123323200100): complete subsection reference.

- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133): complete subsection reference.

- [http2_options](resources--cluster--reference--group-001.md#canonical-1122100310003003-3113331120023312-1013121113333323-1300201130311202-1110110033320221-1200012120120123-3131112312310130-0313133200131313): complete subsection reference.

<a id="canonical-3110322231232111-3103300201231102-3133211302210300-1232320233323231-3212220010031020-3211113031200330-1123002233003120-2120132012232121"></a>

<a id="canonical-1330313100111011-1122120313203313-1122203213331211-0002201331013303-0120131303233001-3120233210112010-2303022210010123-0212030210123121"></a>

## http_idle_timeout property — Property reference / 121231303212 / 10

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1213111021210320-3133020031020330-2033103231020011-3101302300012313-3213133233122330-1212233203213003-0213132011233332-3022233231301021"></a>

<a id="canonical-1011230210133301-1323323321011300-0112121021033200-0331333230030221-1322203210110013-2010312021021302-0230303303032012-3013213333013102"></a>

## ID property — Property reference / 121231303212 / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1022203213321013-1013120123133003-2000031232001231-2101021113223220-1111213201132010-1301112101112130-1120233200222320-0231110012202302"></a>

<a id="canonical-2121112313221032-2333331203013310-1112330122322012-3031302002100020-2310000023022101-0203201232100333-2203332323033312-2033120000202330"></a>

## labels property — Property reference / 121231303212 / 12

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2330320220321013-2220131333132302-1121201133223132-0001200020031213-3020030323121102-2211020220202010-0122011012332000-2022021010332323"></a>

<a id="canonical-0312022312233202-3011323300123012-1012023031110013-0030012031220013-1320302232133331-3203302303310212-3110023312303123-0120033232131311"></a>

## loadbalancer_algorithm property — Property reference / 121231303212 / 13

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3201322022101020-0030203333323112-3000312221332132-2222030322103100-2013220210003300-3003001112020300-2310301322221213-0102020200132033"></a>

<a id="canonical-0312022003232311-0233131120030331-1011022223120122-0322023012132210-0203330311222032-0033202023032022-0132132212203013-3210232111323032"></a>

## max_requests_per_connection property — Property reference / 121231303212 / 14

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\]

Sets the maximum number of requests allowed per connection to the origin server. Enter a value
&gt;=1 to define the request limit per connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

OneOf alternatives in this subsection:

- [max_requests_per_connection](resources--cluster--reference--group-001.md#canonical-3201322022101020-0030203333323112-3000312221332132-2222030322103100-2013220210003300-3003001112020300-2310301322221213-0102020200132033)
- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-2313200231201311-0312120113233312-3033023223102030-1002011022020020-0100333103112202-2031302223131011-0302032220233323-1300123011110120)

Select alternatives according to the provider validators above.

<a id="canonical-0213032200232300-3331110230102232-0302000031211031-0233102023122230-3010031310121103-1100211100222233-3002122032323111-3211301123110303"></a>

<a id="canonical-2211330330202133-3011203300110031-1032332022022001-0103123310022111-3123031121312221-0021032320122311-3121321111001003-3221033320211213"></a>

## name property — Property reference / 121231303212 / 15

Type: `"string"`. Required.

Name of the Cluster. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0103333203123012-1332022312303111-1231300303223023-1113303310121311-2113133033210213-2030200231312021-0210203323200302-3302100030013223"></a>

<a id="canonical-2020312232020000-1231012303200121-1231213113010012-3321302032120020-0230012013120231-3233022213020120-2303103312031023-3131132213122331"></a>

## namespace property — Property reference / 121231303212 / 16

Type: `"string"`. Required.

Namespace where the Cluster is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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
  }
}
```

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322): complete subsection reference.

- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-1222033230013012-0003313031123021-2201020000223133-2130312313211123-2000222100123110-2213333002013013-2112332102110013-3113022211303122): complete subsection reference.

- [outlier_detection](resources--cluster--reference--group-001.md#canonical-1332022310131110-0010322311202313-0132333220211213-1222030130221111-1133102220330211-3122120000022131-3202112111002123-3313332331030110): complete subsection reference.

<a id="canonical-1222321203023301-2201231330122022-0201001302231133-3301131103333010-0133110001010103-0011331332131323-0310023301030212-2230213031102200"></a>

<a id="canonical-3021223011012313-3333122202222223-2320030223011010-3131120000330223-0133220222121020-1111010101000001-1333021113020112-3011112113122022"></a>

## panic_threshold property — Property reference / 121231303212 / 17

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for loadbalancing ignoring its health status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0033223222002132-3130113312332000-3110121121032222-0311212300220331-2111030221111102-1103202003220133-3230012120130011-3222201230122200): complete subsection reference.

- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-1030230022303201-0210032031113222-1010300000311002-1033013033031211-1231331022232311-0022211013203320-1012130030000133-0133112211302022): complete subsection reference.

- [timeouts](resources--cluster--reference--group-001.md#canonical-2330030023321302-2111121231202211-3222113200022000-2010222000312212-1332323122330130-1131000321201000-2221222311312130-0300331200322010): complete subsection reference.

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202): complete subsection reference.

<a id="canonical-3313021202331220-2330213230102231-2232122222210000-0323010302213130-3330021221030213-1102012023322210-3131311222201003-2113123002302320"></a>

## All schema paths — Property reference / 121231303212 / 18

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cluster--reference--group-001.md#canonical-2033300133303132-2210211131221132-2132102000231103-3110100213123321-2303100320002111-2202003120133011-0032211112322300-0223030322023023) |
| `auto_http_config` | [auto_http_config](resources--cluster--reference--group-001.md#canonical-3323112120020132-3032333223220020-2313013301213023-1203123220131011-3301132012131101-2211122220012013-2231213113010303-1211200032100120) |
| `circuit_breaker` | [circuit_breaker](resources--cluster--reference--group-001.md#canonical-0301321011222303-1013230032313131-3223102023222311-1011002131010222-2132021212111111-1100233233123302-1030022111310002-2300132220013102) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](resources--cluster--reference--group-001.md#canonical-3023213033131202-0230102002230323-0221010221320300-1011233103101323-2311230112223231-2331210302223103-3230221112331201-1322321333003312) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](resources--cluster--reference--group-001.md#canonical-0112023131013210-3311313231003103-1131010010003221-0010113111301321-0113322112303203-2211211111101022-0102330233301020-0032330101201021) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](resources--cluster--reference--group-001.md#canonical-1220222100223023-0333333322012211-1103202233100301-0123033032331102-3123212001231223-0302011020020011-2023222030102100-1302000002033203) |
| `circuit_breaker.priority` | [circuit_breaker.priority](resources--cluster--reference--group-001.md#canonical-3221310022012020-1201213323233110-2200231012031013-0222323121123030-2132123201011303-1112212131302312-0330022321010322-2310202311001033) |
| `circuit_breaker.retries` | [circuit_breaker.retries](resources--cluster--reference--group-001.md#canonical-3300221122001100-1130303100002232-0332121320312033-3231330232312123-0302300223100300-1213000131030223-0310212013101302-2212313233200001) |
| `connection_timeout` | [connection_timeout](resources--cluster--reference--group-001.md#canonical-3022333300332021-1013101111331101-0210003030221031-0030220300123311-2233131202133333-3211021312302213-2013000210023213-3121133200200221) |
| `default_subset` | [default_subset](resources--cluster--reference--group-001.md#canonical-3330211002030032-0023131323312110-3112022120010011-3330130001232113-3211210013001002-2220101333031222-3212113303020022-2212100113013003) |
| `description` | [description](resources--cluster--reference--group-001.md#canonical-2231111221222012-0220120303013321-3102121310111323-3010131123030322-0133212230202303-2220130123112323-3332203030311320-2311223320210103) |
| `disable` | [disable](resources--cluster--reference--group-001.md#canonical-1320100213132113-3013103303222103-3130202200223023-0032231210202010-0133331330232332-1020220113121232-3032311030333310-3222321113013322) |
| `disable_proxy_protocol` | [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-2112122012302122-2211122203320030-3323120131210302-0212130210310100-3203011310202032-1010122322002003-2113202312210323-2103311333002323) |
| `endpoint_selection` | [endpoint_selection](resources--cluster--reference--group-001.md#canonical-0132312332330033-3321111211332010-1222023131002303-1110230321220222-1212323313000020-1301112103230223-0200200000132233-2013230202311001) |
| `endpoint_subsets` | [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-0222103230201322-3020333102013131-0031310021001101-0103323022200022-2201230310121230-2033111312212102-3213210302322322-0232320231220230) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](resources--cluster--reference--group-001.md#canonical-0232111112131132-0332131003311213-3331230103202123-0012131220202221-3323302212100112-0323023303002202-1103003321131301-3220031031200233) |
| `endpoints` | [endpoints](resources--cluster--reference--group-001.md#canonical-2313301312133322-1101011333010220-3232202010313231-1312010322012222-0030212033021210-2020121132002013-2312311020303232-3023032132301003) |
| `endpoints.kind` | [endpoints.kind](resources--cluster--reference--group-001.md#canonical-1223001301301311-3131000332011222-0300210101202023-2332022201231230-3123013233320011-0320023202312203-2133313233210013-1213023312122313) |
| `endpoints.name` | [endpoints.name](resources--cluster--reference--group-001.md#canonical-3301133200001033-1013133322121103-1133121000131113-1233130102232132-3221232001312211-2103110001321030-1211120031100020-3002001302103211) |
| `endpoints.namespace` | [endpoints.namespace](resources--cluster--reference--group-001.md#canonical-3102322103330312-2120221103310311-3000213100201012-0113213111320132-0333311312301330-1210321320213202-2323330313110132-1203222331012020) |
| `endpoints.tenant` | [endpoints.tenant](resources--cluster--reference--group-001.md#canonical-1231011300033202-0130200023200222-1002322201021022-0112201223201111-3121033212200011-2310100201230121-1011212200200103-3203133111222112) |
| `endpoints.uid` | [endpoints.uid](resources--cluster--reference--group-001.md#canonical-3022001003013102-0031020111130023-2022300003200010-0310033103320312-2310203102010030-1322033300333020-2301312000322202-3301330011303220) |
| `fallback_policy` | [fallback_policy](resources--cluster--reference--group-001.md#canonical-0233231222121300-2010132212000322-2312113202303231-0011201220303300-1011311003022203-3220103103100220-1230001010021201-1010200000103232) |
| `health_checks` | [health_checks](resources--cluster--reference--group-001.md#canonical-3111322320100103-2202130102211333-3301300131300001-2233002230222232-2100330311003020-3302333001310332-1220220213023302-2233001120333202) |
| `health_checks.kind` | [health_checks.kind](resources--cluster--reference--group-001.md#canonical-3100312022113011-2013012321200332-2230030000221010-2002023220301022-0012223101022001-1211201010321321-2033213331223302-3132332031100003) |
| `health_checks.name` | [health_checks.name](resources--cluster--reference--group-001.md#canonical-2311030122320233-1103030233112023-1312001332102102-2311110023301002-1012311231331131-2232231002231031-0121131012032333-2202032211300201) |
| `health_checks.namespace` | [health_checks.namespace](resources--cluster--reference--group-001.md#canonical-0112300133130313-1333032313031103-0230322130112322-3021212131302031-0310323122113322-3111132323220112-1003213013311003-2113213200332331) |
| `health_checks.tenant` | [health_checks.tenant](resources--cluster--reference--group-001.md#canonical-2113221031313011-0032300122031233-2003232132133323-1031121221210133-2022111111301031-2233332110323230-1203213133013033-3201203021303020) |
| `health_checks.uid` | [health_checks.uid](resources--cluster--reference--group-001.md#canonical-2332111310120313-0000121132011322-1230302021231312-3032233103302013-1013310213123302-2300300330032010-2120303132013123-2222112133132332) |
| `http1_config` | [http1_config](resources--cluster--reference--group-001.md#canonical-2312200301003200-0332313230231001-2033321301231233-3121000302131031-0021223023033331-0023000332232011-3210031320011001-3321003031033200) |
| `http1_config.header_transformation` | [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-0011120112132201-0131112310231302-0013030120301332-0302211232312022-0030003331121101-0002320120101021-3110032230313323-2030012011120210) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](resources--cluster--reference--group-001.md#canonical-3331111300113003-1302233120323130-2332301211320221-1022321022300021-1033133132330333-3102133330001130-0310203232102210-0333111303123033) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-0032220103120333-1223013132322113-3133313101111120-2233232111303131-0320312010230133-2202121101023001-3122110111310023-1121121231203131) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-1330111002203033-0302023033320313-2321101122220133-3020202332213233-3022101200201302-0310000013210331-0021223232123111-2112212300013231) |
| `http2_options` | [http2_options](resources--cluster--reference--group-001.md#canonical-2101222322313120-0223221211121320-0120223111021203-1203032232020011-3221211003003321-1233201120203310-2233131030003202-0001113132332310) |
| `http2_options.enabled` | [http2_options.enabled](resources--cluster--reference--group-001.md#canonical-0330023201103012-1101112302030120-3021100120221331-2200103023300310-3313013111003123-0013213022032212-1020323330133002-0301101131202313) |
| `http_idle_timeout` | [http_idle_timeout](resources--cluster--reference--group-001.md#canonical-3110322231232111-3103300201231102-3133211302210300-1232320233323231-3212220010031020-3211113031200330-1123002233003120-2120132012232121) |
| `id` | [ID](resources--cluster--reference--group-001.md#canonical-1213111021210320-3133020031020330-2033103231020011-3101302300012313-3213133233122330-1212233203213003-0213132011233332-3022233231301021) |
| `labels` | [labels](resources--cluster--reference--group-001.md#canonical-1022203213321013-1013120123133003-2000031232001231-2101021113223220-1111213201132010-1301112101112130-1120233200222320-0231110012202302) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](resources--cluster--reference--group-001.md#canonical-2330320220321013-2220131333132302-1121201133223132-0001200020031213-3020030323121102-2211020220202010-0122011012332000-2022021010332323) |
| `max_requests_per_connection` | [max_requests_per_connection](resources--cluster--reference--group-001.md#canonical-3201322022101020-0030203333323112-3000312221332132-2222030322103100-2013220210003300-3003001112020300-2310301322221213-0102020200132033) |
| `name` | [name](resources--cluster--reference--group-001.md#canonical-0213032200232300-3331110230102232-0302000031211031-0233102023122230-3010031310121103-1100211100222233-3002122032323111-3211301123110303) |
| `namespace` | [namespace](resources--cluster--reference--group-001.md#canonical-0103333203123012-1332022312303111-1231300303223023-1113303310121311-2113133033210213-2030200231312021-0210203323200302-3302100030013223) |
| `no_panic_threshold` | [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-0030112232003133-1300001130002032-1013121032030002-3231003120323012-1112201122312311-2132222303011300-1003203231102100-0002311313103221) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-2313200231201311-0312120113233312-3033023223102030-1002011022020020-0100333103112202-2031302223131011-0302032220233323-1300123011110120) |
| `outlier_detection` | [outlier_detection](resources--cluster--reference--group-001.md#canonical-2000301313012312-2231230331211202-3221001132210122-3001333312022101-3132202023010032-0001112303133132-2011012203323303-1302311011310303) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](resources--cluster--reference--group-001.md#canonical-1200132320100002-1232011030320303-3223333122203302-3133301232030301-2213320121333300-1311201031100300-3313201020301310-2321132303132002) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](resources--cluster--reference--group-001.md#canonical-3012301210332013-0130232133101223-2312320223222222-1002213211222233-1023103011032330-0220123233133231-2210101101100301-0103010212121323) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](resources--cluster--reference--group-001.md#canonical-0013033223010221-0313111303222220-3311311232300013-1300331202310110-3001001023310031-0110123213233112-2010013101131210-1031001311013302) |
| `outlier_detection.interval` | [outlier_detection.interval](resources--cluster--reference--group-001.md#canonical-3021310300001031-1331003323020313-3023020321322002-3103303312031332-2113003130222000-2300310311031002-0101202031012132-1220300231110301) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](resources--cluster--reference--group-001.md#canonical-1330310102212301-3130101201203130-1012220020001333-1212032231112010-0313102113023122-3033213032000000-3200223310132030-3330013113023310) |
| `panic_threshold` | [panic_threshold](resources--cluster--reference--group-001.md#canonical-1222321203023301-2201231330122022-0201001302231133-3301131103333010-0133110001010103-0011331332131323-0310023301030212-2230213031102200) |
| `proxy_protocol_v1` | [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0123022331202012-2203321021103222-1132311223332321-2133323312233121-1221303212030001-2113121211232103-2131113100032230-3031302130011201) |
| `proxy_protocol_v2` | [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-3102031231210012-1020030322123212-3121231203301223-0210133013030310-1320002012120113-2113233320332201-2002120122000132-2133211112203333) |
| `timeouts` | [timeouts](resources--cluster--reference--group-001.md#canonical-0011120012003031-3111002132313321-0231232001302323-1230010031213331-0031110233220101-2010030230013223-1202303111320201-2213011213331120) |
| `timeouts.create` | [timeouts.create](resources--cluster--reference--group-001.md#canonical-1130323101333223-1010032213321333-1133300232113203-3122303333111000-1110021103322021-2121111122322111-3122213313311220-1032011003303232) |
| `timeouts.delete` | [timeouts.delete](resources--cluster--reference--group-001.md#canonical-3012112101011021-0320033303312331-1133031002211310-0021201303310222-1320332103330012-0031023330222121-0102311022131332-3012303221101302) |
| `timeouts.read` | [timeouts.read](resources--cluster--reference--group-001.md#canonical-2131222203323332-1131121222031022-3320212122132111-0023002130320221-1202001131333213-1312311230101111-3322022200231002-0033330230033000) |
| `timeouts.update` | [timeouts.update](resources--cluster--reference--group-001.md#canonical-3200033303013012-1121131301202102-0213012021212131-1120220302011132-1333330301000231-2112111101222310-3012003221303032-1210103322233330) |
| `tls_parameters` | [tls_parameters](resources--cluster--reference--group-001.md#canonical-3131202220133102-1013100100131301-2132120313213100-0321321003131203-2030330031211213-1112033032013313-0103110131021322-1120220330212120) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-2321212120333130-0103220022113300-1022221001323302-3313232232210012-0300331021202221-3111120220030331-3102220023000003-1100122032001031) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](resources--cluster--reference--group-001.md#canonical-2231313331231132-0010220301302320-1101002310113123-3323122121013220-2101022331033012-3331330020330233-2202223202121233-3230211222123320) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](resources--cluster--reference--group-001.md#canonical-3332033231201023-1111103301101203-3321332200012230-0303020333233011-2112211020021031-1113330101100202-2332310110112310-1012221120113012) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](resources--cluster--reference--group-001.md#canonical-1333011113223300-2032333230310321-0222313120212002-0010003223333003-0332311113313030-2133301011133002-3130223321021211-0301101230111322) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](resources--cluster--reference--group-001.md#canonical-3213001020111100-0033111131001010-0033313013013323-1202113020202322-2300121102000123-1303311300012323-3031132231133111-1132301112103103) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](resources--cluster--reference--group-001.md#canonical-2200233121210111-0231301103300202-2131220233331333-3100300100302012-0133000120213011-1310030020311310-3223120023121303-0313032333000013) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](resources--cluster--reference--group-001.md#canonical-3221110212020010-3222330302010201-2212232220221320-1113330020233202-0131300121301201-0201322320103000-2210031300130130-0201300232103102) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](resources--cluster--reference--group-001.md#canonical-1322120020330010-2022311322230223-2023021012220213-2311200122121020-3120330313101231-3112020332112003-0320033313112313-2303020200122021) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](resources--cluster--reference--group-001.md#canonical-1011001323111030-1210011203121131-3103332130112003-1113203002100222-1110332312300323-1030323320133131-0232022223323131-0011110032101133) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](resources--cluster--reference--group-001.md#canonical-2233212232010233-2100121100203131-1332130321011010-0122013121113210-0230000220301123-0000322213210111-3102022023201021-2111000032031131) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](resources--cluster--reference--group-001.md#canonical-2310320120131103-3032122203113332-2012320213331310-1323110111033133-0313210321012131-1302122211302102-3322230332101313-2022233001303023) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-1111221222123022-1311312321201030-3330332220030111-1200001323233030-3230102010132111-1113221320202030-3132103201021310-1200203223313222) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](resources--cluster--reference--group-001.md#canonical-0121331230300021-2131120113111110-0213302301032020-1110131321331233-1001333312301313-1312322231111333-0022130303322013-3021102311010223) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-2121302331021012-1111012132210211-1110322023000113-0323302210201011-1120212101210303-3003102332313332-3123200200121322-3223311130203221) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-001.md#canonical-3200032031202331-3303310300201103-0322120323321011-0233001313110121-2121332331211113-2030232102122110-3323313103000220-1102031200010212) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](resources--cluster--reference--group-001.md#canonical-2302202330020102-1112111232032000-3100020330330013-3300231032212032-2122230221300202-2221120102213111-1223011102301131-3200122201101313) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](resources--cluster--reference--group-001.md#canonical-3112112023102223-2302020210033110-3031231230331132-3010200110110000-0222322020311310-2112313331202001-0132332120020021-3131230332331030) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](resources--cluster--reference--group-001.md#canonical-2003222222313001-2020000110322312-0320300121213332-2131213111123132-0111211202331003-0230110303320210-3200103230033322-1223202122310322) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](resources--cluster--reference--group-001.md#canonical-1212333233312332-3100302131201022-3112113312211001-1030232311200102-1202222120202102-2322202102313301-1202013022200232-0011002213001333) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](resources--cluster--reference--group-001.md#canonical-1102123322002322-1320322100011030-0010320132233300-3231000232301232-2113032101033010-2022012022113021-2230130320012023-0200202301111130) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](resources--cluster--reference--group-001.md#canonical-0003123203233123-2103211200331331-3222100313110322-1203303132102312-1132000032100011-0232333212022332-3010100303101200-0333022303211321) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](resources--cluster--reference--group-001.md#canonical-3021203102003011-2132201220102023-3103210120303002-0230212112123222-0012101202132122-0200130020330221-2221002003102323-2113110210133221) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-0313223202132222-0102310330211003-1122222300322232-0102123033120001-0033313312203210-0331233133120013-0033111120013333-2000213201212232) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-0133021313320021-0231101213111210-1223331302131032-2211021112302021-0111210033212211-0312323103232222-0130100230023033-1102103031021031) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--cluster--reference--group-001.md#canonical-1133211111101322-2021033321312003-1102022311332300-1123322233023100-0323202010300030-2103211332321002-3101113233123301-1131313200112222) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--cluster--reference--group-001.md#canonical-0320231230222310-0031001003033230-1201022023020101-2133312001211130-1231033011233132-3022001000300031-3232021111022133-0203032320021213) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--cluster--reference--group-001.md#canonical-2310310133130010-3323131322122011-3023012102321000-3303220210103213-2120312131130023-0200032022020211-0222120022211013-2023022123131122) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0100100311312333-3331230103112232-3233213303223112-1233021101313233-1321110021213113-1232220323210211-2132130020232133-0023113310302132) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--cluster--reference--group-001.md#canonical-3101311010203203-2023303120120233-2201030100031123-0220333013232222-0030311000132223-2203021231332020-2201023332122111-3002303203102012) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-3101031322202101-3212201230103033-2202123311010221-2101223201322020-3301002232223211-3123003233310010-1103000113333311-2031112213132123) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--cluster--reference--group-001.md#canonical-1230311011232210-2013031031002200-0013231023102303-3130233203233231-3301002030231301-3200121112003311-1221311213012031-3310110210330013) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--cluster--reference--group-001.md#canonical-1022023000001122-3130210331230033-1122200232321113-2232212002202210-0031221200302202-1200113330312301-1102100322322223-3021301032320031) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-1002301011003130-2301121301112233-3232023310303022-2202321022201023-1320221001133132-3321012311003312-1132103320123102-0131202113312312) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-2200322321333323-0311000201202210-3032133223101121-3332230311011312-3013132022222033-1332222213112201-0030322122312002-0022103013103002) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-3311210202122101-3302221201022001-3331323203112112-2000003003132301-3310133013000021-0212222331031220-0003033221332130-3112012011231011) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--cluster--reference--group-001.md#canonical-1133323020300323-2112121112330130-3002001010323102-2011303120200313-3133033102133022-3320110200201111-1020330320201221-1102110103210230) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--cluster--reference--group-001.md#canonical-3130211113302111-0100313000122211-2331010311133100-1313323223303123-3112111212213221-3320223223230023-1131213201230110-0300033322313312) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--cluster--reference--group-001.md#canonical-0131002322032232-2120133203213031-3133003302323223-2133023303132200-0210221102010302-1332323333132311-2201210033133001-0233022100221313) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--reference--group-001.md#canonical-0323230330133100-2323213031330313-3021120112230220-1132220122200001-0030320320030033-2222221111120033-1331331323030133-3221303333013223) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--cluster--reference--group-001.md#canonical-3303210333332310-3202103000302010-1112120201321221-0332122003313230-1120222123222322-2010230302000133-3132331122122123-0100022212113303) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--cluster--reference--group-001.md#canonical-1023301333020013-3301230310130222-1031102122221233-2201003100132010-2220311011000231-3030001013333223-2210031311300221-0021203211322301) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--cluster--reference--group-001.md#canonical-0120003301311011-3311300110113220-3301231322330122-2331031113210112-3122223030032101-3010310323031231-3332213031013023-1103122301132320) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-0133133102312101-1332332333220020-0232232312232220-0021220323221021-0021132030110220-3222222000230321-3201120003120312-1321333333221100) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--cluster--reference--group-002.md#canonical-2211003312223203-2030300222232211-3222310330011031-1320332023303302-0102131300203033-0232310003321230-0311211033300003-3030133333300110) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-0113001301213033-0313332111203011-3221223322213132-0022130301130121-2320330012322003-3102000323101202-2002332201000203-2302030023233101) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-002.md#canonical-1313122211012132-2323103001001011-0212010311201302-1221103302011231-3213123332332012-1223113322323302-0103033010333230-2311102132201233) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--cluster--reference--group-002.md#canonical-3012103321133010-1133133303032000-1031310302000021-1311033132220110-1221023310030311-2002232310120001-3320023323313130-2101322023022232) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--cluster--reference--group-002.md#canonical-2101333123213122-0023033112202223-1013022330210321-0102133302312312-0201223303312123-3012032220321220-0232010010213310-2302203210300010) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--cluster--reference--group-002.md#canonical-1000103232230201-0210021202300100-3030102033020013-0331201222231322-2311122312303331-3202022113301330-0011120131201310-2032023101211310) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--cluster--reference--group-002.md#canonical-0122102332323321-1130221101222111-3212312033312222-2002302321330120-3101201003302113-3033232002312110-1320322320030312-0122210312110213) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--cluster--reference--group-002.md#canonical-3303133130220113-1232100232231223-3023230300311322-3122002310312030-1310232312133203-0031002020131301-3222202323233310-0333312110012122) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--cluster--reference--group-002.md#canonical-1133321011103332-0130311203223102-1033303111022323-3021001010013233-3220202112100111-0323130031302131-1233211222330322-3212121011322212) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--cluster--reference--group-002.md#canonical-2110310211123103-2102100003023020-1003312030333020-0213311011332002-0113213222133001-2300230301123331-2332002030300323-1102303130322300) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](resources--cluster--reference--group-002.md#canonical-0303022213110012-1300303123201131-3303103203230323-0022123313320131-0210210011130022-3223230022013201-1311321323330233-1010212312123302) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-2303301023023213-3311333113121301-1002223201031203-1332130331003033-0322113131301333-3013101131220232-3233111300232102-0233122312021011) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](resources--cluster--reference--group-002.md#canonical-3011031101302320-2132321032112032-3201030210032322-0003233113001212-0133230213031010-0031313020011233-0001000333100302-2003230021320213) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](resources--cluster--reference--group-001.md#canonical-1202102322210022-2133111013230000-2213200211021111-1321202313202311-2100312202102212-0231231110033323-0112003222113231-3311333131013331) |
| `tls_parameters.sni` | [tls_parameters.sni](resources--cluster--reference--group-001.md#canonical-1032233230233223-0302312202012102-1123201223122132-0020032100232220-2202323321120010-1120232223221030-2102123312000102-0320112313302221) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-2012212221130301-0303332301113230-3112332122220113-2220031222031121-0110221211330111-3111022033002232-2012200203321233-2010130310233323) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2123332212301000-2200010021323223-2330213113021121-0221021110000130-0001101030103313-0133101111321103-0322301331033113-0121000000212220) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-0313323233003322-0302021233211011-3111230310312230-3103300120210113-0212200233332101-0102223123320220-0331313020002320-3113233303110330) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-2121321022333202-2233112003101001-3233223201112031-1201321001032030-2020012003332033-3313010300213130-2123000233333301-0313003330003133) |

<a id="canonical-0010331200330331-0112313221021322-1003010130213010-0013000231000203-2213133310230233-2110030012211223-1301102002102230-0013322000323021"></a>

## Next pages — Property reference / 121231303212 / 19

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230)
- [circuit_breaker](resources--cluster--reference--group-001.md#canonical-2221020232120031-3001013211111132-1121322102221222-0312203310302010-0011203101220332-1323012230031023-2230330130032331-3233102132202313)
- [default_subset](resources--cluster--reference--group-001.md#canonical-2203130112112230-2033101230311313-1232010302031100-1220213120323211-1000330212310013-1123310312113230-2123012021101121-1101210100101030)
- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011)
- [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-0130002222020001-0232133310310030-3303113101232230-1302101000000023-2102221011311020-3031020222211032-2220002000132123-3001122011123303)
- [endpoints](resources--cluster--reference--group-001.md#canonical-1122033200033300-2212233112332301-0113231231121220-3011100310230301-2211231323321310-2322020033013220-2001321233013220-1011003202200223)
- [health_checks](resources--cluster--reference--group-001.md#canonical-3333112022010122-1123221201321130-3332220012202132-3212101010213330-3221010130203101-2233333323313123-3210103231011213-1002123323200100)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- [http2_options](resources--cluster--reference--group-001.md#canonical-1122100310003003-3113331120023312-1013121113333323-1300201130311202-1110110033320221-1200012120120123-3131112312310130-0313133200131313)
- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322)
- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-1222033230013012-0003313031123021-2201020000223133-2130312313211123-2000222100123110-2213333002013013-2112332102110013-3113022211303122)
- [outlier_detection](resources--cluster--reference--group-001.md#canonical-1332022310131110-0010322311202313-0132333220211213-1222030130221111-1133102220330211-3122120000022131-3202112111002123-3313332331030110)
- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0033223222002132-3130113312332000-3110121121032222-0311212300220331-2111030221111102-1103202003220133-3230012120130011-3222201230122200)
- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-1030230022303201-0210032031113222-1010300000311002-1033013033031211-1231331022232311-0022211013203320-1012130030000133-0133112211302022)
- [timeouts](resources--cluster--reference--group-001.md#canonical-2330030023321302-2111121231202211-3222113200022000-2010222000312212-1332323122330130-1131000321201000-2221222311312130-0300331200322010)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223100313212220-1102311032020023-0220033233023131-1200103001110211-1001320002133113-0330210032023030-2310002321331023-2210232212210212"></a>

## auto_http_config — auto_http_config / 113331310003 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- auto_http_config

<a id="canonical-3323112120020132-3032333223220020-2313013301213023-1203123220131011-3301132012131101-2211122220012013-2231213113010303-1211200032100120"></a>

Type: `["object", {}]`. Optional.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-3323112120020132-3032333223220020-2313013301213023-1203123220131011-3301132012131101-2211122220012013-2231213113010303-1211200032100120)
- [http1_config](resources--cluster--reference--group-001.md#canonical-2312200301003200-0332313230231001-2033321301231233-3121000302131031-0021223023033331-0023000332232011-3210031320011001-3321003031033200)
- [http2_options](resources--cluster--reference--group-001.md#canonical-2101222322313120-0223221211121320-0120223111021203-1203032232020011-3221211003003321-1233201120203310-2233131030003202-0001113132332310)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
auto_http_config = {}
```

<a id="canonical-0123302221210222-3021001321132331-1303003132013120-1102132300212001-3113202212013320-3221301113123113-2203322121131100-0222101013213122"></a>

## Direct properties — auto_http_config / 113331310003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313201310330200-0332232101302103-2303013120000103-3020233233223033-0203201312310301-0222210222001302-0331000133222122-1213110123131023"></a>

## Next pages — auto_http_config / 113331310003 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2221020232120031-3001013211111132-1121322102221222-0312203310302010-0011203101220332-1323012230031023-2230330130032331-3233102132202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112110221213312-0303103323211011-1132001331231332-3023132220132011-0101123201122032-1230203301220320-1313103230022033-1212100020313001"></a>

## circuit_breaker — circuit_breaker / 110300110221 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- circuit_breaker

<a id="canonical-0301321011222303-1013230032313131-3223102023222311-1011002131010222-2132021212111111-1100233233123302-1030022111310002-2300132220013102"></a>

Type: `"object"`. single nested block, Optional.

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Upstream description:

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
circuit_breaker {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222013130322131-1112213012111010-1001223213320121-0313211122312200-0222132020111021-1311330110000301-2223030103313203-3101202200022010"></a>

## Direct properties — circuit_breaker / 110300110221 / 3

<a id="canonical-3023213033131202-0230102002230323-0221010221320300-1011233103101323-2311230112223231-2331210302223103-3230221112331201-1322321333003312"></a>

<a id="canonical-1301020112013023-3033123330230003-2003303303233301-0200103213020232-3330100012013303-3103031111013003-0121023113202012-3030121333133302"></a>

## connection_limit property — circuit_breaker / 110300110221 / 4

Type: `"number"`. Optional.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0112023131013210-3311313231003103-1131010010003221-0010113111301321-0113322112303203-2211211111101022-0102330233301020-0032330101201021"></a>

<a id="canonical-0132112211213002-1212132311210213-1220033001221303-0113102201121010-3000323100032300-0300130332220222-0030111322301031-3310002220100131"></a>

## max_requests property — circuit_breaker / 110300110221 / 5

Type: `"number"`. Optional.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-1220222100223023-0333333322012211-1103202233100301-0123033032331102-3123212001231223-0302011020020011-2023222030102100-1302000002033203"></a>

<a id="canonical-2103131310121211-2322300132331011-2112003122031302-1131311011120201-1202000330113030-2303102030131223-2331120230321313-0233212222322213"></a>

## pending_requests property — circuit_breaker / 110300110221 / 6

Type: `"number"`. Optional.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately..

Upstream description:

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-3221310022012020-1201213323233110-2200231012031013-0222323121123030-2132123201011303-1112212131302312-0330022321010322-2310202311001033"></a>

<a id="canonical-0110022302000313-0330311011123021-3330221032312313-1123031120330023-2033100320130102-3102012333303311-0113320031031120-0220111021201313"></a>

## priority property — circuit_breaker / 110300110221 / 7

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3300221122001100-1130303100002232-0332121320312033-3231330232312123-0302300223100300-1213000131030223-0310212013101302-2212313233200001"></a>

<a id="canonical-0232303303230111-2110202322203013-2030303002231323-1010030102103202-1030200213331203-3313033023021032-2203020100001103-3222201031320333"></a>

## retries property — circuit_breaker / 110300110221 / 8

Type: `"number"`. Optional.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0032101023100121-0213333133301330-3222012013030001-3303113113020310-3330301003032202-1033103021021311-0113020001031220-1200112111333020"></a>

## Next pages — circuit_breaker / 110300110221 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2203130112112230-2033101230311313-1232010302031100-1220213120323211-1000330212310013-1123310312113230-2123012021101121-1101210100101030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331133113020121-3101330103013021-2032133220312112-0303210321112313-1011333202210231-3302223200303132-1300200002211101-0302011312220013"></a>

## default_subset — default_subset / 220321012123 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- default_subset

<a id="canonical-3330211002030032-0023131323312110-3112022120010011-3330130001232113-3211210013001002-2220101333031222-3212113303020022-2212100113013003"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

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
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

Terraform syntax:

```terraform
default_subset {}
```

<a id="canonical-3220301010233023-3310032021330102-0221031123301030-3221100010012031-0222102111232022-0033011230222200-3110010033221313-2120003232330203"></a>

## Direct properties — default_subset / 220321012123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221323033100132-2130200120000322-3313223003111031-1133232312030230-3202123303313230-1303023233023011-3110213101200303-0211123311003101"></a>

## Next pages — default_subset / 220321012123 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132330120202102-0023323020302212-0023010301212323-3132023323231122-3230332201102310-1332332322020111-0003002122221221-1131001230132001"></a>

## disable_proxy_protocol — disable_proxy_protocol / 000332011331 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- disable_proxy_protocol

<a id="canonical-2112122012302122-2211122203320030-3323120131210302-0212130210310100-3203011310202032-1010122322002003-2113202312210323-2103311333002323"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-2112122012302122-2211122203320030-3323120131210302-0212130210310100-3203011310202032-1010122322002003-2113202312210323-2103311333002323)
- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0123022331202012-2203321021103222-1132311223332321-2133323312233121-1221303212030001-2113121211232103-2131113100032230-3031302130011201)
- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-3102031231210012-1020030322123212-3121231203301223-0210133013030310-1320002012120113-2113233320332201-2002120122000132-2133211112203333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

<a id="canonical-2031020310012223-2023231301103033-0132013001002301-3002031111233132-2023232033031311-0111113022303121-3013201131202231-3010003230311213"></a>

## Direct properties — disable_proxy_protocol / 000332011331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100013310121311-1301102222103222-2333320310213220-3232021300011100-2332113001000202-1101013311213023-1211313321023120-1102032311023022"></a>

## Next pages — disable_proxy_protocol / 000332011331 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0130002222020001-0232133310310030-3303113101232230-1302101000000023-2102221011311020-3031020222211032-2220002000132123-3001122011123303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133120123020212-2010232103302320-2220012013110302-2122200323210220-1031310210310030-2310332102321312-1011310103223031-3233001012023130"></a>

## endpoint_subsets — endpoint_subsets / 121331212133 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- endpoint_subsets

<a id="canonical-0222103230201322-3020333102013131-0031310021001101-0103323022200022-2201230310121230-2033111312212102-3213210302322322-0232320231220230"></a>

Type: `"object"`. list nested block, Optional.

Configure endpoint groups based on metadata labels for traffic routing. Supports weighted
distribution and session affinity across labeled endpoints.

Upstream description:

Cluster may be configured to divide its endpoints into subsets based on metadata attached to the
endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected
by the load balancer.

Endpoint\_subsets is list of subsets for this cluster. Each entry in this list has definition for a
subset (which is collection of keys)

During routing, the route’s metadata match configuration is used to find a specific subset. If there
is a subset with the exact keys and values specified by the route, the subset is used for load
balancing. Otherwise, the fallback policy is used. The cluster’s subset configuration must,
therefore, contain a definition that has the same keys as a given route in order for subset load
balancing to occur.

Example:

RouteConfig

routes: &#8203;- match: &#8203;- headers: \[\] path: path: /1.log query\_params: \[\]
routeDestination: destinations: &#8203;- cluster: &#8203;- kind: cluster.object uid:
00000000-0000-4000-8000-0b50b89d07a2 endpointSubsets: site: india

EndpointConfig

metadata: labels: deployment: debug site: india name: end-1 uid: end-1

ClusterConfig

gcSpec: defaultSubset: stage: production fallbackPolicy: DEFAULT\_SUBSET endpointSubsets: &#8203;-
keys: &#8203;- site &#8203;- keys: &#8203;- stage &#8203;- app

Assume the below endpoints are defined and associated with the cluster.

Endpoint Labels -------- ------

ep1 stage: production, site: india ep2 stage: deployment, site: us ep3 stage: production, app: hr
ep4 site: india

The following table describes some routes and the result of their application to the cluster. The
subset definition for cluster is assumed to be same as given above in the ClusterConfig section

RouteMatch Criteria Subset Reason ------------------- ------ ------

site: india ep1, ep4 Subset of endpoints selected site: us ep2 Subset of endpoints selected app: hr
ep1, ep3 Fallback: No subset selector for "app" alone stage: production, app: hr ep3 Subset of
endpoints selected other: x ep1, ep3 Fallback: No subset selector for “other” (none) ep1, ep3
Fallback: No subset requested.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("keys")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132303310121121-1212032013020132-3122202200321113-1110011033221221-3113123102033222-0212222000333023-2202220013233001-1210311333313012"></a>

## Direct properties — endpoint_subsets / 121331212133 / 3

<a id="canonical-0232111112131132-0332131003311213-3331230103202123-0012131220202221-3323302212100112-0323023303002202-1103003321131301-3220031031200233"></a>

<a id="canonical-1021123322300323-1310100000300331-3212223333101030-0011133233113133-2320323302313111-1020003320313001-2202103211031030-3320212303210103"></a>

## keys property — endpoint_subsets / 121331212133 / 4

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0330133033111120-1230113032033223-1203220132302120-2021232103201232-3311030311231201-3233312023102111-0332012132101021-1300310220031111"></a>

## Next pages — endpoint_subsets / 121331212133 / 5

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1122033200033300-2212233112332301-0113231231121220-3011100310230301-2211231323321310-2322020033013220-2001321233013220-1011003202200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103303120320221-2033132312323033-3312102021333220-0312022213232301-1103132323333110-1212302001111320-1001103200012031-1303100310103001"></a>

## endpoints — endpoints / 220110300320 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- endpoints

<a id="canonical-2313301312133322-1101011333010220-3232202010313231-1312010322012222-0030212033021210-2020121132002013-2312311020303232-3023032132301003"></a>

Type: `"object"`. list nested block, Optional.

List of endpoints for this cluster.

Upstream description:

List of references to all endpoint objects that belong to this cluster.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220100120232102-1312202132000312-3013032031000330-1332230001130120-0213312011121012-0311323033010233-1333203032100313-2230223322332331"></a>

## Direct properties — endpoints / 220110300320 / 3

<a id="canonical-1223001301301311-3131000332011222-0300210101202023-2332022201231230-3123013233320011-0320023202312203-2133313233210013-1213023312122313"></a>

<a id="canonical-0333322121031212-2231220112103220-3333003133311313-2031203012220321-3233121330202113-2032130322111033-0320001331223112-1221002033101210"></a>

## kind property — endpoints / 220110300320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3301133200001033-1013133322121103-1133121000131113-1233130102232132-3221232001312211-2103110001321030-1211120031100020-3002001302103211"></a>

<a id="canonical-1230133002332123-3322311121312220-1222001303003133-1302131033002111-0212330131001022-0302122100333112-2331313022013331-2130303101110012"></a>

## name property — endpoints / 220110300320 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3102322103330312-2120221103310311-3000213100201012-0113213111320132-0333311312301330-1210321320213202-2323330313110132-1203222331012020"></a>

<a id="canonical-2013332332131322-0100020303313003-2123201321200123-0020113223121300-2030322232233300-2030220120121220-3232203202111130-1101231131213020"></a>

## namespace property — endpoints / 220110300320 / 6

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
  }
}
```

<a id="canonical-1231011300033202-0130200023200222-1002322201021022-0112201223201111-3121033212200011-2310100201230121-1011212200200103-3203133111222112"></a>

<a id="canonical-3220133133032003-1233213112003123-3203223112231302-2230030022331020-0010110331111011-2130211203202021-1320010322232121-1201111331103331"></a>

## tenant property — endpoints / 220110300320 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3022001003013102-0031020111130023-2022300003200010-0310033103320312-2310203102010030-1322033300333020-2301312000322202-3301330011303220"></a>

<a id="canonical-1103011221302313-1223222313031000-0232321210100332-1101203021110031-3233100032001312-1223012102110211-1131123023301022-1232021221303011"></a>

## uid property — endpoints / 220110300320 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1011200130232132-0032211212331220-2100220003031313-1103300032123322-2202213012122121-3121121103113130-2212202233223303-1221310203100002"></a>

## Next pages — endpoints / 220110300320 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3333112022010122-1123221201321130-3332220012202132-3212101010213330-3221010130203101-2233333323313123-3210103231011213-1002123323200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122322132111233-3031200211113032-0132002123303300-0311130001130000-3302013210321210-3132010131231113-3112001230230103-1333032301100013"></a>

## health_checks — health_checks / 332221100122 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- health_checks

<a id="canonical-3111322320100103-2202130102211333-3301300131300001-2233002230222232-2100330311003020-3302333001310332-1220220213023302-2233001120333202"></a>

Type: `"object"`. list nested block, Optional.

Health check configuration for backend monitoring.

Upstream description:

List of references to healthcheck object for this cluster.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210130212032131-0322111121203011-2331213302022121-0300012021111121-2323300312022020-1133320233103132-0101210020301000-0012230302023003"></a>

## Direct properties — health_checks / 332221100122 / 3

<a id="canonical-3100312022113011-2013012321200332-2230030000221010-2002023220301022-0012223101022001-1211201010321321-2033213331223302-3132332031100003"></a>

<a id="canonical-3213113302222320-0102102131012030-2300000312323133-3203111332321221-1010322001232101-0322321133313321-3010332331300120-2002021020030211"></a>

## kind property — health_checks / 332221100122 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2311030122320233-1103030233112023-1312001332102102-2311110023301002-1012311231331131-2232231002231031-0121131012032333-2202032211300201"></a>

<a id="canonical-2113331110113212-2213300012312231-1132220203222131-2003110001301110-3112031010121121-3232132103201030-0213023221332322-1200031003333230"></a>

## name property — health_checks / 332221100122 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0112300133130313-1333032313031103-0230322130112322-3021212131302031-0310323122113322-3111132323220112-1003213013311003-2113213200332331"></a>

<a id="canonical-1230321301200330-0333202332300001-0031021033100002-1022223000023330-1321231312303310-3031133300033103-0323000023122010-2033021133330203"></a>

## namespace property — health_checks / 332221100122 / 6

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
  }
}
```

<a id="canonical-2113221031313011-0032300122031233-2003232132133323-1031121221210133-2022111111301031-2233332110323230-1203213133013033-3201203021303020"></a>

<a id="canonical-1120101200203130-3021020313213232-3301200310332010-3300321113303201-0221230022320022-2120312031112122-1002202322110211-0101130013112230"></a>

## tenant property — health_checks / 332221100122 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2332111310120313-0000121132011322-1230302021231312-3032233103302013-1013310213123302-2300300330032010-2120303132013123-2222112133132332"></a>

<a id="canonical-3210102231122332-3010110102130132-1001233013130200-2010132331020211-2302112312110002-0011110320003121-0023002023330112-1212013112331023"></a>

## uid property — health_checks / 332221100122 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3021223213013010-1310313310013202-0322330123113231-3210003121210010-3231232103111203-3023123230202010-0033100132202023-1303221233121203"></a>

## Next pages — health_checks / 332221100122 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002311122300021-3113001233233232-2302233002202120-1220102132223121-3231330301003022-1313322002311022-2012000101022211-1030303332030200"></a>

## http1_config — http1_config / 102111100002 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- http1_config

<a id="canonical-2312200301003200-0332313230231001-2033321301231233-3121000302131031-0021223023033331-0023000332232011-3210031320011001-3321003031033200"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http1_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303021000020030-1013032223100021-3321031330232202-2020103022110301-3210110301311331-2022201023223030-1210100211131130-1111003231310132"></a>

## Direct properties — http1_config / 102111100002 / 3

- [header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011): complete subsection reference.

<a id="canonical-3220022221211232-3032010010212213-0302101120233003-3300101231100133-1001321112131122-2113022300110101-1120000011100120-1033010322333332"></a>

## Next pages — http1_config / 102111100002 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310303333312101-1022011103331103-2301323202123323-0012130101012131-1020230231021120-2203200313011012-3031222312002230-0101232202200312"></a>

## http1_config.header_transformation — header_transformation / 331031111211 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- http1_config.header_transformation

<a id="canonical-0011120112132201-0131112310231302-0013030120301332-0302211232312022-0030003331121101-0002320120101021-3110032230313323-2030012011120210"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1033233021321220-1300012001300020-1111130311210303-2103223111212200-3123020200010332-2002112230312000-2002032311101331-3323301021302111"></a>

## Direct properties — header_transformation / 331031111211 / 3

- [default_header_transformation](resources--cluster--reference--group-001.md#canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201): complete subsection reference.

- [preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131): complete subsection reference.

- [proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120): complete subsection reference.

<a id="canonical-0001001130211133-1212122232012030-2333011113323023-2301013122023332-1213220000300121-2202023113030030-3221230210010020-3312203021113022"></a>

## Next pages — header_transformation / 331031111211 / 4

- [http1_config.header_transformation.default_header_transformation](resources--cluster--reference--group-001.md#canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201)
- [http1_config.header_transformation.preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131)
- [http1_config.header_transformation.proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130313103100311-1300111310003100-3220230201112202-1023210310210100-3210101130311333-0220013113022220-2030213132120133-1321113221032030"></a>

## http1_config.header_transformation.default_header_transformation — default_header_transformation / 021333022303 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- http1_config.header_transformation.default_header_transformation

<a id="canonical-3331111300113003-1302233120323130-2332301211320221-1022321022300021-1033133132330333-3102133330001130-0310203232102210-0333111303123033"></a>

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

<a id="canonical-2011300102212130-1030000011212231-1230203033132231-0022112101223121-1111033330200313-1013212333133003-2111230010130132-2011212003233013"></a>

## Direct properties — default_header_transformation / 021333022303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130301012102331-0012020133210003-3123302102320212-2210200100111201-1022110130100110-3300233210123233-0112220310123033-0122111223322210"></a>

## Next pages — default_header_transformation / 021333022303 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330220131201031-3130012331210113-0030033233130120-0002213230221201-0322332100303200-0000030021033103-0122103101233320-0223113133030203"></a>

## http1_config.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 111011233300 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-0032220103120333-1223013132322113-3133313101111120-2233232111303131-0320312010230133-2202121101023001-3122110111310023-1121121231203131"></a>

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

<a id="canonical-0002331223300101-2122320332102221-1123323222131120-1001202120230313-3212303002103303-2102332311212130-1021132221213222-1120203330331301"></a>

## Direct properties — preserve_case_header_transformation / 111011233300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303020331333133-1002300012221131-0311332332113232-2122131300131223-2101200000331123-2321232233202330-3032002332231003-3231223302233133"></a>

## Next pages — preserve_case_header_transformation / 111011233300 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012122113001023-3000003133313330-0122131231322300-0311122123232312-1111213312310210-1033221101102210-3330121222032033-3020311110320213"></a>

## http1_config.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 223002233313 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-1330111002203033-0302023033320313-2321101122220133-3020202332213233-3022101200201302-0310000013210331-0021223232123111-2112212300013231"></a>

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

<a id="canonical-2302111130033320-1133222021031223-3121103112013322-3202203032212121-3230333011220331-2221210311231010-3133113030133023-1211310103302302"></a>

## Direct properties — proper_case_header_transformation / 223002233313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212220213233202-1010000212101101-2000120013000103-0302230210033010-3320210111213003-3211021221313101-0103321200321332-0010031301221120"></a>

## Next pages — proper_case_header_transformation / 223002233313 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1122100310003003-3113331120023312-1013121113333323-1300201130311202-1110110033320221-1200012120120123-3131112312310130-0313133200131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103103120122322-1212002233310130-1212010332130320-1022310011300230-3111333000113013-2210311010112232-1230223023221022-1021022113111032"></a>

## http2_options — http2_options / 021003310011 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- http2_options

<a id="canonical-2101222322313120-0223221211121320-0120223111021203-1203032232020011-3221211003003321-1233201120203310-2233131030003202-0001113132332310"></a>

Type: `"object"`. single nested block, Optional.

Http2 Protocol OPTIONS for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http2_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300121220323333-2131231023212111-0012313003112322-2112332201200301-2111203312022111-2323020333232010-0011212130331322-0030321322323312"></a>

## Direct properties — http2_options / 021003310011 / 3

<a id="canonical-0330023201103012-1101112302030120-3021100120221331-2200103023300310-3313013111003123-0013213022032212-1020323330133002-0301101131202313"></a>

<a id="canonical-0130220233020010-2010033333132233-2333200312002312-0220100012130030-0233021002332200-1021213221022210-3103120200033303-1030030321130310"></a>

## enabled property — http2_options / 021003310011 / 4

Type: `"bool"`. Optional.

Enable/disable HTTP2 Protocol for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223223201333101-1323310023131023-2301112323000001-2230132330132320-3113033323323000-1023313220222321-2212200201321323-0233222102003330"></a>

## Next pages — http2_options / 021003310011 / 5

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232212202130031-2133131300101213-3223033310201212-0013202332232110-0121332113013032-3121130230332121-2133030001021210-1303133232230232"></a>

## no_panic_threshold — no_panic_threshold / 011113333031 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- no_panic_threshold

<a id="canonical-0030112232003133-1300001130002032-1013121032030002-3231003120323012-1112201122312311-2132222303011300-1003203231102100-0002311313103221"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_panic\_threshold, panic\_threshold; Default: no\_panic\_threshold\] Configuration
parameter for no panic threshold.

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

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-0030112232003133-1300001130002032-1013121032030002-3231003120323012-1112201122312311-2132222303011300-1003203231102100-0002311313103221)
- [panic_threshold](resources--cluster--reference--group-001.md#canonical-1222321203023301-2201231330122022-0201001302231133-3301131103333010-0133110001010103-0011331332131323-0310023301030212-2230213031102200)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_panic_threshold = {}
```

<a id="canonical-2030211011000300-3021200311013103-0223310031301013-3211221022231331-0230313132333231-3302100300031320-0231012302012111-1213133130111213"></a>

## Direct properties — no_panic_threshold / 011113333031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321233020122122-2001113230033213-0320303310021220-3202302233303102-1033022023103311-1200310023322300-0032003120131200-2320311202202033"></a>

## Next pages — no_panic_threshold / 011113333031 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1222033230013012-0003313031123021-2201020000223133-2130312313211123-2000222100123110-2213333002013013-2112332102110013-3113022211303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103023002211323-0221212213112221-1211230010103310-1000020330330031-3202230312000121-3200011112130311-0330103010121223-1211130310333303"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / 313230131111 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- no_request_limit_per_connection

<a id="canonical-2313200231201311-0312120113233312-3033023223102030-1002011022020020-0100333103112202-2031302223131011-0302032220233323-1300123011110120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-1321113132031311-1012322322231102-2031322330221011-1111131211113222-2133000220111020-1003130023302033-2100320231303031-1022120020320133"></a>

## Direct properties — no_request_limit_per_connection / 313230131111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103022233211303-3203002211233220-0310200130321022-2023220202323123-2102321331200320-1320032022332331-0130003131300012-1210222220302203"></a>

## Next pages — no_request_limit_per_connection / 313230131111 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1332022310131110-0010322311202313-0132333220211213-1222030130221111-1133102220330211-3122120000022131-3202112111002123-3313332331030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222101100311230-0213202122033322-0303233113333233-0030131323132120-0211111011303203-2122222121132320-1013311232022022-3331231300110020"></a>

## outlier_detection — outlier_detection / 203001233001 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- outlier_detection

<a id="canonical-2000301313012312-2231230331211202-3221001132210122-3001333312022101-3132202023010032-0001112303133132-2011012203323303-1302311011310303"></a>

Type: `"object"`. single nested block, Optional.

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking. Algorithm 1.

Upstream description:

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking.

Algorithm

&#8203;1. A endpoint is determined to be an outlier (based on configured number of consecutive\_5xx
or consecutive\_gateway\_failures) . &#8203;2. If no endpoints have been ejected, loadbalancer will
eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below
the allowed threshold (specified via max\_ejection\_percent setting). If the number of ejected hosts
is above the threshold, the host is not ejected. &#8203;3. The endpoint is ejected for some number
of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during
load balancing. The number of milliseconds is equal to the base\_ejection\_time value multiplied by
the number of times the host has been ejected. &#8203;4. An ejected endpoint will automatically be
brought back into service after the ejection time has been satisfied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
outlier_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221010212010330-0330003130110113-2122231210032122-1211031020301011-0012311020222130-3202011110102113-1303113323012002-1222231012102020"></a>

## Direct properties — outlier_detection / 203001233001 / 3

<a id="canonical-1200132320100002-1232011030320303-3223333122203302-3133301232030301-2213320121333300-1311201031100300-3313201020301310-2321132303132002"></a>

<a id="canonical-0022313103303000-2100333210303301-0032103111222032-3212013333030332-0103330133302120-3201131003031130-3102232302303311-2131003311231322"></a>

## base_ejection_time property — outlier_detection / 203001233001 / 4

Type: `"number"`. Optional.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail.

Upstream description:

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-3012301210332013-0130232133101223-2312320223222222-1002213211222233-1023103011032330-0220123233133231-2210101101100301-0103010212121323"></a>

<a id="canonical-0130012233032221-2130021011321001-0321101301133233-1233111303112023-0131312321202112-0031011123220022-2102111202312120-2230323202110002"></a>

## consecutive_5xx property — outlier_detection / 203001233001 / 5

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates
the..

Upstream description:

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates the
number of consecutive 5xx responses required before a consecutive 5xx ejection occurs. Defaults to
5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-0013033223010221-0313111303222220-3311311232300013-1300331202310110-3001001023310031-0110123213233112-2010013101131210-1031001311013302"></a>

<a id="canonical-2221232102211320-2013320332031330-1211031200130301-2322023022223303-2012133022231300-1303202101330112-3032001221120320-0311212202302121"></a>

## consecutive_gateway_failure property — outlier_detection / 203001233001 / 6

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.)..

Upstream description:

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-3021310300001031-1331003323020313-3023020321322002-3103303312031332-2113003130222000-2300310311031002-0101202031012132-1220300231110301"></a>

<a id="canonical-3131211322203122-3103020123103322-2131210031122222-0112223003230122-0223332033322200-3012113310111233-3210001113103221-3322220123223033"></a>

## interval property — outlier_detection / 203001233001 / 7

Type: `"number"`. Optional.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Upstream description:

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to 10000ms or 10s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600000),
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1330310102212301-3130101201203130-1012220020001333-1212032231112010-0313102113023122-3033213032000000-3200223310132030-3330013113023310"></a>

<a id="canonical-3320212001132113-0101000023010232-0312023333000133-2310212303320302-2102023121323001-2022201322102123-2300022023133103-1323222133302122"></a>

## max_ejection_percent property — outlier_detection / 203001233001 / 8

Type: `"number"`. Optional.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Upstream description:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0300200003023101-3223311133311312-0011110211031101-2232320123001200-2102030212323022-0031110301111010-3111120031232132-0321303123123023"></a>

## Next pages — outlier_detection / 203001233001 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0033223222002132-3130113312332000-3110121121032222-0311212300220331-2111030221111102-1103202003220133-3230012120130011-3222201230122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221103002232002-3013110332003031-3101332100321023-0003113101331311-1222032130022302-2101010120112303-3023300332200231-3102231221122222"></a>

## proxy_protocol_v1 — proxy_protocol_v1 / 202023013310 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- proxy_protocol_v1

<a id="canonical-0123022331202012-2203321021103222-1132311223332321-2133323312233121-1221303212030001-2113121211232103-2131113100032230-3031302130011201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

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
proxy_protocol_v1 = {}
```

<a id="canonical-1111302301210303-0322001322101300-1111211121232110-3320131001301320-2333222030012022-2230311023013311-3013102123132222-1120113122311022"></a>

## Direct properties — proxy_protocol_v1 / 202023013310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012310032203300-2222022213010001-2010100003011233-3301210333013321-1332011020211100-3303332332301111-2122013202110330-3130000031112203"></a>

## Next pages — proxy_protocol_v1 / 202023013310 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1030230022303201-0210032031113222-1010300000311002-1033013033031211-1231331022232311-0022211013203320-1012130030000133-0133112211302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022120030230112-2133032321030213-3302103322230111-1002132002123201-1321100130122331-1001322321112213-3032303200321032-3013222300313111"></a>

## proxy_protocol_v2 — proxy_protocol_v2 / 020330233131 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- proxy_protocol_v2

<a id="canonical-3102031231210012-1020030322123212-3121231203301223-0210133013030310-1320002012120113-2113233320332201-2002120122000132-2133211112203333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v2.

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
proxy_protocol_v2 = {}
```

<a id="canonical-2103111113122120-1113233110210223-3101022210212110-1033202210311113-3313012120222022-3032111003123121-0212313233012210-0023211031123333"></a>

## Direct properties — proxy_protocol_v2 / 020330233131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320303300313112-2311303203121322-3333000213111232-0320111010023102-0233000013200333-1122022101110233-0210103210211333-2000130313300311"></a>

## Next pages — proxy_protocol_v2 / 020330233131 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2330030023321302-2111121231202211-3222113200022000-2010222000312212-1332323122330130-1131000321201000-2221222311312130-0300331200322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102010102330223-3102300121101233-1120202112213300-3113032123123003-0311123101203230-3221031021113031-2111001121101201-1000203003322113"></a>

## timeouts — timeouts / 132231120323 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- timeouts

<a id="canonical-0011120012003031-3111002132313321-0231232001302323-1230010031213331-0031110233220101-2010030230013223-1202303111320201-2213011213331120"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103012313200202-2203100330222332-0121020230330310-2112001131103331-2323011001230311-2112120133112112-2012010322113020-2000010331321000"></a>

## Direct properties — timeouts / 132231120323 / 3

<a id="canonical-1130323101333223-1010032213321333-1133300232113203-3122303333111000-1110021103322021-2121111122322111-3122213313311220-1032011003303232"></a>

<a id="canonical-3212312113200333-2013312322332021-1313300310130322-1222302300011312-1333033101233211-3322331000110120-3201121313122320-1323023333110312"></a>

## create property — timeouts / 132231120323 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3012112101011021-0320033303312331-1133031002211310-0021201303310222-1320332103330012-0031023330222121-0102311022131332-3012303221101302"></a>

<a id="canonical-2220201120102310-3201022303332031-0113113032221331-3010320301032213-0211121323322130-1203332202313223-1003213302230101-1221123210213233"></a>

## delete property — timeouts / 132231120323 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2131222203323332-1131121222031022-3320212122132111-0023002130320221-1202001131333213-1312311230101111-3322022200231002-0033330230033000"></a>

<a id="canonical-0102313211321011-0321102203213010-2102211013110303-3312011323300333-0233101132203323-2023021213030332-2331303132203212-3022011023133100"></a>

## read property — timeouts / 132231120323 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3200033303013012-1121131301202102-0213012021212131-1120220302011132-1333330301000231-2112111101222310-3012003221303032-1210103322233330"></a>

<a id="canonical-1001101122003023-2112213102121320-2303023302020122-3223120320100002-3312230312003323-1332102223022203-1120001330202203-1021002003201223"></a>

## update property — timeouts / 132231120323 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3211110112000013-1303010030011200-1101110103000012-3203200133132330-1232030020013230-3022023013111132-0233121022101010-0123113211201021"></a>

## Next pages — timeouts / 132231120323 / 8

- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103130112210133-0003131112020023-0233320312222103-2230230312033310-2102233120000321-3102003020200221-2022212122020031-1223300102212010"></a>

## tls_parameters — tls_parameters / 130011002032 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- tls_parameters

<a id="canonical-3131202220133102-1013100100131301-2132120313213100-0321321003131203-2030330031211213-1112033032013313-0103110131021322-1120220330212120"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000221322210310-2103122210222222-0102232031021333-1131220222021321-2121122232321300-2102130122211330-1021320133001320-3200331222133100"></a>

## Direct properties — tls_parameters / 130011002032 / 3

- [cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022): complete subsection reference.

- [common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031): complete subsection reference.

- [default_session_key_caching](resources--cluster--reference--group-002.md#canonical-3111210320132103-3023301033212102-3010212200103223-1330312002311102-2203031033111312-1310030010122203-0223131200231111-2130111222200133): complete subsection reference.

- [disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310): complete subsection reference.

- [disable_sni](resources--cluster--reference--group-002.md#canonical-1231201030310032-2232133130332230-3102203310013132-0231112133233202-2112002222030130-0312121223003313-1302023203330231-0310011232322132): complete subsection reference.

<a id="canonical-1202102322210022-2133111013230000-2213200211021111-1321202313202311-2100312202102212-0231231110033323-0112003222113231-3311333131013331"></a>

<a id="canonical-0222332200030023-3222121231101312-1201311233220000-3131203231021313-3331333332001330-2131232223322103-0113330331331123-0303013212031211"></a>

## max_session_keys property — tls_parameters / 130011002032 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="canonical-1032233230233223-0302312202012102-1123201223122132-0020032100232220-2202323321120010-1120232223221030-2102123312000102-0320112313302221"></a>

<a id="canonical-0131112201103003-1310301221323311-2233100132330003-2222312232202022-1133010223102102-2220110212201033-2302123032102111-2013031320122023"></a>

## sni property — tls_parameters / 130011002032 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-3303030113320302-0313223030232122-1230002312320323-0112201033032021-3032303300003222-3030123202200331-3032022101232333-0330331332030030): complete subsection reference.

<a id="canonical-3230002102002233-3122233113131210-3323212213021133-1200021200303133-3111113331311130-2121313223132111-1203013110100001-0302303003021232"></a>

## Next pages — tls_parameters / 130011002032 / 6

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.default_session_key_caching](resources--cluster--reference--group-002.md#canonical-3111210320132103-3023301033212102-3010212200103223-1330312002311102-2203031033111312-1310030010122203-0223131200231111-2130111222200133)
- [tls_parameters.disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310)
- [tls_parameters.disable_sni](resources--cluster--reference--group-002.md#canonical-1231201030310032-2232133130332230-3102203310013132-0231112133233202-2112002222030130-0312121223003313-1302023203330231-0310011232322132)
- [tls_parameters.use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-3303030113320302-0313223030232122-1230002312320323-0112201033032021-3032303300003222-3030123202200331-3032022101232333-0330331332030030)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233002301202220-1221132313213202-3331013032222303-3311213320233201-0202213021322212-2130003003220222-2130023100100023-0103231321003231"></a>

## tls_parameters.cert_params — cert_params / 233033021111 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.cert_params

<a id="canonical-2321212120333130-0103220022113300-1022221001323302-3313232232210012-0300331021202221-3111120220030331-3102220023000003-1100122032001031"></a>

Type: `"object"`. single nested block, Optional.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "tls_validation_params"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("tls_validation_params",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233010110221023-1231211032003223-3113022312331313-2321022322232201-1203212310222100-3123030113032120-1001300303132102-3130222103023021"></a>

## Direct properties — cert_params / 233033021111 / 3

- [certificates](resources--cluster--reference--group-001.md#canonical-3231221112110221-2222223230122221-0123030212331110-3102133330132221-1323200221220213-2331121203031321-2322013132001213-1231033011033303): complete subsection reference.

<a id="canonical-1322120020330010-2022311322230223-2023021012220213-2311200122121020-3120330313101231-3112020332112003-0320033313112313-2303020200122021"></a>

<a id="canonical-0132031101100110-0021303012003232-0223112232303332-3111031221321300-0223123220011021-1333301112231220-0300311012330233-1323201113033200"></a>

## cipher_suites property — cert_params / 233033021111 / 4

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1011001323111030-1210011203121131-3103332130112003-1113203002100222-1110332312300323-1030323320133131-0232022223323131-0011110032101133"></a>

<a id="canonical-1013200212100323-1301330222122220-2302321100200330-3202022230031020-3003313233332302-3030211232233321-1311212100032330-2030322011200020"></a>

## maximum_protocol_version property — cert_params / 233033021111 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2233212232010233-2100121100203131-1332130321011010-0122013121113210-0230000220301123-0000322213210111-3102022023201021-2111000032031131"></a>

<a id="canonical-1301022122232012-3102203033020230-1321120323200033-3311012003020203-0103300232231002-3311211100132022-0133312202230223-1001010313011231"></a>

## minimum_protocol_version property — cert_params / 233033021111 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
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

- [skip_server_verification](resources--cluster--reference--group-001.md#canonical-3230113212220110-1311312112103302-1122313301110122-2300101023102300-2311300223331201-3302320331023023-1100030020113322-2020202131030311): complete subsection reference.

- [tls_validation_params](resources--cluster--reference--group-001.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111): complete subsection reference.

- [volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120): complete subsection reference.

<a id="canonical-1201010231102103-2021311011320233-0002313212120112-2121300301011030-3332331002303311-3003333132022321-0031232233213013-1103132133200003"></a>

## Next pages — cert_params / 233033021111 / 7

- [tls_parameters.cert_params.certificates](resources--cluster--reference--group-001.md#canonical-3231221112110221-2222223230122221-0123030212331110-3102133330132221-1323200221220213-2331121203031321-2322013132001213-1231033011033303)
- [tls_parameters.cert_params.skip_server_verification](resources--cluster--reference--group-001.md#canonical-3230113212220110-1311312112103302-1122313301110122-2300101023102300-2311300223331201-3302320331023023-1100030020113322-2020202131030311)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- [tls_parameters.cert_params.volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3231221112110221-2222223230122221-0123030212331110-3102133330132221-1323200221220213-2331121203031321-2322013132001213-1231033011033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303123320123211-0213302333022222-2112321330202231-1132303103210102-1302333322122021-1312021023111330-3133312331123022-0013333021000023"></a>

## tls_parameters.cert_params.certificates — certificates / 133133211101 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.certificates

<a id="canonical-2231313331231132-0010220301302320-1101002310113123-3323122121013220-2101022331033012-3331330020330233-2202223202121233-3230211222123320"></a>

Type: `"object"`. list nested block, Optional.

Client TLS Certificate required for mTLS authentication.

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
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003210221231323-2101123310112300-0102210301120030-3122312312023310-2031101330303302-0232000022223210-3202013100032112-1213222211030322"></a>

## Direct properties — certificates / 133133211101 / 3

<a id="canonical-3332033231201023-1111103301101203-3321332200012230-0303020333233011-2112211020021031-1113330101100202-2332310110112310-1012221120113012"></a>

<a id="canonical-2232002002212033-1310001122111030-0012001231322132-3200130123000210-1332003302033110-2331300133212313-2310212203233203-3333012231202333"></a>

## kind property — certificates / 133133211101 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1333011113223300-2032333230310321-0222313120212002-0010003223333003-0332311113313030-2133301011133002-3130223321021211-0301101230111322"></a>

<a id="canonical-1012010313201311-3231112131122330-1332012131002120-1230002231120210-2021221111212112-1312211031310022-2223123113210212-1301210330231131"></a>

## name property — certificates / 133133211101 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3213001020111100-0033111131001010-0033313013013323-1202113020202322-2300121102000123-1303311300012323-3031132231133111-1132301112103103"></a>

<a id="canonical-0311002332002131-0312020111022032-2002031033222311-0212000211201133-0232332212210333-0202120303330203-2222230103133313-0303110131323122"></a>

## namespace property — certificates / 133133211101 / 6

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
  }
}
```

<a id="canonical-2200233121210111-0231301103300202-2131220233331333-3100300100302012-0133000120213011-1310030020311310-3223120023121303-0313032333000013"></a>

<a id="canonical-2133123300221220-0313321200321221-3030001203132132-1000130003323230-0322210220202222-3030131132203113-2310231331312310-3103332210212302"></a>

## tenant property — certificates / 133133211101 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3221110212020010-3222330302010201-2212232220221320-1113330020233202-0131300121301201-0201322320103000-2210031300130130-0201300232103102"></a>

<a id="canonical-0010112132130313-2313313230310102-0130210300100212-1122033231200233-2110101100230220-1021110332232020-3233110133020322-2301311332331021"></a>

## uid property — certificates / 133133211101 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3203313211222232-1002231002102122-1103200311003132-1300132201130010-1301202021221001-0303320113311120-1202303021103303-1222231333133203"></a>

## Next pages — certificates / 133133211101 / 9

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3230113212220110-1311312112103302-1122313301110122-2300101023102300-2311300223331201-3302320331023023-1100030020113322-2020202131030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132120102320221-2120003110131300-0013223000320133-2020111010323003-1133102322113231-2201200330011310-1110023210331200-1301011300020300"></a>

## tls_parameters.cert_params.skip_server_verification — skip_server_verification / 320130012101 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.skip_server_verification

<a id="canonical-2310320120131103-3032122203113332-2012320213331310-1323110111033133-0313210321012131-1302122211302102-3322230332101313-2022233001303023"></a>

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
skip_server_verification = {}
```

<a id="canonical-2303321112222001-3122313103102022-2201021300010322-2333333302112023-1323030031010221-3300111331312320-3333211011211113-0133130022133103"></a>

## Direct properties — skip_server_verification / 320130012101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203122133132233-3230013302133131-1101333103222203-2013320010012212-3202212332322101-1032221302230330-3311011313300113-1103230130110022"></a>

## Next pages — skip_server_verification / 320130012101 / 4

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132323012022112-1312002223222301-2213020020203330-0021011331133132-1320131001202203-3312133122211022-3320033323033130-1111111332330011"></a>

## tls_parameters.cert_params.tls_validation_params — tls_validation_params / 033103303113 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.tls_validation_params

<a id="canonical-1111221222123022-1311312321201030-3330332220030111-1200001323233030-3230102010132111-1113221320202030-3132103201021310-1200203223313222"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
tls_validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222101201320300-2000232103223202-1011310101301021-0021131322222101-1220101133103000-0100133232113223-3022110000233313-3311321333020120"></a>

## Direct properties — tls_validation_params / 033103303113 / 3

<a id="canonical-0121331230300021-2131120113111110-0213302301032020-1110131321331233-1001333312301313-1312322231111333-0022130303322013-3021102311010223"></a>

<a id="canonical-0233212032131100-0232012113022313-0300123132122130-2333010203122221-2112333020102030-0320113132222302-2121021020331210-3102121203122222"></a>

## skip_hostname_verification property — tls_validation_params / 033103303113 / 4

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](resources--cluster--reference--group-001.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322): complete subsection reference.

<a id="canonical-0003123203233123-2103211200331331-3222100313110322-1203303132102312-1132000032100011-0232333212022332-3010100303101200-0333022303211321"></a>

<a id="canonical-0230300323201222-1300303210131203-1231220321310313-2210121000312131-3222320302002333-1322313003012222-3130121101122310-1000312110320320"></a>

## trusted_ca_url property — tls_validation_params / 033103303113 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3021203102003011-2132201220102023-3103210120303002-0230212112123222-0012101202132122-0200130020330221-2221002003102323-2113110210133221"></a>

<a id="canonical-2003033303330332-0100102021002213-1212102203020111-1121002010210020-2133010200321310-1203332030111211-0002031003023021-2231211300012111"></a>

## verify_subject_alt_names property — tls_validation_params / 033103303113 / 6

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002033033312202-0310201201322320-1131323121111320-2201121011121121-2010212310010022-0022003013210332-1102000221321133-2002302031203023"></a>

## Next pages — tls_validation_params / 033103303113 / 7

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120201212232011-1122133123331203-2212130223013020-3330131322013320-1123320202221002-3001001203111001-3321111203312233-2212223333023032"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca — trusted_ca / 222220213322 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="canonical-2121302331021012-1111012132210211-1110322023000113-0323302210201011-1120212101210303-3003102332313332-3123200200121322-3223311130203221"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-2032000011311212-3221201311322020-0113310130311102-3103223022301031-0122202332203302-2203032123303032-2012333033223310-2221320121031231"></a>

## Direct properties — trusted_ca / 222220213322 / 3

- [trusted_ca_list](resources--cluster--reference--group-001.md#canonical-3210202012233222-2222111012311100-0302200210313202-1110231121211201-2133310310320312-1123123022220111-1311011101003222-2302030122313211): complete subsection reference.

<a id="canonical-1230020232123223-2133213201302310-2030231300231310-1031123030132013-2011301320001113-2130313323110201-2202031120132221-1111020112220010"></a>

## Next pages — trusted_ca / 222220213322 / 4

- [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-001.md#canonical-3210202012233222-2222111012311100-0302200210313202-1110231121211201-2133310310320312-1123123022220111-1311011101003222-2302030122313211)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3210202012233222-2222111012311100-0302200210313202-1110231121211201-2133310310320312-1123123022220111-1311011101003222-2302030122313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133110110211033-0031313311211130-1201230132211133-2332103120313132-0010003201200300-2333012330331123-0312103322023200-3313313001323231"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 322212210001 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322)
- tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3200032031202331-3303310300201103-0322120323321011-0233001313110121-2121332331211113-2030232102122110-3323313103000220-1102031200010212"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131030223321101-2230030332331333-2013230103122112-2000031302300013-3023010003030323-0223213130010310-0332203213300231-3000312111332112"></a>

## Direct properties — trusted_ca_list / 322212210001 / 3

<a id="canonical-2302202330020102-1112111232032000-3100020330330013-3300231032212032-2122230221300202-2221120102213111-1223011102301131-3200122201101313"></a>

<a id="canonical-3210010123222102-1102310031220201-2003202223233122-3030211233221013-2003011100302121-3333132021220003-0102310211032302-1002021222303323"></a>

## kind property — trusted_ca_list / 322212210001 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3112112023102223-2302020210033110-3031231230331132-3010200110110000-0222322020311310-2112313331202001-0132332120020021-3131230332331030"></a>

<a id="canonical-1223330312002110-2113220100032031-2102101021313213-2312312220313121-0003033020020100-3000232330300321-2303022303310300-1001133223022320"></a>

## name property — trusted_ca_list / 322212210001 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2003222222313001-2020000110322312-0320300121213332-2131213111123132-0111211202331003-0230110303320210-3200103230033322-1223202122310322"></a>

<a id="canonical-1022202212223103-2033130330220101-1312320020121032-3003202023332031-2131013231310202-0312331311012112-1003310310130030-2200233013133110"></a>

## namespace property — trusted_ca_list / 322212210001 / 6

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
  }
}
```

<a id="canonical-1212333233312332-3100302131201022-3112113312211001-1030232311200102-1202222120202102-2322202102313301-1202013022200232-0011002213001333"></a>

<a id="canonical-2123200211012331-3322300302103133-3202210021010121-0032230032032010-1122233320313202-3302211210302200-3032331020311030-0012100110003232"></a>

## tenant property — trusted_ca_list / 322212210001 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1102123322002322-1320322100011030-0010320132233300-3231000232301232-2113032101033010-2022012022113021-2230130320012023-0200202301111130"></a>

<a id="canonical-0112212121331132-1202230212301002-0020303130010301-1223200302020101-0010110002203113-1313133313103310-2122310120310300-2331103111030030"></a>

## uid property — trusted_ca_list / 322212210001 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0312010120211001-2220010221300202-2333223010331322-2033032023103210-1202011123203210-1300121321033000-0130303110301013-3320132033330210"></a>

## Next pages — trusted_ca_list / 322212210001 / 9

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103101311122033-1331231002132003-0110021321201113-3211120033112001-3221322230122000-3000321102021212-2021123103012332-1210313301330300"></a>

## tls_parameters.cert_params.volterra_trusted_ca — volterra_trusted_ca / 302333113021 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="canonical-0313223202132222-0102310330211003-1122222300322232-0102123033120001-0033313312203210-0331233133120013-0033111120013333-2000213201212232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-2312000101121233-1311010101221200-2320212000011333-2333132212013121-2212233033222033-2023112022023022-2032111023123300-3033331212210312"></a>

## Direct properties — volterra_trusted_ca / 302333113021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122221101003002-1332013311022000-0231132121101301-1002010230220201-0001303121010010-0200130112322033-0230320220331330-2102200300120331"></a>

## Next pages — volterra_trusted_ca / 302333113021 / 4

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300223223231302-3231233330123300-0220102301002231-3011321013133003-1312201321331310-1321330330322113-3112330230132113-0002313303000020"></a>

## tls_parameters.common_params — common_params / 023310212030 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.common_params

<a id="canonical-0133021313320021-0231101213111210-1223331302131032-2211021112302021-0111210033212211-0312323103232222-0130100230023033-1102103031021031"></a>

Type: `"object"`. single nested block, Optional.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220211331121010-3301123031123311-2320000022200220-0312003131103303-3101103021210031-1002020020110020-0230003201230212-1321013330033010"></a>

## Direct properties — common_params / 023310212030 / 3

<a id="canonical-1133211111101322-2021033321312003-1102022311332300-1123322233023100-0323202010300030-2103211332321002-3101113233123301-1131313200112222"></a>

<a id="canonical-0201003001300011-2220121322303203-0302023000212111-2123130100311321-1100200111110333-3203320001301322-2331110331303220-0130100212012323"></a>

## cipher_suites property — common_params / 023310212030 / 4

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320231230222310-0031001003033230-1201022023020101-2133312001211130-1231033011233132-3022001000300031-3232021111022133-0203032320021213"></a>

<a id="canonical-1132233122011311-0302030102101122-0001002212031033-0313023011012333-3311223330310212-2201101210121002-0111211202301311-3121313223112132"></a>

## maximum_protocol_version property — common_params / 023310212030 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2310310133130010-3323131322122011-3023012102321000-3303220210103213-2120312131130023-0200032022020211-0222120022211013-2023022123131122"></a>

<a id="canonical-1213231112001232-0231002203112002-3212031021210102-1223003103203111-2021311020202233-0003201300113321-3311131321301313-2322203012030212"></a>

## minimum_protocol_version property — common_params / 023310212030 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
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

- [tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002): complete subsection reference.

- [validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103): complete subsection reference.

<a id="canonical-2131123011313111-1222300221133202-3212122013001133-2223001333220033-0330022100131032-0030002310032323-1230231103011320-2220220022230320"></a>

## Next pages — common_params / 023310212030 / 7

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231310013231310-0222321023022133-0222300012131200-3002223100330013-0200001331233031-2011212010113321-3132330123120222-3033321230023110"></a>

## tls_parameters.common_params.tls_certificates — tls_certificates / 330311122120 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- tls_parameters.common_params.tls_certificates

<a id="canonical-0100100311312333-3331230103112232-3233213303223112-1233021101313233-1321110021213113-1232220323210211-2132130020232133-0023113310302132"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Provider validators and defaults (from schema source):

```go
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
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130031313212112-1112300130132321-3231303032000002-0020122233221133-1011310201013123-0331302101323130-3022021332310001-1231303123130213"></a>

## Direct properties — tls_certificates / 330311122120 / 3

<a id="canonical-3101311010203203-2023303120120233-2201030100031123-0220333013232222-0030311000132223-2203021231332020-2201023332122111-3002303203102012"></a>

<a id="canonical-3213132120111020-0211310223200122-0331020101002031-2102012200032313-3322010301331103-2023222211023200-0021212013012133-1133231321303202"></a>

## certificate_url property — tls_certificates / 330311122120 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-3100030202202023-3332102123210012-3211221011112033-1010201322312023-3201300103121021-3032033210313123-1010332231032003-2220333022212312): complete subsection reference.

<a id="canonical-1022023000001122-3130210331230033-1122200232321113-2232212002202210-0031221200302202-1200113330312301-1102100322322223-3021301032320031"></a>

<a id="canonical-0231302010123112-3221113230213100-2212021232212033-1233010132101103-2332210323023210-2233322201231222-3011213010303200-1300130323132031"></a>

## description_spec property — tls_certificates / 330311122120 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-1113012131322320-0302131123233120-3101220102133203-2113133101130003-2232130312102120-0013123302000011-3200201321123120-3032320120001011): complete subsection reference.

- [private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331): complete subsection reference.

- [use_system_defaults](resources--cluster--reference--group-001.md#canonical-3312123133213002-2033110322000311-1233120302113100-3332231330331002-2223313020333233-3322023210233300-1001131001321311-3210320211002102): complete subsection reference.

<a id="canonical-0202131111323023-3210303113030223-3312013223212032-0013012030313330-3212133021230333-2103013202230221-1220032123202020-3312200002023031"></a>

## Next pages — tls_certificates / 330311122120 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-3100030202202023-3332102123210012-3211221011112033-1010201322312023-3201300103121021-3032033210313123-1010332231032003-2220333022212312)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-1113012131322320-0302131123233120-3101220102133203-2113133101130003-2232130312102120-0013123302000011-3200201321123120-3032320120001011)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--cluster--reference--group-001.md#canonical-3312123133213002-2033110322000311-1233120302113100-3332231330331002-2223313020333233-3322023210233300-1001131001321311-3210320211002102)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3100030202202023-3332102123210012-3211221011112033-1010201322312023-3201300103121021-3032033210313123-1010332231032003-2220333022212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112120330210332-0013102220002101-3111320232113031-1032333213201323-0222232233110301-1003030031321330-1031321220012122-1230323233213110"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 322010321202 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-3101031322202101-3212201230103033-2202123311010221-2101223201322020-3301002232223211-3123003233310010-1103000113333311-2031112213132123"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2323313220221223-0212113300230301-2010203221232102-1133302212121332-2021130301221001-0103322323211330-1233320100213031-1013203330000321"></a>

## Direct properties — custom_hash_algorithms / 322010321202 / 3

<a id="canonical-1230311011232210-2013031031002200-0013231023102303-3130233203233231-3301002030231301-3200121112003311-1221311213012031-3310110210330013"></a>

<a id="canonical-1200102310002120-2202011031313332-1100321120020231-1333302122123311-1300221113030023-3333032321020310-2200102222031310-0223323012023111"></a>

## hash_algorithms property — custom_hash_algorithms / 322010321202 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2203321032313231-3103301002023011-1102121300102311-3321021002211222-3223110221101013-3223020021330201-3222311020002112-2331201211212110"></a>

## Next pages — custom_hash_algorithms / 322010321202 / 5

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1113012131322320-0302131123233120-3101220102133203-2113133101130003-2232130312102120-0013123302000011-3200201321123120-3032320120001011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013332223110220-1332221130331123-0301213031220133-1203000121120323-0122200000012301-0323003230312032-3133302321201333-2012221033003232"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 020111331323 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1002301011003130-2301121301112233-3232023310303022-2202321022201023-1320221001133132-3321012311003312-1132103320123102-0131202113312312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-2110320210132100-3232123132321330-3223313023221313-3232133321323121-3330203200000202-1311310103101013-2101320233100112-3110111020333100"></a>

## Direct properties — disable_ocsp_stapling / 020111331323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230020112201100-1212221221230222-0200131103303221-2033132002131233-0003130013012033-1013201002031102-2211120110230110-0011232313202122"></a>

## Next pages — disable_ocsp_stapling / 020111331323 / 4

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200302322003003-3200030131102030-3002300311303231-2222233012310233-1130010332332221-0223110331023232-3010122100210030-1231100112132101"></a>

## tls_parameters.common_params.tls_certificates.private_key — private_key / 223202323301 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-2200322321333323-0311000201202210-3032133223101121-3332230311011312-3013132022222033-1332222213112201-0030322122312002-0022103013103002"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131303313200201-0233012331111221-0220133031221110-2332332213033012-0001101210033112-1001232312210111-3032001100122132-1232222000320301"></a>

## Direct properties — private_key / 223202323301 / 3

- [blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-1111201001212120-2100030320133102-3113220213330111-2312322131330022-1100301011312201-0210323001212333-3322202132011012-0132330022213300): complete subsection reference.

- [clear_secret_info](resources--cluster--reference--group-001.md#canonical-1003030120101221-1132033110233013-2333213312110121-0201233200102312-3032112003002113-2132331313311002-2101032023032311-3121200103302010): complete subsection reference.

<a id="canonical-2333203032012103-2330331023003032-3002031330233032-3022301101302013-2211113123223132-0331002320012222-1010010020232312-0021231000301120"></a>

## Next pages — private_key / 223202323301 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-1111201001212120-2100030320133102-3113220213330111-2312322131330022-1100301011312201-0210323001212333-3322202132011012-0132330022213300)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--reference--group-001.md#canonical-1003030120101221-1132033110233013-2333213312110121-0201233200102312-3032112003002113-2132331313311002-2101032023032311-3121200103302010)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1111201001212120-2100030320133102-3113220213330111-2312322131330022-1100301011312201-0210323001212333-3322202132011012-0132330022213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021222332101311-2301232321020331-0021210213331022-1222303220102313-2310313233002300-2010012021131023-3120301330222311-2101312002323213"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 223210131331 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3311210202122101-3302221201022001-3331323203112112-2000003003132301-3310133013000021-0212222331031220-0003033221332130-3112012011231011"></a>

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

<a id="canonical-2300313013113212-2103213031221321-2212121031101221-2223321102311022-0121121333303002-2103220200131220-0120213020130210-1333202331223213"></a>

## Direct properties — blindfold_secret_info / 223210131331 / 3

<a id="canonical-1133323020300323-2112121112330130-3002001010323102-2011303120200313-3133033102133022-3320110200201111-1020330320201221-1102110103210230"></a>

<a id="canonical-2123320233310113-1132333310203313-0121202300330033-1003211002223220-3300000201001312-2312232011301023-2212232333210030-1120113021320030"></a>

## decryption_provider property — blindfold_secret_info / 223210131331 / 4

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

<a id="canonical-3130211113302111-0100313000122211-2331010311133100-1313323223303123-3112111212213221-3320223223230023-1131213201230110-0300033322313312"></a>

<a id="canonical-0022332000031321-0301303311310221-2020022033001022-0120232021202201-0213320231023010-3120013312233111-3021013101010212-2303131020222332"></a>

## location property — blindfold_secret_info / 223210131331 / 5

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

<a id="canonical-0131002322032232-2120133203213031-3133003302323223-2133023303132200-0210221102010302-1332323333132311-2201210033133001-0233022100221313"></a>

<a id="canonical-2301030300202003-2211001231122122-3310103213202000-2222331222011001-2211033333313121-3232232223011101-3301032030211221-2230031311011113"></a>

## store_provider property — blindfold_secret_info / 223210131331 / 6

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

<a id="canonical-1110001100211300-3111330331123130-2012130301132220-0330121012210330-3233313300000133-0221032113300112-0120103022100200-1301011230030301"></a>

## Next pages — blindfold_secret_info / 223210131331 / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-1003030120101221-1132033110233013-2333213312110121-0201233200102312-3032112003002113-2132331313311002-2101032023032311-3121200103302010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020102121202023-0231122101031322-2331301002123111-0301010200330202-1101223321332220-1001030221021022-3030011303323113-1223022111231023"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 013021322110 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0323230330133100-2323213031330313-3021120112230220-1132220122200001-0030320320030033-2222221111120033-1331331323030133-3221303333013223"></a>

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

<a id="canonical-1212302333232331-1002002311111223-3220321222322133-2131012222130030-2312201033113013-2300020032210220-3002221003231302-1323230103211210"></a>

## Direct properties — clear_secret_info / 013021322110 / 3

<a id="canonical-3303210333332310-3202103000302010-1112120201321221-0332122003313230-1120222123222322-2010230302000133-3132331122122123-0100022212113303"></a>

<a id="canonical-2312310113000133-2133012233302121-1000213220212113-2110323223001202-3131200122133301-1202001320030002-1022110022300112-2203320123033212"></a>

## provider_ref property — clear_secret_info / 013021322110 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1023301333020013-3301230310130222-1031102122221233-2201003100132010-2220311011000231-3030001013333223-2210031311300221-0021203211322301"></a>

<a id="canonical-1210020202131123-2212120100010202-0232112230011231-0323300220131101-0103100201313233-1303211330113332-0003213202202130-1131122200012011"></a>

## URL property — clear_secret_info / 013021322110 / 5

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

<a id="canonical-2210131012023121-3223322123222321-2322121013233313-1100301332130012-2122201320233321-0103212303212321-0322220230212132-3021213001301031"></a>

## Next pages — clear_secret_info / 013021322110 / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)

<a id="canonical-3312123133213002-2033110322000311-1233120302113100-3332231330331002-2223313020333233-3322023210233300-1001131001321311-3210320211002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032302230110020-0321131212323221-0110213221112323-2032113211203111-3113003313130233-3203030211311311-3013302112133132-0133013112203000"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — use_system_defaults / 323213311212 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-0120003301311011-3311300110113220-3301231322330122-2331031113210112-3122223030032101-3010310323031231-3332213031013023-1103122301132320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-2202232202210313-2310011320111013-3010113301313212-0311110100220013-2312312222120210-0033113202333332-2113233210321301-0303302330122021"></a>

## Direct properties — use_system_defaults / 323213311212 / 3

This is an empty object or choice marker. It has no direct properties.
