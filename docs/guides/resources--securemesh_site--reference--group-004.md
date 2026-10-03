---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-3112220203321222-1232112020330003-2323202311023111-3110000230313131-3201212312333323-2103320012030010-3220332330231233-1032302023220023"></a>

## Direct properties — static_routes / 033122002300 / 3

- [static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303): complete subsection reference.

<a id="canonical-2032133110213310-1331320032030230-2132213123332331-2030202220230030-2033003022311111-3221311032332320-1112302120120301-2131311313332121"></a>

## Next pages — static_routes / 033122002300 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310001302213113-1303302311110131-0211332201330233-0032003202021313-1310331331020112-1101213010121103-2001002303103122-1220311033031202"></a>

## custom_network_config.slo_config.static_routes.static_routes — static_routes / 033331021133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-3300301301210202-2310233101311223-0321103200232031-2133022300120131-1313020203311123-2330323022113303-0202210033020100-1310203330202000"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120222210202320-1321301221130100-2033212100203031-0231211111323102-3000100212203133-3222113203100320-0201220200223031-2322232002213110"></a>

## Direct properties — static_routes / 033331021133 / 3

<a id="canonical-3330300221000233-2333212101202312-3113230120313133-2201323121322102-1202111300302303-1322231223132103-0102103010331210-0332010120012230"></a>

<a id="canonical-0303311100331120-2212133213030131-0222032123021133-1222013332323001-3131132301011011-1212120221303020-3022312000313332-3003133321330120"></a>

## attrs property — static_routes / 033331021133 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site--reference--group-004.md#canonical-0232033302003131-2133200320202210-3231031330323300-3102112302103022-0131031001300231-1032003203023131-3103003012202200-3023012310323122): complete subsection reference.

<a id="canonical-2133113301212101-2020223000330200-1202223100301122-0012001310023033-3000123002220212-1320220330131303-0323233132030133-2323322300222110"></a>

<a id="canonical-2032002203203332-0030201111030331-0021012030303123-3311111123230333-2001210210002102-2011020103021301-0201030021131301-2111231202313331"></a>

## ip_address property — static_routes / 033331021133 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-3331020132133122-1201131023321200-2020212311022112-2222130033331311-3032311003010031-0123210000103101-0323033123000032-3301101231032113"></a>

<a id="canonical-2331003003003030-0302131213310233-2302130111322130-1031001312210001-3101121222221133-2010220310031031-1231032232231330-3300023333100232"></a>

## ip_prefixes property — static_routes / 033331021133 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site--reference--group-004.md#canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002): complete subsection reference.

<a id="canonical-1123321023101321-1312200020010302-1032122221301213-0322111202313313-0033021301310213-1022231110020010-0020121220102122-1321222211230021"></a>

## Next pages — static_routes / 033331021133 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-0232033302003131-2133200320202210-3231031330323300-3102112302103022-0131031001300231-1032003203023131-3103003012202200-3023012310323122)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0232033302003131-2133200320202210-3231031330323300-3102112302103022-0131031001300231-1032003203023131-3103003012202200-3023012310323122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111222231030032-0023223300231132-1223311203003012-2132302133032010-0322302220011123-1101120021022331-0333121301123001-0313121030020302"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — default_gateway / 212011022301 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-2331333102320030-0331112222021011-3102023033130030-2302321130201123-3100112002032322-2200002211300233-3311102321311121-1221021023230111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-1233103023103023-0302310103320011-2210212120201311-0003331211020212-2301110222003221-2311021231030132-2120122023023321-2133321232020331"></a>

## Direct properties — default_gateway / 212011022301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112203130000211-0333210232113101-3011330200021130-1233303022123132-1111200122300210-2031133213302133-0230323012022001-3001001101332023"></a>

## Next pages — default_gateway / 212011022301 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200000330130210-3030020222231202-2221300323213332-0211022301332233-3010122222103223-3320201220322220-3122101001331103-3100220221221100"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — node_interface / 132023202112 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-1031030230221131-0110111022003123-2030003113212320-1221232301102330-1321300020332223-2303321301033310-1122201010000200-3220102032323032"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233113302202230-3322202103320330-2212210213122023-3022133302020111-0210003112021331-3223231310112122-3023132012333201-0122331000013022"></a>

## Direct properties — node_interface / 132023202112 / 3

