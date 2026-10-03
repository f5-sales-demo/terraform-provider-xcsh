---
page_title: "xcsh_nginx_service_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery reference."
---

# xcsh_nginx_service_discovery reference

<a id="canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302121102121012-3133133012023231-3130001332212121-3110231112111122-0201133133001323-0000320020300113-0230311103112223-0111032131022223"></a>

## Property reference — Property reference / 101313032012 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- Property reference

<a id="canonical-2123303013323011-3230101100022133-1330021303220113-1200232221231021-1121123123320321-2101031132223003-2011120111311313-3333321003210003"></a>

## Direct properties — Property reference / 101313032012 / 3

<a id="canonical-1021313222322213-3130220330113323-2312303021122322-3032021321222011-1021003010220020-3232111022103010-1212100003203122-0233321132301010"></a>

<a id="canonical-0333330320121210-1231302320100211-3201331012110330-0313020133001011-3203013001213131-0010022132212100-0213330102203110-3111022223100213"></a>

## annotations property — Property reference / 101313032012 / 4

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
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

<a id="canonical-3020022132232010-1101331011103113-3210002123002121-3303001330032220-1320121232203111-2120223213132221-2132200101003023-1313213333200312"></a>

<a id="canonical-1300103133113002-3213223222010310-2211120002210210-1233013132222100-2003021122330211-0010211110010310-0332331132101312-3230322232132331"></a>

## description property — Property reference / 101313032012 / 5

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

<a id="canonical-2233002110022312-0313032310102320-0333212123331330-0231101211220221-2201033012220113-2012013003120223-0102312213033033-0300000213130300"></a>

<a id="canonical-2131301033331022-1331133330102002-1000300211121221-3001202113221201-1231320031010200-2323300320033222-3012021031011200-1201023111323233"></a>

## disable property — Property reference / 101313032012 / 6

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

- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312): complete subsection reference.

<a id="canonical-3302231223210022-2133312233002002-2131120030233312-0203221310232111-3230010201112112-0333122030222200-2121013020121222-0030330012130322"></a>

<a id="canonical-1313132323311323-3330331201101221-3233131203203203-3130212233222102-2100033222311103-1321322230303201-1012122113021313-3112313210303232"></a>

## ID property — Property reference / 101313032012 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1023022231013231-2303110301033023-3213101202131320-0221221030102331-1231100122031121-2223102131121301-1303123312000310-3220110012132201"></a>

<a id="canonical-0130301103332320-2322201230320320-1210021122330302-0303223332323201-2013302031100301-3323222122131323-1113233012223200-1113011010002111"></a>

## labels property — Property reference / 101313032012 / 8

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

<a id="canonical-1121002300200112-2101120032201212-2210103121100233-1211003312030221-1201233111232330-2222002331031110-1000030330031013-2003233003012023"></a>

<a id="canonical-1232303310011130-0012021132202311-0330230203213210-3022310100013311-3221010100331313-1300321230321110-3031112002321101-0300212321013221"></a>

## name property — Property reference / 101313032012 / 9

Type: `"string"`. Required.

Name of the Nginx Service Discovery. Must be unique within the namespace.

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

<a id="canonical-2002101032032131-0002012020023311-0312222013222110-0211302012111013-2202123331030122-1231200011131113-3012121130002210-0233222001320301"></a>

<a id="canonical-3132220001322232-3023113321133012-2221221211331113-0023030103013120-0222303021313133-1331122201010103-1131023323303133-2023000020232323"></a>

## namespace property — Property reference / 101313032012 / 10

Type: `"string"`. Required.

Namespace where the Nginx Service Discovery is created.

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

- [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-2222022131131122-3312320112111130-2002010102031033-1121330200131200-3323032021213333-3303212210100132-0101031000301200-2230133310332123): complete subsection reference.

- [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-3332130200322300-2012231113111011-2302332300331203-3311101201333000-2312033011013210-2320322110020100-1123311333112313-3033311321013301): complete subsection reference.

