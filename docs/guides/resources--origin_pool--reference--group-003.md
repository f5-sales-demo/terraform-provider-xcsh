---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0013023223022220-2123102012330222-1303111232221330-1132212122132321-0033302310122232-3202312332011201-1102233003122121-0013023202131123"></a>

## Next pages — timeouts / 321233111120 / 8

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230212333110220-1133211332032130-3210330102110023-2003103031220011-3022221302312033-1112032310001333-3332212313032332-1003322103001123"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 031202022221 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- upstream_conn_pool_reuse_type

<a id="canonical-2020113323212003-1003032002212201-3230111332303132-1132100131321213-1302302110200330-2111203102330131-3033100023112000-2101313201120202"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101300202301313-2202102101121313-2033131212031303-1011331223010213-3012213010211222-3100121232222120-2331312310313031-0110012231231101"></a>

## Direct properties — upstream_conn_pool_reuse_type / 031202022221 / 3

- [disable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323): complete subsection reference.

- [enable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301): complete subsection reference.

<a id="canonical-2311201330213103-1311012131012330-0102321303113331-1130022333330022-2313220132320120-2310233120323221-3230323113003232-3321333232211210"></a>

## Next pages — upstream_conn_pool_reuse_type / 031202022221 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302000002330102-0012233212203220-3221201113230220-2220230233103210-1023232332112230-0200032302023121-3031211212033312-1002312121002132"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — disable_conn_pool_reuse / 100210001000 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1302332230212100-0332010233013101-3300102301101123-0002201113311001-1130033132122002-0201100203010021-1301330033303021-3000221202311030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

<a id="canonical-3001132110223230-3211001223121323-1323030300211221-2001032033220303-2011030000012002-0233322200022031-0210212230313032-1011220211110021"></a>

## Direct properties — disable_conn_pool_reuse / 100210001000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220203013101300-1133320100232313-1032310222310001-3221201213313302-0312103211222011-3110010132332333-2013101322120103-1033311112103033"></a>

## Next pages — disable_conn_pool_reuse / 100210001000 / 4

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110231023303111-3112202101212003-2112120303222102-2220212312312232-1333203232100221-0001211322001132-1202302011200220-1110120301002330"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — enable_conn_pool_reuse / 303022003210 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-2101101012002111-0212122133332312-2211023230011231-2323313110333203-3111100301202012-3302031232001111-1000203100112302-0033120332203101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

<a id="canonical-0303111031321331-2222003002010113-0223332330233000-0301330201203000-0311201213303101-3202310312122022-0333200212333323-1013312323033132"></a>

## Direct properties — enable_conn_pool_reuse / 303022003210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003211013102003-3013020311313123-1103210320122332-1330022211123130-1220321222101310-1000303300112233-3230201101331222-1231321212220330"></a>

## Next pages — enable_conn_pool_reuse / 303022003210 / 4

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033111310100321-1010320213112001-2222323331120000-1211223131302000-3330300130013021-3120200201102310-0013031101223222-1203111013033300"></a>

## use_tls — use_tls / 000231113320 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- use_tls

<a id="canonical-3122122211032132-3132313203100001-1013000203202031-2132002223311123-2211000023310210-0233300322102210-0033100013333133-1022321012001012"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
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

<a id="canonical-0103333122330111-0011013233212321-1121102300232223-2223312321230013-1100230310222320-0300130333300200-0100122203121232-0323211000211331"></a>

## Direct properties — use_tls / 000231113320 / 3

- [default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032): complete subsection reference.

- [disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112): complete subsection reference.

