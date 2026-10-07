---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2212021011102011-1003220310203030-1203200220232232-2001330231011113-2303212012031031-2101100012001130-2122233232002202-1012110130311313"></a>

## `ddos_mitigation_rules.metadata.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_pool

<a id="canonical-1213330012310303-3223033000020112-1132322020000213-3203211100210032-3203202100010003-0323120100333103-3310310021000230-2332331033130013"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: default\_pool, default\_pool\_list; Default: default\_pool\] Configuration parameter for
default pool.

Additional upstream details:

Shape of the origin pool specification.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("automatic_port",
    "lb_port"),
  validators.ConflictingObjectAttributes("automatic_port",
    "port"),
  validators.ConflictingObjectAttributes("health_check_port",
    "same_as_endpoint_port"),
  validators.ConflictingObjectAttributes("lb_port",
    "port"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-health_check_port_choice": "[\"health_check_port\",\"same_as_endpoint_port\"]",
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

OneOf alternatives in this subsection:

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-1213330012310303-3223033000020112-1132322020000213-3203211100210032-3203202100010003-0323120100333103-3310310021000230-2332331033130013)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0130032320003112-3013231322132011-2010113111012220-3312013033110232-0201300301333010-2200230132012313-2122033102033333-3101112132302330)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313311312330121-3100113323301313-2012102320010002-0321311210303111-2210332210131331-0312212202013111-0031132222113121-1000130220311302"></a>

### Direct properties for `default_pool`

- [advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100): complete subsection reference.

