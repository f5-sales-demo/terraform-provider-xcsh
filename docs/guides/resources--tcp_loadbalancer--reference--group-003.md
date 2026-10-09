---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-0233132302001113-1013122023210222-1222022303323200-3120101033031121-0232202032202111-0031100302133121-1303233330230201-1301123310033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.endpoint_subsets

<a id="canonical-3120233110133312-0321210100022022-3000321113133310-3130101101021221-1223301332330133-1103223110230230-2230020003011021-0233130133022230"></a>

Type: `"object"`. single nested block, Optional.

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
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101012032110312-0012103013221021-2000123112202033-2211320112310201-0213110000103010-1321023222113211-3031331212323121-2320012220221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.pool` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.pool

<a id="canonical-2211110010101203-3132220223100012-1211011020310013-3133233312232003-2030302002310022-0201203300032312-2202322100111100-2111033302210000"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021131332322131-2221222301001023-3000301322233223-2133132013103211-0232232001232323-1330021113312320-0022313033230321-1100221302020202"></a>

### Direct properties for `origin_pools_weights.pool`

<a id="canonical-3231133211120113-0210010220020230-3032001333300132-3330221200210310-2003211213221020-2022321330203001-3201322003333331-2230210321333301"></a>

#### `origin_pools_weights.pool.name` property

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

<a id="canonical-1311301302321321-1103201112111120-2303010031013133-3031313320321301-2213220113003101-2320121311011330-2210330332320103-3323320323200323"></a>

<a id="canonical-2110021320201212-0212002011011230-2232322313032331-1010130330011031-3122222020202222-1301211001002100-2210033231001022-0330223132321031"></a>

#### `origin_pools_weights.pool.namespace` property

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

<a id="canonical-0210031323123022-3033220210002201-2122322311203000-2131001200333233-3020321333003202-2100011023013333-1312210001212013-2132103311100202"></a>

<a id="canonical-1011333222013332-1221111032103123-0312012231121111-1013130113201223-1320112333000122-1030003001321332-1101200230300130-3011002110200003"></a>

#### `origin_pools_weights.pool.tenant` property

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

<a id="canonical-2203330232231221-0231100222101101-2032200123103102-3022031012302100-0322123221223231-2330022031223301-3202022122302011-3102020113332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retract_cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- retract_cluster

<a id="canonical-0223212221221323-1003333011230323-2311103300002012-3022112031333333-0211120323003211-0323113322301032-0013033220133100-0321130210120033"></a>

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
retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302000233122001-0330320220303103-0201012112213303-1133102211110130-2333133012210201-3013203312103310-2212130213030102-1122113100303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_policies_from_namespace` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- service_policies_from_namespace

<a id="canonical-3001223103333333-1020031000231122-0311203031231103-3133312332031032-1302211202001202-3211013112120332-1003020031013122-3233111310220230"></a>

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
service_policies_from_namespace = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211312003133330-1223211300322220-2113300203322232-3010131033013022-2332011121303131-3303223200010132-3300002312210210-2000310102320322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- sni

<a id="canonical-2121331132131131-3033031100120302-0330132302022103-3000220102112000-3011311311311201-2200311121231223-3011221000201220-3201310112120322"></a>

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
sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220232031020123-0333221323000013-0330311311100200-2313132103001132-1300312202210020-1103313303133002-0011300213213202-1132012101023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tcp

<a id="canonical-3320220303000012-1212310221333003-2023232000001311-2003113122201032-0320213100330200-3200012310233331-1203312023032303-2033123000332201"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3320220303000012-1212310221333003-2023232000001311-2003113122201032-0320213100330200-3200012310233331-1203312023032303-2033123000332201)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3311310101231101-0312312311123203-2112200103233023-0231013101001013-1232132002021323-0011020322221211-1311301030203010-2212103030203323)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1021002212100302-1311013322023113-0033102003303321-1102223200321033-3111102121313232-0020131233022011-3121111200002222-0320203122133203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tcp = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223011333230122-2303021022201212-1030021321212322-1311000300020110-1130221333112233-3223223303001030-2231022021122211-0113200021221011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- timeouts

<a id="canonical-0231313301332031-0221012333210323-2202220031030022-3300010221320220-0102332011321200-0321222233312130-1133231323310012-1333333313112222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202331302322332-2231203121223220-2130301102301230-3012102112131032-0010201020122312-2133221103330322-2302313100102101-0232132211231012"></a>

### Direct properties for `timeouts`

<a id="canonical-1223233012111011-0320010120330323-2110100223101002-0021021131120322-3212003211111103-0020232031013220-3302212121020312-3332323131102132"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2301320011230031-3113330321132003-2132312201210201-3031201321333030-0312012330232203-0201010223302230-2233330222312123-2132112230212103"></a>

<a id="canonical-1322101011313330-0322103332331101-3133303203032311-3313221031321102-3313303003011100-1101220003130313-1133222103133231-2133131120322312"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2312030213132213-0021202101111112-2102221101011331-0332303033022122-0303123120222001-1001221110032001-0211033113032320-1010122102000231"></a>

<a id="canonical-3000221021220213-3002231100321313-3332331001031012-1010320322112133-1333011303321331-3302301122011001-2302012223330002-0333310303320121"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3323331302032012-3330110330212101-2303232133322130-0223003322013132-2302100202101301-2203312321331313-2011112230130313-1310032331303010"></a>

<a id="canonical-0112010013020330-2023111200021333-1010231002233201-1310301232330101-0112320111201112-3211112022201002-3301001311121310-2030302120222333"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tls_tcp

<a id="canonical-3311310101231101-0312312311123203-2112200103233023-0231013101001013-1232132002021323-0011020322221211-1311301030203010-2212103030203323"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
tls_tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001303022321131-3300200220200203-3300320202311011-1211021011330123-0330130000301033-3101200223121100-2223231111132210-0001102221033023"></a>

### Direct properties for `tls_tcp`

- [tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301): complete subsection reference.

- [tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132): complete subsection reference.

