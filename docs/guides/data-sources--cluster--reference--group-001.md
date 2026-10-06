---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- Property reference

<a id="canonical-0231222131130233-2022213311132201-1213031021311230-0300033313200012-2301100021311320-0133301033100211-1223022123321133-1011303220313203"></a>

### Direct properties for `xcsh_cluster`

<a id="canonical-1101320303220322-2023312031233122-3200323130133222-2303101322310031-3022013033121332-0303110321220020-2203213332330023-0321121213130213"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-0121220033001011-3230302333131331-1120220130301100-2331202102323301-3103232211000023-1201210120201312-0202301131223322-2021003222112331): complete subsection reference.

- [circuit_breaker](data-sources--cluster--reference--group-001.md#canonical-1030311322100312-0021313011330010-0011030103232220-1333113213133000-0223021121231313-2213032200211113-1231102101131311-3333002013010212): complete subsection reference.

<a id="canonical-2000321022302002-0201003312220211-3211032212031002-3012101311231213-2201213033000311-3300322330121112-2002223310232212-2123201121201022"></a>

<a id="canonical-1302311332023003-2320103232322000-2130120200311120-3112030022002322-2201010031010102-0111110330231203-0223020233222300-0222322012302200"></a>

#### `connection_timeout` property

Type: `"number"`. Computed.

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

- [default_subset](data-sources--cluster--reference--group-001.md#canonical-3211100033212002-3323310310011303-3231020001323310-3133321001312311-3031123102313023-3321012012100132-3231002312200331-1303113213312131): complete subsection reference.

<a id="canonical-0220332033330312-0303130021321002-0233300231002323-0102012331232032-0012003322230200-0211110023213101-1300232021201331-3200321313112333"></a>

<a id="canonical-0130133132103003-3222111111311011-1210121220110233-0011100002231011-3201320020113120-2033323301322033-0102020020211301-2120020210021003"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Cluster.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-3032021323222023-2113331010130133-1220333211120201-1312220223131131-1322033303023321-2223300130100302-1001220312222332-1013112312332131): complete subsection reference.

<a id="canonical-0031300321021321-3030223300321312-2111112120001232-2222200230021102-1121032000123231-0132132132113133-0021003213122203-3313103323321123"></a>

<a id="canonical-1313222130123321-1113200001310133-1230312203231200-2003202322013013-3320221300220311-1002111313032322-3112232133133322-3010111003312013"></a>

#### `endpoint_selection` property

Type: `"string"`. Computed.

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

- [endpoint_subsets](data-sources--cluster--reference--group-001.md#canonical-0020130102112320-3222131012320000-2113321330033001-3012022222130010-1230112302230233-2303032003222232-1323001031233101-2003001000131130): complete subsection reference.

- [endpoints](data-sources--cluster--reference--group-001.md#canonical-0112011310121212-0110003113310222-2323213222023033-0321211112303230-2213202000122321-0101220230120010-3221102303012330-2320002330220311): complete subsection reference.

<a id="canonical-0202030220003233-1131100303233102-1111232010300230-0133131013222022-0213322330210111-2222203002123333-2022113133111212-0310332021232322"></a>

<a id="canonical-0130333031223031-0213020001230123-1130220301021031-2022323313030120-2033311033131313-0213013310133310-0000023003013233-0021301132302001"></a>

#### `fallback_policy` property

Type: `"string"`. Computed.

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

- [health_checks](data-sources--cluster--reference--group-001.md#canonical-3202012113122110-1021232210312333-3113020021130130-2321033332112012-3321322013313120-0300123320333310-2313332022323122-1222103012303003): complete subsection reference.

- [http1_config](data-sources--cluster--reference--group-001.md#canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121): complete subsection reference.

- [http2_options](data-sources--cluster--reference--group-001.md#canonical-1030010231200022-3322131033122210-1230033020312003-0111313202211202-3322302012332230-1211221323031332-1123230123011022-3232203321032103): complete subsection reference.

<a id="canonical-1332121023203333-3131100012112321-0011311020321003-2210012331303300-0110300000212102-0300100022221310-3132222102100112-0032212200101021"></a>

<a id="canonical-1102023232321110-3313100300303103-0112121213322030-1030211321023233-2201133331013231-0203112222200131-1213122130333013-0331301022031030"></a>

#### `http_idle_timeout` property

Type: `"number"`. Computed.

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

<a id="canonical-0332221211211103-2301212301233032-0130113113122310-2322233312002301-3033023311310312-2013113023120113-0323313330210102-2331220110311200"></a>

<a id="canonical-3303102010113332-2010103211123322-1303313211212131-0110113323133301-1321130121203221-3113300033030300-2220212110331233-1122102110213311"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0030203133223130-2132103030032203-0031233032121210-3320311320023211-3311003011110113-1013202320123321-2123132322011022-2311110232210031"></a>

<a id="canonical-0202231312132131-0333222020132003-0110313232001330-1320012303223203-2130022230323210-1111311101312010-0101202021332132-1323221013201132"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-3112211023121121-3033100230002232-2330230332202013-1010210123301001-0213300122020121-0033203010110312-1112312120012010-1101110131311020"></a>

<a id="canonical-3111120200130111-0013110022001033-1303103032223113-0232313302213020-2330033321112031-0003230323133110-0023000001003010-1203112331022303"></a>

#### `loadbalancer_algorithm` property

Type: `"string"`. Computed.

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

<a id="canonical-2110110221221120-0222132101222223-1233310103210203-0311032123200010-0003101011233201-3131211122322023-1312120032023221-3133013022001000"></a>

<a id="canonical-2331223003302321-2221300303203131-2033323011021310-2333330023332232-0012333102301133-3010112113311223-2302010332022212-2031321323203132"></a>

#### `max_requests_per_connection` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

OneOf alternatives in this subsection:

- [max_requests_per_connection](data-sources--cluster--reference--group-001.md#canonical-2110110221221120-0222132101222223-1233310103210203-0311032123200010-0003101011233201-3131211122322023-1312120032023221-3133013022001000)
- [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-3020200133211313-0212102330033311-3020210112303000-1112000320311103-1000001100200110-1323303221203003-2313202112212110-2303232031200223)

Select alternatives according to the provider validators above.

<a id="canonical-2331133322311033-2110131310221313-1023010132322331-1122201012303301-0111020223200330-0202000322230122-3300300003213001-0310222112000020"></a>

<a id="canonical-1202201232200121-3233213001302200-1102010013133211-2210321003112121-1020201020202010-1100331030021221-3233133021103101-3212331112100110"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Cluster.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0000110003301331-3233210230312112-3012110310312001-3320113032203221-3213103331001223-1102031332202103-0221330311133213-1023211003231130"></a>

<a id="canonical-0020203112012010-0131133233223120-2332003202112312-1300302013021003-2312012103033110-0223013121012233-0012330000112311-3120122230311110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Cluster exists.

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

- [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-2211301230101320-2002021313220200-3310210323032101-1213102313203211-1100223020000320-0300213121222230-0133300202032003-3232212303130223): complete subsection reference.

- [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-2203320230000303-2221322103212223-2212111131011213-2232212212033022-3230300202221320-0033222323113312-2003021313202221-1223013113002011): complete subsection reference.

- [outlier_detection](data-sources--cluster--reference--group-001.md#canonical-1221013233311223-2121003102312233-2021333310030322-3330222031201202-0121032031312212-3000221001110310-2311002311301120-3232021321201122): complete subsection reference.

<a id="canonical-1331302130111330-1123231310213032-0322001313312222-1330000303121111-1210331121110203-3333213312222301-3203233113121223-0331223123123211"></a>

<a id="canonical-1232302230211301-1011102120331211-1112201331112120-2210233121021023-1133310003020212-2220113132002123-3222110123322212-0322223220300011"></a>

#### `panic_threshold` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-1100331322301132-0320031111010231-0332200212032200-2303202012003200-2313311022333003-1302100231200002-3223231030010220-2132022302330022): complete subsection reference.

- [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-1023012213100031-0232220222310012-3133010122101011-2312120323230312-0310300102000210-2021021123302201-1321110221121130-1220101111021111): complete subsection reference.

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-0313001211201301-0220311332132120-1031132020332300-1331101230303030-2110211232132032-3032132311122313-0201333030230221-2110023012110012): complete subsection reference.

<a id="canonical-2221013100112033-0000202323023313-2212011211323013-1220210201330301-0121230130303213-3103012112322312-1133301111103132-2013230330200131"></a>

### All schema paths for `xcsh_cluster`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cluster--reference--group-001.md#canonical-1101320303220322-2023312031233122-3200323130133222-2303101322310031-3022013033121332-0303110321220020-2203213332330023-0321121213130213) |
| `auto_http_config` | [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-0011300033130230-3211020310101033-1203201332020302-3222210331331332-3011213322111321-2131221222330221-3032032123023010-0011120001103222) |
| `circuit_breaker` | [circuit_breaker](data-sources--cluster--reference--group-001.md#canonical-3000000221001021-1200322013100111-2333020023322222-0130003021121012-1231131330120213-2232101311030033-1012301130030232-3300211132122203) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](data-sources--cluster--reference--group-001.md#canonical-0211003031110322-3010032131213102-1231011121203030-0113210210321220-1213131300301320-3033300020312111-1231320110133010-1002122200300211) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](data-sources--cluster--reference--group-001.md#canonical-2120331032300321-1000111113003331-3230000132301001-0111330310231312-3100333123123023-0000210122110201-3112132122032223-2030300103000332) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](data-sources--cluster--reference--group-001.md#canonical-2010122222330121-2322003103122100-2332312200130310-0132112120000233-1021312003212111-3011001013300022-2113111012110002-3023113022110203) |
| `circuit_breaker.priority` | [circuit_breaker.priority](data-sources--cluster--reference--group-001.md#canonical-3113221123221320-2301003320212332-1323012130000023-0210130302332200-1230322212013303-3300220112021220-2020122220302311-1031202200211222) |
| `circuit_breaker.retries` | [circuit_breaker.retries](data-sources--cluster--reference--group-001.md#canonical-3222002230302101-2201112323320230-3011212130131332-3212233320212002-0032013301220220-2333132311333210-1211100311022321-1312032000121101) |
| `connection_timeout` | [connection_timeout](data-sources--cluster--reference--group-001.md#canonical-2000321022302002-0201003312220211-3211032212031002-3012101311231213-2201213033000311-3300322330121112-2002223310232212-2123201121201022) |
| `default_subset` | [default_subset](data-sources--cluster--reference--group-001.md#canonical-1123020331220310-3100013210210133-2322213222203002-0003323303002113-2312203322022112-0220310303011030-3302003033013203-1223201313332100) |
| `description` | [description](data-sources--cluster--reference--group-001.md#canonical-0220332033330312-0303130021321002-0233300231002323-0102012331232032-0012003322230200-0211110023213101-1300232021201331-3200321313112333) |
| `disable_proxy_protocol` | [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-3311131210223112-2233003101130310-1221311231320320-3201222102230102-1001302211320020-0032211013100022-0030003033113021-1000123120031012) |
| `endpoint_selection` | [endpoint_selection](data-sources--cluster--reference--group-001.md#canonical-0031300321021321-3030223300321312-2111112120001232-2222200230021102-1121032000123231-0132132132113133-0021003213122203-3313103323321123) |
| `endpoint_subsets` | [endpoint_subsets](data-sources--cluster--reference--group-001.md#canonical-2023200031112132-2013130022232130-2200002300002011-0113120210023302-2200323132200132-0202203022130311-1021212310001201-2103233333221132) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](data-sources--cluster--reference--group-001.md#canonical-2322112113300112-3103113330323103-0130210211300003-1211223111002300-3331031100120232-3212121223131212-2232102203321220-2001010002030300) |
| `endpoints` | [endpoints](data-sources--cluster--reference--group-001.md#canonical-3002121301310101-2300223030020122-1210012221001133-3000123010331023-2313201003323111-1301132120110022-2202233220313020-3112220202012111) |
| `endpoints.kind` | [endpoints.kind](data-sources--cluster--reference--group-001.md#canonical-1212212133302132-1301032121331033-3023102023012230-2212311232130123-2013201122000110-2303021100201002-1301121310300300-1310010321100230) |
| `endpoints.name` | [endpoints.name](data-sources--cluster--reference--group-001.md#canonical-2202120123011230-0323120210310013-3211112023313232-0211303001003102-3232303022032223-2032333312120201-3001031123321301-2231032020213031) |
| `endpoints.namespace` | [endpoints.namespace](data-sources--cluster--reference--group-001.md#canonical-1211231100330032-0123201211101303-2230323100232032-1030311231211233-3132200113133331-3202133310330003-2312200223300113-0023221223021210) |
| `endpoints.tenant` | [endpoints.tenant](data-sources--cluster--reference--group-001.md#canonical-0311001022223303-2112221012023103-1333233200122112-0201212312212231-1000100311020022-2312103121211311-2122200111320312-1013013022310212) |
| `endpoints.uid` | [endpoints.uid](data-sources--cluster--reference--group-001.md#canonical-1020112321303313-3110202331103012-0112203232011112-0313221023113002-0312221222210323-3331233101000213-3120101131312212-2023023202110212) |
| `fallback_policy` | [fallback_policy](data-sources--cluster--reference--group-001.md#canonical-0202030220003233-1131100303233102-1111232010300230-0133131013222022-0213322330210111-2222203002123333-2022113133111212-0310332021232322) |
| `health_checks` | [health_checks](data-sources--cluster--reference--group-001.md#canonical-2330230001023101-3032010001303123-2323222313013203-2110120110011031-1132123120230111-2011232122302331-3211121310223230-0213021001221321) |
| `health_checks.kind` | [health_checks.kind](data-sources--cluster--reference--group-001.md#canonical-2011001131213031-1211030233210121-3000311130131321-1101200330233223-0102020002012300-2012031122122211-1202003102110210-1103012101302101) |
| `health_checks.name` | [health_checks.name](data-sources--cluster--reference--group-001.md#canonical-1111031210231101-2132100333310113-1111230023111221-2102310132210232-2203230110220011-2312221231130103-0212231333333122-2000021011132303) |
| `health_checks.namespace` | [health_checks.namespace](data-sources--cluster--reference--group-001.md#canonical-2121111210111110-0133001310323130-0331201233332132-3010010121102201-2210320312203033-1120133233322031-1220232211220132-1111121023100020) |
| `health_checks.tenant` | [health_checks.tenant](data-sources--cluster--reference--group-001.md#canonical-1002202222233331-1033011121000100-1321121033300111-0113030022103122-0020232210213301-2020000310323213-3210322121220020-2013012311232121) |
| `health_checks.uid` | [health_checks.uid](data-sources--cluster--reference--group-001.md#canonical-0302213230011323-1010303032131311-0100221330300333-3022133332000313-0120101220302003-0113022202112122-0332303012303230-0123020122303323) |
| `http1_config` | [http1_config](data-sources--cluster--reference--group-001.md#canonical-2301002310212100-0332002121113223-1023031022121012-2221100333133112-3110310301030103-0131303332113211-3012211333031113-2332122033202022) |
| `http1_config.header_transformation` | [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-2120223221022133-2303010313302310-2231123121002230-1003020321201303-3230212231230311-3000033002132203-1131030210120332-3000212310111303) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](data-sources--cluster--reference--group-001.md#canonical-2031312121130112-1031211131122113-0000311113201112-2231022231221023-3000230010322121-0221123332221231-1012011223210002-1031032000102122) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-3133323302333121-1311113023203332-1223213230011011-2023033102201101-0231101322011022-3031230210320301-0232232232002220-0132022220230032) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-3223331123100022-2232230130020000-1222320300200311-2120020222322033-2330020122132323-3010010201100030-1233302333032233-0121312230313023) |
| `http2_options` | [http2_options](data-sources--cluster--reference--group-001.md#canonical-0100112001100121-1233132000212123-2033013333023121-1231001231032122-3230002221023121-2323001231203302-2030312312102112-3123230301303121) |
| `http2_options.enabled` | [http2_options.enabled](data-sources--cluster--reference--group-001.md#canonical-0103010302202010-0331023112100222-1213221330222122-1320103330132033-2201020031023011-3012331320030101-1210231220010332-0323203312303312) |
| `http_idle_timeout` | [http_idle_timeout](data-sources--cluster--reference--group-001.md#canonical-1332121023203333-3131100012112321-0011311020321003-2210012331303300-0110300000212102-0300100022221310-3132222102100112-0032212200101021) |
| `id` | [ID](data-sources--cluster--reference--group-001.md#canonical-0332221211211103-2301212301233032-0130113113122310-2322233312002301-3033023311310312-2013113023120113-0323313330210102-2331220110311200) |
| `labels` | [labels](data-sources--cluster--reference--group-001.md#canonical-0030203133223130-2132103030032203-0031233032121210-3320311320023211-3311003011110113-1013202320123321-2123132322011022-2311110232210031) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](data-sources--cluster--reference--group-001.md#canonical-3112211023121121-3033100230002232-2330230332202013-1010210123301001-0213300122020121-0033203010110312-1112312120012010-1101110131311020) |
| `max_requests_per_connection` | [max_requests_per_connection](data-sources--cluster--reference--group-001.md#canonical-2110110221221120-0222132101222223-1233310103210203-0311032123200010-0003101011233201-3131211122322023-1312120032023221-3133013022001000) |
| `name` | [name](data-sources--cluster--reference--group-001.md#canonical-2331133322311033-2110131310221313-1023010132322331-1122201012303301-0111020223200330-0202000322230122-3300300003213001-0310222112000020) |
| `namespace` | [namespace](data-sources--cluster--reference--group-001.md#canonical-0000110003301331-3233210230312112-3012110310312001-3320113032203221-3213103331001223-1102031332202103-0221330311133213-1023211003231130) |
| `no_panic_threshold` | [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-1133110030111331-2320333301100321-0212313020013101-0100003211011213-3202013300130031-1233023332010311-3313230211103002-3200332133223132) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-3020200133211313-0212102330033311-3020210112303000-1112000320311103-1000001100200110-1323303221203003-2313202112212110-2303232031200223) |
| `outlier_detection` | [outlier_detection](data-sources--cluster--reference--group-001.md#canonical-3032230020112312-2123120132132002-3012303023003102-0320220102020023-3033023322013200-1030202321302111-0223301100311313-3113201102302111) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](data-sources--cluster--reference--group-001.md#canonical-2131330311113223-1002301310321200-3101002000230022-0303220021130311-3213032023130033-0331332002032101-3022230022120231-1213020330023132) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](data-sources--cluster--reference--group-001.md#canonical-3301032232322231-3220320310201003-1102132221220230-2111310101101320-3303131200021031-2122123213320211-1133210203303232-0210331122233131) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](data-sources--cluster--reference--group-001.md#canonical-0321130113023121-2300013100331330-3210012033013231-3020112201002011-3011231201201102-1232031101113031-2302112313001213-1310030221330100) |
| `outlier_detection.interval` | [outlier_detection.interval](data-sources--cluster--reference--group-001.md#canonical-2101102111310201-1311200303120110-0022031003013310-2231333322030223-3031111100303033-3230023031302313-0220331202032232-1312312032203213) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](data-sources--cluster--reference--group-001.md#canonical-2013222132001220-1022300302333123-0212210231211030-0320322211112130-3300030211201122-2303103221110103-0220312301112011-0312220323232021) |
| `panic_threshold` | [panic_threshold](data-sources--cluster--reference--group-001.md#canonical-1331302130111330-1123231310213032-0322001313312222-1330000303121111-1210331121110203-3333213312222301-3203233113121223-0331223123123211) |
| `proxy_protocol_v1` | [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-1131001111120303-3121233223210101-0231213113022120-0203122021333003-0121032032312033-0201131132021130-0300233011011120-3003301301110100) |
| `proxy_protocol_v2` | [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-1023310220311012-2330002031220301-0131230200210222-1021013003230310-0302002332312033-3311313210120332-2023103223220010-1333220233203211) |
| `tls_parameters` | [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3030010133233120-1232031022131133-0002233333213112-2221011212000120-3230123303222010-3110112132302301-3021233133233202-0030011023220320) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-3020233210311233-1213232021030321-3331213333103221-1123230232220021-2311113110233312-1111013033113210-3221202133200102-1220110020120323) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](data-sources--cluster--reference--group-002.md#canonical-3022330102202210-3123330220312113-0312132232002130-1111231323023032-1232011120320001-3203220210233132-1311120232000002-2322322001002210) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](data-sources--cluster--reference--group-002.md#canonical-1222332013333121-2013213300323201-3210202202232113-0020001322233133-3122312021202221-0321331012012121-2000200203331112-1003030013203130) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](data-sources--cluster--reference--group-002.md#canonical-2121132212133111-3300221213100133-3313320231300011-2001030012221111-3223333311123321-0002001101012333-1113301321100002-0101033220033121) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](data-sources--cluster--reference--group-002.md#canonical-2223313113220012-2131123100102022-0131321023130333-2020102211013302-0210000233201101-3020213302210212-0111003330312130-1003033312323332) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](data-sources--cluster--reference--group-002.md#canonical-0102331323120132-3001112132113332-1211230012011312-3120233121113221-2031020132220023-2123332221023220-2310221113301203-1120330201123310) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](data-sources--cluster--reference--group-002.md#canonical-1103111223313121-2020323202333033-2122012003311331-0233130213110020-3000011021332102-3031012233103103-2312111222303223-2133000222233200) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](data-sources--cluster--reference--group-001.md#canonical-0323130330101232-0230301011312010-0133003123100202-1220110023101222-3000312032002302-2030113023211212-0112001003230331-2303300020133300) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](data-sources--cluster--reference--group-002.md#canonical-2010133011013011-2321002333310321-2022113303113202-3300320203000233-0022230011331101-1000101201031201-0130202123221321-1203301012101331) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](data-sources--cluster--reference--group-002.md#canonical-1202221312032002-1021120001313223-0232231313211033-3222022132031310-0001221231221101-0110131030310012-1231233013023032-3100032120103000) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](data-sources--cluster--reference--group-002.md#canonical-2301303311111120-3032310021011233-2003021203003001-3111220333323103-3331102010030210-1303331012011200-1320023131012332-2122220003321302) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-002.md#canonical-1011101011210320-2100123233010303-2121233123002021-1101123103313012-1101212103212131-3231311102101010-1233301013312121-1302013320113100) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](data-sources--cluster--reference--group-002.md#canonical-1313312222133121-0100233022131132-3113213130111102-0101212020011311-0231001120201130-1103022132231011-2021001211031332-3013021200011022) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-002.md#canonical-0212023021313002-3212032010001011-3211112322220121-1132302212011220-0331022313012321-3320113321121202-1333320121011203-3020122231300122) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-002.md#canonical-0223323032202213-2313010202310121-0231312000012330-0001221022123103-3233103002332310-3032312121001113-2232312000333212-1033232233010232) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--reference--group-002.md#canonical-0023212330332120-2000300333233321-2220113133103233-3230003110110102-2131212020211011-3123013131103223-1103110331202122-0131023320232123) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--reference--group-002.md#canonical-3320100030132230-2210322320001312-1321301333103323-0323333023032021-1320123302013313-3110001202113212-3202311300101221-2210311113202330) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--reference--group-002.md#canonical-3121210121330303-1332330222021232-3203310320232033-3232101220331302-0002320213010213-0322303011112102-1122232131120210-1030010033013032) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--reference--group-002.md#canonical-1011233230030113-1200220311101223-0032122222132322-0032303321322111-0313123210021233-2212032312212023-0201000001021032-3032323333023032) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--reference--group-002.md#canonical-3230002123103210-2111302233121132-2132000030103011-2013223212002120-3211201023100321-1210333110030102-0002310103110021-2213211232103000) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](data-sources--cluster--reference--group-002.md#canonical-1103100231322000-1131303332003023-2312202300102333-3333222330311022-3131113100130312-1322020033133212-2322311310003020-2311213012212332) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--cluster--reference--group-002.md#canonical-2323113302302121-3201200330103312-1003132131303332-0221210221313110-2203002300223012-0113003002203330-2201011121330013-2330130320020110) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](data-sources--cluster--reference--group-002.md#canonical-3133123202302013-2331133120202210-3230012023023100-0131003221300200-1102310033223221-1103111021210120-1302333123331121-3330320113202330) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-3123321333013122-3232130001010230-3132300121003211-3311133311011300-3030121031211001-2123330211121210-0110211021032022-0301300233023302) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--cluster--reference--group-002.md#canonical-0313133310232211-0223111322203322-1320220230012222-1313210210022130-3102311323321320-2123030033221220-0202023332302323-1302023331311021) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--cluster--reference--group-002.md#canonical-0211131202031123-1022120120303223-0302001332103201-0112330002030233-2122101302121030-3133011303222321-2132001302232310-0101021330031133) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--cluster--reference--group-002.md#canonical-2023303003031310-3132233101231002-0330332023201012-3001311100331221-0131322303212333-0302003032031230-2112300311313030-0020331023211220) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-1310312112200001-0111003312110310-1221220032000223-2333122121202220-1203203223012011-2001320231031300-2231012200331331-2313021222330210) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--cluster--reference--group-002.md#canonical-2121031032221002-0001010332023210-0321021211233032-3010201330102002-3122033311001131-0320231312030111-2022013333223232-0030332331222120) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--cluster--reference--group-002.md#canonical-1022021001021210-1102331000011200-1200021100011011-1102323022131112-1002301221112013-1133301311010301-0002212332201330-0332303233211221) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--cluster--reference--group-002.md#canonical-0100111200330331-2312131231313103-1310302203321222-2003210301302323-3131323222310322-1013101302122133-1000312102330230-2000301031030321) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--cluster--reference--group-002.md#canonical-2022200303201210-2233203331313022-1023133203113332-0332123313301210-0333220313210100-0010330212001211-3202331220332313-3302133301123013) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--cluster--reference--group-002.md#canonical-0230131330101331-1213223131333112-3133012001313320-1013210110100232-3022231301003022-3312023202020021-3023213001222123-0223221010111103) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-002.md#canonical-1220301332013011-3020203200102211-3312112101033032-3103030112303233-0311101121210230-1211123102010021-2303033320012333-1102010033231300) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cluster--reference--group-002.md#canonical-0303003102332031-3102312333322202-1322121222033023-2230122310203002-2210103101301032-0201332313223121-3301030032100001-1011202311332000) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--cluster--reference--group-002.md#canonical-0300321200000013-0321130202121000-3030103031000123-0000003031131302-2103132113230330-0220012300121113-1300202323112200-0111130111311100) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--cluster--reference--group-002.md#canonical-1110111221123001-1213120331130003-1302023330123001-3232302110132230-2232303320003201-1010210021311012-3033220122032233-3212111231020130) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--cluster--reference--group-002.md#canonical-3222103111231110-2001133301203021-1320022022111032-3121230302302130-3210232113112200-2230202000113013-2002201131221020-3112231012301201) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--cluster--reference--group-002.md#canonical-3230133002223220-3322210003133223-1001231001012213-3232233120213013-1002022110121103-3302223001333103-3223333213123103-2231021121202102) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--cluster--reference--group-002.md#canonical-2321233323321113-0133220021200321-0021123133303033-1123230220310130-2232023323212321-3031023001132112-3303221203302203-3111032101231213) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--cluster--reference--group-002.md#canonical-2313231013332230-1101332021330321-0330213313031132-3230302302022231-2203122221101022-2203233031222023-3121201212102311-3213310131320120) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--cluster--reference--group-002.md#canonical-0222220232131013-1232322220032302-3231213130201123-3102220102201223-3131022302133303-1111131220021330-1211010321211221-3130303132223220) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-002.md#canonical-0322110113201121-1223001312302130-3221220020333311-2102121120311333-1201312300212232-1001021232023032-2212201223012011-0202313332100230) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--cluster--reference--group-002.md#canonical-2013330002022110-3221233320111022-2113200032223021-2131100333310012-2022101303332013-2112121220032300-1011101133030330-3012122112003203) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-002.md#canonical-1001022012100032-2110022203120003-3001223300010213-1210330122130203-2311223123122123-1012010020130323-3333000331332001-0131012011101101) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-002.md#canonical-3112213033112310-3321000031223300-0223001133022332-1233122302331310-0302101111221020-1031000022220032-2000002220333113-3121322311102031) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--reference--group-002.md#canonical-1202320330213302-1023021011101132-1031201131232203-0200212202022221-2303100011202131-1313012103132123-0313333100213330-0233123120110002) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--reference--group-002.md#canonical-3030211113230023-3231233232030313-0010220100023131-1200102300012110-1222022112312002-1320323323101021-2211211212332212-3003320032330332) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--reference--group-002.md#canonical-3213131101002101-0031213131020203-0202222210230021-0023223100103303-2030331232022202-1103002203120120-2321000011213130-0331330013132003) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--reference--group-002.md#canonical-0301131130020121-1102321302221000-2213122213100100-0110031033331033-1112123312020313-3322220012302102-0231132301303121-0011101311131123) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--reference--group-002.md#canonical-2323321121011121-1323102303012131-2300113320111023-0222203320220011-1232120212213323-1013320011330013-0301031012032103-3323322032211210) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--cluster--reference--group-002.md#canonical-2203102031001022-1232331020313232-2201021032232100-1323211222313033-3110112333211222-0132100313332010-0013133333323313-0223133231332311) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--cluster--reference--group-002.md#canonical-0113323122033110-3033312012301320-3322210232211311-0302120230123012-2100011100011203-3201201302002033-3231030021101332-3112220022000330) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](data-sources--cluster--reference--group-002.md#canonical-0020132322020000-2220201301121210-3130113132111221-2123311202220002-1010313220223102-0220312132310131-1321201312200331-3130110101312233) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](data-sources--cluster--reference--group-002.md#canonical-0013102131022122-0322331031333310-1133301322200323-3130000112013123-1313212320200100-1023012212332300-1100212021011323-0010230221220121) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](data-sources--cluster--reference--group-002.md#canonical-2223002332030032-1012200132302111-3202003203333003-0230030132030210-0221333123123002-2313312130233001-1121312003002202-2132113112132320) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](data-sources--cluster--reference--group-001.md#canonical-3000320110311321-1133023233200201-2223311201323001-0133122100232332-0321003103110232-0002222003103001-3211030301002001-0112230312030122) |
| `tls_parameters.sni` | [tls_parameters.sni](data-sources--cluster--reference--group-001.md#canonical-1103002011211222-3002200211231131-2212102310313333-0313023031302111-1331230201002301-1321132201100230-2030322311021033-2230033020100203) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](data-sources--cluster--reference--group-002.md#canonical-2013110333111220-1222102320022320-3221310000202032-1012033010233330-1133110033133110-3122322311111221-0311121001312330-2100213320030111) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-0022220233122111-0100101301121003-2023000303213331-1302120002102111-1012311112211200-0101133111202031-1133222103113333-1221211022000220) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-0211233012013322-3003020020021130-3312103330213112-3233101213002130-2131003200102012-3211331012023103-2112300133111102-0211203032320210) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-1111133231100310-0321301220320333-2112112032010331-3320133000332222-3020122220111331-3301322223110012-1130221203102232-2223220032131301) |

<a id="canonical-0121220033001011-3230302333131331-1120220130301100-2331202102323301-3103232211000023-1201210120201312-0202301131223322-2021003222112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `auto_http_config` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- auto_http_config

<a id="canonical-0011300033130230-3211020310101033-1203201332020302-3222210331331332-3011213322111321-2131221222330221-3032032123023010-0011120001103222"></a>

Type: `["object", {}]`. Computed.

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

- [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-0011300033130230-3211020310101033-1203201332020302-3222210331331332-3011213322111321-2131221222330221-3032032123023010-0011120001103222)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2301002310212100-0332002121113223-1023031022121012-2221100333133112-3110310301030103-0131303332113211-3012211333031113-2332122033202022)
- [http2_options](data-sources--cluster--reference--group-001.md#canonical-0100112001100121-1233132000212123-2033013333023121-1231001231032122-3230002221023121-2323001231203302-2030312312102112-3123230301303121)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030311322100312-0021313011330010-0011030103232220-1333113213133000-0223021121231313-2213032200211113-1231102101131311-3333002013010212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `circuit_breaker` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- circuit_breaker

<a id="canonical-3000000221001021-1200322013100111-2333020023322222-0130003021121012-1231131330120213-2232101311030033-1012301130030232-3300211132122203"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1201001200220323-3130011001010210-2111011020331322-0033012022022301-3232120120111312-0313012031122201-1001002201331322-2302310020011120"></a>

### Direct properties for `circuit_breaker`

<a id="canonical-0211003031110322-3010032131213102-1231011121203030-0113210210321220-1213131300301320-3033300020312111-1231320110133010-1002122200300211"></a>

#### `circuit_breaker.connection_limit` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-2120331032300321-1000111113003331-3230000132301001-0111330310231312-3100333123123023-0000210122110201-3112132122032223-2030300103000332"></a>

<a id="canonical-1223033012021220-2221220223323100-1320212302031000-2120312112122011-0322030220010332-3032220212020101-0000013103123320-3213200131321332"></a>

#### `circuit_breaker.max_requests` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-2010122222330121-2322003103122100-2332312200130310-0132112120000233-1021312003212111-3011001013300022-2113111012110002-3023113022110203"></a>

<a id="canonical-2100222110322101-0113203321100100-2010111103221222-3301121103110211-3331110001010011-1230220020013110-0021200022033223-2001031323022311"></a>

#### `circuit_breaker.pending_requests` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-3113221123221320-2301003320212332-1323012130000023-0210130302332200-1230322212013303-3300220112021220-2020122220302311-1031202200211222"></a>

<a id="canonical-1100131213212102-3020111321000212-3012223330131230-3212023312230331-0003300103103123-0120013202123111-3213101111232300-1302212301132311"></a>

#### `circuit_breaker.priority` property

Type: `"string"`. Computed.

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

<a id="canonical-3222002230302101-2201112323320230-3011212130131332-3212233320212002-0032013301220220-2333132311333210-1211100311022321-1312032000121101"></a>

<a id="canonical-2023013230230022-0233222332223002-2013230230121102-1233102102313232-1120001300311320-1112203222202022-2033031303300211-1323200210223002"></a>

#### `circuit_breaker.retries` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3211100033212002-3323310310011303-3231020001323310-3133321001312311-3031123102313023-3321012012100132-3231002312200331-1303113213312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_subset` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- default_subset

<a id="canonical-1123020331220310-3100013210210133-2322213222203002-0003323303002113-2312203322022112-0220310303011030-3302003033013203-1223201313332100"></a>

Type: `"single"`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032021323222023-2113331010130133-1220333211120201-1312220223131131-1322033303023321-2223300130100302-1001220312222332-1013112312332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_proxy_protocol` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- disable_proxy_protocol

<a id="canonical-3311131210223112-2233003101130310-1221311231320320-3201222102230102-1001302211320020-0032211013100022-0030003033113021-1000123120031012"></a>

Type: `["object", {}]`. Computed.

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

- [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-3311131210223112-2233003101130310-1221311231320320-3201222102230102-1001302211320020-0032211013100022-0030003033113021-1000123120031012)
- [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-1131001111120303-3121233223210101-0231213113022120-0203122021333003-0121032032312033-0201131132021130-0300233011011120-3003301301110100)
- [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-1023310220311012-2330002031220301-0131230200210222-1021013003230310-0302002332312033-3311313210120332-2023103223220010-1333220233203211)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020130102112320-3222131012320000-2113321330033001-3012022222130010-1230112302230233-2303032003222232-1323001031233101-2003001000131130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_subsets` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- endpoint_subsets

<a id="canonical-2023200031112132-2013130022232130-2200002300002011-0113120210023302-2200323132200132-0202203022130311-1021212310001201-2103233333221132"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0212211222203021-0000302120303122-2320110130211003-2022313022202023-2220101023321133-2030023300001023-1130122133331302-2012323120311112"></a>

### Direct properties for `endpoint_subsets`

<a id="canonical-2322112113300112-3103113330323103-0130210211300003-1211223111002300-3331031100120232-3212121223131212-2232102203321220-2001010002030300"></a>

#### `endpoint_subsets.keys` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0112011310121212-0110003113310222-2323213222023033-0321211112303230-2213202000122321-0101220230120010-3221102303012330-2320002330220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoints` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- endpoints

<a id="canonical-3002121301310101-2300223030020122-1210012221001133-3000123010331023-2313201003323111-1301132120110022-2202233220313020-3112220202012111"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022301000121220-2201130301200132-3230002132100120-3101233010111012-1013001303320333-2133212333012033-3330223132231010-1302002103033231"></a>

### Direct properties for `endpoints`

<a id="canonical-1212212133302132-1301032121331033-3023102023012230-2212311232130123-2013201122000110-2303021100201002-1301121310300300-1310010321100230"></a>

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

<a id="canonical-2202120123011230-0323120210310013-3211112023313232-0211303001003102-3232303022032223-2032333312120201-3001031123321301-2231032020213031"></a>

<a id="canonical-3210121033303221-0131101113113001-3303212103201002-2023230320001002-2202112001032101-3323313103011200-2021113021200013-1100032132302212"></a>

#### `endpoints.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1211231100330032-0123201211101303-2230323100232032-1030311231211233-3132200113133331-3202133310330003-2312200223300113-0023221223021210"></a>

<a id="canonical-2312330121023103-0330122123212131-1312220233023102-0313222123300102-0332312230120202-2030112000120303-1033111303022221-2032110003310320"></a>

#### `endpoints.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-0311001022223303-2112221012023103-1333233200122112-0201212312212231-1000100311020022-2312103121211311-2122200111320312-1013013022310212"></a>

<a id="canonical-2201113203222302-3221201302032222-1033001010313020-3030031102010001-0310111010001223-1121213002220103-3330213223000011-1211020120123321"></a>

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

<a id="canonical-1020112321303313-3110202331103012-0112203232011112-0313221023113002-0312221222210323-3331233101000213-3120101131312212-2023023202110212"></a>

<a id="canonical-3201233231331113-0031131030232023-3020223102230122-2120133031230112-0122133113113112-3132333232230231-2013133001203002-3131300030230002"></a>

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

<a id="canonical-3202012113122110-1021232210312333-3113020021130130-2321033332112012-3321322013313120-0300123320333310-2313332022323122-1222103012303003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `health_checks` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- health_checks

<a id="canonical-2330230001023101-3032010001303123-2323222313013203-2110120110011031-1132123120230111-2011232122302331-3211121310223230-0213021001221321"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1031020301120022-0302000312012321-3231001100332032-0121200022121300-1211201230310012-3221220030020301-0321320321102320-1302131131002300"></a>

### Direct properties for `health_checks`

<a id="canonical-2011001131213031-1211030233210121-3000311130131321-1101200330233223-0102020002012300-2012031122122211-1202003102110210-1103012101302101"></a>

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

<a id="canonical-1111031210231101-2132100333310113-1111230023111221-2102310132210232-2203230110220011-2312221231130103-0212231333333122-2000021011132303"></a>

<a id="canonical-2023111020021121-1233232130112202-0010223122023332-2002111113023111-3203122011222133-3212101120012000-0011020121220223-1030211113123103"></a>

#### `health_checks.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2121111210111110-0133001310323130-0331201233332132-3010010121102201-2210320312203033-1120133233322031-1220232211220132-1111121023100020"></a>

<a id="canonical-2110222330110001-2010323303121303-2030212110303120-2100130110031320-0320002221122100-0223203233102331-3100023021113223-2220102130031323"></a>

#### `health_checks.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-1002202222233331-1033011121000100-1321121033300111-0113030022103122-0020232210213301-2020000310323213-3210322121220020-2013012311232121"></a>

<a id="canonical-0132113222010003-2120320030231313-1232313301111311-1303203333031223-1022202012310001-0313202110221003-0001001230220212-3111310013322013"></a>

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

<a id="canonical-0302213230011323-1010303032131311-0100221330300333-3022133332000313-0120101220302003-0113022202112122-0332303012303230-0123020122303323"></a>

<a id="canonical-3101102211013102-3001202001233100-2320322230113102-2002311103323132-0311231011312013-2112000000001213-0331333301231310-2000103202020232"></a>

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

<a id="canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- http1_config

<a id="canonical-2301002310212100-0332002121113223-1023031022121012-2221100333133112-3110310301030103-0131303332113211-3012211333031113-2332122033202022"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0100011022123011-0200233210320033-0310300000201330-0303121011100330-3333322201112003-1200221300112220-2301100102303303-1203320102201223"></a>

### Direct properties for `http1_config`

- [header_transformation](data-sources--cluster--reference--group-001.md#canonical-3103220002212220-2001001020021333-2102231310103133-3113331311111230-2233331030303323-3220232131032231-0100202330320033-1323222313111111): complete subsection reference.

<a id="canonical-3103220002212220-2001001020021333-2102231310103133-3113331311111230-2233331030303323-3220232131032231-0100202330320033-1323222313111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121)
- http1_config.header_transformation

<a id="canonical-2120223221022133-2303010313302310-2231123121002230-1003020321201303-3230212231230311-3000033002132203-1131030210120332-3000212310111303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2302231300131131-0212102231330112-0111330120121210-3321232013200130-3330321213033230-0330233320103132-3130033032312330-2332321233112201"></a>

### Direct properties for `http1_config.header_transformation`

- [default_header_transformation](data-sources--cluster--reference--group-001.md#canonical-1313010310302303-2230233210303311-3011101232232302-2023231131031120-0311122323121222-3100322120113000-2331132302133002-3002131100123103): complete subsection reference.

- [preserve_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-2220332303313031-2300211031211212-0100222113233010-3103023033001331-1003023010131121-2211302313102021-2203321122132233-0100101113232013): complete subsection reference.

- [proper_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-1321200301101333-2303303311023112-0032322022321203-3203023031100131-3323113321103020-1323033023121220-0313232223222133-0020120233332202): complete subsection reference.

<a id="canonical-1313010310302303-2230233210303311-3011101232232302-2023231131031120-0311122323121222-3100322120113000-2331132302133002-3002131100123103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-3103220002212220-2001001020021333-2102231310103133-3113331311111230-2233331030303323-3220232131032231-0100202330320033-1323222313111111)
- http1_config.header_transformation.default_header_transformation

<a id="canonical-2031312121130112-1031211131122113-0000311113201112-2231022231221023-3000230010322121-0221123332221231-1012011223210002-1031032000102122"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220332303313031-2300211031211212-0100222113233010-3103023033001331-1003023010131121-2211302313102021-2203321122132233-0100101113232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-3103220002212220-2001001020021333-2102231310103133-3113331311111230-2233331030303323-3220232131032231-0100202330320033-1323222313111111)
- http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-3133323302333121-1311113023203332-1223213230011011-2023033102201101-0231101322011022-3031230210320301-0232232232002220-0132022220230032"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321200301101333-2303303311023112-0032322022321203-3203023031100131-3323113321103020-1323033023121220-0313232223222133-0020120233332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http1_config.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-0231221033120032-3310232112322110-2213311020222113-3231330231233122-0300031022330003-2221030311303103-1122002330303331-2301311312200121)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-3103220002212220-2001001020021333-2102231310103133-3113331311111230-2233331030303323-3220232131032231-0100202330320033-1323222313111111)
- http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-3223331123100022-2232230130020000-1222320300200311-2120020222322033-2330020122132323-3010010201100030-1233302333032233-0121312230313023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030010231200022-3322131033122210-1230033020312003-0111313202211202-3322302012332230-1211221323031332-1123230123011022-3232203321032103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http2_options` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- http2_options

<a id="canonical-0100112001100121-1233132000212123-2033013333023121-1231001231032122-3230002221023121-2323001231203302-2030312312102112-3123230301303121"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0002301320231220-1122100223103322-1232001001203213-2112331133033302-0302131310202002-1122322303133100-0023320203133313-0220022213232310"></a>

### Direct properties for `http2_options`

<a id="canonical-0103010302202010-0331023112100222-1213221330222122-1320103330132033-2201020031023011-3012331320030101-1210231220010332-0323203312303312"></a>

#### `http2_options.enabled` property

Type: `"bool"`. Computed.

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

<a id="canonical-2211301230101320-2002021313220200-3310210323032101-1213102313203211-1100223020000320-0300213121222230-0133300202032003-3232212303130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_panic_threshold` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- no_panic_threshold

<a id="canonical-1133110030111331-2320333301100321-0212313020013101-0100003211011213-3202013300130031-1233023332010311-3313230211103002-3200332133223132"></a>

Type: `["object", {}]`. Computed.

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

- [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-1133110030111331-2320333301100321-0212313020013101-0100003211011213-3202013300130031-1233023332010311-3313230211103002-3200332133223132)
- [panic_threshold](data-sources--cluster--reference--group-001.md#canonical-1331302130111330-1123231310213032-0322001313312222-1330000303121111-1210331121110203-3333213312222301-3203233113121223-0331223123123211)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203320230000303-2221322103212223-2212111131011213-2232212212033022-3230300202221320-0033222323113312-2003021313202221-1223013113002011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- no_request_limit_per_connection

<a id="canonical-3020200133211313-0212102330033311-3020210112303000-1112000320311103-1000001100200110-1323303221203003-2313202112212110-2303232031200223"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221013233311223-2121003102312233-2021333310030322-3330222031201202-0121032031312212-3000221001110310-2311002311301120-3232021321201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `outlier_detection` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- outlier_detection

<a id="canonical-3032230020112312-2123120132132002-3012303023003102-0320220102020023-3033023322013200-1030202321302111-0223301100311313-3113201102302111"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0200330003222101-3202012130213121-2121022210203032-0100301301133012-0210320131122201-1013132210122033-3300033201203100-1230011101201300"></a>

### Direct properties for `outlier_detection`

<a id="canonical-2131330311113223-1002301310321200-3101002000230022-0303220021130311-3213032023130033-0331332002032101-3022230022120231-1213020330023132"></a>

#### `outlier_detection.base_ejection_time` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-3301032232322231-3220320310201003-1102132221220230-2111310101101320-3303131200021031-2122123213320211-1133210203303232-0210331122233131"></a>

<a id="canonical-0221223210010300-1111022320032130-2023210222020200-0300221321133033-3313213331000223-2330120221233330-0302100013021012-2213010303212311"></a>

#### `outlier_detection.consecutive_5xx` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-0321130113023121-2300013100331330-3210012033013231-3020112201002011-3011231201201102-1232031101113031-2302112313001213-1310030221330100"></a>

<a id="canonical-3001330002323301-1133033313130213-3230030101031320-1101321133223300-2302122132202211-2123320021232123-0013322100010332-2222001213023021"></a>

#### `outlier_detection.consecutive_gateway_failure` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-2101102111310201-1311200303120110-0022031003013310-2231333322030223-3031111100303033-3230023031302313-0220331202032232-1312312032203213"></a>

<a id="canonical-1000122001320330-3021033011200100-3010121212322133-2223013022001002-3212010332333203-3210330030131222-2102002031311303-0232331233103313"></a>

#### `outlier_detection.interval` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2013222132001220-1022300302333123-0212210231211030-0320322211112130-3300030211201122-2303103221110103-0220312301112011-0312220323232021"></a>

<a id="canonical-2310113110330013-3022203312221033-1010200021310122-1322112122332333-0220000231322011-3001232303033212-2131333232113011-3100032313113123"></a>

#### `outlier_detection.max_ejection_percent` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-1100331322301132-0320031111010231-0332200212032200-2303202012003200-2313311022333003-1302100231200002-3223231030010220-2132022302330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_protocol_v1` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- proxy_protocol_v1

<a id="canonical-1131001111120303-3121233223210101-0231213113022120-0203122021333003-0121032032312033-0201131132021130-0300233011011120-3003301301110100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023012213100031-0232220222310012-3133010122101011-2312120323230312-0310300102000210-2021021123302201-1321110221121130-1220101111021111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_protocol_v2` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- proxy_protocol_v2

<a id="canonical-1023310220311012-2330002031220301-0131230200210222-1021013003230310-0302002332312033-3311313210120332-2023103223220010-1333220233203211"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- tls_parameters

<a id="canonical-3030010133233120-1232031022131133-0002233333213112-2221011212000120-3230123303222010-3110112132302301-3021233133233202-0030011023220320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2123330020103200-2013300133200113-1300022310032002-2121030321031221-2100033030011003-1313022202313030-2202311220320133-2112213230133110"></a>

### Direct properties for `tls_parameters`

- [cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002): complete subsection reference.

- [common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313): complete subsection reference.

- [default_session_key_caching](data-sources--cluster--reference--group-002.md#canonical-0030303201213312-2002233020110002-0103213003202313-0231023111222112-0213131221031201-1302110222212321-3200100310301122-3232332322003330): complete subsection reference.

- [disable_session_key_caching](data-sources--cluster--reference--group-002.md#canonical-3131003313100023-2011121223333200-0211332321102113-3302131121032312-1301111301003330-3020313203010010-1213131322111210-1232103311201020): complete subsection reference.

- [disable_sni](data-sources--cluster--reference--group-002.md#canonical-1130130133200210-1010212021311333-2111122320031101-0333233313010021-0323213330130012-2122211020012332-0212311203233110-0313123311210301): complete subsection reference.

<a id="canonical-3000320110311321-1133023233200201-2223311201323001-0133122100232332-0321003103110232-0002222003103001-3211030301002001-0112230312030122"></a>

<a id="canonical-3202221001311221-2001220011131003-2203223120203023-0202221101320022-2330130233000323-3133002111000213-0230220310101322-1120103221112111"></a>

#### `tls_parameters.max_session_keys` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1103002011211222-3002200211231131-2212102310313333-0313023031302111-1331230201002301-1321132201100230-2030322311021033-2230033020100203"></a>

<a id="canonical-0313333231132121-2023300222222330-0031013332233220-0133212232202110-3011221312001331-2230001321130011-3332212020232301-1001213210002002"></a>

#### `tls_parameters.sni` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](data-sources--cluster--reference--group-002.md#canonical-2223202311322031-0130323103302030-3212303312011321-1130013200232211-2120320001110303-0210223212222320-3120323210123323-3231021123203130): complete subsection reference.

<a id="canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.cert_params

<a id="canonical-3020233210311233-1213232021030321-3331213333103221-1123230232220021-2311113110233312-1111013033113210-3221202133200102-1220110020120323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1130113303121000-0130121211310313-2002112131130122-3003002311321211-0003202230001122-0000031231222003-0021223313221122-0033120013311222"></a>

### Direct properties for `tls_parameters.cert_params`

- [certificates](data-sources--cluster--reference--group-002.md#canonical-3200233330322300-0222311312232113-3213320003030123-1322013221032212-2322301321223233-0222313020122231-3003230032101323-3022130211123332): complete subsection reference.

<a id="canonical-0323130330101232-0230301011312010-0133003123100202-1220110023101222-3000312032002302-2030113023211212-0112001003230331-2303300020133300"></a>
