---
page_title: "xcsh_virtual_network reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network reference."
---

# xcsh_virtual_network reference

<a id="canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221222322330012-3023211033012101-1323121032010103-3100233132011023-0331312333330130-3310312310310121-0001222011103331-3302233231121022"></a>

## Property reference — Property reference / 000031031133 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- Property reference

<a id="canonical-2301233220032202-1112011202103301-2230223332230203-1310101132332300-1232022300213232-2023021103031020-1113311020022032-1200133300210002"></a>

## Direct properties — Property reference / 000031031133 / 3

<a id="canonical-3202031110201221-0312131223330102-0220312002233221-0303002103013111-2132022320320103-0201221230100003-2030212112002123-3333312212203112"></a>

<a id="canonical-1332333013020331-3230331123132130-1012122102213330-0003230313021032-0301023110222312-3030103131133323-1210213231111330-3102301313320032"></a>

## annotations property — Property reference / 000031031133 / 4

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

<a id="canonical-2103322103311332-3200203232223032-3120120112212000-1223333032111011-2102221201001222-1311223233130311-1311302001123012-3102131032322002"></a>

<a id="canonical-2233120333200223-0131221233131311-3302313223131211-2030231313110130-2112201101300033-3312103302221332-2012031230012232-3021121301022210"></a>

## description property — Property reference / 000031031133 / 5

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

<a id="canonical-2211333322001031-0001320122223031-0101221002032002-2301321202112030-1033003222122010-1301010220210020-3211100321020313-2320103311310113"></a>

<a id="canonical-1203222031203331-3331321211123103-1302123120101002-3221032321103023-0110301113121201-3301213133301133-0022203311110323-0222211223213310"></a>

## disable property — Property reference / 000031031133 / 6

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

- [global_network](resources--virtual_network--reference--group-001.md#canonical-2022100300112003-1220202113121001-1131003021302131-3131310112012213-0021220111013333-0303333111203120-3031131131222030-2012313201111120): complete subsection reference.

<a id="canonical-2020230120320201-0021033321221123-2302221201231031-1213111030010133-0221202201113221-3001232120022321-0130111330232321-0011222020031232"></a>

<a id="canonical-1033303202001131-1023220231111123-0122223133202120-3123120232013023-0333010332003331-0022011133331110-2102223302233312-1301011001132322"></a>

## ID property — Property reference / 000031031133 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0123223103020310-0212001021231320-1301023332132311-3010120103311031-3131122010230122-2202000123103323-3222303032122333-2333100120012202"></a>

<a id="canonical-0330332113010231-0301221201010110-3330212323123321-1010133333223112-0001303120211110-2031211331301113-3223332201321201-0320331020200222"></a>

## labels property — Property reference / 000031031133 / 8

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

<a id="canonical-1233103110331230-0020003210331001-2300013301013303-3320121101121302-0203320231103010-2233221100312111-0301001213020300-1131312031311121"></a>

<a id="canonical-2122031232230001-1032203033312131-3030303103230201-3111122020321112-3231313033212200-0312230303311211-3101302323101323-0232121303120213"></a>

## name property — Property reference / 000031031133 / 9

Type: `"string"`. Required.

Name of the Virtual Network. Must be unique within the namespace.

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

<a id="canonical-3312113030022313-3023230103013202-0211133300220101-0221013003121132-1320002103133303-2222313033002103-3033123213200231-0001223122102201"></a>

<a id="canonical-1323332100023022-1122002211212310-1333010021111003-3312120103210201-0221321210213211-2203112200301310-3300033011310122-3002023322100032"></a>

## namespace property — Property reference / 000031031133 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Virtual Network. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-1112332202332020-3320121331310333-0030112220102323-3022022211031210-2111210021233300-1100301133210230-2103221231201332-0212313013201101): complete subsection reference.

- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-2221221203312202-2002322102302002-1023111120132301-3331003331013221-1031132000201122-3011230312213030-3013302022320022-3112321211102311): complete subsection reference.

- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011): complete subsection reference.

