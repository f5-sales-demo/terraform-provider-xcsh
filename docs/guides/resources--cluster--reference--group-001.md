---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- Property reference

<a id="canonical-2222022001221300-2202032033313220-2321130331022200-1030001102101013-2021223013132321-3123130212003223-2231032202201331-3111020003111112"></a>

### Direct properties for `xcsh_cluster`

<a id="canonical-2033300133303132-2210211131221132-2132102000231103-3110100213123321-2303100320002111-2202003120133011-0032211112322300-0223030322023023"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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

<a id="canonical-1121222021110000-1120211223112211-3130303011001033-3201323232001020-0221113002122330-0111123211021211-3210303330332110-1010331122233330"></a>

#### `connection_timeout` property

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

Additional upstream details:

The default value is 2 seconds.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_subset](resources--cluster--reference--group-001.md#canonical-2203130112112230-2033101230311313-1232010302031100-1220213120323211-1000330212310013-1123310312113230-2123012021101121-1101210100101030): complete subsection reference.

<a id="canonical-2231111221222012-0220120303013321-3102121310111323-3010131123030322-0133212230202303-2220130123112323-3332203030311320-2311223320210103"></a>

<a id="canonical-2112200001201011-1232031002110101-3332112333033220-0232012110232122-2322202102031301-2130311000313200-1033321320000130-2330331203333032"></a>

#### `description` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2112211320220203-3332320212133122-0202012332020321-3123321110302231-0101233212323021-0000000231220313-2332031200332131-3311131300221312"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

<a id="canonical-2030211210100022-2210101122023033-2022120112132332-3223021003312131-2220331220030131-2223330001201313-2231233323131202-0121203221223002"></a>

#### `endpoint_selection` property

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`.

Additional upstream details:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

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

<a id="canonical-0121333333020002-2333012220121311-2213002131300320-2101133113100131-0313222012232002-0303111122102120-2113200310001110-2133230020010110"></a>

#### `fallback_policy` property

Type: `"string"`. Optional, Computed.

\[Enum: NO\_FALLBACK|ANY\_ENDPOINT|DEFAULT\_SUBSET\] Enumeration for SubsetFallbackPolicy if subset
match is not found. The request fails as if the cluster had no endpoint matching the subset policy
Any cluster endpoint may be selected if the cluster had no endpoint matching the subset policy Load
balancing is done over endpoints matching.. Possible values are \`NO\_FALLBACK\`, \`ANY\_ENDPOINT\`,
\`DEFAULT\_SUBSET\`. Defaults to \`NO\_FALLBACK\`.

Additional upstream details:

Enumeration for SubsetFallbackPolicy if subset match is not found. The request fails as if the
cluster had no endpoint matching the subset policy Any cluster endpoint may be selected if the
cluster had no endpoint matching the subset policy Load balancing is done over endpoints matching
default\_subset if the cluster had no endpoint matching the subset policy.

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

<a id="canonical-1300032310032220-3322222121100112-1023020123030100-2123021232033130-3132102232331010-2312022330021001-0212300231133301-3310112221221132"></a>

#### `http_idle_timeout` property

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1213111021210320-3133020031020330-2033103231020011-3101302300012313-3213133233122330-1212233203213003-0213132011233332-3022233231301021"></a>

<a id="canonical-1221012001210130-2313123312321011-0112302012030323-2320223130330112-0100003203021322-1302010012113132-0201013020213211-2322231123000231"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1022203213321013-1013120123133003-2000031232001231-2101021113223220-1111213201132010-1301112101112130-1120233200222320-0231110012202302"></a>

<a id="canonical-1330313100111011-1122120313203313-1122203213331211-0002201331013303-0120131303233001-3120233210112010-2303022210010123-0212030210123121"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

<a id="canonical-1011230210133301-1323323321011300-0112121021033200-0331333230030221-1322203210110013-2010312021021302-0230303303032012-3013213333013102"></a>

#### `loadbalancer_algorithm` property

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`.

Additional upstream details:

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

<a id="canonical-2121112313221032-2333331203013310-1112330122322012-3031302002100020-2310000023022101-0203201232100333-2203332323033312-2033120000202330"></a>

#### `max_requests_per_connection` property

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0312022312233202-3011323300123012-1012023031110013-0030012031220013-1320302232133331-3203302303310212-3110023312303123-0120033232131311"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Cluster. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0103333203123012-1332022312303111-1231300303223023-1113303310121311-2113133033210213-2030200231312021-0210203323200302-3302100030013223"></a>

<a id="canonical-0312022003232311-0233131120030331-1011022223120122-0322023012132210-0203330311222032-0033202023032022-0132132212203013-3210232111323032"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Cluster is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
  }
}
```

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322): complete subsection reference.

- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-1222033230013012-0003313031123021-2201020000223133-2130312313211123-2000222100123110-2213333002013013-2112332102110013-3113022211303122): complete subsection reference.

- [outlier_detection](resources--cluster--reference--group-001.md#canonical-1332022310131110-0010322311202313-0132333220211213-1222030130221111-1133102220330211-3122120000022131-3202112111002123-3313332331030110): complete subsection reference.

<a id="canonical-1222321203023301-2201231330122022-0201001302231133-3301131103333010-0133110001010103-0011331332131323-0310023301030212-2230213031102200"></a>

<a id="canonical-2211330330202133-3011203300110031-1032332022022001-0103123310022111-3123031121312221-0021032320122311-3121321111001003-3221033320211213"></a>

#### `panic_threshold` property

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

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

<a id="canonical-2020312232020000-1231012303200121-1231213113010012-3321302032120020-0230012013120231-3233022213020120-2303103312031023-3131132213122331"></a>

### All schema paths for `xcsh_cluster`

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
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](resources--cluster--reference--group-002.md#canonical-3221110212020010-3222330302010201-2212232220221320-1113330020233202-0131300121301201-0201322320103000-2210031300130130-0201300232103102) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](resources--cluster--reference--group-001.md#canonical-1322120020330010-2022311322230223-2023021012220213-2311200122121020-3120330313101231-3112020332112003-0320033313112313-2303020200122021) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](resources--cluster--reference--group-001.md#canonical-1011001323111030-1210011203121131-3103332130112003-1113203002100222-1110332312300323-1030323320133131-0232022223323131-0011110032101133) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](resources--cluster--reference--group-001.md#canonical-2233212232010233-2100121100203131-1332130321011010-0122013121113210-0230000220301123-0000322213210111-3102022023201021-2111000032031131) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](resources--cluster--reference--group-002.md#canonical-2310320120131103-3032122203113332-2012320213331310-1323110111033133-0313210321012131-1302122211302102-3322230332101313-2022233001303023) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-002.md#canonical-1111221222123022-1311312321201030-3330332220030111-1200001323233030-3230102010132111-1113221320202030-3132103201021310-1200203223313222) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](resources--cluster--reference--group-002.md#canonical-0121331230300021-2131120113111110-0213302301032020-1110131321331233-1001333312301313-1312322231111333-0022130303322013-3021102311010223) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2121302331021012-1111012132210211-1110322023000113-0323302210201011-1120212101210303-3003102332313332-3123200200121322-3223311130203221) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-002.md#canonical-3200032031202331-3303310300201103-0322120323321011-0233001313110121-2121332331211113-2030232102122110-3323313103000220-1102031200010212) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](resources--cluster--reference--group-002.md#canonical-2302202330020102-1112111232032000-3100020330330013-3300231032212032-2122230221300202-2221120102213111-1223011102301131-3200122201101313) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](resources--cluster--reference--group-002.md#canonical-3112112023102223-2302020210033110-3031231230331132-3010200110110000-0222322020311310-2112313331202001-0132332120020021-3131230332331030) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](resources--cluster--reference--group-002.md#canonical-2003222222313001-2020000110322312-0320300121213332-2131213111123132-0111211202331003-0230110303320210-3200103230033322-1223202122310322) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](resources--cluster--reference--group-002.md#canonical-1212333233312332-3100302131201022-3112113312211001-1030232311200102-1202222120202102-2322202102313301-1202013022200232-0011002213001333) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](resources--cluster--reference--group-002.md#canonical-1102123322002322-1320322100011030-0010320132233300-3231000232301232-2113032101033010-2022012022113021-2230130320012023-0200202301111130) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](resources--cluster--reference--group-002.md#canonical-0003123203233123-2103211200331331-3222100313110322-1203303132102312-1132000032100011-0232333212022332-3010100303101200-0333022303211321) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](resources--cluster--reference--group-002.md#canonical-3021203102003011-2132201220102023-3103210120303002-0230212112123222-0012101202132122-0200130020330221-2221002003102323-2113110210133221) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](resources--cluster--reference--group-002.md#canonical-0313223202132222-0102310330211003-1122222300322232-0102123033120001-0033313312203210-0331233133120013-0033111120013333-2000213201212232) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-0133021313320021-0231101213111210-1223331302131032-2211021112302021-0111210033212211-0312323103232222-0130100230023033-1102103031021031) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--cluster--reference--group-002.md#canonical-1133211111101322-2021033321312003-1102022311332300-1123322233023100-0323202010300030-2103211332321002-3101113233123301-1131313200112222) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--cluster--reference--group-002.md#canonical-0320231230222310-0031001003033230-1201022023020101-2133312001211130-1231033011233132-3022001000300031-3232021111022133-0203032320021213) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--cluster--reference--group-002.md#canonical-2310310133130010-3323131322122011-3023012102321000-3303220210103213-2120312131130023-0200032022020211-0222120022211013-2023022123131122) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0100100311312333-3331230103112232-3233213303223112-1233021101313233-1321110021213113-1232220323210211-2132130020232133-0023113310302132) |
| `tls_parameters.common_params.tls_certificates.blindfold` | [tls_parameters.common_params.tls_certificates.blindfold](resources--cluster--reference--group-002.md#canonical-0321112000011100-1200221133001212-3300133002031023-0111133111210100-2221011012001131-0122210022310220-1110303223312031-3313233301121230) |
| `tls_parameters.common_params.tls_certificates.blindfold.algorithm` | [tls_parameters.common_params.tls_certificates.blindfold.algorithm](resources--cluster--reference--group-002.md#canonical-1120022001230220-3212130023100132-2303213331012120-1013100012222102-1333133013000102-3310333023300310-2030202231232101-3131223033310011) |
| `tls_parameters.common_params.tls_certificates.blindfold.certificate_file` | [tls_parameters.common_params.tls_certificates.blindfold.certificate_file](resources--cluster--reference--group-002.md#canonical-1312000320220230-0213303210222021-2123030213221313-3011301220002003-2323020001020033-1213121101232220-3101131302211200-1323011031010313) |
| `tls_parameters.common_params.tls_certificates.blindfold.certificate_pem` | [tls_parameters.common_params.tls_certificates.blindfold.certificate_pem](resources--cluster--reference--group-002.md#canonical-1212221010102220-2331132323301310-0020232103331212-2023211110231312-3323203010112100-3313020020110021-3231202032100321-0303222232200021) |
| `tls_parameters.common_params.tls_certificates.blindfold.chain_identity` | [tls_parameters.common_params.tls_certificates.blindfold.chain_identity](resources--cluster--reference--group-002.md#canonical-0221231121303211-2012000202221202-3210113131122110-1313101231213032-3201302310223030-3310223130013221-3032202301012011-0133303021302232) |
| `tls_parameters.common_params.tls_certificates.blindfold.context_digest` | [tls_parameters.common_params.tls_certificates.blindfold.context_digest](resources--cluster--reference--group-002.md#canonical-2323322021302320-1323300121312233-1302301322112310-3321013110201311-0231232133100211-3310311112111202-0002303233102210-0033220330002231) |
| `tls_parameters.common_params.tls_certificates.blindfold.encrypted_location` | [tls_parameters.common_params.tls_certificates.blindfold.encrypted_location](resources--cluster--reference--group-002.md#canonical-1102230010223111-2132111011323330-1200332320200100-2023210302012123-1203021331212200-3030110332001031-1300201320221022-3023012001301322) |
| `tls_parameters.common_params.tls_certificates.blindfold.expires_at` | [tls_parameters.common_params.tls_certificates.blindfold.expires_at](resources--cluster--reference--group-002.md#canonical-3301033023311032-2112303210203231-1113100000033233-3303211331130211-2221300033112003-0120213313013013-3212010123022012-1203101030310030) |
| `tls_parameters.common_params.tls_certificates.blindfold.fingerprint` | [tls_parameters.common_params.tls_certificates.blindfold.fingerprint](resources--cluster--reference--group-002.md#canonical-2102211333312201-1121211022003311-2212332201221033-2321311030133123-3303310022100330-0021131100303322-1001213311300311-2110320020323113) |
| `tls_parameters.common_params.tls_certificates.blindfold.id` | [tls_parameters.common_params.tls_certificates.blindfold.id](resources--cluster--reference--group-002.md#canonical-2302133103011311-0300002112300303-3001121203022010-3303333021311100-3321203312321003-2333320101122132-0322123230031333-2121320323133233) |
| `tls_parameters.common_params.tls_certificates.blindfold.material_version` | [tls_parameters.common_params.tls_certificates.blindfold.material_version](resources--cluster--reference--group-002.md#canonical-0103110233111132-3223321101120120-3030203000223302-1000223300231311-1323302032332123-3330030131100212-0202212333201100-0313101331301031) |
| `tls_parameters.common_params.tls_certificates.blindfold.passphrase_env` | [tls_parameters.common_params.tls_certificates.blindfold.passphrase_env](resources--cluster--reference--group-002.md#canonical-2320212313111203-2123022322132032-1031102321021001-2122013000000203-0000102212113313-2011320223030311-2222122133302203-1033000313210131) |
| `tls_parameters.common_params.tls_certificates.blindfold.passphrase_wo` | [tls_parameters.common_params.tls_certificates.blindfold.passphrase_wo](resources--cluster--reference--group-002.md#canonical-3322023012210021-3002000313223202-2100121310033232-0201201303310200-2122101213323320-1322123221300113-3313311120100022-3001310002012011) |
| `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_file` | [tls_parameters.common_params.tls_certificates.blindfold.pkcs12_file](resources--cluster--reference--group-002.md#canonical-2303313313030312-0333130002331112-0011111102121330-3102031101132322-1330123013111131-0110110001231003-2131313233303103-1030211102213121) |
| `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_wo` | [tls_parameters.common_params.tls_certificates.blindfold.pkcs12_wo](resources--cluster--reference--group-002.md#canonical-3221020220023122-3013310131032323-0321000200123303-1001321103222010-0031110233110320-2033223030303000-1321313113010013-1300212100100133) |
| `tls_parameters.common_params.tls_certificates.blindfold.policy` | [tls_parameters.common_params.tls_certificates.blindfold.policy](resources--cluster--reference--group-002.md#canonical-3302112130112133-1300030310233232-0302300023130023-1113303211211112-2232033013110322-0233231300221010-3111202222313202-3022110123223233) |
| `tls_parameters.common_params.tls_certificates.blindfold.prepared_identity` | [tls_parameters.common_params.tls_certificates.blindfold.prepared_identity](resources--cluster--reference--group-002.md#canonical-3000111021210222-2001112031022230-0133033010320232-3301220333232113-1002323230203100-3110032310213103-3000313103010113-3000302223020130) |
| `tls_parameters.common_params.tls_certificates.blindfold.private_key_file` | [tls_parameters.common_params.tls_certificates.blindfold.private_key_file](resources--cluster--reference--group-002.md#canonical-2220232321023203-2022332030323123-3022131222230112-3320122021121123-3122133322121221-0002301302222200-2101200000101101-2333033321320022) |
| `tls_parameters.common_params.tls_certificates.blindfold.private_key_wo` | [tls_parameters.common_params.tls_certificates.blindfold.private_key_wo](resources--cluster--reference--group-002.md#canonical-2000031120123000-3212230032222023-2213000001012022-1132311133003102-0300121312010031-1222021123203010-3110330212103111-0003302300221330) |
| `tls_parameters.common_params.tls_certificates.blindfold.spki_identity` | [tls_parameters.common_params.tls_certificates.blindfold.spki_identity](resources--cluster--reference--group-002.md#canonical-3130013211301211-0300200030002030-3100110002113101-0011211030000101-2220213113210011-0302123132012133-1111323303133010-1033031322001103) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--cluster--reference--group-002.md#canonical-3101311010203203-2023303120120233-2201030100031123-0220333013232222-0030311000132223-2203021231332020-2201023332122111-3002303203102012) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--cluster--reference--group-002.md#canonical-3101031322202101-3212201230103033-2202123311010221-2101223201322020-3301002232223211-3123003233310010-1103000113333311-2031112213132123) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--cluster--reference--group-002.md#canonical-1230311011232210-2013031031002200-0013231023102303-3130233203233231-3301002030231301-3200121112003311-1221311213012031-3310110210330013) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--cluster--reference--group-002.md#canonical-1022023000001122-3130210331230033-1122200232321113-2232212002202210-0031221200302202-1200113330312301-1102100322322223-3021301032320031) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--cluster--reference--group-002.md#canonical-1002301011003130-2301121301112233-3232023310303022-2202321022201023-1320221001133132-3321012311003312-1132103320123102-0131202113312312) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-002.md#canonical-2200322321333323-0311000201202210-3032133223101121-3332230311011312-3013132022222033-1332222213112201-0030322122312002-0022103013103002) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--reference--group-002.md#canonical-3311210202122101-3302221201022001-3331323203112112-2000003003132301-3310133013000021-0212222331031220-0003033221332130-3112012011231011) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--cluster--reference--group-002.md#canonical-1133323020300323-2112121112330130-3002001010323102-2011303120200313-3133033102133022-3320110200201111-1020330320201221-1102110103210230) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--cluster--reference--group-002.md#canonical-3130211113302111-0100313000122211-2331010311133100-1313323223303123-3112111212213221-3320223223230023-1131213201230110-0300033322313312) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--cluster--reference--group-002.md#canonical-0131002322032232-2120133203213031-3133003302323223-2133023303132200-0210221102010302-1332323333132311-2201210033133001-0233022100221313) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--reference--group-002.md#canonical-0323230330133100-2323213031330313-3021120112230220-1132220122200001-0030320320030033-2222221111120033-1331331323030133-3221303333013223) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--cluster--reference--group-002.md#canonical-3303210333332310-3202103000302010-1112120201321221-0332122003313230-1120222123222322-2010230302000133-3132331122122123-0100022212113303) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--cluster--reference--group-002.md#canonical-1023301333020013-3301230310130222-1031102122221233-2201003100132010-2220311011000231-3030001013333223-2210031311300221-0021203211322301) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--cluster--reference--group-002.md#canonical-0120003301311011-3311300110113220-3301231322330122-2331031113210112-3122223030032101-3010310323031231-3332213031013023-1103122301132320) |
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

<a id="canonical-2322231020220022-3201131011222101-3111130311210132-1100300313031231-3130321322022012-1103100132333333-0022230220220021-1222033112300230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `auto_http_config` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- auto_http_config

<a id="canonical-3323112120020132-3032333223220020-2313013301213023-1203123220131011-3301132012131101-2211122220012013-2231213113010303-1211200032100120"></a>

Type: `["object", {}]`. Optional.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

OneOf alternatives in this subsection:

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-3323112120020132-3032333223220020-2313013301213023-1203123220131011-3301132012131101-2211122220012013-2231213113010303-1211200032100120)
- [http1_config](resources--cluster--reference--group-001.md#canonical-2312200301003200-0332313230231001-2033321301231233-3121000302131031-0021223023033331-0023000332232011-3210031320011001-3321003031033200)
- [http2_options](resources--cluster--reference--group-001.md#canonical-2101222322313120-0223221211121320-0120223111021203-1203032232020011-3221211003003321-1233201120203310-2233131030003202-0001113132332310)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
auto_http_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221020232120031-3001013211111132-1121322102221222-0312203310302010-0011203101220332-1323012230031023-2230330130032331-3233102132202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `circuit_breaker` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- circuit_breaker

<a id="canonical-0301321011222303-1013230032313131-3223102023222311-1011002131010222-2132021212111111-1100233233123302-1030022111310002-2300132220013102"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3112110221213312-0303103323211011-1132001331231332-3023132220132011-0101123201122032-1230203301220320-1313103230022033-1212100020313001"></a>

### Direct properties for `circuit_breaker`

<a id="canonical-3023213033131202-0230102002230323-0221010221320300-1011233103101323-2311230112223231-2331210302223103-3230221112331201-1322321333003312"></a>

#### `circuit_breaker.connection_limit` property

Type: `"number"`. Optional.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0112023131013210-3311313231003103-1131010010003221-0010113111301321-0113322112303203-2211211111101022-0102330233301020-0032330101201021"></a>

<a id="canonical-2222013130322131-1112213012111010-1001223213320121-0313211122312200-0222132020111021-1311330110000301-2223030103313203-3101202200022010"></a>

#### `circuit_breaker.max_requests` property

Type: `"number"`. Optional.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-1220222100223023-0333333322012211-1103202233100301-0123033032331102-3123212001231223-0302011020020011-2023222030102100-1302000002033203"></a>

<a id="canonical-1301020112013023-3033123330230003-2003303303233301-0200103213020232-3330100012013303-3103031111013003-0121023113202012-3030121333133302"></a>

#### `circuit_breaker.pending_requests` property

Type: `"number"`. Optional.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-3221310022012020-1201213323233110-2200231012031013-0222323121123030-2132123201011303-1112212131302312-0330022321010322-2310202311001033"></a>

<a id="canonical-0132112211213002-1212132311210213-1220033001221303-0113102201121010-3000323100032300-0300130332220222-0030111322301031-3310002220100131"></a>

#### `circuit_breaker.priority` property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Additional upstream details:

Priority routing for each request. Default routing mechanism High-Priority routing mechanism.

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

<a id="canonical-2103131310121211-2322300132331011-2112003122031302-1131311011120201-1202000330113030-2303102030131223-2331120230321313-0233212222322213"></a>

#### `circuit_breaker.retries` property

Type: `"number"`. Optional.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2203130112112230-2033101230311313-1232010302031100-1220213120323211-1000330212310013-1123310312113230-2123012021101121-1101210100101030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_subset` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- default_subset

