---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221220012012110-3323110023112323-3330121303222013-0201202031003033-1133200222312132-1313332221221001-1132313313231021-1320322211313121"></a>

## Property reference — Property reference / 120001101131 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- Property reference

<a id="canonical-1230232000210332-2111301013023011-3330121012202303-0200312201103033-3210223323022223-2332120303110313-2300223102012003-0001102032011030"></a>

## Direct properties — Property reference / 120001101131 / 3

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231): complete subsection reference.

<a id="canonical-2100322311221202-1120231030223213-0020212002010332-2230101101000232-1323012111200011-1201322030203021-2310021003013113-1033122011022230"></a>

<a id="canonical-1123333331030233-3030322323301002-0303202003022030-1200022020112110-0012201013323220-1221101003203322-1131232001010131-0100222330011313"></a>

## annotations property — Property reference / 120001101131 / 4

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

- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-1112022203223032-3323221233013132-1121203303031132-1021300130303111-2312333020233311-1132310003313023-1330012320112102-3213232312123011): complete subsection reference.

<a id="canonical-0000200000202123-2113330101101213-1110213113131200-2003022220230302-3020303110123033-0312221222230301-1031221302220210-0031201331222332"></a>

<a id="canonical-1100210233210221-0131232120321231-2221113022203303-1121323131232001-0012130023313201-1022001321332112-3133013301322031-3203203103300122"></a>

## description property — Property reference / 120001101131 / 5

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

<a id="canonical-1101102111030330-3302001311200000-2310001203020230-0220203012033322-0131112223120231-0113132201123023-2122321203120203-0030102022020023"></a>

<a id="canonical-2002212130221231-1210323123200322-2032111001201231-3300202002113011-0033030232231222-0210132221010202-2200112103301123-1300120233031230"></a>

## disable property — Property reference / 120001101131 / 6

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

<a id="canonical-1033200032010330-1131010303302213-0311222023330001-3120232131013322-2201322130131030-3330211222013010-2120232011110310-0312332102311032"></a>

<a id="canonical-1302230233030103-0020223233111010-2033331032032213-2003203123313101-0201201213221133-0130202331332113-3230102230011212-3011130330232301"></a>

## endpoint_selection property — Property reference / 120001101131 / 7

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

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

<a id="canonical-3213010323303321-2133332333032101-3311011211021013-2331301221331003-3031220220201211-2201300222031211-2331112312201313-3221132310000022"></a>

<a id="canonical-3122011133113231-3031121031222023-1030313300232332-2113132113032022-1203020031033310-0120223222230301-3030321233013302-1110101301131021"></a>

## health_check_port property — Property reference / 120001101131 / 8

Type: `"number"`. Optional, Computed.

\[OneOf: health\_check\_port, same\_as\_endpoint\_port\] Exclusive with \[same\_as\_endpoint\_port\]
Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [health_check_port](resources--origin_pool--reference--group-001.md#canonical-3213010323303321-2133332333032101-3311011211021013-2331301221331003-3031220220201211-2201300222031211-2331112312201313-3221132310000022)
- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-2003120102323102-0320322222233123-0110310222321212-1303023300232011-2231113032010222-0322332011312201-0332033220131211-2002332200121120)

Select alternatives according to the provider validators above.