<a id="canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- tls_tcp.tls_cert_params

<a id="canonical-2003212231231230-2122303233200032-0123102210223221-2301333103012021-0022212032211012-2002001201000110-3021200222310310-2223313022132223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103133200032023-2301221133211102-2322230221131311-0122110202230130-0000102133012021-1312330001211121-0000010230131202-1331231001133123"></a>

### Direct properties for `tls_tcp.tls_cert_params`

- [certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-3213110003102033-1303112111202003-2023331320110320-3133112022030112-2203320300332012-0231011200113210-0123111101030111-0201203323200003): complete subsection reference.

- [no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-0011323331201223-2233033222323031-0123021332131023-2101130020120033-1322010103013022-0000022023200202-2333212121323232-1031132032111203): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313): complete subsection reference.

<a id="canonical-3213110003102033-1303112111202003-2023331320110320-3133112022030112-2203320300332012-0231011200113210-0123111101030111-0201203323200003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-1023331330101033-1020020103122030-2123332131230111-2230101121013212-1023301223011112-2011303033120331-2021112200211310-0113223113132333"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331330213020103-2100111223300331-1333023221332313-2213021030122121-1321003110112322-0221001131221223-0323030301021031-2201023212232213"></a>

### Direct properties for `tls_tcp.tls_cert_params.certificates`

<a id="canonical-0012113110101030-3122010132010113-1313312023003013-1023002333303203-2100120033120011-1230300102221222-0021013032200010-2031321030102233"></a>

#### `tls_tcp.tls_cert_params.certificates.name` property

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

<a id="canonical-3201012123201210-1112131301001222-1133201300123010-0213203332033303-3303030130330123-0320301212130210-1233331300132103-1313122220201313"></a>

<a id="canonical-2101032030213231-2122012220211001-1311132131211322-0220120302133033-0230122002212103-0333321123332101-3032003232002330-1313301332103330"></a>

#### `tls_tcp.tls_cert_params.certificates.namespace` property

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

<a id="canonical-3221111031212330-0122002102310230-2301001123303210-3333231002223120-3010123300203330-2320102201202030-2200331320021202-0331020202113010"></a>

<a id="canonical-0102122131133230-1332003033332212-1112213122312213-2223210312003233-3332320003100323-3012223232103311-3322132321300122-2323210101111303"></a>

#### `tls_tcp.tls_cert_params.certificates.tenant` property

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

<a id="canonical-0011323331201223-2233033222323031-0123021332131023-2101130020120033-1322010103013022-0000022023200202-2333212121323232-1031132032111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-0231032331233303-3221202323213112-2122302002313202-0322312221003231-0230200212111233-3222231300120331-0211303122102312-2001133332112332"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-0300232220332032-1221030130223021-3111102102101230-3303020222132332-0221110032000221-1320031310322313-2001122212000022-3332230202330021"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0113220113210132-0302310001112120-1300211000103121-0032020030313201-3131231212013112-2033232122111022-3312310021321330-0103113111131013"></a>

### Direct properties for `tls_tcp.tls_cert_params.tls_config`

- [custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3320220203210203-3010313213202302-0013120331113012-3331210220313000-0303222000322212-3000121230313001-2332320121112310-3112221030310331): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2211123331233312-2131002001130330-1033033000202002-1113120302302011-2303321232232213-0030213231213022-0021123310020203-0212203202100301): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3211203330313011-1301102230023230-1221213021020013-2103100111020231-1231100223233220-3102232113203303-0222110120112032-2133031112020113): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3331020220302000-0002023130121130-3302231310301101-0330001232103201-2220022300103220-2221032222003310-3222021022322002-2312232130120111): complete subsection reference.

<a id="canonical-3320220203210203-3010313213202302-0013120331113012-3331210220313000-0303222000322212-3000121230313001-2332320121112310-3112221030310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-2113323331002013-1311022023212033-3312022212313032-0122213021312121-1303211030101120-1330002303320230-0032333120113002-2323312212101221"></a>

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

<a id="canonical-2031133022013120-2110202131030002-2321222310223123-0012013202101000-0011022233001120-3312310021101000-3313213013012233-0203001312212010"></a>

### Direct properties for `tls_tcp.tls_cert_params.tls_config.custom_security`

<a id="canonical-3121112032303131-2113333330102113-2302002222011123-1321012321213122-2301100112203022-1302000331302101-0021302310222301-1023212320320132"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-2213033221303032-2133232233212220-0010120110122031-1131332230001231-1032003012030000-1201202020100303-2102311131322232-1312010102200300"></a>

<a id="canonical-2132221132233111-1312203223330201-1133132001113320-0032112303222123-3120321133131213-1322212121102012-1100103302123312-3000123013002211"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-1121223320123033-2012032320221312-1303211110000233-3332230021321030-2202030111211230-0313330010310102-0221133303120323-0102111222030332"></a>

<a id="canonical-2122030302022121-1233200200222221-0212312103222233-0200313210201201-1232101022130223-3203131213100303-1112121332213302-1231210111033101"></a>

#### `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-2211123331233312-2131002001130330-1033033000202002-1113120302302011-2303321232232213-0030213231213022-0021123310020203-0212203202100301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-1312101011311300-2121123211331302-0322000312232030-1321002111111312-2211110032231032-0030213020110012-2022321120002130-3313030002300012"></a>

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

<a id="canonical-3211203330313011-1301102230023230-1221213021020013-2103100111020231-1231100223233220-3102232113203303-0222110120112032-2133031112020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-1330110313322011-0310203320203031-3121310203122013-0101231031333102-3201221010321121-2311210210011003-2000223310310022-3131021220323320"></a>

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

<a id="canonical-3331020220302000-0002023130121130-3302231310301101-0330001232103201-2220022300103220-2221032222003310-3222021022322002-2312232130120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-2212022333012322-1311013322202201-0332230112313031-0020113110133021-3200103330203022-1031311301012301-2010030212133332-3100012033212000"></a>

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