<a id="canonical-3330211002030032-0023131323312110-3112022120010011-3330130001232113-3211210013001002-2220101333031222-3212113303020022-2212100113013003"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 32
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "32"
    }
  },
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_proxy_protocol` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- disable_proxy_protocol

<a id="canonical-2112122012302122-2211122203320030-3323120131210302-0212130210310100-3203011310202032-1010122322002003-2113202312210323-2103311333002323"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

OneOf alternatives in this subsection:

- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-2112122012302122-2211122203320030-3323120131210302-0212130210310100-3203011310202032-1010122322002003-2113202312210323-2103311333002323)
- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0123022331202012-2203321021103222-1132311223332321-2133323312233121-1221303212030001-2113121211232103-2131113100032230-3031302130011201)
- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-3102031231210012-1020030322123212-3121231203301223-0210133013030310-1320002012120113-2113233320332201-2002120122000132-2133211112203333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130002222020001-0232133310310030-3303113101232230-1302101000000023-2102221011311020-3031020222211032-2220002000132123-3001122011123303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_subsets` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- endpoint_subsets

<a id="canonical-0222103230201322-3020333102013131-0031310021001101-0103323022200022-2201230310121230-2033111312212102-3213210302322322-0232320231220230"></a>

Type: `"object"`. list nested block, Optional.