- [list](resources--securemesh_site--reference--group-004.md#canonical-3223120221033220-1111210320002212-0212030022330133-3032123320313131-1132001102000310-1320232332231020-1200213121110300-3212021303003231): complete subsection reference.

<a id="canonical-2103312221320331-3031120303132303-3122121220010232-2322122023031010-1030032322303202-0021023310012003-2103201311233023-0020002032012210"></a>

## Next pages — node_interface / 132023202112 / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-3223120221033220-1111210320002212-0212030022330133-3032123320313131-1132001102000310-1320232332231020-1200213121110300-3212021303003231)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3223120221033220-1111210320002212-0212030022330133-3032123320313131-1132001102000310-1320232332231020-1200213121110300-3212021303003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221230302300320-2132232121031233-1031133300301303-1310102330211220-0012010323212031-2203210110303001-2012312110011322-3310302213033002"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — list / 222210133121 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-3131312112320111-3111310011003123-3313010322303233-3323213012031320-0013232100233002-0112311010000021-3031311100202312-0012222011312211"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121011011132130-3010231200101230-3300211233222202-3321230113002301-1301120313331020-1230333033313132-2113123200132113-1101131030321223"></a>

## Direct properties — list / 222210133121 / 3

- [interface](resources--securemesh_site--reference--group-004.md#canonical-3102303202101133-2312032213002220-1103223100012022-2012212131011110-0313201331012020-1320013122233201-0322021302033330-2333022232323033): complete subsection reference.

<a id="canonical-3000010130330231-2321231311010322-0001200302211231-3033023321010300-3333012310331010-1033321311310031-2132113220022130-1011010033320111"></a>

<a id="canonical-1131232231121130-1313213130321102-3110310303112223-0312221303311320-3000033302110200-3220221301202112-0331030013322010-3030133302113302"></a>

## node property — list / 222210133121 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-0221033300321320-0032301321030212-1313200310013301-0100100130133312-0100321110203202-1033130030100212-0221320331120303-1322302310330020"></a>

## Next pages — list / 222210133121 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-3102303202101133-2312032213002220-1103223100012022-2012212131011110-0313201331012020-1320013122233201-0322021302033330-2333022232323033)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3102303202101133-2312032213002220-1103223100012022-2012212131011110-0313201331012020-1320013122233201-0322021302033330-2333022232323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102101213110132-0223233030223022-2000230222232321-1332121110001100-0113023313321210-0002313011012230-2030332132211311-1002121301021203"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 220001323231 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-2203332001031330-3033332130333023-1301201331130203-2303110021031222-2113002021131110-0313200302013213-3210211202022311-2220330321000303)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-0211221131320211-0232321300003031-0323010010030222-1331033010233230-1013221033211303-3211222111103322-3201301012122022-3221333303301002)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-3223120221033220-1111210320002212-0212030022330133-3032123320313131-1132001102000310-1320232332231020-1200213121110300-3212021303003231)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3233113213331302-3123002131133100-3121322212300132-3010033321301012-2202222000221102-2201110031333302-3323013303010331-1213230223133322"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132112123310231-3121311111303201-2323331011323310-2120023330300132-2331211303122310-2311303332110330-2203313332333121-1001210021122110"></a>

## Direct properties — interface / 220001323231 / 3

<a id="canonical-2311232103110012-1132133330230023-3021032001122031-3021132330332333-0202012213312220-2011232122023311-2302221312211210-0101001312312132"></a>

<a id="canonical-0031212121021223-3133331113222122-0032001330100322-2332022121313111-3320231030002100-1213203323111130-3322233133222302-1303023113202220"></a>

## kind property — interface / 220001323231 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0132111203001133-3001303200233332-1013123332330131-2132221003010131-3131302012001320-1322131031030321-2003312301033112-2113020131033010"></a>

<a id="canonical-2013323032300131-1000110331221231-0300013030103003-0031303011100210-0333023013230031-2220201011312323-0033201012220020-3110333303003221"></a>

## name property — interface / 220001323231 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3223303102111101-1001103211211322-3001212023223022-0010312211233132-0030223033230002-1132020010321322-0032101212200111-3010100030200220"></a>

<a id="canonical-0323201112331113-1001010023212211-1113123223010103-1130002202112323-1020011333113121-3011110013321321-2311302232213031-3000323320100012"></a>

## namespace property — interface / 220001323231 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-0002331113211211-1111023300231012-1311322301122212-1223231001331120-2202210023302002-0300222003333330-3232322103311010-1013220333002012"></a>

<a id="canonical-0010132030220132-3300300223231012-2330203002313012-0333201102212221-0313213301121120-2203130101303103-1201013132132112-2123222132102203"></a>

## tenant property — interface / 220001323231 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3303233311313003-2322113120232131-1113021101213031-2011000133321210-1303210333210121-0211231222120122-2033122312032222-1001101123131200"></a>

<a id="canonical-3220333313203210-3121023333021020-2331131030221323-0202211200333230-2130002302323202-2320303012313121-3130320303132312-1300012003212123"></a>

## uid property — interface / 220001323231 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2222003110021023-1100212301003302-1001230200003222-3001000221030320-1310230111012000-2200211302120303-1220132320023231-3231010303011013"></a>

## Next pages — interface / 220001323231 / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-3223120221033220-1111210320002212-0212030022330133-3032123320313131-1132001102000310-1320232332231020-1200213121110300-3212021303003231)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102332203112333-2121110213211013-1033321133023022-0032122001022103-2213110203201322-1220011121223311-3031322120023113-1130323111121212"></a>

## custom_network_config.slo_config.static_v6_routes — static_v6_routes / 203110113330 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-2221001023013331-3121023202313101-0230313310210103-0332123220003022-1122202212130132-3232121131333032-3202213213333123-1302112330322131"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313032311131020-3211322011100210-1321103222132130-3303211321000000-3020132323231003-3130321221132112-3312301312230111-2132312020333123"></a>

## Direct properties — static_v6_routes / 203110113330 / 3

- [static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321): complete subsection reference.

<a id="canonical-2021333220232313-0103210011222203-1032030120112323-2310202030310021-3323031213322103-1310200030321110-0310022222133220-1112010013133030"></a>

## Next pages — static_v6_routes / 203110113330 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221030313023232-2210112123102120-2213230002031233-1021302330122033-0203303310203031-2102013210303112-1132310102233110-3311313101120100"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — static_routes / 112002031311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-0332000210122313-1210231033230200-1002002221010202-0201030111300330-2011122231231331-3221301113120011-1020103012230013-1131233223310031"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202101100233112-3002000031213311-3110222121012032-2013031131133320-1320121100010110-3120011013121122-2221210000112022-3310002013111221"></a>

## Direct properties — static_routes / 112002031311 / 3

<a id="canonical-3210222332312121-3033222321320301-1121021022310133-2033132211102121-3000321003200032-1301030222032130-3233302100210022-0000100221003131"></a>

<a id="canonical-3223012320301113-1123330201133013-2231203002033101-2221102102111021-2212303012023100-2233032332321230-3300230322220131-3012303111113221"></a>

## attrs property — static_routes / 112002031311 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site--reference--group-004.md#canonical-1201332031011132-3213020031323320-3031102302202233-3321030321232110-3121002313110002-2120212103020113-3310001030102011-3021200102302312): complete subsection reference.

<a id="canonical-2230230033300310-2332321211203211-0010123200311132-0212110033010001-2023333232231110-1220113310303322-3120133302131331-1111111301112010"></a>

<a id="canonical-2121122021002212-1200112333213323-3311320112110100-1212300122010331-0122221023232011-1313113031320200-0112020010232113-2220322011123221"></a>

## ip_address property — static_routes / 112002031311 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-2310323310302111-1003030223122331-2120200111211303-0313223010231310-3101113210223122-2031300323220331-3030301222303033-2321301021021002"></a>

<a id="canonical-0323021200102222-1233031031223322-3232012133322032-2303212022122113-3020113102230131-0200230033032030-1000210201121232-2210110210212323"></a>

## ip_prefixes property — static_routes / 112002031311 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site--reference--group-004.md#canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010): complete subsection reference.

<a id="canonical-1200300123312230-3311333122222000-1100332032203101-3102221032110212-3130232020201010-2201131211021333-0131313223012332-3201121330221021"></a>

## Next pages — static_routes / 112002031311 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-1201332031011132-3213020031323320-3031102302202233-3321030321232110-3121002313110002-2120212103020113-3310001030102011-3021200102302312)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1201332031011132-3213020031323320-3031102302202233-3321030321232110-3121002313110002-2120212103020113-3310001030102011-3021200102302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101122000203323-0300113232030220-3323222000113332-3301200302320120-3333203232120100-1223031321120030-1233213103103113-0331020201330332"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 123220201131 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1300233333331023-2210303110323123-2221120202302012-1111021123103212-3002231100123113-2322203333130032-1333201220212022-2131232330022133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-0011301333000231-3121332020123122-0212210013320121-3223221220023033-2231212221023123-3102023011020133-0033013020311202-0100221003130120"></a>

## Direct properties — default_gateway / 123220201131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231222013122010-0121020013320022-3323220123113333-0023121023131310-2302111311122231-0203210213223111-3031130321031322-1313301013333201"></a>

## Next pages — default_gateway / 123220201131 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102233121031022-3112001133133031-3133121332230023-3210322313110110-2010311123333220-2333123022313022-0230330233102103-2301300112020320"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 223330220112 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-0301013031230013-0221300011223232-3131212002000132-1201023100313300-0311132222300213-0211311122223032-2033333201012032-1333032123002020"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232021013322210-0100130300203023-1131102200111021-1121331233001001-2122003110232332-2033311001302312-0101001020233222-1012302121320201"></a>

## Direct properties — node_interface / 223330220112 / 3

- [list](resources--securemesh_site--reference--group-004.md#canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313): complete subsection reference.

<a id="canonical-1233131300120101-3332201032023103-3323012233013013-0011111213322230-1002011300102203-2033210022123212-0313020322010303-0010013302121020"></a>

## Next pages — node_interface / 223330220112 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010301223032031-1221201312121023-3012132031010102-3230100130010322-3333131321033311-2320313033302003-3231113222310131-2130322021131300"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — list / 301222022021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2213002010113230-1203011133313113-3212103123002003-3120322210230110-1122122030233120-3322123101330023-1300130231131031-3012201101330122"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303133210322111-0201010003321210-0222313202221111-1031112012102021-3032131022011230-2010222222320230-2333121300213113-3111101213302003"></a>

## Direct properties — list / 301222022021 / 3

- [interface](resources--securemesh_site--reference--group-004.md#canonical-3110021100113222-3103320311103213-0021213302231000-1303323320230323-1112212223123131-3333200120000100-3100333330130230-3103300303223230): complete subsection reference.

<a id="canonical-1023332213033312-0300121113102002-3013001331000132-2221101312331022-3201310200223033-3013203021120202-1131121022202033-2222101132210113"></a>

<a id="canonical-1321120021110113-3313103321331133-2121102113303221-3302100110303320-2000130233202321-0211331033022230-1203230112031310-3003223101011033"></a>

## node property — list / 301222022021 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3012333313321320-3311233223031303-3210122333223203-3323321311300111-3220001030022222-1023103223132112-1122322020213111-0011122100013233"></a>

## Next pages — list / 301222022021 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-3110021100113222-3103320311103213-0021213302231000-1303323320230323-1112212223123131-3333200120000100-3100333330130230-3103300303223230)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3110021100113222-3103320311103213-0021213302231000-1303323320230323-1112212223123131-3333200120000100-3100333330130230-3103300303223230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203102002033233-0020113330231331-0112133331102232-1130033020301112-3330010100103103-3110302221133301-0332220002100013-0231010102322123"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 123012001303 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3013203132332332-2330212203030010-1103200023030012-1303130101302031-3203012211233231-3200003033013122-0030030001323200-2010201133210321)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-1030121023213012-0011123003033330-0122213113330332-3002221310213013-0020133013300122-0112233233232202-0203320331020100-1310203123032010)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2212130010022032-3313120010132213-2223021220132202-1201332332112202-1132321302211331-3100103122313131-0110120330120131-0133311211303100"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003221123332220-3100302221000321-0121112313212221-3122311030130100-1302320322312000-2012012223302011-2102131332213112-0332301000022010"></a>

## Direct properties — interface / 123012001303 / 3

<a id="canonical-3331213020100011-2302212000003022-2102331103211032-0113113300031202-2321121131132210-2330133130012103-0100200032012210-1210000202221123"></a>

<a id="canonical-0202313103202102-2032121110023021-3101111201002131-0221031031131130-0101230031122130-3103302112003110-2121323210102122-1333013201301031"></a>

## kind property — interface / 123012001303 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3103231232113200-2231100200313311-1030123010011203-3013033032010312-1011131023102102-2300332130333301-1310100010233111-0320130311003021"></a>

<a id="canonical-1101013030102331-1312301303210222-2020103311323103-2022001200123301-3121122120201101-2333213131121030-0213023100110011-1130032222223231"></a>

## name property — interface / 123012001303 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1223121103312310-3201223100320102-0201321113133330-2232133212202221-3330030211330202-1301121132032322-3223111112030320-0010110033132020"></a>

<a id="canonical-2321200100333032-3221202112022112-0111112330123013-2112122013313311-1022002211330020-2330211032233212-2112102312300003-2200221212220322"></a>

## namespace property — interface / 123012001303 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-2102212033103322-0323312321333333-0321222313010330-2321203220101322-2021310111300133-3110301131003210-2300320110013213-3120223220220221"></a>

<a id="canonical-1200012102102113-3133200111300230-1320011010111321-0021322030012220-3212132302303022-0122311321332132-1202131322013212-2131332121331200"></a>

## tenant property — interface / 123012001303 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3123222123313223-2302021031111220-0320213320230332-2013101211220311-3000200120312220-2330320101201033-1211303333130033-0110121213300333"></a>

<a id="canonical-0100100103122301-0202003130300233-2123230113313302-2313103112010022-2223002221313200-1221130122231130-3003110020211011-0301012101312213"></a>

## uid property — interface / 123012001303 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3200031113232210-3233202103330131-0113231321302212-3000112123020201-2332123211313301-1302231011332331-3021301130321111-2012123010122132"></a>

## Next pages — interface / 123012001303 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2131101313011023-3202103200013000-3323310203031110-2210320233112330-0202233231312011-1312202132031201-0231331001012033-3000122001001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310213222023232-3233130133022000-2100121130330221-0312130012303002-1313220033003033-2321300132000001-3323310213233110-1130110112102302"></a>

## custom_network_config.sm_connection_public_ip — sm_connection_public_ip / 301212110303 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.sm_connection_public_ip

<a id="canonical-3300333211213131-2330032100010033-3120000303101223-0012220031333100-1222320113113213-2332001101001011-1231321201333110-2303211031333323"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-0120010030232300-2310100122310121-2033111231120300-3310311231102321-2121333213012321-2303000132312121-2220220221230320-2320222311011331"></a>

## Direct properties — sm_connection_public_ip / 301212110303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311321222123211-0012230123332202-0201113323113033-0102201303321123-1200000121323221-2322313231021012-3112202032130032-0223301212011112"></a>

## Next pages — sm_connection_public_ip / 301212110303 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2201302222100233-1133301323301033-1023210103301301-1330322123030131-2202111223230332-1121203213021131-1100102123030200-2000030033123120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220203000032311-1233111020220223-3120303033131211-1112211331022320-3333311231132000-2203123221121302-1200201223022010-1321120130031331"></a>

## custom_network_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 210000002333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-3121010132212233-2123300123202210-1313310001122332-2113313002320223-0013312211213301-3001230112122220-0033233133123200-3303203012202310"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-2322331102120300-2212223021132133-3131201222033012-0023210013211232-3110200221101132-3012200313130302-3011213111013100-2221132110202230"></a>

## Direct properties — sm_connection_pvt_ip / 210000002333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302313010230231-0213203202122222-1130022321230323-0023122011012213-2230220211200321-1230012110000230-0221003131120202-2010101301213120"></a>

## Next pages — sm_connection_pvt_ip / 210000002333 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2130322011122113-2023213101220320-0310300323120322-1213333313311002-3123113231010313-0110102302312320-1200031302010203-0311320121030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001011130023130-2101101222033002-0322111130000121-2212301011331013-3220122321132233-0301120332110013-3121321001122020-0111132231331232"></a>

## default_blocked_services — default_blocked_services / 313212310311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- default_blocked_services

<a id="canonical-1303223332233202-3102111200020010-1101221113320021-3021211312123002-0022331320030330-2201001130120012-2332202001033010-0222100120332130"></a>

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
default_blocked_services = {}
```

<a id="canonical-0003233331231012-0332322113320031-2120011201323211-1002320032001221-2012312203203301-1323201222312223-1013121100031331-0130023111110101"></a>

## Direct properties — default_blocked_services / 313212310311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220200310212121-1032201232103322-2200001010201232-1320002231333320-3030230113010300-1121303123111101-1203212032101113-2222021212312310"></a>

## Next pages — default_blocked_services / 313212310311 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1211212023033021-3013201123022023-0023321103121133-1222322023300323-0131132032303333-3332330303301122-2032220120223122-2021102112323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103321022303300-2231122103201133-2133121103132212-1132033110213032-2013211203023232-0331331112120121-0010230012312012-3030233313323330"></a>

## default_network_config — default_network_config / 312001332100 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- default_network_config

<a id="canonical-2013013131311232-2133302221020332-3112133011223112-1223202233333032-3300330102101202-3320111211322111-0320131312220332-1201332032332311"></a>

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
default_network_config = {}
```

<a id="canonical-3332311310012013-2021220031203210-1321321002130000-3013123300013033-2302102320103220-0033230303021321-2110333231321321-3303021223202321"></a>

## Direct properties — default_network_config / 312001332100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000303121211000-2232313001220111-1003001022300122-3322022202310002-1202212013012331-3001201232123321-2322203130020020-1323201301301212"></a>

## Next pages — default_network_config / 312001332100 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312111200230130-0232231112021031-1002132032230020-2133010032010230-3202322130310321-0221221213201110-0302200333002312-0020130012222022"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 132112123322 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- kubernetes_upgrade_drain

<a id="canonical-1200220010220323-1103231332321032-2113132203330200-2203233303031201-0212130320011230-3320320203323220-0033002002202010-1022113332333022"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222213013132012-2132113021032132-2310001030201111-1001012311111012-2220213121121002-0202011223212200-0333100030211331-3313112012212111"></a>

## Direct properties — kubernetes_upgrade_drain / 132112123322 / 3

- [disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-0223232203330201-2013032230212300-0231011120000322-1021330320311013-2010123030102020-2103202211102133-1123212010103103-0322113123013130): complete subsection reference.

- [enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010): complete subsection reference.

<a id="canonical-3330232030322332-3112002102011101-3231103121000320-2110202031112233-0113020013302131-2312033323302120-3233110001331021-0211012023320210"></a>

## Next pages — kubernetes_upgrade_drain / 132112123322 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-0223232203330201-2013032230212300-0231011120000322-1021330320311013-2010123030102020-2103202211102133-1123212010103103-0322113123013130)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0223232203330201-2013032230212300-0231011120000322-1021330320311013-2010123030102020-2103202211102133-1123212010103103-0322113123013130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032032311121233-1022332231331130-3223103302011011-3133231301312132-0021001123122332-2132220120132103-2322120233011312-2322131312101213"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 312002323010 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-1101033033232300-1323000102331313-3210313112302012-2210230210313321-0003212132322001-0100103231302000-3313213000111020-0001012032112330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-1020210332121301-3203020100232332-3302012123031311-2313223100100230-0002131030013330-2330030300222311-1011322210323232-0021022002120333"></a>

## Direct properties — disable_upgrade_drain / 312002323010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323222113330302-3223322233013301-0320011022002232-0100310311321130-1223031331111330-3010220131230011-0100312122212123-0020312231021332"></a>

## Next pages — disable_upgrade_drain / 312002323010 / 4

- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221131012020333-3332131102022123-2202201222103112-3300012332113033-2131200333231222-1230011323313100-0023032201213303-2333123132322113"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 013313233031 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-0002223213033113-2113031212023303-1000232021300001-3033001100132300-1023231213231101-3032301301132321-2100020323313013-3003122211120320"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221011110221111-0102003211200231-1213213221133221-0203310132020320-3303333022122101-2133131022203333-1000320121030332-1212031101001033"></a>

## Direct properties — enable_upgrade_drain / 013313233031 / 3

- [disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-1313001311311032-0132123221013301-2010212331111203-2331130023011023-1310003333021031-1333003221110112-2021102121002103-2130330001121110): complete subsection reference.

<a id="canonical-3330311032310010-0003122302023322-1121030031100022-1210123012203332-1232211103320233-0200221000202003-2233010131213101-2121303023111211"></a>

<a id="canonical-0120033133131321-0220101012301010-3130011333332011-3001001321300223-0013131333323311-1002313001100212-0012103322312223-1030013302311032"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 013313233031 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2203313233203013-1211333112112221-2322002102210130-1210312123313111-2132232130231330-3021023302311212-1232022301122211-0221310303332212"></a>

<a id="canonical-2213303333301122-2201313313303230-0310212321100131-1001221031031003-2131321211012303-1102231232020031-0023101203222010-2211010121303230"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 013313233031 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-0112333212002022-0213230212102002-3221303102222032-3221131330300033-0132021033201011-2021011001230002-1202100321311101-2211213321321121"></a>

<a id="canonical-1301102333300020-0023213330312123-2001131120301233-3031112231132010-3013012313021310-1001121223120032-3203310002023111-3322110012010111"></a>

## drain_node_timeout property — enable_upgrade_drain / 013313233031 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-2113122010121131-2302003212300102-1201123033023300-1311312212211000-2122021110211011-1101112202112100-1100302301001321-3332113002300031): complete subsection reference.

<a id="canonical-0000302200010230-2030123222020232-3120310003303130-3020021101222331-3233033100001111-1101022323022022-1220102302321311-2212232220200003"></a>

## Next pages — enable_upgrade_drain / 013313233031 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-1313001311311032-0132123221013301-2010212331111203-2331130023011023-1310003333021031-1333003221110112-2021102121002103-2130330001121110)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-2113122010121131-2302003212300102-1201123033023300-1311312212211000-2122021110211011-1101112202112100-1100302301001321-3332113002300031)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1313001311311032-0132123221013301-2010212331111203-2331130023011023-1310003333021031-1333003221110112-2021102121002103-2130330001121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313332033221103-2322210132121233-1323020222330121-1231100010231321-3300313112312001-2100121130021031-1333123113230102-1201222231012112"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 310223301011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-1222031100203321-1211301220131003-1211320313003011-3323013202223020-3230120022121111-1012320010222110-1002222202110001-0320332233103013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-1000013131212031-1313130202220323-2232321031312023-1010121021003311-3300121030113033-1110130031122122-3131122320010103-3331130213202113"></a>

## Direct properties — disable_vega_upgrade_mode / 310223301011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233332001230111-1313312323011023-3233032013030303-0122333103010300-3333220323122213-0233112311220120-0102123232003110-3332321333001021"></a>

## Next pages — disable_vega_upgrade_mode / 310223301011 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2113122010121131-2302003212300102-1201123033023300-1311312212211000-2122021110211011-1101112202112100-1100302301001321-3332113002300031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130130300313331-2332013301010231-2232123020200013-3101321120220023-0110030233221330-0333112102223100-0101220331322013-2211230330232310"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 133212003112 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2221320002201333-3210113123303202-3321133112213131-2301312302133332-3003203302020102-3221001211312331-2212131033302133-2213222231311221)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-3131300113101223-2112212220111302-3203030302233303-1210310113233133-3130320210223010-2031030323201331-3203103131130002-2102032011100120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-1233210200103001-3200202321223333-1300111022333130-0220311322212031-1201310133200132-0001132320030230-2100131002323302-0223303031222113"></a>

## Direct properties — enable_vega_upgrade_mode / 133212003112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012030232132330-0222031200230330-3330301311113130-1133313011030020-2320101231011312-0312322122312023-1130022211112333-3331300311002303"></a>

## Next pages — enable_vega_upgrade_mode / 133212003112 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-3122301233003211-3133331001103010-2201322213022001-0212011123320133-0000310120113200-1320223310332200-1103220213322133-1302113211232010)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1210231003210032-0210302031033031-1321120333323332-0113030213300300-3003201313113111-2321333000300322-1312122112223131-2303302010130323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131000001121232-3322100323333322-1230011223120200-2011023332011202-1111021221213221-2311322120311312-3300312230130030-3133212212101113"></a>

## log_receiver — log_receiver / 333123033101 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- log_receiver

<a id="canonical-3003100221222100-0203310301302221-0311123203310230-3223013221302213-0203133020303031-3121000331221231-2332200211020200-1130233321233032"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-3003100221222100-0203310301302221-0311123203310230-3223013221302213-0203133020303031-3121000331221231-2332200211020200-1130233321233032)
- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-0122012011201212-3003010312303213-0323233010133011-1321302030301332-3321311103110011-3210001001022311-3203201331001001-1120311203020313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011323033313300-3100101220302101-0211323222323113-3111113232331302-2000213113032230-1121320232002201-0212312300300201-0120313011030303"></a>

## Direct properties — log_receiver / 333123033101 / 3

<a id="canonical-2020000112203102-0130023310330131-0232221103030332-3023013033321023-0330112220221112-3313113320300022-1321023100012313-1010001300312030"></a>

<a id="canonical-3222123313210001-0301003301010111-3110221333021032-1202122120233023-3302022112333300-2100012121103112-2113213202213333-1032233001200032"></a>

## name property — log_receiver / 333123033101 / 4

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

<a id="canonical-1011010000311020-0033222332122130-3022311221301212-1111023310021210-0113033112303020-2123210201122313-1031031110002003-1133101031133021"></a>

<a id="canonical-2003033321100212-3121200010211301-1020231301212202-3212000020100012-1311201203003113-2330223031332000-3213213312012330-0223321222213312"></a>

## namespace property — log_receiver / 333123033101 / 5

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

<a id="canonical-1022103111022303-0302310103121232-1321323030210222-2323122101320211-1301323332132300-1200231110231101-2122202122010033-0121023203031322"></a>

<a id="canonical-0211203103301111-2120120222131000-1022021031303001-1012230001233022-0111112221222312-3130121022111133-2023213013222113-0013111333231100"></a>

## tenant property — log_receiver / 333123033101 / 6

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

<a id="canonical-2030203013200331-1301022211233323-2210302002233213-2111312001322320-0113320103320323-3323311123020313-2022013032213210-2100320322001331"></a>

## Next pages — log_receiver / 333123033101 / 7

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0301322311211110-3313100231331120-3020131300021210-0031233031211302-0223101202102010-3322120202121023-2112032200010220-2203211233300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211101121200223-2210100310300200-1313232102022010-0301013120223121-1230121301203220-0101100331110330-0001320133313032-0002230130113311"></a>

## logs_streaming_disabled — logs_streaming_disabled / 031120202103 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- logs_streaming_disabled

<a id="canonical-0122012011201212-3003010312303213-0323233010133011-1321302030301332-3321311103110011-3210001001022311-3203201331001001-1120311203020313"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-3312032010131113-2111321133212102-3330132313120320-3020302301323030-3032223200213322-3131302302210021-2330122030012301-2213301321203022"></a>

## Direct properties — logs_streaming_disabled / 031120202103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221200111000302-0010302001212120-3333012212121300-0120321332101303-1011003302023232-3003202001221031-2221202331211221-1202320020331013"></a>

## Next pages — logs_streaming_disabled / 031120202103 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2030021313103121-0101111110133002-2130111211310100-2020102210122203-3100100030322221-2033211312031111-3022033203312230-3201003223312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211331130012130-3221022102121132-2133011120011323-3030333011122033-2103101112102333-1323102023111113-3321323122121311-0333101322302112"></a>

## master_node_configuration — master_node_configuration / 002213322111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- master_node_configuration

<a id="canonical-2130211113130203-0203010311303120-2213000303013002-0222132121111333-0330131130030031-0122010020121120-1320203210233323-3332233211022331"></a>

Type: `"object"`. list nested block, Optional.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
master_node_configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022130311022012-1000012303332212-1222232122113330-0333313210001022-0211122121203330-2023321012120210-1211032122302222-2333213323320021"></a>

## Direct properties — master_node_configuration / 002213322111 / 3

<a id="canonical-2012133132133233-3122030223223023-0102301031103232-0322112023000000-2030130302312012-1210310331312001-0102132132121132-2100012102112033"></a>

<a id="canonical-2302331002320202-0220301301211233-2333112000203003-2221212113321332-3330302112233032-2322330333331231-0102031002103331-2212323022203212"></a>

## name property — master_node_configuration / 002213322111 / 4

Type: `"string"`. Optional.

Name. Names of master node.

Upstream description:

Names of master node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-0031102331331311-1112021130321120-0021313130110032-2310232122032313-3023123210013131-1331330000123101-2023033001320320-1010320100011221"></a>

<a id="canonical-2101130111001101-0321330000122033-1320131332302223-0032230232011000-2013323010223322-0210312011332031-0100333323322100-2003301123322323"></a>

## public_ip property — master_node_configuration / 002213322111 / 5

Type: `"string"`. Optional.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

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

<a id="canonical-3202013100122123-3231013033120111-1210132303133231-2202102033323102-0103111303212332-1022000010130302-1131121012201131-1030331013132103"></a>

## Next pages — master_node_configuration / 002213322111 / 6

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0333112101031113-3121130300222310-2130011201332013-2012000233233023-2231013313212021-1321312222230212-0323220201313111-3200133332031330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000202033033222-2231200311233033-2323030220103201-1233233032103013-1220131130200022-1330033132030010-0222103110133311-1003102113212021"></a>

## no_bond_devices — no_bond_devices / 232211000013 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- no_bond_devices

<a id="canonical-0113102101220123-2031110131133030-1302213332233013-2303211033233232-3233310230201330-3110102232321022-3310032031233310-1231321301312313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no bond devices.

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
no_bond_devices = {}
```

<a id="canonical-3121333133000020-0232323002010333-3103111123202302-3023011002220210-2010012130302003-0321222133210333-2101202020021322-1110123232322233"></a>

## Direct properties — no_bond_devices / 232211000013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023213022023033-0313220202121221-3033212331333101-3322020123033222-0303211232122303-2300011223322203-3203010220233020-1301212320321103"></a>

## Next pages — no_bond_devices / 232211000013 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131321200222312-3232103333223111-1000012321022010-0322122302213223-1020232120201033-2103020020331302-1102332230001231-0331001133132111"></a>

## offline_survivability_mode — offline_survivability_mode / 100332020330 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- offline_survivability_mode

<a id="canonical-0122021021232023-1111302211230021-2111220301121202-0200123220133110-0110232302312300-0020212332331122-3103200101012013-0132233322012210"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310023013122231-0212211113331132-1012200011121223-2330132002303001-3011301112132021-3322101010020000-0312000132211133-1222010122103201"></a>

## Direct properties — offline_survivability_mode / 100332020330 / 3

- [enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2202322322200002-3320103121002200-2213133230032333-0000201131311203-3232002222033202-2031133002321330-2212220232130123-0232012210202202): complete subsection reference.

- [no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-0001102322323311-3222023101021232-1201113331231302-1130200003010230-1301330320021221-1112213302233331-1111212300003321-3213002103301020): complete subsection reference.

<a id="canonical-0132002030110110-3021030000132310-2030332201101223-1133312221010232-3112300030133233-0121311200303101-3130233320020021-2010200002233232"></a>

## Next pages — offline_survivability_mode / 100332020330 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2202322322200002-3320103121002200-2213133230032333-0000201131311203-3232002222033202-2031133002321330-2212220232130123-0232012210202202)
- [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-0001102322323311-3222023101021232-1201113331231302-1130200003010230-1301330320021221-1112213302233331-1111212300003321-3213002103301020)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2202322322200002-3320103121002200-2213133230032333-0000201131311203-3232002222033202-2031133002321330-2212220232130123-0232012210202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231232130201023-1223333222010210-1130033123301311-0231212233121302-3322011033202022-1110100201012303-2010022320210001-1021122101031113"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 221220030311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-3100103200221120-3001022213312012-1313120222111213-1230322112023313-3131220211132200-0011330321020331-2111222130201330-0122100320013123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-3231133133022112-2120131333111331-0220102210123320-2322101121113313-2033123220001203-1032232310311133-3323230301211101-0223111131202310"></a>

## Direct properties — enable_offline_survivability_mode / 221220030311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332213210030112-2211220101012211-2320003203211111-0200002123321221-0010012023202211-0002020031122321-1101322301100310-1323000231023330"></a>

## Next pages — enable_offline_survivability_mode / 221220030311 / 4

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0001102322323311-3222023101021232-1201113331231302-1130200003010230-1301330320021221-1112213302233331-1111212300003321-3213002103301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213223103331112-0031012030131211-0033123200201003-3223123110331030-1033320200201231-1223102333032122-1120323213210302-3021023313312103"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 213112011030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-3113330332210203-0122313200031100-1111113332223312-2211000201133312-0032212012332210-3102323331211322-1031103331301131-0113122023311200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-0331100103212313-0212202220301230-3012033001130021-0023010220203031-2200212010320223-3330222012221110-0000212003203030-3232032223021111"></a>

## Direct properties — no_offline_survivability_mode / 213112011030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012323222212220-0301021222332331-2000323300303133-3120130030210203-3003000013202331-3201333121020200-1323122223123321-1223310311113201"></a>

## Next pages — no_offline_survivability_mode / 213112011030 / 4

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-2020131231310113-0023113332222012-2232213233101302-3331333301203331-3301311130123013-1000321231200102-3313312022231113-1101200101031100)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0102133232320131-2231010312223023-0203311201032131-2210230021210322-0310110201031312-2321113100212222-0201202312033330-3000120030223322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120323131211110-1010010020111021-1010311133031010-0223200010132301-2332300222120123-3212133201002023-1121131203001031-3313330313021311"></a>

## os — os / 322202332110 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- os

<a id="canonical-1322303212201302-3130212320301231-0113001101122220-1003031021203100-2220033322332223-0023132320030010-1211131201210110-2102132300211012"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000000131212120-2111002313001132-1220221113133132-1203303001130133-0020220322001002-1033023232021320-1001310100330012-3000302103002133"></a>

## Direct properties — os / 322202332110 / 3

- [default_os_version](resources--securemesh_site--reference--group-004.md#canonical-3130221301333203-1013132201313331-0310132122031210-2120132103113323-3130311032011002-3330301112311023-2330230222120003-2001222200213200): complete subsection reference.

<a id="canonical-0232310003231232-1103300210021021-3022201321301113-1232230302313321-1232231323221032-1212330223331200-0310200103133112-1220110103122232"></a>

<a id="canonical-0103213001121011-1123302101102332-0031320022311203-2333210003331013-3232230121200020-3002230033130320-1120222121020021-1122200333321202"></a>

## operating_system_version property — os / 322202332110 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-1030230111232200-2322313333120033-2201113130302203-0000220212232210-0210210103012122-2110032203123321-1133311131121003-1122320000123302"></a>

## Next pages — os / 322202332110 / 5

- [os.default_os_version](resources--securemesh_site--reference--group-004.md#canonical-3130221301333203-1013132201313331-0310132122031210-2120132103113323-3130311032011002-3330301112311023-2330230222120003-2001222200213200)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3130221301333203-1013132201313331-0310132122031210-2120132103113323-3130311032011002-3330301112311023-2330230222120003-2001222200213200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102003111122120-1113132102031322-2001233003000300-1331122202301101-3123303112031012-2333101331102303-3312032330303222-0001103311031031"></a>

## os.default_os_version — default_os_version / 312001333331 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [os](resources--securemesh_site--reference--group-004.md#canonical-0102133232320131-2231010312223023-0203311201032131-2210230021210322-0310110201031312-2321113100212222-0201202312033330-3000120030223322)
- os.default_os_version

<a id="canonical-3031210222033031-0112020002333111-2102132210011321-1220213331332311-0212103212330001-1003010210113123-2222302331010130-2231202223022033"></a>

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
default_os_version = {}
```

<a id="canonical-1212303122221232-1211032313210330-1022223220000103-0120132110203021-2032103132313000-1132303101113231-0001301020222330-0013001020230201"></a>

## Direct properties — default_os_version / 312001333331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332112001022120-2030021310101213-2332122033121100-2230223133211032-3220201100101333-1321221333330000-3231201331023333-3020333232103132"></a>

## Next pages — default_os_version / 312001333331 / 4

- [os](resources--securemesh_site--reference--group-004.md#canonical-0102133232320131-2231010312223023-0203311201032131-2210230021210322-0310110201031312-2321113100212222-0201202312033330-3000120030223322)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120001330332130-3122213213323203-0212321133322102-2131232200302303-3103031301002231-1220132131011230-2023002333030233-3000223300121031"></a>

## performance_enhancement_mode — performance_enhancement_mode / 322023213011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- performance_enhancement_mode

<a id="canonical-0200201230210303-1312310131231003-0131311210012310-3223120003223203-3112102211300332-3000211023302211-1122020102231101-1003133121023302"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313010330032010-2102223030321310-0012101222000131-3033303120311323-1132031120131012-1301100000212233-1203313101303110-3120203310111030"></a>

## Direct properties — performance_enhancement_mode / 322023213011 / 3

- [perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133): complete subsection reference.

- [perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113): complete subsection reference.

<a id="canonical-0022010133113023-2213230031321103-2311132233020002-2102031301010210-2202102220021220-1023232211102312-0230131010120010-1203300330212012"></a>

## Next pages — performance_enhancement_mode / 322023213011 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320021030211121-3321221322211213-2233302022122120-2311321012003303-3101333031311030-2220300030222322-1211322223203032-0233100310013003"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 101211310133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-1331133330203133-2230131130221321-1323000002212112-3121331301030301-1330111123020010-2031013122233313-3332332233000301-0111121212000332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010313022032131-2332002313311333-2110121221223210-1201021033203330-3131022323100110-0313210030020231-3010203302121123-2033223030122001"></a>

## Direct properties — perf_mode_l3_enhanced / 101211310133 / 3

- [jumbo](resources--securemesh_site--reference--group-004.md#canonical-3112213213200203-0221323100132230-2300031111021101-0321312220231311-1132100000103231-0121333011313332-2302200312000222-2120113131232000): complete subsection reference.

- [no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-2103001320013003-3210313002320212-2211112310011010-1301202000021222-3110112210301130-0213312330322203-2013030010101320-2110231202022120): complete subsection reference.

<a id="canonical-3220113203302111-2103311330020100-0031302202213203-0221003321212201-1303122311222132-1122303131211120-1003300211130223-2022013001031330"></a>

## Next pages — perf_mode_l3_enhanced / 101211310133 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site--reference--group-004.md#canonical-3112213213200203-0221323100132230-2300031111021101-0321312220231311-1132100000103231-0121333011313332-2302200312000222-2120113131232000)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-2103001320013003-3210313002320212-2211112310011010-1301202000021222-3110112210301130-0213312330322203-2013030010101320-2110231202022120)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3112213213200203-0221323100132230-2300031111021101-0321312220231311-1132100000103231-0121333011313332-2302200312000222-2120113131232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320010112323322-1221321000210131-0113020200010331-2333021130200332-0233111011133102-1311030030332021-0102332323121231-3032233111212131"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 111213312332 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-3210230031131113-3223233212002002-0103311323111111-1130010001130231-2120132033020332-3123113231321031-0120000301103112-2201133010201100"></a>

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
jumbo = {}
```

<a id="canonical-0311313231102323-0333212022110031-3010131123122213-1020012322031233-2020202021101123-1122320230013310-3133221310133010-0000021320330021"></a>

## Direct properties — jumbo / 111213312332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002302013103023-1012010221113330-0013011010003323-3002301132101112-1113311133211323-2123013033010133-0333330200300032-2113003122303021"></a>

## Next pages — jumbo / 111213312332 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2103001320013003-3210313002320212-2211112310011010-1301202000021222-3110112210301130-0213312330322203-2013030010101320-2110231202022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201221220222303-2231223230101201-3302033100302132-3210300032031200-0221031000132301-1031331311301011-1330231233101221-2112110032013020"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 011111020212 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-0031023110120211-0200100020310201-2221010121133010-1222022123110330-3112130133311102-1211320013210022-2021303102112013-1213300203002310"></a>

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
no_jumbo = {}
```

<a id="canonical-3211231132000010-3013311211013000-1013333012332321-2311110221203300-1331222232312313-0212202300203313-2303211211102030-3012100000230211"></a>

## Direct properties — no_jumbo / 011111020212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131312333213033-3220312301132030-0132003332001102-3121112120210110-3233112302232311-1012132021203231-3323123201302132-2131013032112212"></a>

## Next pages — no_jumbo / 011111020212 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-2333313212022301-2232303021010000-0032233333133232-1221302020313113-1130010220031303-0010313123000023-0322332132110212-2103320020130133)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002002303011031-2212112210012011-3003323032220122-1021323312010013-1333010203321233-1031330232303322-1130201220123011-1201111133030001"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 101210221030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-1013301130032023-0222200031001211-0222303310113010-2201110033331202-1212330221230000-1130033032220321-2000030111301000-2232332232213022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121301022300220-0233223113323010-1211303011010323-1011010303131310-0303000131301213-1103301231311311-2213112212322330-0020123020122321"></a>

## Direct properties — perf_mode_l7_enhanced / 101210221030 / 3

- [jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-2313023322310013-0133233003203313-2101132102321112-3221232221203322-0231103233100001-2212132330323101-3220311230212323-1110322023130022): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-0112221201313312-0012323212111133-1033223033133331-0022022322012113-2001011300300031-2011110210102012-1032000011020122-1030123031200022): complete subsection reference.

<a id="canonical-2200102213002213-3330213302131023-1012000320201203-3233212110020103-3211230233100013-3002330203220031-3303203202322222-1110020113120110"></a>

## Next pages — perf_mode_l7_enhanced / 101210221030 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-2313023322310013-0133233003203313-2101132102321112-3221232221203322-0231103233100001-2212132330323101-3220311230212323-1110322023130022)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-0112221201313312-0012323212111133-1033223033133331-0022022322012113-2001011300300031-2011110210102012-1032000011020122-1030123031200022)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2313023322310013-0133233003203313-2101132102321112-3221232221203322-0231103233100001-2212132330323101-3220311230212323-1110322023130022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012211112131033-2213031330232223-2212030232120312-1303220131300331-1221301022220310-3130330100030023-2022300213121013-0021203030133312"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 323202303133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3222310222320011-3202320000202203-3312332010020321-1122202301122133-0210102322300222-1100030122101011-0000002113202203-3011012012123211"></a>

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
jumbo_disabled = {}
```

<a id="canonical-3133112332333311-2021132000210311-1213202211101303-2223322113133220-0101211313333330-2011121013023030-1320131132201121-2001302310110222"></a>

## Direct properties — jumbo_disabled / 323202303133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101202121122333-0201003211311102-0011031000110011-2111312012030031-2001231220130300-3231332022203202-3222222233331131-1110121201231110"></a>

## Next pages — jumbo_disabled / 323202303133 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0112221201313312-0012323212111133-1033223033133331-0022022322012113-2001011300300031-2011110210102012-1032000011020122-1030123031200022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003332031020001-3130123210032211-0032323001131212-0231233000030021-1231210302210121-1103003003311211-1220332023102102-0101333021213110"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 112331303322 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-0011232330303313-0020020101001012-3230233302033223-1331120322113100-0123330033123033-0233011110303003-2112011120131023-2130202111131033)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-0132111022322021-2002211322301322-2103112330121223-3031103313110011-3330122200313031-0321233222130212-1131013031033011-1033011220320001"></a>

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
jumbo_enabled = {}
```

<a id="canonical-0102113122021210-0210130120130110-2311010100320121-1230320213311102-2110331111023300-2330020232002301-0221303203011201-2100123301210132"></a>

## Direct properties — jumbo_enabled / 112331303322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302133112011321-2000103010330132-3003321130013321-3233021010202202-3121003311012111-2320003023123103-1012130210112011-2300132112331232"></a>

## Next pages — jumbo_enabled / 112331303322 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-3232120110302303-1311321222312120-3232201102102233-2312321010301102-0323132122213031-2031232123011202-0000220212002002-2122310230113113)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1000012000123120-2210011233010203-3333022101301022-1133130210233321-1101033113332231-1333113032202102-2032212300223010-2110321233023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323120200331101-2223312033023310-0023300223332233-1030330221110233-1331310033202321-2020020002130223-2123102311122133-1200012221230023"></a>

## sw — sw / 332012023000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- sw

<a id="canonical-3223310122101302-0103131020331231-0201231112030130-3130013311322202-2000121210232312-1133210323321010-2320311023031101-1030210223333112"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013020030122030-2033001301002122-2121021221032103-1012301213123232-3322330021330321-2332023102231201-2012013132200103-3010323231231131"></a>

## Direct properties — sw / 332012023000 / 3

- [default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-2130220332131230-1233302312003212-2130313101302312-2223223103213100-0120031232133312-1103320332211223-3301021212311303-3022321223132210): complete subsection reference.

<a id="canonical-1332123030321022-1111110220313130-1122312000200103-0233313210033012-0001013301111010-2200300232011210-1013032122123302-2201313130121013"></a>

<a id="canonical-2113230232002202-2310101123202110-1223323130100032-0020213230233300-2120322301133133-0301233110320222-1022122021232010-1020121333301031"></a>

## volterra_software_version property — sw / 332012023000 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2130310330231310-2203103203010103-1312031202122220-2133323232223123-1201213333212030-2003120033222312-1110013330101310-0000203033323002"></a>

## Next pages — sw / 332012023000 / 5

- [sw.default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-2130220332131230-1233302312003212-2130313101302312-2223223103213100-0120031232133312-1103320332211223-3301021212311303-3022321223132210)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2130220332131230-1233302312003212-2130313101302312-2223223103213100-0120031232133312-1103320332211223-3301021212311303-3022321223132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232132230001130-0212100023202301-0033020132232022-0101223311322122-3333001311202333-3232131030202322-1101012221003021-1120023102130033"></a>

## sw.default_sw_version — default_sw_version / 123223201222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [sw](resources--securemesh_site--reference--group-004.md#canonical-1000012000123120-2210011233010203-3333022101301022-1133130210233321-1101033113332231-1333113032202102-2032212300223010-2110321233023213)
- sw.default_sw_version

<a id="canonical-1102033032300020-1120013312333021-3103231023002121-1322213122033321-2221320122111032-0301003011130011-1202023301003211-3101023322211312"></a>

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
default_sw_version = {}
```

<a id="canonical-3130323030311233-3130211023033103-2311023102022001-2200021320123131-0131122031002302-3101003032303112-2132022113310212-3121103321301223"></a>

## Direct properties — default_sw_version / 123223201222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110331202301202-3102000031201000-2131303323100301-3132211001122213-1013031221000213-1330013100031233-3232201101331113-1022200112112320"></a>

## Next pages — default_sw_version / 123223201222 / 4

- [sw](resources--securemesh_site--reference--group-004.md#canonical-1000012000123120-2210011233010203-3333022101301022-1133130210233321-1101033113332231-1333113032202102-2032212300223010-2110321233023213)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3231302133333333-3210323010223133-3130032303201002-1121332221230320-2013212123011031-2102312110221123-1333231321102102-3213311103101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131312012023310-3322010323102121-3200320311223203-0300300331013330-1031321313201022-3002110200320320-1201010330030332-3200301322232323"></a>

## timeouts — timeouts / 000003202331 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- timeouts

<a id="canonical-0330122001113023-1322112322032123-0131023302203121-2210131202331232-0020311311332023-2301222310022112-1202303321003130-3003031113212310"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200333323222022-3203332322012311-0323202320300201-2212021131100113-2233003030331101-0232203123231011-0221101130130302-0001221332203100"></a>

## Direct properties — timeouts / 000003202331 / 3

<a id="canonical-0233102122122003-3220203011011022-2111033110002010-3130213021103221-0322213102233011-1100112320322203-0103132201001020-0132223032310131"></a>

<a id="canonical-2111312101001312-0020031000323133-3212002320200210-2310322013123210-3023003323230102-1230011320021322-2203130022203231-3012111300131030"></a>

## create property — timeouts / 000003202331 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2101112113221302-2312230030133002-2211211032002113-3013200021111032-2031121200033303-1120113330122001-1330023121112100-3223100133301000"></a>

<a id="canonical-3300213302121201-0123301122022011-2331010330112321-0120022120112220-1223011002221121-3300330012331220-2310031102102332-3033203003210331"></a>

## delete property — timeouts / 000003202331 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2213310131322121-1231022303222220-3010003323321103-0010303112122202-0322021321203331-3310102012331100-2331313232212013-1303110131323122"></a>

<a id="canonical-3332303331323002-3133100230233320-3310120212003101-1032232203022013-2033322333203332-3320203231001232-1012301233333310-2022132012010032"></a>

## read property — timeouts / 000003202331 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3210202220122102-3120101303101210-0000101123102300-1220000321321331-2021000200300200-3120002320200100-2203311232001123-1312110132022312"></a>

<a id="canonical-3123321131310303-3112122110312310-1201020222122210-1132233032021200-3211220223300033-3030223123313111-3310003022220321-1230001011022131"></a>

## update property — timeouts / 000003202331 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0000000111023322-3201312103113130-2230200302031021-3110300201320300-3113322111011231-3100100223122110-2231213133031320-1323031313101222"></a>

## Next pages — timeouts / 000003202331 / 8

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232111203023203-3010331303212122-2030321112101310-1233332332310132-3122232220203202-1010010321111202-2301020220032022-1330133330223203"></a>

## waf_signatures — waf_signatures / 103200323220 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- waf_signatures

<a id="canonical-1333331221112111-1230332223010200-1203221320210122-2111211122200321-3021230220001210-3132232112311312-0100312311231331-0000012200301310"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302311222220231-1211021232233312-0313322112100233-1231310333132322-3111213122332113-3200213120231221-3030301213212200-2312333210100313"></a>

## Direct properties — waf_signatures / 103200323220 / 3

- [automatic](resources--securemesh_site--reference--group-004.md#canonical-1120111213131011-2001322322102112-2323203033113120-3220013133210230-3001310302110030-1313310102231011-2002031122320000-0322121320312013): complete subsection reference.

- [manual](resources--securemesh_site--reference--group-004.md#canonical-0301232231103220-2013210121011130-0212132301310310-3110313332333231-0323233003332131-0131221331033102-3202131230231123-1120001033033133): complete subsection reference.

<a id="canonical-3102223132330132-1111012332123001-0211323301032323-1310100030232222-3320103321313130-0133101333120223-1323131222202313-3111203011103302"></a>

## Next pages — waf_signatures / 103200323220 / 4

- [waf_signatures.automatic](resources--securemesh_site--reference--group-004.md#canonical-1120111213131011-2001322322102112-2323203033113120-3220013133210230-3001310302110030-1313310102231011-2002031122320000-0322121320312013)
- [waf_signatures.manual](resources--securemesh_site--reference--group-004.md#canonical-0301232231103220-2013210121011130-0212132301310310-3110313332333231-0323233003332131-0131221331033102-3202131230231123-1120001033033133)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1120111213131011-2001322322102112-2323203033113120-3220013133210230-3001310302110030-1313310102231011-2002031122320000-0322121320312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002003132330000-1100032021010320-1030223320203201-3122202013230020-2333112213322100-1222011013122232-1010033330130132-1032202303032203"></a>

## waf_signatures.automatic — automatic / 112220302301 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200)
- waf_signatures.automatic

<a id="canonical-3200220232003323-0003021002223332-1233031032022010-3131103322022310-2110223332132100-0200113312031002-3101103022302303-0030223102002320"></a>

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
automatic = {}
```

<a id="canonical-3001233023032122-0032320011332102-1220232210021332-0031221202031111-2203100003121213-1003000132232301-3121110013100102-3113203031111223"></a>

## Direct properties — automatic / 112220302301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111122111010302-3200002232012101-1220220300321021-0113213003312112-2232231013232103-1100302332111300-2311201303130022-2211320103000003"></a>

## Next pages — automatic / 112220302301 / 4

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0301232231103220-2013210121011130-0212132301310310-3110313332333231-0323233003332131-0131221331033102-3202131230231123-1120001033033133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111133130201300-0221131323022312-3222213202333332-3032330032011211-1100112120112013-2120030133110210-3030023123000201-2220000330121023"></a>

## waf_signatures.manual — manual / 122133033012 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200)
- waf_signatures.manual

<a id="canonical-2133132112202113-3222122211021301-1022132222111203-2212210111001332-3311123303100321-3210200311001013-2120231303000212-3002031110012222"></a>

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
manual = {}
```

<a id="canonical-1303312032022121-1303002222132231-1333333020330003-1110001132302321-3131201002133210-2300011110320300-1302321013313232-2131003001022133"></a>

## Direct properties — manual / 122133033012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233220320302201-3302122311133033-2102210112102210-3313221233312200-3032223112121321-3311113232202122-1311331222110133-2110231133030301"></a>

## Next pages — manual / 122133033012 / 4

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-3203310323033131-0010220212131332-1113111313110122-1002100223302222-1322303202020202-2131112033300232-3102133213010130-3330222230133200)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