<a id="canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-3111121020110031-3020010003330322-1122023201122301-3322201021112202-1223113202201010-1103102120111200-3022123010301001-1213101302223032"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222203122203030-0001202123331220-1021310031023301-2312013001132302-0313021113031211-0233010130200100-2202231112200131-1131133303323112"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls`

<a id="canonical-3013320211222232-1113312010311002-2013331322303032-3122101112232213-2121333130330020-2313232131331221-1231201031202110-2000330200313310"></a>

#### `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0013203232221122-3131111332200030-1203322032032201-3020301221113322-2031221333332022-1223321321032033-1230301223121200-3020112202333022): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-3112230111023311-3101102123323312-0012002232001030-2310033001231110-3230312221301102-2001013123020332-1113120033111310-0200211220101232): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-0200003100030200-2110123332302230-3021233032202021-2113003301313222-1010010331203122-1331032301012223-3233003001322033-3201021231230232): complete subsection reference.

<a id="canonical-1331110031120010-0133112021222233-3310123203101103-1120030233300330-2102123221100213-2032103221311323-3021201201200321-3002003101100213"></a>

<a id="canonical-0012300033110121-3323302102332111-2323310133022122-1223321332103231-1330132331131323-0112020221211021-3231210130001033-1303100333021001"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-0232003220003202-2001302003310012-3101300112231123-2000131303013013-0230313130011112-2201230112233231-2311110123132210-0221213112031311): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-2211233203111001-0112213331212320-2030100120211101-1123203223220031-0213311102100010-1322132321202230-1210010301011010-3231020331123320): complete subsection reference.

<a id="canonical-0013203232221122-3131111332200030-1203322032032201-3020301221113322-2031221333332022-1223321321032033-1230301223121200-3020112202333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-0201000232020212-2202311101120211-1303001110313122-1323113011333302-0023300000331133-2101223132022230-2001321002132032-1111132022101323"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010120302302122-2223322333122133-3101023312100320-1030032131122132-0023201011002111-1011132033012000-0023300032332303-0000020110210301"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.crl`

<a id="canonical-0122220321213333-3032112020213112-0000131113302311-0130313203103131-0031312022112301-2313202322130323-2111333000333032-2133223233320111"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-3222211200132001-3031220310033032-1310310032033320-3300220203220201-1103321122010233-1023001202031002-2032003321031202-2223222313132122"></a>

<a id="canonical-1111132213230213-0311231310131323-1312011221013133-3301002312302010-1202120312021300-3320121112000031-3102032021311331-1320333101010320"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-3202100131322103-0100202301112031-2233311310302301-3323223101332211-3020310300300033-2211113331313103-2311312300112033-3302303033300002"></a>

<a id="canonical-1311220130102012-1302121203021030-3320232230110322-1121300311210323-0001021000302230-3211122000120103-0303330321012201-0033222011013022"></a>

#### `tls_tcp.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-3112230111023311-3101102123323312-0012002232001030-2310033001231110-3230312221301102-2001013123020332-1113120033111310-0200211220101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-3321012202302322-2300022233010223-0002002320103100-2201212300331212-2131232302233030-2311321311113112-1100223113313200-2203112310201333"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200003100030200-2110123332302230-3021233032202021-2113003301313222-1010010331203122-1331032301012223-3233003001322033-3201021231230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-2001121202320001-2202231211201110-2003012013201230-0220313302201133-2330001003222200-1012133301012132-2332233323330302-1302031300000212"></a>

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

<a id="canonical-2331222022232120-3013210101003101-2211013213231222-3033311131113121-3031033121311113-3011023033120033-0203232103332100-2220221301033222"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-0030003332002301-3033213103121332-2202100321303020-3032001133100103-1320313301321003-1103233113131113-2120211020010311-0331221112301033"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-3033303232333313-3122210031010110-1102013100233333-3221003022213232-1133233010020222-1101122133102123-0320200000011021-1313122332030222"></a>

<a id="canonical-0021133212211202-3010133232303302-2031000130220031-1123212301232131-3301023310231100-1121111013103011-3111032213023103-0020233203222313"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0133131301311303-0200121022110032-1313302030123033-2000333320301323-2023030002320232-1003302101033332-2112302300111012-0000133021222213"></a>

<a id="canonical-3233232312301212-3213031302233020-2322232221330132-1132203112322023-2033113112131130-1030333333112123-2223313000011103-3331113222021003"></a>

#### `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-0232003220003202-2001302003310012-3101300112231123-2000131303013013-0230313130011112-2201230112233231-2311110123132210-0221213112031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1102003200120212-0001032231120020-3010012232313310-2102131301303203-2110213321321002-0103100013232032-1012012330132120-3012111012230230"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211233203111001-0112213331212320-2030100120211101-1123203223220031-0213311102100010-1322132321202230-1210010301011010-3231020331123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-003.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2001133311013203-0331320222002031-1020230000001303-1230320001331322-0033131323332332-2110301023013303-0333023120112033-3333330002310022"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101303300332021-1112133310330111-0022121333233231-1220320110330022-2320131232322321-2022121000011001-3223322311110012-3221201320001101"></a>

### Direct properties for `tls_tcp.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-3110021330311023-2330101103023010-2310200210001001-3330102302120121-1231201313101111-0212112322133232-3112023211112133-0223203312230311"></a>

#### `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- tls_tcp.tls_parameters

<a id="canonical-3110131221302012-0020201332212302-2232113113030133-1301112203103333-1123112003212201-2021231011313230-3030310112330201-3023113233130131"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130213001110132-2323010012310223-3030020120322112-2322202003123101-3301322303001133-2122003131322222-2003033221312111-3303311311003032"></a>

### Direct properties for `tls_tcp.tls_parameters`

- [no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312032102310132-3313333033231033-2033013201212031-0312003323222322-1333102130012332-3313300133122000-1013213111230231-1333131222103221): complete subsection reference.

- [tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000): complete subsection reference.