Configure endpoint groups based on metadata labels for traffic routing. Supports weighted
distribution and session affinity across labeled endpoints.

Additional upstream details:

Cluster may be configured to divide its endpoints into subsets based on metadata attached to the
endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected
by the load balancer. Endpoint\_subsets is list of subsets for this cluster. Each entry in this list
has definition for a subset (which is collection of keys)

During routing, the route’s metadata match configuration is used to find a specific subset. If there
is a subset with the exact keys and values specified by the route, the subset is used for load
balancing. Otherwise, the fallback policy is used. The cluster’s subset configuration must,
therefore, contain a definition that has the same keys as a given route in order for subset load
balancing to occur. Example:

RouteConfig

routes: &#8203;- match: &#8203;- headers: \[\] path: path: /1.log query\_params: \[\]
routeDestination: destinations: &#8203;- cluster: &#8203;- kind: cluster.object uid:
00000000-0000-4000-8000-0b50b89d07a2 endpointSubsets: site: india

EndpointConfig

metadata: labels: deployment: debug site: india name: end-1 uid: end-1

ClusterConfig

gcSpec: defaultSubset: stage: production fallbackPolicy: DEFAULT\_SUBSET endpointSubsets: &#8203;-
keys: &#8203;- site &#8203;- keys: &#8203;- stage &#8203;- app

