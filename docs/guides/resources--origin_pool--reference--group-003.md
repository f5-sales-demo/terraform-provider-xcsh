---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0012301331302022-1303130000001202-3200103020002222-0320223322203303-0320201320021013-2222221200112120-2200011131132120-2020211101323323"></a>

## `origin_servers.vn_private_name.private_network.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3203203021110231-0111012212301333-0231202321023031-0200023321203132-2122323200121130-1202232210232212-2112323113223002-3222230212232320"></a>

<a id="canonical-0303321311102333-1230110211213001-3123232212213210-2333121021122233-1223223203203332-0330001220230130-3311303133033201-0013012223113021"></a>

## `origin_servers.vn_private_name.private_network.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110220132232202-3223311123032232-3333333030113010-2133200310011011-3010332223210221-3113000210333321-1221001103212033-0210001201112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- same_as_endpoint_port

<a id="canonical-2003120102323102-0320322222233123-0110310222321212-1303023300232011-2231113032010222-0322332011312201-0332033220131211-2002332200121120"></a>

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
same_as_endpoint_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322321330320022-2021311033121221-0210313231232031-3331330210233122-2333300130103032-3230223313322102-2023330330022001-1110122123303032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- timeouts

<a id="canonical-0000231121303330-2003202121011221-0020300020231001-1223302100322112-1221333321301112-3331212200332003-2111332002333011-3211101330111331"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222033010120310-0112013213211130-1311130211110030-1001312102332212-1210033102013201-1102332001012213-1211213131200033-0211013102003032"></a>

### Direct properties for `timeouts`

<a id="canonical-0322021132111321-3112303003031330-0130110311120000-2323112211313121-2012101131003231-2310330113310310-0320302320102011-2221132032313132"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1123132211211221-0212322301210311-3001010203232213-1212101120321000-2222310231130132-1231023213030003-0031133032232203-3212022030111323"></a>

<a id="canonical-1033323101210003-0330031130031231-1220133032102130-3330202302223232-3011223311212223-3323012303020112-2301212100130213-3123313001310010"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0112102231203010-1002213233010322-3013322001331311-1200201112200033-3200131003022301-2112112000322322-2232133101312103-0233312112020123"></a>

<a id="canonical-0200310130121112-1303133311330033-3120301030230323-3123012032013021-0202232101001212-3021321100201000-3200101222010201-2003022222232221"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0110313103330301-1331101223223003-0022300330300231-0323203311011100-3322001332200002-3310012203101213-2201112330320131-0311333030233120"></a>

<a id="canonical-0102012122130023-2202211032033233-3011000210113033-3111223021201001-3121123211130323-3132211320121101-1231121030320331-1322313120233131"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- upstream_conn_pool_reuse_type

<a id="canonical-2020113323212003-1003032002212201-3230111332303132-1132100131321213-1302302110200330-2111203102330131-3033100023112000-2101313201120202"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230212333110220-1133211332032130-3210330102110023-2003103031220011-3022221302312033-1112032310001333-3332212313032332-1003322103001123"></a>

### Direct properties for `upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323): complete subsection reference.

- [enable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301): complete subsection reference.

<a id="canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1302332230212100-0332010233013101-3300102301101123-0002201113311001-1130033132122002-0201100203010021-1301330033303021-3000221202311030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-2101101012002111-0212122133332312-2211023230011231-2323313110333203-3111100301202012-3302031232001111-1000203100112302-0033120332203101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- use_tls

<a id="canonical-3122122211032132-3132313203100001-1013000203202031-2132002223311123-2211000023310210-0233300322102210-0033100013333133-1022321012001012"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033111310100321-1010320213112001-2222323331120000-1211223131302000-3330300130013021-3120200201102310-0013031101223222-1203111013033300"></a>

### Direct properties for `use_tls`

- [default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032): complete subsection reference.

- [disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112): complete subsection reference.

