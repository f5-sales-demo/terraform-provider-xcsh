---
page_title: "xcsh_nginx_service_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery reference."
---

# xcsh_nginx_service_discovery reference

<a id="canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- Property reference

<a id="canonical-1331230230032031-2302023110220331-2102203321203302-3211003230300121-0202000101302001-2322022001101131-1221120230113232-1110021200231011"></a>

### Direct properties for `xcsh_nginx_service_discovery`

<a id="canonical-0002122230222003-1022110003212331-0212210312200002-0110321031011200-2221103010302012-0313211203023131-1210131233312122-3211112003210122"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-2211131110111210-0110212101303300-0301100310131221-0122100013232022-0232120331011012-1021002120220001-2102211331233111-1032103310112010"></a>

<a id="canonical-1020321132002322-2333230023223233-2303300210330320-3222132021012002-0331003003122111-2313122111221113-0201001321022203-2033300322020203"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NginxServiceDiscovery.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211): complete subsection reference.

<a id="canonical-1112123103000031-0000201010200032-1002330231003003-0101331103122223-3233002012120210-0012322213232030-0123100031002213-1110320001200133"></a>

<a id="canonical-3101011200101022-1310113103233230-3221310022001202-0113113020331122-2312222331313021-3310022032230113-1221330112101300-3331002211233221"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0132102110300322-3122103113021132-0310230021100122-1301102210311321-0123310113121130-1210332022202333-2212313303211001-2003313113310133"></a>

<a id="canonical-2320001311303323-2330311313130112-3100211100113233-2303312320223232-0122000012313323-3221003310133012-0131030210132032-3022311103211301"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

<a id="canonical-1203011130000002-2030211013301202-3323110133201233-2030311332011123-3313311033031220-3232121231001020-2023032301102221-3003320301123322"></a>

<a id="canonical-2303113320121033-2310313223010323-0132120103123223-2203033330110303-0113011130213231-1131230111120003-3221102310023203-2031021002233321"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NginxServiceDiscovery.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2232212111223022-3030313203320222-2011321031130300-0200211132232213-0301001021010231-0033222000012110-2122002113210223-2123303230122230"></a>

<a id="canonical-1000212020002033-3120133321233310-0312331220000302-2213001123112101-2231302332120202-2221022100221330-0223231233110110-1313333330310300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the NginxServiceDiscovery exists.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [server_block_filters](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3301113020113113-0301200303022012-1223200003223313-2333133121110221-3213132022300322-3323132030003003-2202003121023320-1113212320113302): complete subsection reference.

<a id="canonical-1333222210313010-3223003003331013-0301030032210002-2130030010220103-3220323321132032-3330122103320121-0112101221120010-1020103212301210"></a>