Assume the below endpoints are defined and associated with the cluster. Endpoint Labels --------
&#8203;------

ep1 stage: production, site: india ep2 stage: deployment, site: us ep3 stage: production, app: hr
ep4 site: india

The following table describes some routes and the result of their application to the cluster. The
subset definition for cluster is assumed to be same as given above in the ClusterConfig section

RouteMatch Criteria Subset Reason ------------------- ------ ------

site: india ep1, ep4 Subset of endpoints selected site: us ep2 Subset of endpoints selected app: hr
ep1, ep3 Fallback: No subset selector for "app" alone stage: production, app: hr ep3 Subset of
endpoints selected other: x ep1, ep3 Fallback: No subset selector for “other” (none) ep1, ep3
Fallback: No subset requested.

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

<a id="canonical-2133120123020212-2010232103302320-2220012013110302-2122200323210220-1031310210310030-2310332102321312-1011310103223031-3233001012023130"></a>

### Direct properties for `endpoint_subsets`

<a id="canonical-0232111112131132-0332131003311213-3331230103202123-0012131220202221-3323302212100112-0323023303002202-1103003321131301-3220031031200233"></a>

#### `endpoint_subsets.keys` property

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

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

<a id="canonical-1122033200033300-2212233112332301-0113231231121220-3011100310230301-2211231323321310-2322020033013220-2001321233013220-1011003202200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoints` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- endpoints