- [disable_sni](resources--origin_pool--reference--group-003.md#canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103): complete subsection reference.

<a id="canonical-1010332301301203-1302100023030012-0012312120112131-0223010111101231-0301233110003032-1101022330100330-2331001011130013-2200321301310220"></a>

<a id="canonical-3222033320102121-1120201033110323-1222123210103102-3013032023300211-2022323103210310-0301301132031001-1003022220110021-0233133302000103"></a>

## max_session_keys property — use_tls / 000231113320 / 4

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

- [no_mtls](resources--origin_pool--reference--group-003.md#canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221): complete subsection reference.

- [skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201): complete subsection reference.

<a id="canonical-1220200011310322-1313110000203111-0131320231230110-1312300003202311-2021312133301200-1203030202220211-2233131120002033-1310221211332200"></a>

<a id="canonical-2321100102030032-0310321120323212-0112221023001301-0000231103222020-2122301033022122-0213120100312130-0320010102203010-1032020010023221"></a>

## sni property — use_tls / 000231113320 / 5

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

- [tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131): complete subsection reference.

- [use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022): complete subsection reference.

- [use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011): complete subsection reference.

- [use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110): complete subsection reference.

- [use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002): complete subsection reference.

- [volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110): complete subsection reference.

<a id="canonical-2230333223023200-1322010213302101-1233113121202330-2211330113213200-0020230302212313-1231110011303012-1332320330210123-3201302221131202"></a>

## Next pages — use_tls / 000231113320 / 6

- [use_tls.default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032)
- [use_tls.disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112)
- [use_tls.disable_sni](resources--origin_pool--reference--group-003.md#canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103)
- [use_tls.no_mtls](resources--origin_pool--reference--group-003.md#canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221)
- [use_tls.skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- [use_tls.use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002)
- [use_tls.volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120000320210333-0221003133111032-2233303103103210-0133102203112311-0120102103021031-2330321203113030-1212231321103122-0201211110111223"></a>

## use_tls.default_session_key_caching — default_session_key_caching / 301110310110 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.default_session_key_caching

<a id="canonical-3030232203023330-2211203000103303-0030123222002110-1210003113321012-3223303023211312-0300132233212333-0312103033022102-0100002311022323"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
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
default_session_key_caching = {}
```

<a id="canonical-3020100010212332-0013031321123023-2200023100012232-3222100232200300-2113033230012010-3003200223232213-2111332120003032-2002221032001102"></a>

## Direct properties — default_session_key_caching / 301110310110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321203320012303-3000232021203213-0020020330233232-3212002013120131-1323212221020131-3022213021310022-1222213231221222-0231213221222121"></a>

## Next pages — default_session_key_caching / 301110310110 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332123133212003-3320101023100320-0232231010300112-0101212020030311-3023112233001133-3111132003230002-2331133121302232-2131210231122103"></a>

## use_tls.disable_session_key_caching — disable_session_key_caching / 212121132311 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_session_key_caching

<a id="canonical-0133330110333312-2101233020121221-1212201013003111-0213300012112330-3332101120100012-2122111011331212-2102213012002122-3212032202130203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

<a id="canonical-1300021022222231-0113112303133002-2322230210033320-0311302310133322-3031313133200031-3002022010031132-2030221210310300-0222122201320032"></a>

## Direct properties — disable_session_key_caching / 212121132311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000030101301000-3002200100322122-1300130002302320-3110012022011010-2333201200020021-1321030210122300-3322223001113231-0012320021012031"></a>

## Next pages — disable_session_key_caching / 212121132311 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003301113120230-1312300021303313-3023121122310032-1101203300210101-2203123102032000-1100112013202302-2000230033311330-2333201323002300"></a>

## use_tls.disable_sni — disable_sni / 001210221123 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_sni

<a id="canonical-3121221222210030-0011023312311000-0213231013120333-3321330130232201-1122111133021030-0322220021233033-1301012022233032-0331111232122201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-1200203202222313-1203023201113233-3213321320300331-1321012121011011-2102032032130031-1200231133122202-3133011013212320-3122123303112102"></a>

## Direct properties — disable_sni / 001210221123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102110223201003-1002313120300100-2031212200101010-2331330210231310-1030100223031322-0011020012132123-2121210223131032-3122212222310221"></a>

## Next pages — disable_sni / 001210221123 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203033312202033-2302023222233010-3031201000333203-0001103211331013-2130233301133221-0100102220201313-0301133212211031-2002230213201101"></a>

## use_tls.no_mtls — no_mtls / 210111303121 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.no_mtls

<a id="canonical-3011203232233111-0210021203020023-3032203220101031-1110110203331323-3133231000132032-0101010211121311-1223020302132201-1013232312121222"></a>

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
no_mtls = {}
```

<a id="canonical-2131002120223110-0020320032131303-2030232123332123-0201203012312012-1101131332312110-3203031120220110-1223230033233300-2321013233020021"></a>

## Direct properties — no_mtls / 210111303121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221211122031333-0032202122012222-3323113321212122-1321033021211212-2321223212020120-2333221210221122-0312210132212312-3330002100001203"></a>

## Next pages — no_mtls / 210111303121 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022233033221030-1031200100222313-2213333221113013-0231032103330031-3001101010303131-0311102330003302-3232311303201202-2130111331211003"></a>

## use_tls.skip_server_verification — skip_server_verification / 221223220323 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.skip_server_verification

<a id="canonical-2222233003100111-0302003010100213-2011310100011121-3002311212312332-2311313311313130-2211231120332133-1223313220232221-0113201012312301"></a>

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

<a id="canonical-0033101333112121-3003330300323021-0131033203121031-1231010211110120-2300202010321110-3232210210200031-3323322320310133-0223231203233213"></a>

## Direct properties — skip_server_verification / 221223220323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213310332131023-2111221303230131-1013213131312202-2320110132012220-3101131333012122-2010130121010030-3012002203011033-3300130001320203"></a>

## Next pages — skip_server_verification / 221223220323 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210312202132303-3201213002011322-0221132123010203-1300003333102100-0312002311021132-1223330120310303-2112230201210000-1313332011102131"></a>

## use_tls.tls_config — tls_config / 213012101010 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.tls_config

<a id="canonical-1020312000221113-0332112223333302-3122020211012330-2001011231100112-2100330310322121-0023311223022300-1202133123232313-2202220123133300"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2232302023230100-3030323200201132-1320010101112232-2102110132101010-1120212301300101-2310123301022101-1203333100332311-3030303201212003"></a>

## Direct properties — tls_config / 213012101010 / 3

- [custom_security](resources--origin_pool--reference--group-003.md#canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231): complete subsection reference.

- [default_security](resources--origin_pool--reference--group-003.md#canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312): complete subsection reference.

- [low_security](resources--origin_pool--reference--group-003.md#canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110): complete subsection reference.

- [medium_security](resources--origin_pool--reference--group-003.md#canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221): complete subsection reference.

<a id="canonical-1222030303021113-1223233203232111-2223012312320321-2102013013332332-1333103330232033-0221112013232110-2023012231022201-3220131110032223"></a>

## Next pages — tls_config / 213012101010 / 4

- [use_tls.tls_config.custom_security](resources--origin_pool--reference--group-003.md#canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231)
- [use_tls.tls_config.default_security](resources--origin_pool--reference--group-003.md#canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312)
- [use_tls.tls_config.low_security](resources--origin_pool--reference--group-003.md#canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110)
- [use_tls.tls_config.medium_security](resources--origin_pool--reference--group-003.md#canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223322230311021-0323220302223212-2010213221232002-3320313211230223-1211030230323310-3011020031333211-0310013322203110-2010010010300332"></a>

## use_tls.tls_config.custom_security — custom_security / 030210302302 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.custom_security

<a id="canonical-0203303001311310-2322002212132332-1010120123231232-1113000203111122-2201232211310022-1032120102202233-0332023301120102-1323302313333132"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112121020131121-1103112213320331-2012030100300110-2001312110111221-1323010301231011-1200211231023222-0010011002310033-3333330301213221"></a>

## Direct properties — custom_security / 030210302302 / 3

<a id="canonical-3203013203033320-2103302313120031-2110301221221120-1030331013001031-0321112320310212-1020103003323321-2002032100001303-0312103212130330"></a>

<a id="canonical-0210110200022301-2231302301030220-3121302330312103-3321102202131323-0330123322233013-0312030233301310-0210220321023121-3030303130021012"></a>

## cipher_suites property — custom_security / 030210302302 / 4

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

<a id="canonical-3221210302130321-1101101333331121-2222331013001312-0310211001000003-2023002323311112-2310302321333030-1310120032211113-3320300030200212"></a>

## max_version property — custom_security / 030210302302 / 5

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

<a id="canonical-1233311302332033-3103001123311321-3313220110300313-3012312233103323-1211312223130202-0123001303100212-2233321301012023-3333203132331202"></a>

<a id="canonical-3101213130101213-2130302122030332-3111212020123021-2221211310230200-0010312001133312-0233222221133210-3201233200220100-1022221112133001"></a>

## min_version property — custom_security / 030210302302 / 6

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

<a id="canonical-1022223322232030-3031212110030003-1011301203221012-3302301132223131-3301332231103203-0313002120111113-1031020111113002-3100202300131122"></a>

## Next pages — custom_security / 030210302302 / 7

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100030233133133-2113202330132311-1002000210220113-2310321010303100-1231221101111230-2010231203211220-0010332122131302-1033321012023130"></a>

## use_tls.tls_config.default_security — default_security / 022112030100 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.default_security

<a id="canonical-2200230203011232-1012310231101221-1302222003121322-0000213033303310-1321100332211312-3202202332303220-3101323313203001-1331301111100332"></a>

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
default_security = {}
```

<a id="canonical-2131013222233000-2233323232121013-3010232332232300-1100122023112030-2030033210002210-1022020202000112-1201133200101013-0000132331203232"></a>

## Direct properties — default_security / 022112030100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322213000301212-1200021200113312-1212031330123030-3230112112222313-2312200232230300-0111320313123110-0001221320220101-1312122212220323"></a>

## Next pages — default_security / 022112030100 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033220122030323-2002223002310223-0031001201033331-3311210231213333-2030000213312132-0221120312221201-2303231313131023-3213233020310312"></a>

## use_tls.tls_config.low_security — low_security / 230332222203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.low_security

<a id="canonical-2300211011100131-2312131300032321-3133020322230003-3030323201100300-0020003002220020-1031212102213310-2202030002123013-0032322311122213"></a>

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
low_security = {}
```

<a id="canonical-2112121332222213-2323321033031120-2033200331223133-2103333201313320-3033112313032220-1221220302033223-3232332301000203-0121001310001003"></a>

## Direct properties — low_security / 230332222203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300300002000001-1131132023302230-2032323023113110-1120232000302311-3323120213330223-3202020311030121-2213221032233212-0300131322101123"></a>

## Next pages — low_security / 230332222203 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020023110202321-0322203210232220-3012302232030232-0223130123202032-1012221330320031-1232113222200203-2130011121233200-2133330132303000"></a>

## use_tls.tls_config.medium_security — medium_security / 311020103120 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.medium_security

<a id="canonical-3220202323330311-0003102032032322-0223033101210321-2013332030011010-2213021230323322-0302133022011001-2330131022102303-1020303313021110"></a>

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
medium_security = {}
```

<a id="canonical-0210131203210113-1133011233133011-3201303321113222-2223203230132231-0211330013233303-0300003223011022-1221202033100200-0300200210010330"></a>

## Direct properties — medium_security / 311020103120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021033100203320-3031023113231131-0203201023132211-0100031310202332-3322112211121231-2332022233011010-2122011032233211-3100230130113110"></a>

## Next pages — medium_security / 311020103120 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002003223210302-3222221232120003-0333002332323323-0311023203021010-3303223011332213-0130302213113212-2232030322312120-1211112032230032"></a>

## use_tls.use_host_header_as_sni — use_host_header_as_sni / 312302011131 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_host_header_as_sni

<a id="canonical-3010233110232013-0320031332313011-3002333211230103-3232321131201122-1230310031110203-0133221320122110-3131120320021100-2303232111101310"></a>

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
use_host_header_as_sni = {}
```

<a id="canonical-2032131000002032-0202313222313330-0102123100030103-1232210221230210-2103202221121113-2232121110220122-2030222012000213-2122112200212212"></a>

## Direct properties — use_host_header_as_sni / 312302011131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113022211233003-0112013312331120-1231322100310011-1132210103002231-1032023030000032-2203300313011332-2113023232221220-2310220002332020"></a>

## Next pages — use_host_header_as_sni / 312302011131 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030303000323103-0230230023020220-0001103110131003-2302200033012211-2212231002301303-3332222302113211-0113312032310333-1231231312322230"></a>

## use_tls.use_mtls — use_mtls / 111230220002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls

<a id="canonical-1011223301223132-2013212102312122-0221212033123111-3331012332113120-1212201210230330-1203130033210303-0303021021321123-0333022103120211"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113203022200222-1212300221103112-0202201312033013-0321322311203233-1031100301220130-0202120021320230-0223233113210103-0013301131223020"></a>

## Direct properties — use_mtls / 111230220002 / 3

- [tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320): complete subsection reference.

<a id="canonical-1301120221222100-2311031331223301-2123301323000331-0223212321002213-3021301312122313-2033020100110131-2110301103022333-3202202221301213"></a>

## Next pages — use_mtls / 111230220002 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322013210011121-1132221301321101-3012130223222221-0113333300031111-3321102112131331-3003300020320223-2013013223212231-0313012100200300"></a>

## use_tls.use_mtls.tls_certificates — tls_certificates / 000132123122 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- use_tls.use_mtls.tls_certificates

<a id="canonical-3322232030000201-2023133321112003-1230111113231113-0103010032120300-0011131231103233-2012100003111330-1021213301220103-2120010120303223"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1030313333323100-1303221130011233-3312333020120031-2223233312133101-0011002231012113-0300321313012332-1131103023310132-2113001133031302"></a>

## Direct properties — tls_certificates / 000132123122 / 3

<a id="canonical-1020333333103300-3203001211023011-0331100312033032-1331323202120103-1003032303313233-2322333023311310-0003202103221112-1101012322013010"></a>

<a id="canonical-2210313213203032-0203210303012001-0103003131210000-3303123200232121-1331000231211001-0222133312023201-1233200020213311-2203312331122021"></a>

## certificate_url property — tls_certificates / 000132123122 / 4

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

- [custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210): complete subsection reference.

<a id="canonical-2313010232001222-1321021230103222-1000002003103000-3111022130302323-2211203311020331-1320202010132122-1013002033320032-3122310031333311"></a>

<a id="canonical-0220032302212111-0310311132222303-3210322031230103-2033003120121331-2321102302003020-1222230020313231-2310022031330320-2210203332222332"></a>

## description_spec property — tls_certificates / 000132123122 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222): complete subsection reference.

- [private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312): complete subsection reference.

- [use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113): complete subsection reference.

<a id="canonical-3013201301023123-3003033111030310-1133212223010121-1001223011323203-3010301331002131-2120313002320000-0211110311321022-1202023121200233"></a>

## Next pages — tls_certificates / 000132123122 / 6

- [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210)
- [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330221122302013-2110222132001213-2010031001221220-1100100202312111-1011212332031120-0013111321133300-1000332231002313-2213303202132010"></a>

## use_tls.use_mtls.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 102312220103 / 2

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

<a id="canonical-3221131103300220-2033012231131321-1101131312023310-1212210223002332-0031030113031303-2300221021332103-0102313212113120-3320203312101032"></a>

## Direct properties — custom_hash_algorithms / 102312220103 / 3

<a id="canonical-0321110130030203-3123231300301322-0100330000011001-3203112331003102-3021331301133322-0300331222030123-2231130120300120-3121221330333033"></a>

<a id="canonical-2323312023002103-1213220210010013-1033213301112021-2222123330112331-2321201203200023-2131012200100320-1213322130132301-2232030211012320"></a>

## hash_algorithms property — custom_hash_algorithms / 102312220103 / 4

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

<a id="canonical-0102122312322112-0223020101001233-3231123213221333-3222000031302110-3310123320321303-2100301300022213-0010211302003210-1201123311300202"></a>

## Next pages — custom_hash_algorithms / 102312220103 / 5

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002103012220002-0312111103002133-0122311333023321-1022020121123031-2311132220203330-3111223031112103-1300220002323232-1233112212310201"></a>

## use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 302310310312 / 2

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

<a id="canonical-0021123301111323-0033213012322022-3203322030022332-1312311222322110-2210202131022031-3111033210221111-3002031231023212-3103222220101211"></a>

## Direct properties — disable_ocsp_stapling / 302310310312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003123101232133-2122303321021332-3122322300321202-2231333121100132-1130111331212102-0230112223322200-1020201300333022-0111000331111101"></a>

## Next pages — disable_ocsp_stapling / 302310310312 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300101123330201-0030121031103211-3332333212100113-0020330033031222-0102310102302330-1321330222232120-0002102301022031-1320233013223302"></a>

## use_tls.use_mtls.tls_certificates.private_key — private_key / 323221020011 / 2

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

<a id="canonical-0203332033123010-0012312323321211-0301121011322300-2313231213322132-2230200220211322-1000120230103333-0333121320112033-2000022102222132"></a>

## Direct properties — private_key / 323221020011 / 3

- [blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000): complete subsection reference.

- [clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000): complete subsection reference.

<a id="canonical-2221033111003002-2231323210003313-0133321222000011-0311032303332300-0330013230231213-0121233323223131-2110033000212231-2031030003203230"></a>

## Next pages — private_key / 323221020011 / 4

- [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000)
- [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203211223020122-1111122122021333-1132001031023320-2222130031212230-2110322311132102-3013130233102120-3223220012210100-3221102320312013"></a>

## use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 232133301323 / 2

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

<a id="canonical-0021033101313312-0121320300300032-2123221032023323-1120202112033012-0301210200103122-3102003200312133-2123022233211323-1103232203333112"></a>

## Direct properties — blindfold_secret_info / 232133301323 / 3

<a id="canonical-2013331301120103-2210112033020212-0312000203011300-3120003131103213-3203013231030210-3330031202200330-3233002212013021-0301223113122001"></a>

<a id="canonical-2303230222000201-3213001300103022-2221120031101221-3012222203232132-2202222313002201-3101302201013233-1123202022230333-2210022001220231"></a>

## decryption_provider property — blindfold_secret_info / 232133301323 / 4

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

<a id="canonical-3102221000300000-1010023013101110-0132000101123120-1131113230111302-2213302333033232-2030201122210021-2213233022022031-3030001212001030"></a>

<a id="canonical-0322001130010023-0031322211202032-0132121131031112-3031123100032122-2212112101011030-1121030031020030-1223113320300331-1011320002000100"></a>

## location property — blindfold_secret_info / 232133301323 / 5

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

<a id="canonical-0033311012223112-1130122113201000-2123313110033211-0330223023223333-2212033331030011-3203322312323213-2232033003201121-3013302131310202"></a>

<a id="canonical-1223122202003222-2300031201310031-0030100333333320-0332022221310311-3213221301020033-3103103101031100-1012211231022012-1131032311033311"></a>

## store_provider property — blindfold_secret_info / 232133301323 / 6

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

<a id="canonical-1200102232321332-1302332130230320-0030101113321210-3332311332211201-2220211310101303-2112100020213210-2233111233021310-1133001133230303"></a>

## Next pages — blindfold_secret_info / 232133301323 / 7

- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033112132010113-1110233011330131-2302021002011121-1122030112123103-0032031221112000-0002012100223310-3202221001320311-3311213103301011"></a>

## use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — clear_secret_info / 101221022020 / 2

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

<a id="canonical-1212301102231203-1310103301311220-2110330031120100-1233300201222213-3133103130003301-1212310323312133-2330033230001020-2103312030020302"></a>

## Direct properties — clear_secret_info / 101221022020 / 3

<a id="canonical-0033032221002213-2122323020200123-2232322102333232-0323201332200320-3122021310323232-0230012231030103-3132210321021211-2012200313310033"></a>

<a id="canonical-2203010002012222-1111130322302030-1321223201023113-3323230322102311-1212223100021212-1133321102120130-1120323102302222-0002111010231301"></a>

## provider_ref property — clear_secret_info / 101221022020 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1211332112310200-1232002021223003-3120011200102233-1332111032220302-2322202210213020-0200202211320031-2010220010302203-0221300031211030"></a>

<a id="canonical-1000300320322323-2212112310303301-0222220120132121-1113111201113212-3303132000332022-2103221210121100-1020001213001301-0211021020333101"></a>

## URL property — clear_secret_info / 101221022020 / 5

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

<a id="canonical-0100223222311131-1222220313222313-2323211122023033-0100323102120100-1323000210202323-3102110010323112-0110031212122213-1121101103322001"></a>

## Next pages — clear_secret_info / 101221022020 / 6

- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101222203023021-0330321132313333-2200111201132031-2322113131001111-1013012233223112-0023202030210322-0303131113212131-3323111310202021"></a>

## use_tls.use_mtls.tls_certificates.use_system_defaults — use_system_defaults / 211323312200 / 2

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

<a id="canonical-2120332303013032-3133320021112333-1011312220011312-0233111012230013-3233231012130102-3302111112201030-3021023010311002-2103231131000302"></a>

## Direct properties — use_system_defaults / 211323312200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230002100033022-3101231133303123-1333010313233202-3333312311100001-2303021131022113-2313201022321002-0011222330111223-3213202110021000"></a>

## Next pages — use_system_defaults / 211323312200 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301322212311213-3230032230200211-3230012231133220-0331102331011031-0122201231032000-0110312310013223-1323213113000022-0212031323002112"></a>

## use_tls.use_mtls_obj — use_mtls_obj / 032223013103 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls_obj

<a id="canonical-0311020032113212-1203103100333300-3120011330330310-2133030131113310-1200133213303001-0103330110212320-0311211222012323-1000301320200020"></a>

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301311312303232-3213212013330200-0320100033230113-3023201320322112-1212112220100310-0211013032230302-2032321130030302-3300201311221222"></a>

## Direct properties — use_mtls_obj / 032223013103 / 3

<a id="canonical-0231211130302032-2202100003223131-3133021022010120-1031222032311213-3310111011332112-3103221020013201-2111031200302330-0323030313333001"></a>

<a id="canonical-2013313221011112-2230020132313131-0031211202202220-0302123022302021-3010220203022332-1103101220332000-1032202102312211-0112100133333130"></a>

## name property — use_mtls_obj / 032223013103 / 4

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

<a id="canonical-3210200223021301-1323211003102302-3230323233302133-3221102000111102-3223120130322211-2001232023212122-0233312112011310-1233230002321120"></a>

<a id="canonical-1200103312103332-0200010010332212-2300002322202312-3203231301103232-1330310303010002-3221303023011031-2132212211130120-2233221232000231"></a>

## namespace property — use_mtls_obj / 032223013103 / 5

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

<a id="canonical-3200110220221112-3132202102011101-3023302001331222-0320032203131012-2132103012300313-0101322313011012-1233011231000111-1333011302033102"></a>

<a id="canonical-3133222212100332-0231300032310121-1231122030233122-3223011230113032-3322122003333232-3300332333101330-2220230033322330-3020211111310212"></a>

## tenant property — use_mtls_obj / 032223013103 / 6

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

<a id="canonical-3000133100020002-0222022022320001-3230103202013202-3122112011000220-3113132101112333-1320220130110202-2123103302310010-3131012230310201"></a>

## Next pages — use_mtls_obj / 032223013103 / 7

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230130300220332-2220112133333010-2330133300102111-2303113300000313-2302211333123123-0223002301021111-1211003112223033-3120311222001311"></a>

## use_tls.use_server_verification — use_server_verification / 101101202131 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_server_verification

<a id="canonical-1202022113131322-3223333122031311-2012231003203110-2200120033103313-0303123113012001-1021131200222210-3302220021012302-2121023331223031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

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
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223100100003123-1232231022032302-3300010033210300-0213101020200111-1221002032320301-0303223221331132-0331022123113131-3112211221003130"></a>

## Direct properties — use_server_verification / 101101202131 / 3

- [trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211): complete subsection reference.

<a id="canonical-1220133023112011-0133103323321013-2230232111003000-1311033011331123-2002332230231132-0210313333022312-1133232021122222-0311333013311320"></a>

<a id="canonical-2022011333302021-2131323030233222-0202112110321320-0122233213301100-2212002112312132-1212011112323331-2133313321121101-1210201111311230"></a>

## trusted_ca_url property — use_server_verification / 101101202131 / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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

<a id="canonical-2212031200132122-0223002031011333-3003030102312113-2222320231120122-3103101032332321-2133212002203210-0230233123313032-3023103323002123"></a>

## Next pages — use_server_verification / 101101202131 / 5

- [use_tls.use_server_verification.trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310322322300221-1013011011301211-0131221202310311-2130033322231100-3233111220133121-1201333021223121-0210002100210021-3333031121020132"></a>

## use_tls.use_server_verification.trusted_ca — trusted_ca / 012200112010 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002)
- use_tls.use_server_verification.trusted_ca

<a id="canonical-3211120110011201-0121003012003331-3200121102030103-3100212002221031-3100111133302033-0212211223332131-0003023220323223-0303330120232302"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230303213202211-3310231021222022-3132223300113103-0211030113012310-2321102002112223-2212220323331023-3202003113131111-0321032200210332"></a>

## Direct properties — trusted_ca / 012200112010 / 3

<a id="canonical-2300000132123033-2103011001320010-0003201212202311-2131202021300121-0100301332031130-3100031203003111-3213000321133131-3220001111032321"></a>

<a id="canonical-1320222031003013-1333202031030121-0110103021102000-0112102233101131-1313202221310320-1212213301130112-2313202002233303-0111131101013213"></a>

## name property — trusted_ca / 012200112010 / 4

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

<a id="canonical-0121211112011311-3022002120201003-1200133210002120-3103310112300321-1331332330013221-1322110233131121-2221323113132132-0330103213300331"></a>

<a id="canonical-2021203133311122-2013022110033021-0102233210233201-0102033202221232-1201023312010202-2102322210030012-0330233011101222-0311311313200320"></a>

## namespace property — trusted_ca / 012200112010 / 5

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

<a id="canonical-3323101123113223-3112300031313201-0323233213330222-1122012021031001-1213311231311210-1321012323131120-1220113301313301-2232122310332010"></a>

<a id="canonical-0133223221333121-0321003021322232-3312233021013300-1321101103200301-1021101203012232-0110032023212130-1301033322323101-2232011010001220"></a>

## tenant property — trusted_ca / 012200112010 / 6

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

<a id="canonical-1312302022102310-0211021321100232-2233302031331211-2030123321111320-0030330233330222-0201100011031130-3020321301312303-3012333130202233"></a>

## Next pages — trusted_ca / 012200112010 / 7

- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220330212103200-2031223323032212-3130321030111310-3003013021302011-0310100301312313-3332003111030323-0123212020321121-3003313322321132"></a>

## use_tls.volterra_trusted_ca — volterra_trusted_ca / 101330121021 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.volterra_trusted_ca

<a id="canonical-1002302310301002-1302021231211331-0023032120010312-1013223211301323-1323121220130010-2023101130223113-3023221113032221-2012303103232033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
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
volterra_trusted_ca = {}
```

<a id="canonical-2211121322232200-2021123020200312-3030321030223002-1333313311300201-1333320332222333-0002210031312101-2113100210302131-1233123310113100"></a>

## Direct properties — volterra_trusted_ca / 101330121021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223000000323232-2111122331102121-2020232111131200-3012311131120323-1232230311023103-3200001221202302-2222010211231112-0233332033131122"></a>

## Next pages — volterra_trusted_ca / 101330121021 / 4

- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
