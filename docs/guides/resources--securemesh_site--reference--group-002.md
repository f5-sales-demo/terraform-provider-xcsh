---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-2321121210331020-1110021221021021-3123003132011002-2311012312020012-2330303200331100-0312122123220321-2003022001120033-3111212013210233"></a>

## name property — network_policies / 012303233312 / 4

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

<a id="canonical-0303201310031031-2323102001221311-3002030221031101-3110221232010120-3011303132222200-2031020103202213-3000233313230313-3320102323310231"></a>

<a id="canonical-3030032020212230-0301331323003202-1333232230022103-2010031332031123-0103303031101032-1021131310203111-1311200200323133-2213021103331323"></a>

## namespace property — network_policies / 012303233312 / 5

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

<a id="canonical-2010102330302122-2133321020013333-0023311313301131-3133013012022310-2313000203231010-2223102121032133-3121213101002113-3123322303113011"></a>

<a id="canonical-3221230002100003-2311221223001300-3002023301102131-0122021311021202-1203022203200131-1300233121202033-3211003221333122-0312003031101123"></a>

## tenant property — network_policies / 012303233312 / 6

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

<a id="canonical-3202012313332331-2221102211001223-1112132312122201-1101202213232312-2113232300203322-0022022122131200-2130203032223002-2113211002230221"></a>

## Next pages — network_policies / 012303233312 / 7

- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-3023131331003221-2030010320001301-1121003103323033-2220220101231222-3320232031203110-3232001230223213-1002021211312311-3003233011102220)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1003133003322303-2221311020000231-2132020123222101-2023312312123323-0130023221221313-2230022323203032-0201303102000312-2200122002000332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123001130000330-2311100310033001-2031310032111000-1330110301000003-0203313031213102-2231120312302033-3203121113022232-0102033132133022"></a>

## custom_network_config.default_config — default_config / 001203320103 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.default_config

<a id="canonical-0132110310222130-3311203300231032-3201323020201333-1202321122103233-1203330130233232-2100130003210300-1323110131102001-3332312212033011"></a>

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
default_config = {}
```

<a id="canonical-1013110011102211-0220210113233201-2030211200313222-2002003010111010-0022222112020012-2010301233220132-3213333120330312-1122331031012212"></a>

## Direct properties — default_config / 001203320103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323213201323221-2212332110131121-1023102023201120-3033110221333110-1033023202023312-1321221121133133-1120210033031320-1200103332130222"></a>

## Next pages — default_config / 001203320103 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3022201322330102-2133020001321012-0320312322210022-3212030112002322-1110002223002100-1020111001131323-2132012322101210-0121001203020112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222313203012330-1202123111102003-1020020000211332-0223303110031221-0003230000121130-2103121021100000-1203033100303333-2003130311210033"></a>

## custom_network_config.default_interface_config — default_interface_config / 303001210222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.default_interface_config

<a id="canonical-3213011112013032-3310111303321210-3011002102311303-2222032110200133-2123000302231211-3311013101000330-3032331031210203-1200101200200202"></a>

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
default_interface_config = {}
```

<a id="canonical-2113320032211321-3102113030011232-0132202211123132-2111210222202101-0102323233200003-3110021132300200-1000302002213103-2221303311220123"></a>

## Direct properties — default_interface_config / 303001210222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300330121213213-2310210030232210-3030110333332223-3213231123111121-2130013130333022-3330220322021300-1023310103223323-1103101222011022"></a>

## Next pages — default_interface_config / 303001210222 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2221000332132000-3211222221213320-0312003330022103-3233322233133113-0012301330310133-0012213130011033-0113021130103333-0222302012203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001220200021322-1302333230120300-0112101230031102-0121213133021103-3132222011220033-2120311131132312-3131132030213311-1131330011012032"></a>

## custom_network_config.default_sli_config — default_sli_config / 213313000000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.default_sli_config