<a id="canonical-2313301312133322-1101011333010220-3232202010313231-1312010322012222-0030212033021210-2020121132002013-2312311020303232-3023032132301003"></a>

Type: `"object"`. list nested block, Optional.

List of endpoints for this cluster.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0103303120320221-2033132312323033-3312102021333220-0312022213232301-1103132323333110-1212302001111320-1001103200012031-1303100310103001"></a>

### Direct properties for `endpoints`

<a id="canonical-1223001301301311-3131000332011222-0300210101202023-2332022201231230-3123013233320011-0320023202312203-2133313233210013-1213023312122313"></a>

#### `endpoints.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3220100120232102-1312202132000312-3013032031000330-1332230001130120-0213312011121012-0311323033010233-1333203032100313-2230223322332331"></a>

#### `endpoints.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0333322121031212-2231220112103220-3333003133311313-2031203012220321-3233121330202113-2032130322111033-0320001331223112-1221002033101210"></a>

#### `endpoints.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  }
}
```

<a id="canonical-1231011300033202-0130200023200222-1002322201021022-0112201223201111-3121033212200011-2310100201230121-1011212200200103-3203133111222112"></a>

<a id="canonical-1230133002332123-3322311121312220-1222001303003133-1302131033002111-0212330131001022-0302122100333112-2331313022013331-2130303101110012"></a>

#### `endpoints.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2013332332131322-0100020303313003-2123201321200123-0020113223121300-2030322232233300-2030220120121220-3232203202111130-1101231131213020"></a>

