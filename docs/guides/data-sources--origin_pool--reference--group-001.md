---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310021212323322-1022103230011121-1231023130113301-1302023131331130-1132102312213021-1023000333022203-0202310231020131-0301003330332013"></a>

## Property reference — Property reference / 232011232202 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- Property reference

<a id="canonical-2133033201112200-3300322022131002-1220111303123313-1010211310312312-3222020123121212-2220101000033100-1312203133121232-0220112102231000"></a>

## Direct properties — Property reference / 232011232202 / 3

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130): complete subsection reference.

<a id="canonical-2010223130200322-1211322313021311-1111020200213230-3201023200131021-2322303330030002-2202301213213020-0230020232210023-3232111311323113"></a>

<a id="canonical-3230032001332332-1132103300330300-2321302021323311-0023000203322102-2111203233321200-3003320103003131-0003221202101313-3323222232212113"></a>

## annotations property — Property reference / 232011232202 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-3232200323013011-3012121113003102-1101211132122231-3303101001222131-3132321221221112-0012322000210223-2010320011332333-2212232121313131): complete subsection reference.

<a id="canonical-0021310000312203-0013233030301230-0023310022103230-1120232113102221-2030201302013321-1113303133301100-1320330111020311-2010130320012131"></a>

<a id="canonical-2033121331002221-2231202300303011-1103000232122221-1110213202131201-1203332220331111-2002103020221203-2101003102012102-0001011220320120"></a>

## description property — Property reference / 232011232202 / 5

Type: `"string"`. Computed.

Description of the OriginPool.

Upstream description:

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

<a id="canonical-2112103202122012-0331210012310032-2311122033101321-2110321202222322-1031331100322102-0223232012030320-0312031100103313-0012020110200131"></a>

<a id="canonical-2300123012320233-0320103001221013-2122223111211120-0302210232101102-2203221032311131-0220212333232220-1012211103303220-3033311101111311"></a>

## endpoint_selection property — Property reference / 232011232202 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1311220212320020-3111132100331120-0003010001013110-2003122232110332-2110123003232013-3222012200321112-3112233232203333-2300000232201002"></a>

<a id="canonical-3030323120222303-1222220020203330-2133030032230201-2023222032102012-2003133010123011-3021220220211132-2301113022332312-2100330113120210"></a>

## health_check_port property — Property reference / 232011232202 / 7

Type: `"number"`. Computed.

\[OneOf: health\_check\_port, same\_as\_endpoint\_port\] Exclusive with \[same\_as\_endpoint\_port\]
Port used for performing health check.

Upstream description:

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

OneOf alternatives in this subsection:

- [health_check_port](data-sources--origin_pool--reference--group-001.md#canonical-1311220212320020-3111132100331120-0003010001013110-2003122232110332-2110123003232013-3222012200321112-3112233232203333-2300000232201002)
- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-3230202001231011-0231003131211320-3300230323330013-3310330210031113-0323111321313021-1120103003020320-2233213133003321-0013202003001030)

Select alternatives according to the provider validators above.

- [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-3321301103022100-0002302311013132-3122331220210333-2233013212333133-1323332021333233-0320103330123010-3013120233010203-2201003110133022): complete subsection reference.

<a id="canonical-3323101103011031-1132322003301333-3310311111030120-1012233131023000-3330033133212011-2000202031333023-3213233112130000-0220020202030133"></a>

<a id="canonical-1321022121310233-1223211313213201-2320102133313111-2120320333120023-2221011332330123-0301222302130130-1132212132001221-3232003111313210"></a>

## ID property — Property reference / 232011232202 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2000122330320302-3133213012000111-3311302221021100-2301001200322333-1233330022001131-3220011310032022-0031332103121033-3111010330302103"></a>

<a id="canonical-0302001022313210-2211011110321032-3301313020132102-1312332103310112-2320231022130110-2312202321100023-2232223111112132-0002202220203232"></a>

## labels property — Property reference / 232011232202 / 9

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-2321331312131123-3012123331102100-2201101123302212-1013031113132232-0023332220033133-3221300001120322-1131200220002203-2121000221033310): complete subsection reference.

<a id="canonical-3033223001320132-0000103210110031-0333011222123132-1323201300212021-1113010303132102-0320103221122002-0330320220120201-1122301010303301"></a>

<a id="canonical-1102130210210321-2023103230013002-3221123202021200-1033132031311320-3113012001233021-0112231313022231-3113112013101003-0231300001000112"></a>

## loadbalancer_algorithm property — Property reference / 232011232202 / 10

Type: `"string"`. Computed.

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

<a id="canonical-0002121323113202-2232013120002221-3000013202310033-1310230032033031-1333000110212313-0202210333131310-0300113023223030-1101301112113002"></a>

<a id="canonical-3302220200232103-0210321303103101-2233222333100122-3310200011223122-2210231003212011-3011102011212003-1331111031002002-2301122032230210"></a>

## name property — Property reference / 232011232202 / 11

Type: `"string"`. Required.

Name of the OriginPool.

Upstream description:

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

<a id="canonical-3012022133132110-2233331310301310-0302303300133220-3200112012230010-1110313331232003-0310122111201313-3023211223121111-3301111200122110"></a>

<a id="canonical-3233211110203202-2112023013332203-1022332320220233-3112103302103103-3002000100330123-2032332110102230-1230011313111223-3233002110000033"></a>

## namespace property — Property reference / 232011232202 / 12

Type: `"string"`. Required.

Namespace where the OriginPool exists.

Upstream description:

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

- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-3003002212002303-1310102203201220-1232202311101013-2320113213310102-0313212211221112-0022010223033133-1302122300202210-0233020303132231): complete subsection reference.

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211): complete subsection reference.

<a id="canonical-1331233022212123-0232020012320101-0313013123210003-2023013132030210-1213332332113330-0220222032331222-1133021111301032-2011132211203233"></a>

<a id="canonical-1332201021130031-0222122221030020-2331202033303233-3103012332100101-0021231103102023-0011203330130323-0322203302012313-2232313123111123"></a>

## port property — Property reference / 232011232202 / 13

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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

- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-3130012011302310-0100000020300011-2333213101123221-2323300231131231-3230213133122123-3221221202130031-0122330100113130-1102320312113033): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031): complete subsection reference.

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301): complete subsection reference.

<a id="canonical-3112121000310220-2001100310321310-3320320003223121-2013123120332200-0330300031213122-1110121223230111-1121211002011313-0003300031013322"></a>