<a id="canonical-3120010312003000-3021112032333311-3123232332031133-3020013323000023-0320020211133203-3323300311012121-3030230110232111-3203110000221320"></a>

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
default_sli_config = {}
```

<a id="canonical-0020222300022223-0110132010331221-3012330210313231-0333131333013212-3323120011123013-2100133022120033-3232230023022313-2012302332332133"></a>

## Direct properties — default_sli_config / 213313000000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330300210132030-0212312010220322-0111222303130113-3121201222213130-1020203122022200-2313210022131130-0201132211101030-3333111103100023"></a>

## Next pages — default_sli_config / 213313000000 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1032010001122222-1030130223330222-0330032002313232-1123210200131223-3232031003302033-3223312313301133-1013323200313311-2120020330132101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302203220301120-3111131221032103-1132223332130101-1020302221013120-1232020033030310-1021312130210010-1210131330313220-3312312021031233"></a>

## custom_network_config.forward_proxy_allow_all — forward_proxy_allow_all / 210320001310 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.forward_proxy_allow_all

<a id="canonical-0033113133211033-2200312302301331-2232032200220213-0332331032231221-3111013102222010-3120021202010333-1213203322322230-3313122332332220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-2232122211012020-3202031202111133-2032021212030002-2122102330013303-1122221321201211-1323221133230112-2212001130202010-1332312312101322"></a>

## Direct properties — forward_proxy_allow_all / 210320001310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301013231132023-3220310003032300-0313030331030331-2323311110121202-3313010010001011-2013222301103012-2111333322121021-1332302212020300"></a>

## Next pages — forward_proxy_allow_all / 210320001310 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220031003201213-3313010330010103-2222102222301011-0022203020232323-2013202332022101-2212210120020220-1113001321320002-3122210202101333"></a>

## custom_network_config.global_network_list — global_network_list / 222221132320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.global_network_list

<a id="canonical-3102110030132111-0123331011120312-2233011232332200-2320231102310003-0010002303131303-1233100000031103-3311120310001303-2233002123001201"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230120103233100-2322310123233001-1001113312032313-3131033330233030-3033232311103221-2121303203202311-0021331310230110-3012313213322300"></a>

## Direct properties — global_network_list / 222221132320 / 3

- [global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032): complete subsection reference.

<a id="canonical-3021000010213212-3333213112212122-1111002130100123-0311220220110312-3301320313030230-3310012123132231-2111230033120121-3310131212322112"></a>

## Next pages — global_network_list / 222221132320 / 4

- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220232301331303-2013100031133122-3033131321223111-2103122131302121-2320320331310022-1322222103223311-1212302322101121-1022232332303112"></a>

## custom_network_config.global_network_list.global_network_connections — global_network_connections / 211130122131 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- custom_network_config.global_network_list.global_network_connections

<a id="canonical-2200020230022100-2032000100333202-3320322010131102-3211113300331333-1310100010211333-3113133121231212-2313300133013330-3312112110110122"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103130330011030-1303311331020023-1222323122002312-1210100312131120-0220320111210121-3121132110222302-1112230310130123-3200211210003032"></a>

## Direct properties — global_network_connections / 211130122131 / 3

- [sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0321112010033003-1021323232122012-3031013201131003-2301221132111013-0322221323023203-2022330332113110-2030110223030331-1123332013232202): complete subsection reference.

- [slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0123201112001121-1320200113233302-1211211301231030-0111111133102211-2233123101103200-2311330331130231-2312220321021310-0001101221130002): complete subsection reference.

<a id="canonical-3100233100121321-0012332123221221-1332031211110302-0231313330230002-3223303012223322-2102230122232123-3110033231332133-2202310113031110"></a>

## Next pages — global_network_connections / 211130122131 / 4

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0321112010033003-1021323232122012-3031013201131003-2301221132111013-0322221323023203-2022330332113110-2030110223030331-1123332013232202)
- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0123201112001121-1320200113233302-1211211301231030-0111111133102211-2233123101103200-2311330331130231-2312220321021310-0001101221130002)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0321112010033003-1021323232122012-3031013201131003-2301221132111013-0322221323023203-2022330332113110-2030110223030331-1123332013232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122023212312210-3133201121201123-2233121211110303-1101323101200332-1313130002213223-2003303231311301-1113220330330220-1101322013122133"></a>

## custom_network_config.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 212311300103 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- custom_network_config.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3132131320231210-0103103230123002-0211023021122210-2230200101033332-1130101100203222-2330200133100301-0010301132000012-3333203032000132"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010200220203312-3013102202110103-0013021002333231-1130210221222332-3300022010203113-0313223203331022-2301013310110110-3221110202331313"></a>

## Direct properties — sli_to_global_dr / 212311300103 / 3

- [global_vn](resources--securemesh_site--reference--group-002.md#canonical-1022212013100213-2010330103020202-0102311233123322-1303032323211111-0311321001200011-2231213310333202-3133032003221121-3013021123022223): complete subsection reference.

<a id="canonical-0133302000310123-0131102331032030-2113302030022111-3232010213001120-1101123101130310-3031211101222020-3313300312222122-3113332122101122"></a>

## Next pages — sli_to_global_dr / 212311300103 / 4

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-1022212013100213-2010330103020202-0102311233123322-1303032323211111-0311321001200011-2231213310333202-3133032003221121-3013021123022223)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1022212013100213-2010330103020202-0102311233123322-1303032323211111-0311321001200011-2231213310333202-3133032003221121-3013021123022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202101310102333-3202231002030021-0312313033122222-0223001120321222-1122331013101222-0330231100222130-2333001230020100-0123133311200021"></a>

## custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 103230301102 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0321112010033003-1021323232122012-3031013201131003-2301221132111013-0322221323023203-2022330332113110-2030110223030331-1123332013232202)
- custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-1110133321120000-0102112023330021-1223302002203030-1013223112331232-2022131330233321-2223200300033221-1123220101113232-2211303233010330"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333120123220023-2120020121031021-1211133212223022-0102010332223130-1103133133302123-1131011313032113-3212012131033020-3122120312110121"></a>

## Direct properties — global_vn / 103230301102 / 3

<a id="canonical-1322013322330210-0312000331223311-2021203131013032-1302303320330220-1231222022101022-0112201023322333-1212131201123011-0013202100223223"></a>

<a id="canonical-3012300031131223-3222313003110121-2110030220323032-3333122313212222-2023203001313132-0030213112221231-3201020232110120-0313231313123303"></a>

## name property — global_vn / 103230301102 / 4

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

<a id="canonical-2113311123223020-2120232210331312-3000310211123223-3113102232223020-1032333333100001-2200101131130323-1133333103120212-2110223222313202"></a>

<a id="canonical-1023221230012200-1012121030302311-3322010002011100-1213123202230033-0001022212033233-0332302302312133-2021322022120221-1001222200312033"></a>

## namespace property — global_vn / 103230301102 / 5

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

<a id="canonical-2111003011223233-2122022213112300-1111001113103031-0032223131013311-3223311132311122-2101331122130320-1132123333032323-1230133123202131"></a>

<a id="canonical-2323032120123203-2212131012223003-2223233223112232-1210001003233023-0101233220133223-2112312223311203-3303023320122122-0332211121220120"></a>

## tenant property — global_vn / 103230301102 / 6

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

<a id="canonical-3303301121030002-1201123202301221-3303121332210002-2220103012212203-2123310133022023-1120032010001110-1121011330210103-3100302013323200"></a>

## Next pages — global_vn / 103230301102 / 7

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0321112010033003-1021323232122012-3031013201131003-2301221132111013-0322221323023203-2022330332113110-2030110223030331-1123332013232202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0123201112001121-1320200113233302-1211211301231030-0111111133102211-2233123101103200-2311330331130231-2312220321021310-0001101221130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320331300112300-0330002220103120-1323323010232033-2023100202112302-2200130010201012-0212003100130012-1213033303002211-2030220211311302"></a>

## custom_network_config.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 333233031010 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- custom_network_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1130122202111221-2101311313322203-2232022112210300-2020310233312121-2203201211230022-2311232323103210-2311131102111300-0210220020320221"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113320231203022-0331331033203200-3033100102130123-2211123201221312-2102112133311302-1131022323123213-2130123123221301-2110123103120221"></a>

## Direct properties — slo_to_global_dr / 333233031010 / 3

- [global_vn](resources--securemesh_site--reference--group-002.md#canonical-0202211233003202-3031103210321022-2020331132101121-0300123012100322-3301120322123133-2232020323303031-3113000201332011-0101232023113023): complete subsection reference.

<a id="canonical-3303031023203210-1101213303322332-2102202021322032-3030223001012300-1230030013023131-2333233230321130-2013021203312111-1013101122322330"></a>

## Next pages — slo_to_global_dr / 333233031010 / 4

- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-0202211233003202-3031103210321022-2020331132101121-0300123012100322-3301120322123133-2232020323303031-3113000201332011-0101232023113023)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0202211233003202-3031103210321022-2020331132101121-0300123012100322-3301120322123133-2232020323303031-3113000201332011-0101232023113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311033310111100-1130013202123001-2132320213032301-0331123111122210-0202110212311230-2330020203000203-0332110213112110-0300330033302332"></a>

## custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 331212200332 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-2202012302002002-1110010133131331-1121220233201310-2132103122130102-0302131132112003-2221222123010113-3300230332030031-1232112101330203)
- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-0001010201123010-0231110111230303-3103130303123011-0202201122032122-1233322230300110-0200220121232020-3233000122001022-0130102223120032)
- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0123201112001121-1320200113233302-1211211301231030-0111111133102211-2233123101103200-2311330331130231-2312220321021310-0001101221130002)
- custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0123230103131231-0220320002332001-1101322321131320-3103021202022210-3230202301002010-2023303012333122-1210000221002102-0333002012020003"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232202002320321-0021111112330013-3000322110203301-3322211111102121-2011231001321221-0130300121223123-0222030130220210-1102111123221212"></a>

## Direct properties — global_vn / 331212200332 / 3

<a id="canonical-0032211002133100-0321320030223013-2023021220302212-3213021312332310-1020201033332330-2320103221200001-0313231301200312-2211233102323133"></a>

<a id="canonical-3032220303200103-3030210332003001-3130102321110113-1100023122132230-3232130320023211-1320002112223333-0121230312313223-3110132133000302"></a>

## name property — global_vn / 331212200332 / 4

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

<a id="canonical-0101030131012133-3322101031123203-3000112230332212-3133130323202300-0012121211013332-0101101323023211-1131022111023233-2001230021110020"></a>

<a id="canonical-3101312113020203-2312313303231112-3132200231120231-1132200313110032-3320110013101111-0032131122110210-1301210311302011-2302230330230113"></a>

## namespace property — global_vn / 331212200332 / 5

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

<a id="canonical-3131012021301011-3032322101212033-2122311232212130-0003320331131133-1202131231131101-2120321310232031-0301313303023022-3322213330000333"></a>

<a id="canonical-2213302013003332-3030212332312110-2021121011110311-1233331000022231-0132101003211220-3333213322021111-2031221123211100-1310222101200222"></a>

## tenant property — global_vn / 331212200332 / 6

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

<a id="canonical-3322320331023121-0030323130032002-0333013011313213-1330210020021233-0012022011212020-3132310220101220-2330311123210000-0223212200112211"></a>

## Next pages — global_vn / 331212200332 / 7

- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-0123201112001121-1320200113233302-1211211301231030-0111111133102211-2233123101103200-2311330331130231-2312220321021310-0001101221130002)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120001201103022-3302100103033322-3321301131330103-1303233210320222-0130233113203223-2322032021300033-2210332223112301-1223330120012331"></a>

## custom_network_config.interface_list — interface_list / 030322000013 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.interface_list

<a id="canonical-3031231313103301-0111321211211020-0310322321322320-2210002113230301-1121212122023203-0000032233330210-1302203201113233-1103020210020100"></a>

Type: `"object"`. single nested block, Optional.

Configure network interfaces for this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210222310122222-2032103233230203-1121110111130112-1311113012312033-1201022223103322-1210001220110121-1303213031333232-1112210133332033"></a>

## Direct properties — interface_list / 030322000013 / 3

- [interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022): complete subsection reference.

<a id="canonical-2310312112301110-2122120223311322-0211213120002221-3210223321003212-0010113102012330-1110200132230130-3021101312013031-3112310020032323"></a>

## Next pages — interface_list / 030322000013 / 4

- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100132301120222-2133302133232030-2123120101200300-3332022011000321-2032001213113222-2030003221110231-3003132003232333-2310330222120111"></a>

## custom_network_config.interface_list.interfaces — interfaces / 201001210312 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- custom_network_config.interface_list.interfaces

<a id="canonical-2032133103223310-2032211132321000-0022021112112130-1312323203332222-2121331223311332-0031212310001033-1131210001302312-1111112312223201"></a>

Type: `"object"`. list nested block, Optional.

Configure network interfaces for this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dc_cluster_group_connectivity_interface_disabled",
    "dc_cluster_group_connectivity_interface_enabled"),
  validators.ConflictingListObjectAttributes("dedicated_interface",
    "dedicated_management_interface"),
  validators.ConflictingListObjectAttributes("dedicated_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("dedicated_management_interface",
    "ethernet_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031213032230131-3031120313030300-2113001320200323-3130302013230002-3010112032022023-1101011121013311-0033232102130310-0330322123303210"></a>

## Direct properties — interfaces / 201001210312 / 3

- [dc_cluster_group_connectivity_interface_disabled](resources--securemesh_site--reference--group-002.md#canonical-2222010301121031-1023120200100223-1131121120131203-3231223001213230-3130203022201023-2223222320212321-3100122203131131-1300301110012131): complete subsection reference.

- [dc_cluster_group_connectivity_interface_enabled](resources--securemesh_site--reference--group-002.md#canonical-0320203022022220-1301330022111213-1332110231301211-2000021300311101-2133323311101203-1010001133120323-2223222120220021-1032032033111220): complete subsection reference.

- [dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120): complete subsection reference.

- [dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-1303123102103231-1033032230222310-0030330020130000-1000211121021101-3320010013003201-1110023022030023-0130021310332223-1311313032202201): complete subsection reference.

<a id="canonical-3230221312230332-0200232300013121-0332122200300321-3322232102032101-0010001022030331-2011120130203321-1223020003103323-2003200213120330"></a>

<a id="canonical-2003030210221122-3112112200130333-1221231320111133-3020133030112021-1010010023220002-2322021131301332-2110202122323110-2330033302321212"></a>

## description_spec property — interfaces / 201001210312 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233): complete subsection reference.

<a id="canonical-0122001301203001-2120331302231300-3120313223320112-2330311000031311-3210313132011131-3230300130303122-1030303300330232-3012103022021211"></a>

<a id="canonical-2332131203033132-2000013012113022-3123003202320023-2202302010100020-3303022312230131-2022121031012002-0200222201200103-2303030101011322"></a>

## labels property — interfaces / 201001210312 / 5

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3021021332023220-3200022133300201-0320302201332102-3000223332032222-2011032223313031-0312333022210321-2113111113110203-2013131120312323"></a>

## Next pages — interfaces / 201001210312 / 6

- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](resources--securemesh_site--reference--group-002.md#canonical-2222010301121031-1023120200100223-1131121120131203-3231223001213230-3130203022201023-2223222320212321-3100122203131131-1300301110012131)
- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](resources--securemesh_site--reference--group-002.md#canonical-0320203022022220-1301330022111213-1332110231301211-2000021300311101-2133323311101203-1010001133120323-2223222120220021-1032032033111220)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-1303123102103231-1033032230222310-0030330020130000-1000211121021101-3320010013003201-1110023022030023-0130021310332223-1311313032202201)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2222010301121031-1023120200100223-1131121120131203-3231223001213230-3130203022201023-2223222320212321-3100122203131131-1300301110012131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103002213033030-0002103023022313-1310313120130201-1211022333232110-3311321101321000-3110113303301312-3012011231222323-0321320223220202"></a>

## custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled — dc_cluster_group_connectivity_interface_disabled / 003001201323 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled

<a id="canonical-1121023011212102-2011312132112320-2033210202101033-3321033133231032-1221021013322203-3000030003021222-2002010213111110-1313030100311100"></a>

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
dc_cluster_group_connectivity_interface_disabled = {}
```