<a id="canonical-3211311313301003-1120012122112311-2000111013111313-2231100020110021-3111303033023113-1300320232303233-1211031030221121-2132011320110220"></a>

## All schema paths — Property reference / 101313032012 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nginx_service_discovery--reference--group-001.md#canonical-1021313222322213-3130220330113323-2312303021122322-3032021321222011-1021003010220020-3232111022103010-1212100003203122-0233321132301010) |
| `description` | [description](resources--nginx_service_discovery--reference--group-001.md#canonical-3020022132232010-1101331011103113-3210002123002121-3303001330032220-1320121232203111-2120223213132221-2132200101003023-1313213333200312) |
| `disable` | [disable](resources--nginx_service_discovery--reference--group-001.md#canonical-2233002110022312-0313032310102320-0333212123331330-0231101211220221-2201033012220113-2012013003120223-0102312213033033-0300000213130300) |
| `discovery_target` | [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2022331312010000-2320111232000323-2222123032313012-1100020321313132-3102332003321122-0113023121213001-1103223303310323-1000213132212213) |
| `discovery_target.config_sync_group` | [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-0000113122231230-3331130020203110-3310132332332311-0221222332211100-3133210020323303-2122310022201103-1311220032020301-2313221302323032) |
| `discovery_target.config_sync_group.config_sync_group` | [discovery_target.config_sync_group.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-0112331002023033-2202221012331133-2030320201332233-3211212320222302-1301110302230012-1120321131122021-0113330232023032-0211311231213331) |
| `discovery_target.config_sync_group.config_sync_group.kind` | [discovery_target.config_sync_group.config_sync_group.kind](resources--nginx_service_discovery--reference--group-001.md#canonical-0213010300011013-3010100311002310-0033221331001011-3110222302132221-3022221011223103-1221030223333312-0320212321012313-2212320130121330) |
| `discovery_target.config_sync_group.config_sync_group.name` | [discovery_target.config_sync_group.config_sync_group.name](resources--nginx_service_discovery--reference--group-001.md#canonical-0213212033121003-0300013012003300-0000332013211110-1232312101300230-3023222002013333-3211133120020032-2200220020313222-1031022221201321) |
| `discovery_target.config_sync_group.config_sync_group.namespace` | [discovery_target.config_sync_group.config_sync_group.namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-1331130022213331-3100203003131233-2323233130231012-3003233231010033-1312111313102222-3130311333312300-0203303212223300-3323320311013130) |
| `discovery_target.config_sync_group.config_sync_group.tenant` | [discovery_target.config_sync_group.config_sync_group.tenant](resources--nginx_service_discovery--reference--group-001.md#canonical-0200112120320210-1102112330102212-3313320030212102-3012000120323102-2122110020311201-0202122223313121-1201023331113132-3231103130301201) |
| `discovery_target.config_sync_group.config_sync_group.uid` | [discovery_target.config_sync_group.config_sync_group.uid](resources--nginx_service_discovery--reference--group-001.md#canonical-1202110231222222-2011230331122123-1232001311221010-0332211211111232-2123001333320023-2331302002311233-1120201213032200-3123211030003322) |
| `discovery_target.nginx_instance` | [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-1320110011322101-0022301323313120-3111332302320213-3231111031000012-0001213031023230-1301202201333213-2000012020201113-3110133021230023) |
| `discovery_target.nginx_instance.nginx_instance` | [discovery_target.nginx_instance.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3220002121311012-1203130230003121-3123232210010002-2111121101202322-1133012201222033-2133323323301120-2133312120312311-3102313212322132) |
| `discovery_target.nginx_instance.nginx_instance.kind` | [discovery_target.nginx_instance.nginx_instance.kind](resources--nginx_service_discovery--reference--group-001.md#canonical-2311233212232003-1302020323003320-0330201211000311-3201021223121302-3213323023033321-3210320331233133-3120132323230310-3102210222112211) |
| `discovery_target.nginx_instance.nginx_instance.name` | [discovery_target.nginx_instance.nginx_instance.name](resources--nginx_service_discovery--reference--group-001.md#canonical-2131112333332233-3313032232321333-0131023230023031-2122013002000212-3213231232002011-2300311231311313-2313122132133213-2120311223013102) |
| `discovery_target.nginx_instance.nginx_instance.namespace` | [discovery_target.nginx_instance.nginx_instance.namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-3130131001303233-1022211120211221-2110330123221133-0133122001100222-2311300322320111-0113032333121131-2102011321023121-2011220212231320) |
| `discovery_target.nginx_instance.nginx_instance.tenant` | [discovery_target.nginx_instance.nginx_instance.tenant](resources--nginx_service_discovery--reference--group-001.md#canonical-1300012101322311-3012100331213003-3130121212230331-2131331130013220-3010111210111310-3011201331111101-3220213030332001-3031022021113001) |
| `discovery_target.nginx_instance.nginx_instance.uid` | [discovery_target.nginx_instance.nginx_instance.uid](resources--nginx_service_discovery--reference--group-001.md#canonical-1032130002111133-2331200132232200-1222023201011331-2033011313222301-1230230200003311-3213011033302132-3130320020322222-3313222312202213) |
| `id` | [ID](resources--nginx_service_discovery--reference--group-001.md#canonical-3302231223210022-2133312233002002-2131120030233312-0203221310232111-3230010201112112-0333122030222200-2121013020121222-0030330012130322) |
| `labels` | [labels](resources--nginx_service_discovery--reference--group-001.md#canonical-1023022231013231-2303110301033023-3213101202131320-0221221030102331-1231100122031121-2223102131121301-1303123312000310-3220110012132201) |
| `name` | [name](resources--nginx_service_discovery--reference--group-001.md#canonical-1121002300200112-2101120032201212-2210103121100233-1211003312030221-1201233111232330-2222002331031110-1000030330031013-2003233003012023) |
| `namespace` | [namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-2002101032032131-0002012020023311-0312222013222110-0211302012111013-2202123331030122-1231200011131113-3012121130002210-0233222001320301) |
| `server_block_filters` | [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-3020122212222123-3301322201302030-3103120202332222-3333000112233312-2332333003023013-3100322001021121-0222000310220210-1031123110302022) |
| `server_block_filters.name_regex` | [server_block_filters.name_regex](resources--nginx_service_discovery--reference--group-001.md#canonical-2113012313232200-0000303121210103-2210230320222300-2020320233003133-2312003220131213-0311303122011020-2013221203002000-3000333020010031) |
| `server_block_filters.port_ranges` | [server_block_filters.port_ranges](resources--nginx_service_discovery--reference--group-001.md#canonical-1323120310312232-0113002131111120-0312011321322220-2201101230203131-2230213303302300-1130322232213233-1100231331121122-1302212123331201) |
| `timeouts` | [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-3100333022112002-2010200132313011-0030022133103121-0111222303211121-2000330020003123-2102230321333311-0133202321112131-1330323120020231) |
| `timeouts.create` | [timeouts.create](resources--nginx_service_discovery--reference--group-001.md#canonical-2330032321210203-2211010111302113-3122020123010002-3131321301111311-3122120322100313-3321321302113132-2201322232001102-3202110110320020) |
| `timeouts.delete` | [timeouts.delete](resources--nginx_service_discovery--reference--group-001.md#canonical-1000221201331212-2230103301132301-1202001022230320-1213210302230230-3330013121033123-3201033320220122-1100121333030031-3232110111323213) |
| `timeouts.read` | [timeouts.read](resources--nginx_service_discovery--reference--group-001.md#canonical-1320011113320130-0122131222013120-3200030331030021-0133030102321311-1233203332213210-3300003212030002-2100012233003112-0232013130010200) |
| `timeouts.update` | [timeouts.update](resources--nginx_service_discovery--reference--group-001.md#canonical-3223120021311031-1301311303330210-0001133233213032-0120200123313312-2001101311303023-3023120122100013-1321012321211113-0231030233303001) |

<a id="canonical-0103130123123013-3301333210020213-3013210120112303-2100210302330133-0113111102132033-0321310100202132-0102303303232132-1323032002023201"></a>

## Next pages — Property reference / 101313032012 / 12

- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-2222022131131122-3312320112111130-2002010102031033-1121330200131200-3323032021213333-3303212210100132-0101031000301200-2230133310332123)
- [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-3332130200322300-2012231113111011-2302332300331203-3311101201333000-2312033011013210-2320322110020100-1123311333112313-3033311321013301)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320303000012001-3100323302321331-1113331023120022-3223110120203003-1121303120233031-3233221332332320-3232203122001321-0220001123030312"></a>

## discovery_target — discovery_target / 200030001330 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- discovery_target

<a id="canonical-2022331312010000-2320111232000323-2222123032313012-1100020321313132-3102332003321122-0113023121213001-1103223303310323-1000213132212213"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("config_sync_group",
    "nginx_instance")}
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
  "x-ves-oneof-field-target": "[\"config_sync_group\",\"nginx_instance\"]"
}
```

Terraform syntax:

```terraform
discovery_target {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011213130312310-3321302313120213-0110132122113222-2102211131122331-0003110010331101-1213111122303003-2121122313220231-3113123001103211"></a>

## Direct properties — discovery_target / 200030001330 / 3

- [config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122): complete subsection reference.

- [nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133): complete subsection reference.

<a id="canonical-1020103300020102-2200100001121011-3121313201301301-2012221323303311-3123210332333323-3002123202323022-3331213100322222-2313312231133023"></a>

## Next pages — discovery_target / 200030001330 / 4

- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122)
- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321302223232013-2201313021103001-1011233202212313-3201103001220100-1230312322033031-2202033020130231-1032311312032212-1032031332301213"></a>

## discovery_target.config_sync_group — config_sync_group / 000100302103 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- discovery_target.config_sync_group

<a id="canonical-0000113122231230-3331130020203110-3310132332332311-0221222332211100-3133210020323303-2122310022201103-1311220032020301-2313221302323032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for config sync group.

Upstream description:

Select new ConfigSyncGroup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("config_sync_group")}
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
config_sync_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302013322302303-0231311012303132-0201033332322312-1033331030121110-1100011031130003-0222101032112322-1102213331100110-1032033222222120"></a>

## Direct properties — config_sync_group / 000100302103 / 3

- [config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-0132230021310122-3103332103310000-2112033023321020-3323003103122001-0021000102032233-0210300202232130-2312102030000031-2230211132303030): complete subsection reference.

<a id="canonical-2332132231201102-2120000113322300-3101331033231123-0311231131013231-0113111003101102-0300100102230112-0122032223103033-1113022212203200"></a>

## Next pages — config_sync_group / 000100302103 / 4

- [discovery_target.config_sync_group.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-0132230021310122-3103332103310000-2112033023321020-3323003103122001-0021000102032233-0210300202232130-2312102030000031-2230211132303030)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-0132230021310122-3103332103310000-2112033023321020-3323003103122001-0021000102032233-0210300202232130-2312102030000031-2230211132303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003201033031211-0001130313302002-2211103231311020-2310301022322100-3232120330033130-0310223012200313-1311300110333230-1021303203110303"></a>

## discovery_target.config_sync_group.config_sync_group — config_sync_group / 003311003321 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122)
- discovery_target.config_sync_group.config_sync_group

<a id="canonical-0112331002023033-2202221012331133-2030320201332233-3211212320222302-1301110302230012-1120321131122021-0113330232023032-0211311231213331"></a>

Type: `"object"`. list nested block, Optional.

Reference. Select new ConfigSyncGroup.

Upstream description:

Select new ConfigSyncGroup.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
config_sync_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230331013121232-2312300023130102-1220103213201132-0301230011221333-0102303212020331-0110032021010213-2123221320000031-2122013102123003"></a>

## Direct properties — config_sync_group / 003311003321 / 3

<a id="canonical-0213010300011013-3010100311002310-0033221331001011-3110222302132221-3022221011223103-1221030223333312-0320212321012313-2212320130121330"></a>

<a id="canonical-2110213120323133-0220012011110321-2002102311323133-1331300330001220-2230331012313003-1022203220003033-3332320112303103-1313210332002023"></a>

## kind property — config_sync_group / 003311003321 / 4

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

<a id="canonical-0213212033121003-0300013012003300-0000332013211110-1232312101300230-3023222002013333-3211133120020032-2200220020313222-1031022221201321"></a>

<a id="canonical-1103212320120022-3003220010201302-1222031322023301-0201023310123310-0313212222132222-0030330220222223-1133103130212000-2001123230102010"></a>

## name property — config_sync_group / 003311003321 / 5

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

<a id="canonical-1331130022213331-3100203003131233-2323233130231012-3003233231010033-1312111313102222-3130311333312300-0203303212223300-3323320311013130"></a>

<a id="canonical-2030201002311223-0133232013221013-0133320300302131-3231012111123333-0203100222302312-0113122123321023-2312331320321230-0130332301033223"></a>

## namespace property — config_sync_group / 003311003321 / 6

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

<a id="canonical-0200112120320210-1102112330102212-3313320030212102-3012000120323102-2122110020311201-0202122223313121-1201023331113132-3231103130301201"></a>

<a id="canonical-0210102222100031-1221201032321021-2030113231313031-3030203202100103-3100221032132112-2101103002301100-3112120021000011-2012003333223001"></a>

## tenant property — config_sync_group / 003311003321 / 7

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

<a id="canonical-1202110231222222-2011230331122123-1232001311221010-0332211211111232-2123001333320023-2331302002311233-1120201213032200-3123211030003322"></a>

<a id="canonical-3210220010223001-0222112021201320-2011222210222323-1233133103233200-1220021321322011-1133010213031331-2220023233210222-2232103113201302"></a>

## uid property — config_sync_group / 003311003321 / 8

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

<a id="canonical-2222323232202003-2333112330012122-3313033330020231-0102023230030101-0302232023010211-1103011120121211-3332002132111101-0133000332131231"></a>

## Next pages — config_sync_group / 003311003321 / 9

- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-2011212030103323-1211101011023030-1210222010122301-0322112032010033-1113232200111120-0030013221311330-3123122103230032-1221201330220122)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003330222330220-3110232022101003-2020011312033220-3221333000201323-2210303123121302-1011321311033002-1002010011223313-1003112122110230"></a>

## discovery_target.nginx_instance — nginx_instance / 001323020021 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- discovery_target.nginx_instance

<a id="canonical-1320110011322101-0022301323313120-3111332302320213-3231111031000012-0001213031023230-1301202201333213-2000012020201113-3110133021230023"></a>

Type: `"object"`. single nested block, Optional.

NGINXInstance Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nginx_instance")}
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
nginx_instance {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102211203013102-2223023100021123-2300032002032001-2010121330032123-0122000113123230-3203002323212030-1323133013003222-1312131312211320"></a>

## Direct properties — nginx_instance / 001323020021 / 3

- [nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-0330103203002321-2201120233130313-1222012203203013-0233221012200212-0130121131202133-2333331001203333-0210333221031012-0133202011221210): complete subsection reference.

<a id="canonical-3002222021302230-1001103012303031-2022223123220132-0002303113010222-3021313220000030-1013000312212022-2303330230112013-1231010011232312"></a>

## Next pages — nginx_instance / 001323020021 / 4

- [discovery_target.nginx_instance.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-0330103203002321-2201120233130313-1222012203203013-0233221012200212-0130121131202133-2333331001203333-0210333221031012-0133202011221210)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-0330103203002321-2201120233130313-1222012203203013-0233221012200212-0130121131202133-2333331001203333-0210333221031012-0133202011221210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103300121323101-1332102213110101-1022011230211311-3311313232110200-0311123330231230-1331030310011113-3031231300210023-1031200113322020"></a>

## discovery_target.nginx_instance.nginx_instance — nginx_instance / 131211323213 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-0200111022000313-3021311210312113-2123133211211313-0300212021112210-0021332122331022-2330302200002331-0210301332130032-1132030132233312)
- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133)
- discovery_target.nginx_instance.nginx_instance

<a id="canonical-3220002121311012-1203130230003121-3123232210010002-2111121101202322-1133012201222033-2133323323301120-2133312120312311-3102313212322132"></a>

Type: `"object"`. list nested block, Optional.

Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
nginx_instance {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010010322202310-2130211331302201-0100323333123320-2133323211003311-2110131033013312-1231123133302101-0013330213020132-0022211033103231"></a>

## Direct properties — nginx_instance / 131211323213 / 3

<a id="canonical-2311233212232003-1302020323003320-0330201211000311-3201021223121302-3213323023033321-3210320331233133-3120132323230310-3102210222112211"></a>

<a id="canonical-2113101021231320-0120030232120203-0112332100223231-3311323202302212-1123012201010222-0101230111100320-2022223211103012-3012222210102012"></a>

## kind property — nginx_instance / 131211323213 / 4

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

<a id="canonical-2131112333332233-3313032232321333-0131023230023031-2122013002000212-3213231232002011-2300311231311313-2313122132133213-2120311223013102"></a>

<a id="canonical-0123320231032233-2220121103133233-2311111210100003-0322133232233233-0013232302110000-2110233210011302-0331120203130123-0331123001310020"></a>

## name property — nginx_instance / 131211323213 / 5

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

<a id="canonical-3130131001303233-1022211120211221-2110330123221133-0133122001100222-2311300322320111-0113032333121131-2102011321023121-2011220212231320"></a>

<a id="canonical-0123032332101202-0103030021323023-0200332120011302-3031032213201222-3210122012031123-0322131233331131-0223310332011331-3222211021201302"></a>

## namespace property — nginx_instance / 131211323213 / 6

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

<a id="canonical-1300012101322311-3012100331213003-3130121212230331-2131331130013220-3010111210111310-3011201331111101-3220213030332001-3031022021113001"></a>

<a id="canonical-2132003310221223-2102321201003110-2103203320000212-2331323031012320-3231111223200313-3120002011221331-1333221033032130-2133302333230032"></a>

## tenant property — nginx_instance / 131211323213 / 7

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

<a id="canonical-1032130002111133-2331200132232200-1222023201011331-2033011313222301-1230230200003311-3213011033302132-3130320020322222-3313222312202213"></a>

<a id="canonical-1013110330023110-2333233233212332-2122211113333321-1222321301012210-1002010233020303-1312023321231121-3031203023331223-0333013323221022"></a>

## uid property — nginx_instance / 131211323213 / 8

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

<a id="canonical-2222031120303202-3000122311332323-1320111302311013-2201011131200001-1001100003230212-1111030310332230-0210122133002122-1013002213112212"></a>

## Next pages — nginx_instance / 131211323213 / 9

- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3002103300302203-3233002012210030-2131002122302221-0203011202221231-3100111103220331-1220330203313221-0322022212312332-0230302020002133)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-2222022131131122-3312320112111130-2002010102031033-1121330200131200-3323032021213333-3303212210100132-0101031000301200-2230133310332123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310032131032020-0310211003110330-0020012131000300-1112110212321000-3031233103202222-0332330322331001-1011230333230210-2102310221122223"></a>

## server_block_filters — server_block_filters / 321332210030 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- server_block_filters

<a id="canonical-3020122212222123-3301322201302030-3103120202332222-3333000112233312-2332333003023013-3100322001021121-0222000310220210-1031123110302022"></a>

Type: `"object"`. list nested block, Optional.

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks
will be discovered by default.

Upstream description:

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter.

X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_block_filters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231311020300231-2101003033330312-0022202230131021-2222202320323013-0211023302010200-1102012231133331-0312123112203310-0113130120120311"></a>

## Direct properties — server_block_filters / 321332210030 / 3

<a id="canonical-2113012313232200-0000303121210103-2210230320222300-2020320233003133-2312003220131213-0311303122011020-2013221203002000-3000333020010031"></a>

<a id="canonical-2110302312101310-2021003111032212-3011323300222022-0330111332132200-3013123020211221-2330012000132302-1212111322210322-3220320121113102"></a>

## name_regex property — server_block_filters / 321332210030 / 4

Type: `"string"`. Optional.

Regular expression to match the server name or domain that must be discovered.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1323120310312232-0113002131111120-0312011321322220-2201101230203131-2230213303302300-1130322232213233-1100231331121122-1302212123331201"></a>

<a id="canonical-3132032003020122-2311321212300330-0001322212102311-3311102120303220-3203133101020102-0121302122302333-1211021213022101-1100332231011300"></a>

## port_ranges property — server_block_filters / 321332210030 / 5

Type: `"string"`. Optional.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Upstream description:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". For example, 8000-8191. Maximum
number of ports allowed is 1024.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-0130030021331211-3321102112012200-1132031312313231-3330330022203001-0123120023331030-1212020222320101-1121101300201033-0102002130210200"></a>

## Next pages — server_block_filters / 321332210030 / 6

- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)

<a id="canonical-3332130200322300-2012231113111011-2302332300331203-3311101201333000-2312033011013210-2320322110020100-1123311333112313-3033311321013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302302312233321-0122120303023303-1203012031133021-3232101003013021-2223102111130011-3012001133210003-1322031123300200-1113313302333213"></a>

## timeouts — timeouts / 301320101230 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- timeouts

<a id="canonical-3100333022112002-2010200132313011-0030022133103121-0111222303211121-2000330020003123-2102230321333311-0133202321112131-1330323120020231"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210121201001213-1123113012122203-2113302112133312-2231102331120132-2303130230103002-2021312212210332-0032021331300322-0021332222031301"></a>

## Direct properties — timeouts / 301320101230 / 3

<a id="canonical-2330032321210203-2211010111302113-3122020123010002-3131321301111311-3122120322100313-3321321302113132-2201322232001102-3202110110320020"></a>

<a id="canonical-3212210022130030-3002310011101003-2031120330320012-2103300111033311-3012301100003101-2213020031111132-1022000300322211-2133333301103122"></a>

## create property — timeouts / 301320101230 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1000221201331212-2230103301132301-1202001022230320-1213210302230230-3330013121033123-3201033320220122-1100121333030031-3232110111323213"></a>

<a id="canonical-3012330021223002-1123010232331020-2120222003212022-0333011202000201-0113110223301202-1103222233102010-1032300031023133-1032211032323200"></a>

## delete property — timeouts / 301320101230 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1320011113320130-0122131222013120-3200030331030021-0133030102321311-1233203332213210-3300003212030002-2100012233003112-0232013130010200"></a>

<a id="canonical-3302201113011301-3221311021312312-1122202000120311-3233100203003201-3023332332302331-2332031230031021-2303311203030203-3112313022121322"></a>

## read property — timeouts / 301320101230 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3223120021311031-1301311303330210-0001133233213032-0120200123313312-2001101311303023-3023120122100013-1321012321211113-0231030233303001"></a>

<a id="canonical-3301103333313310-0333320203303203-0232312203313323-2030013032000203-3230211223210320-2202211331020100-3223023011020200-1032133123131012"></a>

## update property — timeouts / 301320101230 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0330333320122033-1200033210233002-3311312333320023-2330001030211300-0003023030232212-1010211002303100-2202133203330233-1202122130011211"></a>

## Next pages — timeouts / 301320101230 / 8

- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