- [healthcheck](resources--origin_pool--reference--group-001.md#canonical-0220303133013032-2302323323212323-0311200001213123-3021130113330311-2212330212003003-3231310310300112-2222012213121021-0030322010120013): complete subsection reference.

<a id="canonical-1103110211131123-1001200020120223-2013131001232202-3220010013313230-3132303322101013-0122331230303030-1113333331110212-2231002010313113"></a>

<a id="canonical-0233111300213323-2012233303202122-1032300331212122-2330001103313102-0333311321032013-2220133201120121-3331303132201021-1313113131211233"></a>

## ID property — Property reference / 120001101131 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0030131313221310-2123001103101032-0211102002001022-0323202013222021-1203332321330120-1220001230132330-1112030312210331-1303321232213120"></a>

<a id="canonical-2301210032203233-1022332000030010-0313023212300110-2332010233303121-1000010030303301-2110233300003331-3220210331012132-0132322222202322"></a>

## labels property — Property reference / 120001101131 / 10

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

- [lb_port](resources--origin_pool--reference--group-001.md#canonical-0020300233212120-1303102312023111-0212101131203331-1302022110112113-0033130113202023-0130301212232010-1022111220323023-0020230101001100): complete subsection reference.

<a id="canonical-1010102003323021-1301112120233030-1233102213223112-2123020123200103-2303203201223321-0232002133112311-2311021320121332-0131221333023101"></a>

<a id="canonical-2212202122030021-1030020202002232-3010010110313333-2101130212123012-3130001133033003-1121312101211212-2033021310321222-3233101101313122"></a>

## loadbalancer_algorithm property — Property reference / 120001101131 / 11

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

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

<a id="canonical-0200313001011322-3023100130023333-1220202120003332-0201032101122020-1022123102303010-0122100120132021-2022110113103121-3233231102100232"></a>

<a id="canonical-2322323010130332-0102220220321013-2100303113233122-0321003223332112-1311030320002212-3013222103011233-0011132301033101-2030331023111120"></a>

## name property — Property reference / 120001101131 / 12

Type: `"string"`. Required.

Name of the Origin Pool. Must be unique within the namespace.

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

<a id="canonical-1001320313030032-1002211022020002-1020312330001121-3333103332110211-1331000133332011-0222321003300300-2323103330010303-1110110102221000"></a>

<a id="canonical-3031302210021311-3122322231023200-3113303010011132-0223212203030112-1032101100030230-2122310232231132-1100123111132233-1012232200013122"></a>

## namespace property — Property reference / 120001101131 / 13

Type: `"string"`. Required.

Namespace where the Origin Pool is created.

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

- [no_tls](resources--origin_pool--reference--group-001.md#canonical-3232201211033121-2320120332221003-2102323230100223-1300312333300223-3332330230003310-0213003320030302-3013211011131001-3021000212231031): complete subsection reference.

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120): complete subsection reference.

<a id="canonical-2332110001010220-3311320331201032-1221232122321010-0331301231310313-0122021030131110-0201203221123012-1211002121112313-2231113213132103"></a>

<a id="canonical-2013332003310312-2233321012020211-0321020100021221-2321320223300013-3130210223023211-2320213121020230-1021332211331303-0220220213102303"></a>

## port property — Property reference / 120001101131 / 14

Type: `"number"`. Optional, Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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

- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-3110220132232202-3223311123032232-3333333030113010-2133200310011011-3010332223210221-3113000210333321-1221001103212033-0210001201112033): complete subsection reference.

- [timeouts](resources--origin_pool--reference--group-002.md#canonical-1322321330320022-2021311033121221-0210313231232031-3331330210233122-2333300130103032-3230223313322102-2023330330022001-1110122123303032): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320): complete subsection reference.

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133): complete subsection reference.

<a id="canonical-2110121321232313-0113230233203212-1012102220213303-0332333221301020-2333130223132000-0023221101110000-0300130103110310-1221213100211333"></a>

## All schema paths — Property reference / 120001101131 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](resources--origin_pool--reference--group-001.md#canonical-0012213120000312-1203300231333311-1021321221123332-2021123001320201-1123332310021300-0223131021320232-0001120311122000-0300020000303222) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](resources--origin_pool--reference--group-001.md#canonical-3133330103202322-1023102201202321-3313110101131032-2230311012100111-0002112332301222-0103102121032031-0223133212323023-0111211121331031) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-3121121110033323-1330220010212321-0130030222312201-3003233131223112-1100203202103013-3010130323013330-3011331020003312-3120022123232233) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](resources--origin_pool--reference--group-001.md#canonical-2102313010133333-1220221113221301-0131312030000000-1102333210123102-2022320331202232-0011221321330113-0020032203211112-0220133112332113) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](resources--origin_pool--reference--group-001.md#canonical-3031301032133230-1201023103020103-1102221120110032-3200223001303300-0022222032020023-3120213022332321-0230021000000111-1022100022030121) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](resources--origin_pool--reference--group-001.md#canonical-1313203122102302-3313012223123011-2011213302212212-1211222222333322-0011212013010331-3012331322223021-0013112313323003-1131023203020032) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](resources--origin_pool--reference--group-001.md#canonical-2022221301110011-1001232211213020-2330232332300110-3132103331030002-2112221013121331-3122122313023103-3032233333333113-2221323200021120) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](resources--origin_pool--reference--group-001.md#canonical-2301013301303032-0120103130032313-2030222321012310-1311322300300002-2022020003030122-1013333123211130-1330220212022030-2002231021020122) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](resources--origin_pool--reference--group-001.md#canonical-1332301311310313-0202321222030203-1112311111123010-2323220202131223-0323333212023000-0203321031210321-3010032310122032-3212312112113102) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-2311233310122001-2000133121133031-0130303003233303-0002202010330301-0120300310033102-3012120131321102-2331013213123302-1333320123303121) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-1201223203033000-0212200210113002-2123012022103202-1323311030032110-3000333111011022-1023201301323013-1102130202113203-0320312121103003) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-3312211303112211-2122002021021103-1112212232302331-2132111331022013-1202011013223311-3302000212300212-0132113101330010-0121331121232312) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-2212201332103112-2302001020103013-1132313320003233-0300213312212312-0133202221110012-2130320122232013-0312202013323213-3002323223233101) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-1223213201110302-3203023132210132-3321002121300031-3122203113302013-2030112122202131-3202003103330011-0330311133032133-1010331221030333) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](resources--origin_pool--reference--group-001.md#canonical-2221002121003202-0110323122320323-0312031200322111-0133202102021233-1201000233003103-3232111033000131-2203132332312322-1021312232200213) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-3300113001032211-0103032032303033-2311013312203333-3000322131122030-2023130130303202-2121231202130313-3122011231212221-3222332330203013) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0130213121302133-2302311313233002-3133011200110221-2312031121330231-2021233031010111-0330102012311323-2120302100122000-0023222310232122) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--reference--group-001.md#canonical-0112103101202323-3021313231233213-0220331031313320-0102102330103002-0023313023323213-3020031113032313-2020033212111321-3210123312201201) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-3121100310212312-3312122233302321-1001200330030022-3232330223331100-0232210012102020-0202331311323113-1121113132231221-3202010300311023) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--reference--group-001.md#canonical-0100213033301133-0233012321333201-1223031300021022-1212133120202303-0222302300231032-0110022032212202-1301320001121312-2020221011202023) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-3003103130030232-0230000031012213-1301031101001233-2210110120010023-2223331030120211-1121123133300123-0323221230011023-2121012000202312) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](resources--origin_pool--reference--group-001.md#canonical-1201301112230012-1311311020221001-2332033101303231-0303102010020213-3112203310100332-1303100232123233-2232312302213233-3322303330232013) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](resources--origin_pool--reference--group-001.md#canonical-3121222233001100-3031133233230220-0323303312010233-0010131132031320-3033222030310130-2123102131212232-3020002133201121-2203110330222211) |
| `advanced_options.http1_config` | [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-2132020013013033-2323101233311313-3211303313122120-3102302022100333-0022202022113202-2012032011321231-2302330103122302-0022032131123222) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-3121033213333332-3112123233111100-2021210313333332-2212300333203100-3020011233300311-0001013211303123-3113000132232123-0011023321303101) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-1030313301202002-3313010133302201-2212221203301002-2312110331010101-2110331020102220-2033101321130103-2201030003333002-2132233011332002) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-2113021231033330-1222230130333113-3232331122330003-0311100210231011-1222100200111020-3233302113013123-3200120232200310-3000231131301200) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0203033301120032-1302311331200121-0133313200231021-0131312332232300-3112102211000320-0310010132003312-0320321132020112-3123032130231213) |
| `advanced_options.http2_options` | [advanced_options.http2_options](resources--origin_pool--reference--group-001.md#canonical-2223312110301032-2021330112213312-3032301011213023-1310031223110232-2021203303111330-0012332111231101-0302122133111103-2002002000022301) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](resources--origin_pool--reference--group-001.md#canonical-0001032332222230-1232102121000202-0223132112001333-0301133131132123-3302031331201113-1212301110000123-0100310303122321-0220022202310212) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](resources--origin_pool--reference--group-001.md#canonical-1321010201322222-1130133001201211-3121010322233010-3330321322232321-3133201010030331-3033221133311122-3312201231002102-2213330131233312) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](resources--origin_pool--reference--group-001.md#canonical-0130231302123131-1313302130002302-1120311231300200-1122311220210332-2223002102212200-0131303313120211-3311033101321000-2101011110101211) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-3022211111030303-3303210112231032-3021322222331213-0112210010001200-3301230202232230-1333221121031111-3310020003001123-3221033220022021) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-0333321010033130-3320313302301220-2122011100131333-3213123332220331-2232001112011220-1120330033123332-0302232303012031-1233011330211012) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](resources--origin_pool--reference--group-001.md#canonical-3303000202220013-3031232033223310-1321323003321201-1010000202210313-1202211123013022-0303123122302011-3300002131010223-1100212211121220) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](resources--origin_pool--reference--group-001.md#canonical-3222013223202212-3021331003111312-1122020202020233-1223310211021313-3333132332131022-1111030030331122-2313200201332021-3321323232013330) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](resources--origin_pool--reference--group-001.md#canonical-2203321212323023-1033222332001110-2212310130113320-3020323000203022-3310323033013031-0302132223110303-2313321311313323-0103322133211113) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](resources--origin_pool--reference--group-001.md#canonical-2032210130223023-2102101312113131-3032201012121110-2133300213313112-3023001221202213-3013300002310031-0133323100001111-3323313111103212) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](resources--origin_pool--reference--group-001.md#canonical-2130211122313303-3100312301022030-2113233133123103-2321321302223121-3330121120122021-2031222311233021-0213233320112322-1021001232101223) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](resources--origin_pool--reference--group-001.md#canonical-2122121332233031-1110123232210002-3012111131220131-0011000203331013-2302333122222202-3112322331313021-2310221000331300-1210301323303112) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](resources--origin_pool--reference--group-001.md#canonical-0112331030230032-3201013022032332-1120103003101220-1323221003010321-1230112030200203-3121333210301131-2020220232203201-0121103031133120) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-1112010023001330-1123320210030303-1111013103121033-2333113200220132-2321321103032213-0230230010023212-1222230223232233-1302030120202033) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-1130332102113212-0021332032233120-2200102202030301-0030013222102321-3123301113201332-0232222011303302-2100100222311301-1003221222301131) |
| `annotations` | [annotations](resources--origin_pool--reference--group-001.md#canonical-2100322311221202-1120231030223213-0020212002010332-2230101101000232-1323012111200011-1201322030203021-2310021003013113-1033122011022230) |
| `automatic_port` | [automatic_port](resources--origin_pool--reference--group-001.md#canonical-1303022022111110-2200312312101231-1001130323122311-0202303312301223-0312030312212121-1001200100312002-2111111312223331-3333331320211201) |
| `description` | [description](resources--origin_pool--reference--group-001.md#canonical-0000200000202123-2113330101101213-1110213113131200-2003022220230302-3020303110123033-0312221222230301-1031221302220210-0031201331222332) |
| `disable` | [disable](resources--origin_pool--reference--group-001.md#canonical-1101102111030330-3302001311200000-2310001203020230-0220203012033322-0131112223120231-0113132201123023-2122321203120203-0030102022020023) |
| `endpoint_selection` | [endpoint_selection](resources--origin_pool--reference--group-001.md#canonical-1033200032010330-1131010303302213-0311222023330001-3120232131013322-2201322130131030-3330211222013010-2120232011110310-0312332102311032) |
| `health_check_port` | [health_check_port](resources--origin_pool--reference--group-001.md#canonical-3213010323303321-2133332333032101-3311011211021013-2331301221331003-3031220220201211-2201300222031211-2331112312201313-3221132310000022) |
| `healthcheck` | [healthcheck](resources--origin_pool--reference--group-001.md#canonical-3311000013100201-2012002213332223-1010102011213333-0021131001203021-2122012002233132-2211211221330320-3003300301311301-1333121120220323) |
| `healthcheck.name` | [healthcheck.name](resources--origin_pool--reference--group-001.md#canonical-3102130011331331-1210122203312100-2200033302232230-1020113322120222-0131301331003312-3211132330031110-0233103330331213-1311202013012001) |
| `healthcheck.namespace` | [healthcheck.namespace](resources--origin_pool--reference--group-001.md#canonical-3232130210112223-0331032011003301-3133012330113111-3210012013333021-3120100223320020-2131213033021323-2320003311013231-1211011203221012) |
| `healthcheck.tenant` | [healthcheck.tenant](resources--origin_pool--reference--group-001.md#canonical-1103333033322212-0313232301030200-3310102110331110-0320220033213023-1311033110223013-2210010333111213-0131132232331103-3301100032030222) |
| `id` | [ID](resources--origin_pool--reference--group-001.md#canonical-1103110211131123-1001200020120223-2013131001232202-3220010013313230-3132303322101013-0122331230303030-1113333331110212-2231002010313113) |
| `labels` | [labels](resources--origin_pool--reference--group-001.md#canonical-0030131313221310-2123001103101032-0211102002001022-0323202013222021-1203332321330120-1220001230132330-1112030312210331-1303321232213120) |
| `lb_port` | [lb_port](resources--origin_pool--reference--group-001.md#canonical-1203333113211100-0000103313212022-1300332310021033-3013101120332310-1310310323010220-0222223000033030-0110200123301112-0211322003111020) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](resources--origin_pool--reference--group-001.md#canonical-1010102003323021-1301112120233030-1233102213223112-2123020123200103-2303203201223321-0232002133112311-2311021320121332-0131221333023101) |
| `name` | [name](resources--origin_pool--reference--group-001.md#canonical-0200313001011322-3023100130023333-1220202120003332-0201032101122020-1022123102303010-0122100120132021-2022110113103121-3233231102100232) |
| `namespace` | [namespace](resources--origin_pool--reference--group-001.md#canonical-1001320313030032-1002211022020002-1020312330001121-3333103332110211-1331000133332011-0222321003300300-2323103330010303-1110110102221000) |
| `no_tls` | [no_tls](resources--origin_pool--reference--group-001.md#canonical-0031212003112103-0032021023323003-0332121111010323-1200301212121200-2113222033331332-1102230131001230-0012333220100001-2103312132322120) |
| `origin_servers` | [origin_servers](resources--origin_pool--reference--group-001.md#canonical-2001322203132131-3022210213212113-3312012003100000-0030021333032022-0213330213213133-3303011303133031-2200012302331213-1200003120130011) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](resources--origin_pool--reference--group-002.md#canonical-2232011323023031-2123202320003122-3001223113102021-0300230212012002-2302023131212310-0111230202310022-0330201131312223-3300232200032111) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](resources--origin_pool--reference--group-002.md#canonical-2122320211101032-2021033213233330-0032001101221103-3123301221332330-3111312222001230-2133120330200002-1231313132021133-0133331103123121) |
| `origin_servers.consul_service` | [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-2333113120233112-1200100221101203-0113323303001231-0100302303021323-2001233003012230-0230313310333100-2311302022213231-3223011030233031) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-1110003230323333-3131233003322202-3030200013122101-0010032013131331-3233122313003310-1131122003320020-0020331133123233-3320333111001322) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-2232001021110332-2302033301112320-1120323103323230-0013222133123302-2233200031101031-3002003300003130-3333002210133323-0331331331012220) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](resources--origin_pool--reference--group-002.md#canonical-3331121103330031-1122123103111100-1022301220121030-2033300032111112-0000032312020012-0202330133300332-1201103120300330-2322010231020003) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-2310223132002331-1020201010020121-2312303330130331-0103132010130222-1303311030100202-3122103303022111-2231301021212331-0213310312013010) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-0102200331031133-2200301213330313-2022230102112212-3200231331300203-2122202133333020-2021003201001112-3211021131320011-1133231100030030) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-0230133131031313-0020111122332201-0020233003303311-0033013122321333-2111022222011123-3020232022131123-3203320220031031-0103320321101012) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-1320320210002112-3300303032132200-3130312230203123-2221223121311322-2121212001010102-0021130300220122-2223001210123221-1332303013312222) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-1222232010212320-2203033013203330-0310302322212333-2102021022303033-3032030230213112-2312322023023231-2303010121230031-2011203301333002) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-1021233000001300-3231302203020232-0310123302131333-0313320330311122-2011300001133002-2021222311202230-0031000233303033-2203112002201112) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-0103103310212302-1122012323220112-2213131202322203-0032231032310320-3101011003101031-2020100200332310-1220333200011201-0221000123020300) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-3201013012023030-0011210031222311-0223113032203201-3012301322020320-3022111333111132-1131031132223230-2113320032123311-1012200033021133) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-0330031120312222-2102221012032032-1332201113313100-1101011301112313-0210002301311302-1011101120020310-1211313302013311-1220312201301132) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-3310112312023012-2030333223002320-3223031313200013-3132013031131332-0011203013002010-0320303321120323-2023303001223110-0120100321122332) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-1211011012230320-1213331000002222-3332122112303332-0320110033300211-2331230010122230-1222003100012031-3311012222110131-1113130022031221) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0211201212313032-1220211223222131-2221020302131322-3033301111320201-3122311310023333-1202131220131030-0321032330022113-0020231023331000) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-2020120101220132-1133213231212013-1230130013023000-0032220131233001-2202322011030131-0301310311112132-2012320021032132-3100332302323020) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-1302022212330303-2203322202131021-3122332203210233-0223300301103330-1033111103332132-0013311211103211-3121130131121103-0320231132011301) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--reference--group-002.md#canonical-2220012331112233-1013030220230023-2202120000200202-3202200203332230-0100201030310220-1230133131232303-3022312310213330-1223130103311121) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](resources--origin_pool--reference--group-002.md#canonical-0003320112113300-0233311102103212-0032210321031003-1311112023212300-2021230322230020-2311330032120001-0231210233323020-0121012232203121) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](resources--origin_pool--reference--group-002.md#canonical-1213101011112013-1233103301132302-2003300300212122-1133220220002211-1333110113112221-1233221232000313-0300111311300232-0312003120102112) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](resources--origin_pool--reference--group-002.md#canonical-3330231212331013-2101100320123121-2312213221032033-1100233211312111-0320330002222031-3200032130101023-3130301011122012-2210101201330221) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-1030300312310013-0232032320031233-0321120322100000-1203211033132131-0011131303210313-3312332103010313-1112021202002011-0333111102120233) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-1023300302332022-1002212333203103-0200330012333210-0302332010031111-1320013012311210-0100021132200333-1022300201013233-2313233023220331) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-0200213003113132-2023222231331310-2230113323301202-2211222031132011-0033332200031130-2013120323322321-1330102032233232-3101213331030010) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](resources--origin_pool--reference--group-002.md#canonical-0311121300123202-2221210100230022-1233300103220023-2013022122201102-0101123301122032-0120131013101312-0112020211120133-1230322023012303) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](resources--origin_pool--reference--group-002.md#canonical-3121002331113031-3230020310313222-2210212222223132-3211103300021220-0001213102332021-0302210231311112-0303111122301210-2130131333030202) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-0210110101223301-1332211132103003-2122101201000300-3232202220003030-0323132200322021-1031030230301100-3013303132331112-1301002231102332) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-3132212002001320-1011333321223332-1320123123221213-3111323223031202-3012133121203111-0202303002033130-3131332300331001-2333322330123022) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-1231302310302222-1101203001232232-1130302320321033-0000320130103230-0001002231012230-3311130121303031-3032110323331223-1113013012233012) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-2000013301321012-2001230112000330-0201112100203003-1113213330132032-1300121001130102-0122300001311330-0002333212332130-0011100301221220) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-2210010001330323-0311122232312312-2111301223332213-2200102000133300-0032322003322331-3211200212121033-2333002310233103-2320202311320212) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-2202122132021231-3321020200323200-1310202030313310-1330230003113200-2023120102022020-3101323010320321-0320133022010001-3312132023330131) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-1111333110031032-3213113231122321-3132230132020123-1123302333233320-2201113113032032-3301013313120111-0211131021101103-3011311232300232) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-0312030111303223-0203213101223030-1001010012130330-3010221020030021-0031302012311123-2301331321021112-1202202012301222-0111103021310012) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-2113321233311132-1131200212200302-2320031233131202-1111221313012210-0131110123231321-3201303030032220-0202212002031331-1203222023210300) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1313332303200303-0131030312220311-3230121111020130-1203211302210213-3301211032330300-0012033312120030-3301312322223121-1032312222131033) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3210010303322003-3021021200000033-3323212120110303-1332000331023212-1133233102021230-1322212231333012-2220131303130310-3120202011301203) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2301002222010202-0121131311200100-3203331322232323-3313020223203210-2101213201012213-2113030120331002-0013111231101320-3202312001002031) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-0232203132122110-0010012323011322-1033300302032200-2222120301013310-1310102210222222-0033302213333223-0023221332213222-1123212130320123) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-3232303302233102-2031023011300221-3110021223111212-0022120200222200-0311213000033220-2201201100032011-0100022002301201-2221320203122020) |
| `origin_servers.labels` | [origin_servers.labels](resources--origin_pool--reference--group-001.md#canonical-0222322311301012-0302213113300121-1303133220100113-2310201011020013-1133022301330223-1211000303231312-2232032222303300-2231312012013323) |
| `origin_servers.private_ip` | [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0132122003323203-1113202003201132-3321320022031212-2230321230311222-2210101323121301-3303110010030232-3303110003233000-0303302222121003) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](resources--origin_pool--reference--group-002.md#canonical-1010312132210302-2122330213320320-3002121000121223-2321111221333320-3233220102120023-3111220020033121-0022020003220201-2003211121111023) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](resources--origin_pool--reference--group-002.md#canonical-3131321310122213-2013013222233311-2013100113320312-1111123221322122-3013030033122032-2032133200310211-2223132120100033-2301131030132321) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](resources--origin_pool--reference--group-002.md#canonical-3020123121020013-1233223013301111-3001332131013133-2303021113203001-3120321223301031-2112101011011222-0221013312230300-3030013213102321) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](resources--origin_pool--reference--group-002.md#canonical-0122201332003012-3133322332232022-1210211032033221-1113102110020313-3132031203113300-1100300320102320-1110112012323103-2100131212133310) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](resources--origin_pool--reference--group-002.md#canonical-1223002113120231-2231221222212323-3200032101100123-0100301000333301-1310301122013232-0130201212230011-1121333303122332-1211330012331332) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](resources--origin_pool--reference--group-002.md#canonical-0002023023112303-3000002202301301-2122112011220210-2300210311301112-1310303122333301-3023310112222330-3112020232133100-3013231300223301) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](resources--origin_pool--reference--group-002.md#canonical-2211231312012130-3330102313111132-0010213002332311-3333322201120203-3210311203321223-2202121223213200-0210312203332122-3122031211132310) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-2300231310100223-1133011131133232-1022201302332233-0201232012331201-2310030012112201-1312110221032323-1012213021333220-1123101310031303) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-2331020232213023-3031110031113100-3210011130022031-1310130031320000-1301021211133300-0320112000033032-2013320111122100-1133023031220313) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-0213220102003313-2321021101003033-0112323000222111-3312212011301311-1330003232101003-1220312113223211-1331121211323313-3121012013301312) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-0310332310121302-0012001110320332-2133332133310001-3302331031221212-1122023203031010-3121012220302020-0011011002012001-3331221202100202) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-3333130012100110-3133010102330332-2230012130101130-2032111033113120-2101210230323002-3121102032110030-0031220330012002-1022302023022111) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-2222230231021133-1310332222203132-2213203023213220-0203330233023131-0122101031012301-0203303032321110-2230021233220330-2121021123221222) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-1013210020122120-3113310033130023-2203233010131300-1232120113110230-1012120320130103-2210000321001023-2232103311023030-1331321022012011) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-2200220020123203-3033213310312002-0113221220033300-1103023110203201-3032000302331022-2001131023203233-0200213310311000-0133101003022121) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-0323020000231022-2321300331100223-1221131022231231-3322133121100023-3202320023222230-0110102231120112-0202310122123211-0220310023123201) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1230022110130033-0122300323220323-2233303320012120-2123001031323003-0331000322013032-3031212033032211-0210003203200313-3110000013232222) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-1213000033011202-1103310310032130-1003133221100132-2020230011311220-1103111231331110-0213111131321230-2330010102223102-1132132222213302) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0033010020000203-0223123013022223-0200013323230300-3310100102021313-3100202201211321-3000122310222333-3233331313312012-0111022110330203) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-2310332121303123-3333230133231031-0211120300333113-1202323232220302-0032213011303110-1313113201221230-2313031222100132-2103101033132003) |
| `origin_servers.private_name` | [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-3022300021030132-2220222012003231-1032303111232231-2212212013120222-2323130212232113-0112320023321113-2313230210023201-3013310312203012) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-0030320020312221-2112213011032231-1010223120203032-0030320313030110-2330203033200010-2000333313312010-1000133300301332-0301233013232223) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](resources--origin_pool--reference--group-002.md#canonical-2322312020100021-0201033330302321-0022320311110130-2203302111310130-2112111212210332-0030012002320200-2020220333231022-0100020021213131) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](resources--origin_pool--reference--group-002.md#canonical-0013231003130331-3120123023103021-2320131211031233-3223331122101021-2103123112122130-2020233001233000-0013323130033302-2011301100101231) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](resources--origin_pool--reference--group-002.md#canonical-1001321013320301-3231130133303213-1012133100211202-0202310003112322-0320013300303023-2210013102000202-3300202320012331-0013231020013230) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](resources--origin_pool--reference--group-002.md#canonical-1133101023013231-2111002010302323-2132111310121133-1101003132011232-1310123201023212-3333020132201122-2000001333102003-0200320101132013) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](resources--origin_pool--reference--group-002.md#canonical-3130300300300122-2021121312321213-2203003013113330-1312322123211312-2132330301132333-0003220131030032-1010311023201132-0303113001021320) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](resources--origin_pool--reference--group-002.md#canonical-2213202322130032-1302102012000111-0130011113033110-0332003132300102-3013320111020321-0123202013202103-2200010222222002-1313303132130320) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](resources--origin_pool--reference--group-002.md#canonical-0020132303013213-0000202322312013-0130223103003201-2320022211103132-1011010001103111-3210213313122021-0323000330112332-0230310211022122) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-1020001220021113-3201030222113010-3233123302013110-2033130012033221-0330122021331100-1113302312232110-2312031322302123-3113203103030302) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-2310110312030231-0323110012101103-1231313132110123-3322003102003131-1320220002300020-0210320122211021-3233202103231300-0312122000232233) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-1233332330010203-2103101120011230-3302022311211002-3131021321120101-1030011110033031-3201111031230123-3322103220002102-0023103113220323) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-2030110230122011-0121312231332013-0103303202223333-0202112032330303-0111300300121002-0020021300130230-3302131213012013-1302221311122312) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-3122311300323320-0211222202200211-3102330002320302-3333310303332213-1212202113332222-3222331232300233-2232323000323223-1033321013103023) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-0310202213011000-3111221033301021-0000311230003001-3133110320333332-2332303200200020-1133002033223110-1113111032333320-0303223221312122) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-3230020311221133-3320203223320221-3003211320333331-0100110201013203-0013132002113213-1033002220010200-0032321303320130-3003202000212202) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-1211202113232100-2103013212012101-3223031333022230-2020222132122330-1013333210103312-2133112122232022-0231130221203310-1301032203320222) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-3333120331021200-0122122122112221-1033232311130311-2001001031121012-1323101121200221-3031223233321322-1020232202213333-0032031311223013) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1012322211101010-0203023132110100-0321333121033313-1121311033211030-0012132301213012-1133130003230120-0020130231013323-3113120111010133) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-0113203312310213-2110203230230023-2212211321322013-2102102302112231-3013121131100003-2120303003130323-0233330323032322-0223131100133121) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-1012022110102030-0322310000331011-1121232011213022-0223112103323312-2103220000323222-3113313200023112-3120001311231113-0031212113330001) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-0022211003233101-1123312210101003-0022330320112012-0103323033222030-1222232031131222-1020220031211132-0033330032101220-2332200002003220) |
| `origin_servers.public_ip` | [origin_servers.public_ip](resources--origin_pool--reference--group-002.md#canonical-2330102103323312-2311122332333212-2111322303321021-1202231123132011-2230321120110033-1323002102021311-3322132202133231-3130033013112100) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](resources--origin_pool--reference--group-002.md#canonical-0301113230322103-0122131333211003-2200221332132002-0002323200320120-0322023121113332-2030330012133033-0131102333011030-1221302101232132) |
| `origin_servers.public_name` | [origin_servers.public_name](resources--origin_pool--reference--group-002.md#canonical-0222302133210320-2001221203003010-3022300302110110-3000333100030103-2313301012110103-2212210120131013-1012130222122313-3202312023112330) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-2111122303201022-2233210133031300-0320303100012233-1103103233312021-2032231230310321-1120300132031101-2002033321232112-1233310112133232) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](resources--origin_pool--reference--group-002.md#canonical-3021102323312312-2101333023321220-1102102211203121-3133300301203323-3213311103001301-2121213022101223-1023111301321210-1011010022020222) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-1212301103220103-0222133311031102-2310030020003310-3011032130010130-0333113121203022-1121200301130110-1121333133231213-2122303322012222) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](resources--origin_pool--reference--group-002.md#canonical-0303031013321122-3312002010330312-3300200032023111-1022212011213113-2232202312122132-0131302301220001-3101202011123200-0312301013023023) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](resources--origin_pool--reference--group-002.md#canonical-0003212300322202-3222000200122020-3203010013001101-1311323200003200-2102323331103133-0230313302203300-3302211300312303-3202332311301130) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](resources--origin_pool--reference--group-002.md#canonical-0000202302112102-1112012101000233-3302013312100000-0023221022201010-0201113330200201-2003201103322210-1330030120230101-2002210032322331) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](resources--origin_pool--reference--group-002.md#canonical-0223223101122021-2213133211011111-1110303311033212-2202212310201012-1232122321032311-0022200133002320-1203111330100322-3332031233320120) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](resources--origin_pool--reference--group-002.md#canonical-0111322301320300-0322201111323203-1212030301001200-0211310021223012-3123201121032322-2030010332303230-2231323310021331-0120323133322133) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1133331232233003-1320111130012311-1123222301130121-1020122210302320-1322110011220312-3110310013210300-1130202300310203-1332113133000131) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-3002030022032132-3311301011121132-2213132122130322-0202002203100100-0201311301120102-1020000023121010-2020131020121000-2302221200332000) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](resources--origin_pool--reference--group-002.md#canonical-1133303321023212-2001300031121022-1011301211231302-0001021121033021-2323331011202103-0022221131021002-1223211000123212-0002221220103230) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](resources--origin_pool--reference--group-002.md#canonical-0211012010133233-1230101313033231-2203330313202001-0230213001201123-0023222333323120-1021301221110101-3130101312333010-0022323203212030) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](resources--origin_pool--reference--group-002.md#canonical-3001031121122201-2021001100001030-1223130321201021-2103201022330133-0031030202301102-2003332030311030-1013301002031230-0333123312320300) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](resources--origin_pool--reference--group-002.md#canonical-3203203021110231-0111012212301333-0231202321023031-0200023321203132-2122323200121130-1202232210232212-2112323113223002-3222230212232320) |
| `port` | [port](resources--origin_pool--reference--group-001.md#canonical-2332110001010220-3311320331201032-1221232122321010-0331301231310313-0122021030131110-0201203221123012-1211002121112313-2231113213132103) |
| `same_as_endpoint_port` | [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-2003120102323102-0320322222233123-0110310222321212-1303023300232011-2231113032010222-0322332011312201-0332033220131211-2002332200121120) |
| `timeouts` | [timeouts](resources--origin_pool--reference--group-002.md#canonical-0000231121303330-2003202121011221-0020300020231001-1223302100322112-1221333321301112-3331212200332003-2111332002333011-3211101330111331) |
| `timeouts.create` | [timeouts.create](resources--origin_pool--reference--group-002.md#canonical-0322021132111321-3112303003031330-0130110311120000-2323112211313121-2012101131003231-2310330113310310-0320302320102011-2221132032313132) |
| `timeouts.delete` | [timeouts.delete](resources--origin_pool--reference--group-002.md#canonical-1123132211211221-0212322301210311-3001010203232213-1212101120321000-2222310231130132-1231023213030003-0031133032232203-3212022030111323) |
| `timeouts.read` | [timeouts.read](resources--origin_pool--reference--group-002.md#canonical-0112102231203010-1002213233010322-3013322001331311-1200201112200033-3200131003022301-2112112000322322-2232133101312103-0233312112020123) |
| `timeouts.update` | [timeouts.update](resources--origin_pool--reference--group-002.md#canonical-0110313103330301-1331101223223003-0022300330300231-0323203311011100-3322001332200002-3310012203101213-2201112330320131-0311333030233120) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2020113323212003-1003032002212201-3230111332303132-1132100131321213-1302302110200330-2111203102330131-3033100023112000-2101313201120202) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-1302332230212100-0332010233013101-3300102301101123-0002201113311001-1130033132122002-0201100203010021-1301330033303021-3000221202311030) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-2101101012002111-0212122133332312-2211023230011231-2323313110333203-3111100301202012-3302031232001111-1000203100112302-0033120332203101) |
| `use_tls` | [use_tls](resources--origin_pool--reference--group-003.md#canonical-3122122211032132-3132313203100001-1013000203202031-2132002223311123-2211000023310210-0233300322102210-0033100013333133-1022321012001012) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-3030232203023330-2211203000103303-0030123222002110-1210003113321012-3223303023211312-0300132233212333-0312103033022102-0100002311022323) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-0133330110333312-2101233020121221-1212201013003111-0213300012112330-3332101120100012-2122111011331212-2102213012002122-3212032202130203) |
| `use_tls.disable_sni` | [use_tls.disable_sni](resources--origin_pool--reference--group-003.md#canonical-3121221222210030-0011023312311000-0213231013120333-3321330130232201-1122111133021030-0322220021233033-1301012022233032-0331111232122201) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](resources--origin_pool--reference--group-003.md#canonical-1010332301301203-1302100023030012-0012312120112131-0223010111101231-0301233110003032-1101022330100330-2331001011130013-2200321301310220) |
| `use_tls.no_mtls` | [use_tls.no_mtls](resources--origin_pool--reference--group-003.md#canonical-3011203232233111-0210021203020023-3032203220101031-1110110203331323-3133231000132032-0101010211121311-1223020302132201-1013232312121222) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-2222233003100111-0302003010100213-2011310100011121-3002311212312332-2311313311313130-2211231120332133-1223313220232221-0113201012312301) |
| `use_tls.sni` | [use_tls.sni](resources--origin_pool--reference--group-003.md#canonical-1220200011310322-1313110000203111-0131320231230110-1312300003202311-2021312133301200-1203030202220211-2233131120002033-1310221211332200) |
| `use_tls.tls_config` | [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1020312000221113-0332112223333302-3122020211012330-2001011231100112-2100330310322121-0023311223022300-1202133123232313-2202220123133300) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](resources--origin_pool--reference--group-003.md#canonical-0203303001311310-2322002212132332-1010120123231232-1113000203111122-2201232211310022-1032120102202233-0332023301120102-1323302313333132) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](resources--origin_pool--reference--group-003.md#canonical-3203013203033320-2103302313120031-2110301221221120-1030331013001031-0321112320310212-1020103003323321-2002032100001303-0312103212130330) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](resources--origin_pool--reference--group-003.md#canonical-0122213110200311-0101230001332320-0130020201122132-3113320102110003-3223310321312132-3232133210333113-2323320131111230-3101002032132012) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](resources--origin_pool--reference--group-003.md#canonical-1233311302332033-3103001123311321-3313220110300313-3012312233103323-1211312223130202-0123001303100212-2233321301012023-3333203132331202) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](resources--origin_pool--reference--group-003.md#canonical-2200230203011232-1012310231101221-1302222003121322-0000213033303310-1321100332211312-3202202332303220-3101323313203001-1331301111100332) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](resources--origin_pool--reference--group-003.md#canonical-2300211011100131-2312131300032321-3133020322230003-3030323201100300-0020003002220020-1031212102213310-2202030002123013-0032322311122213) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](resources--origin_pool--reference--group-003.md#canonical-3220202323330311-0003102032032322-0223033101210321-2013332030011010-2213021230323322-0302133022011001-2330131022102303-1020303313021110) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-3010233110232013-0320031332313011-3002333211230103-3232321131201122-1230310031110203-0133221320122110-3131120320021100-2303232111101310) |
| `use_tls.use_mtls` | [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1011223301223132-2013212102312122-0221212033123111-3331012332113120-1212201210230330-1203130033210303-0303021021321123-0333022103120211) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3322232030000201-2023133321112003-1230111113231113-0103010032120300-0011131231103233-2012100003111330-1021213301220103-2120010120303223) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](resources--origin_pool--reference--group-003.md#canonical-1020333333103300-3203001211023011-0331100312033032-1331323202120103-1003032303313233-2322333023311310-0003202103221112-1101012322013010) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-2213223211110223-3021301202303321-3000101311331301-0110113113100000-0031030230003220-2102010111213120-0332200330001012-2310011311303112) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-0321110130030203-3123231300301322-0100330000011001-3203112331003102-3021331301133322-0300331222030123-2231130120300120-3121221330333033) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](resources--origin_pool--reference--group-003.md#canonical-2313010232001222-1321021230103222-1000002003103000-3111022130302323-2211203311020331-1320202010132122-1013002033320032-3122310031333311) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-3333300201211233-1032301121302000-2022223332122122-2223111331102203-1323103332103020-3302230030122232-1212130120321021-1111203300330332) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-2312122300111210-2320203233333123-0110010220000333-2133221212322120-3332123321213233-2112300310000012-2032231020312032-0001321213022110) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-1321030203013330-2010233100103102-0122331220200012-1133022023030330-0113333100231223-3131102101133033-2303111220320013-3013030233123210) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--origin_pool--reference--group-003.md#canonical-2013331301120103-2210112033020212-0312000203011300-3120003131103213-3203013231030210-3330031202200330-3233002212013021-0301223113122001) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](resources--origin_pool--reference--group-003.md#canonical-3102221000300000-1010023013101110-0132000101123120-1131113230111302-2213302333033232-2030201122210021-2213233022022031-3030001212001030) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--origin_pool--reference--group-003.md#canonical-0033311012223112-1130122113201000-2123313110033211-0330223023223333-2212033331030011-3203322312323213-2232033003201121-3013302131310202) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-2111132313310011-2312201030120001-1123021332221032-2101010113212013-3212022320002303-2313331301003330-2331210133313200-3023131301312032) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](resources--origin_pool--reference--group-003.md#canonical-0033032221002213-2122323020200123-2232322102333232-0323201332200320-3122021310323232-0230012231030103-3132210321021211-2012200313310033) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](resources--origin_pool--reference--group-003.md#canonical-1211332112310200-1232002021223003-3120011200102233-1332111032220302-2322202210213020-0200202211320031-2010220010302203-0221300031211030) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-0301312001311001-3013220001311121-1010021310302213-1333300031230210-0110113323013013-3201201200133221-3332230303220222-3303022020031233) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-0311020032113212-1203103100333300-3120011330330310-2133030131113310-1200133213303001-0103330110212320-0311211222012323-1000301320200020) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](resources--origin_pool--reference--group-003.md#canonical-0231211130302032-2202100003223131-3133021022010120-1031222032311213-3310111011332112-3103221020013201-2111031200302330-0323030313333001) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](resources--origin_pool--reference--group-003.md#canonical-3210200223021301-1323211003102302-3230323233302133-3221102000111102-3223120130322211-2001232023212122-0233312112011310-1233230002321120) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](resources--origin_pool--reference--group-003.md#canonical-3200110220221112-3132202102011101-3023302001331222-0320032203131012-2132103012300313-0101322313011012-1233011231000111-1333011302033102) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-1202022113131322-3223333122031311-2012231003203110-2200120033103313-0303123113012001-1021131200222210-3302220021012302-2121023331223031) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](resources--origin_pool--reference--group-003.md#canonical-3211120110011201-0121003012003331-3200121102030103-3100212002221031-3100111133302033-0212211223332131-0003023220323223-0303330120232302) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](resources--origin_pool--reference--group-003.md#canonical-2300000132123033-2103011001320010-0003201212202311-2131202021300121-0100301332031130-3100031203003111-3213000321133131-3220001111032321) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](resources--origin_pool--reference--group-003.md#canonical-0121211112011311-3022002120201003-1200133210002120-3103310112300321-1331332330013221-1322110233131121-2221323113132132-0330103213300331) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](resources--origin_pool--reference--group-003.md#canonical-3323101123113223-3112300031313201-0323233213330222-1122012021031001-1213311231311210-1321012323131120-1220113301313301-2232122310332010) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](resources--origin_pool--reference--group-003.md#canonical-1220133023112011-0133103323321013-2230232111003000-1311033011331123-2002332230231132-0210313333022312-1133232021122222-0311333013311320) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-1002302310301002-1302021231211331-0023032120010312-1013223211301323-1323121220130010-2023101130223113-3023221113032221-2012303103232033) |

<a id="canonical-3303102313133102-2102000312200011-1302230310233231-1233133131111331-1330110310322110-1002233133332302-2323000112001322-1020202101220100"></a>

## Next pages — Property reference / 120001101131 / 16

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-1112022203223032-3323221233013132-1121203303031132-1021300130303111-2312333020233311-1132310003313023-1330012320112102-3213232312123011)
- [healthcheck](resources--origin_pool--reference--group-001.md#canonical-0220303133013032-2302323323212323-0311200001213123-3021130113330311-2212330212003003-3231310310300112-2222012213121021-0030322010120013)
- [lb_port](resources--origin_pool--reference--group-001.md#canonical-0020300233212120-1303102312023111-0212101131203331-1302022110112113-0033130113202023-0130301212232010-1022111220323023-0020230101001100)
- [no_tls](resources--origin_pool--reference--group-001.md#canonical-3232201211033121-2320120332221003-2102323230100223-1300312333300223-3332330230003310-0213003320030302-3013211011131001-3021000212231031)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-3110220132232202-3223311123032232-3333333030113010-2133200310011011-3010332223210221-3113000210333321-1221001103212033-0210001201112033)
- [timeouts](resources--origin_pool--reference--group-002.md#canonical-1322321330320022-2021311033121221-0210313231232031-3331330210233122-2333300130103032-3230223313322102-2023330330022001-1110122123303032)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310322122232213-3112233113123320-0322010332120300-0322302133112103-3203031132022003-0231002033221232-1222230332313320-0032020112110123"></a>

## advanced_options — advanced_options / 303012332232 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- advanced_options

<a id="canonical-0012213120000312-1203300231333311-1021321221123332-2021123001320201-1123332310021300-0223131021320232-0001120311122000-0300020000303222"></a>

Type: `"object"`. single nested block, Optional.

Configure Advanced OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_http_config",
    "http1_config"),
  validators.ConflictingObjectAttributes("auto_http_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "default_circuit_breaker"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("default_circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("disable_lb_source_ip_persistence",
    "enable_lb_source_ip_persistence"),
  validators.ConflictingObjectAttributes("disable_outlier_detection",
    "outlier_detection"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v1"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v2"),
  validators.ConflictingObjectAttributes("disable_subsets",
    "enable_subsets"),
  validators.ConflictingObjectAttributes("http1_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection"),
  validators.ConflictingObjectAttributes("no_panic_threshold",
    "panic_threshold"),
  validators.ConflictingObjectAttributes("proxy_protocol_v1",
    "proxy_protocol_v2")}
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
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010133221230310-0230013333121113-0303001300231133-3332232112033221-1330320311100321-3330001011101332-3333122103203102-1033003230322203"></a>

## Direct properties — advanced_options / 303012332232 / 3

- [auto_http_config](resources--origin_pool--reference--group-001.md#canonical-3331110012302303-2330201011230023-3123103101203032-3003211202201210-0012201112120222-1212013010223121-0010202022300122-3210301122110302): complete subsection reference.

- [circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-0033113010002302-1233333120200310-2012223323210222-3313302231001230-3301111232102122-3311103200320333-0330213023110303-0113203212020020): complete subsection reference.

<a id="canonical-1332301311310313-0202321222030203-1112311111123010-2323220202131223-0323333212023000-0203321031210321-3010032310122032-3212312112113102"></a>

<a id="canonical-1111222232112000-1120122121323001-1301031322001201-1212000032020321-1031113322001211-1313132023213012-1220300133002031-1031101333221321"></a>

## connection_timeout property — advanced_options / 303012332232 / 4

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-3101120010020013-3330113221002001-2220200100231320-1133013223213133-0022122002312331-3303231030210132-0013132132200000-0212213313333311): complete subsection reference.

- [disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-1000203102031302-1012300312210303-1320311100001012-3003220233033001-0221232232101100-1230211123332100-3300113033201021-3122130231222211): complete subsection reference.

- [disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-1110033310111221-0311103030303223-3211303210220123-3033011120133100-2031102313201310-2302221330013212-2211323331302113-2330023011311230): complete subsection reference.

- [disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-0303111031121330-3030232011000130-3000031202321210-3313200311012320-0233221123302133-3200321002133322-3222232213001122-3301202011203131): complete subsection reference.

- [disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-2213300033233331-3203002012201223-1020300032131122-3110201221031120-1233033311101332-1000313131023221-2131220002000133-2000301223322131): complete subsection reference.

- [disable_subsets](resources--origin_pool--reference--group-001.md#canonical-0132102210122311-3313023112101111-2013130010001303-0031232103110330-3010201102310202-2032033320210210-0303113213130210-3301233212211230): complete subsection reference.

- [enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-0003003221233012-1310223021303120-2211013133202102-3001322232012120-0123301231303010-2022223113222201-3311201000011110-0230022321023230): complete subsection reference.

- [enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103): complete subsection reference.

- [http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000): complete subsection reference.

- [http2_options](resources--origin_pool--reference--group-001.md#canonical-3110221330232310-0033030320300111-0032333121300122-2203022202112131-2002121100333022-3130111233223112-1030012232320323-0122011120320120): complete subsection reference.

<a id="canonical-1321010201322222-1130133001201211-3121010322233010-3330321322232321-3133201010030331-3033221133311122-3312201231002102-2213330131233312"></a>

<a id="canonical-0031301313313233-2210130000302200-3022331233230013-1212313232313310-1322033212322223-0301030032112003-0321213031323230-3210013210221220"></a>

## http_idle_timeout property — advanced_options / 303012332232 / 5

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

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

<a id="canonical-0130231302123131-1313302130002302-1120311231300200-1122311220210332-2223002102212200-0131303313120211-3311033101321000-2101011110101211"></a>

<a id="canonical-2113020030020220-1013320020231023-0311101212001021-3213020313321013-2212013112332120-0320000133033210-1303231313103310-3020300112300111"></a>

## max_requests_per_connection property — advanced_options / 303012332232 / 6

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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

- [no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-0202112120201020-0312010230000131-1100030322023231-3232212311110032-1220210322222103-2332221300001302-3110032302100201-0210013330221330): complete subsection reference.

- [no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-1330213333313110-3100031301022021-2123130200313332-3201013222020131-2223030121213030-2220120203131331-1221101000313013-1112103003000113): complete subsection reference.

- [outlier_detection](resources--origin_pool--reference--group-001.md#canonical-3000232203110021-1031230213122011-1123333023221112-2112132033221221-0233111301322212-1122212030103130-0001131010111133-3133220123103122): complete subsection reference.

<a id="canonical-0112331030230032-3201013022032332-1120103003101220-1323221003010321-1230112030200203-3121333210301131-2020220232203201-0121103031133120"></a>

<a id="canonical-0220232220333102-2323033130122223-1301103311221301-0103010322030213-1031213202220001-2202201301101313-2302130000031213-2203221311100221"></a>

## panic_threshold property — advanced_options / 303012332232 / 7

Type: `"number"`. Optional.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-2231302013311332-2302020300210101-3013021032120300-1322020323111001-1021312211133123-0001210223111212-2210221303330123-2132302330031012): complete subsection reference.

- [proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-0320031013302220-3332100213221311-2123202301102332-2010033001202221-1032332013201311-3211112203311011-3001213113311213-2320133302110011): complete subsection reference.

<a id="canonical-0033233220311111-0213322100233033-0233303133103233-1010203123311102-1032302123003211-2330210021231312-0223233300111202-3001120130212323"></a>

## Next pages — advanced_options / 303012332232 / 8

- [advanced_options.auto_http_config](resources--origin_pool--reference--group-001.md#canonical-3331110012302303-2330201011230023-3123103101203032-3003211202201210-0012201112120222-1212013010223121-0010202022300122-3210301122110302)
- [advanced_options.circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-0033113010002302-1233333120200310-2012223323210222-3313302231001230-3301111232102122-3311103200320333-0330213023110303-0113203212020020)
- [advanced_options.default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-3101120010020013-3330113221002001-2220200100231320-1133013223213133-0022122002312331-3303231030210132-0013132132200000-0212213313333311)
- [advanced_options.disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-1000203102031302-1012300312210303-1320311100001012-3003220233033001-0221232232101100-1230211123332100-3300113033201021-3122130231222211)
- [advanced_options.disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-1110033310111221-0311103030303223-3211303210220123-3033011120133100-2031102313201310-2302221330013212-2211323331302113-2330023011311230)
- [advanced_options.disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-0303111031121330-3030232011000130-3000031202321210-3313200311012320-0233221123302133-3200321002133322-3222232213001122-3301202011203131)
- [advanced_options.disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-2213300033233331-3203002012201223-1020300032131122-3110201221031120-1233033311101332-1000313131023221-2131220002000133-2000301223322131)
- [advanced_options.disable_subsets](resources--origin_pool--reference--group-001.md#canonical-0132102210122311-3313023112101111-2013130010001303-0031232103110330-3010201102310202-2032033320210210-0303113213130210-3301233212211230)
- [advanced_options.enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-0003003221233012-1310223021303120-2211013133202102-3001322232012120-0123301231303010-2022223113222201-3311201000011110-0230022321023230)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- [advanced_options.http2_options](resources--origin_pool--reference--group-001.md#canonical-3110221330232310-0033030320300111-0032333121300122-2203022202112131-2002121100333022-3130111233223112-1030012232320323-0122011120320120)
- [advanced_options.no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-0202112120201020-0312010230000131-1100030322023231-3232212311110032-1220210322222103-2332221300001302-3110032302100201-0210013330221330)
- [advanced_options.no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-1330213333313110-3100031301022021-2123130200313332-3201013222020131-2223030121213030-2220120203131331-1221101000313013-1112103003000113)
- [advanced_options.outlier_detection](resources--origin_pool--reference--group-001.md#canonical-3000232203110021-1031230213122011-1123333023221112-2112132033221221-0233111301322212-1122212030103130-0001131010111133-3133220123103122)
- [advanced_options.proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-2231302013311332-2302020300210101-3013021032120300-1322020323111001-1021312211133123-0001210223111212-2210221303330123-2132302330031012)
- [advanced_options.proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-0320031013302220-3332100213221311-2123202301102332-2010033001202221-1032332013201311-3211112203311011-3001213113311213-2320133302110011)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3331110012302303-2330201011230023-3123103101203032-3003211202201210-0012201112120222-1212013010223121-0010202022300122-3210301122110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133213200113310-0212012312212033-3203031022311230-2310331233221301-1213023312210003-2113011113212113-2322030030220221-1300023102032233"></a>

## advanced_options.auto_http_config — auto_http_config / 123031130112 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.auto_http_config

<a id="canonical-3133330103202322-1023102201202321-3313110101131032-2230311012100111-0002112332301222-0103102121032031-0223133212323023-0111211121331031"></a>

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
auto_http_config = {}
```

<a id="canonical-3021233023013322-1020330300112122-0212320112131123-3203030032201201-2331201310113133-3132123313013233-3031220132323023-2100221112320210"></a>

## Direct properties — auto_http_config / 123031130112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322213010302200-3023322021331123-0111103130320301-3113100222200213-1331021223223010-1032310212020122-3011203112131113-1002333020231021"></a>

## Next pages — auto_http_config / 123031130112 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0033113010002302-1233333120200310-2012223323210222-3313302231001230-3301111232102122-3311103200320333-0330213023110303-0113203212020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032010023100130-3112303020121011-0203100032332210-0212320333012030-1111030230113122-2112032332230133-1203122132031133-3003212332110102"></a>

## advanced_options.circuit_breaker — circuit_breaker / 332221011212 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.circuit_breaker

<a id="canonical-3121121110033323-1330220010212321-0130030222312201-3003233131223112-1100203202103013-3010130323013330-3011331020003312-3120022123232233"></a>

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

<a id="canonical-0122331031131231-1131100021223121-2210210310000033-3303130020212201-1130212221001022-3322331103212020-3123022120200310-1010122333032031"></a>

## Direct properties — circuit_breaker / 332221011212 / 3

<a id="canonical-2102313010133333-1220221113221301-0131312030000000-1102333210123102-2022320331202232-0011221321330113-0020032203211112-0220133112332113"></a>

<a id="canonical-0333212320133203-0222122130102030-0202220232232300-2022032200230013-3131333023113113-0032030201133320-1221303333310210-1020330133031331"></a>

## connection_limit property — circuit_breaker / 332221011212 / 4

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

<a id="canonical-3031301032133230-1201023103020103-1102221120110032-3200223001303300-0022222032020023-3120213022332321-0230021000000111-1022100022030121"></a>

<a id="canonical-2320012322300020-0102010303122022-0232131222331132-0003300032102302-3332222000123010-3101020320221012-3112123110310030-0111013332233223"></a>

## max_requests property — circuit_breaker / 332221011212 / 5

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

<a id="canonical-1313203122102302-3313012223123011-2011213302212212-1211222222333322-0011212013010331-3012331322223021-0013112313323003-1131023203020032"></a>

<a id="canonical-2303323023211002-3323110003122001-0111323101212000-2200212011321212-1011220221033210-2011121021112003-0223331231013333-1331213300333103"></a>

## pending_requests property — circuit_breaker / 332221011212 / 6

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

<a id="canonical-2022221301110011-1001232211213020-2330232332300110-3132103331030002-2112221013121331-3122122313023103-3032233333333113-2221323200021120"></a>

<a id="canonical-3203010221213300-1031132221310120-0031030100233331-1020112011032210-0003033200133300-0231231321123021-3122003212221100-1212333223313333"></a>

## priority property — circuit_breaker / 332221011212 / 7

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

<a id="canonical-2301013301303032-0120103130032313-2030222321012310-1311322300300002-2022020003030122-1013333123211130-1330220212022030-2002231021020122"></a>

<a id="canonical-2233033013121031-1200020330122203-3313131001222002-1332133313330320-1013032023101330-2011210313113321-0101113212111002-0121001011023200"></a>

## retries property — circuit_breaker / 332221011212 / 8

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

<a id="canonical-3312323103030111-3322130032020200-0323300023033133-0110213030111323-1101201313233302-2312302213201212-0322301202110313-1232012000021312"></a>

## Next pages — circuit_breaker / 332221011212 / 9

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3101120010020013-3330113221002001-2220200100231320-1133013223213133-0022122002312331-3303231030210132-0013132132200000-0212213313333311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101331001233322-3231320113031022-1130302313310332-2200333312132220-3102030113113112-0003101220031230-1130131031003112-2130113010203133"></a>

## advanced_options.default_circuit_breaker — default_circuit_breaker / 202113333100 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.default_circuit_breaker

<a id="canonical-2311233310122001-2000133121133031-0130303003233303-0002202010330301-0120300310033102-3012120131321102-2331013213123302-1333320123303121"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
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
default_circuit_breaker = {}
```

<a id="canonical-2313010031322102-2322332223132101-2210120101133211-2132113131123021-0133303303121332-2002232131101103-2032031330031021-0210023021331023"></a>

## Direct properties — default_circuit_breaker / 202113333100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111011330002213-2332210212113111-2030323233130111-1313132013001033-3233121002131212-1331111012303232-1132321131103213-3022023320220200"></a>

## Next pages — default_circuit_breaker / 202113333100 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1000203102031302-1012300312210303-1320311100001012-3003220233033001-0221232232101100-1230211123332100-3300113033201021-3122130231222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133100211121001-3002301300030301-3231022221323023-3211310011030203-3210213331310131-2220013031030312-3122122132322302-2322320000201311"></a>

## advanced_options.disable_circuit_breaker — disable_circuit_breaker / 030303301323 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.disable_circuit_breaker

<a id="canonical-1201223203033000-0212200210113002-2123012022103202-1323311030032110-3000333111011022-1023201301323013-1102130202113203-0320312121103003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable circuit breaker.

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
disable_circuit_breaker = {}
```

<a id="canonical-1223013321213300-2030020112132132-3320313132303213-2112330332001230-2332002113033303-2113002310011033-2322201230032321-1122202023310302"></a>

## Direct properties — disable_circuit_breaker / 030303301323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230323013001101-2101112333023301-3133030111120130-1123333021031231-2130100010022002-2033031012012311-0132111111202111-2200112332122000"></a>

## Next pages — disable_circuit_breaker / 030303301323 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1110033310111221-0311103030303223-3211303210220123-3033011120133100-2031102313201310-2302221330013212-2211323331302113-2330023011311230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333302233023102-0031103110302212-3213001233212202-3002220313100330-3131032121201102-1112132103330300-3121232210012100-1302331231030323"></a>

## advanced_options.disable_lb_source_ip_persistence — disable_lb_source_ip_persistence / 130000030330 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.disable_lb_source_ip_persistence

<a id="canonical-3312211303112211-2122002021021103-1112212232302331-2132111331022013-1202011013223311-3302000212300212-0132113101330010-0121331121232312"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
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
disable_lb_source_ip_persistence = {}
```

<a id="canonical-3201021312233011-0323120110210032-2123233311033230-3113302103230212-3332133030112030-1303302113223312-3133311231332011-2332200120011332"></a>

## Direct properties — disable_lb_source_ip_persistence / 130000030330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022220013300321-2233033011102203-1121311233232132-3023230113232221-0133321012032033-0113233033033000-3200021033100110-3221132220001313"></a>

## Next pages — disable_lb_source_ip_persistence / 130000030330 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0303111031121330-3030232011000130-3000031202321210-3313200311012320-0233221123302133-3200321002133322-3222232213001122-3301202011203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133102302023300-1202321300302312-0301312023110021-0312320033201221-2112201300213303-0032122303212021-1130013030222231-0103312102201032"></a>

## advanced_options.disable_outlier_detection — disable_outlier_detection / 301202132012 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.disable_outlier_detection

<a id="canonical-2212201332103112-2302001020103013-1132313320003233-0300213312212312-0133202221110012-2130320122232013-0312202013323213-3002323223233101"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
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

Terraform syntax:

```terraform
disable_outlier_detection = {}
```

<a id="canonical-3002230020032213-3101013121001001-0110010311002230-2131210322101102-0320020013311001-3112002222020333-0212110223002331-3312332212233330"></a>

## Direct properties — disable_outlier_detection / 301202132012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032000110123123-0322022120210103-0220220313210231-3112203202010313-0213012213303312-1322100000331230-2320332302213122-1111100203101020"></a>

## Next pages — disable_outlier_detection / 301202132012 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2213300033233331-3203002012201223-1020300032131122-3110201221031120-1233033311101332-1000313131023221-2131220002000133-2000301223322131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202320332332233-3012032331000012-3310023101030030-2010313201213223-0210001203120303-0201112101123030-1302110210132000-3231130200022333"></a>

## advanced_options.disable_proxy_protocol — disable_proxy_protocol / 131333333132 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.disable_proxy_protocol

<a id="canonical-1223213201110302-3203023132210132-3321002121300031-3122203113302013-2030112122202131-3202003103330011-0330311133032133-1010331221030333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable proxy protocol.

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
disable_proxy_protocol = {}
```

<a id="canonical-2113133030032302-1221313002123132-1001212323313230-1012112121100032-2031312320113112-3201221313303121-1222220301100223-2310112032233130"></a>

## Direct properties — disable_proxy_protocol / 131333333132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313131323021023-2113203322110313-2013012313201111-3021323201013011-1031001233113032-0333223023003202-1120223121333201-2020113002123003"></a>

## Next pages — disable_proxy_protocol / 131333333132 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0132102210122311-3313023112101111-2013130010001303-0031232103110330-3010201102310202-2032033320210210-0303113213130210-3301233212211230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022321103133121-0113202020011222-2001110132121210-3031212031312120-0311323320212233-3123311220013023-1300211231300031-3121030332110010"></a>

## advanced_options.disable_subsets — disable_subsets / 103233101200 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.disable_subsets

<a id="canonical-2221002121003202-0110323122320323-0312031200322111-0133202102021233-1201000233003103-3232111033000131-2203132332312322-1021312232200213"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
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
disable_subsets = {}
```

<a id="canonical-3002323200212300-1111010301110310-1103221131023101-3222030131201131-1303300000310300-0202233300331010-2232001301103101-3022022032110000"></a>

## Direct properties — disable_subsets / 103233101200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300030002120003-3030212013111220-2232122203332101-3213100011012001-2021032131203221-1213022200310222-0123201313220230-1133111110312003"></a>

## Next pages — disable_subsets / 103233101200 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0003003221233012-1310223021303120-2211013133202102-3001322232012120-0123301231303010-2022223113222201-3311201000011110-0230022321023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313231112232012-0131312023233310-1110033011000303-1021011331320122-0010132320323321-0113031033123003-1102013203011133-1132130000122203"></a>

## advanced_options.enable_lb_source_ip_persistence — enable_lb_source_ip_persistence / 000032230113 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.enable_lb_source_ip_persistence

<a id="canonical-3300113001032211-0103032032303033-2311013312203333-3000322131122030-2023130130303202-2121231202130313-3122011231212221-3222332330203013"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
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
enable_lb_source_ip_persistence = {}
```

<a id="canonical-0120111312100301-1331011213303203-2101200001021021-1003332312233030-2132132302302103-2033321013011312-0201200331313320-2321002021233123"></a>

## Direct properties — enable_lb_source_ip_persistence / 000032230113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312211210201102-0310133131322033-2013322101223311-3302000032200033-2001303100103101-1213310311101222-2312002323232013-1213003000133023"></a>

## Next pages — enable_lb_source_ip_persistence / 000032230113 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231023301311002-3132232203300213-0011213102202230-1000113323310121-1100033031221230-0031300000003220-0231100231202031-0003030202021003"></a>

## advanced_options.enable_subsets — enable_subsets / 012133231002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.enable_subsets

<a id="canonical-0130213121302133-2302311313233002-3133011200110221-2312031121330231-2021233031010111-0330102012311323-2120302100122000-0023222310232122"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint_subsets"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "default_subset"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "fail_request"),
  validators.ConflictingObjectAttributes("default_subset",
    "fail_request")}
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
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

Terraform syntax:

```terraform
enable_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302312333130320-3000003102013212-3030032302133110-3331012133301102-2300010210222010-2321231012100130-3301222220223300-1110213201312231"></a>

## Direct properties — enable_subsets / 012133231002 / 3

- [any_endpoint](resources--origin_pool--reference--group-001.md#canonical-1121303203301233-0330131022131321-2331213130301012-3101333030030020-3301212310113020-3023213033120033-3011120312103120-0222120021211220): complete subsection reference.

- [default_subset](resources--origin_pool--reference--group-001.md#canonical-1321103031213013-0311223021230033-3212121130120222-0002212110223121-3110121312032001-3131302223131121-0332000321202300-3020010030013333): complete subsection reference.

- [endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-3333211223132133-1132232010121131-3323222002121213-3033230302120322-2231132313021333-3130333231100130-3002003220033303-1223300032102121): complete subsection reference.

- [fail_request](resources--origin_pool--reference--group-001.md#canonical-2220010100211223-2000310001213033-0103311300112330-2311301203323102-2032301030100221-2120123031202322-1202102232221010-3123322123111330): complete subsection reference.

<a id="canonical-2113230113231213-1302230000220032-1120102100030322-0330210300203310-3313022330322000-3003002330303033-2011112300102333-2331223002023230"></a>

## Next pages — enable_subsets / 012133231002 / 4

- [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--reference--group-001.md#canonical-1121303203301233-0330131022131321-2331213130301012-3101333030030020-3301212310113020-3023213033120033-3011120312103120-0222120021211220)
- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-1321103031213013-0311223021230033-3212121130120222-0002212110223121-3110121312032001-3131302223131121-0332000321202300-3020010030013333)
- [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-3333211223132133-1132232010121131-3323222002121213-3033230302120322-2231132313021333-3130333231100130-3002003220033303-1223300032102121)
- [advanced_options.enable_subsets.fail_request](resources--origin_pool--reference--group-001.md#canonical-2220010100211223-2000310001213033-0103311300112330-2311301203323102-2032301030100221-2120123031202322-1202102232221010-3123322123111330)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1121303203301233-0330131022131321-2331213130301012-3101333030030020-3301212310113020-3023213033120033-3011120312103120-0222120021211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222133223332333-3001031210331111-0303322033202011-0031322300321301-0033232221333120-0033301310332313-2233213200123201-1013332300303032"></a>

## advanced_options.enable_subsets.any_endpoint — any_endpoint / 103210120223 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- advanced_options.enable_subsets.any_endpoint

<a id="canonical-0112103101202323-3021313231233213-0220331031313320-0102102330103002-0023313023323213-3020031113032313-2020033212111321-3210123312201201"></a>

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
any_endpoint = {}
```

<a id="canonical-3111020303223231-2210002130003312-1312203323332311-0303323203220303-0023220023232322-3011300113102022-3103233320112030-0023212333223301"></a>

## Direct properties — any_endpoint / 103210120223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203103212120012-2201111310221201-1002013331213131-0322103301221321-1233212121100223-0102212112313022-2320333003123100-2300031000101010"></a>

## Next pages — any_endpoint / 103210120223 / 4

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1321103031213013-0311223021230033-3212121130120222-0002212110223121-3110121312032001-3131302223131121-0332000321202300-3020010030013333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001001120331102-3102322122103232-1300002200202322-2200300103303220-3330113010001131-0103032000110131-0110201123213111-0022301222123001"></a>

## advanced_options.enable_subsets.default_subset — default_subset / 100002200311 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- advanced_options.enable_subsets.default_subset

<a id="canonical-3121100310212312-3312122233302321-1001200330030022-3232330223331100-0232210012102020-0202331311323113-1121113132231221-3202010300311023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

Receipt-pinned upstream constraints:

```json
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
default_subset {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111231120331203-0302123233021300-2100312300130001-0330233001023332-0103110102201131-0232321202300331-0221020301102323-0232001102213023"></a>

## Direct properties — default_subset / 100002200311 / 3

- [default_subset](resources--origin_pool--reference--group-001.md#canonical-3201201213002301-3203310321121111-0111211221202232-0312321103313130-1222132231232200-2202002113302312-2033010222031010-1232011321332002): complete subsection reference.

<a id="canonical-3200302212300230-2213223100023101-2113002321231322-3321321323213102-1100320031012200-1211123101323313-1301010330112003-0330023330130102"></a>

## Next pages — default_subset / 100002200311 / 4

- [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--reference--group-001.md#canonical-3201201213002301-3203310321121111-0111211221202232-0312321103313130-1222132231232200-2202002113302312-2033010222031010-1232011321332002)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3201201213002301-3203310321121111-0111211221202232-0312321103313130-1222132231232200-2202002113302312-2033010222031010-1232011321332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103220100120222-0332302333222023-3312310331223022-2031131022122001-1123012000012332-2212320323320102-1202100312121112-1203212002021021"></a>

## advanced_options.enable_subsets.default_subset.default_subset — default_subset / 102033331110 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-1321103031213013-0311223021230033-3212121130120222-0002212110223121-3110121312032001-3131302223131121-0332000321202300-3020010030013333)
- advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-0100213033301133-0233012321333201-1223031300021022-1212133120202303-0222302300231032-0110022032212202-1301320001121312-2020221011202023"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 32,
    "metadata": {
      "confidence": 0.75,
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

<a id="canonical-3001212303303302-2113002331122023-0131230330212013-1101303310212231-1101110132232002-2130103022023211-2112322020131302-0001321233131311"></a>

## Direct properties — default_subset / 102033331110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310111000103103-1331010232212223-1332231131022332-1000313323002302-0220301121200220-1112230300231210-0023022200102132-1302112130322130"></a>

## Next pages — default_subset / 102033331110 / 4

- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-1321103031213013-0311223021230033-3212121130120222-0002212110223121-3110121312032001-3131302223131121-0332000321202300-3020010030013333)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3333211223132133-1132232010121131-3323222002121213-3033230302120322-2231132313021333-3130333231100130-3002003220033303-1223300032102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013233132300131-3002333333220032-3032330113300232-1002332133330230-1132200003003221-2303130001110133-1313220113202132-3112111211133121"></a>

## advanced_options.enable_subsets.endpoint_subsets — endpoint_subsets / 302113312001 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-3003103130030232-0230000031012213-1301031101001233-2210110120010023-2223331030120211-1121123133300123-0323221230011023-2121012000202312"></a>

Type: `"object"`. list nested block, Optional.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

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
endpoint_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212102001320211-3233110130332113-3020220130101120-1301133111331203-1022232030333000-0311221003333320-3231300011111110-1122220000013021"></a>

## Direct properties — endpoint_subsets / 302113312001 / 3

<a id="canonical-1201301112230012-1311311020221001-2332033101303231-0303102010020213-3112203310100332-1303100232123233-2232312302213233-3322303330232013"></a>

<a id="canonical-3123102220000200-1031021232200330-1020012303130030-3000133223312102-1223221203203013-2000313301132022-2311000012330320-3300033123021231"></a>

## keys property — endpoint_subsets / 302113312001 / 4

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

<a id="canonical-1020132310311010-1000333202202333-2230010133210012-1012303333303323-0330332002111102-2101001310131320-3301233310320222-2130131311333010"></a>

## Next pages — endpoint_subsets / 302113312001 / 5

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2220010100211223-2000310001213033-0103311300112330-2311301203323102-2032301030100221-2120123031202322-1202102232221010-3123322123111330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001320013031130-2120111302211320-0332311233303023-2022232132010113-1213100101332232-0132313333000220-1001121332033020-2203213323003323"></a>

## advanced_options.enable_subsets.fail_request — fail_request / 203203302001 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- advanced_options.enable_subsets.fail_request

<a id="canonical-3121222233001100-3031133233230220-0323303312010233-0010131132031320-3033222030310130-2123102131212232-3020002133201121-2203110330222211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fail request.

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
fail_request = {}
```

<a id="canonical-1232302213110012-2023200220220311-0301111120023220-2321001212122303-3330022113210020-1131123102131030-2313032112100122-1201303131330003"></a>

## Direct properties — fail_request / 203203302001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003011330300211-2230031333313031-3311122122231120-1003200221121230-2203130100232313-1133000132231222-2313233213121123-3120200320131302"></a>

## Next pages — fail_request / 203203302001 / 4

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200021312132020-2120023201322021-2102111232211311-1211222210121232-2200301332000113-3301132222231233-0212022030002233-3132111313110313"></a>

## advanced_options.http1_config — http1_config / 331011130310 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.http1_config

<a id="canonical-2132020013013033-2323101233311313-3211303313122120-3102302022100333-0022202022113202-2012032011321231-2302330103122302-0022032131123222"></a>

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

<a id="canonical-1031000130103011-1231112200101322-2132233323031330-0102102312113323-2221033022021313-3013331323131110-0330120001210030-2201230012310003"></a>

## Direct properties — http1_config / 331011130310 / 3

- [header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113): complete subsection reference.

<a id="canonical-2030112203320231-2310222032010333-2303000322102320-1323210113011023-1221213002210121-3230100323301113-3331102310321231-3311013310213002"></a>

## Next pages — http1_config / 331011130310 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323302130110001-0223233313200101-2310111110331013-0023002030022300-1032302222101113-0222103000310300-0230122121033312-0003133221113002"></a>

## advanced_options.http1_config.header_transformation — header_transformation / 033201032323 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- advanced_options.http1_config.header_transformation

<a id="canonical-3121033213333332-3112123233111100-2021210313333332-2212300333203100-3020011233300311-0001013211303123-3113000132232123-0011023321303101"></a>

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

<a id="canonical-2223300332330020-2313002131203231-3113220132331331-2020210020013132-1022000101320312-2032011312212233-0003323100113211-2031023333212022"></a>

## Direct properties — header_transformation / 033201032323 / 3

- [default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0000101230130110-2020003133003030-3312202133003231-0031122321123201-1010011102103130-3103201233223100-1023100311130223-0003000120213322): complete subsection reference.

- [preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0330021003111221-1332122300203310-0123000012003123-0223121210122230-3231211101203232-2312110300330112-2123001120301023-3030102300123220): complete subsection reference.

- [proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-2321111301112201-3001110011210122-0311313321332031-2210310131221130-3123323313022111-1231012212212200-2322313302312102-2231011113000111): complete subsection reference.

<a id="canonical-2010222121022301-1122323333112132-0213310002000011-2200332031112300-2031202120130030-3013330210320223-3320133223011322-0121200032030320"></a>

## Next pages — header_transformation / 033201032323 / 4

- [advanced_options.http1_config.header_transformation.default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0000101230130110-2020003133003030-3312202133003231-0031122321123201-1010011102103130-3103201233223100-1023100311130223-0003000120213322)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0330021003111221-1332122300203310-0123000012003123-0223121210122230-3231211101203232-2312110300330112-2123001120301023-3030102300123220)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-2321111301112201-3001110011210122-0311313321332031-2210310131221130-3123323313022111-1231012212212200-2322313302312102-2231011113000111)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0000101230130110-2020003133003030-3312202133003231-0031122321123201-1010011102103130-3103201233223100-1023100311130223-0003000120213322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000323031232212-0302300032310000-1100030100230222-3021131230313221-1031201111320032-3200031330032213-2032321113322012-0112220021321302"></a>

## advanced_options.http1_config.header_transformation.default_header_transformation — default_header_transformation / 202211022323 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-1030313301202002-3313010133302201-2212221203301002-2312110331010101-2110331020102220-2033101321130103-2201030003333002-2132233011332002"></a>

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

<a id="canonical-2031302000003332-0201301110231321-3023332201000221-3201010313202112-0131232223202201-2030020233012120-0330113331200132-0021012321101123"></a>

## Direct properties — default_header_transformation / 202211022323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032231321101222-3101010311132202-2020320001211023-3100210022311133-3002100113012102-2203323331213003-2100201323231223-0201300133313022"></a>

## Next pages — default_header_transformation / 202211022323 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0330021003111221-1332122300203310-0123000012003123-0223121210122230-3231211101203232-2312110300330112-2123001120301023-3030102300123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332000302213002-0001033000311011-2213023012101032-2103003220203031-2102221110103313-1322132000002101-1030133120021232-0123310231013123"></a>

## advanced_options.http1_config.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 103223202302 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-2113021231033330-1222230130333113-3232331122330003-0311100210231011-1222100200111020-3233302113013123-3200120232200310-3000231131301200"></a>

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

<a id="canonical-0133320301311210-3330222332032211-3101102223121032-1031020201121113-2202200210200020-3201113022233132-0120232203232331-0100210131030023"></a>

## Direct properties — preserve_case_header_transformation / 103223202302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010012013223032-2101011311103313-3323032032001211-3310210301123333-3132023202302101-3231112113120322-0130302220333323-3300322021033201"></a>

## Next pages — preserve_case_header_transformation / 103223202302 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2321111301112201-3001110011210122-0311313321332031-2210310131221130-3123323313022111-1231012212212200-2322313302312102-2231011113000111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022000333011123-3103223313231333-1122201101220331-2301323132131300-2100330323322102-1221101310020210-2310101221232331-2110312021230321"></a>

## advanced_options.http1_config.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 121032303011 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-1120132330211033-0012321110310010-0011021112300002-3030000212012022-2333301001233332-2022230020331222-0113112211030310-1322332201011000)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-0203033301120032-1302311331200121-0133313200231021-0131312332232300-3112102211000320-0310010132003312-0320321132020112-3123032130231213"></a>

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

<a id="canonical-1130300302310321-0211013232032231-1331223000022003-0021102223103113-0232032323010000-3000101101120002-2113233121000221-2133211123230112"></a>

## Direct properties — proper_case_header_transformation / 121032303011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013130112200300-1221013313013133-0301232211330231-0223123200030230-0200231002301102-2131120212220210-0110233102020112-1021300023313303"></a>

## Next pages — proper_case_header_transformation / 121032303011 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3110221330232310-0033030320300111-0032333121300122-2203022202112131-2002121100333022-3130111233223112-1030012232320323-0122011120320120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310123202210030-2123202123330123-2320013211231100-2321201303313102-2310122112113212-3200032320121002-1012322322310231-1301303022311020"></a>

## advanced_options.http2_options — http2_options / 013303111322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.http2_options

<a id="canonical-2223312110301032-2021330112213312-3032301011213023-1310031223110232-2021203303111330-0012332111231101-0302122133111103-2002002000022301"></a>

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

<a id="canonical-0122130313013111-0130131001021011-0021022022131021-0011232101102210-3200003320312233-1132101031023012-2210331203030210-0020123022111233"></a>

## Direct properties — http2_options / 013303111322 / 3

<a id="canonical-0001032332222230-1232102121000202-0223132112001333-0301133131132123-3302031331201113-1212301110000123-0100310303122321-0220022202310212"></a>

<a id="canonical-3031103330122023-0321313312033223-2111320222102322-3121331021031132-2133222210101232-3301120220123221-0131233002301222-0220302111332231"></a>

## enabled property — http2_options / 013303111322 / 4

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

<a id="canonical-0133312022323132-1303012311310020-3302102201222102-0121110222031101-1003131120301322-2101232010222300-2111203233102300-2012321132122132"></a>

## Next pages — http2_options / 013303111322 / 5

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0202112120201020-0312010230000131-1100030322023231-3232212311110032-1220210322222103-2332221300001302-3110032302100201-0210013330221330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101013121020223-2023322100233111-0220232302231123-0212330233121221-3232211013200201-1022230310330013-2203100120012000-3113311212132020"></a>

## advanced_options.no_panic_threshold — no_panic_threshold / 221122213233 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.no_panic_threshold

<a id="canonical-3022211111030303-3303210112231032-3021322222331213-0112210010001200-3301230202232230-1333221121031111-3310020003001123-3221033220022021"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no panic threshold. Defaults to \`map\[\]\`. Server applies default when
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
no_panic_threshold = {}
```

<a id="canonical-2110201231132032-2321033221110121-3120121202133022-1322232031310101-2111002101223313-2011033322131210-3000232301332020-3333220203310303"></a>

## Direct properties — no_panic_threshold / 221122213233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233301301211021-3301233222212331-2331031302230220-2110223011312320-2021210233311333-3130021130201100-0200330310210322-2031222212103331"></a>

## Next pages — no_panic_threshold / 221122213233 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1330213333313110-3100031301022021-2123130200313332-3201013222020131-2223030121213030-2220120203131331-1221101000313013-1112103003000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012323113220111-1011123303212031-2113131002102233-2013230133320311-0200211011332333-0200211313203301-1123020221001021-0113032103030310"></a>

## advanced_options.no_request_limit_per_connection — no_request_limit_per_connection / 233021300221 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.no_request_limit_per_connection

<a id="canonical-0333321010033130-3320313302301220-2122011100131333-3213123332220331-2232001112011220-1120330033123332-0302232303012031-1233011330211012"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
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

Terraform syntax:

```terraform
no_request_limit_per_connection = {}
```

<a id="canonical-3113210223213310-3213301212213112-2233220120203102-2133201022030211-2332223002022130-3222220012313212-3201332001120120-2100012223121012"></a>

## Direct properties — no_request_limit_per_connection / 233021300221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003123122212112-1031233023313231-2100300001233102-2103120023200212-0230013002031132-1213321032030020-3330013222221103-2211101220020201"></a>

## Next pages — no_request_limit_per_connection / 233021300221 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3000232203110021-1031230213122011-1123333023221112-2112132033221221-0233111301322212-1122212030103130-0001131010111133-3133220123103122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213322200021232-0202233312023023-0211330232102210-2231101212302201-0130203020211313-1113302101320023-2123002303011111-0313333320013100"></a>

## advanced_options.outlier_detection — outlier_detection / 323112202222 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.outlier_detection

<a id="canonical-3303000202220013-3031232033223310-1321323003321201-1010000202210313-1202211123013022-0303123122302011-3300002131010223-1100212211121220"></a>

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

<a id="canonical-1021123001210121-3200232001111332-3033123123232012-1021222220001333-2000102033222110-0310323212100003-3332002202211013-3231213212323331"></a>

## Direct properties — outlier_detection / 323112202222 / 3

<a id="canonical-3222013223202212-3021331003111312-1122020202020233-1223310211021313-3333132332131022-1111030030331122-2313200201332021-3321323232013330"></a>

<a id="canonical-0023301003023303-0023130030323101-0113321302000303-2110130213221032-1202030133212133-0323212311122102-3222132313211120-2310122221011333"></a>

## base_ejection_time property — outlier_detection / 323112202222 / 4

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

<a id="canonical-2203321212323023-1033222332001110-2212310130113320-3020323000203022-3310323033013031-0302132223110303-2313321311313323-0103322133211113"></a>

<a id="canonical-3131120131100011-2321310221133201-1000112010211201-1302130132322031-2311020301130000-1221212013031313-2323301210000310-3302011210032332"></a>

## consecutive_5xx property — outlier_detection / 323112202222 / 5

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

<a id="canonical-2032210130223023-2102101312113131-3032201012121110-2133300213313112-3023001221202213-3013300002310031-0133323100001111-3323313111103212"></a>

<a id="canonical-0020031130012321-3211011020300100-3010313021302211-0212132113203113-3311003113021013-3212112333032213-1230011130011213-0133331023130223"></a>

## consecutive_gateway_failure property — outlier_detection / 323112202222 / 6

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

<a id="canonical-2130211122313303-3100312301022030-2113233133123103-2321321302223121-3330121120122021-2031222311233021-0213233320112322-1021001232101223"></a>

<a id="canonical-3033302130002021-0233132323333033-0200310333210310-3301132112131033-3203113100333203-2132033122100121-3021330200320133-1221122312311013"></a>

## interval property — outlier_detection / 323112202222 / 7

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

<a id="canonical-2122121332233031-1110123232210002-3012111131220131-0011000203331013-2302333122222202-3112322331313021-2310221000331300-1210301323303112"></a>

<a id="canonical-2103121221033211-1003203302133312-0321122121231231-0030102133330013-1032301301230220-3331223231021132-2211210203122322-0001022330112130"></a>

## max_ejection_percent property — outlier_detection / 323112202222 / 8

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

<a id="canonical-0330033113001222-1330303110012010-2221123123230002-2311020233121102-2023002112330101-0231221311103100-2220322213331112-3102103013011213"></a>

## Next pages — outlier_detection / 323112202222 / 9

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2231302013311332-2302020300210101-3013021032120300-1322020323111001-1021312211133123-0001210223111212-2210221303330123-2132302330031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203302022310000-3210331212331310-3331320022030103-2100132200321200-1120302130202203-2003201113230322-2133320301100232-2030322300123312"></a>

## advanced_options.proxy_protocol_v1 — proxy_protocol_v1 / 232123331231 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.proxy_protocol_v1

<a id="canonical-1112010023001330-1123320210030303-1111013103121033-2333113200220132-2321321103032213-0230230010023212-1222230223232233-1302030120202033"></a>

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

<a id="canonical-0030010310000202-3322021123213222-1313131030212212-2301333322031001-2132122302102321-2300132030132030-2022200323302012-2322021020110333"></a>

## Direct properties — proxy_protocol_v1 / 232123331231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320323011223310-1131032221123102-0233001331103320-2323302303333000-0212212103321321-0320310102210223-1010020003113311-3101333201012332"></a>

## Next pages — proxy_protocol_v1 / 232123331231 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0320031013302220-3332100213221311-2123202301102332-2010033001202221-1032332013201311-3211112203311011-3001213113311213-2320133302110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231332023001210-1203122233031333-1101012203122021-0332230102121202-0013031102131322-1201213023310222-0322001233103131-1201322130203113"></a>

## advanced_options.proxy_protocol_v2 — proxy_protocol_v2 / 322202302330 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- advanced_options.proxy_protocol_v2

<a id="canonical-1130332102113212-0021332032233120-2200102202030301-0030013222102321-3123301113201332-0232222011303302-2100100222311301-1003221222301131"></a>

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

<a id="canonical-0022110023202032-0023101102213313-3332013010232233-1312333230233022-2020221220123022-3003332221101000-2132222233231301-3132123010113332"></a>

## Direct properties — proxy_protocol_v2 / 322202302330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132202221032210-1320330330300133-1320113033232321-0221222300223230-0001230023132301-2212110121300322-0103031213203201-3010001130122001"></a>

## Next pages — proxy_protocol_v2 / 322202302330 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1112022203223032-3323221233013132-1121203303031132-1021300130303111-2312333020233311-1132310003313023-1330012320112102-3213232312123011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111212320211030-0131021000010120-2022031200132323-3230302332221132-3131122312111231-1012233303332223-1132303310022010-0110231002011130"></a>

## automatic_port — automatic_port / 032123112032 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- automatic_port

<a id="canonical-1303022022111110-2200312312101231-1001130323122311-0202303312301223-0312030312212121-1001200100312002-2111111312223331-3333331320211201"></a>

Type: `["object", {}]`. Optional.

\[OneOf: automatic\_port, lb\_port, port\] Enable this option

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

- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-1303022022111110-2200312312101231-1001130323122311-0202303312301223-0312030312212121-1001200100312002-2111111312223331-3333331320211201)
- [lb_port](resources--origin_pool--reference--group-001.md#canonical-1203333113211100-0000103313212022-1300332310021033-3013101120332310-1310310323010220-0222223000033030-0110200123301112-0211322003111020)
- [port](resources--origin_pool--reference--group-001.md#canonical-2332110001010220-3311320331201032-1221232122321010-0331301231310313-0122021030131110-0201203221123012-1211002121112313-2231113213132103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
automatic_port = {}
```

<a id="canonical-3032331103022311-0223232333120231-0301120021213022-0322122030020131-3210010111330023-0100333113222313-0313233032021120-3132333331303323"></a>

## Direct properties — automatic_port / 032123112032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131320221311321-2033231320112233-1221122232223031-3231212302023200-2120003211313100-3111002322132013-0200020010030231-1200322013011110"></a>

## Next pages — automatic_port / 032123112032 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0220303133013032-2302323323212323-0311200001213123-3021130113330311-2212330212003003-3231310310300112-2222012213121021-0030322010120013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303101031021201-3101111030211230-0221313100303103-3123121313203210-0121113000321333-2112032323233021-1130211020200320-0210003202203131"></a>

## healthcheck — healthcheck / 313023132131 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- healthcheck

<a id="canonical-3311000013100201-2012002213332223-1010102011213333-0021131001203021-2122012002233132-2211211221330320-3003300301311301-1333121120220323"></a>

Type: `"object"`. list nested block, Optional.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Reference to healthcheck configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
healthcheck {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000322323032121-2202230222030310-2321013100331210-2013012011030230-1331220033201020-2201123201322333-3021102111101300-0020232313330133"></a>

## Direct properties — healthcheck / 313023132131 / 3

<a id="canonical-3102130011331331-1210122203312100-2200033302232230-1020113322120222-0131301331003312-3211132330031110-0233103330331213-1311202013012001"></a>

<a id="canonical-3030212300323003-0020230021101300-2303031100101302-2011321301200030-2213123323233312-3201321120020312-0122222333111323-3313122033322100"></a>

## name property — healthcheck / 313023132131 / 4

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

<a id="canonical-3232130210112223-0331032011003301-3133012330113111-3210012013333021-3120100223320020-2131213033021323-2320003311013231-1211011203221012"></a>

<a id="canonical-2202312023110223-2211031132203010-2022003111201312-2312013021132323-3020023213020021-0313223332002013-2002312020001333-3200330222122211"></a>

## namespace property — healthcheck / 313023132131 / 5

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

<a id="canonical-1103333033322212-0313232301030200-3310102110331110-0320220033213023-1311033110223013-2210010333111213-0131132232331103-3301100032030222"></a>

<a id="canonical-2001232231012130-2132122000021230-0300202101322233-1333100120111201-0312012132331311-1330223020211003-0322310110002130-1312010311333303"></a>

## tenant property — healthcheck / 313023132131 / 6

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

<a id="canonical-0110310010111233-3122000131023212-3302303021102313-1103113311222313-2030012213110213-0110332022211011-3331310112030012-0031021121203023"></a>

## Next pages — healthcheck / 313023132131 / 7

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0020300233212120-1303102312023111-0212101131203331-1302022110112113-0033130113202023-0130301212232010-1022111220323023-0020230101001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220333331131023-3011322032333021-2133211332300210-0032113120302223-0003102222022232-1023212203302110-2133112023221310-3302322113102213"></a>

## lb_port — lb_port / 120330313213 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- lb_port

<a id="canonical-1203333113211100-0000103313212022-1300332310021033-3013101120332310-1310310323010220-0222223000033030-0110200123301112-0211322003111020"></a>

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
lb_port = {}
```

<a id="canonical-2332313110212003-0201003030111323-0001012310032100-0233321022330302-1032312223001320-0333011021201122-2322030021321310-2320111223213001"></a>

## Direct properties — lb_port / 120330313213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010223311320313-3120003331213330-0200320013333013-0101223033000000-2122333033122331-2231130112320131-3110201002102200-0201233022211200"></a>

## Next pages — lb_port / 120330313213 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3232201211033121-2320120332221003-2102323230100223-1300312333300223-3332330230003310-0213003320030302-3013211011131001-3021000212231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222122131301011-0110321110031101-3223111303132332-0000113212101121-2312100222311130-2022221031030111-0221200322203332-1022022010332100"></a>

## no_tls — no_tls / 031031300102 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- no_tls

<a id="canonical-0031212003112103-0032021023323003-0332121111010323-1200301212121200-2113222033331332-1102230131001230-0012333220100001-2103312132322120"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: no\_tls, use\_tls; Default: no\_tls\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [no_tls](resources--origin_pool--reference--group-001.md#canonical-0031212003112103-0032021023323003-0332121111010323-1200301212121200-2113222033331332-1102230131001230-0012333220100001-2103312132322120)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-3122122211032132-3132313203100001-1013000203202031-2132002223311123-2211000023310210-0233300322102210-0033100013333133-1022321012001012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_tls = {}
```

<a id="canonical-1323010210203200-2233112303110022-1133002012033221-2311132100321032-3011131132303102-3002312101011110-0020023003112003-0322000201233033"></a>

## Direct properties — no_tls / 031031300102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202333211201303-3202122102012303-3030101312131010-1003321230130110-1233201312220120-3312201303123303-2230201310332232-0100202331102123"></a>

## Next pages — no_tls / 031031300102 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333302111121330-0300331102000302-2233022110030232-2321023311233300-0230001111110102-3132230213201123-0212301213110301-3333001322120002"></a>

## origin_servers — origin_servers / 330013001011 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- origin_servers

<a id="canonical-2001322203132131-3022210213212113-3312012003100000-0030021333032022-0213330213213133-3303011303133031-2200012302331213-1200003120130011"></a>

Type: `"object"`. list nested block, Optional.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cbip_service",
    "consul_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("consul_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("vn_private_ip",
    "vn_private_name")}
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
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012201102320023-2312313330331323-3020011130303331-2020132220012222-2211113012230033-0112022203102102-3112122330322303-0000102330300130"></a>

## Direct properties — origin_servers / 330013001011 / 3

- [cbip_service](resources--origin_pool--reference--group-002.md#canonical-1010331102312323-2231313310123113-1310231000100331-3110100312012302-3331111003312110-2102012030220303-1220320003131210-3032300320110001): complete subsection reference.

- [consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331): complete subsection reference.

- [custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-1013001331311011-2021321010021121-0102302220012123-3123220033111332-2222001121311212-2212111320022231-0232110210110023-2030010232123321): complete subsection reference.

- [k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300): complete subsection reference.

<a id="canonical-0222322311301012-0302213113300121-1303133220100113-2310201011020013-1133022301330223-1211000303231312-2232032222303300-2231312012013323"></a>

<a id="canonical-0303320322232021-3230331021221230-3233300121123203-0323131322210333-1211213200013022-3121333212320111-1131212232300133-2132123230303020"></a>

## labels property — origin_servers / 330013001011 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this origin server, these labels can be used to form subset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003): complete subsection reference.

- [private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323): complete subsection reference.

- [public_ip](resources--origin_pool--reference--group-002.md#canonical-2012202102220023-2022233311031000-2131130202320231-1200001300020311-2032322312323113-2110322301123032-3132132030003331-0200031213133012): complete subsection reference.

- [public_name](resources--origin_pool--reference--group-002.md#canonical-0101123032132221-2211032101001201-1000311221111203-2133111213011322-0333020102301313-3100023301322230-0212222321323322-2203130310321323): complete subsection reference.

- [vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313): complete subsection reference.

- [vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112): complete subsection reference.