<a id="canonical-3131303232301132-3212312121000222-3032112313313201-2323330211012333-3010133030312021-1113330213013131-1111130112220120-1211222130110101"></a>

## Direct properties — dc_cluster_group_connectivity_interface_disabled / 003001201323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131122103213301-1312330311120101-3222131320121332-2110211233233111-0103230012311231-2210310333132130-2003130130023211-3113020323332223"></a>

## Next pages — dc_cluster_group_connectivity_interface_disabled / 003001201323 / 4

- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0320203022022220-1301330022111213-1332110231301211-2000021300311101-2133323311101203-1010001133120323-2223222120220021-1032032033111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321213010032121-3231012112012321-0321300210113331-0303312331101303-3230312103331232-1123332003131211-0023302232310112-0203101310010201"></a>

## custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled — dc_cluster_group_connectivity_interface_enabled / 103101200011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled

<a id="canonical-3122003012100110-1100012112003233-2313330013223110-3222102000203102-0030122200232132-3313030122120121-0210133223110033-3101311123113332"></a>

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
dc_cluster_group_connectivity_interface_enabled = {}
```

<a id="canonical-1232220311121300-0331232100031023-0203111221131300-3000131110111303-2010100310012113-3231103331013103-1313010110102211-0312020003212131"></a>

## Direct properties — dc_cluster_group_connectivity_interface_enabled / 103101200011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331333233221311-1033203012332031-2103223233030020-3320001333010311-2103022333330331-0312031103100302-3233222031002313-1222102211201313"></a>

## Next pages — dc_cluster_group_connectivity_interface_enabled / 103101200011 / 4

- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031020110101230-1131221313303032-3321220221112333-0031222013013123-0211213001331323-0333032311101110-3103302012013111-1000212211112222"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface — dedicated_interface / 202033321112 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- custom_network_config.interface_list.interfaces.dedicated_interface

<a id="canonical-2013203101311331-1211130221323023-1000302320233223-3231100023000130-2033113230300323-0231301301133133-0032103303000111-1000110232323300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled")}
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
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]"
}
```

Terraform syntax:

```terraform
dedicated_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311232032312212-0132232323120333-1233013302022322-0300120331222003-0020202010120300-2132221123213332-2101110211323021-3202011222001032"></a>

## Direct properties — dedicated_interface / 202033321112 / 3

- [cluster](resources--securemesh_site--reference--group-002.md#canonical-1101300313131230-2012121300323302-0331330221112220-3310021333113201-3111333111020200-0330103101230033-3001122202131311-0322121322211302): complete subsection reference.

<a id="canonical-1111323103100302-3133203303312120-1231222003113333-1213310022230320-0032313320210301-3032031310312112-2310101131113112-2031231010111011"></a>

<a id="canonical-3100233103203121-0201022100210020-3010332110100132-2303213001021213-1133001113300033-0233103313103032-3321123133002331-1320300010332200"></a>

## device property — dedicated_interface / 202033321112 / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [is_primary](resources--securemesh_site--reference--group-002.md#canonical-0232310302012232-0320212302312201-2011233232312101-0231023001311133-2320101303201103-3022111111132231-1202003222330313-2102333313220121): complete subsection reference.

- [monitor](resources--securemesh_site--reference--group-002.md#canonical-3111320130120313-3211000330320030-1213233322213221-3211120300202021-2133330110232103-3010133130131031-0211330330211031-3220321102023211): complete subsection reference.

- [monitor_disabled](resources--securemesh_site--reference--group-002.md#canonical-2323213021222301-1230213212313233-2322303331023213-3313310132323032-2233031133123123-0032332010232320-1300013031220122-2310031233331010): complete subsection reference.

<a id="canonical-3032012030202202-0311003023312100-3301331332313131-2111113333212212-3231132313200323-0133322320132131-1213331001011333-0210102013231113"></a>

<a id="canonical-1203001233211213-3230012000013103-3012110200220122-2032221033032010-1322312210012312-0020203232103232-2010010030202103-3201321003212203"></a>

## mtu property — dedicated_interface / 202033321112 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-3101120222122310-0231112312301010-0101211103121030-2000231100210211-0223022330130020-2221021013233132-2312203110021121-2200023310120130"></a>

<a id="canonical-2333010221222030-3231113003302222-2303310303113321-3210223203323112-3002022323132122-2321020130202322-1213230031323313-3312202321002013"></a>

## node property — dedicated_interface / 202033321112 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--securemesh_site--reference--group-002.md#canonical-3213301023321323-2003211301300000-2310112100122020-3313321132313321-0120332002312112-3001322110003310-1012000332213213-0032200222111030): complete subsection reference.

<a id="canonical-0222101323122002-1110012223332123-2311210212311021-2100120111001202-2132010323001221-3201022120122021-3132300202222012-1303130112023122"></a>

<a id="canonical-1132312013311320-3003130101030031-0030112130212011-3222112222303313-2302030213323032-3110321211202303-0011213311212032-3012210012223232"></a>

## priority property — dedicated_interface / 202033321112 / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1012310100200011-2130212021030100-1112021322221110-1233033220033111-2322100123323212-2022301112133201-2203000100223132-2220000211033302"></a>

## Next pages — dedicated_interface / 202033321112 / 8

- [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-1101300313131230-2012121300323302-0331330221112220-3310021333113201-3111333111020200-0330103101230033-3001122202131311-0322121322211302)
- [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](resources--securemesh_site--reference--group-002.md#canonical-0232310302012232-0320212302312201-2011233232312101-0231023001311133-2320101303201103-3022111111132231-1202003222330313-2102333313220121)
- [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](resources--securemesh_site--reference--group-002.md#canonical-3111320130120313-3211000330320030-1213233322213221-3211120300202021-2133330110232103-3010133130131031-0211330330211031-3220321102023211)
- [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](resources--securemesh_site--reference--group-002.md#canonical-2323213021222301-1230213212313233-2322303331023213-3313310132323032-2233031133123123-0032332010232320-1300013031220122-2310031233331010)
- [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](resources--securemesh_site--reference--group-002.md#canonical-3213301023321323-2003211301300000-2310112100122020-3313321132313321-0120332002312112-3001322110003310-1012000332213213-0032200222111030)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1101300313131230-2012121300323302-0331330221112220-3310021333113201-3111333111020200-0330103101230033-3001122202131311-0322121322211302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123020231112130-1232000120301202-0131030320232223-0203320011032201-1032333222311003-0120002031201233-3210023130231302-0023133002231130"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface.cluster — cluster / 322221231122 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- custom_network_config.interface_list.interfaces.dedicated_interface.cluster

<a id="canonical-3100332220122303-1310013311312013-0111303013330300-0033331031131003-0121001122310111-0313322021301320-1331211300201110-3002102213110320"></a>

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
cluster = {}
```