<a id="canonical-3312032102310132-3313333033231033-2033013201212031-0312003323222322-1333102130012332-3313300133122000-1013213111230231-1333131222103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-1101321203110110-2003011002113030-3132001310310233-1003003302213320-0301113232132202-1101023221222301-3222011302000021-2013123121323323"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-1202231212120100-0230313232012210-0101220333230100-3133010011111033-3100212131003130-2032111121032021-3213120013320200-1223223221323012"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
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

<a id="canonical-1131132211231033-3321233100121200-1012010322331223-3320021132310323-1121320000100111-3231202201021203-1021300112311300-2302023133111130"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates`

- [blindfold](resources--tcp_loadbalancer--reference--group-003.md#canonical-0131210121222222-0221023212310113-3310003331313320-3133010103122132-0111131221312000-1321111232110022-1201313003103332-1201311011302332): complete subsection reference.

<a id="canonical-3323312300322303-0101303101133320-0113031230010102-0112102212033023-1100213323020213-0233230302230231-2031210032231230-3001200130112212"></a>

<a id="canonical-0100023133333322-2113003030103310-1120233210023100-3012231123010123-2100312213313321-2030302320201203-2332101110011303-1033131231302111"></a>

#### `tls_tcp.tls_parameters.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-003.md#canonical-3301303122202032-0100230000013222-3021011212311010-0230130121121012-1321031020230022-2221312323210312-2333010103101231-0313112121312113): complete subsection reference.

<a id="canonical-2213300101033130-3302203302213113-0322332331210012-0201321200032210-3020333310031000-0133011302021313-1233002210320112-3113323033122330"></a>

<a id="canonical-2221302132200103-1303000122031332-2223131333302131-1103120112033330-1030231131210011-3011211003331200-3103100301322230-0333330211210302"></a>

#### `tls_tcp.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-003.md#canonical-1130133232111030-1313232203303110-3131200030333223-0333033223000000-0222210123323020-0202130300123102-1122110121333233-1003022103200212): complete subsection reference.

- [private_key](resources--tcp_loadbalancer--reference--group-003.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013): complete subsection reference.

