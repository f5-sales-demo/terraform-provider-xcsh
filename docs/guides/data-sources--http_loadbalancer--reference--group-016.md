---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3222121322103120-3302200030310232-0011123313010203-3112113122203333-0212312003332112-3022230312110013-2030212111331221-1133000310133031"></a>

## `default_pool.health_check_port` property

Type: `"number"`. Computed.

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

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

- [healthcheck](data-sources--http_loadbalancer--reference--group-016.md#canonical-1202331312130223-3323033003031000-1300032201301212-2133102012302220-1012021120110103-2231131032011311-0213332212020202-2021031122333020): complete subsection reference.

- [lb_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-2122312022231123-2122221001303110-1111302333031202-2020023322122331-0020010213103110-0012000101322211-1200121210030020-1320232013101313): complete subsection reference.

<a id="canonical-3201120301112012-3101322101210201-3103000300300123-2201300330100000-1031203121212222-2133232010120212-2020111233030003-1332020332311322"></a>

<a id="canonical-3131110110111002-1232332301213121-1021130131010221-0021201233213311-3020102331310203-0132203223300220-0011012211310323-0202110220330312"></a>

## `default_pool.loadbalancer_algorithm` property

Type: `"string"`. Computed.

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

- [no_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-1102212210131000-1000101230031101-0023123120321223-0123033000022131-0222321112012111-0001230102130023-2110102101300223-1203122202213132): complete subsection reference.

- [origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000): complete subsection reference.

<a id="canonical-2022100303132311-1033010221001301-0310233200232222-0002003122000012-1211022211021322-3002002021301230-0220302310311031-3010201002213020"></a>

<a id="canonical-3131020313302000-3020010200101022-1221303303212021-0120313223121122-2303103132333303-2313310021102112-2212322211310321-2020213020123012"></a>

## `default_pool.port` property

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

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

- [same_as_endpoint_port](data-sources--http_loadbalancer--reference--group-017.md#canonical-1220231310023223-2021212032323232-0022231222211011-1010322033032201-0313321021233313-2333033123131031-3012210101010020-1022012223113311): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-017.md#canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012): complete subsection reference.

- [use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031): complete subsection reference.

- [view_internal](data-sources--http_loadbalancer--reference--group-017.md#canonical-2322220323320321-1023031031123201-0303333212100213-0313203031010132-0222033300002102-3131312330012013-2002323331211102-0113131333012333): complete subsection reference.

<a id="canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.advanced_options

<a id="canonical-1213011323322132-3010010030122001-3022303000113300-2213101320030312-1221223103103220-2311330332110122-0302030120203210-0310033332231120"></a>

Type: `"single"`. Computed.

Configure Advanced OPTIONS for origin pool.

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

<a id="canonical-2300201212313221-1003000321321132-0120022230001130-2022013331212333-3331302332102121-2011233212220113-1331013102201302-2033310231003120"></a>

### Direct properties for `default_pool.advanced_options`

- [auto_http_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-1130221102311203-0231221321000220-1010101302131120-3122131031311211-1022130022230031-1203101122023121-2310200000010220-1133033110333012): complete subsection reference.

- [circuit_breaker](data-sources--http_loadbalancer--reference--group-016.md#canonical-0232020321002221-1233132132022200-0013031023102223-3210100000312120-0101013022202003-0033322120303102-3303002223033210-3201111133320202): complete subsection reference.

<a id="canonical-0132330233100200-0211310020333001-3302031320000121-2211022212210102-3022321330300300-3312000212113220-0022001203002230-3211011112121111"></a>

<a id="canonical-1330203002131101-2332103232333201-3230103133232323-1002023311102101-1233332221332020-0110111123003101-3002001101200113-0331102301311003"></a>

#### `default_pool.advanced_options.connection_timeout` property

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

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

- [default_circuit_breaker](data-sources--http_loadbalancer--reference--group-016.md#canonical-1132203130230100-0021230222302023-1012032232112121-3112301022311233-1010303330331103-0121010203330113-2100020332123203-1302311223302003): complete subsection reference.

- [disable_circuit_breaker](data-sources--http_loadbalancer--reference--group-016.md#canonical-1103122131020223-0233032003331303-1203212013001231-0131222300130203-3021303201233003-1211300321313311-2102120320200231-0201002130102103): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-016.md#canonical-3020030101330100-1312220330202101-0111001313103013-0331013030011111-3112230002222211-1112210013113123-3220132020200020-2031212222123123): complete subsection reference.

- [disable_outlier_detection](data-sources--http_loadbalancer--reference--group-016.md#canonical-1332012311033103-1321022311010313-0101213221232232-0132003113002213-3021102132022102-0220110203312233-3301313112203210-1011220103310130): complete subsection reference.

- [disable_proxy_protocol](data-sources--http_loadbalancer--reference--group-016.md#canonical-1030313332320200-1303011110223202-3320200013010123-1213002201221302-0322320302210223-3011323333221012-1011213301321111-1030131230021230): complete subsection reference.

- [disable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-0321332300211123-3230221211223301-3331112313102200-1103202311312112-3110110211112113-1131120230022300-2112221130003120-3310200211021222): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-016.md#canonical-3020213032033211-0032113103233022-0223003002111333-3021113232030003-2300200121010101-0213221210001011-3032132003022122-2110211110222232): complete subsection reference.

- [enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003): complete subsection reference.

- [http1_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023): complete subsection reference.

- [http2_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-0302103210203123-2300200322030323-3303020103012020-2031120223010323-2313200232200231-3203333211320133-2033203220101011-1023222100033000): complete subsection reference.

<a id="canonical-2000112112220300-3312013130222213-2003213001031320-2021331312022232-0020032120210012-2013201303203033-2210013021201122-2211121030301311"></a>

<a id="canonical-0130101123322211-0200121112320032-3011113023003311-2202222010222223-2130200121121132-0202330211133112-1022301330233033-3220333123201203"></a>

#### `default_pool.advanced_options.http_idle_timeout` property

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Additional upstream details:

Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 5 minutes.

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

<a id="canonical-0112210022031110-3312312021113132-2030002010031131-3201013000230020-3023022231223312-1032133123013111-0012220011130020-1033221122133311"></a>

<a id="canonical-1003112301113222-1232113321031002-2221120220230023-1020003032202332-1332233003032033-2123312113303000-0303100033213111-3210201330120102"></a>

#### `default_pool.advanced_options.max_requests_per_connection` property

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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

- [no_panic_threshold](data-sources--http_loadbalancer--reference--group-016.md#canonical-0212013320121110-3033210311131000-3201212002231122-1110302332211220-0112213030302123-2320021233211303-0110130210300021-1323000111111223): complete subsection reference.

- [no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-016.md#canonical-1321012101122211-2102000320220313-1331310322231032-1203123320313111-1311130321320303-1020331231312103-1330303100211131-0032021020020300): complete subsection reference.

- [outlier_detection](data-sources--http_loadbalancer--reference--group-016.md#canonical-0233100000211130-1311321230320220-1220101003130100-1202010131212223-3301032010002020-3231201110022231-3002222332232320-2213210303031120): complete subsection reference.

<a id="canonical-0330333010122110-1021103333321221-1023201000303013-2112002023222000-3001030321111022-1213023030031231-2230232331211023-1223102123322000"></a>

<a id="canonical-0301101213212233-2312111232121222-1123332001130131-3232132021031123-1231101121132331-1300121012213331-0310332001001312-2300313012310103"></a>

#### `default_pool.advanced_options.panic_threshold` property

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--http_loadbalancer--reference--group-016.md#canonical-1320312022011223-0112101033132110-2322303210023201-1320121013000330-1331321233310011-3220111102322200-3021312012310331-0301233032003001): complete subsection reference.

- [proxy_protocol_v2](data-sources--http_loadbalancer--reference--group-016.md#canonical-0211103313321032-0201122021031121-2311101000312300-2230110230320310-3213220020320023-2010000121013233-2311011010210301-3323120323303331): complete subsection reference.

<a id="canonical-1130221102311203-0231221321000220-1010101302131120-3122131031311211-1022130022230031-1203101122023121-2310200000010220-1133033110333012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.auto_http_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.auto_http_config

<a id="canonical-2303133030223031-0313032012122100-1231110001311200-0000001303113020-2222022002213112-3232300113033033-1301131010133312-0012202300023130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232020321002221-1233132132022200-0013031023102223-3210100000312120-0101013022202003-0033322120303102-3303002223033210-3201111133320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.circuit_breaker

<a id="canonical-0121223110230212-3121203013333032-1332033013002311-2331112200111322-1211332213231312-2332312321132130-2300010011023322-2233023320003310"></a>

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

<a id="canonical-2030121300032332-3230332202211013-1301030232220033-0031333113013122-3121230321113331-0213303200113103-3333001203001021-0022013332310021"></a>

### Direct properties for `default_pool.advanced_options.circuit_breaker`

<a id="canonical-0323330021323031-2122112203011211-2012120120021210-0300011020122120-3110221322121010-2333222133221303-3232301102202113-0210121203120221"></a>

#### `default_pool.advanced_options.circuit_breaker.connection_limit` property

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

<a id="canonical-1311203223221131-1213203021321233-0210011133223123-2102002123311320-2132213110311003-2303211212032233-3330133230113111-1233122131303110"></a>

<a id="canonical-2310231333321022-2023312212310230-1133200221212121-1200012100313321-3120023130212300-3110132122301203-2322113200233022-2121331033222302"></a>

#### `default_pool.advanced_options.circuit_breaker.max_requests` property

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

<a id="canonical-1003030011012210-2320001222001202-1213121213102020-2200002222010211-0330210212303012-0323310201213030-3320321010202102-0031113000100321"></a>

<a id="canonical-3302032131133313-3211311100211233-2300121020133320-1310320132322300-0003001310131132-3210031003301030-0330220201200033-2012000313003112"></a>

#### `default_pool.advanced_options.circuit_breaker.pending_requests` property

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

<a id="canonical-0210321322231000-3302132111022212-3013100033130330-0333102223330112-1031000202313323-2300311211101132-2321112113302313-1012311012231213"></a>

<a id="canonical-1132301300212303-2220022031000221-0030300123030123-2202123122232111-3202100000103100-1013313223101200-2331203333121233-1123120313131332"></a>

#### `default_pool.advanced_options.circuit_breaker.priority` property

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

<a id="canonical-0021330111130221-3011231023023132-1221103110100320-3003011111101211-0002103313332311-0132313030031203-0132112212200230-3230110311000301"></a>

<a id="canonical-1010333032132010-1013323302103032-0033133313302100-0032020002020323-1100210212101023-3030102122103300-2331331011210101-2301303120302130"></a>

#### `default_pool.advanced_options.circuit_breaker.retries` property

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

<a id="canonical-1132203130230100-0021230222302023-1012032232112121-3112301022311233-1010303330331103-0121010203330113-2100020332123203-1302311223302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.default_circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.default_circuit_breaker

<a id="canonical-3323333213022201-3233113132221222-2212302201312103-0001322321332010-1130211102211013-3021311121110300-2011321012311203-2022020332320002"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103122131020223-0233032003331303-1203212013001231-0131222300130203-3021303201233003-1211300321313311-2102120320200231-0201002130102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_circuit_breaker` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_circuit_breaker

<a id="canonical-0133311100122000-2110323231233301-3310232123110330-2113213201100121-3022330322021003-1033232311103210-3302113000333020-0320100102120002"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020030101330100-1312220330202101-0111001313103013-0331013030011111-3112230002222211-1112210013113123-3220132020200020-2031212222123123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_lb_source_ip_persistence` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_lb_source_ip_persistence

<a id="canonical-1111123310003113-3230003100000131-3023333111311210-1222201302212032-0111320222212121-3111101223302221-2023023033021331-3233201020101022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332012311033103-1321022311010313-0101213221232232-0132003113002213-3021102132022102-0220110203312233-3301313112203210-1011220103310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_outlier_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_outlier_detection

<a id="canonical-0111323303223211-0301331211031012-0331100132212102-2321333223310231-1030132023201032-3333100123331321-3123020333110322-1231230323112222"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030313332320200-1303011110223202-3320200013010123-1213002201221302-0322320302210223-3011323333221012-1011213301321111-1030131230021230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_proxy_protocol` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_proxy_protocol

<a id="canonical-2112322110223113-1231230330100332-0222010013223310-0311302311301013-1221231321103232-3113303221011303-2203301322320033-3301121013201302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321332300211123-3230221211223301-3331112313102200-1103202311312112-3110110211112113-1131120230022300-2112221130003120-3310200211021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.disable_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_subsets

<a id="canonical-2222333131110113-0130133330011002-2331330103022231-1011322200100000-2210133113300022-3031030030322023-3110212221013310-1020201221232103"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020213032033211-0032113103233022-0223003002111333-3021113232030003-2300200121010101-0213221210001011-3032132003022122-2110211110222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_lb_source_ip_persistence` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.enable_lb_source_ip_persistence

<a id="canonical-2110330023312300-2113130223210221-0023220220113303-2021103302133212-3010322101201221-3100301333103321-1100020022203310-2201010103302003"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.enable_subsets

<a id="canonical-0022320222203203-2202021010102323-3113301101302003-2222130322332132-0231311331202012-1020220022010203-2002120312330120-3110223100223201"></a>

Type: `"single"`. Computed.

Configure subset OPTIONS for origin pool.

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

<a id="canonical-0032033231111003-1103010133110130-2200213211230301-2332230321333232-0120130023220113-1121131320233210-2022110012130020-1320103003312001"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets`

- [any_endpoint](data-sources--http_loadbalancer--reference--group-016.md#canonical-0033000210210213-1033200310020203-0323331100133030-2000220212322321-3103221110220030-0120323132332102-1302311330200321-2003320122230112): complete subsection reference.

- [default_subset](data-sources--http_loadbalancer--reference--group-016.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2121322022211212-1333302100133030-0301201230333113-3032010200012231-1121031001010220-0233210000022211-1200323002312122-2132002133332223): complete subsection reference.

- [fail_request](data-sources--http_loadbalancer--reference--group-016.md#canonical-1021103000030322-0131231321303110-3133120333303230-1013320313021013-1320103003030210-0130323110332113-0022001201131110-3003322320132213): complete subsection reference.

<a id="canonical-0033000210210213-1033200310020203-0323331100133030-2000220212322321-3103221110220030-0120323132332102-1302311330200321-2003320122230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.any_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.any_endpoint

<a id="canonical-0311300213223111-2201202121002221-0122000013231000-0303221003130133-2130123231122221-3220201101202122-0122121112213103-2321300321003022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.default_subset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="canonical-2220211111101212-2133122301022030-3320221133101133-2312200212330211-3130123022203321-1112301220030203-2230323121103100-3211002303222203"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1000120121133102-3301110212121103-1203302322121002-1003001323213013-2320303032022000-0122333120132110-2222011121213232-3312223310102310"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets.default_subset`

- [default_subset](data-sources--http_loadbalancer--reference--group-016.md#canonical-0120210320232100-2212123202222221-3213201133300210-1121313230211211-3331230033212003-2321021021210032-3023030110313013-1120310032212120): complete subsection reference.

<a id="canonical-0120210320232100-2212123202222221-3213201133300210-1121313230211211-3331230033212003-2321021021210032-3023030110313013-1120310032212120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.default_subset.default_subset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-016.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-1301120023233123-0110021332212222-1322121213131100-2321322300122233-1012010110101021-2012230201020301-3302002330022231-3221111223032333"></a>

Type: `"single"`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121322022211212-1333302100133030-0301201230333113-3032010200012231-1121031001010220-0233210000022211-1200323002312122-2132002133332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-3201033213033012-0332321133203303-3230001303231313-3331013211332013-3010212232121031-2122031012332021-1133132220023310-0122200023210001"></a>

Type: `"list"`. Computed.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

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

<a id="canonical-2131300101120113-0322232122302010-0311302301010312-2322101020222012-0003332131321332-0132231321010300-2232121012123320-0131113033231032"></a>

### Direct properties for `default_pool.advanced_options.enable_subsets.endpoint_subsets`

<a id="canonical-1122303030032221-2022330332030310-1102310103331303-0013013320203310-3313221030310031-2102121010212112-3231110200002130-2110012323210111"></a>

#### `default_pool.advanced_options.enable_subsets.endpoint_subsets.keys` property

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

<a id="canonical-1021103000030322-0131231321303110-3133120333303230-1013320313021013-1320103003030210-0130323110332113-0022001201131110-3003322320132213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.enable_subsets.fail_request` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.fail_request

<a id="canonical-2313303003321112-0000210310032302-2211132133222033-1132100100302020-3131012101102313-1211313222310123-2210132123222300-0122122202210100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.http1_config

<a id="canonical-3121213230322101-3320303123221221-3003231201320123-2113331231122033-3033203001323001-0211012003001210-0300001013330303-0302221100100122"></a>

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

<a id="canonical-1221131010001301-3013103201312302-0123211310323303-0003211211101003-1100311121130303-3102100231221130-2003303131331203-3333030123123330"></a>

### Direct properties for `default_pool.advanced_options.http1_config`

- [header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220): complete subsection reference.

<a id="canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- default_pool.advanced_options.http1_config.header_transformation

<a id="canonical-3212331102111021-1102132100213031-3331312211201303-1111201302333210-3302111331303103-2233033113113030-1131322132320022-2002110131313100"></a>

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

<a id="canonical-3021022012211113-3213033302120013-0332102033201211-3312303021003330-0201221112321202-2123200311303000-3110222030303113-0331323210102333"></a>

### Direct properties for `default_pool.advanced_options.http1_config.header_transformation`

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-1031301133331213-2300012111223000-2130200222220110-0203311230222302-2132131323010111-1032200131122133-3332211210231010-1333323220323000): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-2322110023300022-1112030202333011-1220031323001120-0023100110112310-1331110010231311-2231310233230220-3310132123233210-0012132303311201): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-1323033233133112-2221211333320013-3022331323213222-0002321132302020-3131022033110130-0232310130213020-0331010110220112-1221033122230311): complete subsection reference.

<a id="canonical-1031301133331213-2300012111223000-2130200222220110-0203311230222302-2132131323010111-1032200131122133-3332211210231010-1333323220323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-1230103222331211-0033223020323100-0213012330200223-1202122223311022-3210133122213112-1131013013301320-2231023331010333-1111030103211030"></a>

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

<a id="canonical-2322110023300022-1112030202333011-1220031323001120-0023100110112310-1331110010231311-2231310233230220-3310132123233210-0012132303311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-2313122323023020-2221202012000010-0100231220120310-0312231122210002-0231013033332023-2231321210113010-1203202010223101-3021122302123110"></a>

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

<a id="canonical-1323033233133112-2221211333320013-3022331323213222-0002321132302020-3131022033110130-0232310130213020-0331010110220112-1221033122230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-016.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-0301010220313301-0000200200011023-0032313211231001-0322221220232302-3230101331201313-2033003103310002-3102123130322210-1001230203132313"></a>

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

<a id="canonical-0302103210203123-2300200322030323-3303020103012020-2031120223010323-2313200232200231-3203333211320133-2033203220101011-1023222100033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.http2_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.http2_options

<a id="canonical-1031310232203301-1103211232010213-1030121201031122-1002133103113210-0132111323003320-2213332302122311-2123332201312010-3330323330230003"></a>

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

<a id="canonical-2210020201233232-2001333123230232-3013230032233323-0312031321011211-1222020323232310-2121330302303202-1211320322121101-2033300011030131"></a>

### Direct properties for `default_pool.advanced_options.http2_options`

<a id="canonical-3033132012123301-2301031330001031-3033230032203230-2120310100222232-1013103330300012-3000111123110221-0221033111223131-2223022321020123"></a>

#### `default_pool.advanced_options.http2_options.enabled` property

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

<a id="canonical-0212013320121110-3033210311131000-3201212002231122-1110302332211220-0112213030302123-2320021233211303-0110130210300021-1323000111111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.no_panic_threshold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.no_panic_threshold

<a id="canonical-2102001132321320-2131332001300320-2221023030201300-3301330333303210-3312220002301102-2122310021031132-3220113301121300-1303312120203001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321012101122211-2102000320220313-1331310322231032-1203123320313111-1311130321320303-1020331231312103-1330303100211131-0032021020020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.no_request_limit_per_connection

<a id="canonical-3213313013103333-1010210020121031-3333032203130212-0012021030212213-1333233031223100-2201011212000112-3212332102200331-3221210232302212"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233100000211130-1311321230320220-1220101003130100-1202010131212223-3301032010002020-3231201110022231-3002222332232320-2213210303031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.outlier_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.outlier_detection

<a id="canonical-2021121330101131-0023221102321201-1102111202100320-2213210223231213-3303201011110310-2021233211300103-1301213331003112-0322130123323202"></a>

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

<a id="canonical-3121120333102223-3030310013032113-3120000110201022-1123101003220230-1121102313001211-3101113310303001-1102220013123003-2333230013021312"></a>

### Direct properties for `default_pool.advanced_options.outlier_detection`

<a id="canonical-0321123221032321-1121212112332200-1103312113133212-2211203220020331-2212110023120201-0123313302102011-0221130102311013-0230132333133310"></a>

#### `default_pool.advanced_options.outlier_detection.base_ejection_time` property

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

<a id="canonical-2102233030131311-1311020320200111-1020313003302232-1210301320013112-0122122021120033-0122012320302332-1121213221123303-0031123123101213"></a>

<a id="canonical-3111321122132102-1033223000022110-2030011321300120-1131231133012332-3031331201123223-2201030103221122-1133311211223032-3100011330232030"></a>

#### `default_pool.advanced_options.outlier_detection.consecutive_5xx` property

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

<a id="canonical-2322210233320332-3230213220110320-1120033111132103-1111123312002330-0001203033123231-0300230312200212-2022023310102311-2210011320031202"></a>

<a id="canonical-1311233110222110-3110230120302013-3011030303000023-2202121010123020-2030221030300101-2203301020100311-0032211111202331-3223120130031333"></a>

#### `default_pool.advanced_options.outlier_detection.consecutive_gateway_failure` property

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

<a id="canonical-0022200230312302-3323200320110333-0330031113103121-2212231021111323-1333100010320202-1331300332012020-1322232012132311-1310101233220102"></a>

<a id="canonical-1121201023111100-0302300023103110-2031010022333312-2320310130132013-0202321202020213-2120131112302003-3320321023101201-2210312131233213"></a>

#### `default_pool.advanced_options.outlier_detection.interval` property

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

<a id="canonical-2111013133301020-1031100000120202-2210030012232122-1002023222233111-0332013021131330-1312310133311213-2300321113210131-3223333321020120"></a>

<a id="canonical-0021320303113310-2331001212210220-2231101201302121-1220012213332002-3230132132303002-2120020231113302-3101113222230021-2210030030023323"></a>

#### `default_pool.advanced_options.outlier_detection.max_ejection_percent` property

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

<a id="canonical-1320312022011223-0112101033132110-2322303210023201-1320121013000330-1331321233310011-3220111102322200-3021312012310331-0301233032003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.proxy_protocol_v1` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.proxy_protocol_v1

<a id="canonical-1022300310301110-1013222230001303-1321223121122031-3302210203201232-3202011021200000-3002030130101030-2301333203102020-0221103110320020"></a>

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

<a id="canonical-0211103313321032-0201122021031121-2311101000312300-2230110230320310-3213220020320023-2010000121013233-2311011010210301-3323120323303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.advanced_options.proxy_protocol_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.proxy_protocol_v2

<a id="canonical-2221233201202201-3321310323030200-2331030221001122-1011021121021120-2123313101110221-1113020303123332-1032202100322312-3322032302111101"></a>

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

<a id="canonical-3333201113312220-0211103023020231-1231022023021132-3130330021322021-1011333013330100-1221301202121312-1123312313021112-0032303012213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.automatic_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.automatic_port

<a id="canonical-2000202230023223-2221130111130032-0330112333113212-2012022301223222-3023331313033032-1321120333233210-1033330330300323-3001110330311302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202331312130223-3323033003031000-1300032201301212-2133102012302220-1012021120110103-2231131032011311-0213332212020202-2021031122333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.healthcheck` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.healthcheck

<a id="canonical-0222331313333111-1100300002202300-1310203123203011-3211013210000120-0301320100310221-3200112003012230-0132301121212120-1122003300103131"></a>

Type: `"list"`. Computed.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0133020203000330-2300333101123121-1121313300030233-3331222232031302-3222033131311032-1331033131032210-3033121211110222-0003232320010131"></a>

### Direct properties for `default_pool.healthcheck`

<a id="canonical-2133012220020301-3333113212231020-2231102300202120-1120323023331122-2200232123101133-1221033021321111-2331032112120321-2130123113112003"></a>

#### `default_pool.healthcheck.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0311101012013130-3303321230202133-2113233312230310-3120002323122213-2330220110321230-3222201010103230-2221201110131202-0131111330301001"></a>

<a id="canonical-3323132303323130-0013000111030120-1112300102101131-3213300221223221-2300300323312030-3200001120020121-1103321130312330-2100333312032000"></a>

#### `default_pool.healthcheck.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3110113022031032-0321313003300313-2333103213331203-3021200222302030-1220120001322231-2130131312032230-0133213221222000-1211112331230310"></a>

<a id="canonical-0323210112320201-0202001210332032-1101120112003211-1233121121011312-1132133302301320-3030210100223210-1321021310321103-0113132002320310"></a>

#### `default_pool.healthcheck.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2122312022231123-2122221001303110-1111302333031202-2020023322122331-0020010213103110-0012000101322211-1200121210030020-1320232013101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.lb_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.lb_port

<a id="canonical-1323032232231033-1120102212003320-3222332333022333-2203011002113032-0231033220031303-2322212202133333-1221313001031302-1130210123313201"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102212210131000-1000101230031101-0023123120321223-0123033000022131-0222321112012111-0001230102130023-2110102101300223-1203122202213132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.no_tls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.no_tls

<a id="canonical-2010010321203031-1132003130232303-3311120001013322-3032312231122132-3213321031332211-2123300121002303-1112132113123100-3221131331300303"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.origin_servers

<a id="canonical-3200221210222012-1011330212322330-1011303021020311-1130231023312210-3032131323013121-1020022313021301-2000100030321332-1203110230010022"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

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

<a id="canonical-0232230331222022-3231010233203031-1020123032202131-1012311221103101-3211002020231203-0203210030230320-0102202011313213-0323310200223001"></a>

### Direct properties for `default_pool.origin_servers`

- [cbip_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1100110213210332-0331202020032201-1112330111231031-3202003310322001-0123113030213003-1233321303111302-0212313031020331-3112211233012012): complete subsection reference.

- [consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012): complete subsection reference.

- [custom_endpoint_object](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011112221011021-1222300212032111-2121332331230003-0310301311130002-2230221332312323-2332003230101112-2222223001102001-0201111233222103): complete subsection reference.

- [k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321): complete subsection reference.

<a id="canonical-3220012212232102-0122011121020002-2221120121012223-2301313121101320-3221121310023233-1131223123132232-2232311212030300-2000322013203300"></a>

<a id="canonical-3020020233022102-0230023310031021-0201232103230111-2303101131010222-0302222322132132-0212300013023121-0023102123333321-0323033303331131"></a>

#### `default_pool.origin_servers.labels` property

Type: `["map", "string"]`. Computed.

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

- [private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033): complete subsection reference.

- [private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222): complete subsection reference.

- [public_ip](data-sources--http_loadbalancer--reference--group-017.md#canonical-3230232123101031-2001223303123020-1213102222300331-0202333300210101-2132110133113330-1212330031022321-1002123133032130-2211123112002303): complete subsection reference.

- [public_name](data-sources--http_loadbalancer--reference--group-017.md#canonical-3213313232020122-3010103102320133-2122200332231332-1300302012030303-1213021123010331-2023230013331002-2031230031231102-3020232021202232): complete subsection reference.

- [vn_private_ip](data-sources--http_loadbalancer--reference--group-017.md#canonical-0211323132213300-2233131133130022-1221132031330120-3031032002223301-0212313121002022-3330103333233300-1120331120002031-0013013203231112): complete subsection reference.

- [vn_private_name](data-sources--http_loadbalancer--reference--group-017.md#canonical-0010223201122332-2321021313212022-0223002110212131-1011310123201121-3013320132013212-3232112322022332-1201100103213213-2331200010213312): complete subsection reference.

<a id="canonical-1100110213210332-0331202020032201-1112330111231031-3202003310322001-0123113030213003-1233321303111302-0212313031020331-3112211233012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.cbip_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.cbip_service

<a id="canonical-1023032123012012-1223301032311301-3103011301200020-0213200010003201-3201203123331103-2301311121122002-0010232201220022-3221110221003032"></a>

Type: `"single"`. Computed.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3230332012223123-2322222211131102-1111332020233222-2112003301322032-1131303121110013-1312133122010331-0332220322332223-2131110231131211"></a>

### Direct properties for `default_pool.origin_servers.cbip_service`

<a id="canonical-2301032302323030-1023111100022130-1331003333031131-1302012101003321-0231023333311130-2131032301013220-2300033130000222-1100302201303101"></a>

#### `default_pool.origin_servers.cbip_service.service_name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.consul_service

<a id="canonical-1011102303010031-3200003132020001-2032010332311102-1001022101210323-2130200220122201-1101130310103103-1122332232300133-1313311102112321"></a>

Type: `"single"`. Computed.

Specify origin server with HashiCorp Consul service name and site information.

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

<a id="canonical-2303110231203201-2130130311003332-2130033132220311-3311321111023302-1321201323131311-0132322221120121-2010021322232322-0310120120033120"></a>

### Direct properties for `default_pool.origin_servers.consul_service`

- [inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-2303011000103121-1202321131031130-1312032030123001-0100322010011312-1103002003010023-2203233111322212-1332012232203120-1123020331123103): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-2123003131002121-2101100002331012-0231302322200311-0131202231323330-3210011131000003-0220031210113203-1003133130103120-2331231203300310): complete subsection reference.

<a id="canonical-1231130201330020-3320201331200011-0203322212230100-3202023201202310-2220100300110000-3300201213020023-2323311020322303-0003303222330003"></a>

<a id="canonical-0000311321120203-0001323302132130-2023332120010131-0101130022001103-2223020010102101-3230002130302300-2330032301131301-0312003132332100"></a>

#### `default_pool.origin_servers.consul_service.service_name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131): complete subsection reference.

<a id="canonical-2303011000103121-1202321131031130-1312032030123001-0100322010011312-1103002003010023-2203233111322212-1332012232203120-1123020331123103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.inside_network

<a id="canonical-2303213312213120-1202023200323020-2213112323201103-2132000031312322-1100000201232230-2013103313011112-0001020321033002-1021330323020201"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123003131002121-2101100002331012-0231302322200311-0131202231323330-3210011131000003-0220031210113203-1003133130103120-2331231203300310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.outside_network

<a id="canonical-1231310033033000-0221220230321230-0313111002213120-2211111010110310-2333122113313001-3223310330032230-1200233220320220-3333313312031200"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.site_locator

<a id="canonical-3311002110203303-3221311131012203-2000030203233202-2102320103101003-3021323131012102-0233032331113323-2132000230123112-3320023231233001"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

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

<a id="canonical-2232201331222211-1223013301121330-2332030003023300-1300003032111312-0221231003312100-3230230012231202-0023130101022231-1321031203121110"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator`

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-0113133213010031-0221212130213112-1101012102300010-3020102310301112-1102021210300300-2220312222302102-1123231320022231-2210313003012232): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-0211211332011303-3233123201231311-3302120303322113-3031011333200011-1233213031111300-3020313112312331-2023000033332212-0132111131313211): complete subsection reference.

<a id="canonical-0113133213010031-0221212130213112-1101012102300010-3020102310301112-1102021210300300-2220312222302102-1123231320022231-2210313003012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- default_pool.origin_servers.consul_service.site_locator.site

<a id="canonical-3030001010123121-2100323101030311-2121121322133021-2122332330103001-2320323233123012-1300011230322010-3200222031200101-2003121101322301"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131133330211210-1033303011110322-1023003322132111-1101323323103200-2232111331202111-1301233211210311-3311203002322203-0032001111130320"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator.site`

<a id="canonical-0303031302132022-2122200122203010-0031102112212303-0103322321033102-0033212202221021-3012222202111212-3112212002303220-0330121010322300"></a>

#### `default_pool.origin_servers.consul_service.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3223113320211321-3102313112122100-3002221312133032-1212233311301203-0132201133303130-1202101320313033-2333331012311011-2200222301232322"></a>

<a id="canonical-3233031301003222-1023033322110000-1100030020113210-1233323123321210-2023020122313020-1032120222001210-1010010323000131-3231012123102010"></a>

#### `default_pool.origin_servers.consul_service.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0310303223130012-0012022310012113-1312212131130211-0301222133033012-2203020000102131-3210023012332011-1010321300233332-1010233103022130"></a>

<a id="canonical-3332131123303210-3311122321023031-0203003120031220-3333130220322131-1023130203200333-2310201101223233-3330213123012311-3210201110033230"></a>

#### `default_pool.origin_servers.consul_service.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0211211332011303-3233123201231311-3302120303322113-3031011333200011-1233213031111300-3020313112312331-2023000033332212-0132111131313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- default_pool.origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-1022002303110200-0201012003302300-1200331232123321-0033313122102300-1012310211302203-0131232030111320-0210320221012123-2031301011330033"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001122310011330-3113110021022220-3232311211213322-2300203201032313-0320221002230332-3313103023321020-1211311020110030-2100002200220232"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator.virtual_site`

<a id="canonical-2013331321002110-3320233022020002-2123111221023213-3333031132012213-0223120211301021-3032102132013212-2203231210030300-3132111232122001"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1131322131332303-0230032100320113-2201310302320010-3001211200311200-1312201323332200-1331032200113201-0013223321133232-0321012133201331"></a>

<a id="canonical-1130311202103121-0303201213311033-2023210112302022-3333311033323120-0011113033113201-2113112122120222-0231333010122103-0112112102321312"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3031112031222100-3121303012212312-1202201103122322-0001103202021131-2333030122100123-2302321300020201-2212312122310000-3201302021131000"></a>

<a id="canonical-0300011130312223-0232230311302011-0302123332112102-3030202021113220-2300302320332030-1313000210132031-1331323230303003-0011333331212023"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.snat_pool

<a id="canonical-0023021120003331-3212301010031133-2222311001200003-2022203111230023-1220030322212133-3202132213302111-2102020020332013-3301133013122330"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-3300220023323122-2030221131323121-3210310313120022-1332301111220223-3110200113013232-2022113233211003-3121130121123132-2221120113201102"></a>

### Direct properties for `default_pool.origin_servers.consul_service.snat_pool`

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-0323211101000312-2323202102110332-1032311331312130-0120203121233322-0130233213310201-3203220321211013-2332301302321013-3332000333312300): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1021332232010022-1231313332103001-0132110111321031-3200001013102112-2120001113120313-2110212112201313-2013232301203333-3010111102022313): complete subsection reference.

<a id="canonical-0323211101000312-2323202102110332-1032311331312130-0120203121233322-0130233213310201-3203220321211013-2332301302321013-3332000333312300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131)
- default_pool.origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-0023220123302333-1330200223203320-1100131213000233-3000210333102111-0320300130223322-2112030330031020-0131332303332211-2011133121110021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-1021332232010022-1231313332103001-0132110111321031-3200001013102112-2120001113120313-2110212112201313-2013232301203333-3010111102022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131)
- default_pool.origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-0111211103320113-2103212012213013-2332220312032021-1321012202113311-0002202233210310-3312032202030222-0023013222002303-2123010330303131"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3020131302320022-3010001112101210-0012233013013122-3311233231032320-0133223012302130-1232101220313003-2030113130111313-1032103223100231"></a>

### Direct properties for `default_pool.origin_servers.consul_service.snat_pool.snat_pool`

<a id="canonical-2321232122103030-2001123011112012-1201121211030233-0312001101333020-1212112110132121-1022100323321013-0112232103131200-2013303203200022"></a>

#### `default_pool.origin_servers.consul_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0011112221011021-1222300212032111-2121332331230003-0310301311130002-2230221332312323-2332003230101112-2222223001102001-0201111233222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.custom_endpoint_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.custom_endpoint_object

<a id="canonical-3103233322200131-3122233012311311-0213332020221031-0101023101200012-3021013311303102-1213231301000000-2333003122130202-0213320231213100"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0121232323210003-2003333231233022-2131231000110200-1302133213332200-1113013303100223-1231110111110322-1003133320121220-2331230123110000"></a>

### Direct properties for `default_pool.origin_servers.custom_endpoint_object`

- [endpoint](data-sources--http_loadbalancer--reference--group-016.md#canonical-0102023031322330-1323301331001022-2013323300230330-2302313011230031-0331121232021100-2323220211311302-1030322231301001-3211033200303101): complete subsection reference.

<a id="canonical-0102023031322330-1323301331001022-2013323300230330-2302313011230031-0331121232021100-2323220211311302-1030322231301001-3211033200303101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.custom_endpoint_object.endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.custom_endpoint_object](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011112221011021-1222300212032111-2121332331230003-0310301311130002-2230221332312323-2332003230101112-2222223001102001-0201111233222103)
- default_pool.origin_servers.custom_endpoint_object.endpoint

<a id="canonical-0331101120031030-0310220102220333-2100222200221113-3231123000121131-0330313030002000-3112323032012121-1100121323010333-1131000201113032"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002313133332202-3011123311102221-1003101230131122-0201113122202113-3102101133112000-2322123210122122-1112220321302013-2033231002103321"></a>

### Direct properties for `default_pool.origin_servers.custom_endpoint_object.endpoint`

<a id="canonical-0211330331213323-3212321233200203-2102101300321310-1330321120133032-0121033320301023-3033211211323201-2303032220011210-0220131132103223"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1100303202321012-1111113130010221-0311233032120031-1001131132320120-1323222103323210-2320312033312012-0232313203233033-3300112132122020"></a>

<a id="canonical-0301133012223021-2030110100120131-2100333313322033-0332001302230123-3131111202223333-2200013100330120-3201332222131030-2320323123112002"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3020032321231313-3223332321120120-0210130201311311-3300133302132001-0223100200210312-3133303123323120-3000001233111131-1321303231213133"></a>

<a id="canonical-1022300113000113-0320211332332300-1113303213330302-2020330103230212-2001210333031230-3303311010211323-3313323323002130-1131302301132112"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.k8s_service

<a id="canonical-1030003203011212-0123311230113330-0020313003102122-1302202330202002-2333310231311133-3032303030230311-1303123312133201-3001313311320020"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-2301103310313033-3102300101332200-3311100233312111-3030300103201330-0012202102321303-3110222210110023-2333302201323101-3111232333032332"></a>

### Direct properties for `default_pool.origin_servers.k8s_service`

- [inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-0002312232330202-3221202330100323-1002000020230331-3020301210221320-3321121010223111-0003230103211312-0130010222310301-0121220002233203): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-1321131231123103-3002222222213103-1120330010220222-3033101221123322-0131210101221033-3322212312111110-3223013231200230-0110011210102101): complete subsection reference.

<a id="canonical-0310331211200100-0020101030302302-2330322203013312-3021322000210102-2013211131223003-3100031211130301-3210010013020231-2133320023110303"></a>

<a id="canonical-1201212002300020-0102113013031111-2230012020123021-2101100021121102-2210012210001300-2331030203220121-1122313301131200-0311300211210302"></a>

#### `default_pool.origin_servers.k8s_service.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203202001232131-1110221300100312-3200330333133232-3212022011220322-0123113200321212-3223231230013100-2011113322303030-0102300323322131"></a>

<a id="canonical-3123321221310302-1100302333000101-2130320023213011-2210131033220010-0002313110102121-3310101301222133-2033000210300033-2331023002132223"></a>

#### `default_pool.origin_servers.k8s_service.service_name` property

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0313302230312132-1301212333103200-3332212301310110-3232203111233330-3133313111033301-2112123021003321-1133012201323323-2312333033232303): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1213311303022021-0001331121110013-2012122022233320-2012330323110121-2233123112021010-1222231033101223-1332322131233003-0001102201212133): complete subsection reference.

- [vk8s_networks](data-sources--http_loadbalancer--reference--group-016.md#canonical-1020233122012133-3100332120233030-3301311022103101-1221220131011330-2221001202232301-3231221303233233-1002121130231203-0231222220030221): complete subsection reference.

<a id="canonical-0002312232330202-3221202330100323-1002000020230331-3020301210221320-3321121010223111-0003230103211312-0130010222310301-0121220002233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- default_pool.origin_servers.k8s_service.inside_network

<a id="canonical-2313122230313111-0031013321030210-3002103321202000-2002203311210331-0111112003310320-3003212003021201-0003202121120303-3323331113230320"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321131231123103-3002222222213103-1120330010220222-3033101221123322-0131210101221033-3322212312111110-3223013231200230-0110011210102101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- default_pool.origin_servers.k8s_service.outside_network

<a id="canonical-2022013301111232-1233313113003022-2131220212213111-3212201132310333-1313132332131032-0201200131003331-0010203232022001-1102103003002130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313302230312132-1301212333103200-3332212301310110-3232203111233330-3133313111033301-2112123021003321-1133012201323323-2312333033232303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- default_pool.origin_servers.k8s_service.site_locator

<a id="canonical-3303131301211010-2022301212002211-0212130232101220-2021220131013120-2130020323320011-1032012231012311-1200331202113322-0210122222210231"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

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

<a id="canonical-3320022302121100-2031013110110003-2030101132020321-3232101023000321-3302133222033101-0110100301122221-3022311312113313-2230012223033230"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator`

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-3301003333001002-1233121011321233-1222032033030103-3230101022100223-0010321001212032-0300120221003301-2201101230331332-3232202230032303): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-0123302200113131-0301330201230322-2013010221130121-0232201330123002-3103202123120310-2310101112233201-3210220210212311-0031211021100332): complete subsection reference.

<a id="canonical-3301003333001002-1233121011321233-1222032033030103-3230101022100223-0010321001212032-0300120221003301-2201101230331332-3232202230032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0313302230312132-1301212333103200-3332212301310110-3232203111233330-3133313111033301-2112123021003321-1133012201323323-2312333033232303)
- default_pool.origin_servers.k8s_service.site_locator.site

<a id="canonical-0212001323212212-2101301312003331-0310301022023002-1121003210300100-1321212133312132-1303011113033003-1133320012100233-3212120302221323"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301321022001031-1123213303333030-3032023132330301-1202202010003131-3103003023211303-1033232300331100-0203003022232101-0022011030300102"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator.site`

<a id="canonical-3220033300212113-0303231102233010-0010131101120330-0030023123312310-0303331121312032-2000210200311111-3122022332012123-1030211131300010"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3322002101222330-1032103310121302-1131013300011122-3113300010121102-2112233202132322-3012203312312332-2101333330301202-3233201011200031"></a>

<a id="canonical-2230103033311313-2030110003312000-3223332312211330-3313133221330210-1210120101013330-0201331231110311-2221310233010303-3011221013313102"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2121333330202333-1312113203023203-3021210111003123-2120313011331232-3133113223223110-0312220011003300-1233323223223211-3001121221113330"></a>

<a id="canonical-0022132211010333-1222311301231313-1132222311230011-3020101310121301-0103332122132010-0222303010311230-3302230020011203-1232301010022001"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0123302200113131-0301330201230322-2013010221130121-0232201330123002-3103202123120310-2310101112233201-3210220210212311-0031211021100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0313302230312132-1301212333103200-3332212301310110-3232203111233330-3133313111033301-2112123021003321-1133012201323323-2312333033232303)
- default_pool.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-3112001222113332-0013232303212031-1013023202113323-2211303311131121-2110303311000111-3010133213121323-2102323301030031-2323022030123230"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131012120130223-2111131303211112-1101320113023122-2233210323213232-0320110230233100-3212312213100303-2213121113220323-0131201322201003"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-2300210213210222-2103323301333011-0203103213213023-1231023023211210-1031320323332013-2332111033320013-2032201213311211-3033211231311221"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3111032230101113-0001022300120130-3332012022112332-3202322232330321-2012330331031131-0330103132102021-2221021100202000-0300100112021200"></a>

<a id="canonical-2121011331330310-2130323023312111-2021002222333330-1022010000013133-3212022021000130-0133131300202002-2000121312021011-1111231311330232"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3332101121012202-1030310003210023-2021031313121131-3213202121122010-0210202212220111-3003322111313302-0023213100300211-0320312202022321"></a>

<a id="canonical-2133302323133312-2001211212301123-3020103110310110-1123122232213010-1020201203313223-0330111011312023-3322102131213311-1002202230131121"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1213311303022021-0001331121110013-2012122022233320-2012330323110121-2233123112021010-1222231033101223-1332322131233003-0001102201212133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="canonical-1323232013201101-2033210012312031-3023210110112012-0302121133222023-1010212331333301-2302101131013112-1202320022322122-0303033213203232"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-3210310233301333-1110032023120213-3312113001101133-1221111320220301-1332322323311002-2233112333233323-2011222112202332-1013030001220213"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1030330121000011-3202012300333201-3100021233000223-3212332121030011-1130310032012003-0012113202031131-0102322313302301-0322033322310122): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-2000321112121233-2110001001133000-0312122323133202-1203211210013211-0030020313212120-1121212332302110-1210113220320122-0000031031210223): complete subsection reference.

<a id="canonical-1030330121000011-3202012300333201-3100021233000223-3212332121030011-1130310032012003-0012113202031131-0102322313302301-0322033322310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1213311303022021-0001331121110013-2012122022233320-2012330323110121-2233123112021010-1222231033101223-1332322131233003-0001102201212133)
- default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-3031313112301313-3200021110222021-2202231031300130-1300033203210020-3322020222231230-3223201300100101-3312110332132131-2133002311132212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-2000321112121233-2110001001133000-0312122323133202-1203211210013211-0030020313212120-1121212332302110-1210113220320122-0000031031210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1213311303022021-0001331121110013-2012122022233320-2012330323110121-2233123112021010-1222231033101223-1332322131233003-0001102201212133)
- default_pool.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-0112202210101310-0103211322010332-0213313122021132-1010120013333133-3223102321101320-1010202021131001-3300322331032002-3012313033210000"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3130123231320203-3210211323111222-2031132310211010-3113321303001202-0103213032211000-1231023311033003-0123123213001130-2231331100000212"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-1020302103102300-2121031103221132-0021120013020020-0201001333302221-2003121010322010-0223001112012200-2100330331102000-3202132133223220"></a>

#### `default_pool.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020233122012133-3100332120233030-3301311022103101-1221220131011330-2221001202232301-3231221303233233-1002121130231203-0231222220030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- default_pool.origin_servers.k8s_service.vk8s_networks

<a id="canonical-3030331022132222-3233330221301223-3021001111023210-2201000212320112-0210201333123130-1202010301313332-2130023023321302-1213302210003322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

<a id="canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.private_ip

<a id="canonical-3003322011131020-0002031202212030-0102123221230230-3333332332333021-0022121331223223-3101221021212133-3103102132201211-2130123002223310"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-3221133003222313-3013131213111311-2121030202313302-2232130211232013-0121201320101100-1121312103022031-1012321002300212-2331122122120211"></a>

### Direct properties for `default_pool.origin_servers.private_ip`

- [inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-1133201323110323-1100130032102132-0211101030312220-3312131313300200-2031022003010113-2121101232022330-2111132223201330-0200302210213320): complete subsection reference.

<a id="canonical-0022023233003120-1130013113221031-1033222232222333-1332102201011222-1223112300221320-0020100333302120-0322200232132220-2200212031001000"></a>

<a id="canonical-3301002121031103-2212311122313003-3123013301230203-0030000322023030-1100101311300222-3231200121201111-2031300221111131-2001330120120311"></a>

#### `default_pool.origin_servers.private_ip.ip` property

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

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

- [outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-3331033001303231-2131030022320231-2011231332313102-1021213123101311-0020103330011023-2031133202111320-0013202022133132-1012133333121311): complete subsection reference.

- [segment](data-sources--http_loadbalancer--reference--group-016.md#canonical-3302222020032232-1322133301021331-3213020130001213-3002330233202310-2020033202221032-2231102130010303-0231301112033313-2122301131101121): complete subsection reference.

- [site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0120002023322211-3032122001223103-3031322302112111-0310022223012103-3223232323333212-0233011302022002-3011020312333322-2023013000023111): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-3001113110231102-3003033002211210-2201102232330112-3031010131033220-3332131332112130-2321232310010212-0312212112213212-3212203111311002): complete subsection reference.

<a id="canonical-1133201323110323-1100130032102132-0211101030312220-3312131313300200-2031022003010113-2121101232022330-2111132223201330-0200302210213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- default_pool.origin_servers.private_ip.inside_network

<a id="canonical-3203010223020001-3230212010002030-3123021012112000-1213120022011231-0002031023310022-3002300303333123-2010110213020030-2221223202122001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331033001303231-2131030022320231-2011231332313102-1021213123101311-0020103330011023-2031133202111320-0013202022133132-1012133333121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- default_pool.origin_servers.private_ip.outside_network

<a id="canonical-3233032231023220-3222311222320023-3112212001113332-1003223230320103-2002120312133220-3011131222220122-0113223020123122-3131303131321201"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302222020032232-1322133301021331-3213020130001213-3002330233202310-2020033202221032-2231102130010303-0231301112033313-2122301131101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- default_pool.origin_servers.private_ip.segment

<a id="canonical-2300123303201300-0111222131023021-3221132033102021-3120300130232323-0103213131010222-2010120023000131-2110100122022113-1130211130323220"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3120211323220121-2110020212020122-2000230202232110-2011213312031210-0322123302201221-0001022120132220-0023003011001102-0020122013020000"></a>

### Direct properties for `default_pool.origin_servers.private_ip.segment`

<a id="canonical-0222323022212102-2330130020022100-2011212233133211-2331231230010130-1301213111132200-3320033222230333-0331113122031123-2221231021202213"></a>

#### `default_pool.origin_servers.private_ip.segment.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0200321330211020-3221101322223103-3133212302101200-2300110331233211-0212310101033112-1133002122320230-1203223300223120-2101033012320021"></a>

<a id="canonical-0323322113021031-3300221313012331-2303322013122232-3311200031303301-0310211230323033-3312321101112011-2011201333022302-0011202322021022"></a>

#### `default_pool.origin_servers.private_ip.segment.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3321311111002231-2123320310223021-2121321033211103-0230221121211000-0310301132123030-0210311331121313-1200232233312031-2213231201021220"></a>

<a id="canonical-0302220110031231-1012123300131001-2321221122020022-2013302302302310-2321231330023311-3301032103033030-2100233031201111-3200231333000133"></a>

#### `default_pool.origin_servers.private_ip.segment.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0120002023322211-3032122001223103-3031322302112111-0310022223012103-3223232323333212-0233011302022002-3011020312333322-2023013000023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- default_pool.origin_servers.private_ip.site_locator

<a id="canonical-1301133222302313-2013103112131111-2031003320101310-1130300320313101-1021331001221033-0013232100000210-2030301303322231-2112133120000020"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

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

<a id="canonical-0100332000200001-0322201000313331-2002231132211301-3132310331210031-0322031220312200-0330201320123230-2313003002333000-3213032131030132"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator`

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-1312301201012321-1313012310003010-3321221001011020-0012023323233303-1101322031012201-2131130021301233-1213120112223331-3322211102031130): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-1331201331311221-2112321121102100-0120201231012303-3100013323021322-0020000210313111-3303012123230021-3212223231320100-1213203210012003): complete subsection reference.

<a id="canonical-1312301201012321-1313012310003010-3321221001011020-0012023323233303-1101322031012201-2131130021301233-1213120112223331-3322211102031130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0120002023322211-3032122001223103-3031322302112111-0310022223012103-3223232323333212-0233011302022002-3011020312333322-2023013000023111)
- default_pool.origin_servers.private_ip.site_locator.site

<a id="canonical-1210330121300032-0311103220331211-0120310100011001-0300302333133200-3331023221111120-3023323122003331-0213131232211131-1010122131111231"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033112232121002-2300011102221222-2222322311313131-0030311333121303-0332101221033102-2031003221113302-2101311022033020-3003200232202211"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator.site`

<a id="canonical-1013310012130121-0300231102103020-3313132000122301-0133112131321232-3031310122021201-3322010302220001-0333013322120121-0300133301012101"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3131013303003231-1200232300200301-2230332212012130-1100021310233201-1111313131303212-2221320331222110-3011330113221203-2101331311313030"></a>

<a id="canonical-3212322020120000-3101102232331322-3200033203121132-0312323201322231-2303003233100132-2331030130100000-2122321021101330-2232012201101231"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0323233313211310-2032300200000311-0010202132022031-3111121130100321-1112013203012303-2021103003220220-0330333133212000-0102112302302123"></a>

<a id="canonical-3302301201300231-1012130311301323-0320212002130212-3210021013102200-3113313331330133-0202302223132131-0020332001022233-1022113103220310"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1331201331311221-2112321121102100-0120201231012303-3100013323021322-0020000210313111-3303012123230021-3212223231320100-1213203210012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0120002023322211-3032122001223103-3031322302112111-0310022223012103-3223232323333212-0233011302022002-3011020312333322-2023013000023111)
- default_pool.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2013302222202023-2003020300232110-0100122003130113-1113212211113320-2321221112312330-0323130112202210-3112101300332133-1320230103123110"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021320322230112-0103023101112013-1302121132033302-0300320302220332-0000220212110001-3103201233110302-3003111203311112-2222112313200332"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-3003013301011021-3012322001013332-2230002022200201-1030120033201301-3133012233300032-2302212230000110-3013132300130202-2001210123312332"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2122002303322001-2310203211332312-3021120111121011-1002011121322012-0333120320022301-2230110012230320-3032100322003311-1122003102230200"></a>

<a id="canonical-1211232223233332-1311131101232021-3121312232122211-3323332333211103-1002102113023220-0033320230033331-0031133312122023-2112130300200222"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3321223232303133-2210203132331301-1031111033230302-2100222020320002-3212031103002313-3231133322113313-1301021310032223-2110320010021012"></a>

<a id="canonical-2202301302230333-2133132112102213-2311130002011010-2111231223130222-3300103121320330-3023123102202022-3203321022231320-3021100112311332"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3001113110231102-3003033002211210-2201102232330112-3031010131033220-3332131332112130-2321232310010212-0312212112213212-3212203111311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- default_pool.origin_servers.private_ip.snat_pool

<a id="canonical-2021101111211111-2122312122133321-3001222033232120-2113033001301300-2021001300300202-1131320330033232-3011303313100211-0100102230101323"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-3102213011203201-3030103213333022-0000313302212331-3101331101102133-3210212122333323-3311211033022013-2203001033210211-2012210033201013"></a>

### Direct properties for `default_pool.origin_servers.private_ip.snat_pool`

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1220310022310300-1001010312311001-1232022321200220-1030223331022011-3332220010232021-3310200200030023-2211220321313212-0312023322033111): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-1300103100202130-0203111203032100-2133313020112000-2013310223113000-1002122100302000-0212322332302133-2110100312221222-2201221300300312): complete subsection reference.

<a id="canonical-1220310022310300-1001010312311001-1232022321200220-1030223331022011-3332220010232021-3310200200030023-2211220321313212-0312023322033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-3001113110231102-3003033002211210-2201102232330112-3031010131033220-3332131332112130-2321232310010212-0312212112213212-3212203111311002)
- default_pool.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-1101010121010113-3011103103301003-0310320311320211-2211113310021320-1121021032011131-2003330123210322-3333100320200202-3233003221333302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-1300103100202130-0203111203032100-2133313020112000-2013310223113000-1002122100302000-0212322332302133-2110100312221222-2201221300300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-3001113110231102-3003033002211210-2201102232330112-3031010131033220-3332131332112130-2321232310010212-0312212112213212-3212203111311002)
- default_pool.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-3102312221210212-2323030103021010-3330122120021032-0200330110111032-3121312203001021-1201123320003312-3122030033020100-3020030111323131"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210310310012202-1333330223000320-3333133033130322-2203030303232203-0232020213122003-3002321211301203-0003113120121003-1322031131031110"></a>

### Direct properties for `default_pool.origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-2220200203111023-1232332222031212-0223113112223303-3322300332131311-1230203023031102-2212230312223231-3220303133011132-1213233202020000"></a>

#### `default_pool.origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.private_name

<a id="canonical-1032103001011312-2101111221203100-3131230203300002-3032211032220021-3201103103020012-2112223022203103-3003313302223323-1030223013300311"></a>

Type: `"single"`. Computed.

Specify origin server with private or public DNS name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

<a id="canonical-0131003131122123-2202133212200312-3223130213131132-2033303221230100-1210123030323022-3331020111113120-3322211200112032-0111001230001111"></a>

### Direct properties for `default_pool.origin_servers.private_name`

<a id="canonical-1332012313012111-0201212021231110-0231330101102112-1213231120211230-3212310301013123-3233321221021031-3232230223203231-0331120323223221"></a>

#### `default_pool.origin_servers.private_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-2000002120302101-0232300120201013-3321203121112313-1002031031331013-2133002020310202-0102123211110031-3331101313312301-0002031301330312): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-0003020030222011-3303101323110003-3233213133102223-1220313011011302-1130002231123002-3200313010121233-1333102330301202-1112000000101213): complete subsection reference.

<a id="canonical-3210223233012121-3123232202031033-2122010330323230-1210300213321001-3021230321200112-0113113210232222-1010302133023222-0221013123222023"></a>

<a id="canonical-3213033012213211-3231123101321331-3320010333322131-3201000332211231-1001001211112321-0103112333223320-3102211212001130-0032322112102102"></a>

#### `default_pool.origin_servers.private_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](data-sources--http_loadbalancer--reference--group-016.md#canonical-3320330210323021-2121033300300223-3331030011112113-2132013013323023-1220323132030132-3100310100121303-3210022033312320-1233032132221011): complete subsection reference.

- [site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0310320011032122-2311301301130313-0303200313201211-3022331013021321-0233302120202012-1122031222203303-0311100322233033-0301130233121123): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-3310022110103212-0302203330000030-3203123300323012-3232222312021232-3012311130300313-0231003000303220-1303031211212313-2321333311033331): complete subsection reference.

<a id="canonical-2000002120302101-0232300120201013-3321203121112313-1002031031331013-2133002020310202-0102123211110031-3331101313312301-0002031301330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-2333213113031313-2122033131303030-2322013320130003-0200003231312211-0103011100223210-0122013030122301-0231322311322313-2221210301201002"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003020030222011-3303101323110003-3233213133102223-1220313011011302-1130002231123002-3200313010121233-1333102330301202-1112000000101213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-1231310330222121-0330211032030103-2023211222312020-0123333330203002-2030203331132120-2003230232133203-1200221202110213-0300303320010011"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320330210323021-2121033300300223-3331030011112113-2132013013323023-1220323132030132-3100310100121303-3210022033312320-1233032132221011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.segment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- default_pool.origin_servers.private_name.segment

<a id="canonical-1100300331231023-3233312010220022-1001221002113230-2102021231321312-3233201132330012-1330113203223102-2200112322331101-2110222203233011"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311022020122330-2132302033210112-3201032311232300-3211321221102103-2210030302101322-0230013013330312-0111033112033131-1222212133233123"></a>

### Direct properties for `default_pool.origin_servers.private_name.segment`

<a id="canonical-3331312302212121-0213011323101111-0210200323213122-0312010221013102-0331311230301111-3023010331012130-0011303222102031-3321213021003020"></a>

#### `default_pool.origin_servers.private_name.segment.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1303012000031213-3233203031211130-0001210200322032-0333113023111111-0120331002201212-2130333300003030-2100013232000220-2131100031322122"></a>

<a id="canonical-3200320300122232-1111301101102222-3331001003022301-3302130023131032-3233302303132222-0101022130332222-1012230003131310-1221032210320011"></a>

#### `default_pool.origin_servers.private_name.segment.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0302313122223330-1121323302211110-2131323133310133-1100232201313322-2020202131302103-0012131323131313-3012003210310203-2002330233132302"></a>

<a id="canonical-3120110030110113-0011130101220023-1223120122132330-0231213110323100-3312300200321021-2132121223311130-0030113233311321-1101133222113133"></a>

#### `default_pool.origin_servers.private_name.segment.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0310320011032122-2311301301130313-0303200313201211-3022331013021321-0233302120202012-1122031222203303-0311100322233033-0301130233121123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-0033332223012301-2300330132120010-3103023213131323-0002200112330130-3333133021031020-1321320113030032-1312233312013332-3000321121121313"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

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

<a id="canonical-1222000330332132-3001020103302222-2112020332330030-3202102122211311-2113133332233001-0100300331221212-2032012230131311-1322102023312200"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator`

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-2213112212202231-0223110103000012-2031002123313323-1220313210301111-2200023222200000-0220110121010130-1231033323103310-3113330010212020): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-1020013321323300-1122202322201320-1000303300312023-3200112001012210-1321031131322323-3001211103302301-1320131102203300-2010012202032020): complete subsection reference.

<a id="canonical-2213112212202231-0223110103000012-2031002123313323-1220313210301111-2200023222200000-0220110121010130-1231033323103310-3113330010212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0310320011032122-2311301301130313-0303200313201211-3022331013021321-0233302120202012-1122031222203303-0311100322233033-0301130233121123)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-1130311033002322-2012222022132332-1202103110131220-2020302232330213-1131102200111323-1103030103303110-0222323231002111-1012030031011022"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1100030223012222-1233100013323013-3201211101021200-2132001030333301-2212112021200330-2321203130210223-2022130011013013-2002321311010002"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator.site`

<a id="canonical-1010113201213033-2231121222231121-1302102020012211-0232111112331313-2112212111210303-0112322022223312-1113302130012030-1320333222012033"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0312002313311232-2030310211031321-3121010210331230-0032131011133331-0001230200301302-2230000012323112-1020100302002011-1112322203130012"></a>

<a id="canonical-1333301112230200-0331320221332310-1333100021121300-2122232033121323-3221230230030202-2013012333233321-1203321310231300-2021102021102031"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2323331103230112-3321021231100330-0323233111313203-1320221230333332-1221311332102323-3013023101110202-1121331201220320-2332321003302231"></a>

<a id="canonical-3121010103110311-1020211101312013-0032222231210121-3112211301322031-2023123300203202-2302003003020132-2231320012213113-1023201333222020"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1020013321323300-1122202322201320-1000303300312023-3200112001012210-1321031131322323-3001211103302301-1320131102203300-2010012202032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-0310320011032122-2311301301130313-0303200313201211-3022331013021321-0233302120202012-1122031222203303-0311100322233033-0301130233121123)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0331003331313100-2302021320302203-3312302123022200-2321230203210301-0022220020000133-1322013122320221-1033232303101222-1210323003030312"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