<a id="canonical-3222122133000111-3023132203211031-3010033112301133-2012103132320322-1330200233313200-1313310012002112-1301232030000332-0310223301312330"></a>

## Direct properties — cluster / 322221231122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132313121200222-0002123322013112-1201002121101102-0302302323333023-1211122321311322-0102032021312022-1302222031010011-2003021131020210"></a>

## Next pages — cluster / 322221231122 / 4

- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0232310302012232-0320212302312201-2011233232312101-0231023001311133-2320101303201103-3022111111132231-1202003222330313-2102333313220121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321002322202333-0131300210031232-1020110130221200-1023310300110121-1321001113011330-2211031103220110-3221311013002221-2120023132113121"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface.is_primary — is_primary / 100022112001 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- custom_network_config.interface_list.interfaces.dedicated_interface.is_primary

<a id="canonical-3000301222323323-1302031123031020-3202010120303012-1120332201002322-3102012033101210-1023122003300332-1323113032310301-0021133212123031"></a>

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
is_primary = {}
```

<a id="canonical-3332110220011302-3333221102310000-2310121311320301-1322132021123223-1013020003031133-1101302331313133-2133000110113232-3103133330003300"></a>

## Direct properties — is_primary / 100022112001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010030122330000-3010100112233031-2211321201023233-2210330012300031-2332132310130231-2312110023010112-1312302200033310-3323012022331312"></a>

## Next pages — is_primary / 100022112001 / 4

- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3111320130120313-3211000330320030-1213233322213221-3211120300202021-2133330110232103-3010133130131031-0211330330211031-3220321102023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200103133311230-1332003322122313-0332123320210301-2212332211301303-1221310321213031-3223120133312203-0121120211030300-3323332133131200"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface.monitor — monitor / 202333231301 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- custom_network_config.interface_list.interfaces.dedicated_interface.monitor

<a id="canonical-3312021020220211-3330222320233310-1100110221020301-3231202201213013-2232322120232200-2031011131313021-3003302332331033-3312301123113211"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

Receipt-pinned upstream constraints:

```json
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
monitor = {}
```

<a id="canonical-1220210311303012-3311312322210200-0001310331233113-0212021013020301-0210322221020220-1233202313212001-0203233021120220-0011020112032020"></a>

## Direct properties — monitor / 202333231301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203321010320211-2021231221323123-1113123033033002-2202112112120112-2102223100213330-2122202102000221-1230312002210323-0110222233220103"></a>

## Next pages — monitor / 202333231301 / 4

- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2323213021222301-1230213212313233-2322303331023213-3313310132323032-2233031133123123-0032332010232320-1300013031220122-2310031233331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021130321221021-1331001211030322-1133212330201013-2210300132021232-0311213001210232-0332202103303321-0200110020131133-3131310002133020"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled — monitor_disabled / 331111311121 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled

<a id="canonical-2012110200103223-0030010333123323-1302023313130023-1202010030122202-2013011111223013-3233210323132111-0230321312111310-2201113122112221"></a>

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
monitor_disabled = {}
```

<a id="canonical-2331033202333011-1010012100231022-0212332303313210-1330121003130002-0031230200032323-2330311013021101-2112133120333021-2310133133212101"></a>

## Direct properties — monitor_disabled / 331111311121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232220210331210-2231022131310232-1220023312222323-3020022121332311-0000111331013230-2202120033231001-2000202213033133-0121013112133233"></a>

## Next pages — monitor_disabled / 331111311121 / 4

- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3213301023321323-2003211301300000-2310112100122020-3313321132313321-0120332002312112-3001322110003310-1012000332213213-0032200222111030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123333120002223-2330111102101012-3220022132211302-2003303332132110-1113232012122213-1030112002110001-0321222312212203-3003132222012201"></a>

## custom_network_config.interface_list.interfaces.dedicated_interface.not_primary — not_primary / 223230003302 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- custom_network_config.interface_list.interfaces.dedicated_interface.not_primary

<a id="canonical-3321200320221230-3303003003200202-1231122002003020-2031213320130231-2222021223320103-0002103311112110-3030201310213131-1300132203101213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for not primary.

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
not_primary = {}
```

<a id="canonical-3031032301133020-0112113012113101-3202332303100313-2032131331033123-1330033200222312-1302121123303111-2332000313013230-0020202200113310"></a>

## Direct properties — not_primary / 223230003302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131033021020302-0323121231200131-1232202102103121-0002203223302212-0102331222111323-3033201200120230-0121223033221030-3332001031210300"></a>

## Next pages — not_primary / 223230003302 / 4

- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-1331231123132301-2202103323302333-3110213000001002-1011102131321001-2002330121110010-1213332212303012-3301331031131312-3232223013133120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1303123102103231-1033032230222310-0030330020130000-1000211121021101-3320010013003201-1110023022030023-0130021310332223-1311313032202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301003020321232-2003223002121300-1220300000323102-0200331100202332-0103002303302213-3122301321001312-3313010010313011-3033010221301311"></a>

## custom_network_config.interface_list.interfaces.dedicated_management_interface — dedicated_management_interface / 321223032233 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- custom_network_config.interface_list.interfaces.dedicated_management_interface

<a id="canonical-3301323030000332-1201131333002310-3330220203332131-3010200300123023-0022001030202301-3131001332213323-0012302101012310-0003100202012022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node")}
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
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

Terraform syntax:

```terraform
dedicated_management_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120213322020332-2333021211000222-0022122210103022-0311211221330222-2033221222233133-1131011220232222-0321232331202002-1301033301112100"></a>

## Direct properties — dedicated_management_interface / 321223032233 / 3