#### `endpoints.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3333112022010122-1123221201321130-3332220012202132-3212101010213330-3221010130203101-2233333323313123-3210103231011213-1002123323200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `health_checks` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- health_checks

<a id="canonical-3111322320100103-2202130102211333-3301300131300001-2233002230222232-2100330311003020-3302333001310332-1220220213023302-2233001120333202"></a>

Type: `"object"`. list nested block, Optional.

Health check configuration for backend monitoring.

Additional upstream details:

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

<a id="canonical-3122322132111233-3031200211113032-0132002123303300-0311130001130000-3302013210321210-3132010131231113-3112001230230103-1333032301100013"></a>

### Direct properties for `health_checks`

<a id="canonical-3100312022113011-2013012321200332-2230030000221010-2002023220301022-0012223101022001-1211201010321321-2033213331223302-3132332031100003"></a>

#### `health_checks.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2210130212032131-0322111121203011-2331213302022121-0300012021111121-2323300312022020-1133320233103132-0101210020301000-0012230302023003"></a>

#### `health_checks.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3213113302222320-0102102131012030-2300000312323133-3203111332321221-1010322001232101-0322321133313321-3010332331300120-2002021020030211"></a>

#### `health_checks.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  }
}
```

<a id="canonical-2113221031313011-0032300122031233-2003232132133323-1031121221210133-2022111111301031-2233332110323230-1203213133013033-3201203021303020"></a>

<a id="canonical-2113331110113212-2213300012312231-1132220203222131-2003110001301110-3112031010121121-3232132103201030-0213023221332322-1200031003333230"></a>

#### `health_checks.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1230321301200330-0333202332300001-0031021033100002-1022223000023330-1321231312303310-3031133300033103-0323000023122010-2033021133330203"></a>

#### `health_checks.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config` properties

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

<a id="canonical-1002311122300021-3113001233233232-2302233002202120-1220102132223121-3231330301003022-1313322002311022-2012000101022211-1030303332030200"></a>

### Direct properties for `http1_config`