- [use_system_defaults](resources--tcp_loadbalancer--reference--group-003.md#canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132): complete subsection reference.

<a id="canonical-0131210121222222-0221023212310113-3310003331313320-3133010103122132-0111131221312000-1321111232110022-1201313003103332-1201311011302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.blindfold

<a id="canonical-2000130330312033-3220121003000323-2122003300123031-3230032212233233-3333132000233313-3300203231103023-3323320201233021-2003101130202300"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3100222321003203-1212220030232311-3131102112321032-0023300220013102-0313300302221103-3133312121100302-0333303220012332-1320011303113313"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.blindfold`

<a id="canonical-0111212221020332-3033021322321210-1110100023031112-2132203130321321-3320030133213012-0233223330020320-1100211323331000-1130131132000222"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-1333321012323031-1211122332003213-1112303221012203-0102230211010020-1322233013022303-0322201101131110-0020232102130110-3300311323303110"></a>

<a id="canonical-2303211103020200-0013122103133131-1032221111132313-2112111020130320-0022123203312023-1111032022220012-3101311003302033-1210220032013110"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-0133231131302231-3121201331013101-0101221033010230-3221031021233220-2220011120331312-0100220023012000-1331033200201333-3202220010103220"></a>

<a id="canonical-1321111221111101-2210100011130222-3213222221321022-3322021322013300-2311131012212120-2201103223100103-1001230220023123-3030310323202130"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0111000303120012-2231103010303232-1200132130320222-2021013333233111-1320033020030222-1100021132331333-3331013310030232-3300212313220133"></a>

<a id="canonical-1330012003333230-0032100120330021-1000233001302033-2102121002000023-0020213112101023-3212122030230321-0310003002312031-0033233003232221"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-3021302021330110-2020311323003203-0233210302020032-3112303310001312-3103323220230023-0033033133023000-2223020301320212-3331132211303011"></a>

<a id="canonical-1311013233322102-2022112231110032-3033113132031233-0032030202033333-2320033301310232-1131331323010311-2321101312011210-2132212111222330"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-1231311212312221-3112221122203032-1110323313111111-1011233330332220-2102202212223332-1032331333302313-1331031010030021-2111133222001022"></a>

<a id="canonical-0012103311020100-1010123102233023-1313021300323102-3111112212020230-0033121131132213-3131132210300302-1302220021120203-1132221332322003"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0310030100212201-1213033101230232-1312033201132122-3123311023101122-0323131032103331-0313122101220031-3231210023222021-2111013032030003"></a>

<a id="canonical-1233031103113221-0203132112010122-3013021203212130-0132301131020101-3210320113333300-3023101332010121-3130330023132123-3202233202231332"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-0031121221213132-3333022303102202-2110333120312101-1033301011301330-2131002112301213-0203233001301330-0312121013312310-2310200320222212"></a>

<a id="canonical-0113112332100301-3000000010300300-2122033201101230-3132113020121001-0311122333010321-0013321112310131-3120213032133131-2112230312233123"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-0021000231032301-0101013012332001-2313130122233100-1122332013101010-1213101103100101-1032201113021320-0312221123002110-0002331022110103"></a>

<a id="canonical-3103320201213100-2301101110320312-2123332103013321-2103022011033003-1332133100313231-2310111320233121-1330031232211000-1211202132011033"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-2302311230031133-2010020310113020-1321333113112003-2003303320333202-3321123012021130-3013320331023120-3323232313123221-2132020003311131"></a>

<a id="canonical-1132112313033002-2021113221320032-2000013213102313-2023100023201331-2021100303211002-3112133131001131-1323131220310331-1320130301030231"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-1313110130322302-2113113132032132-3220031220023113-0331332223000200-0221301033201211-0311031130321311-3001103330131103-2200311201222332"></a>

<a id="canonical-1032210011323222-2232122123231220-2033202113111210-2223320102313201-3223321210210020-1300132131001023-1030303331000023-1010001332113223"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3103122130332200-2113310132210121-0111310120013213-0321130020321233-0021030103123033-2203110200212213-3202022123221113-2021010333003323"></a>

<a id="canonical-2210002323010313-3110012121301323-0301003312232202-2012133300003110-2133103312030002-3203230230123013-1331130022023130-1030232123120122"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1103013322302033-0310023012002103-1331111313130220-0310013323301333-0101213321113120-1211023033032030-1222200112132231-2120112123023010"></a>

<a id="canonical-1033331223230112-1120012112132023-0222131113331131-3013233113120001-2010201101211230-2320123111120030-2210000200303003-1001222113101221"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-0322111033102113-2031332323310130-3310103112202111-2321232022102330-3102112010123312-3000022010311122-0002131012021230-1330223020123020"></a>

<a id="canonical-2312020321112201-1220033112211111-2000033000100312-3010330131223202-2322210021000313-2112103222233301-2131312130023320-0200220301031013"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2123333002011103-0120130202032002-1031120103021211-3313132213202313-1100201032212211-2211321203300022-1303223302230300-3330101203100313"></a>

<a id="canonical-1010123111131321-0332300113220001-3303233022031301-2133311003131232-0032202000321033-1100123333232331-2213122330110222-0231033013103212"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-0222332220013201-2312301313213312-3100121112031220-0103323200232133-1231213213013102-1232302130331120-3133333230213322-2032022002012201"></a>

<a id="canonical-2222332002222310-2331023223301013-0213233323020333-1131132332103301-1302213212101120-1132110021233331-2221313310101101-3031103301020300"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-2320110223102101-1313213310213202-1213220211332021-2031301121221321-1113230311332013-1201102122133323-1221123202120132-2001011021121021"></a>

<a id="canonical-3331212311030312-1022111133010321-2223101120110320-3233301033220130-1023332000120212-1320320312112211-1212003222011221-1213200111322200"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-2031012221223203-0313100313000122-1223101121323330-0131002220322332-3331003111101202-1213013300202312-2023311130031101-3111130130030300"></a>

<a id="canonical-3023030310302210-0302110133200003-0133200031200100-0010102232112111-3110101311233113-2310223202221202-1313223022222123-0221020233110130"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1332330323332233-3232203122120321-2101200001221201-0302210322122300-3110220313200003-1100203233001102-1003023313131212-3221013003013333"></a>

<a id="canonical-2220312112220301-3122330300202120-3010232302233200-1300220300303120-0313310231031312-0002010303321331-1131230233133200-2231023320311002"></a>

#### `tls_tcp.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-3301303122202032-0100230000013222-3021011212311010-0230130121121012-1321031020230022-2221312323210312-2333010103101231-0313112121312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3023032010200202-0123321213302310-0200011320123020-2301033322202011-3221011013213010-3000231003323200-3101222231311010-0100003033232233"></a>

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

<a id="canonical-1230231131031123-3232211302102332-3123130022010301-3210021102212112-3230133130133223-2310321331012201-2202033022023113-2212322223021201"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-2212220120123321-2133031300330331-0111223332223332-3231033131323101-2102203211122021-1000202233110013-3331303113201222-0100103310000211"></a>

#### `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-1130133232111030-1313232203303110-3131200030333223-0333033223000000-0222210123323020-0202130300123102-1122110121333233-1003022103200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3023311011102113-3021223201133112-1323211121000232-0303210323200303-2330130013033010-3002032023100001-0212323313231232-2121232130122333"></a>

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

<a id="canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-1123011310203113-3021011221132332-1033103301311113-1000100031023220-1333112013232012-0000112221132213-3322033303103231-2223231110110032"></a>

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

<a id="canonical-2122113230102120-0223231100210221-2130230031011022-3121001232021312-3002230132013033-2100333021321302-3203312001221301-2030110230111110"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--tcp_loadbalancer--reference--group-003.md#canonical-0320313320210230-1002123322100231-1303211120310011-1001023000201010-0110021301010220-1120210310223301-3320002323332001-0323323232130210): complete subsection reference.

- [clear_secret_info](resources--tcp_loadbalancer--reference--group-003.md#canonical-0203120320001120-1303310322101331-1333232131022333-1101032230333221-2201100331122131-0311223122332330-0203003300131013-2010200203233322): complete subsection reference.

<a id="canonical-0320313320210230-1002123322100231-1303211120310011-1001023000201010-0110021301010220-1120210310223301-3320002323332001-0323323232130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-003.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3332301012313030-0020232102210020-3000311111333222-2133112010222120-3222012301032002-0310232101232002-3323230131122013-2331232301032132"></a>

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

<a id="canonical-3203233022131121-3213000111221331-2130010023032333-0332103331321231-0032000002210212-0222331330200330-3030121103200333-1033313230331300"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2012210230101302-0132203223022232-1220023022312111-2032112132230330-2301222033320111-0330020313110321-0200003003200313-0011213000113331"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3303301233102000-2312323022220232-0223230133231013-2101201011133311-3113222303012022-0013323103231121-3121230313121131-3011300303220132"></a>

<a id="canonical-3300012230213133-1302130123321202-3013122231110333-3330232313020111-3132212003222131-0002333101003000-3100302212233131-3120023003012000"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-1023331223321223-0021203233231033-3110103223203020-0303221210102323-1232203131220130-2113000202030023-1020013013113100-2031022321022321"></a>

<a id="canonical-1330012120023103-1100302113200232-2211320332313320-3133321320332102-2030321211303223-3210223221321033-3003100332120122-0211111032332321"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0203120320001120-1303310322101331-1333232131022333-1101032230333221-2201100331122131-0311223122332330-0203003300131013-2010200203233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-003.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1121332121001303-1330031030003332-2332023200211220-1113111101212011-0113332002311202-2312203002010102-3000023000101023-1120111233020003"></a>

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

<a id="canonical-3113222200321323-0012112222113022-2101111303012010-2231220100333211-2113030321021020-1320201031203022-2112311023332332-2303212100322233"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0102311301232332-0233123210030330-2203000022130033-2032202210202330-0331102020232103-3313133111031201-0203030310332121-3210013321223232"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0130202101332222-3202202301103120-0212010100131101-1120020103112031-3033201300100321-2003300301101320-1102130200313011-3131223233102310"></a>

<a id="canonical-3310332011330120-1113301102011202-0020322132121330-3332102313213231-0031131011212121-2001320332332003-3333310103102330-1331203233232321"></a>

#### `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-003.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3321332112202331-0332220020203310-0303033222211031-3111032001010021-3022330202330210-0230302202001231-1203010221201321-3320220322313121"></a>

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

<a id="canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-0212300030220000-1030102030120032-3230012030010030-0000021233232103-0220012112012103-0232213020123322-2232021021001302-0012002033001322"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3001323033032222-2331303031031202-1010002210033302-0101030011033001-1102121331203212-0132020010132202-1210103112300002-3122221122033302"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_config`

- [custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-0130120303200203-2030112320030203-0133122321202013-1233130133222211-2322200122103223-0120023212110033-0231012031200131-1233012301231213): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3011313102103121-3231233210201221-0022213002310313-1032322300233321-0200302013232133-3300313101210033-2330210102022303-2303013013130210): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2300003002320102-2022222323333210-2301101111110221-0202120131321310-1011003021132121-1012223013100001-0302333020233012-0111332122032303): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-0122101212100122-0123222332303022-3031313101311112-3132323113331202-1021023112312302-0003111030202031-1320222133300130-2222121221213212): complete subsection reference.