### All schema paths for `xcsh_nginx_service_discovery`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0002122230222003-1022110003212331-0212210312200002-0110321031011200-2221103010302012-0313211203023131-1210131233312122-3211112003210122) |
| `description` | [description](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2211131110111210-0110212101303300-0301100310131221-0122100013232022-0232120331011012-1021002120220001-2102211331233111-1032103310112010) |
| `discovery_target` | [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3133110032212122-3121011221010311-3103013212111303-2322223001300310-1020112023311111-1210111200303311-3101011013222013-0303110121212033) |
| `discovery_target.config_sync_group` | [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0112121132211101-0331001112202133-3221132230101023-2232013131112310-1133303200333130-2333002301010203-1031022220102312-0011130322112130) |
| `discovery_target.config_sync_group.config_sync_group` | [discovery_target.config_sync_group.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3131220222131311-3010201003100302-2113030102331001-2313000320233123-1102232121321233-0113112102133010-1131300231131332-1313023000332321) |
| `discovery_target.config_sync_group.config_sync_group.kind` | [discovery_target.config_sync_group.config_sync_group.kind](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3331111113031301-3222313133110231-2011203101332232-0303120301210332-1213031223323312-3101300013301331-3211121033230133-1202323232213303) |
| `discovery_target.config_sync_group.config_sync_group.name` | [discovery_target.config_sync_group.config_sync_group.name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2301003110122022-1333321010123230-2031001211101322-0221222032230231-0213023001231112-1030012011323021-3001303311113200-0122023323300002) |
| `discovery_target.config_sync_group.config_sync_group.namespace` | [discovery_target.config_sync_group.config_sync_group.namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2100133321021101-1300110212022332-2222333120123021-2001111222003001-1001310203313111-3323023121003331-2210231011223212-1131001033103020) |
| `discovery_target.config_sync_group.config_sync_group.tenant` | [discovery_target.config_sync_group.config_sync_group.tenant](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3331012230112112-3221003200032332-3213220332313203-0203110003002132-3112311001020321-1021000033100003-1210201001122313-2010103123230123) |
| `discovery_target.config_sync_group.config_sync_group.uid` | [discovery_target.config_sync_group.config_sync_group.uid](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0233300000313031-3003102122302211-0021021000101310-1001211301121022-0100022020110130-2112223313113003-3311231302032011-3221210002113010) |
| `discovery_target.nginx_instance` | [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0223023122200020-1123122201012103-0203220002122010-3103023003123033-2100130310311333-1222320033010020-2013302323101111-1110233310303001) |
| `discovery_target.nginx_instance.nginx_instance` | [discovery_target.nginx_instance.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2103221321200103-1212212330320321-1332223102013023-2100011222210010-1223323021230110-2000220133131031-1303001311310032-1033233010032110) |
| `discovery_target.nginx_instance.nginx_instance.kind` | [discovery_target.nginx_instance.nginx_instance.kind](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0202313132230222-1132312003003001-0332201211300012-2113231213321031-1012110033303112-2011030232110330-3220003111330123-0120010321001300) |
| `discovery_target.nginx_instance.nginx_instance.name` | [discovery_target.nginx_instance.nginx_instance.name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0131131223232332-3031331130113333-0033020332003203-1011132012321031-1023130303330033-3101220122102313-3230133130002012-3211120113311303) |
| `discovery_target.nginx_instance.nginx_instance.namespace` | [discovery_target.nginx_instance.nginx_instance.namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3332130211121221-2313230222202013-1103100203031031-2322100322321221-2000220303021333-0033011022220110-1211023323221233-2310002133331010) |
| `discovery_target.nginx_instance.nginx_instance.tenant` | [discovery_target.nginx_instance.nginx_instance.tenant](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0132031030320132-0010330133100321-2101000011023302-0131021122013302-0101020010113112-0121002213303312-0101001120110302-0303031210202332) |
| `discovery_target.nginx_instance.nginx_instance.uid` | [discovery_target.nginx_instance.nginx_instance.uid](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0110232111010322-3001312131011302-0121200023303323-0122322231133013-2300112332100020-0303110012303222-2103011221100131-2010300310232330) |
| `id` | [ID](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1112123103000031-0000201010200032-1002330231003003-0101331103122223-3233002012120210-0012322213232030-0123100031002213-1110320001200133) |
| `labels` | [labels](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0132102110300322-3122103113021132-0310230021100122-1301102210311321-0123310113121130-1210332022202333-2212313303211001-2003313113310133) |
| `name` | [name](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1203011130000002-2030211013301202-3323110133201233-2030311332011123-3313311033031220-3232121231001020-2023032301102221-3003320301123322) |
| `namespace` | [namespace](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2232212111223022-3030313203320222-2011321031130300-0200211132232213-0301001021010231-0033222000012110-2122002113210223-2123303230122230) |
| `server_block_filters` | [server_block_filters](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2020303002011313-2102331303233232-2201002201313033-2131000101121103-0302321200312003-0103330102120002-3301122010011213-3331121100213112) |
| `server_block_filters.name_regex` | [server_block_filters.name_regex](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3211310110210133-3200122133303302-3211002223212210-2312113013103321-0321312233013031-1212102131233100-2102130231110222-2130012202312111) |
| `server_block_filters.port_ranges` | [server_block_filters.port_ranges](data-sources--nginx_service_discovery--reference--group-001.md#canonical-3333211202200010-0221001202021110-2313211120120310-1013010331312330-3202111132223201-1222300233201303-2322021223012112-0133332320103023) |

<a id="canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_target` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- discovery_target

<a id="canonical-3133110032212122-3121011221010311-3103013212111303-2322223001300310-1020112023311111-1210111200303311-3101011013222013-0303110121212033"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery target.

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

<a id="canonical-3331212233112032-1000222310132211-2022212033122113-3232013311303132-2122020331311313-1212221313332121-3220021113123002-3212203130223023"></a>

### Direct properties for `discovery_target`

- [config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1121200212020111-0103012222000111-2012131130013021-0123231212310310-3021032022021223-0201223001210012-2110003321002012-2202001121022221): complete subsection reference.

- [nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0300232101130030-0001232300211210-0303010113013301-3322131213033201-2233001322220221-2202230200021122-1232030300100130-1303110002110131): complete subsection reference.

<a id="canonical-1121200212020111-0103012222000111-2012131130013021-0123231212310310-3021032022021223-0201223001210012-2110003321002012-2202001121022221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_target.config_sync_group` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211)
- discovery_target.config_sync_group

<a id="canonical-0112121132211101-0331001112202133-3221132230101023-2232013131112310-1133303200333130-2333002301010203-1031022220102312-0011130322112130"></a>

Type: `"single"`. Computed.

Configuration parameter for config sync group.

Additional upstream details:

Select new ConfigSyncGroup.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313020220010123-0121300303033111-0023323023323303-3020133213322031-2012003123112130-2200223121031313-0120131310332331-1313132322021303"></a>

### Direct properties for `discovery_target.config_sync_group`

- [config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1202011210031202-0031323121201132-3312110021112101-0330123302103021-0020123113102303-2332223222031122-1023030320311212-2313133131220033): complete subsection reference.

<a id="canonical-1202011210031202-0031323121201132-3312110021112101-0330123302103021-0020123113102303-2332223222031122-1023030320311212-2313133131220033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_target.config_sync_group.config_sync_group` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211)
- [discovery_target.config_sync_group](data-sources--nginx_service_discovery--reference--group-001.md#canonical-1121200212020111-0103012222000111-2012131130013021-0123231212310310-3021032022021223-0201223001210012-2110003321002012-2202001121022221)
- discovery_target.config_sync_group.config_sync_group

<a id="canonical-3131220222131311-3010201003100302-2113030102331001-2313000320233123-1102232121321233-0113112102133010-1131300231131332-1313023000332321"></a>

Type: `"list"`. Computed.

Reference. Select new ConfigSyncGroup.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300213312113100-1301302211112030-1022223300100101-3030201120220101-0203033211312120-2232301100200103-1233301003220231-2221213130302103"></a>

### Direct properties for `discovery_target.config_sync_group.config_sync_group`

<a id="canonical-3331111113031301-3222313133110231-2011203101332232-0303120301210332-1213031223323312-3101300013301331-3211121033230133-1202323232213303"></a>

#### `discovery_target.config_sync_group.config_sync_group.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2301003110122022-1333321010123230-2031001211101322-0221222032230231-0213023001231112-1030012011323021-3001303311113200-0122023323300002"></a>

<a id="canonical-2333031231000312-3101030020202202-0012210132132201-3023210202223300-0230021130000330-1230220131111120-3300203203010020-3122201323312210"></a>

#### `discovery_target.config_sync_group.config_sync_group.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2100133321021101-1300110212022332-2222333120123021-2001111222003001-1001310203313111-3323023121003331-2210231011223212-1131001033103020"></a>

<a id="canonical-3030301330013021-2001021033102021-2033010301100311-3212012121232101-0201012120021203-1233100230201211-1210321323010232-3331311313322103"></a>

#### `discovery_target.config_sync_group.config_sync_group.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3331012230112112-3221003200032332-3213220332313203-0203110003002132-3112311001020321-1021000033100003-1210201001122313-2010103123230123"></a>

<a id="canonical-0120312022113330-1020123200121103-0220103020032113-3112221111023323-2311302020300020-2223132003022310-0111021100210322-3132100111103200"></a>

#### `discovery_target.config_sync_group.config_sync_group.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0233300000313031-3003102122302211-0021021000101310-1001211301121022-0100022020110130-2112223313113003-3311231302032011-3221210002113010"></a>

<a id="canonical-0333233212012002-2222313013120031-3322132002022031-0013210001300012-2222101331003101-1310112322223021-3320333322212222-2310213233132131"></a>

#### `discovery_target.config_sync_group.config_sync_group.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0300232101130030-0001232300211210-0303010113013301-3322131213033201-2233001322220221-2202230200021122-1232030300100130-1303110002110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_target.nginx_instance` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211)
- discovery_target.nginx_instance

<a id="canonical-0223023122200020-1123122201012103-0203220002122010-3103023003123033-2100130310311333-1222320033010020-2013302323101111-1110233310303001"></a>

Type: `"single"`. Computed.

NGINXInstance Reference. Select new NGINX Instance.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0310133311012132-1313121122001223-2302001130033103-2323300321000301-2022323111033220-0102123330222301-0330001203301022-1010022311211131"></a>

### Direct properties for `discovery_target.nginx_instance`

- [nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2021222132230222-1310130221020031-2110122030333130-3320321230100233-1130103301030022-0112212320211032-0011221121303313-3322020320011323): complete subsection reference.

<a id="canonical-2021222132230222-1310130221020031-2110122030333130-3320321230100233-1130103301030022-0112212320211032-0011221121303313-3322020320011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_target.nginx_instance.nginx_instance` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- [discovery_target](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0220312301013130-1323112333310301-1023311301010200-1111320222132102-2112102102200100-3200110003311031-3031212331233013-0020102033031211)
- [discovery_target.nginx_instance](data-sources--nginx_service_discovery--reference--group-001.md#canonical-0300232101130030-0001232300211210-0303010113013301-3322131213033201-2233001322220221-2202230200021122-1232030300100130-1303110002110131)
- discovery_target.nginx_instance.nginx_instance

<a id="canonical-2103221321200103-1212212330320321-1332223102013023-2100011222210010-1223323021230110-2000220133131031-1303001311310032-1033233010032110"></a>

Type: `"list"`. Computed.

Reference. Select new NGINX Instance.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2301333300231301-3130210330003132-3003301103033013-3323103013033111-3010103302323021-0120111210130000-0330131120211010-3312310121001132"></a>

### Direct properties for `discovery_target.nginx_instance.nginx_instance`

<a id="canonical-0202313132230222-1132312003003001-0332201211300012-2113231213321031-1012110033303112-2011030232110330-3220003111330123-0120010321001300"></a>

#### `discovery_target.nginx_instance.nginx_instance.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0131131223232332-3031331130113333-0033020332003203-1011132012321031-1023130303330033-3101220122102313-3230133130002012-3211120113311303"></a>

<a id="canonical-1103121131110022-0020121010231300-0001230030133201-3012133130111101-2222100220310110-1110210223021213-0222121021200111-0323122103331122"></a>

#### `discovery_target.nginx_instance.nginx_instance.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3332130211121221-2313230222202013-1103100203031031-2322100322321221-2000220303021333-0033011022220110-1211023323221233-2310002133331010"></a>

<a id="canonical-1132323300331302-2331102300031021-2230221220122232-1112220031203012-3302012201033301-1321230313123330-3232021322110332-2123002300002121"></a>

#### `discovery_target.nginx_instance.nginx_instance.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0132031030320132-0010330133100321-2101000011023302-0131021122013302-0101020010113112-0121002213303312-0101001120110302-0303031210202332"></a>

<a id="canonical-1322111120132003-1032331311022002-3302330103010311-1301232311102031-0322130013302002-3112210231303132-0123300011221120-2001323232012030"></a>

#### `discovery_target.nginx_instance.nginx_instance.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0110232111010322-3001312131011302-0121200023303323-0122322231133013-2300112332100020-0303110012303222-2103011221100131-2010300310232330"></a>

<a id="canonical-2013310220100313-3102133121012110-0111132301331321-0132003100301312-3310100103110331-2123013133332301-3300133012310230-2111111023000310"></a>

#### `discovery_target.nginx_instance.nginx_instance.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3301113020113113-0301200303022012-1223200003223313-2333133121110221-3213132022300322-3323132030003003-2202003121023320-1113212320113302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `server_block_filters` properties

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Property reference](data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- server_block_filters

<a id="canonical-2020303002011313-2102331303233232-2201002201313033-2131000101121103-0302321200312003-0103330102120002-3301122010011213-3331121100213112"></a>

Type: `"list"`. Computed.

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks
will be discovered by default.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012013301231300-1112210210130030-3031330022023211-0003222121332302-1033112320332300-3031312003302121-2330323312311010-3320312123320122"></a>

### Direct properties for `server_block_filters`

<a id="canonical-3211310110210133-3200122133303302-3211002223212210-2312113013103321-0321312233013031-1212102131233100-2102130231110222-2130012202312111"></a>

#### `server_block_filters.name_regex` property

Type: `"string"`. Computed.

Regular expression to match the server name or domain that must be discovered.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3333211202200010-0221001202021110-2313211120120310-1013010331312330-3202111132223201-1222300233201303-2322021223012112-0133332320103023"></a>

<a id="canonical-1331221303100111-3232321132031301-2212103333300333-0012301013020001-0312123203220111-0213031123103101-1230221021332331-1012212221320232"></a>

#### `server_block_filters.port_ranges` property

Type: `"string"`. Computed.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Additional upstream details:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". Maximum number of ports allowed is
1024.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