- [timeouts](resources--virtual_network--reference--group-001.md#canonical-0303303121130200-2111222013322103-2221220102302100-1210333321103132-0102210313121311-0100033232232312-1020112202101030-2202320112020332): complete subsection reference.

<a id="canonical-0103330013333003-1303321232011103-2103303003303002-3213322000301311-0323201102313022-2332330030112320-3313003003110310-3101323220211311"></a>

## All schema paths — Property reference / 000031031133 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_network--reference--group-001.md#canonical-3202031110201221-0312131223330102-0220312002233221-0303002103013111-2132022320320103-0201221230100003-2030212112002123-3333312212203112) |
| `description` | [description](resources--virtual_network--reference--group-001.md#canonical-2103322103311332-3200203232223032-3120120112212000-1223333032111011-2102221201001222-1311223233130311-1311302001123012-3102131032322002) |
| `disable` | [disable](resources--virtual_network--reference--group-001.md#canonical-2211333322001031-0001320122223031-0101221002032002-2301321202112030-1033003222122010-1301010220210020-3211100321020313-2320103311310113) |
| `global_network` | [global_network](resources--virtual_network--reference--group-001.md#canonical-0230331330003200-3232222031030022-1301311200302223-2301230032002112-2332003030001103-1112331330101210-2231122230132030-2000113032210322) |
| `id` | [ID](resources--virtual_network--reference--group-001.md#canonical-2020230120320201-0021033321221123-2302221201231031-1213111030010133-0221202201113221-3001232120022321-0130111330232321-0011222020031232) |
| `labels` | [labels](resources--virtual_network--reference--group-001.md#canonical-0123223103020310-0212001021231320-1301023332132311-3010120103311031-3131122010230122-2202000123103323-3222303032122333-2333100120012202) |
| `name` | [name](resources--virtual_network--reference--group-001.md#canonical-1233103110331230-0020003210331001-2300013301013303-3320121101121302-0203320231103010-2233221100312111-0301001213020300-1131312031311121) |
| `namespace` | [namespace](resources--virtual_network--reference--group-001.md#canonical-3312113030022313-3023230103013202-0211133300220101-0221013003121132-1320002103133303-2222313033002103-3033123213200231-0001223122102201) |
| `site_local_inside_network` | [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-1321012110111033-2233010013110111-0330300110313313-3232130013323110-2312211112303202-3001322310110003-2121201002133002-2201000123033222) |
| `site_local_network` | [site_local_network](resources--virtual_network--reference--group-001.md#canonical-2210131023030233-2213202201131012-2323001102303233-2031210031320223-3003121303223030-1101213300211111-3130001010112330-1223132220102213) |
| `static_routes` | [static_routes](resources--virtual_network--reference--group-001.md#canonical-3101313321120111-2332333223033330-2330033200300120-3310320211013332-2222233322332100-2333122131230132-1223122211110002-0103310332222233) |
| `static_routes.attrs` | [static_routes.attrs](resources--virtual_network--reference--group-001.md#canonical-0203323301132033-0201211110112321-1001000220230230-0112311031323332-3213031023002033-0032210020112100-1032113123110022-0020220011310123) |
| `static_routes.default_gateway` | [static_routes.default_gateway](resources--virtual_network--reference--group-001.md#canonical-3130023313112230-3200212320103332-1003332310122021-2303031003010301-1232322003013230-0001300032120022-0302300031132202-3231130001023212) |
| `static_routes.ip_address` | [static_routes.ip_address](resources--virtual_network--reference--group-001.md#canonical-0032110131203211-2312222001123222-3003210001323122-2332230011233010-1233213022200231-1022232333010132-2200311033201121-3200302110223011) |
| `static_routes.ip_prefixes` | [static_routes.ip_prefixes](resources--virtual_network--reference--group-001.md#canonical-1310113100100122-0202211021131013-1110232210322130-0301323233310231-0231111200111103-3112301321211111-0312300311131313-3111233121021232) |
| `static_routes.node_interface` | [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-1132022012321230-1321321023133133-0221000233220033-1023111300002012-0322212330321212-1123231111133323-0022012213120333-2331012100301023) |
| `static_routes.node_interface.list` | [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-2303023231023032-3321133212013331-2120230201212021-3102313000303102-0033133310100121-0312220102023010-1010322120200012-1303211112322220) |
| `static_routes.node_interface.list.interface` | [static_routes.node_interface.list.interface](resources--virtual_network--reference--group-001.md#canonical-3011023120001230-3220003130133012-2212013300113103-3212231033000122-2231333103312231-0011003230332120-1001323001130033-1110232212211221) |
| `static_routes.node_interface.list.interface.kind` | [static_routes.node_interface.list.interface.kind](resources--virtual_network--reference--group-001.md#canonical-0221031321333112-0000011312223111-1212021032023012-1223323023221121-3321312303231123-1300002201323201-1311132013110021-2301023000102122) |
| `static_routes.node_interface.list.interface.name` | [static_routes.node_interface.list.interface.name](resources--virtual_network--reference--group-001.md#canonical-2001301310302023-3200100131330303-0002231033010112-0113321213122000-2213220133312211-3301131222311102-1021000022113131-3323012100101120) |
| `static_routes.node_interface.list.interface.namespace` | [static_routes.node_interface.list.interface.namespace](resources--virtual_network--reference--group-001.md#canonical-3300322200230211-1333132312221303-2030331100033002-1030121210013333-0002312202003201-2320330313202103-3200000312231223-3202320022313121) |
| `static_routes.node_interface.list.interface.tenant` | [static_routes.node_interface.list.interface.tenant](resources--virtual_network--reference--group-001.md#canonical-2113323101331222-3012212212211322-0202212221030032-3031300202223223-0023113011221323-1013132100112003-0331110010331111-0220303102010220) |
| `static_routes.node_interface.list.interface.uid` | [static_routes.node_interface.list.interface.uid](resources--virtual_network--reference--group-001.md#canonical-3201010232110330-3012103202211100-1302110301322012-1121023000033031-1321300230121102-3300310321131313-1330230213321302-3020101230111101) |
| `static_routes.node_interface.list.node` | [static_routes.node_interface.list.node](resources--virtual_network--reference--group-001.md#canonical-3010301030012210-3231200101200023-3133321032233130-2230311033330312-2023230133322233-1132102011002113-0311322233121002-1333200302220122) |
| `timeouts` | [timeouts](resources--virtual_network--reference--group-001.md#canonical-1210032230123131-1032010330032202-3123203200020303-3230322023100233-0023231221101333-3121212111122023-0222310311310200-1000300131032322) |
| `timeouts.create` | [timeouts.create](resources--virtual_network--reference--group-001.md#canonical-1201233021301313-2030022332020121-3313232211132021-1233201102202222-2201001302100001-2011222323323101-3133130102322301-1233131210013112) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_network--reference--group-001.md#canonical-2122333010012210-1303011111210130-0233220211132110-0202211211013002-0120122020102023-3322131010222300-1230320202011300-0333221320011233) |
| `timeouts.read` | [timeouts.read](resources--virtual_network--reference--group-001.md#canonical-0200112133323303-1311223022331112-1321203202302213-1002230301131302-3102111020303201-2203303122213323-2120330101302033-0320311001202203) |
| `timeouts.update` | [timeouts.update](resources--virtual_network--reference--group-001.md#canonical-1330030133310223-2111203303022201-0122002133000233-2131323033000311-0211203313111103-3312130031300322-3013311112130121-1211302020222320) |

<a id="canonical-1013110133200122-2223213103102012-2212312022033221-2131001203220132-1323022302112112-1031121102211031-3312100002222232-1201231312112012"></a>

## Next pages — Property reference / 000031031133 / 12

- [global_network](resources--virtual_network--reference--group-001.md#canonical-2022100300112003-1220202113121001-1131003021302131-3131310112012213-0021220111013333-0303333111203120-3031131131222030-2012313201111120)
- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-1112332202332020-3320121331310333-0030112220102323-3022022211031210-2111210021233300-1100301133210230-2103221231201332-0212313013201101)
- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-2221221203312202-2002322102302002-1023111120132301-3331003331013221-1031132000201122-3011230312213030-3013302022320022-3112321211102311)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- [timeouts](resources--virtual_network--reference--group-001.md#canonical-0303303121130200-2111222013322103-2221220102302100-1210333321103132-0102210313121311-0100033232232312-1020112202101030-2202320112020332)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-2022100300112003-1220202113121001-1131003021302131-3131310112012213-0021220111013333-0303333111203120-3031131131222030-2012313201111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221002131331210-3023111332211023-0303200210333300-3131223120233111-3102130030210012-3231022313020221-3200122122223210-3310213023211112"></a>

## global_network — global_network / 031120031203 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- global_network

<a id="canonical-0230331330003200-3232222031030022-1301311200302223-2301230032002112-2332003030001103-1112331330101210-2231122230132030-2000113032210322"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

Receipt-pinned upstream constraints:

```json
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

- [global_network](resources--virtual_network--reference--group-001.md#canonical-0230331330003200-3232222031030022-1301311200302223-2301230032002112-2332003030001103-1112331330101210-2231122230132030-2000113032210322)
- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-1321012110111033-2233010013110111-0330300110313313-3232130013323110-2312211112303202-3001322310110003-2121201002133002-2201000123033222)
- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-2210131023030233-2213202201131012-2323001102303233-2031210031320223-3003121303223030-1101213300211111-3130001010112330-1223132220102213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_network = {}
```

<a id="canonical-3023313232031323-0231023313320011-2222321313020130-3023231321012100-2322113121123013-1123111213102003-1320211331010013-1301230210131032"></a>

## Direct properties — global_network / 031120031203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133002122130030-0221120323331112-3121212230332000-1211311022131111-0022233312111201-2002221000121321-0300231221302313-3113232321013011"></a>

## Next pages — global_network / 031120031203 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-1112332202332020-3320121331310333-0030112220102323-3022022211031210-2111210021233300-1100301133210230-2103221231201332-0212313013201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322011022323100-2130333223013200-1312033013021322-0202123311320021-0321120211302003-2001013321100100-3220100013322011-2232233300102331"></a>

## site_local_inside_network — site_local_inside_network / 012101013302 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- site_local_inside_network

<a id="canonical-1321012110111033-2233010013110111-0330300110313313-3232130013323110-2312211112303202-3001322310110003-2121201002133002-2201000123033222"></a>

Type: `["object", {}]`. Optional.

Select the site-local inside network for site-internal connectivity.

Receipt-pinned upstream constraints:

```json
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
site_local_inside_network = {}
```

<a id="canonical-1133221211200323-3233331113020300-1111131313003200-0322203022322032-3023021011001230-0111122332230001-1002033133020203-0113120331033002"></a>

## Direct properties — site_local_inside_network / 012101013302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130001233310021-0312101302222030-3300021212132020-3020220312130322-2023210302123131-1113013132131323-0111302002213210-3332013000301221"></a>

## Next pages — site_local_inside_network / 012101013302 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-2221221203312202-2002322102302002-1023111120132301-3331003331013221-1031132000201122-3011230312213030-3013302022320022-3112321211102311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231131003331033-2111311121031130-1002122210100003-3003220021310003-2022002202230031-2323021003200201-2020313223332032-2130303033132211"></a>

## site_local_network — site_local_network / 232220123002 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- site_local_network

<a id="canonical-2210131023030233-2213202201131012-2323001102303233-2031210031320223-3003121303223030-1101213300211111-3130001010112330-1223132220102213"></a>

Type: `["object", {}]`. Optional.

Select a site-local virtual network when connectivity must remain within one site.

Receipt-pinned upstream constraints:

```json
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
site_local_network = {}
```

<a id="canonical-0101231331200321-1201003101031012-1030321132033022-0213310122111130-0121202322132202-2222033322312223-3011222011233231-0211021332013000"></a>

## Direct properties — site_local_network / 232220123002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200010311210013-0000002223302332-1010223132201103-2231311130321323-0310223033130030-1121022131303122-0033313310020303-3110212022303333"></a>

## Next pages — site_local_network / 232220123002 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211032313000203-0122221232033333-1223201312203010-0103311222100321-1300312020023103-2211300300333003-2031202003200201-1032222331100033"></a>

## static_routes — static_routes / 213100102332 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- static_routes

<a id="canonical-3101313321120111-2332333223033330-2330033200300120-3310320211013332-2222233322332100-2333122131230132-1223122211110002-0103310332222233"></a>

Type: `"object"`. list nested block, Optional.

List of static routes on the virtual network.

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
  "maxItems": 165,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 165,
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
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "165",
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

<a id="canonical-1022123121113113-0302302320320321-2022211313111323-0033011121210123-2132131022313212-0200203012131130-0312003312133102-3201021022000220"></a>

## Direct properties — static_routes / 213100102332 / 3

<a id="canonical-0203323301132033-0201211110112321-1001000220230230-0112311031323332-3213031023002033-0032210020112100-1032113123110022-0020220011310123"></a>

<a id="canonical-1233230131233230-3133333022023230-1112021022003323-1013303022111131-0120311313212030-2031003200232213-0022332233323200-0023202111320221"></a>

## attrs property — static_routes / 213100102332 / 4

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

- [default_gateway](resources--virtual_network--reference--group-001.md#canonical-3221203221010200-0123123213133132-0330102021123113-2001033320003001-2213101113330113-2000003033002131-3011331102000311-2013033232022230): complete subsection reference.

<a id="canonical-0032110131203211-2312222001123222-3003210001323122-2332230011233010-1233213022200231-1022232333010132-2200311033201121-3200302110223011"></a>

<a id="canonical-2233131113332231-0201313210301021-2312133300302003-1003102120031303-1200232131022331-0121000231301020-0111311012333003-0211213333223323"></a>

## ip_address property — static_routes / 213100102332 / 5

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

<a id="canonical-1310113100100122-0202211021131013-1110232210322130-0301323233310231-0231111200111103-3112301321211111-0312300311131313-3111233121021232"></a>

<a id="canonical-0203121132321203-2030302213313332-1231110103221013-2200230310321233-1323022103320113-0332001222211020-0030110330322310-3020110213002320"></a>

## ip_prefixes property — static_routes / 213100102332 / 6

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

- [node_interface](resources--virtual_network--reference--group-001.md#canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020): complete subsection reference.

<a id="canonical-3001312013112303-1012122110030033-3232223111330001-3031231231313223-1000322120202303-3213310033123111-1113213102121103-0001313231200000"></a>

## Next pages — static_routes / 213100102332 / 7

- [static_routes.default_gateway](resources--virtual_network--reference--group-001.md#canonical-3221203221010200-0123123213133132-0330102021123113-2001033320003001-2213101113330113-2000003033002131-3011331102000311-2013033232022230)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-3221203221010200-0123123213133132-0330102021123113-2001033320003001-2213101113330113-2000003033002131-3011331102000311-2013033232022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020320000311133-2221310303032022-3113100122230212-0111132032333232-3002100200333021-2002232131000320-1113101130332102-1000320231333202"></a>

## static_routes.default_gateway — default_gateway / 112320022003 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- static_routes.default_gateway

<a id="canonical-3130023313112230-3200212320103332-1003332310122021-2303031003010301-1232322003013230-0001300032120022-0302300031132202-3231130001023212"></a>

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

<a id="canonical-3111101313022133-3131122333232232-3111133130032212-1231102100001330-0033330313020121-3212101203220102-1012012010021301-0122020300110301"></a>

## Direct properties — default_gateway / 112320022003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010110223003201-2231300123200233-0023213222030231-3310232221013323-2020311201302231-0303223102003013-2022100223002302-1030210021131021"></a>

## Next pages — default_gateway / 112320022003 / 4

- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103211123033032-3302131311210211-1132131023300200-3211111021112120-2320032131020020-1322010323300131-1133211120010230-3131220231103020"></a>

## static_routes.node_interface — node_interface / 221212303332 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- static_routes.node_interface

<a id="canonical-1132022012321230-1321321023133133-0221000233220033-1023111300002012-0322212330321212-1123231111133323-0022012213120333-2331012100301023"></a>

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

<a id="canonical-2123211223320311-0200032301100003-2233302200300011-2321132123300000-2310023312113012-2313323331030301-1000331322110322-2002212222023120"></a>

## Direct properties — node_interface / 221212303332 / 3

- [list](resources--virtual_network--reference--group-001.md#canonical-3113030110330221-1133011111002203-0202221330332200-2022120100231200-0120201121330031-2101111300211323-1233310210323111-0121102112121213): complete subsection reference.

<a id="canonical-3010213112000303-1133231201202113-2131223202231313-2320113322000001-0013013203122032-3021231212001021-2011002021232220-3010332333202012"></a>

## Next pages — node_interface / 221212303332 / 4

- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-3113030110330221-1133011111002203-0202221330332200-2022120100231200-0120201121330031-2101111300211323-1233310210323111-0121102112121213)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-3113030110330221-1133011111002203-0202221330332200-2022120100231200-0120201121330031-2101111300211323-1233310210323111-0121102112121213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320212210022203-2022122112122010-3320232002303101-2213213012201230-1120232303023122-0003002202222023-2323121112300231-3320221303330201"></a>

## static_routes.node_interface.list — list / 030322230103 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020)
- static_routes.node_interface.list

<a id="canonical-2303023231023032-3321133212013331-2120230201212021-3102313000303102-0033133310100121-0312220102023010-1010322120200012-1303211112322220"></a>

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

<a id="canonical-3102310103001001-0311121331231013-3230300130023311-1300331112020202-3000212111103200-3012330300012000-3020011000213221-0212112331132312"></a>

## Direct properties — list / 030322230103 / 3

- [interface](resources--virtual_network--reference--group-001.md#canonical-1230233132111323-0223113232101121-0310223102031210-2233033033222232-3223201330231322-0022021122323021-3100310122231332-0001310113120013): complete subsection reference.

<a id="canonical-3010301030012210-3231200101200023-3133321032233130-2230311033330312-2023230133322233-1132102011002113-0311322233121002-1333200302220122"></a>

<a id="canonical-1033223023330123-0030023223300023-1132211213213200-1311103120310231-1233001212110120-3333121130222220-1323203012301012-3332022321031320"></a>

## node property — list / 030322230103 / 4

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

<a id="canonical-1000100003331303-0022332112211332-2112302222121310-2111112003332310-1111000323213212-1010001123132111-1112311133023031-3202130333011303"></a>

## Next pages — list / 030322230103 / 5

- [static_routes.node_interface.list.interface](resources--virtual_network--reference--group-001.md#canonical-1230233132111323-0223113232101121-0310223102031210-2233033033222232-3223201330231322-0022021122323021-3100310122231332-0001310113120013)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-1230233132111323-0223113232101121-0310223102031210-2233033033222232-3223201330231322-0022021122323021-3100310122231332-0001310113120013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221211320232332-0201102111130033-2003103320333221-0313311322223211-1232231230012200-1020331113220333-2221230002322103-3102102332000302"></a>

## static_routes.node_interface.list.interface — interface / 032300311313 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-2213023321010103-2233012001101320-3233300021131202-2111023330200032-0322011301321203-2230301113203323-0130330321331233-2010201022003011)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-3302111130120200-0103131230233103-1101230202332102-0221033032201123-1022023123100223-0101111232103302-2232123032231101-2100221310320020)
- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-3113030110330221-1133011111002203-0202221330332200-2022120100231200-0120201121330031-2101111300211323-1233310210323111-0121102112121213)
- static_routes.node_interface.list.interface

<a id="canonical-3011023120001230-3220003130133012-2212013300113103-3212231033000122-2231333103312231-0011003230332120-1001323001130033-1110232212211221"></a>

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

<a id="canonical-0021332113131110-0001310311021020-3131023213210321-0000202330313112-3012310112221130-3201213222303301-2102320212100302-0323131101230121"></a>

## Direct properties — interface / 032300311313 / 3

<a id="canonical-0221031321333112-0000011312223111-1212021032023012-1223323023221121-3321312303231123-1300002201323201-1311132013110021-2301023000102122"></a>

<a id="canonical-2220023301303303-2320030030102102-0101232221001023-2322301332203311-3233120102030200-1331300032220003-1303333102210310-1331110210110301"></a>

## kind property — interface / 032300311313 / 4

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

<a id="canonical-2001301310302023-3200100131330303-0002231033010112-0113321213122000-2213220133312211-3301131222311102-1021000022113131-3323012100101120"></a>

<a id="canonical-2221223123200121-3132230333320132-0233000101123213-1303011133001233-0003230233330103-1130210130203301-2210202331303321-0021231211321001"></a>

## name property — interface / 032300311313 / 5

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

<a id="canonical-3300322200230211-1333132312221303-2030331100033002-1030121210013333-0002312202003201-2320330313202103-3200000312231223-3202320022313121"></a>

<a id="canonical-0101322210223301-2030210202320233-3101321110002123-2100022210001021-3102313012212013-2203123221021330-3122023323213222-1222333020003201"></a>

## namespace property — interface / 032300311313 / 6

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

<a id="canonical-2113323101331222-3012212212211322-0202212221030032-3031300202223223-0023113011221323-1013132100112003-0331110010331111-0220303102010220"></a>

<a id="canonical-3332300032320103-2121101120003012-0012221130032203-0313012310020230-2230101231321023-2022222223302131-1313010100033211-2130113012002130"></a>

## tenant property — interface / 032300311313 / 7

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

<a id="canonical-3201010232110330-3012103202211100-1302110301322012-1121023000033031-1321300230121102-3300310321131313-1330230213321302-3020101230111101"></a>

<a id="canonical-0320223013113010-1121221231100022-1321031232001331-0211303003021231-1311221232203222-0030103122230130-2330032200120011-0010231002013111"></a>

## uid property — interface / 032300311313 / 8

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

<a id="canonical-1130311123331321-2300000221330323-3103100013200332-1102222321221010-2131111301203031-3102130222201321-3110210330113220-3222220012221123"></a>

## Next pages — interface / 032300311313 / 9

- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-3113030110330221-1133011111002203-0202221330332200-2022120100231200-0120201121330031-2101111300211323-1233310210323111-0121102112121213)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)

<a id="canonical-0303303121130200-2111222013322103-2221220102302100-1210333321103132-0102210313121311-0100033232232312-1020112202101030-2202320112020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211201203003110-0332131121130331-2033223113133123-3033031213012222-3301131313202012-2212312130212003-3232312330002323-2333301130203001"></a>

## timeouts — timeouts / 121003320130 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- timeouts

<a id="canonical-1210032230123131-1032010330032202-3123203200020303-3230322023100233-0023231221101333-3121212111122023-0222310311310200-1000300131032322"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023313011202120-3001110201301322-2322023313002332-0233200032110313-1233210311120223-0222303011312332-2220010212110223-0113030013130022"></a>

## Direct properties — timeouts / 121003320130 / 3

<a id="canonical-1201233021301313-2030022332020121-3313232211132021-1233201102202222-2201001302100001-2011222323323101-3133130102322301-1233131210013112"></a>

<a id="canonical-0333210030312003-3013130230111313-1211100110032100-0300202321222101-0311101213120212-3001102123201022-2111202221130110-0230002322323112"></a>

## create property — timeouts / 121003320130 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2122333010012210-1303011111210130-0233220211132110-0202211211013002-0120122020102023-3322131010222300-1230320202011300-0333221320011233"></a>

<a id="canonical-3101231202120201-3100331310211200-0023303122012200-2201232030322210-1100212203102210-2223310003010310-2110202021130112-3110032320232303"></a>

## delete property — timeouts / 121003320130 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0200112133323303-1311223022331112-1321203202302213-1002230301131302-3102111020303201-2203303122213323-2120330101302033-0320311001202203"></a>

<a id="canonical-1311113220031230-2210103102003302-0202032011310301-0002303212102021-2300201220222231-0022210313321210-2213333233132313-2231322321023221"></a>

## read property — timeouts / 121003320130 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1330030133310223-2111203303022201-0122002133000233-2131323033000311-0211203313111103-3312130031300322-3013311112130121-1211302020222320"></a>

<a id="canonical-3033312231012010-0223132130011212-3323003301023221-0021323232103323-1222230233112220-2122113320001322-3010000032021012-2121133231202103"></a>

## update property — timeouts / 121003320130 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2201120130333110-2222320332132010-3132331013100213-1231133111013002-2100210021300303-2332012323132113-0130012131210200-0301030302201200"></a>

## Next pages — timeouts / 121003320130 / 8

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