## All schema paths — Property reference / 232011232202 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-2033231303001123-0222123202223310-3321131000020131-3121222301003020-2032103133233130-2203203323002113-2303000301312103-2200301121001123) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-0023303322122330-0110201011033232-3102231020133113-3122301100133200-0231120303001121-3111203303331233-0201212231123221-2200101331233130) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-0201310102212200-3201020120222301-0323301123232203-3103332113133313-3023110123312213-3101100121230101-3311000111133012-0122033013133331) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](data-sources--origin_pool--reference--group-001.md#canonical-2033013232330200-3033032001001310-1202012030230122-0101333330121102-1123202121020011-3130322202101010-2001110323023001-1100303000300100) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](data-sources--origin_pool--reference--group-001.md#canonical-1122322311330123-1211201200303032-2110321131103310-3203023312033121-0023230321012013-2102031013001131-3312130233212200-2301133013310233) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](data-sources--origin_pool--reference--group-001.md#canonical-0302312112110122-3111301320001222-1203011133012212-0012301320303331-2133320222203211-1131210210300011-3023112321302201-2120222311232320) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](data-sources--origin_pool--reference--group-001.md#canonical-3332120101123321-1112320230310303-2023303103033332-2232010202202131-0000010201310203-2000112010110313-3311233000212200-1022302022211110) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](data-sources--origin_pool--reference--group-001.md#canonical-3322131120003331-3300031133303232-2233322213132202-1100203220122323-2010202203000301-3332023103231320-3312121002131322-0120121310220312) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](data-sources--origin_pool--reference--group-001.md#canonical-1313301212211312-1203221231330311-3021213133321323-3300030300113211-0310322021121232-0020300322022212-0113220313221330-2313230312020330) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-3003210100122113-1110202321020020-0020131122311320-1031101122321322-3322303111123000-3012210231132200-3033322102233203-3332012001213103) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-2103331103012022-0121331233320100-1020203320112232-3210223322111113-1213031101110022-0101012332203022-1213333311123221-0122210332123102) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-0223101220202022-3201110002131132-0122010220112231-3301120301221001-0212113002220300-3033312012320010-3123220031212012-2022002321201222) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-0132022131122020-2010013331030323-2310300321310021-3031032022133033-2331022203030000-2032221003120202-1300221212301210-2130313122203131) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-0311111311310132-2011221212102202-3222033123112001-0302100021321331-1102102210211332-3311002212032032-0110103332330213-3130103121032333) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1113202310302103-3013020212022223-1020030320331323-3002322130230202-3023020131022313-1113033303332233-3133230133010330-1121130210301200) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-1013230220212023-1033133123012201-3121301321333211-3200232311033122-0022123203332110-1012022202223230-3033302331301232-2012300313112002) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-3300120222132113-1333330221301102-0131331033332022-0221221121232101-3230123333212213-2030101030302132-0302013221122033-1100302213300000) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-3101222313123233-2330322022200332-3121011001200233-3121102123013230-2222311033021030-0330011322303113-1131020133130102-1211301210101303) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-1330213222300200-2130012201301121-1320230232022300-3132112103222300-2032121013021020-1223021000322323-1003300110023223-2213120213222220) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-0310233311020332-2331022310012210-3113102110100020-2331122310132122-0222300222201222-2312331110232113-3012210233111230-3310132203201102) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-3001123213133310-2001013202331113-3312010310011100-1031103222311313-1131002321233123-3011112112211130-0302312211202013-3001330321030331) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](data-sources--origin_pool--reference--group-001.md#canonical-2322123120302300-0211003003203231-2312033221013001-2022220311213222-1010202322310003-0230131002011212-2000023303020033-2122120213133111) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](data-sources--origin_pool--reference--group-001.md#canonical-1222310232302021-3020123031233333-3032032201331123-2130323000222032-0111132203132113-3200212033030221-0101212003011331-2330311102310023) |
| `advanced_options.http1_config` | [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-3011311102322130-1323231223210330-0122202001203201-2003130102112122-0212332013131102-1202203000221220-3023222020201101-2200333303210233) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0223111232112221-2300021320012132-2323122301302123-2131101233312010-1302312101002332-1332031301023002-3113013330131321-1222333000103013) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0313213332222132-1003002112123213-1203100012030113-0100031311000123-3133211201023112-1312133322300321-2311111030100023-2102103100202002) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-2220101031321301-1310012031011300-3012132203210131-0011000002010322-0302023213313231-0323100113332312-1300312010133320-3010001100323133) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-3000232000131123-2231230033310122-0013001113112331-0211303013030230-3133232030111120-1122311133222312-1022213232232321-3003311100233313) |
| `advanced_options.http2_options` | [advanced_options.http2_options](data-sources--origin_pool--reference--group-001.md#canonical-0030122322133321-1301003300103003-3301131221031302-3010312203022202-0300230022233012-3131033313110200-2112333103220030-2031023320121022) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](data-sources--origin_pool--reference--group-001.md#canonical-3302103020113132-2330302202103300-1130023211332221-1313133110113012-2130030302002100-3112230130322001-0122331310312323-0313133212132120) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](data-sources--origin_pool--reference--group-001.md#canonical-0001321102101331-2130031230130033-0330200313111310-0103033010122232-1130101230122231-3233131231312130-0120021303012221-3100200220120203) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-3031031223322031-1211103131131121-1023201322100221-2332333233111322-3231323203213022-0131230112330300-2213131112033110-1130220331203213) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-1313132231132013-0133211333033201-2033001300013312-0223210103313312-1000310003033100-1331212301202332-1202221113102231-3101031301202320) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-1330331121013031-3322030210320321-3103131130331333-3213212002231222-1232120213031333-3100313112320100-1100300012212112-3102333032210223) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-0120230001000133-3203310223211112-3132133011001011-2002232100121213-2202003300213112-1023012233021110-0302231011331111-1213103310302231) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](data-sources--origin_pool--reference--group-001.md#canonical-1032230222031213-0223102201311333-2212313202302122-3320000333200121-2123011222133123-3312103113232312-0212023201333121-2223022103132121) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](data-sources--origin_pool--reference--group-001.md#canonical-3310110022021010-2211033323122201-1232021122003332-3100101022120222-1102330102123003-2230032122000211-1322200131210013-1330001100012301) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](data-sources--origin_pool--reference--group-001.md#canonical-2021221200003200-2123031132201220-3102221102300311-1332110030330122-3220002022233112-2311302312221033-1222231120222020-0113012320312133) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](data-sources--origin_pool--reference--group-001.md#canonical-0113320002320321-3020322133233101-2201201130231111-2230033232130112-2002122123201212-2000201211100322-0223320302102311-0203312201333100) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](data-sources--origin_pool--reference--group-001.md#canonical-3211300220032332-2220121022233030-2302332031013213-2322211311310211-2210211111312330-2032230032123301-1102201323103123-1300220220332201) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-1030100211333100-0033323220102303-2223000021212002-2101320111130322-3111003021010300-1311123303222021-1132112312000103-2032103121320233) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-3131021033021122-3332300331321230-2213310130323231-0332020030322311-3310112321122301-3023232230203212-3210233003132203-2111213220202321) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-3303032000000002-1211212303130132-3133011323011230-1210302310232032-2203102030322323-0211220011222302-1322211323103132-2010322102210113) |
| `annotations` | [annotations](data-sources--origin_pool--reference--group-001.md#canonical-2010223130200322-1211322313021311-1111020200213230-3201023200131021-2322303330030002-2202301213213020-0230020232210023-3232111311323113) |
| `automatic_port` | [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-3330030101213132-1023001113001113-2323003003322122-3003122331131110-1132010030223132-1131332022230111-3132311313331113-2130313202000312) |
| `description` | [description](data-sources--origin_pool--reference--group-001.md#canonical-0021310000312203-0013233030301230-0023310022103230-1120232113102221-2030201302013321-1113303133301100-1320330111020311-2010130320012131) |
| `endpoint_selection` | [endpoint_selection](data-sources--origin_pool--reference--group-001.md#canonical-2112103202122012-0331210012310032-2311122033101321-2110321202222322-1031331100322102-0223232012030320-0312031100103313-0012020110200131) |
| `health_check_port` | [health_check_port](data-sources--origin_pool--reference--group-001.md#canonical-1311220212320020-3111132100331120-0003010001013110-2003122232110332-2110123003232013-3222012200321112-3112233232203333-2300000232201002) |
| `healthcheck` | [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-1013131122022033-3120111310102210-0122223231303100-1002203002313300-3112312220330112-0123221302122033-2111003202132002-2200220232301001) |
| `healthcheck.name` | [healthcheck.name](data-sources--origin_pool--reference--group-001.md#canonical-1020202033220110-3030301031212233-1331211302231211-0022012012313121-1313000312202122-3213323223122230-0013220011313121-0320012130000011) |
| `healthcheck.namespace` | [healthcheck.namespace](data-sources--origin_pool--reference--group-001.md#canonical-2132130030010021-2022302112231131-1233200313111231-3221201011001111-0220233100220123-2033332301330200-2313220132112001-0122122033122303) |
| `healthcheck.tenant` | [healthcheck.tenant](data-sources--origin_pool--reference--group-001.md#canonical-3333122232012021-0202223110210203-1313302332213200-1122211113331230-3221012013200330-0102023300002010-3111010311122323-1033012331310313) |
| `id` | [ID](data-sources--origin_pool--reference--group-001.md#canonical-3323101103011031-1132322003301333-3310311111030120-1012233131023000-3330033133212011-2000202031333023-3213233112130000-0220020202030133) |
| `labels` | [labels](data-sources--origin_pool--reference--group-001.md#canonical-2000122330320302-3133213012000111-3311302221021100-2301001200322333-1233330022001131-3220011310032022-0031332103121033-3111010330302103) |
| `lb_port` | [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-3332320330310202-3001331103323210-0100103311013120-0002100310001103-2000333310000320-3010022210312130-1113032332102000-2210230130200022) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](data-sources--origin_pool--reference--group-001.md#canonical-3033223001320132-0000103210110031-0333011222123132-1323201300212021-1113010303132102-0320103221122002-0330320220120201-1122301010303301) |
| `name` | [name](data-sources--origin_pool--reference--group-001.md#canonical-0002121323113202-2232013120002221-3000013202310033-1310230032033031-1333000110212313-0202210333131310-0300113023223030-1101301112113002) |
| `namespace` | [namespace](data-sources--origin_pool--reference--group-001.md#canonical-3012022133132110-2233331310301310-0302303300133220-3200112012230010-1110313331232003-0310122111201313-3023211223121111-3301111200122110) |
| `no_tls` | [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-3003013113222220-2211123101222221-3301200210020030-1013302300212120-2102113320000233-3323230131232012-1312300231010103-1233101032222222) |
| `origin_servers` | [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1202000233133133-3013023232001221-0213032112311332-2212331201233022-0011000211132213-3032110030112230-2120311230221030-2001112121001331) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-0233032331300323-3222131020331213-1312130022210001-2212330232123302-1033120331010313-2201020010203311-0001110000200022-2101001211221212) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](data-sources--origin_pool--reference--group-001.md#canonical-3333012200100220-3211011203300311-0023133011123002-3000100012303212-1122110200000110-2330203112020310-2021320000210013-1232111223310102) |
| `origin_servers.consul_service` | [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-0210310100023010-3001201121100323-2233111213111200-1110031130123110-1013203030001320-0312331000200223-2023121102221301-0122302120030023) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1320231211131113-0312330100313120-0000133233102123-0232312210131300-1202121323203102-3332212312233131-0332223332023321-0113223122313111) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2001322200002122-1231301123332322-0300121011121113-2111202032020310-2100131213210122-0221000112331123-3302031003120303-1122102210222113) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](data-sources--origin_pool--reference--group-002.md#canonical-1300121122331122-1123001012133210-2000202031321320-1323001021210011-1211001310113113-1230021010000203-2312313322320211-1112211122322231) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2201002203020211-2222200021030213-2303312333023310-1303230031112302-2123323031022112-3020032000231233-0020323030031323-2211133311111301) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-1313003323320102-2303213312302230-3302110132132131-1110022102103022-1203310010303222-3230231330220332-3330203002223202-0232011311001021) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-3000122233210221-0013332312311303-2001121210122023-1230331330203320-0111323012222032-3110013000332120-0103221311011303-2032122120322030) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-0003010221112231-3301122210021031-2311333302322103-0223033330010021-3021221110202200-1101302222200231-0033200320333131-2011223230231030) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-3211212202030121-3200311312212202-3232102102133030-3003213201311312-2301321112113100-2110002312012113-2003111200233312-0013002212013212) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0020302332213002-0211232210010332-0210230023321103-3012221020333233-3021321021313323-2331131012123100-0111302121230001-2322121202303023) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-3311231130121310-1113010133301031-3110023330122112-0011201101203001-2212121300210130-1022032332122003-1223013001021303-0331233101011330) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-0323311030133302-1020031231202313-2300133333321203-3012101332012111-0101001103021222-3302310211232223-0302200110133302-3012021120303010) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-2031011002312332-0112122212112020-2021313002121220-1231113333220302-3002100213102213-2023133130021302-1113131023321112-2031000313113103) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2002312100030302-3002022232311220-3301132331131212-0010310320111231-1113213010131230-3023023210130222-3222233301023132-3203022211022202) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3202201023102303-2012010023330120-3312200223102300-3321033012123312-1000320132320311-2203022332211232-3230313130120111-3031211232132110) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3022220313021103-1130003021313320-1220333003113330-0332112323300013-1321022310011020-0200012202331312-2003202332120302-3203310202303332) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-0210132323210231-2130202320012330-1103212300013211-1303200213300003-2011130211020101-2113031321300113-2113133111232103-3113201213020222) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3333021302313233-2132310320013112-0113132311331331-2123210023300020-2030301130113303-3310200033000000-3230303211313120-3220310031310203) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](data-sources--origin_pool--reference--group-002.md#canonical-1213110323211231-0333323132100100-0111033023230312-2120123301202333-1032110200313211-0223020233100211-2322033020203321-3333100211133201) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](data-sources--origin_pool--reference--group-002.md#canonical-3131322332321103-3233133102210301-3233120011202133-0002123111003220-2322230232312331-1201201301020111-0130232021332013-3232031000112323) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](data-sources--origin_pool--reference--group-002.md#canonical-0320122000303100-2211122300113100-1231000112232003-1100111233232010-0300033333232333-3120330230313123-2223100212332311-3320221232202122) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](data-sources--origin_pool--reference--group-002.md#canonical-3133030202000321-3301010003320331-3110322022313022-2001231320013112-2132002020322202-2301020022310311-3003201201021101-0311020122320122) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-1223233020022101-3220211210102003-1322331302233310-3210312321210332-2022313212330223-3032211030303222-1023120103333200-1113332132013312) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2203311331100022-3222230321303230-2120012031212131-1200031230202303-0301032213011323-0233003220100333-0111101122312331-1300222230302132) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-3012021223332031-2331031010321322-1321023001130222-1110130231311222-1030302202321020-0010220322233210-0200002000311313-0030121002132231) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](data-sources--origin_pool--reference--group-002.md#canonical-2033110102321231-2011330302322202-1233110020223330-2221122203002103-0300002033001031-3320102302330233-0332002101233021-3003323032022011) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](data-sources--origin_pool--reference--group-002.md#canonical-2203033320222001-2103023111123222-2123233202333022-1302301030303113-0032310111031133-3231203233112102-2313112330103011-0010231113212013) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0022112310330022-0110101102111300-2022230203223313-1021300231102100-1320312212011221-0220311212012120-0021321021212233-1122023130231201) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-1132230003030301-1031133202102201-3111000112233130-2111001033320033-0323131211130233-2231310022112003-0220312223332322-0203131100022311) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-2020201212221111-3011330313010203-2021202000321221-1103133213210120-1123021302330132-0030122133300010-1112113130033231-0132212231001120) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-0202233210233213-0013230131322323-0121011223032222-3012120221000302-0120002202022021-0130001311111213-0123123020123322-1200003023220321) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-2222012313310210-0022032013100010-1211230330120201-1111302301322301-3113122201331221-3001023023003210-1031121200020003-1131223333333211) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0310032322220213-0023320311333122-3300011313031132-3332003333030021-3202203330303110-1320232211123233-2113232133332122-2302112132103200) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-3121121032012310-3323221103112303-3033102120210112-1312011301331103-1011030030011113-3313210023300001-3311121031021013-3333031222000131) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3100000122100212-0311032300110231-2102310101333300-3101201323130231-2330121223111103-2020301113203031-1122031231022020-1213022011303232) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-1030131222233221-0131322020130322-1013021132020111-2133202323001322-1101221232121133-2021213021212031-2033130130231003-0232303312300202) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2300132112011133-1001222102231111-1021033011211330-1013213202333130-1031330203222122-1031221000101110-0232123121022112-2111300211032002) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2112133302313122-0101020213222322-2101220231122310-1330323112021031-2030200331033020-1031013313100001-2013232120021223-0320212020313322) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0301323033010103-2011122000221231-0203013123323002-2111221332012030-2323013333123330-0303301231023202-2322322230201030-1122122000331310) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-1200301233210100-3313112203023312-2100103121232311-2202010012221101-2330320031030010-3031220212201302-3132221301132002-1030022001210021) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-1121320221330002-3130133132310313-3123111112030121-2200020223033312-3302202200233301-1203100203000022-2113333233111312-0022330210032111) |
| `origin_servers.labels` | [origin_servers.labels](data-sources--origin_pool--reference--group-001.md#canonical-1102110302003202-2233112110210002-0203310333030000-1232303012101122-3200332233322221-1200132132012213-0003110321322113-3113303101332101) |
| `origin_servers.private_ip` | [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1323132032102301-0213302210220300-1020121233103023-1020311222111112-1021313101312113-1100120021102132-1233223302202231-0132123313103321) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1332333132310022-3003132113211323-0322120123232001-1331322020021023-3322100031003020-3013333211202122-3030113233110100-1321000310023030) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-0320121111021001-2203023121330110-3022120020120210-2221003230002312-3113001022013212-2213122121300300-1322013112110201-3332221331003322) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-3031033110331132-3203320200310102-3330333312113313-0203111020203111-0102332121203313-0001011113021332-3132111312320230-3333102001033230) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](data-sources--origin_pool--reference--group-002.md#canonical-3201111022301122-3212230121302012-1222322302333010-2301231110023000-2310203230000220-3201031222233121-1110110313332120-2123331211011100) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](data-sources--origin_pool--reference--group-002.md#canonical-1121332010322010-0331001213013201-2131310002211300-2222320222312010-2102333233133131-1130310232202313-0022013011121321-0310220023123313) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3302312100203322-3331010132203231-1003021303133023-0020113022020203-3011220013333200-2233100110332001-1330202010303030-0202121321332011) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](data-sources--origin_pool--reference--group-002.md#canonical-3233332010001012-1110311120001213-3120222011223023-0333222232121211-0000220033100120-0221202320211231-2333003012020310-1123022233113212) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1122123131011031-0313203023211013-3301232131102222-0032133132333001-3213212231100211-0032200303000013-2221221320000103-0130032103113303) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-0210110200112220-0223300031210231-3200011201331330-0013221232313333-3331130210310133-3320221312023331-0301221002100220-1112123322000300) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-2011333133310033-3111112013320011-3311200323000133-0010313130202330-2012112031033000-1110111312320022-1222221033022123-1112023023123102) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-1222002321100201-1101201301133331-1032323133222112-0213211133312233-0022003212323222-3312033300302221-3000210221322310-0302320302222200) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-0012123211112011-3230103211311313-3213113312212130-1022200021003102-1320213220020100-0332331223203301-2221122022102330-2203021222331130) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-2322223312023130-3220100111322322-2331222033031323-2020102013312300-1021033311223033-1331102312103133-0121121111322321-0311300313201102) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-2122220122323033-3011313232113110-1312231103032221-3021110310220201-0203022032302212-0313101011012210-3311303333111112-2111311031003330) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3010330231003330-1332110232022312-1132231023111310-0013013003321012-1101130010312333-3221001211232020-0131332221213310-1300220133211200) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-0130000200222002-0203230011223020-0103030211123323-3122302010302333-3200032002020011-2032123202321320-1311320230023023-1120103102102012) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1230021313222323-3120023031131200-1233113301020201-2232222010011223-0032021312332211-3100132312300121-2202330322301313-1303211031130322) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0003302321133221-3130232212332103-1122210001002331-0010132303311003-1020030123022012-0012330002212132-2232113131203303-2021031103210122) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2232333332130303-0330120201011332-0022211113221003-2332200023030010-0101132031231212-3132020113213211-2110030323233031-0133030103013230) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-2223130122312001-3221130033133232-3002313333331210-2202010202231213-0300330220003120-2310133320022001-1202230123122322-2301032302202013) |
| `origin_servers.private_name` | [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1212131321023331-3323221213100213-3113030213130331-1022001321131223-0300113331232220-0221212323102122-0132231303200201-0101122223310000) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-0033110221311322-0232002212203030-0222030202220121-0233103311122200-2321020110302131-0120212321323311-3003313013211121-1110232232323213) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-3312013330201201-3132311111230322-2232223022212333-1013203331200012-3033300123203022-0000032033020033-2030333301231201-1330312212233300) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-0113310102130113-3100201313013101-2030200220102110-3133110031312120-0122202002021032-1033001331120103-3012332133022000-2321311332323123) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](data-sources--origin_pool--reference--group-002.md#canonical-0010301302220223-0122112222213313-1312231113030322-0333132012223023-3011222311030123-0201203210113311-1022202231103030-2102203101220321) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](data-sources--origin_pool--reference--group-002.md#canonical-1221001011021022-0033011011130002-2021222002011013-3120212333000223-1003102102031003-2130213321103222-2300300011213110-0222103112022322) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](data-sources--origin_pool--reference--group-002.md#canonical-1033202133100110-1322011323001330-3121210233103123-3000223222300302-0022212132321211-3301003002033303-0331130131221200-2100313113312010) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](data-sources--origin_pool--reference--group-002.md#canonical-2321123322120313-0312230231131300-3203012323333301-0203313112112332-1323113320100331-2022010103130201-2210113323133033-0012221331203010) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](data-sources--origin_pool--reference--group-002.md#canonical-1210202331220112-1210321233303002-3021313223331210-2311030032232011-0032010311333002-0010132130313133-2122303313121013-1213230132101223) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1323033030323301-1000213232102213-1222113200010031-2233202002201211-0312021122012220-3011323321322103-1222233021120013-3212331312020123) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-2122103123032102-2131333320013322-0213003013010301-3310201212001002-3211000232030113-1101112032102111-0331323101012220-3003001010121201) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-1211233032113220-1001202021313232-0311202130333311-2311300302311102-2231311323110320-0001210120300302-1013101012110102-3130110300020010) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-1120032123131211-0002033102121322-0331300123112000-1313233111010033-3111132113013230-3102001012122031-1120231131332003-3110312321231120) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-1233203330101211-3310220211203310-1331130003213323-2010021330032230-3012231023201013-0111020013232323-3231231320110022-3131321310111031) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-1030132131011010-1212123123013211-0221012011230212-1223121332332322-1032232202022121-3011331310213102-0103033321011323-0012301013230133) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-3110033030101323-3233302100100230-0200301100201200-0131232232012032-2230100221120101-3122323311002221-0222300103032322-0222100011303122) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-0333221002022120-0210210333101213-2011102100013021-3230220323022303-1111212231133021-0232103021320203-0231313231220120-0333233121230300) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-1200122031023023-1331310113102020-1301021313002211-1130213000132332-1322121113230110-2320102132233220-3013232220001133-0232313131000023) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1331233212211320-2330123312023213-0133103211121212-0120101132022101-0030332333002120-1303321100121230-3122021023133032-1301312121300213) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1231201333130132-1002102100110300-1031113030320232-0200000330233032-0012221031123231-3120101323313200-3122130021321332-3203010113331223) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2000120023311112-2013031201210131-3020110013222322-2120332123222321-2232100202201123-3313300131131100-1201032210210303-1012102101332300) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-0233103211231133-2110222003232101-0320133301200002-3030012213211320-1033011002033131-3300111023323230-1211310033212312-3120300302301312) |
| `origin_servers.public_ip` | [origin_servers.public_ip](data-sources--origin_pool--reference--group-002.md#canonical-1123110022211300-0110332003213013-3212002113001201-0213030321132330-0010033001123102-2130022002302130-1032031203320030-0321203102322033) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-2112313312012323-2310022330233310-3112311113003102-1303230123010002-1331011111322201-0333202020101120-3201102313122030-0100112212232122) |
| `origin_servers.public_name` | [origin_servers.public_name](data-sources--origin_pool--reference--group-002.md#canonical-1233231001013031-2011111132012112-3312030331030032-0110110203203310-3233312102102323-3130123013132101-1330232321023212-2230100130011023) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-1200212100313122-0230222122110333-3032013211322133-2223123001330111-1103312130303231-1030200121030133-0301323223203322-3302113023330332) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](data-sources--origin_pool--reference--group-002.md#canonical-3222221303103003-2201330000303113-2100313002211111-0331321033003220-2130311010103001-3200120333100130-3323121003230033-0023233130303021) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1201112113021310-0222213200333021-1033212320200202-3230321213113231-0133112033123230-1022003020032320-2122211222312012-1100330331002123) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-2003223332223001-3232311123032310-1312001010002300-2030002132223000-0202011113321011-3220300313123232-2312120221301013-2021130320222030) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-0002021132223321-2003211320303002-1103121321312311-2223210300112331-0300321112200111-3133303122132323-0323330102103021-3030211012303332) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](data-sources--origin_pool--reference--group-002.md#canonical-1012013330302213-0231110132123330-2133213131023301-2023331220323003-3232332330210031-2130022132022322-2300311020013000-3203103032110201) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3231102213103123-3000012112222310-0203210022120331-2322102213111103-0200311221132311-3320010012311030-1121210020030313-2020301002103123) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](data-sources--origin_pool--reference--group-002.md#canonical-3312222332103020-2012233203111323-3030232021022110-2312330121102231-1313223332120003-1213102110130332-1022201211112203-3223201102303300) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-3312300023020300-2221333211123333-2330333201320022-3320130123211313-2112111202221302-2223111321301023-3120123230023022-0031013011113231) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-1101311331303030-0032012310223201-1030113001021100-3223200320231312-2300221231321110-3211030022002012-0323333220300300-1120312102213313) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](data-sources--origin_pool--reference--group-002.md#canonical-2022001023201223-1213310013210002-2333312233011020-1133010220111313-0232120012022232-2322002013233023-3321233023222332-3030101333030010) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](data-sources--origin_pool--reference--group-002.md#canonical-3131030122030032-2223021030311310-3020011103310230-0100010132233300-3332032022303322-0313220320031220-3311102032113020-2333232131200011) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](data-sources--origin_pool--reference--group-002.md#canonical-1100122120222203-1201312111313100-1320332211123330-1023231021110333-2322133023220030-2020202003310122-1313211032121220-0133020330002110) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](data-sources--origin_pool--reference--group-002.md#canonical-2122133133222333-1331222130230200-3020130132330020-1121001131220300-0230331012033311-3331133333031012-1012302232023131-2121122220201202) |
| `port` | [port](data-sources--origin_pool--reference--group-001.md#canonical-1331233022212123-0232020012320101-0313013123210003-2023013132030210-1213332332113330-0220222032331222-1133021111301032-2011132211203233) |
| `same_as_endpoint_port` | [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-3230202001231011-0231003131211320-3300230323330013-3310330210031113-0323111321313021-1120103003020320-2233213133003321-0013202003001030) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-1223310203212231-2312201220222003-0032121130323132-1001322123023120-3030102001303133-0131203102310120-0233101220301003-3210220300301332) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1221113302121301-2313020121331320-0221203333123323-2123200223300002-1330031110322133-3313200133100312-1310111322321002-2123221332010213) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1303312032131320-2021212201121011-2021212032121302-0213111331132120-0201200332100220-0201322223303113-2130120213131233-3303120103001132) |
| `use_tls` | [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-0002222103101132-0331331011030000-0221212201230213-1311030121013132-0223020023313332-2311112031102200-3132032023331103-2122132001320021) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-2322202020222130-1232110023322133-3333101113001113-3210232200231331-1231120123332320-3010301033212312-3103012211310233-0311213210031323) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](data-sources--origin_pool--reference--group-003.md#canonical-2003010311021311-0332020301111223-3123232101102112-3103113001010103-0333200233003302-2102011021303311-1112312022322313-1302011130230300) |
| `use_tls.disable_sni` | [use_tls.disable_sni](data-sources--origin_pool--reference--group-003.md#canonical-1212122200212010-3013213213022221-3021312132133231-1021313013010111-3031231222111033-0211300322212000-2333213222123033-2220012003130200) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](data-sources--origin_pool--reference--group-002.md#canonical-0233103101220320-1020311303110113-1100311111211321-1231332102301202-2320032130222232-0033111113112220-2100121012331321-0130112230133010) |
| `use_tls.no_mtls` | [use_tls.no_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2121302331023120-3120231203230320-1111220312012332-1030122211200213-2330101321313131-1023131222221102-0300313213311031-0111103120033123) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3331020311210023-1121220230320302-1302202211103233-2310322312021232-3132220231332313-3321331123311103-3200303312312113-2001130112031321) |
| `use_tls.sni` | [use_tls.sni](data-sources--origin_pool--reference--group-002.md#canonical-3313012030230223-0020202211000122-2020100322223211-2300020233303232-3202222230130322-1233213120013233-0002012022211210-0321132221022123) |
| `use_tls.tls_config` | [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-2221121021212003-3311211133031031-3322231011231023-2021100112101121-1130032323313321-1022313211131100-1033020031230122-2332321020033203) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](data-sources--origin_pool--reference--group-003.md#canonical-3103212201032200-3132123333012101-3303303013300030-2113221030022312-3111111231031111-1130120113103330-2101111111302113-3003322113131103) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](data-sources--origin_pool--reference--group-003.md#canonical-1133310020130221-0011030031032112-0230300021102011-3333000223122333-2102211003312202-3033132231022221-3013102330002300-3113023000223230) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](data-sources--origin_pool--reference--group-003.md#canonical-3323313112120033-2112103233111132-3132310012221221-0100020120032033-1221231132003020-3012032210330030-0023120100120310-3223220112113200) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](data-sources--origin_pool--reference--group-003.md#canonical-1222322212103203-3310011001332312-2233323112233022-0232112103330322-3221101221022300-0132231221031332-3032330132002131-0012002231200003) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](data-sources--origin_pool--reference--group-003.md#canonical-3011131100330232-2331223001030001-2100020233001001-0033112210302331-3123310210321131-3023221333003230-3102230110302311-3332321000321101) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](data-sources--origin_pool--reference--group-003.md#canonical-3001030021000000-0123220133112320-3002232333203323-2120131120102120-3230310331232212-1211131130102033-3302113330313330-2202310312031031) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](data-sources--origin_pool--reference--group-003.md#canonical-1231213011331323-3030123110201313-3103022200330311-3312021113003211-2123033032331133-2323011323210000-1103310323202321-1132012001303232) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-3002002300202330-0332220022012011-3022022020123221-1033202333233112-3210233212112023-3210131231311110-1212330303233012-0223130102000013) |
| `use_tls.use_mtls` | [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-0033031012120110-2120022223112020-3023203121021012-3023302113320212-1232031232003020-0312120013033313-1003003313101323-1112303001113332) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-2202212233123011-2303103030233031-3000003203010000-2300010200213231-3013301111123220-0312101101002122-1230132311310123-1321203332102332) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](data-sources--origin_pool--reference--group-003.md#canonical-1223331010111320-0112301120111003-3312112000101332-3233130011303312-1211233011232113-2130100220213221-2113133230300321-2310100030001012) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](data-sources--origin_pool--reference--group-003.md#canonical-1220102320202303-1111031101221022-0202303032210332-0131203011230003-0113113012202331-0223020030013103-3123203223302003-2320000301233032) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--origin_pool--reference--group-003.md#canonical-1220012221230213-1013230221122030-0112110300201302-3132001013300213-0232133123213203-1230022011111103-0323120103032330-3033101021002000) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](data-sources--origin_pool--reference--group-003.md#canonical-1031213332303213-1220020112222332-2300302013112023-3012013211023210-1023331301303220-3232002013123311-3121321102320012-3322122223210321) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](data-sources--origin_pool--reference--group-003.md#canonical-0110213323332312-3031023021013130-2122032320303101-3032133230023111-3223111320302313-1332022212313312-3230033031310223-1300123101012211) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](data-sources--origin_pool--reference--group-003.md#canonical-2301130302322321-3131122203022311-3103012320011131-1202210201223201-2202202203122120-0030102031112032-0332003000032301-2012103133130013) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-0023300230121122-1013012102031220-0033103301023100-0321233132332203-3332230001301100-2101012100113110-0333113013221301-0222223230111130) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--origin_pool--reference--group-003.md#canonical-0301303131210231-1022000223200132-0001121001100210-3222333213221002-1330330323300021-0331311003102212-2321223003220232-3321221223012130) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](data-sources--origin_pool--reference--group-003.md#canonical-1000001120232231-2000101030032122-2203213222130300-1230232222111222-0221012231233210-3113302301013310-2311302213223120-2023102220021333) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--origin_pool--reference--group-003.md#canonical-2300023330023010-0210211132203102-3013321003332021-2323201003212200-1130200321302102-1201030031313221-2112110300103320-0202203123212301) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-2303113123201301-2030130200000221-3012311131022123-2330101110113102-2023122210031212-3100000212000200-1302232322201302-3131210211211130) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--origin_pool--reference--group-003.md#canonical-0110120331231132-2020230322002312-0300020030303202-0201330303232203-3233221332133102-2232212310200111-1230221212331322-2123032321031122) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](data-sources--origin_pool--reference--group-003.md#canonical-0232203120222331-2123332013313223-1031102202300012-2101011031220220-0213203110223213-3033322203111102-0122103332112032-1331201233302323) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](data-sources--origin_pool--reference--group-003.md#canonical-0031310212103130-1310313100023112-3012131302223101-1321300211310021-2300130120123321-1010223012202000-1113010232212021-3231031110032022) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-3221232002222233-1031100031012131-0031311201022132-1102211130230203-1110312313123220-2203333133012202-3120212111302220-0210312301230002) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](data-sources--origin_pool--reference--group-003.md#canonical-0131030113033223-1123123110001032-1333303333212002-2131322113223303-0201312222332331-0222001100132022-0103320010131001-2022021003103010) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](data-sources--origin_pool--reference--group-003.md#canonical-1100220001210103-1322301223000031-0030301333010321-0000201011023233-2122131202231120-2010120001133301-3232233000220220-0122212310313223) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](data-sources--origin_pool--reference--group-003.md#canonical-0300310203232220-0033212011213220-3023322010122113-3231202031203032-2320330230102210-1031213123133131-2032100002222201-3013123300233213) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-2332232310221103-3020011210223113-3010300002012100-2220313031000030-0330012230320200-3300333103111231-3201321020123012-3101212223012013) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-2333330121103121-1211202222332020-2220102122231311-0133230102301211-0202001303010211-0101212000031101-2330211321000301-1311330112030110) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](data-sources--origin_pool--reference--group-003.md#canonical-3101010321130322-1101022122111133-1122313113112132-3131001013313230-2300230300220313-0123203031320000-3311132120012022-3120030313020001) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](data-sources--origin_pool--reference--group-003.md#canonical-3023211323021121-3302133322231222-0013230011320232-2333121231111320-0323312020212331-3311302300212133-3121003022101030-3222100131032332) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](data-sources--origin_pool--reference--group-003.md#canonical-0131102013021100-3200112220313010-2112013330001312-0100201302000030-2332220112203202-1101331021302131-2333000222130200-0320213311132131) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](data-sources--origin_pool--reference--group-003.md#canonical-2112300001303223-1333103223233110-2022201121123023-3131010003320103-1112123121311320-0312231133021030-3010010331310213-0000122010133112) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-2312213023123311-1023103300000102-3313110223220323-1212011230113131-3201101233232303-3110122023021312-1010000301122132-2011300003022123) |

<a id="canonical-0312332122112201-3302310013110223-2003033012010003-3112321230330112-3302030321210121-1301331321302010-3101000221121311-3030021020132131"></a>

## Next pages — Property reference / 232011232202 / 15

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-3232200323013011-3012121113003102-1101211132122231-3303101001222131-3132321221221112-0012322000210223-2010320011332333-2212232121313131)
- [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-3321301103022100-0002302311013132-3122331220210333-2233013212333133-1323332021333233-0320103330123010-3013120233010203-2201003110133022)
- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-2321331312131123-3012123331102100-2201101123302212-1013031113132232-0023332220033133-3221300001120322-1131200220002203-2121000221033310)
- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-3003002212002303-1310102203201220-1232202311101013-2320113213310102-0313212211221112-0022010223033133-1302122300202210-0233020303132231)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-3130012011302310-0100000020300011-2333213101123221-2323300231131231-3230213133122123-3221221202130031-0122330100113130-1102320312113033)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332302213001121-3131232111300130-0033212230021201-2002221231213032-0131030113000102-1331132222221132-2331031023112033-3203010102102323"></a>