<a id="canonical-0130120303200203-2030112320030203-0133122321202013-1233130133222211-2322200122103223-0120023212110033-0231012031200131-1233012301231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-2101203230103003-0022303032202231-0131223223111012-3233012312103213-1001023021033001-2100022031011221-1203102001323321-1301100001220121"></a>

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

<a id="canonical-0210121110010032-3212123100020220-1310213330303222-3011213020013123-1221330302232031-0210001032022131-2022310131023022-0231332320221312"></a>

### Direct properties for `tls_tcp.tls_parameters.tls_config.custom_security`

<a id="canonical-2120111230032212-3301001021200020-1023133233032212-1110303201110011-2021022230221201-1131020200030133-2223232121220313-2103311220122113"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-1031333132300100-0133123202123002-0000101321001333-2031122131003303-2022221001220200-0233300233000321-2332323322132110-1313021210333320"></a>

<a id="canonical-2211100213231113-3300203112120113-3321300012302120-0013332123302100-1210210222210320-2103202033030233-1313312333103020-2212300003123103"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-0332013010102312-0233000311211020-2331000222113201-0120313002200032-1213122300031232-1003222000210023-3122233313232021-0310303231010021"></a>

<a id="canonical-2320112313111021-1221022220123302-1020213133120032-1311333321003101-0201232332133231-0320201113221233-2001002222320021-1002103211233230"></a>

#### `tls_tcp.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-3011313102103121-3231233210201221-0022213002310313-1032322300233321-0200302013232133-3300313101210033-2330210102022303-2303013013130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-2100100011331032-1320201311013202-1220313213333021-0132130000330331-3020102312322220-2201223031321111-3113031332123203-3321301231210332"></a>

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

<a id="canonical-2300003002320102-2022222323333210-2301101111110221-0202120131321310-1011003021132121-1012223013100001-0302333020233012-0111332122032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-1221332221122311-3130011130300221-1202023201232230-3013132212022033-2022330030303302-2230311211331020-3013323230032002-2121232033311301"></a>

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

<a id="canonical-0122101212100122-0123222332303022-3031313101311112-3132323113331202-1021023112312302-0003111030202031-1320222133300130-2222121221213212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-1020031022132103-0201111302022203-3133231013000123-1122132022031103-2230013201200003-1031232100303201-1001212120302120-0321332000232002"></a>

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

<a id="canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-1232221302231213-0031210310023113-0211110311201223-1120002312213120-2011321300110012-3200030013001222-2100103323222132-3302122312012030"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320001103232103-1132122010310113-1311010303001311-3013002220132320-1221202031120131-1331011230322200-3000200010023002-0130022303312300"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls`

<a id="canonical-3221111013201111-3213023210020303-3131030321232001-3203102023003103-0102001010302210-2012103231122331-1203201220231100-3221211132120123"></a>

#### `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0221011232132102-3013223103032110-1212203323321123-1200223110001020-2321033103122302-2120320330031323-2303023100221123-2200033122213033): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0123330020221011-0130332120131212-0123313133232112-3111230010033011-2221002101021100-0200001231202201-1212210333330011-0011131110130111): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-2011233303322103-3011101331001312-3210102103312000-2200300102303003-0101022230330103-1013222332313000-2002201113122021-1203021300001330): complete subsection reference.

<a id="canonical-2003211222130021-0102013022031220-0311111202133301-1223000212300330-0120033121121310-2102321021313102-1300102312222321-2321210202130323"></a>

<a id="canonical-0300203002021301-0130032201210120-2100121300201022-3100221032103212-1020200321310231-3013112032032102-0012002300102131-1122000121121303"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3221120302133200-2013113331332203-0323332322332331-2113030121313132-3201101323233220-1002123231331100-3030233310332120-3033231201103012): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-0000312003001302-1122232200100212-2311230011331001-2201013210031300-0230220320310321-3013203311333033-3322330201331111-2213231322023111): complete subsection reference.

<a id="canonical-0221011232132102-3013223103032110-1212203323321123-1200223110001020-2321033103122302-2120320330031323-2303023100221123-2200033122213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-2231002300022321-0310320020013321-3013231001302330-0222022201020003-1333312021222020-0123002203132031-3221023102230103-2323210312230300"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031212221223211-3022121223331103-0212203303202220-3232013103202312-2022000123230122-0021023010332023-1113320201002322-0220033022121021"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.crl`

<a id="canonical-0230012313210303-0033010223133013-0300200112331132-0321122221223033-3132301012030211-3302203013101313-3231320322233230-3000032003213223"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-0113221301022030-1221222321132301-0010123121103202-0300021231303212-0212320221200121-1223111003120310-3321311201112000-2313202122122223"></a>