- [disable_sni](resources--origin_pool--reference--group-003.md#canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103): complete subsection reference.

<a id="canonical-1010332301301203-1302100023030012-0012312120112131-0223010111101231-0301233110003032-1101022330100330-2331001011130013-2200321301310220"></a>

<a id="canonical-0103333122330111-0011013233212321-1121102300232223-2223312321230013-1100230310222320-0300130333300200-0100122203121232-0323211000211331"></a>

#### `use_tls.max_session_keys` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [no_mtls](resources--origin_pool--reference--group-003.md#canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221): complete subsection reference.

- [skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201): complete subsection reference.

<a id="canonical-1220200011310322-1313110000203111-0131320231230110-1312300003202311-2021312133301200-1203030202220211-2233131120002033-1310221211332200"></a>

<a id="canonical-3222033320102121-1120201033110323-1222123210103102-3013032023300211-2022323103210310-0301301132031001-1003022220110021-0233133302000103"></a>

#### `use_tls.sni` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131): complete subsection reference.

- [use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022): complete subsection reference.

- [use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011): complete subsection reference.

- [use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110): complete subsection reference.

- [use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002): complete subsection reference.

- [volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110): complete subsection reference.

<a id="canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.default_session_key_caching

<a id="canonical-3030232203023330-2211203000103303-0030123222002110-1210003113321012-3223303023211312-0300132233212333-0312103033022102-0100002311022323"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
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
default_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_session_key_caching

<a id="canonical-0133330110333312-2101233020121221-1212201013003111-0213300012112330-3332101120100012-2122111011331212-2102213012002122-3212032202130203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_sni

<a id="canonical-3121221222210030-0011023312311000-0213231013120333-3321330130232201-1122111133021030-0322220021233033-1301012022233032-0331111232122201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.no_mtls

<a id="canonical-3011203232233111-0210021203020023-3032203220101031-1110110203331323-3133231000132032-0101010211121311-1223020302132201-1013232312121222"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.skip_server_verification

<a id="canonical-2222233003100111-0302003010100213-2011310100011121-3002311212312332-2311313311313130-2211231120332133-1223313220232221-0113201012312301"></a>

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
skip_server_verification = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.tls_config

<a id="canonical-1020312000221113-0332112223333302-3122020211012330-2001011231100112-2100330310322121-0023311223022300-1202133123232313-2202220123133300"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210312202132303-3201213002011322-0221132123010203-1300003333102100-0312002311021132-1223330120310303-2112230201210000-1313332011102131"></a>

### Direct properties for `use_tls.tls_config`

- [custom_security](resources--origin_pool--reference--group-003.md#canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231): complete subsection reference.

- [default_security](resources--origin_pool--reference--group-003.md#canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312): complete subsection reference.

- [low_security](resources--origin_pool--reference--group-003.md#canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110): complete subsection reference.

- [medium_security](resources--origin_pool--reference--group-003.md#canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221): complete subsection reference.

<a id="canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.custom_security

<a id="canonical-0203303001311310-2322002212132332-1010120123231232-1113000203111122-2201232211310022-1032120102202233-0332023301120102-1323302313333132"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223322230311021-0323220302223212-2010213221232002-3320313211230223-1211030230323310-3011020031333211-0310013322203110-2010010010300332"></a>

### Direct properties for `use_tls.tls_config.custom_security`

<a id="canonical-3203013203033320-2103302313120031-2110301221221120-1030331013001031-0321112320310212-1020103003323321-2002032100001303-0312103212130330"></a>

#### `use_tls.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0122213110200311-0101230001332320-0130020201122132-3113320102110003-3223310321312132-3232133210333113-2323320131111230-3101002032132012"></a>

<a id="canonical-3112121020131121-1103112213320331-2012030100300110-2001312110111221-1323010301231011-1200211231023222-0010011002310033-3333330301213221"></a>

#### `use_tls.tls_config.custom_security.max_version` property

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

<a id="canonical-1233311302332033-3103001123311321-3313220110300313-3012312233103323-1211312223130202-0123001303100212-2233321301012023-3333203132331202"></a>

<a id="canonical-0210110200022301-2231302301030220-3121302330312103-3321102202131323-0330123322233013-0312030233301310-0210220321023121-3030303130021012"></a>

#### `use_tls.tls_config.custom_security.min_version` property

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

<a id="canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.default_security

<a id="canonical-2200230203011232-1012310231101221-1302222003121322-0000213033303310-1321100332211312-3202202332303220-3101323313203001-1331301111100332"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.low_security

<a id="canonical-2300211011100131-2312131300032321-3133020322230003-3030323201100300-0020003002220020-1031212102213310-2202030002123013-0032322311122213"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.medium_security

<a id="canonical-3220202323330311-0003102032032322-0223033101210321-2013332030011010-2213021230323322-0302133022011001-2330131022102303-1020303313021110"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_host_header_as_sni

<a id="canonical-3010233110232013-0320031332313011-3002333211230103-3232321131201122-1230310031110203-0133221320122110-3131120320021100-2303232111101310"></a>

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
use_host_header_as_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls

<a id="canonical-1011223301223132-2013212102312122-0221212033123111-3331012332113120-1212201210230330-1203130033210303-0303021021321123-0333022103120211"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030303000323103-0230230023020220-0001103110131003-2302200033012211-2212231002301303-3332222302113211-0113312032310333-1231231312322230"></a>

### Direct properties for `use_tls.use_mtls`

- [tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320): complete subsection reference.

<a id="canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- use_tls.use_mtls.tls_certificates

<a id="canonical-3322232030000201-2023133321112003-1230111113231113-0103010032120300-0011131231103233-2012100003111330-1021213301220103-2120010120303223"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322013210011121-1132221301321101-3012130223222221-0113333300031111-3321102112131331-3003300020320223-2013013223212231-0313012100200300"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates`

- [blindfold](resources--origin_pool--reference--group-003.md#canonical-0030211202303032-3031321103000020-1102100110020031-2313330233301123-3113201012313030-3030012133112230-1031211033123301-1311021332110033): complete subsection reference.

<a id="canonical-1020333333103300-3203001211023011-0331100312033032-1331323202120103-1003032303313233-2322333023311310-0003202103221112-1101012322013010"></a>

<a id="canonical-1030313333323100-1303221130011233-3312333020120031-2223233312133101-0011002231012113-0300321313012332-1131103023310132-2113001133031302"></a>

#### `use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210): complete subsection reference.

<a id="canonical-2313010232001222-1321021230103222-1000002003103000-3111022130302323-2211203311020331-1320202010132122-1013002033320032-3122310031333311"></a>

<a id="canonical-2210313213203032-0203210303012001-0103003131210000-3303123200232121-1331000231211001-0222133312023201-1233200020213311-2203312331122021"></a>

#### `use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222): complete subsection reference.

- [private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312): complete subsection reference.

- [use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113): complete subsection reference.

<a id="canonical-0030211202303032-3031321103000020-1102100110020031-2313330233301123-3113201012313030-3030012133112230-1031211033123301-1311021332110033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.blindfold

<a id="canonical-3331101001101333-0000102231022311-1131210231112322-0323322113203002-0123212102031302-2022032221032013-3212332033303331-2110030301132231"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3202330100120002-0013011230200303-0002332001212221-3131100022110200-3133211322001010-3031010033013030-2200012303313110-0023330130133021"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.blindfold`

<a id="canonical-3301300132021022-0223302320010120-1013031032000101-0030221212020220-1120303302031133-3233300230013233-0233131320012102-3131211022003111"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2000300101212202-0302131320131312-0001302201200113-1113202013020310-2020122213221013-2330321313003202-3312113123103120-3113000101332102"></a>

<a id="canonical-0131101130230110-3302212302332003-1303012232321100-3221312010321101-2031211100312230-2322121322331033-1301112113200221-0121101023013121"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-3113332113231123-1010322121013031-3111003011033220-3332021002220003-1132021223212202-2110303011200330-2302030310332030-2321233323222312"></a>

<a id="canonical-3200121032213003-0132201131130130-2232130031303132-1020032330233201-3322012123311202-2212021313220032-0201301221021200-1003012233001331"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-1033131020322210-1220000322112032-3231313222021222-3010301022302010-2113032001102223-0302312220020113-3230113313323001-1100030223213312"></a>

<a id="canonical-2223323323031111-2200133323233212-2320203220112300-2102020030321330-2032331331031232-0312020110303101-1100130231100123-3231130022331013"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-3300303130013121-3221310221021310-1023101123301233-1003032112020310-0001301120231210-0301211300212332-2110221320312021-3301023013203003"></a>

<a id="canonical-0223330222003031-1220000211333321-3210030303213011-0121203301132302-1213112201223333-3111220313230333-2031031011010101-3001303212333100"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-3131021121323021-3202213101030310-3330213331213300-2320303333231122-2330203003333300-3133201030003212-0333300021321203-2212021031323301"></a>

<a id="canonical-0103332131333232-0110310310123011-2202012301321330-3332321230020131-2001320210100112-1110123102021212-1002230100010323-3133202123110023"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0231322202121003-1012332003013301-1111022021212221-1332000201321223-2033012123321000-1103313312230111-1301131220101102-2230222100030213"></a>

<a id="canonical-1220233332333333-0110010212301223-0030220013210103-3002123220111230-2310332331011323-0003303022312213-2332130013000321-2321300003102023"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1303332301120013-2303132200221201-2232100310233202-2000101200301333-0323020221120331-2311331031102021-0202030023201133-0300123121001203"></a>

<a id="canonical-0332130210230110-2222133303022100-3013311110331301-2221311211010222-3110201322213230-2302123201133000-3232113002322120-1313300301220210"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-2332000022100203-1230320212321223-2313202132213333-1313220221130320-0011131110332130-2131232323302120-3220203333000201-1232131331113122"></a>

<a id="canonical-1130011300032222-1331311013110330-0002330032032302-1122311233313003-2101310210022023-2022332312312022-2121020021201203-0103232231332223"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3231011311303211-0032213331200102-2300102100221011-1012232020233113-2010202013132233-2201120222030012-0221300002003122-2331113312313201"></a>

<a id="canonical-1201302300210203-1100320201223333-1112132310322012-3330223033212230-1310300113300010-1232112022210233-2203111311130100-2110212331212012"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3200100210203320-1130121310000321-3032123201002130-3322203022013130-3120021020302201-1001121312112213-2201131020122011-1011101100110103"></a>

<a id="canonical-2122103111020031-0211111113133220-2202011131100303-0231121123032332-1223203220211113-2232312331302032-0013102323330032-2312220203311303"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3031321130210011-1311130130203121-0032321021320302-1313203001001032-2213030010310201-3230011212130323-0111201000310013-0112120122100030"></a>

<a id="canonical-2021010332202332-3101303302230232-0213023323113313-2130000323200001-1011030301301011-2321110302230031-1233123330221022-0231010202103222"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1301102202222213-3113120330020303-1103300022020030-2101213221013120-2212211122013202-3123120212312003-2213002210122112-0023323231133313"></a>

<a id="canonical-2301210101222212-0213020232123111-3021201033101021-0333231111132200-0213231131202221-1221122223322211-2313112203121013-1311331020012302"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2233102212302202-2323021101321223-3101022101132000-1203212313211310-1323113100021112-0332222221132013-2303132100223312-0011312200032023"></a>

<a id="canonical-1222231113010302-1323231031022131-3210201010130221-3130222331112200-1322112311323002-0322232331212230-0302313332021200-2313210200011333"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0000310110011312-3222332133331100-3210222020203213-2303010332002022-2112031132200233-3201100123333232-1021211233202123-1033222203002233"></a>

<a id="canonical-0202000030202011-0232131202210133-1100113220321030-0213303020302330-0302102121321030-0220200322310210-1001230230111220-3210102103312121"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-2213023322301230-3310010233021330-1210322321323120-1321130021111332-3321302333300230-1233333200030230-0330311022302301-0323210111232230"></a>

<a id="canonical-0300133000303122-3210133330300101-2031113302002210-2003000322333030-0123201221322122-2001211232103313-3020113311121310-2100100122132313"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-2213303331323303-2220222322323101-1011300211032030-3023031033313011-0000302213311122-2010022213313032-2112133212013311-0232222203321220"></a>

<a id="canonical-1113320313210000-1310112000311130-3302101031322122-2230102011120011-0312212330233210-1130121301032000-3300200031332031-3303112002000312"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-1100113020321203-1302120333130223-3220121123133232-1120132231023313-1101002231211033-0301120233020331-0223333133023302-2231320132110132"></a>

<a id="canonical-3222323301011333-1203303020311130-0101033023320303-3231211023101300-0013231231232311-3210002030020210-0223322133031322-1000133202312132"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2120130203101221-2131110301320303-1021033000120233-0132330210211001-2232212123130200-3321100200100331-3013010230233301-1311123120232233"></a>

<a id="canonical-3313202210123011-0011222103320310-2001010331202133-2033112331300112-0030112100031200-1102130310200023-0312120012313323-2231330320230232"></a>

#### `use_tls.use_mtls.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-2213223211110223-3021301202303321-3000101311331301-0110113113100000-0031030230003220-2102010111213120-0332200330001012-2310011311303112"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2330221122302013-2110222132001213-2010031001221220-1100100202312111-1011212332031120-0013111321133300-1000332231002313-2213303202132010"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-0321110130030203-3123231300301322-0100330000011001-3203112331003102-3021331301133322-0300331222030123-2231130120300120-3121221330333033"></a>

#### `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-3333300201211233-1032301121302000-2022223332122122-2223111331102203-1323103332103020-3302230030122232-1212130120321021-1111203300330332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2312122300111210-2320203233333123-0110010220000333-2133221212322120-3332123321213233-2112300310000012-2032231020312032-0001321213022110"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1300101123330201-0030121031103211-3332333212100113-0020330033031222-0102310102302330-1321330222232120-0002102301022031-1320233013223302"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000): complete subsection reference.

- [clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000): complete subsection reference.

<a id="canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1321030203013330-2010233100103102-0122331220200012-1133022023030330-0113333100231223-3131102101133033-2303111220320013-3013030233123210"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1203211223020122-1111122122021333-1132001031023320-2222130031212230-2110322311132102-3013130233102120-3223220012210100-3221102320312013"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2013331301120103-2210112033020212-0312000203011300-3120003131103213-3203013231030210-3330031202200330-3233002212013021-0301223113122001"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102221000300000-1010023013101110-0132000101123120-1131113230111302-2213302333033232-2030201122210021-2213233022022031-3030001212001030"></a>

<a id="canonical-0021033101313312-0121320300300032-2123221032023323-1120202112033012-0301210200103122-3102003200312133-2123022233211323-1103232203333112"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033311012223112-1130122113201000-2123313110033211-0330223023223333-2212033331030011-3203322312323213-2232033003201121-3013302131310202"></a>

<a id="canonical-2303230222000201-3213001300103022-2221120031101221-3012222203232132-2202222313002201-3101302201013233-1123202022230333-2210022001220231"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-2111132313310011-2312201030120001-1123021332221032-2101010113212013-3212022320002303-2313331301003330-2331210133313200-3023131301312032"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1033112132010113-1110233011330131-2302021002011121-1122030112123103-0032031221112000-0002012100223310-3202221001320311-3311213103301011"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0033032221002213-2122323020200123-2232322102333232-0323201332200320-3122021310323232-0230012231030103-3132210321021211-2012200313310033"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1211332112310200-1232002021223003-3120011200102233-1332111032220302-2322202210213020-0200202211320031-2010220010302203-0221300031211030"></a>

<a id="canonical-1212301102231203-1310103301311220-2110330031120100-1233300201222213-3133103130003301-1212310323312133-2330033230001020-2103312030020302"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0301312001311001-3013220001311121-1010021310302213-1333300031230210-0110113323013013-3201201200133221-3332230303220222-3303022020031233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls_obj

<a id="canonical-0311020032113212-1203103100333300-3120011330330310-2133030131113310-1200133213303001-0103330110212320-0311211222012323-1000301320200020"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301322212311213-3230032230200211-3230012231133220-0331102331011031-0122201231032000-0110312310013223-1323213113000022-0212031323002112"></a>

### Direct properties for `use_tls.use_mtls_obj`

<a id="canonical-0231211130302032-2202100003223131-3133021022010120-1031222032311213-3310111011332112-3103221020013201-2111031200302330-0323030313333001"></a>

#### `use_tls.use_mtls_obj.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3210200223021301-1323211003102302-3230323233302133-3221102000111102-3223120130322211-2001232023212122-0233312112011310-1233230002321120"></a>

<a id="canonical-2301311312303232-3213212013330200-0320100033230113-3023201320322112-1212112220100310-0211013032230302-2032321130030302-3300201311221222"></a>

#### `use_tls.use_mtls_obj.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3200110220221112-3132202102011101-3023302001331222-0320032203131012-2132103012300313-0101322313011012-1233011231000111-1333011302033102"></a>

<a id="canonical-2013313221011112-2230020132313131-0031211202202220-0302123022302021-3010220203022332-1103101220332000-1032202102312211-0112100133333130"></a>

#### `use_tls.use_mtls_obj.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_server_verification

<a id="canonical-1202022113131322-3223333122031311-2012231003203110-2200120033103313-0303123113012001-1021131200222210-3302220021012302-2121023331223031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

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
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230130300220332-2220112133333010-2330133300102111-2303113300000313-2302211333123123-0223002301021111-1211003112223033-3120311222001311"></a>

### Direct properties for `use_tls.use_server_verification`

- [trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211): complete subsection reference.

<a id="canonical-1220133023112011-0133103323321013-2230232111003000-1311033011331123-2002332230231132-0210313333022312-1133232021122222-0311333013311320"></a>

<a id="canonical-2223100100003123-1232231022032302-3300010033210300-0213101020200111-1221002032320301-0303223221331132-0331022123113131-3112211221003130"></a>

#### `use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002)
- use_tls.use_server_verification.trusted_ca

<a id="canonical-3211120110011201-0121003012003331-3200121102030103-3100212002221031-3100111133302033-0212211223332131-0003023220323223-0303330120232302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310322322300221-1013011011301211-0131221202310311-2130033322231100-3233111220133121-1201333021223121-0210002100210021-3333031121020132"></a>

### Direct properties for `use_tls.use_server_verification.trusted_ca`

<a id="canonical-2300000132123033-2103011001320010-0003201212202311-2131202021300121-0100301332031130-3100031203003111-3213000321133131-3220001111032321"></a>

#### `use_tls.use_server_verification.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0121211112011311-3022002120201003-1200133210002120-3103310112300321-1331332330013221-1322110233131121-2221323113132132-0330103213300331"></a>

<a id="canonical-1230303213202211-3310231021222022-3132223300113103-0211030113012310-2321102002112223-2212220323331023-3202003113131111-0321032200210332"></a>

#### `use_tls.use_server_verification.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3323101123113223-3112300031313201-0323233213330222-1122012021031001-1213311231311210-1321012323131120-1220113301313301-2232122310332010"></a>

<a id="canonical-1320222031003013-1333202031030121-0110103021102000-0112102233101131-1313202221310320-1212213301130112-2313202002233303-0111131101013213"></a>

#### `use_tls.use_server_verification.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.volterra_trusted_ca

<a id="canonical-1002302310301002-1302021231211331-0023032120010312-1013223211301323-1323121220130010-2023101130223113-3023221113032221-2012303103232033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.