- [header_transformation](resources--cluster--reference--group-001.md#canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011): complete subsection reference.

<a id="canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [http1_config](resources--cluster--reference--group-001.md#canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133)
- http1_config.header_transformation

<a id="canonical-0011120112132201-0131112310231302-0013030120301332-0302211232312022-0030003331121101-0002320120101021-3110032230313323-2030012011120210"></a>

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

<a id="canonical-0310303333312101-1022011103331103-2301323202123323-0012130101012131-1020230231021120-2203200313011012-3031222312002230-0101232202200312"></a>

### Direct properties for `http1_config.header_transformation`

- [default_header_transformation](resources--cluster--reference--group-001.md#canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201): complete subsection reference.

- [preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131): complete subsection reference.

- [proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120): complete subsection reference.

<a id="canonical-3210111201011333-3303233101013001-1220302301313223-2002332233223022-0331321212130333-2311311023333103-3300021323021322-0323322311021201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.default_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233030001120322-0200313113322200-1313032121203330-2322320303213222-1213122030323010-1013033223110220-0331000121023311-1202330301032131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.preserve_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231100321330203-1210202101302300-1303003222032121-0312313011333021-3211020021011222-2122211033331200-2002101332232122-1202020021103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.proper_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122100310003003-3113331120023312-1013121113333323-1300201130311202-1110110033320221-1200012120120123-3131112312310130-0313133200131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http2_options` properties

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

<a id="canonical-3103103120122322-1212002233310130-1212010332130320-1022310011300230-3111333000113013-2210311010112232-1230223023221022-1021022113111032"></a>

### Direct properties for `http2_options`

<a id="canonical-0330023201103012-1101112302030120-3021100120221331-2200103023300310-3313013111003123-0013213022032212-1020323330133002-0301101131202313"></a>

#### `http2_options.enabled` property

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

<a id="canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_panic_threshold` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- no_panic_threshold

<a id="canonical-0030112232003133-1300001130002032-1013121032030002-3231003120323012-1112201122312311-2132222303011300-1003203231102100-0002311313103221"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_panic\_threshold, panic\_threshold; Default: no\_panic\_threshold\] Configuration
parameter for no panic threshold.

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

OneOf alternatives in this subsection:

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-0030112232003133-1300001130002032-1013121032030002-3231003120323012-1112201122312311-2132222303011300-1003203231102100-0002311313103221)
- [panic_threshold](resources--cluster--reference--group-001.md#canonical-1222321203023301-2201231330122022-0201001302231133-3301131103333010-0133110001010103-0011331332131323-0310023301030212-2230213031102200)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_panic_threshold = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222033230013012-0003313031123021-2201020000223133-2130312313211123-2000222100123110-2213333002013013-2112332102110013-3113022211303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- no_request_limit_per_connection

<a id="canonical-2313200231201311-0312120113233312-3033023223102030-1002011022020020-0100333103112202-2031302223131011-0302032220233323-1300123011110120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332022310131110-0010322311202313-0132333220211213-1222030130221111-1133102220330211-3122120000022131-3202112111002123-3313332331030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `outlier_detection` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- outlier_detection

<a id="canonical-2000301313012312-2231230331211202-3221001132210122-3001333312022101-3132202023010032-0001112303133132-2011012203323303-1302311011310303"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2222101100311230-0213202122033322-0303233113333233-0030131323132120-0211111011303203-2122222121132320-1013311232022022-3331231300110020"></a>

### Direct properties for `outlier_detection`

<a id="canonical-1200132320100002-1232011030320303-3223333122203302-3133301232030301-2213320121333300-1311201031100300-3313201020301310-2321132303132002"></a>

#### `outlier_detection.base_ejection_time` property

Type: `"number"`. Optional.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-3012301210332013-0130232133101223-2312320223222222-1002213211222233-1023103011032330-0220123233133231-2210101101100301-0103010212121323"></a>

<a id="canonical-2221010212010330-0330003130110113-2122231210032122-1211031020301011-0012311020222130-3202011110102113-1303113323012002-1222231012102020"></a>

#### `outlier_detection.consecutive_5xx` property

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates the
number of consecutive 5xx responses required before a consecutive 5xx ejection occurs. Defaults to
5.

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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-0013033223010221-0313111303222220-3311311232300013-1300331202310110-3001001023310031-0110123213233112-2010013101131210-1031001311013302"></a>

<a id="canonical-0022313103303000-2100333210303301-0032103111222032-3212013333030332-0103330133302120-3201131003031130-3102232302303311-2131003311231322"></a>

#### `outlier_detection.consecutive_gateway_failure` property

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-3021310300001031-1331003323020313-3023020321322002-3103303312031332-2113003130222000-2300310311031002-0101202031012132-1220300231110301"></a>

<a id="canonical-0130012233032221-2130021011321001-0321101301133233-1233111303112023-0131312321202112-0031011123220022-2102111202312120-2230323202110002"></a>

#### `outlier_detection.interval` property

Type: `"number"`. Optional.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Additional upstream details:

Defaults to 10000ms or 10s. Specified in milliseconds.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2221232102211320-2013320332031330-1211031200130301-2322023022223303-2012133022231300-1303202101330112-3032001221120320-0311212202302121"></a>

#### `outlier_detection.max_ejection_percent` property

Type: `"number"`. Optional.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Additional upstream details:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0033223222002132-3130113312332000-3110121121032222-0311212300220331-2111030221111102-1103202003220133-3230012120130011-3222201230122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_protocol_v1` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- proxy_protocol_v1

<a id="canonical-0123022331202012-2203321021103222-1132311223332321-2133323312233121-1221303212030001-2113121211232103-2131113100032230-3031302130011201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

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
proxy_protocol_v1 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030230022303201-0210032031113222-1010300000311002-1033013033031211-1231331022232311-0022211013203320-1012130030000133-0133112211302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_protocol_v2` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- proxy_protocol_v2

<a id="canonical-3102031231210012-1020030322123212-3121231203301223-0210133013030310-1320002012120113-2113233320332201-2002120122000132-2133211112203333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v2.

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
proxy_protocol_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330030023321302-2111121231202211-3222113200022000-2010222000312212-1332323122330130-1131000321201000-2221222311312130-0300331200322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-2102010102330223-3102300121101233-1120202112213300-3113032123123003-0311123101203230-3221031021113031-2111001121101201-1000203003322113"></a>

### Direct properties for `timeouts`

<a id="canonical-1130323101333223-1010032213321333-1133300232113203-3122303333111000-1110021103322021-2121111122322111-3122213313311220-1032011003303232"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3012112101011021-0320033303312331-1133031002211310-0021201303310222-1320332103330012-0031023330222121-0102311022131332-3012303221101302"></a>

<a id="canonical-3103012313200202-2203100330222332-0121020230330310-2112001131103331-2323011001230311-2112120133112112-2012010322113020-2000010331321000"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2131222203323332-1131121222031022-3320212122132111-0023002130320221-1202001131333213-1312311230101111-3322022200231002-0033330230033000"></a>

<a id="canonical-3212312113200333-2013312322332021-1313300310130322-1222302300011312-1333033101233211-3322331000110120-3201121313122320-1323023333110312"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3200033303013012-1121131301202102-0213012021212131-1120220302011132-1333330301000231-2112111101222310-3012003221303032-1210103322233330"></a>

<a id="canonical-2220201120102310-3201022303332031-0113113032221331-3010320301032213-0211121323322130-1203332202313223-1003213302230101-1221123210213233"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- tls_parameters

<a id="canonical-3131202220133102-1013100100131301-2132120313213100-0321321003131203-2030330031211213-1112033032013313-0103110131021322-1120220330212120"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

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

<a id="canonical-1103130112210133-0003131112020023-0233320312222103-2230230312033310-2102233120000321-3102003020200221-2022212122020031-1223300102212010"></a>

### Direct properties for `tls_parameters`

- [cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022): complete subsection reference.

- [common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031): complete subsection reference.

- [default_session_key_caching](resources--cluster--reference--group-002.md#canonical-3111210320132103-3023301033212102-3010212200103223-1330312002311102-2203031033111312-1310030010122203-0223131200231111-2130111222200133): complete subsection reference.

- [disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310): complete subsection reference.

- [disable_sni](resources--cluster--reference--group-002.md#canonical-1231201030310032-2232133130332230-3102203310013132-0231112133233202-2112002222030130-0312121223003313-1302023203330231-0310011232322132): complete subsection reference.

<a id="canonical-1202102322210022-2133111013230000-2213200211021111-1321202313202311-2100312202102212-0231231110033323-0112003222113231-3311333131013331"></a>

<a id="canonical-0000221322210310-2103122210222222-0102232031021333-1131220222021321-2121122232321300-2102130122211330-1021320133001320-3200331222133100"></a>

#### `tls_parameters.max_session_keys` property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0222332200030023-3222121231101312-1201311233220000-3131203231021313-3331333332001330-2131232223322103-0113330331331123-0303013212031211"></a>

#### `tls_parameters.sni` property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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

<a id="canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.cert_params

<a id="canonical-2321212120333130-0103220022113300-1022221001323302-3313232232210012-0300331021202221-3111120220030331-3102220023000003-1100122032001031"></a>

Type: `"object"`. single nested block, Optional.

Certificate Parameters for authentication, TLS ciphers, and trust store.

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

<a id="canonical-2233002301202220-1221132313213202-3331013032222303-3311213320233201-0202213021322212-2130003003220222-2130023100100023-0103231321003231"></a>

### Direct properties for `tls_parameters.cert_params`

- [certificates](resources--cluster--reference--group-001.md#canonical-3231221112110221-2222223230122221-0123030212331110-3102133330132221-1323200221220213-2331121203031321-2322013132001213-1231033011033303): complete subsection reference.

<a id="canonical-1322120020330010-2022311322230223-2023021012220213-2311200122121020-3120330313101231-3112020332112003-0320033313112313-2303020200122021"></a>

<a id="canonical-0233010110221023-1231211032003223-3113022312331313-2321022322232201-1203212310222100-3123030113032120-1001300303132102-3130222103023021"></a>

#### `tls_parameters.cert_params.cipher_suites` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0132031101100110-0021303012003232-0223112232303332-3111031221321300-0223123220011021-1333301112231220-0300311012330233-1323201113033200"></a>

#### `tls_parameters.cert_params.maximum_protocol_version` property

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

<a id="canonical-2233212232010233-2100121100203131-1332130321011010-0122013121113210-0230000220301123-0000322213210111-3102022023201021-2111000032031131"></a>

<a id="canonical-1013200212100323-1301330222122220-2302321100200330-3202022230031020-3003313233332302-3030211232233321-1311212100032330-2030322011200020"></a>

#### `tls_parameters.cert_params.minimum_protocol_version` property

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

- [skip_server_verification](resources--cluster--reference--group-002.md#canonical-3230113212220110-1311312112103302-1122313301110122-2300101023102300-2311300223331201-3302320331023023-1100030020113322-2020202131030311): complete subsection reference.

- [tls_validation_params](resources--cluster--reference--group-002.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111): complete subsection reference.

- [volterra_trusted_ca](resources--cluster--reference--group-002.md#canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120): complete subsection reference.

<a id="canonical-3231221112110221-2222223230122221-0123030212331110-3102133330132221-1323200221220213-2331121203031321-2322013132001213-1231033011033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.certificates` properties

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

<a id="canonical-1303123320123211-0213302333022222-2112321330202231-1132303103210102-1302333322122021-1312021023111330-3133312331123022-0013333021000023"></a>

### Direct properties for `tls_parameters.cert_params.certificates`

<a id="canonical-3332033231201023-1111103301101203-3321332200012230-0303020333233011-2112211020021031-1113330101100202-2332310110112310-1012221120113012"></a>

#### `tls_parameters.cert_params.certificates.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3003210221231323-2101123310112300-0102210301120030-3122312312023310-2031101330303302-0232000022223210-3202013100032112-1213222211030322"></a>

#### `tls_parameters.cert_params.certificates.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2232002002212033-1310001122111030-0012001231322132-3200130123000210-1332003302033110-2331300133212313-2310212203233203-3333012231202333"></a>

#### `tls_parameters.cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  }
}
```

<a id="canonical-2200233121210111-0231301103300202-2131220233331333-3100300100302012-0133000120213011-1310030020311310-3223120023121303-0313032333000013"></a>