## advanced_options — advanced_options / 120302001110 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- advanced_options

<a id="canonical-2033231303001123-0222123202223310-3321131000020131-3121222301003020-2032103133233130-2203203323002113-2303000301312103-2200301121001123"></a>

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

<a id="canonical-3200331311030021-3021123102000312-2233230321130212-1103130201330033-3220102002021331-1232101310321322-1023311000110102-0210212000311203"></a>

## Direct properties — advanced_options / 120302001110 / 3

- [auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-2200031022002030-1300302112231220-0013002331313110-0112033132001021-1021200020301311-1010033310023321-2012333202220333-3223011312020110): complete subsection reference.

- [circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-3303211032331102-3000131332302232-2111333311103331-0022330012002013-1323300033100223-0311323110301102-1033101131302121-0030030232021011): complete subsection reference.

<a id="canonical-1313301212211312-1203221231330311-3021213133321323-3300030300113211-0310322021121232-0020300322022212-0113220313221330-2313230312020330"></a>

<a id="canonical-3121201100212002-1011132110230023-1123221021110211-0222222221032323-2211202203223203-2121112230122030-1121312300212200-2101303001131213"></a>

## connection_timeout property — advanced_options / 120302001110 / 4

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-2222312123013123-0220032133031323-3221323121132332-0120133323032000-3112221210010233-3331023332110331-0303023302233322-2002211122113112): complete subsection reference.