- [cluster](resources--securemesh_site--reference--group-002.md#canonical-0113313011001101-0211312021330012-0233120332120320-3101322211133313-1113123110010311-3231200133100110-3231321021202202-2203301130130300): complete subsection reference.

<a id="canonical-0323302323210231-2102222131112033-0101320021232012-0000312132123131-3302113331303320-0300013101231012-0123133333312011-3100203321003000"></a>

<a id="canonical-1223123213132232-1120003103112222-3133101221213123-1331331321332212-0221332111133220-3212020320302033-3311202112313020-0310302222213130"></a>

## device property — dedicated_management_interface / 321223032233 / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3121031200102111-1030300130031101-1312101303213013-0233332230103222-3011202233110013-1121111202013110-0002002013120022-3233122221031310"></a>

<a id="canonical-2032101232223100-1213201232013312-0020132020202331-3010212023121110-1030010322222212-1100130230122212-0010101211021201-1130223221220302"></a>

## mtu property — dedicated_management_interface / 321223032233 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-1023321210203333-1332333231223211-2302110302200230-0231220303030103-0221123012021310-0330332230022030-3333202320123030-0232101210013233"></a>

<a id="canonical-0301111113103230-1120030120110003-1120002113331332-3213232213311101-1102000221033110-1331321100003221-2230232033023133-2031131023030303"></a>

## node property — dedicated_management_interface / 321223032233 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3321322332320323-1223301220211102-2131302112100332-3230210031101120-2100112103323323-1120002011133131-2032311231022321-0122212120320303"></a>

## Next pages — dedicated_management_interface / 321223032233 / 7

- [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-0113313011001101-0211312021330012-0233120332120320-3101322211133313-1113123110010311-3231200133100110-3231321021202202-2203301130130300)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0113313011001101-0211312021330012-0233120332120320-3101322211133313-1113123110010311-3231200133100110-3231321021202202-2203301130130300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211321023331203-1213213133232210-3001021230333312-3323301333303123-2202310300103101-3121002200013313-2100311200130001-1233233201132001"></a>

## custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster — cluster / 210113210011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-1303123102103231-1033032230222310-0030330020130000-1000211121021101-3320010013003201-1110023022030023-0130021310332223-1311313032202201)
- custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster

<a id="canonical-2323233313312111-2332001212000031-2012033202230313-2131121233032213-0031121330222221-0001210011301103-3232021203230212-1130232213121032"></a>

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
cluster = {}
```

<a id="canonical-1102002001303120-3223032033332331-0120203012103122-2101211122113131-3123133322312113-0123031101101031-2022301302222212-0222021322313223"></a>

## Direct properties — cluster / 210113210011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002120203330133-1322103213033200-2101210222212333-2311333132213222-0002332231032101-1222333020100212-1130330221130103-1113100022300021"></a>

## Next pages — cluster / 210113210011 / 4

- [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-1303123102103231-1033032230222310-0030330020130000-1000211121021101-3320010013003201-1110023022030023-0130021310332223-1311313032202201)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333130100210220-0131301331213131-2312000022011301-0301201031021311-0303030101213211-2200321020302211-3320011113021012-2001103222013022"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface — ethernet_interface / 303031020213 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- custom_network_config.interface_list.interfaces.ethernet_interface

<a id="canonical-1031003210123321-0013200222022231-3031103030300110-0213213302101332-3100112313233323-3032200010011031-0023010122201203-3032211002201023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
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
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123002213310331-2231102303200000-2233132020120233-0311013310232022-3213022100123032-1000231313210133-1023112232010212-3100321310221312"></a>

## Direct properties — ethernet_interface / 303031020213 / 3

- [cluster](resources--securemesh_site--reference--group-002.md#canonical-1110113300022313-2212023022220122-3331031230301112-0013331331123030-2032022312202302-1302323202033330-2120211103001310-3003022323003330): complete subsection reference.

<a id="canonical-0122132122330103-2033200201213032-2232331020001230-3133202001133322-2213110212111022-1003320222203133-0111101130331332-3020223102212102"></a>

<a id="canonical-2232331113113022-2201312112211103-0010101210231200-1333213321120201-2100200031232311-2013320300310333-0212131032103110-2312302201120020"></a>

## device property — ethernet_interface / 303031020213 / 4

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](resources--securemesh_site--reference--group-002.md#canonical-3211121331120122-0231223000220130-2301113012031112-1100012321023200-1321212311233012-0103110101031203-3331321203102012-0122331120302110): complete subsection reference.

- [dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133): complete subsection reference.

- [is_primary](resources--securemesh_site--reference--group-003.md#canonical-3011030111020311-3300213332323121-3231230031321023-0301121230030300-3100021313302312-2201021132221213-3123112332332103-0233110332301002): complete subsection reference.

- [monitor](resources--securemesh_site--reference--group-003.md#canonical-1020021101002011-1000113331021323-2023213103023313-0102213210333203-2100120231331211-2123023102011310-0013200131000330-1011301211211103): complete subsection reference.

- [monitor_disabled](resources--securemesh_site--reference--group-003.md#canonical-1303013201121301-0130130300113220-1020321001031001-0213012011113330-1103230113211333-1303231201331133-0221212120232313-2231322111002312): complete subsection reference.

<a id="canonical-2101122022313310-2313110020003331-3212032311103300-0313212332013320-3021111221223031-3111311123201123-3322323012121211-0001112210200303"></a>

<a id="canonical-1320232221001303-2322303201022301-0113212221201110-3212022111221122-0101111303000122-3033132221010322-3133221121203110-3032103302202321"></a>

## mtu property — ethernet_interface / 303031020213 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-2112233221302010-0212003311133123-3303222201001222-1011223313303222-0030211210333221-1022322202310310-1002110233021312-3330132111213123): complete subsection reference.

<a id="canonical-0221303232100311-3011011022203013-1322310332312031-0022201322033303-0122112023312000-0033320202311301-2211133133013000-2011202031032320"></a>

<a id="canonical-0312022021223130-1211213013030221-0231001323232202-0200301131100310-1021300203102011-3122222013130131-1331320010202323-2011012331232103"></a>

## node property — ethernet_interface / 303031020213 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--securemesh_site--reference--group-003.md#canonical-1312002312303010-0021100322330231-3103212331202213-3012013002111132-3203123230100320-1221132011312223-0310301021313303-2001301030221133): complete subsection reference.

<a id="canonical-0031023222123301-3102222211030221-3200323020333021-1111131110221323-3220333012303012-2103222321223021-1332013230200233-3132331101213020"></a>

<a id="canonical-1212211203100211-3003321002313310-3030322220121003-0033011032311330-1102221031121112-3020222013111112-1201122102311013-2311311320230021"></a>

## priority property — ethernet_interface / 303031020213 / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](resources--securemesh_site--reference--group-003.md#canonical-3221212220111121-0221231002312232-2211313013313322-0200322021111302-1003113223232130-0111332001032000-3112323122221001-3311321022032121): complete subsection reference.

- [site_local_network](resources--securemesh_site--reference--group-003.md#canonical-0302131022023122-2201211001033022-0033312030231003-0231210331200321-0311320000313032-3301231031020302-1132323102000021-3202320201212200): complete subsection reference.

- [static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210): complete subsection reference.

- [storage_network](resources--securemesh_site--reference--group-003.md#canonical-0302221212130113-3210032322320332-0131033110232323-2232202101222201-0120031231212101-1031312131332010-0122112232022201-1011320133013212): complete subsection reference.

- [untagged](resources--securemesh_site--reference--group-003.md#canonical-2030323210031111-2122220222013201-1221133130030112-2201300133321132-2131032320202110-0220030001210203-1202200303313102-0312132100233013): complete subsection reference.

<a id="canonical-3021202021101201-3231211133003011-1330012200031233-2300231001300223-0212020002022033-2132101113220232-2321021122022001-2132031131022030"></a>

<a id="canonical-2133001123201131-3230323131102330-1312312321013323-3330123311023211-1103300110023211-3312030003031322-3333330223032122-1131023223002203"></a>

## vlan_id property — ethernet_interface / 303031020213 / 8

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-1022020101011230-2221333310130000-3303331200020000-1320233231001120-0110223131113332-3110022121123302-0023111100010110-0130122301313303"></a>

## Next pages — ethernet_interface / 303031020213 / 9

- [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-1110113300022313-2212023022220122-3331031230301112-0013331331123030-2032022312202302-1302323202033330-2120211103001310-3003022323003330)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](resources--securemesh_site--reference--group-002.md#canonical-3211121331120122-0231223000220130-2301113012031112-1100012321023200-1321212311233012-0103110101031203-3331321203102012-0122331120302110)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](resources--securemesh_site--reference--group-003.md#canonical-3011030111020311-3300213332323121-3231230031321023-0301121230030300-3100021313302312-2201021132221213-3123112332332103-0233110332301002)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](resources--securemesh_site--reference--group-003.md#canonical-1020021101002011-1000113331021323-2023213103023313-0102213210333203-2100120231331211-2123023102011310-0013200131000330-1011301211211103)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](resources--securemesh_site--reference--group-003.md#canonical-1303013201121301-0130130300113220-1020321001031001-0213012011113330-1103230113211333-1303231201331133-0221212120232313-2231322111002312)
- [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-2112233221302010-0212003311133123-3303222201001222-1011223313303222-0030211210333221-1022322202310310-1002110233021312-3330132111213123)
- [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](resources--securemesh_site--reference--group-003.md#canonical-1312002312303010-0021100322330231-3103212331202213-3012013002111132-3203123230100320-1221132011312223-0310301021313303-2001301030221133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](resources--securemesh_site--reference--group-003.md#canonical-3221212220111121-0221231002312232-2211313013313322-0200322021111302-1003113223232130-0111332001032000-3112323122221001-3311321022032121)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](resources--securemesh_site--reference--group-003.md#canonical-0302131022023122-2201211001033022-0033312030231003-0231210331200321-0311320000313032-3301231031020302-1132323102000021-3202320201212200)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210)
- [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](resources--securemesh_site--reference--group-003.md#canonical-0302221212130113-3210032322320332-0131033110232323-2232202101222201-0120031231212101-1031312131332010-0122112232022201-1011320133013212)
- [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](resources--securemesh_site--reference--group-003.md#canonical-2030323210031111-2122220222013201-1221133130030112-2201300133321132-2131032320202110-0220030001210203-1202200303313102-0312132100233013)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1110113300022313-2212023022220122-3331031230301112-0013331331123030-2032022312202302-1302323202033330-2120211103001310-3003022323003330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111332300331223-1020333322032301-3123233230200211-2132101312210010-0010012123022301-0021223021201103-1313012133331210-3122321313012212"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.cluster — cluster / 201032222222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.cluster

<a id="canonical-0132131003210032-0322301220223313-3200103312230232-2123002313310202-0330233133122203-3201302021233122-3211013320110323-1300301000023000"></a>

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
cluster = {}
```

<a id="canonical-2233001313222131-1222330310210100-2102020022110312-1220103313222001-3210001310132003-2100001301313123-2110033321302321-3301201222032132"></a>

## Direct properties — cluster / 201032222222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223212120033220-0033303220301112-0201321212330122-3103230030103232-3203303221203202-1113100033200002-0230002221202201-3303323030010110"></a>

## Next pages — cluster / 201032222222 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3211121331120122-0231223000220130-2301113012031112-1100012321023200-1321212311233012-0103110101031203-3331321203102012-0122331120302110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000230211011303-1002300121121012-3322320113211330-3301331032121100-2213131322231030-3203331101103232-1000311313310121-0303021332030130"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client — dhcp_client / 101113222000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client

<a id="canonical-3200333213221023-0000220311321133-2123030023232122-1230130203320230-1301323031231033-2320332232011200-3331132133122331-2223232301130023"></a>

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
dhcp_client = {}
```

<a id="canonical-0220222101012000-3113113000220023-1022221013001133-1303302233200123-3323231000231132-2233311211121232-1032113003032301-2312002230233003"></a>

## Direct properties — dhcp_client / 101113222000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310301201011012-3333020231221000-2312310313230200-3231231000233221-1203013202333213-2320231101311030-3122003311321002-0333321333102311"></a>

## Next pages — dhcp_client / 101113222000 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302202302211033-3330033023213013-1110113131311001-0031013221310022-0332300120220102-1230212112123321-3123332301330232-3212002233323223"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server — dhcp_server / 200231031213 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server

<a id="canonical-1122322021331212-1031311233110231-3032010032220012-0102103312001022-0211301022132232-0031232333002212-2212103333202123-3030130313201111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312220003000102-2310021022331200-1331323011101020-1102122203333231-2133100211123003-3012313013033023-0220013310312302-2320122111331111"></a>

## Direct properties — dhcp_server / 200231031213 / 3

- [automatic_from_end](resources--securemesh_site--reference--group-002.md#canonical-0211302201230033-1322011213332122-0110032211103302-2210212103122100-3132332110112331-3101210033303003-0020132012330211-2122111003213000): complete subsection reference.

- [automatic_from_start](resources--securemesh_site--reference--group-002.md#canonical-1310130230211022-2212022111000201-0131220132320311-0103121021022010-2130030022023122-1302001101322001-3232202313110301-0210012230033231): complete subsection reference.

- [dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131): complete subsection reference.

<a id="canonical-0011031333313230-0322311121113301-0301102132032303-3013202031121303-0123321110103322-3302333013112101-0030303201023322-2123201013232232"></a>

<a id="canonical-1111111000022102-0211203012001210-2200132100302230-2313022131330311-1002313121133223-3111300210202223-1200012331101312-3033231302300020"></a>

## dhcp_option82_tag property — dhcp_server / 200231031213 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0311333122132032-3303033003103102-1231301200030011-3323103133013002-3222222103321313-2031212012120323-1210321223231210-2013030310220321"></a>

<a id="canonical-1121020223003320-3202013211332331-3333301031133102-0330213131032100-1031331321132033-3111322120113330-1100013022021131-3012022201202323"></a>

## fixed_ip_map property — dhcp_server / 200231031213 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-3223010300203233-2311113032020013-3331022000231112-2033230313131000-0231303322222000-3233133132230201-3300020122313133-1203311331230002): complete subsection reference.

<a id="canonical-2223211333332112-2231000230123203-3312032103021313-1110002332133212-0311201313312233-1001232112331033-3022111131311232-1233301232300330"></a>

## Next pages — dhcp_server / 200231031213 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--securemesh_site--reference--group-002.md#canonical-0211302201230033-1322011213332122-0110032211103302-2210212103122100-3132332110112331-3101210033303003-0020132012330211-2122111003213000)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--securemesh_site--reference--group-002.md#canonical-1310130230211022-2212022111000201-0131220132320311-0103121021022010-2130030022023122-1302001101322001-3232202313110301-0210012230033231)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-3223010300203233-2311113032020013-3331022000231112-2033230313131000-0231303322222000-3233133132230201-3300020122313133-1203311331230002)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0211302201230033-1322011213332122-0110032211103302-2210212103122100-3132332110112331-3101210033303003-0020132012330211-2122111003213000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003320012232023-0102311230132310-2131210300121012-1211330003013022-2303231112122103-0310312330000311-1112313232200330-2213020033120212"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end — automatic_from_end / 103012011120 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-2101233300302011-3310303230221022-0120023323222221-0122100332120312-3010130311321322-2332011000011230-0213113121021312-3223010031200132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-1231321323313022-3113201020112031-1230333202322221-1113312320020111-2030211103012001-3231321022123022-1332233012230300-0130131233002200"></a>

## Direct properties — automatic_from_end / 103012011120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113313321311200-1010310202010330-0020221301301223-2032232011002102-1223100132133203-0132220213220322-2013031032211022-1001111302130232"></a>

## Next pages — automatic_from_end / 103012011120 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1310130230211022-2212022111000201-0131220132320311-0103121021022010-2130030022023122-1302001101322001-3232202313110301-0210012230033231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112111322230022-0211210011120033-3113303012131112-3030312212303031-0330301010310300-0230231323311132-0330132202201110-2303322230221012"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start — automatic_from_start / 232210212122 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-0102121110330323-1303023131033002-0310131303122233-2111100023220021-2100113011232221-0113120330200131-3121222103031201-1130022120021313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-1101022010332232-0333313110211120-2101113133131020-3133111031303220-1313023112000023-3103220320223003-1301132122312031-3013010002232223"></a>

## Direct properties — automatic_from_start / 232210212122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121220201013201-2121002123010021-2231201313113223-3120310101013300-2132300111310300-2321022311000222-2022300120100211-0202100322102022"></a>

## Next pages — automatic_from_start / 232210212122 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212323302200120-2300003021131130-0212001301310112-3213220123301211-2200200131233211-0211013323300222-0022312322301101-2222003022031321"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks — dhcp_networks / 020133013201 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-1130333220231003-1332233031213033-3330113001131022-1022021213210302-1100212312223220-2330121201331322-0313013010321132-0300222002111023"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100110222010033-3011130001130221-3113023032022120-3322021012112311-2003121311121100-2130110102220103-1233213133101323-2301310322301232"></a>

## Direct properties — dhcp_networks / 020133013201 / 3

<a id="canonical-0320332111002330-0233001103313032-0122013332322211-2301333003321210-0233111000220033-0103320013202212-1210300102200131-0223130313102111"></a>

<a id="canonical-3313101322212020-0000220330002100-0032003312132122-2100120201123032-0031031310113221-2311321213011231-0103001020020202-1231321010100200"></a>

## dgw_address property — dhcp_networks / 020133013201 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3302203233101223-2023100312030332-0131023032202123-0201200100213010-3312232311100012-0001221311311130-0223303011132333-1031202302101110"></a>

<a id="canonical-2131300113001332-3203203310122303-3230123130330221-1003230331322033-0110201003031132-1332303033312313-1303223301203232-0123132001221220"></a>

## dns_address property — dhcp_networks / 020133013201 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [first_address](resources--securemesh_site--reference--group-002.md#canonical-2230231331211311-3332122002322021-1000012121323003-3332303212231231-0111222300313231-0021301001221331-3323312132231021-0201113010331033): complete subsection reference.

- [last_address](resources--securemesh_site--reference--group-002.md#canonical-0203322223302130-3103011320310221-0230113330301323-0103020320330122-3223031032231232-0123313202033323-2313203312202322-0130312221130122): complete subsection reference.

<a id="canonical-0020013120211222-0213300222232123-1330233332323112-0001332112323310-0023130211233020-2101130011200212-0103121332333031-1301330103300321"></a>

<a id="canonical-1000021203220003-2311230213021133-0003012230300003-1031023120121002-3100123321102303-0321121102221223-2100233301103213-2002010012020311"></a>

## network_prefix property — dhcp_networks / 020133013201 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0131320311000331-1323233300232312-0220101030320001-1203232100001132-1102001301101230-1100323303230323-2301132220110032-0301033003113321"></a>

<a id="canonical-1230001102101113-2121300120220310-2022302030000320-2111310030113320-0112122120201300-0131031110120323-3221312112310210-1103303201330110"></a>

## pool_settings property — dhcp_networks / 020133013201 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site--reference--group-002.md#canonical-1032010023100300-0213020111130133-0000213120112120-1130021111130211-1010300011130210-3223300000201323-2001200310212130-1022103121221321): complete subsection reference.

- [same_as_dgw](resources--securemesh_site--reference--group-002.md#canonical-2033133102003212-2311113330001332-1031312312033320-0203233313330303-3322311101300033-3122100030121131-3333023132012000-1210010031200230): complete subsection reference.

<a id="canonical-1202223231010100-0202000121100023-2201121201231300-0302302203223111-1133333022032300-2003330020020123-3302031102213232-1113201331300122"></a>

## Next pages — dhcp_networks / 020133013201 / 8

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--securemesh_site--reference--group-002.md#canonical-2230231331211311-3332122002322021-1000012121323003-3332303212231231-0111222300313231-0021301001221331-3323312132231021-0201113010331033)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--securemesh_site--reference--group-002.md#canonical-0203322223302130-3103011320310221-0230113330301323-0103020320330122-3223031032231232-0123313202033323-2313203312202322-0130312221130122)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](resources--securemesh_site--reference--group-002.md#canonical-1032010023100300-0213020111130133-0000213120112120-1130021111130211-1010300011130210-3223300000201323-2001200310212130-1022103121221321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site--reference--group-002.md#canonical-2033133102003212-2311113330001332-1031312312033320-0203233313330303-3322311101300033-3122100030121131-3333023132012000-1210010031200230)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2230231331211311-3332122002322021-1000012121323003-3332303212231231-0111222300313231-0021301001221331-3323312132231021-0201113010331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231003320303011-1121100022322033-3031001110112311-3332130311331200-0320122022003233-1302202102330030-1212001313101002-1330120212211121"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address — first_address / 212323023210 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-0131230321121231-2002302311312131-3203111110031303-0103203102201111-0310021211301012-1120301200201031-1020123122012321-1303112202220012"></a>

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
first_address = {}
```

<a id="canonical-3231000312332000-3303322111200313-0331330030020213-2312032020320122-3023331003020123-2211302013101231-0132120232212110-0110320003132221"></a>

## Direct properties — first_address / 212323023210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332122033022120-3031231021133233-1120122112332212-2300301301030321-0010020303201313-1323223233231223-2031232300122320-1032102033110001"></a>

## Next pages — first_address / 212323023210 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0203322223302130-3103011320310221-0230113330301323-0103020320330122-3223031032231232-0123313202033323-2313203312202322-0130312221130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033330113021332-3021301211310102-2033133033233232-2212111313233030-1221020112103000-1331020013332002-2332231203130023-2121010003201321"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address — last_address / 010032233303 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-0202132222010003-0313200331313120-0111130221120300-3111112200111131-3233223011320032-1023223220223213-1131202102020223-0331133010222312"></a>

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
last_address = {}
```

<a id="canonical-0223100113112322-3332322310023231-0302322022320203-0023311000201000-0121312323011132-3033322122201310-2111203312011212-2102023221222333"></a>

## Direct properties — last_address / 010032233303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022302012011313-0311121123330332-0303222020021213-2210110032001333-3102200321232320-1320231301123030-2020330320112012-1111123000100213"></a>

## Next pages — last_address / 010032233303 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1032010023100300-0213020111130133-0000213120112120-1130021111130211-1010300011130210-3223300000201323-2001200310212130-1022103121221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311002203102303-0200133112020031-3031110312033220-0112230322310213-3222323302310023-1020231023323132-1202002232200000-0122030101300030"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools — pools / 212322033022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-3030213331123103-2012002003222313-0230110003333311-0132302111211031-3102220301310333-3303310213121131-2113321110233000-0132130112013022"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013003030122103-1223120303131120-2120302213033300-3112111122123303-2030223320100010-0311203100010111-3211121332122310-1203222011111122"></a>

## Direct properties — pools / 212322033022 / 3

<a id="canonical-0312122120030101-1212003330220130-0322013130300321-1301310332332103-3032102033322303-2011302031013112-0023321331013330-0202123020332120"></a>

<a id="canonical-3002010303002201-3220111202132221-2131211321000010-3001233102132201-2113210321313300-2130223213021333-1030012000213202-2001231213313202"></a>

## end_ip property — pools / 212322033022 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3131012130201310-2312333213231121-3313020313212321-1102021331223100-1001220112332202-3320210103032303-0330221333312303-2303231130213232"></a>

<a id="canonical-1233123213031112-0300232310101203-0333201312333012-1220112201233012-1201221322300113-1001010021032211-3010221102000031-3110031300323110"></a>

## exclude property — pools / 212322033022 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1011011133313003-1022010010012123-3120022011101101-2102021133122001-1112320121120322-0031322331231001-2132212123202000-0030331110200100"></a>

<a id="canonical-0311012303231033-1231322020313223-2213111130010110-0001221320103310-0221012103302312-3100112331011213-2201120310233100-0120210230033320"></a>

## start_ip property — pools / 212322033022 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3231201111232331-1001100211220303-0223222131003313-1103110113131102-3210011201331001-2021203302332323-1230312333130101-1301030232012132"></a>

## Next pages — pools / 212322033022 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2033133102003212-2311113330001332-1031312312033320-0203233313330303-3322311101300033-3122100030121131-3333023132012000-1210010031200230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200131021121213-3123021223013301-0313221023212010-0303112310030002-1300123200222202-2100032122032013-1021113230103221-0120223333020213"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 310031220210 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1221211020131221-1110022301031000-0222020001303110-0203110120210121-0302111303211102-0023333131333032-1102102123232301-1033111332302333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-2013030133330321-2023121130121103-2011100123200123-1310013202302030-0331201220331023-0101300022222310-3332313213020330-0022103022213310"></a>

## Direct properties — same_as_dgw / 310031220210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132320222310010-1122112220003103-0013012203320303-1032133032022331-1132020232131033-2232123221310211-1100221123213333-3312123333222230"></a>

## Next pages — same_as_dgw / 310031220210 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-2110210110003313-2121020233120210-2202322310321233-1320233233322132-3233030202213313-2212021201103010-0220033033120120-0131211300022131)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3223010300203233-2311113032020013-3331022000231112-2033230313131000-0231303322222000-3233133132230201-3300020122313133-1203311331230002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233323312333130-0031200013033003-0330032302032103-1030210312003033-1333102210012031-3020223332320003-0102221201103201-1313033300330322"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map — interface_ip_map / 202323122221 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-0212233001123230-2121202202310120-2032300302011100-3203111322322113-3220211131221333-2201330302103311-3200333110311303-3012213330232232"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

Receipt-pinned upstream constraints:

```json
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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002120131230132-2311203121112200-3022223000303000-3013332133031220-0120212320312303-2303303113320203-2300222221232103-2321132012030132"></a>

## Direct properties — interface_ip_map / 202323122221 / 3

<a id="canonical-3111021021313321-3213113012001222-3030111210111002-1211011121203120-3122220201330100-1110030133213022-3201120230231132-3122001300030320"></a>

<a id="canonical-2311012032121222-1202121232032202-1330011331023332-3002200323322011-1320033010101111-1220332132022102-2231320030332202-2003231101020123"></a>

## interface_ip_map property — interface_ip_map / 202323122221 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-2333323301002312-0212321131210023-1220232203123003-1020320230101021-3030321031330330-2230131130321102-2131111120221230-1101211320101212"></a>

## Next pages — interface_ip_map / 202323122221 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-1320200002133002-3302013113010131-0222002301321110-3302120321201130-0233013302200330-0003203321113202-1113002011223122-3111322222333233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031000302113001-1030213013201231-3100020103032311-2000212301330031-2233301313122133-0122021223332120-0110033220021222-0233130033113201"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config — ipv6_auto_config / 100003033003 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

<a id="canonical-3321203313230201-3120112203113011-0300311020111331-3301233232103322-1331013222333001-3312123211003103-3203111221112002-3022213033332123"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232211223310222-2111323213323131-1033021013012023-1122011030012130-0223000122013130-1000030123120111-1112233232323223-0321332122201331"></a>

## Direct properties — ipv6_auto_config / 100003033003 / 3

- [host](resources--securemesh_site--reference--group-002.md#canonical-1023111110203211-3001020322213233-2313102330112331-0111303320203100-2310332133133132-0120130213132232-1033330011221120-3230301301330103): complete subsection reference.

- [router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201): complete subsection reference.

<a id="canonical-2023120121032313-0013201333301033-0112222321331222-2333020033230232-3033121301210320-2301333113110220-2323302103110313-2120331332113111"></a>

## Next pages — ipv6_auto_config / 100003033003 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](resources--securemesh_site--reference--group-002.md#canonical-1023111110203211-3001020322213233-2313102330112331-0111303320203100-2310332133133132-0120130213132232-1033330011221120-3230301301330103)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1023111110203211-3001020322213233-2313102330112331-0111303320203100-2310332133133132-0120130213132232-1033330011221120-3230301301330103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321100032201300-2301121320001203-2320123012130233-3032301313321021-1222313303123331-2121031210231231-0100300212231102-2021003002300032"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host — host / 133131013022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host

<a id="canonical-3322011220321130-1131210333103210-1232113033331130-2223231200300300-1131002032303022-1220000202303002-0000013211322300-2231232012121332"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-0203110330211322-2233003203020303-1122222030200132-3322010110320131-1223133012210300-2230011330131102-1301110230310220-1200130201112123"></a>

## Direct properties — host / 133131013022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031102220130333-0310230033210311-1300213120011101-2112330131312323-3233210113313010-2230133120100330-3130301310233332-3322031301021000"></a>

## Next pages — host / 133131013022 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121210200130131-2031013303103023-2331011003223032-2022332030320010-1320100032223012-1120101322302231-2002103333210103-2302200302101012"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router — router / 301131103101 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

<a id="canonical-1233102012103223-0130110331103211-0123202023003203-3332022303122000-0122033020133101-3210212301133021-3131313222002200-1321032120222001"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332013313221222-2111120103102303-0320233313013101-2332022200103002-0311313031310223-0032200003020203-2103030221203023-3111113302022101"></a>

## Direct properties — router / 301131103101 / 3

- [dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111): complete subsection reference.

<a id="canonical-3010001202013200-0231331333020133-1310210233202112-3102230101011033-1311131321130310-2212332011300302-2022301320030212-0333211312233200"></a>

<a id="canonical-3222131103013223-3133000323223303-1230220313103110-2110033232011211-1101131202302022-3232202020320023-1221030010221131-0311133011030222"></a>

## network_prefix property — router / 301131103101 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211): complete subsection reference.

<a id="canonical-2200233300113130-2321031300303303-0120303212120011-1222013133123033-0311200132321123-3130333103103202-3001020012112320-1002222332032020"></a>

## Next pages — router / 301131103101 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333201102032132-2002301130302011-2320301022301323-0113222313023232-3110133232331312-2102031302010010-3130321321211022-2023221312312300"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config — dns_config / 300313133022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-3021332123121111-3022312301011110-2301221001131200-0201313010233203-3121212331312120-1313303320322310-0333000222221323-3222123211203100"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302130123333232-2222000233321120-2101321222003222-1030022222223332-0320023000302212-2123033103012310-1320020330232211-0102133110002202"></a>

## Direct properties — dns_config / 300313133022 / 3

- [configured_list](resources--securemesh_site--reference--group-002.md#canonical-2331213301210120-1302322102032323-3123332121322133-1033122333222003-3332130121133003-2000101201301320-3001103122133001-1323220201132021): complete subsection reference.

- [local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230): complete subsection reference.

<a id="canonical-2213102223322333-0201221211011233-0322223311322330-2320203302021300-1202011012203130-2321010133233011-0330301222110211-2332113101313132"></a>

## Next pages — dns_config / 300313133022 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site--reference--group-002.md#canonical-2331213301210120-1302322102032323-3123332121322133-1033122333222003-3332130121133003-2000101201301320-3001103122133001-1323220201132021)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2331213301210120-1302322102032323-3123332121322133-1033122333222003-3332130121133003-2000101201301320-3001103122133001-1323220201132021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213110011100333-3222202312122010-3111020211220202-2022332020211220-3221110332213310-0000012333120331-0033310300132011-3230232120310133"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 203030132203 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3121033201310302-1032032112201122-0230231203131010-0121200001200322-1031132220223233-2300130021333132-1210301222323212-3220300303221232"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211221221210101-1132231111130201-0223313223233302-3333322010201201-0010012023302102-0001132101210012-3213211301111123-2120233112133211"></a>

## Direct properties — configured_list / 203030132203 / 3

<a id="canonical-2221202212202133-0230123100320113-1201123103130222-1212002203111210-1231301321312230-1223320320331003-0321120002211200-2332232123033011"></a>

<a id="canonical-0311010030120203-3222030221120113-2333120003332020-1332011332221001-1003322213100012-2213031201321303-3023013012112101-1333201120320000"></a>

## dns_list property — configured_list / 203030132203 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1121113102011222-2223230002333030-0210133102100322-1010133220000223-3020331303023300-0220030213330331-0301223030011130-1132033312023101"></a>

## Next pages — configured_list / 203030132203 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021210301310331-3203321210222312-3222312110223000-3233332330011210-3032030221302011-3130220002220020-0020021210322120-2301213330302211"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 312202012120 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2212211130133100-0212200111310312-1132221210322010-0232013230311020-1001202001321320-1330012130000313-3221030213203231-2310021111312311"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303331323300202-1103231213122023-1222223213302113-3030131310113320-2133022131103120-1102322101301222-3301103321012011-1001132111103303"></a>

## Direct properties — local_dns / 312202012120 / 3

<a id="canonical-2000312132020223-1003310323222320-1030013211111330-1230113101130232-1010001312113021-0002313021312313-0231003103021331-0313000031322133"></a>

<a id="canonical-1123012303013300-1201112100311022-3310220102320303-0022120021323311-1230301131211113-0002330222132211-2110110320233301-1130210122121232"></a>

## configured_address property — local_dns / 312202012120 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](resources--securemesh_site--reference--group-002.md#canonical-0011321203212233-0022123001323211-3021123213112231-0003231210322013-0023232103333322-0101321220220313-1012012302003212-0110310033103221): complete subsection reference.

- [last_address](resources--securemesh_site--reference--group-002.md#canonical-1022303230122311-1323130101212023-2220133113230013-3003001010313002-0233200100122023-3033032210333001-0103313020133131-0121211112010221): complete subsection reference.

<a id="canonical-3222312030131222-0020001101221023-1201122131221000-3013111323010011-2213302130200312-3022101033130031-0110230300233322-1312223020323321"></a>

## Next pages — local_dns / 312202012120 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site--reference--group-002.md#canonical-0011321203212233-0022123001323211-3021123213112231-0003231210322013-0023232103333322-0101321220220313-1012012302003212-0110310033103221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site--reference--group-002.md#canonical-1022303230122311-1323130101212023-2220133113230013-3003001010313002-0233200100122023-3033032210333001-0103313020133131-0121211112010221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0011321203212233-0022123001323211-3021123213112231-0003231210322013-0023232103333322-0101321220220313-1012012302003212-0110310033103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102203031213132-1211121133213101-1111333323303213-2221102212120111-2133312112010121-0131302110230302-0323102201213032-0002210032302011"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 302002320333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2110230330110231-2233131100101202-2213300101322102-1322330100203230-3310330111113000-0033031330131232-3332332022201302-0132102002220321"></a>

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
first_address = {}
```

<a id="canonical-3011021331130310-2111022121101202-0231220333113101-1103301001320330-2302120332032122-1231310333023003-1313323100130123-2203010013133320"></a>

## Direct properties — first_address / 302002320333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220021103322120-3013322203323110-2203011000232100-0022110310103001-0322210232232000-1333131232300202-3023330013031123-0110330201201330"></a>

## Next pages — first_address / 302002320333 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1022303230122311-1323130101212023-2220133113230013-3003001010313002-0233200100122023-3033032210333001-0103313020133131-0121211112010221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210330332103210-0000320132300303-0310332220113123-2033111300103330-1121123313001131-1033122233202320-3021100233123322-0311002300012101"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 200222311222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-3133311213310303-1112212323210122-1011001213222330-3021211033203101-0132111030312233-2300320301103123-2020001012003302-2010330310313111)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0321031132221123-0132231221330120-0030101322320221-2003031330300131-1222030003211131-2121231300013231-3311021333013221-1313300031121320"></a>

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
last_address = {}
```

<a id="canonical-1113303122203332-3112323130102110-2310133220221323-3231321022010311-2112021231123201-0211013130303320-0000102011221000-2332013120120023"></a>

## Direct properties — last_address / 200222311222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111101000122230-0203302201120100-0232131213132012-1103233230302301-1223321101023120-1013230230312213-2012323230201201-0132022213231332"></a>

## Next pages — last_address / 200222311222 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-3223120010113301-0231132011230021-0230332333311120-3313021220300110-2321212011001321-3032323321223201-1030220122320221-3023003222123230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321113123113110-1312003231303032-1332002203321312-1332221130113033-0320020333120231-3230311022012231-3313122010231201-1202130223010131"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful — stateful / 030303031331 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-0112113122003223-3232120223310300-2122331202030123-1120110020323212-0210202301003021-0203212201010232-1132012121131020-3103020200332320"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```