- [automatic_port](resources--http_loadbalancer--reference--group-015.md#canonical-1322002312102330-3201331032321312-3300300133011331-3120132210110301-3210102232133203-0002132201013013-2023203103322300-1312001000200312): complete subsection reference.

<a id="canonical-1220013021102311-1210300212021323-3030313021111112-2331311200001322-1301030222003002-2332032220130322-3233323111021113-2321012103002111"></a>

<a id="canonical-1200033130233022-2111031030233113-2033123120032200-0301320231120233-0010303201100120-0331202202132202-3010221022100131-1023133002221332"></a>

#### `default_pool.endpoint_selection` property

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Additional upstream details:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DISTRIBUTED","LOCAL_ONLY","LOCAL_PREFERRED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-2333031322131122-1333231311010031-0023200220303113-2223032011333030-1203001000013012-0333300323010101-2103102303202333-2212032221013220"></a>

<a id="canonical-0122133123300312-3202000120222003-1030023200201002-2233203202220001-1112212130113013-2333100113022003-0321013222222331-0001031003002212"></a>

#### `default_pool.health_check_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [healthcheck](resources--http_loadbalancer--reference--group-015.md#canonical-1221110002132303-0012322203331031-0200320230330313-0111101330333233-2122220212333113-0311121023131102-2021213213213203-2231002020013210): complete subsection reference.

- [lb_port](resources--http_loadbalancer--reference--group-015.md#canonical-1023332221023002-1120001320231221-2230003000223121-1203332032211030-3112300112131312-2203210121000020-1110310100200211-0012031020100313): complete subsection reference.

<a id="canonical-2001303000020210-3130102121230130-1031001003223123-1300321230302002-0023213122001213-2022000210011132-0033113233123233-3202100031231222"></a>

<a id="canonical-1030000100233130-1011130013132331-1301121230320021-3331112313131312-0320130023233230-1301222002330310-2133313332132123-1302031122033231"></a>

#### `default_pool.loadbalancer_algorithm` property

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LB_OVERRIDE","LEAST_REQUEST","RANDOM","RING_HASH","ROUND_ROBIN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [no_tls](resources--http_loadbalancer--reference--group-015.md#canonical-3222232123033110-0311020222300102-2311031131230111-0311001112220210-0122011130233121-0322223310331112-0222133023233112-0031210201010010): complete subsection reference.

- [origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031): complete subsection reference.

<a id="canonical-1110332222020032-1030112012011212-0120103200222313-2331302310323003-2322213313231001-1300233203310123-0212302001322020-3231203111022120"></a>

<a id="canonical-0202221012133213-3120301033212222-3303130032003223-2121332010112030-2221232112000020-0300321131331000-3131222033311032-1010010221121300"></a>

#### `default_pool.port` property

Type: `"number"`. Optional.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [same_as_endpoint_port](resources--http_loadbalancer--reference--group-016.md#canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230): complete subsection reference.

- [use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022): complete subsection reference.

- [view_internal](resources--http_loadbalancer--reference--group-017.md#canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332): complete subsection reference.

<a id="canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.advanced_options

<a id="canonical-2201323003203213-3213010002212030-3101030122020221-1021231203300201-1221330132003001-2201133313112133-2323222210220222-1133020302333331"></a>

Type: `"object"`. single nested block, Optional.

Configure Advanced OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0203020001311110-2120133110033003-3322002213123033-1121113031003113-1113300011201231-3310331120322203-0210320330010003-2212300222021300"></a>

### Direct properties for `default_pool.advanced_options`

- [auto_http_config](resources--http_loadbalancer--reference--group-015.md#canonical-3112103233111313-3303220101001113-3022220201103300-2310102331122013-0202002302022030-1310031130301100-1231313211312123-1131333023120221): complete subsection reference.

- [circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-1031212032101202-3130331233121230-1022333233112301-0011311210120111-0130221013310320-2201133030023101-1021133320211103-0302331222320211): complete subsection reference.

<a id="canonical-0002230021000320-2212101133003322-0233132212110213-2310101011321021-0311211000003020-1102212120121103-0231202302003033-1101311110300322"></a>

<a id="canonical-0222110112333101-0303201202033301-0233121231323020-0200130130320022-2110130320310210-3110320021221211-2022301312210221-2233223102323113"></a>

#### `default_pool.advanced_options.connection_timeout` property

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [default_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-0220231330001103-1312111311013211-2301111133301230-0123001132310332-0121313113212333-0001211212112332-2201120003102123-2223202111113133): complete subsection reference.

- [disable_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-3030013113131131-3103010310302203-2123231330311322-1231031222331233-1321202223300003-3001203202013202-3000310231301223-3101032101210223): complete subsection reference.

- [disable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-3113100222020202-3133203310100220-0330332303302301-2031303313121332-0323121303222111-3223132221322032-0310310300022112-3030312000202003): complete subsection reference.

- [disable_outlier_detection](resources--http_loadbalancer--reference--group-015.md#canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202): complete subsection reference.

- [disable_proxy_protocol](resources--http_loadbalancer--reference--group-015.md#canonical-1323211300100232-0123211012202031-0232200001130201-1010020100302322-1122010113130203-0321323222203113-0330110330301230-2203311200223113): complete subsection reference.

- [disable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-0023103001332003-0303111121333030-1302202301203023-2300202001223132-3000303201112022-3312000113212101-3222132031232321-1110111201302101): complete subsection reference.

- [enable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-1030023020110210-2012011100303020-3003123221320323-3003020312113033-1230031030231320-0130321233030101-1032013131300003-2201032222030012): complete subsection reference.

- [enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230): complete subsection reference.

- [http1_config](resources--http_loadbalancer--reference--group-015.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000): complete subsection reference.

- [http2_options](resources--http_loadbalancer--reference--group-015.md#canonical-3113102323303220-1222012231233210-3132230022230202-2102202133132231-2231111233120333-0131210111011331-3220111022033123-3110000201103023): complete subsection reference.

<a id="canonical-3101121222220101-1333020313101022-3010101311133101-2010131210230212-0102100210212100-3213233303001020-0111120223101123-1130022001221311"></a>

<a id="canonical-3222110212332003-3202333312232012-2300333121233311-0120300002232323-2032331100220223-3021000100322033-2032221222122032-1211333331233101"></a>

#### `default_pool.advanced_options.http_idle_timeout` property

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Additional upstream details:

Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 5 minutes.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3312313120101212-3310113103331102-2030021330123121-1202121211210203-2100131131303331-3103211310312030-3230031121311332-2032002201300033"></a>

<a id="canonical-1111231000003202-1230310223130123-0223121133222011-2001330003230122-2100321100120230-1031021223221232-2130022013113310-2210032002333121"></a>

#### `default_pool.advanced_options.max_requests_per_connection` property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [no_panic_threshold](resources--http_loadbalancer--reference--group-015.md#canonical-3010321022203332-1133303230302303-2331101002312232-0110221102133000-2100020203120200-3333120023011003-2312121111020320-1223232102333011): complete subsection reference.

- [no_request_limit_per_connection](resources--http_loadbalancer--reference--group-015.md#canonical-3202011012323122-2310202220132112-1203221321333321-3232310230130231-0223011220002020-3131220013013302-1112021212030131-2212302110230313): complete subsection reference.

- [outlier_detection](resources--http_loadbalancer--reference--group-015.md#canonical-3122120002230323-1331302213332300-0221022000101211-0233001322022311-3200121333120233-0120103223222330-0001121011111302-0100313133112132): complete subsection reference.

<a id="canonical-1011102130302100-3311233002102131-1210231013102012-3301233233202330-0331002201332113-2102312333332312-2200130212312232-1311032321222202"></a>

<a id="canonical-0123033020210300-2220132011233120-2013221330103221-1012333212003302-0131310302033212-3211002312322102-1322211100131213-3320011121013120"></a>

#### `default_pool.advanced_options.panic_threshold` property

Type: `"number"`. Optional.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [proxy_protocol_v1](resources--http_loadbalancer--reference--group-015.md#canonical-0320332213030303-2110221221131312-1112220010132111-3111231331132203-1012323131311113-3203232033330032-3300331111032003-0132310330010030): complete subsection reference.

- [proxy_protocol_v2](resources--http_loadbalancer--reference--group-015.md#canonical-2002203333000021-3223210132201331-1111301110300313-0211110110221123-3233022321110020-0320011110202330-3210010203121323-3010312231222231): complete subsection reference.

<a id="canonical-3112103233111313-3303220101001113-3022220201103300-2310102331122013-0202002302022030-1310031130301100-1231313211312123-1131333023120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.auto_http_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.auto_http_config

<a id="canonical-2222110010010322-0311011300222101-1321130230021303-0012031022002130-1303301110220031-0232210132221202-3333222033322221-0011010213221202"></a>

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
auto_http_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031212032101202-3130331233121230-1022333233112301-0011311210120111-0130221013310320-2201133030023101-1021133320211103-0302331222320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.circuit_breaker

<a id="canonical-3233111032021020-3230121313312001-1120032120103330-1002003010201123-0022010201321001-2032101002202110-2203010031122031-1212133000010020"></a>

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

<a id="canonical-1221103120111131-2221033002320213-1232223002011222-1212322133031003-1011213203310033-3113233022103201-2323002010222032-3001120000202111"></a>

### Direct properties for `default_pool.advanced_options.circuit_breaker`

<a id="canonical-0133112222330102-1032011100210302-0031311111011000-1212112133011122-0121002331103103-1222220330020120-1323113323301021-1120113021000211"></a>

#### `default_pool.advanced_options.circuit_breaker.connection_limit` property

Type: `"number"`. Optional.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0333100220301322-1203012213120213-0310200333002312-2111110213201020-2003333032011303-0000102130122213-0323211021030211-0103130301313232"></a>

<a id="canonical-2101213313303200-2323310131233003-1010130021213010-0212022331021201-1023013220112301-3100000210131220-1212311322303010-1232101103133300"></a>

#### `default_pool.advanced_options.circuit_breaker.max_requests` property

Type: `"number"`. Optional.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2111101031232003-3220223212221012-0223220231013213-2103022321333021-3230020330302000-3311303020313031-0120013023210221-3020132333002110"></a>

<a id="canonical-3302121332111030-3311211112032323-2132130301331000-1331003210131111-0132211222131112-3012233101213111-2033100012110210-3320211122002033"></a>

#### `default_pool.advanced_options.circuit_breaker.pending_requests` property

Type: `"number"`. Optional.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2031113110313111-1030203300320212-2120323032103331-1303331112120122-3300132313212001-0320113200212013-0212313131121332-1013200031133311"></a>

<a id="canonical-1313312300101110-1133011322223203-2212123132030010-0020003030021230-3001133121121230-0312033301322131-0111202231200122-0313311020231000"></a>

#### `default_pool.advanced_options.circuit_breaker.priority` property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Additional upstream details:

Priority routing for each request. Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DEFAULT","HIGH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3022133031223010-0321131220313232-0000323221221321-2100300112331301-0113112000323012-3123323321001211-2111130031313231-1122311122012123"></a>

<a id="canonical-1230313122110020-3133123221122111-3212311311131022-3302310231012113-2303322211312003-0210313203212201-3313111333112133-0133213231312220"></a>

#### `default_pool.advanced_options.circuit_breaker.retries` property

Type: `"number"`. Optional.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0220231330001103-1312111311013211-2301111133301230-0123001132310332-0121313113212333-0001211212112332-2201120003102123-2223202111113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.default_circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.default_circuit_breaker

<a id="canonical-2311223321130021-2322220301201322-3103101100201223-0120313123301311-2321110002000220-2312122131102120-0312000013010231-1231030012332031"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
default_circuit_breaker = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030013113131131-3103010310302203-2123231330311322-1231031222331233-1321202223300003-3001203202013202-3000310231301223-3101032101210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_circuit_breaker

<a id="canonical-2012333203011332-2210122221231131-2320332013312011-2020133202331001-1001232330302121-1220010312001323-2323222202230333-3213103103132000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable circuit breaker.

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
disable_circuit_breaker = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113100222020202-3133203310100220-0330332303302301-2031303313121332-0323121303222111-3223132221322032-0310310300022112-3030312000202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_lb_source_ip_persistence` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_lb_source_ip_persistence

<a id="canonical-3311300012320322-0121211301313010-1001312211033110-2112333221022112-2000221032332123-3333222120211231-3021032220312201-1010313013033110"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_outlier_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_outlier_detection

<a id="canonical-1323102303233300-0000323231203132-3203201112023100-2200332202320100-3123011331310031-1202232022310313-3112323231002230-3212232131030022"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
disable_outlier_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323211300100232-0123211012202031-0232200001130201-1010020100302322-1122010113130203-0321323222203113-0330110330301230-2203311200223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_proxy_protocol` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_proxy_protocol

<a id="canonical-1023212212311013-1211022333031122-3310032221123012-3233100303313210-3223223220022321-2120110130031302-0323111013200320-1132031110022002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable proxy protocol.

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
disable_proxy_protocol = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023103001332003-0303111121333030-1302202301203023-2300202001223132-3000303201112022-3312000113212101-3222132031232321-1110111201302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_subsets

<a id="canonical-0102132000213013-3322002221102033-3100320030031203-0311232200232330-2012212020100231-1022012212232312-1031101312203130-3022202110311302"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
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
disable_subsets = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030023020110210-2012011100303020-3003123221320323-3003020312113033-1230031030231320-0130321233030101-1032013131300003-2201032222030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_lb_source_ip_persistence` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.enable_lb_source_ip_persistence

<a id="canonical-0202213131211030-2213121123310313-1232230113113203-1130231122332210-2311232210013010-3330101020133100-3210222033313003-0003033310313113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.enable_subsets

<a id="canonical-2230231312121130-1032301133102111-3111133223133033-0203011113020101-0313311202133022-2203333022200212-1022103132100131-0013001010332212"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1310113220101322-3111031123201002-1233010321102100-2323002200302210-2031231320101113-0212210313232022-0131320032033001-1321310000123020"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets`

- [any_endpoint](resources--http_loadbalancer--reference--group-015.md#canonical-3233312033110200-1111010030213011-3320211103320013-3233012111330131-3103031321022103-1332102101221111-0033311313130002-0332111310210222): complete subsection reference.

- [default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233): complete subsection reference.

- [fail_request](resources--http_loadbalancer--reference--group-015.md#canonical-1333011222302002-0211100111302313-0332200031311133-1223021130221131-2020220212132020-2332130122030302-0320320000130321-1012020200232321): complete subsection reference.

<a id="canonical-3233312033110200-1111010030213011-3320211103320013-3233012111330131-3103031321022103-1332102101221111-0033311313130002-0332111310210222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.any_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.any_endpoint

<a id="canonical-1000303030013302-0323313313321130-0112221323000311-0001222213133300-3232131301100302-1233322320321222-3212021133220311-0130101322100013"></a>

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
any_endpoint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.default_subset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="canonical-2210003123013121-2200321322330231-1321031021100120-3303020130133000-3022121011322111-3220312321220133-3110011210223020-0303013221212000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Additional upstream details:

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

<a id="canonical-1212323300030301-2101230323032332-1332323102200012-3013120303031313-0220020310302313-0123213122022323-0230200210322021-3021011022110232"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets.default_subset`

- [default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-3100312012123303-1223333121231133-0202011211212120-2012000330200330-1212201010233222-2202021010100311-0322010101113001-0002110020100032): complete subsection reference.

<a id="canonical-3100312012123303-1223333121231133-0202011211212120-2012000330200330-1212201010233222-2202021010100311-0322010101113001-0002110020100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.default_subset.default_subset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [default_pool.advanced_options.enable_subsets.default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-1111223122221220-2231311230203031-1310320333000310-1332302120233002-3131222012030031-0320111110202232-0233111213101133-2322211110312302"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

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

<a id="canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-1310130021001011-1033003231003212-2010212211212232-2033303132320232-3131022211323031-3030201220333200-1112320002232101-2331300232211331"></a>

Type: `"object"`. list nested block, Optional.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3101302323312131-1330310013001031-0332303233021231-3020221221211133-2101332303320330-2311022212121133-1023113132222233-3023010030233330"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets.endpoint_subsets`

<a id="canonical-0010310311122032-1331031001123323-0320021331122000-0312200133331120-0323102213112301-0333113231032111-0122130132200220-0300312330311310"></a>

#### `default_pool.advanced_options.enable_subsets.endpoint_subsets.keys` property

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1333011222302002-0211100111302313-0332200031311133-1223021130221131-2020220212132020-2332130122030302-0320320000130321-1012020200232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.fail_request` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.fail_request

<a id="canonical-3332023132101122-3133211002232233-3321102022000122-2023031000031301-2111020231211022-0333222213000122-1030220223302101-2020232133333012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fail request.

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
fail_request = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.http1_config

<a id="canonical-2221310333013012-0112023230002320-1220133001210101-1011302103303010-2030033120011312-3332030320200210-0232223130003102-2003031030013012"></a>

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

<a id="canonical-3200123012322231-3112020112303202-1233213220112300-2200013023230221-1221011113121112-1222200113201011-0012232231000300-2001120102332122"></a>

### Direct properties for `default_pool.advanced_options.http1_config`

- [header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301): complete subsection reference.

<a id="canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--reference--group-015.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000)
- default_pool.advanced_options.http1_config.header_transformation

<a id="canonical-1000113330203301-1100310033330203-1323013023031332-3022122120003021-0023132023302111-0303311001230100-0313113110103132-3112212200131003"></a>

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

<a id="canonical-0131012320333131-0013222130331011-3222030203313010-2230301221110311-0002220132232120-0123212201223023-3120202312222112-0122003220010330"></a>

### Direct properties for `default_pool.advanced_options.http1_config.header_transformation`

- [default_header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-3312102333213120-3222103231310230-3100210002322120-2123000120100030-3023002010123200-2112202120010131-2132332032100001-0213002223200221): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-2001111330010100-1023133232013203-3031230000022101-3213032101203111-0230103100112332-0110303220232220-0120103130022012-0022023111012013): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-1202032102112013-2121032320000201-0232132011133103-3320333120322221-3100322103120321-3310103101031233-3231312132201021-1222133231213032): complete subsection reference.

<a id="canonical-3312102333213120-3222103231310230-3100210002322120-2123000120100030-3023002010123200-2112202120010131-2132332032100001-0213002223200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--reference--group-015.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000)
- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301)
- default_pool.advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-1031033102101311-3020130122212213-2303100323331023-2002230302011231-0201122202221202-0322001310203323-2202020033023230-3111030322311321"></a>

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

<a id="canonical-2001111330010100-1023133232013203-3031230000022101-3213032101203111-0230103100112332-0110303220232220-0120103130022012-0022023111012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--reference--group-015.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000)
- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301)
- default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-2103033032222132-1120220002133221-0333000321000122-1211032112303110-2022313301313132-2300313002023233-3002000313330330-3330111302123203"></a>

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

<a id="canonical-1202032102112013-2121032320000201-0232132011133103-3320333120322221-3100322103120321-3310103101031233-3231312132201021-1222133231213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--reference--group-015.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000)
- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--reference--group-015.md#canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-0000222203303112-0122213022011113-3030201022123002-0000101131102201-3322002223322003-1232121223002120-2312100200023120-3000003211023302"></a>

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

<a id="canonical-3113102323303220-1222012231233210-3132230022230202-2102202133132231-2231111233120333-0131210111011331-3220111022033123-3110000201103023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http2_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.http2_options

<a id="canonical-0310020333302010-1022102012022203-0220011012212123-0003210220202020-1003300332320321-0303320121313212-3330130321032323-2113301001212100"></a>

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

<a id="canonical-0203031223211302-3003303130123131-0022201013231201-1120330300002300-1131222012311111-3221003133332003-2303132101310032-1230112012022332"></a>

### Direct properties for `default_pool.advanced_options.http2_options`

<a id="canonical-0113322321003310-3031203221002122-0312001121012023-3302322332101222-3110231332130231-2021020220320221-1310002013202113-1000112201122032"></a>

#### `default_pool.advanced_options.http2_options.enabled` property

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

<a id="canonical-3010321022203332-1133303230302303-2331101002312232-0110221102133000-2100020203120200-3333120023011003-2312121111020320-1223232102333011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.no_panic_threshold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.no_panic_threshold

<a id="canonical-1022211332103313-3020222102122012-3312232213020301-1031331320310032-2310123103321230-0201130000300002-0132123001123331-3011022332011223"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no panic threshold. Defaults to \`map\[\]\`. Server applies default when
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
no_panic_threshold = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202011012323122-2310202220132112-1203221321333321-3232310230130231-0223011220002020-3131220013013302-1112021212030131-2212302110230313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.no_request_limit_per_connection

<a id="canonical-0232033110230333-1121230303102131-3231103211320301-1001102220101103-1200202132200112-2001221021002130-0201000313020000-1023312220323033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-3122120002230323-1331302213332300-0221022000101211-0233001322022311-3200121333120233-0120103223222330-0001121011111302-0100313133112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.outlier_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.outlier_detection

<a id="canonical-0322033302102001-0111111010013331-0310230033311320-1211011211033023-2110233221133002-1221112103100020-2111010030323303-3232130210033131"></a>

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

<a id="canonical-3232003102221213-0212031130312033-1212331112101112-2122210233103230-3332102122011111-1211123312233321-2312223333202332-0320131320332322"></a>

### Direct properties for `default_pool.advanced_options.outlier_detection`

<a id="canonical-2122112033122330-3200310113233113-1031001130002321-0031203322011012-0000102103033110-3321301223010010-2020120132332310-1103101030200032"></a>

#### `default_pool.advanced_options.outlier_detection.base_ejection_time` property

Type: `"number"`. Optional.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3122221300330312-2212122102210131-0003123201103213-3321001011310233-1120203131222312-2023320001201320-0003023100022323-0200000330323100"></a>

<a id="canonical-3001201303110223-0223203001012023-1330030330132232-1133331013333311-1220332333131120-3300333011310333-0031310111331111-0002301100131300"></a>

#### `default_pool.advanced_options.outlier_detection.consecutive_5xx` property

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates the
number of consecutive 5xx responses required before a consecutive 5xx ejection occurs. Defaults to
5.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100121320010132-1332331303302122-2011032023232101-0200021012000302-0203123333330231-2203303331311031-1222231222101100-1003032303130221"></a>

<a id="canonical-1211111201123220-1001312013301120-3210213121100333-1221123032223303-2101133131002022-3001001122030302-1031102102310000-1310113221301220"></a>

#### `default_pool.advanced_options.outlier_detection.consecutive_gateway_failure` property

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3322103320220021-3012301020112121-0022102121320100-2200100010323333-2013303111333130-1202103321003321-3221313002210111-2131332013030313"></a>

<a id="canonical-0033001102232131-1111223000300302-1322203103302303-0121013322310210-0030102323330021-1200202123011203-3003331301131020-0211012210010232"></a>

#### `default_pool.advanced_options.outlier_detection.interval` property

Type: `"number"`. Optional.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Additional upstream details:

Defaults to 10000ms or 10s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2130130021103030-1020002001300001-1321311011230201-1023322131333222-3232321231032111-3032330211020331-2330001123131302-1333303201113123"></a>

<a id="canonical-3233323112020323-1122010220132201-3231322121312232-1312023213321020-3010300223200223-3223210333222202-3310130121003130-3122310233001132"></a>

#### `default_pool.advanced_options.outlier_detection.max_ejection_percent` property

Type: `"number"`. Optional.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Additional upstream details:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320332213030303-2110221221131312-1112220010132111-3111231331132203-1012323131311113-3203232033330032-3300331111032003-0132310330010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.proxy_protocol_v1` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.proxy_protocol_v1

<a id="canonical-1100222112113012-3331221121020132-0002130310313301-2031133103032001-1123033001100031-2130333133322222-2211303320132010-0033123222311220"></a>

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

<a id="canonical-2002203333000021-3223210132201331-1111301110300313-0211110110221123-3233022321110020-0320011110202330-3210010203121323-3010312231222231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.proxy_protocol_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.proxy_protocol_v2

<a id="canonical-0131302130110311-0323020023003313-3312112231113110-1321203003332230-0232213210200303-3201131313312003-2102203020001222-3311210232020033"></a>

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

<a id="canonical-1322002312102330-3201331032321312-3300300133011331-3120132210110301-3210102232133203-0002132201013013-2023203103322300-1312001000200312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.automatic_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.automatic_port

<a id="canonical-2313110132311100-3111321122021032-3123313000302213-1213001120002121-2312123132031000-3002131231002120-3122032130323102-3221311302001001"></a>

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
automatic_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221110002132303-0012322203331031-0200320230330313-0111101330333233-2122220212333113-0311121023131102-2021213213213203-2231002020013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.healthcheck` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.healthcheck

<a id="canonical-3321201032103203-0010132301030130-0101111301123300-3232003021233132-3102003223230202-2233101310232110-2131000130131013-0131200331133313"></a>

Type: `"object"`. list nested block, Optional.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1330311201101220-2220111230332220-3000132313301112-2202330212323032-0330333123222130-0133030023311220-1030200330322313-2110303201301003"></a>

### Direct properties for `default_pool.healthcheck`

<a id="canonical-1313133020001221-2223020222113332-3131300311221012-0111230220233312-2220230012121302-2110103203033201-2303230130232130-0302330001121121"></a>

#### `default_pool.healthcheck.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0022021320323231-0101113022022222-2111121010211011-0021330332023110-1121133002003113-3210032023013133-0111223222222020-1212330220321012"></a>

<a id="canonical-2331012321331122-0122030210221323-1122221232213012-0033000021101121-0102203122121211-1201313021023002-1130023313332222-2311312010022203"></a>

#### `default_pool.healthcheck.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3321301232101322-2212310111232331-2200201302213123-3111301112123321-0322112302203230-3203130320111231-0102213000102333-2020000211220113"></a>

<a id="canonical-3230111113310201-0231220112203030-2112133013023213-2120333301000022-1313121300301101-1130211302012012-1331213202211322-0033331313112312"></a>

#### `default_pool.healthcheck.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1023332221023002-1120001320231221-2230003000223121-1203332032211030-3112300112131312-2203210121000020-1110310100200211-0012031020100313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.lb_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.lb_port

<a id="canonical-2110013230303102-3103011311021101-0103233330101302-2331301030331121-2333030110030320-1023233022120300-3321012212011110-1300010023230300"></a>

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
lb_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222232123033110-0311020222300102-2311031131230111-0311001112220210-0122011130233121-0322223310331112-0222133023233112-0031210201010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.no_tls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.no_tls

<a id="canonical-3201312012232201-1110033123111301-1020312022211332-3310331103120332-2103221222002020-0120230301023211-0001112030231102-0322002331022010"></a>

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
no_tls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.origin_servers

<a id="canonical-2120323112222332-2100013231020103-0321331203103333-0110132300202113-0322332002332001-1212121003303002-3031013232333230-1020230033331222"></a>

Type: `"object"`. list nested block, Optional.

Origin Servers. List of origin servers in this pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2330310100311201-3322223033132102-3230330310231013-3022330313100011-3303111330302031-1203201131310301-1202222222231022-2210221302212003"></a>

### Direct properties for `default_pool.origin_servers`

- [cbip_service](resources--http_loadbalancer--reference--group-015.md#canonical-1023000113222130-0300113230213310-0020100321213020-2110120302131213-3103233330311111-0222011323033301-0310203310220213-3332301122030311): complete subsection reference.

- [consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003): complete subsection reference.

- [custom_endpoint_object](resources--http_loadbalancer--reference--group-016.md#canonical-3320101200101333-0101332020212110-3023110121121203-0211302010232331-0301203130200233-3221233212020201-1322313330211100-1110201120130123): complete subsection reference.

- [k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213): complete subsection reference.

<a id="canonical-3001302100123112-2100303311201230-0002103212010023-1321133303012200-0323123112311220-3101320132033002-2000300200002123-0300202002020102"></a>

<a id="canonical-3011323010233330-0313323133232102-2103010332310101-0020102303010032-0020221231221202-0111320011113330-2113321211211311-1211203002113300"></a>

#### `default_pool.origin_servers.labels` property

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

- [private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021): complete subsection reference.

- [private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032): complete subsection reference.

- [public_ip](resources--http_loadbalancer--reference--group-016.md#canonical-1123323200112132-2011301210333320-0230333321233220-3120200223212323-0221012302113321-1102011330332121-0221103113230111-0011130112303131): complete subsection reference.

- [public_name](resources--http_loadbalancer--reference--group-016.md#canonical-3302320302223102-1023033333230010-2222202133303121-2101111202221230-2113030322031221-1332101010000231-1202031111300230-1003313332002300): complete subsection reference.

- [vn_private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122): complete subsection reference.

- [vn_private_name](resources--http_loadbalancer--reference--group-016.md#canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312): complete subsection reference.

<a id="canonical-1023000113222130-0300113230213310-0020100321213020-2110120302131213-3103233330311111-0222011323033301-0310203310220213-3332301122030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.cbip_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.cbip_service

<a id="canonical-3130003001130021-3222113232111131-0210012031232131-1322310103013203-3321332310131011-2230022311022212-3011312120020221-1202330100210330"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name")}
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
cbip_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220110101013213-2003220300320001-2201111130103002-3321133202030000-1022333200213011-3210322211011122-0333331331012032-3331230200112221"></a>

### Direct properties for `default_pool.origin_servers.cbip_service`

<a id="canonical-2001113012021001-2112200213011121-0100300311022320-2030022110301032-0213003113030330-3001310311333103-1021021300031112-0012131130100223"></a>

#### `default_pool.origin_servers.cbip_service.service_name` property

Type: `"string"`. Optional.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.consul_service

<a id="canonical-1110013101000210-0200322111110131-1101202321331121-1112102011033122-0302012211103303-3321113233030120-3133021032031010-1032223220223100"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with HashiCorp Consul service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

Terraform syntax:

```terraform
consul_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201300231032120-1101213131203323-3101123301322321-1013103210333311-3012323102321020-2113333110302202-2011311111300101-0333211121222002"></a>

### Direct properties for `default_pool.origin_servers.consul_service`

- [inside_network](resources--http_loadbalancer--reference--group-015.md#canonical-0021012310103003-3301013203103103-2103030200210222-2210112310201122-2220132020123330-0332113323120102-2000131110202233-3103320222100133): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-015.md#canonical-2211012211311332-0231231002020013-1011033211320212-2320300030231020-1130220012033121-3130000031030213-1031212001111112-1233202130213220): complete subsection reference.

<a id="canonical-3102232233311230-2122111220231301-3103013002030200-1232112230231311-0033301111103013-0200013131313133-0220300113323010-2032200132311010"></a>

<a id="canonical-3320213001233132-1300123323032132-2222111220131203-3131323113113321-0130000212312020-1310132211310232-0100000030120300-3123021322321130"></a>

#### `default_pool.origin_servers.consul_service.service_name` property

Type: `"string"`. Optional.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-1323111100223303-1312011332301001-3103112212130120-1123302201120002-2201130210332220-3203331113133221-3111311301322111-1101213123011332): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1203200311313231-3123130203231211-0212312001001320-2000333120210300-0003011112231121-2210223300312313-2213133112232123-2212023301012001): complete subsection reference.

<a id="canonical-0021012310103003-3301013203103103-2103030200210222-2210112310201122-2220132020123330-0332113323120102-2000131110202233-3103320222100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- default_pool.origin_servers.consul_service.inside_network

<a id="canonical-0012000312013211-1130231130120322-2312313323302003-3120231311313110-3300100332302112-2211000003011102-0313332101003000-0131212221233132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211012211311332-0231231002020013-1011033211320212-2320300030231020-1130220012033121-3130000031030213-1031212001111112-1233202130213220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- default_pool.origin_servers.consul_service.outside_network

<a id="canonical-2132133103031121-1110233121330100-3012312030003203-0121030123210001-3312210200021311-3303221201303222-0011130303023130-1213002033122013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323111100223303-1312011332301001-3103112212130120-1123302201120002-2201130210332220-3203331113133221-3111311301322111-1101213123011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- default_pool.origin_servers.consul_service.site_locator

<a id="canonical-3033002220133033-1110311030032122-2321130101232333-3212000301202133-1310331133012133-2231101313131232-1333200213223113-1201113012102123"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220011101131312-0303200001110200-0330313103303321-2021123232103030-1102013030031133-3200032001032133-2123211121000031-3231011112201200"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator`

- [site](resources--http_loadbalancer--reference--group-015.md#canonical-2030012200213100-0032323231132322-0220132010110331-1100102333013123-1330312330003113-1220211020120123-0211022231001203-3231002303021200): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-0133302031203232-3121110331002233-0133300101033002-1311320120221330-2330211033311312-3111021032210110-3212310010101333-1232332132302231): complete subsection reference.

<a id="canonical-2030012200213100-0032323231132322-0220132010110331-1100102333013123-1330312330003113-1220211020120123-0211022231001203-3231002303021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- [default_pool.origin_servers.consul_service.site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-1323111100223303-1312011332301001-3103112212130120-1123302201120002-2201130210332220-3203331113133221-3111311301322111-1101213123011332)
- default_pool.origin_servers.consul_service.site_locator.site

<a id="canonical-2300233021000212-2210000013220103-2223332013003230-2221312332120211-2313022112003200-3322231331130130-1020100113230333-1103313003010330"></a>

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

<a id="canonical-3113113013233311-0311222301021003-3223110121033330-1012100202100111-3220001023001121-1321232300000320-3100202311023231-2322302222132332"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator.site`

<a id="canonical-3331320100210100-2220133322223102-1011220212212033-1333133131310222-0112201220002020-0203002101210032-0013001311211320-3203121011323330"></a>

#### `default_pool.origin_servers.consul_service.site_locator.site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3213330300002223-2323030230303113-2322102032111120-1102003202203311-3221321330000211-2033233030322203-1312200131213102-1031120100132203"></a>

<a id="canonical-0122311303330313-3123321030130112-3112030112100023-1023202320301120-3211022233111222-3201013333202020-3121100011031122-1103021001232103"></a>

#### `default_pool.origin_servers.consul_service.site_locator.site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022121013300332-1213011100000302-2100321100030020-1001132213333013-0132033202123001-0112012302003102-2220031301330221-0130312032320122"></a>