- [disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-2303331221323232-2131033011320020-1312313203230231-3032101322010010-2132002201310132-2123133113010210-2101203231010123-2102332030103213): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-1213032331232001-0223202331302310-0133021133332212-0023201202211120-1231111323121031-1300030031220231-3332200231131121-0303133203132203): complete subsection reference.

- [disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-0331332310320032-3231220232003000-0202001201230330-1313101231011332-1011333230012321-3230222000110213-2312000213112300-1212020312130023): complete subsection reference.

- [disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-1211030213101131-1002300102001220-1323100123030002-0303312212323331-3023030200230131-1133230332230033-0120213031101232-3011311123010313): complete subsection reference.

- [disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-3201000333232310-1103033203122230-1130310001320233-1131313131333102-3202221000112030-1113122030003330-0330210122231321-1120213033123212): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-2031332023330111-0003133313331220-2300212112321120-0232132200120331-1022021123000030-0202010113022010-3333201013330233-0030200302312021): complete subsection reference.

- [enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321): complete subsection reference.

- [http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032): complete subsection reference.

- [http2_options](data-sources--origin_pool--reference--group-001.md#canonical-0311301212323323-0122213312310323-1212313311033002-1220320200332223-1013033201133131-3022323000020023-1211333021022130-0113320113112110): complete subsection reference.

<a id="canonical-0001321102101331-2130031230130033-0330200313111310-0103033010122232-1130101230122231-3233131231312130-0120021303012221-3100200220120203"></a>

<a id="canonical-0133210311002023-1222003212332013-0223233302123023-2331113002001323-2302111113021212-3201210332113102-1120110031310130-2113313330203220"></a>

## http_idle_timeout property — advanced_options / 120302001110 / 5

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

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

<a id="canonical-3031031223322031-1211103131131121-1023201322100221-2332333233111322-3231323203213022-0131230112330300-2213131112033110-1130220331203213"></a>

<a id="canonical-2020102102311210-3121311102321330-2103213233030231-1110113313113231-2030122330321301-1322332111010203-3001020222030312-2131321123130022"></a>

## max_requests_per_connection property — advanced_options / 120302001110 / 6

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

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

- [no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-0332123032230300-2020011130111022-3322100302022001-3222223331100103-3021233312230333-3233213303320033-0103033300000310-1023002202101333): complete subsection reference.

- [no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-0230231030102102-2331000031100323-1232332011230021-2310131010313322-3211010311030002-2010002313311200-1111003323322203-1220030310231102): complete subsection reference.

- [outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-1020123211100122-1111222303123220-0323332010023123-3013331232213022-2112231021033202-0310323232120133-2211012330103210-3023022112220122): complete subsection reference.

<a id="canonical-1030100211333100-0033323220102303-2223000021212002-2101320111130322-3111003021010300-1311123303222021-1132112312000103-2032103121320233"></a>

<a id="canonical-0110303102020100-1130031211102100-0021133303023230-1233322032220330-2130231133002312-3231223130322031-3123302003121112-3123011302122112"></a>

## panic_threshold property — advanced_options / 120302001110 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-2112302233203232-2312130221113230-0302122201023112-3103002001321120-3131312111202222-2011003132202023-3012300321002230-1313002012030031): complete subsection reference.

- [proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-1310101030312333-0010032203021230-0132232233023132-1230310123232110-0212203312131220-3121132023122030-0103100100313123-0301130111123320): complete subsection reference.

<a id="canonical-3202213302112201-1201333120003022-3233012111210332-0202131332020020-0212211223211230-0123033120310122-3223110321130021-0012023102230223"></a>

## Next pages — advanced_options / 120302001110 / 8

- [advanced_options.auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-2200031022002030-1300302112231220-0013002331313110-0112033132001021-1021200020301311-1010033310023321-2012333202220333-3223011312020110)
- [advanced_options.circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-3303211032331102-3000131332302232-2111333311103331-0022330012002013-1323300033100223-0311323110301102-1033101131302121-0030030232021011)
- [advanced_options.default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-2222312123013123-0220032133031323-3221323121132332-0120133323032000-3112221210010233-3331023332110331-0303023302233322-2002211122113112)
- [advanced_options.disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-2303331221323232-2131033011320020-1312313203230231-3032101322010010-2132002201310132-2123133113010210-2101203231010123-2102332030103213)
- [advanced_options.disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-1213032331232001-0223202331302310-0133021133332212-0023201202211120-1231111323121031-1300030031220231-3332200231131121-0303133203132203)
- [advanced_options.disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-0331332310320032-3231220232003000-0202001201230330-1313101231011332-1011333230012321-3230222000110213-2312000213112300-1212020312130023)
- [advanced_options.disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-1211030213101131-1002300102001220-1323100123030002-0303312212323331-3023030200230131-1133230332230033-0120213031101232-3011311123010313)
- [advanced_options.disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-3201000333232310-1103033203122230-1130310001320233-1131313131333102-3202221000112030-1113122030003330-0330210122231321-1120213033123212)
- [advanced_options.enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-2031332023330111-0003133313331220-2300212112321120-0232132200120331-1022021123000030-0202010113022010-3333201013330233-0030200302312021)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- [advanced_options.http2_options](data-sources--origin_pool--reference--group-001.md#canonical-0311301212323323-0122213312310323-1212313311033002-1220320200332223-1013033201133131-3022323000020023-1211333021022130-0113320113112110)
- [advanced_options.no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-0332123032230300-2020011130111022-3322100302022001-3222223331100103-3021233312230333-3233213303320033-0103033300000310-1023002202101333)
- [advanced_options.no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-0230231030102102-2331000031100323-1232332011230021-2310131010313322-3211010311030002-2010002313311200-1111003323322203-1220030310231102)
- [advanced_options.outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-1020123211100122-1111222303123220-0323332010023123-3013331232213022-2112231021033202-0310323232120133-2211012330103210-3023022112220122)
- [advanced_options.proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-2112302233203232-2312130221113230-0302122201023112-3103002001321120-3131312111202222-2011003132202023-3012300321002230-1313002012030031)
- [advanced_options.proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-1310101030312333-0010032203021230-0132232233023132-1230310123232110-0212203312131220-3121132023122030-0103100100313123-0301130111123320)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2200031022002030-1300302112231220-0013002331313110-0112033132001021-1021200020301311-1010033310023321-2012333202220333-3223011312020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031023330002111-3212133231232013-0201101233032003-3000013021022120-0110120023022230-3210300330331000-3311220000000232-2223233221120030"></a>

## advanced_options.auto_http_config — auto_http_config / 332110123103 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.auto_http_config

<a id="canonical-0023303322122330-0110201011033232-3102231020133113-3122301100133200-0231120303001121-3111203303331233-0201212231123221-2200101331233130"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3332230321003313-0331200331113332-3232330123301330-2322020322133203-3010222033212132-0111123021310333-1112132013221310-0310000120121201"></a>

## Direct properties — auto_http_config / 332110123103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230110022233323-0101232222110011-1323210233202010-0213031330000000-2030331130232132-3001033231002301-1223313133000330-3200221200210120"></a>

## Next pages — auto_http_config / 332110123103 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3303211032331102-3000131332302232-2111333311103331-0022330012002013-1323300033100223-0311323110301102-1033101131302121-0030030232021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013202203112000-2133303202013101-0000223012202312-1233032221020110-2330302211203110-0203113132113213-3222211222102022-3323321110210320"></a>

## advanced_options.circuit_breaker — circuit_breaker / 330321112203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.circuit_breaker

<a id="canonical-0201310102212200-3201020120222301-0323301123232203-3103332113133313-3023110123312213-3101100121230101-3311000111133012-0122033013133331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0031101102322211-0033100320213312-0132123322311000-3300112320133030-3132322333123001-3130121102000022-0202303323220210-0200200120202300"></a>

## Direct properties — circuit_breaker / 330321112203 / 3

<a id="canonical-2033013232330200-3033032001001310-1202012030230122-0101333330121102-1123202121020011-3130322202101010-2001110323023001-1100303000300100"></a>

<a id="canonical-0033330320231210-1030313001103313-2201303100110131-3322311100310003-3023112312200200-3312310212112323-1132333103330103-2101303300221230"></a>

## connection_limit property — circuit_breaker / 330321112203 / 4

Type: `"number"`. Computed.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

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

<a id="canonical-1122322311330123-1211201200303032-2110321131103310-3203023312033121-0023230321012013-2102031013001131-3312130233212200-2301133013310233"></a>

<a id="canonical-1333303011133300-1030133220010330-2300110321313210-1020023132231223-0202121002221002-2101131100221123-2203230000022001-0001230031320021"></a>

## max_requests property — circuit_breaker / 330321112203 / 5

Type: `"number"`. Computed.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

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

<a id="canonical-0302312112110122-3111301320001222-1203011133012212-0012301320303331-2133320222203211-1131210210300011-3023112321302201-2120222311232320"></a>

<a id="canonical-2312313102331311-1022313131202200-2213230010101310-1303132331110011-2332022210213212-1113200013223012-1003310222020231-0322013011103113"></a>

## pending_requests property — circuit_breaker / 330321112203 / 6

Type: `"number"`. Computed.

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

<a id="canonical-3332120101123321-1112320230310303-2023303103033332-2232010202202131-0000010201310203-2000112010110313-3311233000212200-1022302022211110"></a>

<a id="canonical-0310020212211013-0232031331202223-2332210012103222-1331301131211303-0232131020011213-0013031310110320-3011121301203330-1221220132230301"></a>

## priority property — circuit_breaker / 330321112203 / 7

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

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

<a id="canonical-3322131120003331-3300031133303232-2233322213132202-1100203220122323-2010202203000301-3332023103231320-3312121002131322-0120121310220312"></a>

<a id="canonical-3302032232122120-2211312303101123-0000223333301003-2231333011133302-3033222113013121-1120302032110211-1023020203210232-2211323003113231"></a>

## retries property — circuit_breaker / 330321112203 / 8

Type: `"number"`. Computed.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

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

<a id="canonical-2211110311323220-1102112320222000-2202111210201230-0203301013333201-3113331211112201-2301301301021310-2330003022323012-1030013230231113"></a>

## Next pages — circuit_breaker / 330321112203 / 9

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2222312123013123-0220032133031323-3221323121132332-0120133323032000-3112221210010233-3331023332110331-0303023302233322-2002211122113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123201013212323-1030212020310033-3021112233333202-0133213021121333-3110231233231013-3300103122020312-3133330012011011-1030000330300123"></a>

## advanced_options.default_circuit_breaker — default_circuit_breaker / 103012300121 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.default_circuit_breaker

<a id="canonical-3003210100122113-1110202321020020-0020131122311320-1031101122321322-3322303111123000-3012210231132200-3033322102233203-3332012001213103"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2232120102001033-1303000223120100-3310020132032330-0210333202003331-0110012030002223-3110100210233331-1233111300231032-0310131233311223"></a>

## Direct properties — default_circuit_breaker / 103012300121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010023013220223-3303203310331330-0020323112021200-2113112123223301-1322112313013020-1031210211003100-3133131222012223-0213010122010202"></a>

## Next pages — default_circuit_breaker / 103012300121 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2303331221323232-2131033011320020-1312313203230231-3032101322010010-2132002201310132-2123133113010210-2101203231010123-2102332030103213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112110200120230-2021030211302113-0003202012301200-3100332002001033-2323300300320212-0012102011310010-0233121233102313-3201011103130022"></a>

## advanced_options.disable_circuit_breaker — disable_circuit_breaker / 200131111231 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.disable_circuit_breaker

<a id="canonical-2103331103012022-0121331233320100-1020203320112232-3210223322111113-1213031101110022-0101012332203022-1213333311123221-0122210332123102"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0101000023032103-2223330012110221-0211211310230022-2210011102301031-2112222223231100-3122231133110212-0022000103003201-1021131032301033"></a>

## Direct properties — disable_circuit_breaker / 200131111231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200011230221312-1022030032301020-1123321121202103-1310211331200332-3323303232030131-3113011321002123-3103013230011202-1011102320013323"></a>

## Next pages — disable_circuit_breaker / 200131111231 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1213032331232001-0223202331302310-0133021133332212-0023201202211120-1231111323121031-1300030031220231-3332200231131121-0303133203132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233022110133023-2001032201230330-0021023002322212-2122331121110103-0112120001033121-0332000121202232-3320121120022303-0311031232310310"></a>

## advanced_options.disable_lb_source_ip_persistence — disable_lb_source_ip_persistence / 120030123020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.disable_lb_source_ip_persistence

<a id="canonical-0223101220202022-3201110002131132-0122010220112231-3301120301221001-0212113002220300-3033312012320010-3123220031212012-2022002321201222"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3200011102120001-1032003301200333-2221023120001002-1230333212200301-2200121110122021-3311303233211333-2132012313222220-0101121020113012"></a>

## Direct properties — disable_lb_source_ip_persistence / 120030123020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100331303023312-0321113103322320-3032212000202220-3232232122123321-0123023100210011-0022201223012223-2200121210200102-0210330100002013"></a>

## Next pages — disable_lb_source_ip_persistence / 120030123020 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0331332310320032-3231220232003000-0202001201230330-1313101231011332-1011333230012321-3230222000110213-2312000213112300-1212020312130023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030122122013003-2121213320002001-3131020231032320-2120320010021131-1301303000201131-0033330201202122-1121223320320223-1203001011333323"></a>

## advanced_options.disable_outlier_detection — disable_outlier_detection / 323330112300 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.disable_outlier_detection

<a id="canonical-0132022131122020-2010013331030323-2310300321310021-3031032022133033-2331022203030000-2032221003120202-1300221212301210-2130313122203131"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0312330011003111-2113011031200332-3023012022121321-2232232021103211-0213120221213331-3321111232130223-3100130001310223-1213313131120000"></a>

## Direct properties — disable_outlier_detection / 323330112300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020111201203010-3322203101012323-2111303033000230-1032113203301120-1001030303201112-1011120030313200-1001213321011203-0321110021013031"></a>

## Next pages — disable_outlier_detection / 323330112300 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1211030213101131-1002300102001220-1323100123030002-0303312212323331-3023030200230131-1133230332230033-0120213031101232-3011311123010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222121130313300-0220303303123110-1213030312100233-3003003313313021-3333210221220101-0211110110211113-0203301121302233-1011210223200011"></a>

## advanced_options.disable_proxy_protocol — disable_proxy_protocol / 123220320322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.disable_proxy_protocol

<a id="canonical-0311111311310132-2011221212102202-3222033123112001-0302100021321331-1102102210211332-3311002212032032-0110103332330213-3130103121032333"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3321111320132220-2012113030312033-2000230020221221-2012231130301030-3103230312131332-2100113202203320-2133220323032232-3023300210210320"></a>

## Direct properties — disable_proxy_protocol / 123220320322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210010230201102-1021223330022313-3130111010003312-3020023312131300-1200003311121132-1221323213303313-1323012013123303-3121023301212320"></a>

## Next pages — disable_proxy_protocol / 123220320322 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3201000333232310-1103033203122230-1130310001320233-1131313131333102-3202221000112030-1113122030003330-0330210122231321-1120213033123212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133333221213200-3330113303211023-1333103231113232-0331312112320021-0223023202021103-0231301231123131-0223023200300301-3310022103020011"></a>

## advanced_options.disable_subsets — disable_subsets / 020102302033 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.disable_subsets

<a id="canonical-1113202310302103-3013020212022223-1020030320331323-3002322130230202-3023020131022313-1113033303332233-3133230133010330-1121130210301200"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0320312003302021-3213301231210212-3011003011023033-3211322222012032-2302233303302022-2302213321233123-3332020023212103-1020110203032332"></a>

## Direct properties — disable_subsets / 020102302033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033323113232122-3123131322222133-2222110020002312-0231132233032302-3002123011201322-2223230111212112-1020110010102111-2311302220032003"></a>

## Next pages — disable_subsets / 020102302033 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2031332023330111-0003133313331220-2300212112321120-0232132200120331-1022021123000030-0202010113022010-3333201013330233-0030200302312021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002221123123122-0012001222112231-1102312012122133-1011231231222332-1130010312203332-2110230322021220-1101001210302213-1303000000031011"></a>

## advanced_options.enable_lb_source_ip_persistence — enable_lb_source_ip_persistence / 003201313020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.enable_lb_source_ip_persistence

<a id="canonical-1013230220212023-1033133123012201-3121301321333211-3200232311033122-0022123203332110-1012022202223230-3033302331301232-2012300313112002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0301023113112010-2030032311230101-3313323102220321-2032020333302132-2223231302313111-2032222002113232-1232120311221131-0003123302112111"></a>

## Direct properties — enable_lb_source_ip_persistence / 003201313020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220302212230103-3020311203133301-2203122230233102-0231023310132022-2011200022132002-1000211212120003-3112202033113230-0302103122021103"></a>

## Next pages — enable_lb_source_ip_persistence / 003201313020 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112203200203120-3330023320303231-3021111310131103-0300213012031103-3202223102220232-2212002010030330-1311311002021123-3021020302312022"></a>

## advanced_options.enable_subsets — enable_subsets / 021311201310 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.enable_subsets

<a id="canonical-3300120222132113-1333330221301102-0131331033332022-0221221121232101-3230123333212213-2030101030302132-0302013221122033-1100302213300000"></a>

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

<a id="canonical-3002103111331220-3201002320202203-3222232031002312-1022212012001313-1331120002303222-3111100123033013-2131033201111021-1211220031132330"></a>

## Direct properties — enable_subsets / 021311201310 / 3

- [any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-2002322200012102-3033002212332233-2222220010121203-2031100311333130-3011032131002030-1201021123003103-2002203333110220-0131111221121311): complete subsection reference.

- [default_subset](data-sources--origin_pool--reference--group-001.md#canonical-0211221101102112-2232313130313102-1130122133222302-3320123202201000-0022132021130312-2233330130310012-0112011330213202-2313211102100132): complete subsection reference.

- [endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1213232022232131-2112013323330233-1311213330120132-3303012110020032-3112120311311000-0110000133320333-2211123131232003-1333103222212210): complete subsection reference.

- [fail_request](data-sources--origin_pool--reference--group-001.md#canonical-2332210123231121-1331322021121303-1200310010003331-1312202011002330-3302121323302123-3231031323010332-3032303221322211-3132023332011010): complete subsection reference.

<a id="canonical-1000120320132020-2033120013121202-3110302330133332-1311302321231102-0133132320333112-3130211310202000-2212003023332103-3100322103213032"></a>

## Next pages — enable_subsets / 021311201310 / 4

- [advanced_options.enable_subsets.any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-2002322200012102-3033002212332233-2222220010121203-2031100311333130-3011032131002030-1201021123003103-2002203333110220-0131111221121311)
- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-0211221101102112-2232313130313102-1130122133222302-3320123202201000-0022132021130312-2233330130310012-0112011330213202-2313211102100132)
- [advanced_options.enable_subsets.endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1213232022232131-2112013323330233-1311213330120132-3303012110020032-3112120311311000-0110000133320333-2211123131232003-1333103222212210)
- [advanced_options.enable_subsets.fail_request](data-sources--origin_pool--reference--group-001.md#canonical-2332210123231121-1331322021121303-1200310010003331-1312202011002330-3302121323302123-3231031323010332-3032303221322211-3132023332011010)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2002322200012102-3033002212332233-2222220010121203-2031100311333130-3011032131002030-1201021123003103-2002203333110220-0131111221121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302031313112112-3103132201332032-2012113131310131-3221303302031000-3312233132230320-3213121113111212-0222013300121330-1030213213123302"></a>

## advanced_options.enable_subsets.any_endpoint — any_endpoint / 313313332203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- advanced_options.enable_subsets.any_endpoint

<a id="canonical-3101222313123233-2330322022200332-3121011001200233-3121102123013230-2222311033021030-0330011322303113-1131020133130102-1211301210101303"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1332003203131222-0231010022302100-0020112232233012-0122022000102021-3220132320212012-3300213013203131-0222023231211113-3312110302103311"></a>

## Direct properties — any_endpoint / 313313332203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232321000313011-1311223320302303-1321020011110121-3202322131331232-2302323322113121-1332232121320203-2200223031132321-1320001311210302"></a>

## Next pages — any_endpoint / 313313332203 / 4

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0211221101102112-2232313130313102-1130122133222302-3320123202201000-0022132021130312-2233330130310012-0112011330213202-2313211102100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032201002132230-2220331313232133-2313220131232301-3031112331233011-0023331202102200-0113310033001300-1003031110321120-0112331313000021"></a>

## advanced_options.enable_subsets.default_subset — default_subset / 210321022113 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- advanced_options.enable_subsets.default_subset

<a id="canonical-1330213222300200-2130012201301121-1320230232022300-3132112103222300-2032121013021020-1223021000322323-1003300110023223-2213120213222220"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0110013130200220-3222233122223132-3322123222001011-2112312310330231-0213112201330022-0002100303311211-2320000212321201-1310310330223131"></a>

## Direct properties — default_subset / 210321022113 / 3

- [default_subset](data-sources--origin_pool--reference--group-001.md#canonical-1001211333300301-2012222022200220-3012303011312123-2300231033101320-2331202022312222-2032330233011210-2033311222220220-0032121301211302): complete subsection reference.

<a id="canonical-1012312100103233-1232001122113121-3001112312121133-1320300210230003-1202230103000321-1303000232022330-3001010002332312-2102222010203103"></a>

## Next pages — default_subset / 210321022113 / 4

- [advanced_options.enable_subsets.default_subset.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-1001211333300301-2012222022200220-3012303011312123-2300231033101320-2331202022312222-2032330233011210-2033311222220220-0032121301211302)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1001211333300301-2012222022200220-3012303011312123-2300231033101320-2331202022312222-2032330233011210-2033311222220220-0032121301211302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002011230303212-3310001201133212-1223331230231220-0220010300221131-3203322013033220-1223301020203120-1311021303001213-2102012230101021"></a>

## advanced_options.enable_subsets.default_subset.default_subset — default_subset / 103220212012 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-0211221101102112-2232313130313102-1130122133222302-3320123202201000-0022132021130312-2233330130310012-0112011330213202-2313211102100132)
- advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-0310233311020332-2331022310012210-3113102110100020-2331122310132122-0222300222201222-2312331110232113-3012210233111230-3310132203201102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3012120311221113-0303012002300110-3101020223031030-3101230313102011-3211102222221131-0021103120111010-2303102030200332-3302110111101333"></a>

## Direct properties — default_subset / 103220212012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302322312022111-2103100013131112-2111213113123302-1221210131021202-3112132202022230-0011211102233022-3012221122010003-0122022103220031"></a>

## Next pages — default_subset / 103220212012 / 4

- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-0211221101102112-2232313130313102-1130122133222302-3320123202201000-0022132021130312-2233330130310012-0112011330213202-2313211102100132)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1213232022232131-2112013323330233-1311213330120132-3303012110020032-3112120311311000-0110000133320333-2211123131232003-1333103222212210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120313033122001-0013313300331203-2013230010011221-3322300013222030-1011111200002122-3303131101231302-2001311100112113-3001200120000302"></a>

## advanced_options.enable_subsets.endpoint_subsets — endpoint_subsets / 313321330230 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-3001123213133310-2001013202331113-3312010310011100-1031103222311313-1131002321233123-3011112112211130-0302312211202013-3001330321030331"></a>

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

<a id="canonical-0000112222112323-0032022221322303-2203111012322222-3221102011022202-1333033331302012-3010011030003023-2021220222301303-1312020130011030"></a>

## Direct properties — endpoint_subsets / 313321330230 / 3

<a id="canonical-2322123120302300-0211003003203231-2312033221013001-2022220311213222-1010202322310003-0230131002011212-2000023303020033-2122120213133111"></a>

<a id="canonical-1300033013213001-3021103300112311-1300301000313332-0301112310122301-1100323221013030-0021233201121221-3113132320121302-2103121230301310"></a>

## keys property — endpoint_subsets / 313321330230 / 4

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

<a id="canonical-3100132200010120-3121100321302113-1331000000201121-0222210102100303-3103310203201202-1033111013323133-2030320010032002-2310210213132301"></a>

## Next pages — endpoint_subsets / 313321330230 / 5

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2332210123231121-1331322021121303-1200310010003331-1312202011002330-3302121323302123-3231031323010332-3032303221322211-3132023332011010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203302132200310-1213003112001312-3013330133100131-2031122212313323-0323322111232002-1003002000313002-3221313021102121-2212100032110122"></a>

## advanced_options.enable_subsets.fail_request — fail_request / 220332312003 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- advanced_options.enable_subsets.fail_request

<a id="canonical-1222310232302021-3020123031233333-3032032201331123-2130323000222032-0111132203132113-3200212033030221-0101212003011331-2330311102310023"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3333012323222031-0131320120131110-2130232121022000-1032001111133322-2233100220310102-2333232033203103-0112302321033133-1130101222322311"></a>

## Direct properties — fail_request / 220332312003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020203322200231-1223000300032300-3100303111103332-1020021301012212-1232322310133310-3122231112100301-3123032130221103-3110103120003221"></a>

## Next pages — fail_request / 220332312003 / 4

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201311011000110-1323303011122013-2201311222120210-1033121313212130-3102031113330030-2002200321322001-1012200100103323-2330101302013333"></a>

## advanced_options.http1_config — http1_config / 020202100331 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.http1_config

<a id="canonical-3011311102322130-1323231223210330-0122202001203201-2003130102112122-0212332013131102-1202203000221220-3023222020201101-2200333303210233"></a>

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

<a id="canonical-2132210232213213-2013201200313033-2301033213020211-3331130303032222-0011300131101023-2133311120332232-3320110331021330-1322322313103223"></a>

## Direct properties — http1_config / 020202100331 / 3

- [header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132): complete subsection reference.

<a id="canonical-2211303030103210-0300112301112321-2110101100123312-1110003210021232-0021100031232210-0031220132112201-2020100120010303-0022112213202121"></a>

## Next pages — http1_config / 020202100331 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321120233221311-1102023323031023-1313223233222112-1312233232222230-2231023200212233-0030123010100300-0302212130233200-0332103001223313"></a>

## advanced_options.http1_config.header_transformation — header_transformation / 213033223001 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- advanced_options.http1_config.header_transformation

<a id="canonical-0223111232112221-2300021320012132-2323122301302123-2131101233312010-1302312101002332-1332031301023002-3113013330131321-1222333000103013"></a>

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

<a id="canonical-3310203011100103-1121323233231013-2310313330322120-3312220030113002-1023221302122222-1032301203220100-1330231330010320-0122220030111113"></a>

## Direct properties — header_transformation / 213033223001 / 3

- [default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-2230302220133213-1221113122202131-2102011112330112-3013310132212103-1002212222220002-1033011133310211-2300201023323212-0122023111021133): complete subsection reference.

- [preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0203032222222022-0133302011332232-2110102201203003-0310310133210310-2302201110310212-2013323233311002-2033202332120331-1210323012121102): complete subsection reference.

- [proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0333133130310311-1230002311120020-1320211103011102-3200032321003322-3310022300313210-3033122333021032-1100021000322113-1232302203122300): complete subsection reference.

<a id="canonical-0013333132032111-2111122330210223-0111123023112010-3023303210100210-2103033202310010-0001011010222223-0232030100312201-0210210010210031"></a>

## Next pages — header_transformation / 213033223001 / 4

- [advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-2230302220133213-1221113122202131-2102011112330112-3013310132212103-1002212222220002-1033011133310211-2300201023323212-0122023111021133)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0203032222222022-0133302011332232-2110102201203003-0310310133210310-2302201110310212-2013323233311002-2033202332120331-1210323012121102)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-0333133130310311-1230002311120020-1320211103011102-3200032321003322-3310022300313210-3033122333021032-1100021000322113-1232302203122300)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2230302220133213-1221113122202131-2102011112330112-3013310132212103-1002212222220002-1033011133310211-2300201023323212-0122023111021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320230332212333-3232100001202233-3230030210200020-0102302031310132-2202330113332133-0210230121110103-2102102320120202-3330313322222231"></a>

## advanced_options.http1_config.header_transformation.default_header_transformation — default_header_transformation / 023000120222 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-0313213332222132-1003002112123213-1203100012030113-0100031311000123-3133211201023112-1312133322300321-2311111030100023-2102103100202002"></a>

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

<a id="canonical-3003003302103213-3300011102300310-1313320121110210-1032213002220002-1102201010331103-1112313020212032-0100032323130130-0302131021320221"></a>

## Direct properties — default_header_transformation / 023000120222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012313312031003-3132013223332001-1122322023312031-1211100010031020-1222322331001111-3203011331320021-1220323201302321-2022303120222210"></a>

## Next pages — default_header_transformation / 023000120222 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0203032222222022-0133302011332232-2110102201203003-0310310133210310-2302201110310212-2013323233311002-2033202332120331-1210323012121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001020000122200-2022103002120320-3212223021011003-2211001301222101-1013312230003232-3313113120320323-3033310302000311-2303310001020023"></a>

## advanced_options.http1_config.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 033310333203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-2220101031321301-1310012031011300-3012132203210131-0011000002010322-0302023213313231-0323100113332312-1300312010133320-3010001100323133"></a>

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

<a id="canonical-2311212130122021-1303132210012222-3120201011023312-2021000032001111-2000331232303323-1223031020012300-2220030031320021-0333302031011323"></a>

## Direct properties — preserve_case_header_transformation / 033310333203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001221110033110-0323013020320303-2200033130123100-2302230130322133-2031230232330021-3300212221020300-3313132111213221-0323001113030020"></a>

## Next pages — preserve_case_header_transformation / 033310333203 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0333133130310311-1230002311120020-1320211103011102-3200032321003322-3310022300313210-3033122333021032-1100021000322113-1232302203122300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102023303310202-1331032131021320-2302213003221230-1031223332223232-0033110230121203-1111312300120122-0212132301132120-0331123011213331"></a>

## advanced_options.http1_config.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 311101310300 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-3000232000131123-2231230033310122-0013001113112331-0211303013030230-3133232030111120-1122311133222312-1022213232232321-3003311100233313"></a>

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

<a id="canonical-2131001232032213-0332222323030033-2203101130310131-0102212031001302-0221313201113032-2331013332332032-2222101103021330-1301221003030022"></a>

## Direct properties — proper_case_header_transformation / 311101310300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233220211200222-1330220031033201-2110002333211320-3023210030333213-1011331030331010-3230031313333322-0300221112221300-1332322002321010"></a>

## Next pages — proper_case_header_transformation / 311101310300 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-1320030132230032-2220023120120100-2303100020220101-3203000132101012-2203012331113301-3023220202131130-2302111311100130-3310121000100132)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0311301212323323-0122213312310323-1212313311033002-1220320200332223-1013033201133131-3022323000020023-1211333021022130-0113320113112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020333010320201-2230121311021013-0223010002221213-2002103313311121-3102330332013111-1110103032232011-1200203013323133-1111203313131211"></a>

## advanced_options.http2_options — http2_options / 301010301031 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.http2_options

<a id="canonical-0030122322133321-1301003300103003-3301131221031302-3010312203022202-0300230022233012-3131033313110200-2112333103220030-2031023320121022"></a>

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

<a id="canonical-3213120311112211-2202010010313222-1303333211220302-0302223003113123-3102212313201312-3033123201013230-3231100232212200-2002312122212223"></a>

## Direct properties — http2_options / 301010301031 / 3

<a id="canonical-3302103020113132-2330302202103300-1130023211332221-1313133110113012-2130030302002100-3112230130322001-0122331310312323-0313133212132120"></a>

<a id="canonical-1321121032011012-3313212303112122-3313013121130332-1131330023120131-3333213120201120-3100302320000101-0323022230123332-1311133321122130"></a>

## enabled property — http2_options / 301010301031 / 4

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

<a id="canonical-1211232310101021-1012022122010033-3131120100100001-2033101231022211-0003130003303001-3321222220122102-1110300311022233-3131011133223211"></a>

## Next pages — http2_options / 301010301031 / 5

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0332123032230300-2020011130111022-3322100302022001-3222223331100103-3021233312230333-3233213303320033-0103033300000310-1023002202101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322332033200131-0121100221021202-3333330131313232-3321101322120111-3300221001113220-0000230101113230-0002102231300232-1213000222020113"></a>

## advanced_options.no_panic_threshold — no_panic_threshold / 310130231201 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.no_panic_threshold

<a id="canonical-1313132231132013-0133211333033201-2033001300013312-0223210103313312-1000310003033100-1331212301202332-1202221113102231-3101031301202320"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1003302101010120-3100023000113023-1001132013302223-2211201231220102-2110122201210020-3201233230232033-1010021011323030-0233113223332200"></a>

## Direct properties — no_panic_threshold / 310130231201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200300333230031-0222013110301221-1133132132011130-0022022030221021-2101323310233221-0000323012111123-3113131113110233-0301232223300011"></a>

## Next pages — no_panic_threshold / 310130231201 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0230231030102102-2331000031100323-1232332011230021-2310131010313322-3211010311030002-2010002313311200-1111003323322203-1220030310231102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123213223022032-2301132302312112-2233232130332301-1312031003311122-0323020031112011-1321102333322233-1332223121220111-1012130200230302"></a>

## advanced_options.no_request_limit_per_connection — no_request_limit_per_connection / 101023200301 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.no_request_limit_per_connection

<a id="canonical-1330331121013031-3322030210320321-3103131130331333-3213212002231222-1232120213031333-3100313112320100-1100300012212112-3102333032210223"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0313232031011332-0121220110021330-0111001123310023-0102310002331001-0003102222121012-1233200110002013-1300303330210022-3021211031121003"></a>

## Direct properties — no_request_limit_per_connection / 101023200301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231231101232122-2121233001031312-2302013301320031-1011331321012101-2132013023202221-1232312021212322-1302300122031331-0132200023230123"></a>

## Next pages — no_request_limit_per_connection / 101023200301 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1020123211100122-1111222303123220-0323332010023123-3013331232213022-2112231021033202-0310323232120133-2211012330103210-3023022112220122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133302300132320-0120012002311222-3020321020021302-3032011321203221-1300022022023101-1231122023100022-2133300110213001-3132001210300232"></a>

## advanced_options.outlier_detection — outlier_detection / 110003311031 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.outlier_detection

<a id="canonical-0120230001000133-3203310223211112-3132133011001011-2002232100121213-2202003300213112-1023012233021110-0302231011331111-1213103310302231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0031111011012121-1110113233112212-0313021000211302-0120132020210101-1220322101311023-1210111213322320-2332120130101203-2120333130001112"></a>

## Direct properties — outlier_detection / 110003311031 / 3

<a id="canonical-1032230222031213-0223102201311333-2212313202302122-3320000333200121-2123011222133123-3312103113232312-0212023201333121-2223022103132121"></a>

<a id="canonical-0030012311113231-2132331032100200-0213202301120101-2322211313312320-1120200330113211-0132032212011312-2021211302302011-3311303303130131"></a>

## base_ejection_time property — outlier_detection / 110003311031 / 4

Type: `"number"`. Computed.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail.

Upstream description:

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

<a id="canonical-3310110022021010-2211033323122201-1232021122003332-3100101022120222-1102330102123003-2230032122000211-1322200131210013-1330001100012301"></a>

<a id="canonical-1003333212002300-3030000011313003-1311110323002131-2130000202303213-3330032302323203-3310021231100302-3320210001222200-2323122202131032"></a>

## consecutive_5xx property — outlier_detection / 110003311031 / 5

Type: `"number"`. Computed.

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

<a id="canonical-2021221200003200-2123031132201220-3102221102300311-1332110030330122-3220002022233112-2311302312221033-1222231120222020-0113012320312133"></a>

<a id="canonical-3123013221320002-0010100231203031-3212002323121003-0303101230331013-2032231112131311-0313300010231310-2011101231013130-3132121112030320"></a>

## consecutive_gateway_failure property — outlier_detection / 110003311031 / 6

Type: `"number"`. Computed.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.)..

Upstream description:

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

<a id="canonical-0113320002320321-3020322133233101-2201201130231111-2230033232130112-2002122123201212-2000201211100322-0223320302102311-0203312201333100"></a>

<a id="canonical-0302000311323313-0321230230332023-2120322010123003-2001303012011022-0131000001110111-3213223320220201-1310330301131032-3213112002121111"></a>

## interval property — outlier_detection / 110003311031 / 7

Type: `"number"`. Computed.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Upstream description:

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to 10000ms or 10s. Specified in milliseconds.

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

<a id="canonical-3211300220032332-2220121022233030-2302332031013213-2322211311310211-2210211111312330-2032230032123301-1102201323103123-1300220220332201"></a>

<a id="canonical-0121221031122222-3112320323303111-1300221013331201-3103223103223320-1131310202310132-3330103001003201-3322332033310212-0123311322123113"></a>

## max_ejection_percent property — outlier_detection / 110003311031 / 8

Type: `"number"`. Computed.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Upstream description:

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

<a id="canonical-3013233322030032-1123123313221331-1102102230313223-3203120013210212-1011101213213123-2020223212032031-2023000211110032-3121330010330130"></a>

## Next pages — outlier_detection / 110003311031 / 9

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2112302233203232-2312130221113230-0302122201023112-3103002001321120-3131312111202222-2011003132202023-3012300321002230-1313002012030031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300333201031223-1110231300211122-2132100112300033-0133033132202221-0222211032111030-0133203300233232-2302010220121303-1122220321030000"></a>

## advanced_options.proxy_protocol_v1 — proxy_protocol_v1 / 003300222310 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.proxy_protocol_v1

<a id="canonical-3131021033021122-3332300331321230-2213310130323231-0332020030322311-3310112321122301-3023232230203212-3210233003132203-2111213220202321"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2303230032212031-0233233302101001-1111321210301030-0030131013000232-1310003023232023-3101131212330220-1212010233231130-3313332011103003"></a>

## Direct properties — proxy_protocol_v1 / 003300222310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100012113120331-1122213122322121-0303011301031000-3333012233010032-2300301223320121-1120220120322333-3213230123211223-3013020333020120"></a>

## Next pages — proxy_protocol_v1 / 003300222310 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1310101030312333-0010032203021230-0132232233023132-1230310123232110-0212203312131220-3121132023122030-0103100100313123-0301130111123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112311333201002-0123130320003132-1111103031120211-3101310301022003-0331112212313320-3312232310223231-1221030202112302-0113313322201010"></a>

## advanced_options.proxy_protocol_v2 — proxy_protocol_v2 / 021012312311 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.proxy_protocol_v2

<a id="canonical-3303032000000002-1211212303130132-3133011323011230-1210302310232032-2203102030322323-0211220011222302-1322211323103132-2010322102210113"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0123213101301300-1021212203220222-2210112313223311-1121132332230102-1000312133313212-0012023032030123-3311220230331131-2213123222202310"></a>

## Direct properties — proxy_protocol_v2 / 021012312311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222132101111001-2302101311222322-2003233230002321-3110230210231310-0300323323101201-0110010301021223-1200021233231110-1110102023313330"></a>

## Next pages — proxy_protocol_v2 / 021012312311 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3232200323013011-3012121113003102-1101211132122231-3303101001222131-3132321221221112-0012322000210223-2010320011332333-2212232121313131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022212322233310-3002221102033100-1202222300131310-2302003022230132-0133211213320222-1202213333033231-0002300303100112-2302120311113222"></a>

## automatic_port — automatic_port / 333222131312 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- automatic_port

<a id="canonical-3330030101213132-1023001113001113-2323003003322122-3003122331131110-1132010030223132-1131332022230111-3132311313331113-2130313202000312"></a>

Type: `["object", {}]`. Computed.

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

- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-3330030101213132-1023001113001113-2323003003322122-3003122331131110-1132010030223132-1131332022230111-3132311313331113-2130313202000312)
- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-3332320330310202-3001331103323210-0100103311013120-0002100310001103-2000333310000320-3010022210312130-1113032332102000-2210230130200022)
- [port](data-sources--origin_pool--reference--group-001.md#canonical-1331233022212123-0232020012320101-0313013123210003-2023013132030210-1213332332113330-0220222032331222-1133021111301032-2011132211203233)

Select alternatives according to the provider validators above.

<a id="canonical-2321312033033032-0331222202230232-0332011202330303-3130030002312302-3031132302203322-2303302002212313-0031003010312320-3123220030330113"></a>

## Direct properties — automatic_port / 333222131312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103212113232001-3131233020132130-3332312303333232-0211321322131301-0000102332301132-3301303230211323-1300003300100231-1011000210303133"></a>

## Next pages — automatic_port / 333222131312 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3321301103022100-0002302311013132-3122331220210333-2233013212333133-1323332021333233-0320103330123010-3013120233010203-2201003110133022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111232223103301-0203322220312303-0003132301003133-1330301210203022-3010002311303311-0202202310011101-0122322222311203-2231131010120132"></a>

## healthcheck — healthcheck / 003130223111 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- healthcheck

<a id="canonical-1013131122022033-3120111310102210-0122223231303100-1002203002313300-3112312220330112-0123221302122033-2111003202132002-2200220232301001"></a>

Type: `"list"`. Computed.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Reference to healthcheck configuration objects.

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

<a id="canonical-3311003322132300-1321110032320310-2230220323223210-3101221101103121-0323100002313002-1023323231202102-3013202111010130-0330221111201003"></a>

## Direct properties — healthcheck / 003130223111 / 3

<a id="canonical-1020202033220110-3030301031212233-1331211302231211-0022012012313121-1313000312202122-3213323223122230-0013220011313121-0320012130000011"></a>

<a id="canonical-0031301200230233-1311102332013221-0111232230033111-1210012030120010-0311110122211203-1212320213122020-1311003021210021-2222022110311213"></a>

## name property — healthcheck / 003130223111 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2132130030010021-2022302112231131-1233200313111231-3221201011001111-0220233100220123-2033332301330200-2313220132112001-0122122033122303"></a>

<a id="canonical-0221203120200321-3233213130302103-2311122311032111-1101203131222312-0131020230313022-3133203031022213-1010111131311221-3312211101322132"></a>

## namespace property — healthcheck / 003130223111 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3333122232012021-0202223110210203-1313302332213200-1122211113331230-3221012013200330-0102023300002010-3111010311122323-1033012331310313"></a>

<a id="canonical-0233313000013301-2200233221201123-2011223121113021-0113312020133300-2012200201313221-1010211000002311-3101132201121121-0111332310122203"></a>

## tenant property — healthcheck / 003130223111 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0232321133202312-3330312122103332-0201023131310030-1121321320303223-2030030212310332-2120021231203102-1201330312003301-1003232103330311"></a>

## Next pages — healthcheck / 003130223111 / 7

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2321331312131123-3012123331102100-2201101123302212-1013031113132232-0023332220033133-3221300001120322-1131200220002203-2121000221033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212122300331331-2033010303111100-0212111131012200-0233313322302311-3211112230111003-1322021011012332-2210223023332130-1131020333212023"></a>

## lb_port — lb_port / 111210201020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- lb_port

<a id="canonical-3332320330310202-3001331103323210-0100103311013120-0002100310001103-2000333310000320-3010022210312130-1113032332102000-2210230130200022"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0203313130311123-3312010233332233-1032313012001000-0211302011030322-1020321223300101-3301133030331220-0310102233310013-2020020201222133"></a>

## Direct properties — lb_port / 111210201020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113320213113221-1001321132323132-1331021001301122-2323013331020212-3012213002230201-1131123120232331-0301033120233011-3212121010332012"></a>

## Next pages — lb_port / 111210201020 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3003002212002303-1310102203201220-1232202311101013-2320113213310102-0313212211221112-0022010223033133-1302122300202210-0233020303132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311123323231002-0020101111211230-3200101303100013-0202133200130231-2212133031130123-0122212111303130-2313202011010233-2300312133231201"></a>

## no_tls — no_tls / 112312013123 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- no_tls

<a id="canonical-3003013113222220-2211123101222221-3301200210020030-1013302300212120-2102113320000233-3323230131232012-1312300231010103-1233101032222222"></a>

Type: `["object", {}]`. Computed.

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

- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-3003013113222220-2211123101222221-3301200210020030-1013302300212120-2102113320000233-3323230131232012-1312300231010103-1233101032222222)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-0002222103101132-0331331011030000-0221212201230213-1311030121013132-0223020023313332-2311112031102200-3132032023331103-2122132001320021)