<a id="canonical-0101103223011323-2011231101203212-0003011320031323-0033201130313332-2131110102203133-3123030020013130-0133100102130032-3203130323133013"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-3321103023302322-1130302103302300-3232302121212231-3213320201031033-3232213100110303-2101303022312333-2103112303122232-1220320212122000"></a>

<a id="canonical-3003213021111132-1001321322101300-3101020123012231-2203202323131113-3321100222121231-2101102002000113-2322331132313321-1232012031130113"></a>

#### `tls_tcp.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-0123330020221011-0130332120131212-0123313133232112-3111230010033011-2221002101021100-0200001231202201-1212210333330011-0011131110130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-3130121333221110-2303030212130233-0101213010223232-3212112032320032-3310033333332110-0330323323320023-1122000222002203-2230112223222001"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011233303322103-3011101331001312-3210102103312000-2200300102303003-0101022230330103-1013222332313000-2002201113122021-1203021300001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0223202131300212-0310312222130113-3132132013321133-1113010120111302-1112023223102000-3113203211000312-1300033232131000-2303101302210311"></a>

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

<a id="canonical-0212210231001311-0200231202130201-0130323031010200-0232033231003320-0230203302111310-3003322123321213-1300020003303020-0310231221131021"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-2212212023121202-2313023330321321-0231321121110211-3123301023001322-3130220200231021-3212313322310221-3313202210101101-1232031212223231"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-2230130301232230-3001011121023102-2302011103023023-2003121212313023-1110102030102030-2003031323222233-0120002222012302-1220033330133113"></a>

<a id="canonical-2221301000131000-0310010010313313-1103002201211110-1030322101221103-0331103103033121-2230122113030212-1021102310303012-0312011020020110"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0130123031211303-0212032231230031-3231110103101211-0001310013220021-3032222201300001-2001313132300123-2200003031320000-3221223311320333"></a>

<a id="canonical-0322022222301102-0031100211113212-2103313120302220-1012212002131133-1113130112021300-0330130201131020-1011203022030213-3131223100233231"></a>

#### `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3221120302133200-2013113331332203-0323332322332331-2113030121313132-3201101323233220-1002123231331100-3030233310332120-3033231201103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1301301113111110-2233321110301200-3201113132333301-3123023200120333-2130030330321230-3130101213131301-3320033210001100-2302023131330223"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000312003001302-1122232200100212-2311230011331001-2201013210031300-0230220320310321-3013203311333033-3322330201331111-2213231322023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-003.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-003.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1232011003223131-1231312010203002-0320300222000333-2020333200020212-0013010110113301-0321220123313213-3331330222001311-0121200312220102"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303331130132023-0322103302211320-0103003011133022-3211031233221103-1232130203330322-3122232123213030-3312013233102202-0120212031100112"></a>

### Direct properties for `tls_tcp.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-1231010103323111-0133022222121131-3031321113032032-2113331320323301-0010101132333022-1312231130002100-1331200300323310-3001332033133021"></a>

#### `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tls_tcp_auto_cert

<a id="canonical-1021002212100302-1311013322023113-0033102003303321-1102223200321033-3111102121313232-0020131233022011-3121111200002222-0320203122133203"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_tcp_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301332223213023-2323011010011022-2133100013322011-1312230132311011-3212030120332121-0110022011222231-1300013003302110-3033222102321220"></a>

### Direct properties for `tls_tcp_auto_cert`

- [no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330): complete subsection reference.

<a id="canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-2313120012003131-0311100200113110-0123033033123200-3003200132300220-1021000011003221-3200000101330201-2110012131301202-3132133220230013"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.tls_config

<a id="canonical-3321310201213312-1102103322302112-2213212202200032-3302231101100330-0213221333111211-1023311110133123-0020313120313112-1300130202230102"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2322031011113223-1021330102311222-2013112020130210-3033133323321302-3000230321102233-1232103112233223-2321330112333111-0233032210322131"></a>

### Direct properties for `tls_tcp_auto_cert.tls_config`

- [custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2313013020003011-3223231210320332-2110311102211120-3033132323001101-0011210322131033-0132220020103033-0003120020233123-2220023113001312): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2201112121202132-0000011113332030-3232020202200320-2131120132123300-0331302111333333-2313220001121022-3111212221132303-1112311301131322): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3000323120232101-0221013010300110-2213311012323201-0213132003100230-2121333300332311-1101023310121302-3121012021212101-3330311103103320): complete subsection reference.

<a id="canonical-2313013020003011-3223231210320332-2110311102211120-3033132323001101-0011210322131033-0132220020103033-0003120020233123-2220023113001312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-3122320210213120-3132310300203321-0023203133132221-2012132122221032-1132021010222230-3110333322211003-1200110210233200-3202113110210333"></a>

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

<a id="canonical-2030232132203002-0001202221202120-1300000010223123-2311112301121220-1300330323022332-1232213030231230-0302030301303213-0331202333222320"></a>

### Direct properties for `tls_tcp_auto_cert.tls_config.custom_security`

<a id="canonical-3110230210310221-1011323300032322-0113000130312010-2022231123113023-1022302023122221-3111300133033221-3101321133322100-3233221210211200"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-3302230203202020-1020012130110030-2020113300330120-3012211330212032-2212103232133231-1230220210132333-3301023232002323-1023002132323012"></a>

<a id="canonical-1200111332222101-0133033130231233-3000312010123303-3102022103131022-1201302121312102-3222203003032301-2010233102021002-1013133131022013"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-2331012331102330-2222311331122201-3102103312030101-0320011333220032-0103132130020020-1210103003231030-0300130312113113-3122112322122300"></a>

<a id="canonical-0113001330012013-1301121111210002-3013130131310121-2230210222222310-1330032311200131-1010122210232001-3301023021133233-2230300322221110"></a>