Select alternatives according to the provider validators above.

<a id="canonical-1030213032121302-2033121211021022-3330003000300113-2232100020302203-2233113322213000-0102303021102322-0113010321032233-3133311200311010"></a>

## Direct properties — no_tls / 112312013123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303321133033333-1113121211220100-3130212001031003-1321312231211230-0202230011310312-1101321013003021-2303010010212231-1302312203121012"></a>

## Next pages — no_tls / 112312013123 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003112032023133-2320302323030301-2130300123231223-2011322102330210-1031030120021030-2202230111231303-3203011032011222-1021233003022232"></a>

## origin_servers — origin_servers / 210033021313 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- origin_servers

<a id="canonical-1202000233133133-3013023232001221-0213032112311332-2212331201233022-0011000211132213-3032110030112230-2120311230221030-2001112121001331"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

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

<a id="canonical-1002320313212030-0221031330201123-0033032200200210-2333000303230110-2020030312303023-3103112332202000-2011302021123321-1202120120330302"></a>

## Direct properties — origin_servers / 210033021313 / 3

- [cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-1131022223301121-0113110032012312-3110321130113110-1311223333131310-0030021302100133-1223302232123022-1321310233220231-3103231012021331): complete subsection reference.

- [consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101): complete subsection reference.

- [custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231): complete subsection reference.