#### `tls_tcp_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-2010010302012332-2230323211012202-0312230201203102-3300001130302322-1330311131021002-0113302200111113-2122213232321110-2211120031233012"></a>

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

<a id="canonical-2201112121202132-0000011113332030-3232020202200320-2131120132123300-0331302111333333-2313220001121022-3111212221132303-1112311301131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-0012331301303113-0122013031120210-3111133330002022-1223232130010220-0213102121003312-3333303311320231-2330320030021323-1132012003211232"></a>

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

<a id="canonical-3000323120232101-0221013010300110-2213311012323201-0213132003100230-2121333300332311-1101023310121302-3121012021212101-3330311103103320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-2203331123101331-3200220000230021-0112033122002013-0333232221023131-1003012123322012-1001322221220201-2000000113100103-2111033323100312"></a>

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

<a id="canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.use_mtls

<a id="canonical-2122123310312001-1332302032133120-0112220022001113-0332021122030022-0331230110000112-3330102331101112-1023032300031022-0133033333131032"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011033001030203-0123300321112112-2203333031332221-2002331230032310-2002131100101113-1330103113100010-1001313310222210-2230333032213002"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls`

<a id="canonical-1021332202221031-0011322033110003-0000322201110132-2312203201221221-1310102321101111-2323221212331321-3000212001013023-3303203103330332"></a>

#### `tls_tcp_auto_cert.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-1211131111020302-1000023331233231-2300202002002211-2312321110020022-2222322300110003-2030101231311323-2120313123031031-3123103313130203): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0221331323323322-2233111213111121-2103120211102130-2112133132101331-2100213203012113-1213303221011011-2112132230030001-2022231220113302): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-1313322021032031-2133113312113013-0313333312300101-0000313030010021-2101100033003031-2103231100121022-1311023121103301-0312120231221130): complete subsection reference.

<a id="canonical-1301233321023301-3221301123322013-3120111200100111-3213131121132220-1322131002130010-1021133020131331-2222210230130311-1313302230033120"></a>

<a id="canonical-1020221220303300-3321131202233221-1313021322320002-0230322013013130-0322131302303101-2323321323312033-2323202232102300-1113310313011012"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3332232033013000-2032222110022033-0001300033312133-0211031021233110-0110331002111312-1311101022331130-3223133323323212-3031212323210310): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300): complete subsection reference.

<a id="canonical-1211131111020302-1000023331233231-2300202002002211-2312321110020022-2222322300110003-2030101231311323-2120313123031031-3123103313130203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-1020103222230313-0322312213133322-0121301323231031-0302030102111222-0300002203120120-0010103310112202-3130101220001322-3000220213230313"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022030212303100-2333232200111031-2131311232023312-0301021330031221-3030010133111021-0023323122020221-2102310222033120-1102131213301300"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.crl`

<a id="canonical-2022111101313010-0313322233121013-1331001202300120-1223110330130112-1030131213103012-2132110103021201-3013121013320201-0010010332220303"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-2123213131033113-0101303223000211-1112011200112121-3110332222321100-3213022001232231-1300122121202230-0012311230210033-3333321322223111"></a>

<a id="canonical-1013320201303231-1232322313001231-0021101100113011-3001200012100030-1220121212221112-1333223121031121-1010312022302112-2211013103131110"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-3023300012323011-3112212113020112-3203202020332121-3202100302232031-3000323012331313-1211002031231033-1323023311223000-3001221322120300"></a>

<a id="canonical-1000000121011311-0002133320130131-3021031132122332-2331322113220001-1303213112131320-1112312220222333-2133033223223320-3223031122031121"></a>

#### `tls_tcp_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-0221331323323322-2233111213111121-2103120211102130-2112133132101331-2100213203012113-1213303221011011-2112132230030001-2022231220113302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-1331203211100310-2320012102132103-0030010000223320-3300333111012011-2010310131300000-1010323123033112-1010210100203101-2310320223223302"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313322021032031-2133113312113013-0313333312300101-0000313030010021-2101100033003031-2103231100121022-1311023121103301-0312120231221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-0310300202211220-2130103020100302-0102210122321330-2030310221102210-3211111311311132-3303103211003233-1323202302101323-1103012012131011"></a>

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

<a id="canonical-0331230310010010-1022010332200010-3233301010013112-3230123031313013-1200032013021011-3001100330221212-0131103322110102-1320131123110133"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.trusted_ca`

<a id="canonical-0101112133310323-1100032200311133-2011111120122012-0120033102303020-1103302330232031-0323101331133123-0210322302301331-0113113013302222"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-3100333102033203-1230122020112301-3010023213303000-1120310003030230-3131012112101221-3333023301131013-0101220110332130-0130312003203320"></a>

<a id="canonical-0122010000201231-2123022320210111-2201331213103123-0321002231211102-1330221211122201-2332202023321302-2112103213310310-2002023310121111"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3132002233303301-3102011330212212-1201132302202201-0213033231101332-0111010120002031-0330003031010021-3213110011101303-1011031101020301"></a>

<a id="canonical-0103310301311323-1312232123120103-3102120301203303-3003232112330000-1000111113131032-0302010332330100-1133012000313220-2330013131221200"></a>

#### `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3332232033013000-2032222110022033-0001300033312133-0211031021233110-0110331002111312-1311101022331130-3223133323323212-3031212323210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0013103230120133-2032003020333120-1113103020030311-3300020222213333-1202230333110320-1210013302101223-1323322211113010-3321320103021301"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.xfcc_options

<a id="canonical-2213032130123000-1133101312211231-0130100322231020-3133332321101210-2202113113101213-2133221323233300-3123203031201133-1322320303110102"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202011012321113-2231330323231232-3131311021233020-0233012302320002-0112030002002303-2302211012131010-3112022310131011-1313132232231101"></a>

### Direct properties for `tls_tcp_auto_cert.use_mtls.xfcc_options`

<a id="canonical-1122302311011201-3233002131113212-0111022023223311-3002001100333310-2132111310012313-0201110102321230-2332103233021031-1231211232113132"></a>

#### `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```