- [k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201): complete subsection reference.

<a id="canonical-1102110302003202-2233112110210002-0203310333030000-1232303012101122-3200332233322221-1200132132012213-0003110321322113-3113303101332101"></a>

<a id="canonical-0231111213223121-0012000022302312-3210003121023013-3210330133031233-3310212102120002-1223332302012101-0010333023101011-3230132301311003"></a>

## labels property — origin_servers / 210033021313 / 4

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

- [private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102): complete subsection reference.

- [private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110): complete subsection reference.

- [public_ip](data-sources--origin_pool--reference--group-002.md#canonical-3233223332020121-3100021120202320-2031230003101211-3023100123103102-0220333130210323-2201111232211122-1033211313103113-3001033210113321): complete subsection reference.

- [public_name](data-sources--origin_pool--reference--group-002.md#canonical-3023121012300320-1113333033031232-1003130321112112-1203102212132000-1310200323133230-2031320320302202-3121312332032031-2323130012131111): complete subsection reference.

- [vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202): complete subsection reference.

- [vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330): complete subsection reference.

<a id="canonical-0300103031121210-1033001212310000-0020110102310313-1330100331110130-0223300221133021-1330330123031000-1112001203323000-0122002002201102"></a>

## Next pages — origin_servers / 210033021313 / 5

- [origin_servers.cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-1131022223301121-0113110032012312-3110321130113110-1311223333131310-0030021302100133-1223302232123022-1321310233220231-3103231012021331)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.public_ip](data-sources--origin_pool--reference--group-002.md#canonical-3233223332020121-3100021120202320-2031230003101211-3023100123103102-0220333130210323-2201111232211122-1033211313103113-3001033210113321)
- [origin_servers.public_name](data-sources--origin_pool--reference--group-002.md#canonical-3023121012300320-1113333033031232-1003130321112112-1203102212132000-1310200323133230-2031320320302202-3121312332032031-2323130012131111)
- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202)
- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1131022223301121-0113110032012312-3110321130113110-1311223333131310-0030021302100133-1223302232123022-1321310233220231-3103231012021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021313333031100-3002312210320131-2000131002002311-1333031301332031-0213320300332322-1333022022123303-1033130320003022-0011222023100202"></a>

## origin_servers.cbip_service — cbip_service / 311313302021 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.cbip_service

<a id="canonical-0233032331300323-3222131020331213-1312130022210001-2212330232123302-1033120331010313-2201020010203311-0001110000200022-2101001211221212"></a>

Type: `"single"`. Computed.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0300211213331222-0010122332023001-3333232230011321-0300312232021311-1203222033013013-1203011003010302-2201231031002312-0321333023300330"></a>

## Direct properties — cbip_service / 311313302021 / 3

<a id="canonical-3333012200100220-3211011203300311-0023133011123002-3000100012303212-1122110200000110-2330203112020310-2021320000210013-1232111223310102"></a>

<a id="canonical-0300303130103211-3332323223201202-0012202222100013-0322213020110203-0303203111032203-1231330000223020-1010230300322221-0012132202122222"></a>

## service_name property — cbip_service / 311313302021 / 4

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

<a id="canonical-0030000010213320-1110113313221323-2211312333323302-3111032312022221-0033132000123331-2002322100320320-1222222131002232-1323333233023332"></a>

## Next pages — cbip_service / 311313302021 / 5

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012303011302210-3121102302031133-2323130310031201-0002001220032331-1322033331232203-0030330212011333-1021113122301301-1301130033201032"></a>

## origin_servers.consul_service — consul_service / 222303003123 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.consul_service

<a id="canonical-0210310100023010-3001201121100323-2233111213111200-1110031130123110-1013203030001320-0312331000200223-2023121102221301-0122302120030023"></a>

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
