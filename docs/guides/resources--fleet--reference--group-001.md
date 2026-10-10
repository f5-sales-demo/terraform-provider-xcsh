---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- Property reference

<a id="canonical-2322212301121222-0323221021330032-1202231302230103-3103012110221210-3210013330002223-2220002213321001-2000301221332122-1001031303011031"></a>

### Direct properties for `xcsh_fleet`

- [allow_all_usb](resources--fleet--reference--group-001.md#canonical-0233020232022221-0012000111130132-3331103020110010-1132313232302301-3302110232202113-3033221010233312-1302312102121301-2003110120333123): complete subsection reference.

<a id="canonical-3331213132101021-2000221300121330-1322000022102033-1312322013303300-0022110022102332-3310303102220110-0130123331020003-2301203022033233"></a>

<a id="canonical-1321311020111303-0310230110210023-2130020203010130-3220210111233303-2330003202202112-1333101220100323-3122102033333320-2022121302032100"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233): complete subsection reference.

- [bond_device_list](resources--fleet--reference--group-001.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331): complete subsection reference.

- [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-2220011112301110-2101322302212032-2331222020113323-3333113013301010-2200210103313131-3003210110031021-1011111332222330-3232220020111203): complete subsection reference.

- [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-1022322123131331-3133331231203323-3213011213101211-0310133103101302-3003011111020032-0112210323033012-3222000332033020-0103003311002033): complete subsection reference.

- [default_config](resources--fleet--reference--group-002.md#canonical-3310002022222312-0033100302223102-0112021232310302-1233201223101302-0221021103321211-0101113223023222-3000011131123213-1121030200103313): complete subsection reference.

- [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-0321320013020113-2010101112101310-3001121223201313-1212300320113031-2233220102201110-2003223123110010-3213112211322133-3220222221031301): complete subsection reference.

- [default_storage_class](resources--fleet--reference--group-002.md#canonical-2311012123131011-1011230030130023-3102322012232302-1222231101331323-1232103111123230-0012022301103233-1003210001001231-0212032003000100): complete subsection reference.

- [deny_all_usb](resources--fleet--reference--group-002.md#canonical-2012100320202111-1003330223212301-3032022132110231-1312231300111330-1332100132012122-1310200213231333-3313103222231312-0121120321130320): complete subsection reference.

<a id="canonical-2132102100113132-0322202213230031-0001331003233100-1102233023320332-1111102222111023-3231222033101222-3330201210232032-2200122213320110"></a>

<a id="canonical-3021310022121231-1132211031020123-2102200111112103-2333202201012132-3033031103212211-3110200300302023-1113330111130031-0301031333133121"></a>

#### `description` property

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

- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303): complete subsection reference.

<a id="canonical-3021101001322222-1220020003032322-1323131010300233-2310031202130012-2202030100033133-2120333003102323-2003103312223023-1132211110303130"></a>

<a id="canonical-1113103223331111-2323022103212203-1301331001301103-1330012220000332-3311013101201110-2232223020300131-1020200330002033-2230313220132301"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [disable_gpu](resources--fleet--reference--group-002.md#canonical-0123032211233331-0131332231333212-3200202011200112-1001000300113120-0122200030312231-1201011230021221-0213033021011323-3032133220031003): complete subsection reference.

- [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-0030313313212311-2001301220310203-0020001123321003-2121032212103232-1131303332211222-3020011200213223-1213020113230210-0100321121202200): complete subsection reference.

- [disable_vm](resources--fleet--reference--group-002.md#canonical-0321001020131001-1211220132331302-1322103320333300-0031202311220102-1230303330102113-1311000031300313-1122113011301220-2122300130120201): complete subsection reference.

<a id="canonical-0000010333010311-0122213121310300-1101221020103130-0222233203223000-3010031321122202-0200121010113331-3320330100210013-2103002211230211"></a>

<a id="canonical-3231012321101321-0322000011221211-3331022323331332-3230333301011301-0201022220222122-1033312002331212-2020012202102300-1320101211230213"></a>

#### `enable_default_fleet_config_download` property

Type: `"bool"`. Optional, Computed.

Enable default fleet config, It must be set for storage config and GPU config.

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

- [enable_gpu](resources--fleet--reference--group-002.md#canonical-0002001130122322-1321022131203221-1122200222031023-1020213032321102-0222000230321311-3133232323322212-0110120030211330-3132231013210310): complete subsection reference.

- [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-1300130001320220-1113103202210133-1323332200222222-1011232231123110-0001121122301303-2210010201302013-1133102312333222-1032303221030230): complete subsection reference.

- [enable_vgpu](resources--fleet--reference--group-002.md#canonical-0023220101031001-1331310302022102-0213011031321233-3112103310211311-1212211111023103-3221232112323313-3332022200013311-1333313210022010): complete subsection reference.

- [enable_vm](resources--fleet--reference--group-002.md#canonical-0322133113031201-3211102313221333-0210203311332212-0110112033311332-0120013321023303-1301023021020300-0230103122131203-2030002233213303): complete subsection reference.

<a id="canonical-2131203023220223-1220213233212310-0200120023011212-2032013213320133-1123102133310030-3233010232323001-2002221331011102-1302122223221001"></a>

<a id="canonical-0120211101113031-2020211202011313-1323222111000310-0210233023333023-3321013330200212-3111313032111222-2212111110130213-1322033323011133"></a>

#### `fleet_label` property

Type: `"string"`. Required.

Fleet\_label value is used to create known\_label 'F5 XC/fleet=&lt;fleet\_label&gt;' The
known\_label is created in the 'shared' namespace for the tenant. A virtual\_site object with name
&lt;fleet\_label&gt; is also created in 'shared' namespace for tenant. The virtual\_site object will
select all sites..

Additional upstream details:

Fleet\_label value is used to create known\_label "F5 XC/fleet=&lt;fleet\_label&gt;" The
known\_label is created in the "shared" namespace for the tenant. A virtual\_site object with name
&lt;fleet\_label&gt; is also created in "shared" namespace for tenant. The virtual\_site object will
select all sites configured with the known\_label above fleet\_label with "sfo" will create a
known\_label "F5 XC/fleet=sfo" in tenant for the fleet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  }
}
```

<a id="canonical-1101032023031023-1133210000013233-3230302323032002-3200012013312223-3032120332132232-1320100221233320-3333132131212233-3013202000220223"></a>

<a id="canonical-1111233110210102-3212123201132211-3131101013103300-0311213132132301-1310011202132320-2021203122021320-3021331010100000-1113103100121221"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [inside_virtual_network](resources--fleet--reference--group-002.md#canonical-3333200312232013-0113002003132011-3330301312033033-2120022120220222-1123100201032211-1221013003232322-1100332121322133-3130123311000222): complete subsection reference.

- [interface_list](resources--fleet--reference--group-002.md#canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302): complete subsection reference.

- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300): complete subsection reference.

<a id="canonical-1110230313122310-3120023020031301-0012300321123230-1111130312121220-1132331332020313-3133022311301020-0003302002002221-0013102012123120"></a>

<a id="canonical-1013333133302221-2021012321021211-1233023012312230-2330122113133313-1130033313232100-0001202022302322-3312210220003222-3003333021302213"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

- [log_receiver](resources--fleet--reference--group-002.md#canonical-0331330130330132-3110332130130011-0101221033032011-1321323001303111-1310212222202120-1000000110001011-0200233010121032-2221330200010033): complete subsection reference.

- [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-3212213123002023-0202231202121031-0011310312010030-0202200201110132-0022133203021121-1110002111002320-0011010213122003-0312100022301211): complete subsection reference.

<a id="canonical-0121133220020032-0313210013012102-0201133230331030-1203312103132312-0121022001112230-2211220302122300-0201332310211010-2221003010203221"></a>

<a id="canonical-1123221333102321-3013122021303021-3002322132013231-3330012130021233-1213100011103313-2100330221311300-0310212203021121-3130130011120023"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Fleet. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1120331003331031-2121131021113130-0312013002033031-1212133201023113-3223310113332121-1121203203131010-1233223330331233-2123300303203122"></a>

<a id="canonical-0231313230013023-2202103133010132-0223201022222220-2121222003322132-3330100332113303-3313332032331013-2132330013211202-3313121200200101"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Fleet is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [network_connectors](resources--fleet--reference--group-002.md#canonical-3323111111030122-1303021120311223-1110003221103022-2020332013330231-3132231113011220-1121100020211332-0213211322111100-3303110210031121): complete subsection reference.

- [network_firewall](resources--fleet--reference--group-002.md#canonical-2101130020120112-0133200133323012-0232110022111200-1113311010303213-1313213332310323-1200301201112022-1100302111320120-3222312321231213): complete subsection reference.

- [no_bond_devices](resources--fleet--reference--group-002.md#canonical-3203320110130321-1121100220230211-2020333200023102-0312320112101333-3102032302000203-1112313033231111-3102003021133322-0232123301133201): complete subsection reference.

- [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-3002230220210021-3122231112230020-3301210303100111-2000220001202332-3133200133012031-0222312213013001-0021203302213202-2210201320011200): complete subsection reference.

- [no_storage_device](resources--fleet--reference--group-002.md#canonical-1323333231331210-1331121020203033-1001232023202221-0323221132032011-2122012311233313-0333321001230223-1230201132012120-3320022013311213): complete subsection reference.

- [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-3021101310113300-2330120320311311-0103112023302012-0111200030121332-2122032220110113-0122101022233212-3300032121002220-0102132112203230): complete subsection reference.

- [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-2231011202030311-3303111022101200-2001103313313330-0212102002001302-2030113131220333-1311323330233023-2120212100232203-0031032331001130): complete subsection reference.

<a id="canonical-3232010320330012-3033321031230233-0101301030010133-1323022102213310-1012211222232133-3221122313112121-2030323103200000-0213213221102222"></a>

<a id="canonical-1302232330221313-1123132231320202-2220021100100010-1103213131113011-2023210321032303-0201112003032223-0112301130321223-2010133200211232"></a>

#### `operating_system_version` property

Type: `"string"`. Optional, Computed.

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [outside_virtual_network](resources--fleet--reference--group-002.md#canonical-3223331122001220-3011311231202031-2020330101213300-3133211103110113-2113210013003020-2001322230121133-3212301331022221-2100121222312301): complete subsection reference.

- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333): complete subsection reference.

- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021): complete subsection reference.

- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023): complete subsection reference.

- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101): complete subsection reference.

- [storage_interface_list](resources--fleet--reference--group-003.md#canonical-2200223121321313-3202113301130230-3113322333033203-2122021133112132-3302320001211330-3311111102111033-2201322232210330-0102122100010132): complete subsection reference.

- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330): complete subsection reference.

- [timeouts](resources--fleet--reference--group-004.md#canonical-3013312232300101-1023222303301232-0210002302033021-2230300221300103-1230132310201312-1210112221101000-0010122001311202-1320000331121210): complete subsection reference.

- [usb_policy](resources--fleet--reference--group-004.md#canonical-1313133033021303-0322312200331123-0101000033230330-1300123032023012-1022323323332131-3002113220222232-0302220132222213-1022313213313302): complete subsection reference.

<a id="canonical-1112311233323303-0300320300110001-2131011211002222-2113300313322103-1132133203321220-3302313320200210-0333231020001200-1202231201133302"></a>

<a id="canonical-2111331233013221-0113232213311311-2220223003120122-0121131303002000-1112223302033301-3100212211023202-0313111302231200-2013101222021303"></a>

#### `volterra_software_version` property

Type: `"string"`. Optional, Computed.

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0021112332233103-0210300222111211-0302112213023100-2012122020320211-2021021320221231-1212302331101310-1022332122112221-1101031002130200"></a>

### All schema paths for `xcsh_fleet`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_usb` | [allow_all_usb](resources--fleet--reference--group-001.md#canonical-0011122200232321-0331210222021203-3211310010331331-0313201100121033-0210102200112232-2220223223100313-1001011001120021-3220221212103101) |
| `annotations` | [annotations](resources--fleet--reference--group-001.md#canonical-3331213132101021-2000221300121330-1322000022102033-1312322013303300-0022110022102332-3310303102220110-0130123331020003-2301203022033233) |
| `blocked_services` | [blocked_services](resources--fleet--reference--group-001.md#canonical-2300213201220031-0101132121033233-3121011212322211-3303123131131133-1203331103123303-2331223201001111-2121021312131201-1321123122123132) |
| `blocked_services.dns` | [blocked_services.dns](resources--fleet--reference--group-001.md#canonical-3001003313321332-3330212121120221-2231001312302113-0130210130101030-1223311221031011-1202322303232120-3103302032023322-0021233022113320) |
| `blocked_services.network_type` | [blocked_services.network_type](resources--fleet--reference--group-001.md#canonical-2213010100023002-1333330311200302-3200201323202200-0210010100210332-1001103132033320-0331312232110230-3113122031033311-0230031130001232) |
| `blocked_services.ssh` | [blocked_services.ssh](resources--fleet--reference--group-001.md#canonical-0302223331101002-0213123231100103-3000030300221332-3300313302033202-0212203220321130-1330103300322313-0321101230120100-0002110202021030) |
| `blocked_services.web_user_interface` | [blocked_services.web_user_interface](resources--fleet--reference--group-001.md#canonical-1123031312320312-1220130012132210-0133032010322123-2101323021222330-2023233200223023-2311122321002303-0030311213221031-2003303232021033) |
| `bond_device_list` | [bond_device_list](resources--fleet--reference--group-001.md#canonical-3112331033303132-1010323223233021-0001113321020203-0020000321010310-0312210031322111-2212121113112120-0012220002121030-2332000033322102) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](resources--fleet--reference--group-001.md#canonical-3233230000001102-0311301000203233-2111011133110120-3233103113333020-3021302233123110-2002020302302210-3331023230132110-2230111323213220) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](resources--fleet--reference--group-002.md#canonical-2132113223210212-3220131323133320-0112021122002113-1331132211002101-2212032312213130-1230303323303103-3132122000332110-3231200233220012) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](resources--fleet--reference--group-001.md#canonical-1102101300033333-1222302001120331-3001102112133100-1113221021010223-1011033102031323-2001201332011123-3021100212311201-3021223202312132) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](resources--fleet--reference--group-002.md#canonical-3313323230011200-3112222003013333-1332013011020223-2101133212133213-0133030103102010-3231322320010231-0110131020223102-3201102113332030) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](resources--fleet--reference--group-002.md#canonical-0121011021031200-1203113213112302-2311030033311221-0001200330000302-2302020222001220-1313333320120131-2133220231103311-3311213310231231) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](resources--fleet--reference--group-001.md#canonical-3220201301033322-3120000332322333-0213203203003312-3232200020010033-2303012020310131-2302311233300131-0000013010231312-3013130100312323) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](resources--fleet--reference--group-001.md#canonical-2112131103330123-0303001101002111-1102112103200311-1113220212110022-2303210120033112-0031200132110330-1222020030102232-3102030331313323) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](resources--fleet--reference--group-001.md#canonical-2133233103031220-1330321001113020-1103312233020121-0100122200001233-1103032120012000-3130012302010322-3313003331032211-2100312203130113) |
| `dc_cluster_group` | [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-3222132011302110-3223233110303022-0010111003310331-1232121213200230-0212010203312230-2220311223301131-1013331123331132-0011233232032100) |
| `dc_cluster_group.name` | [dc_cluster_group.name](resources--fleet--reference--group-002.md#canonical-0101131011112223-1230130100133332-1220213300302321-2300002103000222-3121030302332303-2313202311120022-3112022200100121-2300011020133120) |
| `dc_cluster_group.namespace` | [dc_cluster_group.namespace](resources--fleet--reference--group-002.md#canonical-2133010021202033-1133110231202332-3311230000223312-0103230123323001-2003333320312222-0331320310323013-0303101222201001-1323102123303313) |
| `dc_cluster_group.tenant` | [dc_cluster_group.tenant](resources--fleet--reference--group-002.md#canonical-2023322113123002-3100320322200311-3001321002020130-3332211000200200-3223000030112330-3303213001022330-1221003223012102-0221221123233202) |
| `dc_cluster_group_inside` | [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-3113032020311002-1032323211212323-0330330230100021-2112321321021100-2313331222110031-1213131220011002-2013013013100021-3210322011302131) |
| `dc_cluster_group_inside.name` | [dc_cluster_group_inside.name](resources--fleet--reference--group-002.md#canonical-2011323303112120-1312121233303233-1311311230113121-0221023112223103-2320311333333331-0203021123121111-3322023001102331-1102313203031032) |
| `dc_cluster_group_inside.namespace` | [dc_cluster_group_inside.namespace](resources--fleet--reference--group-002.md#canonical-2100220030012020-0321102223000001-0102112221223102-1222121232320330-1020103111112312-1221001121120222-3010133001003222-0130001013213001) |
| `dc_cluster_group_inside.tenant` | [dc_cluster_group_inside.tenant](resources--fleet--reference--group-002.md#canonical-2012030333200310-0333322331231021-3020333102301233-1322113333332233-1121311121112010-2121323022000230-3023201203213011-2113201032312211) |
| `default_config` | [default_config](resources--fleet--reference--group-002.md#canonical-0222331031323120-0233230301313130-2222300312211303-2131330122312100-3030030103000312-3321300032012020-1221213113311200-3302320011013321) |
| `default_sriov_interface` | [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-3120231221133033-3133113300203332-3300212000321212-2230302302231033-3110323323131002-3333002303210003-1111030102211010-3220300203133231) |
| `default_storage_class` | [default_storage_class](resources--fleet--reference--group-002.md#canonical-3232310021320301-3232001110331123-2122322312002112-3020302312022233-3221320130031123-1100300210011002-2302313031231330-2332332332120130) |
| `deny_all_usb` | [deny_all_usb](resources--fleet--reference--group-002.md#canonical-3031222031130102-0220323323303102-0001201301130320-0310333220032103-2020133202021100-3213120313202223-2013331220210232-1002130131230330) |
| `description` | [description](resources--fleet--reference--group-001.md#canonical-2132102100113132-0322202213230031-0001331003233100-1102233023320332-1111102222111023-3231222033101222-3330201210232032-2200122213320110) |
| `device_list` | [device_list](resources--fleet--reference--group-002.md#canonical-3200322202323330-0323133003013331-1122130122320213-3221210232101013-1200323220311002-0232323111303111-2110131211011210-0020213222312313) |
| `device_list.devices` | [device_list.devices](resources--fleet--reference--group-002.md#canonical-0231031102013112-0023111223302101-2021310021322003-1200312221122032-1300223112001013-2331211030032020-3300200103100133-2033202112010213) |
| `device_list.devices.name` | [device_list.devices.name](resources--fleet--reference--group-002.md#canonical-1100231202123210-2011110011112001-2113330223130011-2212213222220313-3201110322220330-3000120231022133-1111321100131301-0003123202132031) |
| `device_list.devices.network_device` | [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-1233122212201230-3130021332033113-3121011121310113-3221100021110321-3012311120233122-0100013233333031-0122123300302103-3222001210333012) |
| `device_list.devices.network_device.interface` | [device_list.devices.network_device.interface](resources--fleet--reference--group-002.md#canonical-3100120003312303-0003110011302032-3320031212231121-2220210003310112-2103210012021210-3102323323031331-0133312302300322-1002303103203112) |
| `device_list.devices.network_device.interface.kind` | [device_list.devices.network_device.interface.kind](resources--fleet--reference--group-002.md#canonical-1012010012103110-3030010120023321-1000231213302133-3332203032230213-0132012311012310-3320232210202323-3103000102220023-1013202021123323) |
| `device_list.devices.network_device.interface.name` | [device_list.devices.network_device.interface.name](resources--fleet--reference--group-002.md#canonical-3233333311000113-0020001313212012-1200231301321133-1222110310101332-1022220111201131-1012201330230012-1311132232333331-2033201310000320) |
| `device_list.devices.network_device.interface.namespace` | [device_list.devices.network_device.interface.namespace](resources--fleet--reference--group-002.md#canonical-2213021333121212-1021213211210102-0122021231233013-3310220120110300-3122232133132230-2013200320003131-1021210103032312-0202312102310021) |
| `device_list.devices.network_device.interface.tenant` | [device_list.devices.network_device.interface.tenant](resources--fleet--reference--group-002.md#canonical-0230321303203311-1001303000233233-3313032111020212-3230212302231332-0223323100130231-3000333322323200-3001030112332013-0301123220030223) |
| `device_list.devices.network_device.interface.uid` | [device_list.devices.network_device.interface.uid](resources--fleet--reference--group-002.md#canonical-0233221033313310-3010220003120213-2212031120312101-0220202203310302-0331302111111313-0012001102223032-3333320201013310-1123322110301202) |
| `device_list.devices.network_device.use` | [device_list.devices.network_device.use](resources--fleet--reference--group-002.md#canonical-2231211212012001-1230322122322023-2101323312122103-0023012020120213-1103321033103103-2130100023103020-3130312223002122-3120003011331310) |
| `device_list.devices.owner` | [device_list.devices.owner](resources--fleet--reference--group-002.md#canonical-2211213033230302-1031233322012201-3022301211020123-0221130310223300-2010203301032032-3020303323311103-2233211130100012-1112111100021233) |
| `disable` | [disable](resources--fleet--reference--group-001.md#canonical-3021101001322222-1220020003032322-1323131010300233-2310031202130012-2202030100033133-2120333003102323-2003103312223023-1132211110303130) |
| `disable_gpu` | [disable_gpu](resources--fleet--reference--group-002.md#canonical-0213033021223211-2011031122232000-0031331130102311-0223100313001222-0032002333002023-3222030023220001-1220103333102311-3323302230201330) |
| `disable_log_anonymization` | [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-2121001302333101-3211331033211213-3233213331231013-2301210300232123-0033110031101222-0312002133222130-0120100323223301-0331300320023001) |
| `disable_vm` | [disable_vm](resources--fleet--reference--group-002.md#canonical-3230331213222102-3010132211202323-0302300222121300-1322023031220311-3323322213002302-1122103131213313-1321302120133121-2211100120320133) |
| `enable_default_fleet_config_download` | [enable_default_fleet_config_download](resources--fleet--reference--group-001.md#canonical-0000010333010311-0122213121310300-1101221020103130-0222233203223000-3010031321122202-0200121010113331-3320330100210013-2103002211230211) |
| `enable_gpu` | [enable_gpu](resources--fleet--reference--group-002.md#canonical-2200303220000331-1001222020122121-1313202222301130-3330333133022123-3302320023303020-0201020112330320-2301302111030103-3220132132312202) |
| `enable_log_anonymization` | [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-3200320200030321-3033013222300332-2010230132310221-3032313333233322-0332313113200121-1201222331033303-3120202212101301-3020103131101003) |
| `enable_vgpu` | [enable_vgpu](resources--fleet--reference--group-002.md#canonical-2022102100100133-3233222033211321-3200233210103333-1131101321130311-2120202203330101-1302300011010321-3003031331100131-3122303110120030) |
| `enable_vgpu.feature_type` | [enable_vgpu.feature_type](resources--fleet--reference--group-002.md#canonical-2023210113311333-0320211122000010-1023321012002001-3210010223100003-1321020212111320-0122033120111313-0332033011031222-0102020333322310) |
| `enable_vgpu.server_address` | [enable_vgpu.server_address](resources--fleet--reference--group-002.md#canonical-1023131222202213-0331220323130333-0131111200110203-2033123101001111-3121102110021223-0101032120223231-0121001010101202-3330220300231212) |
| `enable_vgpu.server_port` | [enable_vgpu.server_port](resources--fleet--reference--group-002.md#canonical-0110302323212200-1232132013132232-0011133101132231-1011111220312201-2020020031121032-0320011033100232-0210111310221232-0201022112313332) |
| `enable_vm` | [enable_vm](resources--fleet--reference--group-002.md#canonical-0300310111033333-3323012300233121-0223320333212221-2033211212300122-0320122222321001-2130221000220223-1131213121203130-2103020131022232) |
| `fleet_label` | [fleet_label](resources--fleet--reference--group-001.md#canonical-2131203023220223-1220213233212310-0200120023011212-2032013213320133-1123102133310030-3233010232323001-2002221331011102-1302122223221001) |
| `id` | [ID](resources--fleet--reference--group-001.md#canonical-1101032023031023-1133210000013233-3230302323032002-3200012013312223-3032120332132232-1320100221233320-3333132131212233-3013202000220223) |
| `inside_virtual_network` | [inside_virtual_network](resources--fleet--reference--group-002.md#canonical-2122322131012313-3130301003123021-1212130323302111-0203212231233110-0003213231120130-0211202223132033-2230322211113002-3013330103212130) |
| `inside_virtual_network.kind` | [inside_virtual_network.kind](resources--fleet--reference--group-002.md#canonical-2230030320120311-1301002313331101-1033031031133302-3130011021332011-1203003022220020-1011231331303022-0010212112311232-2230211222220020) |
| `inside_virtual_network.name` | [inside_virtual_network.name](resources--fleet--reference--group-002.md#canonical-1011023102110022-2333212123222312-1202220112210103-3232201131211031-0133131312230321-2111002112331103-1120030101022102-3121000312223003) |
| `inside_virtual_network.namespace` | [inside_virtual_network.namespace](resources--fleet--reference--group-002.md#canonical-2111022200220232-1310232023101322-2300133301202011-1320312010121333-2030313313300120-0303122210120001-0022301332210111-3123113233212033) |
| `inside_virtual_network.tenant` | [inside_virtual_network.tenant](resources--fleet--reference--group-002.md#canonical-2022230213020013-0333210331102312-2310211033222213-3330313003233032-3102003110321220-3011030031032203-1210003232121031-2130033312131303) |
| `inside_virtual_network.uid` | [inside_virtual_network.uid](resources--fleet--reference--group-002.md#canonical-3221222313112131-3131320132212203-3111322330230132-0023230111333013-3312113210322012-3021321332013130-1210013322202333-1213001110000322) |
| `interface_list` | [interface_list](resources--fleet--reference--group-002.md#canonical-1022102312212002-1321333300332120-1002202302210311-1003202220302323-2023331213223200-0220022300211302-2130200100001300-3021321322102123) |
| `interface_list.interfaces` | [interface_list.interfaces](resources--fleet--reference--group-002.md#canonical-2231333023223222-2300022113210131-0103201221101230-1302323202012100-0130121122330333-1001010010030123-0302102232132101-1103210332202212) |
| `interface_list.interfaces.name` | [interface_list.interfaces.name](resources--fleet--reference--group-002.md#canonical-0012002302301001-1302031001230033-0101120022101302-3220031301122001-3000330122031021-2112003213300001-0121301102032112-1110011131000000) |
| `interface_list.interfaces.namespace` | [interface_list.interfaces.namespace](resources--fleet--reference--group-002.md#canonical-1000221230233223-2330200323332330-2233300233233020-0223011131213102-1113102002130023-0000211223313220-2312231210132112-0330130220230133) |
| `interface_list.interfaces.tenant` | [interface_list.interfaces.tenant](resources--fleet--reference--group-002.md#canonical-0131003000100302-2302222210203211-1121332333131202-2003223101221013-1100033212023303-1333133232102301-0003312333312100-1131023003200221) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2021130120130131-0131201122032111-0100011213212223-3331321110111101-3311000033113323-1332022231022230-0220132320100133-0221011220232023) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-1230301023033301-2231012323222023-1022233031331311-2121301231200022-2032231011102332-1103112103023220-3100110302212212-3223221220113201) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2200003321111012-2102312032003130-0331322023323100-1033223013033010-0132110322312003-2121320100203112-2220001120330032-0321201312231310) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-3303103200232103-1120111120302102-2101210203313300-1121032310221130-1213002132012230-2233011102031033-0001200020220030-2212021010031011) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--fleet--reference--group-002.md#canonical-0312303200220300-2110120203113210-0133330021211203-3022113320111311-2211100210022330-1322023011232212-0220103030102123-2031210133020101) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--fleet--reference--group-002.md#canonical-1330022231112010-3133223023113013-1310120132011001-0331200003013230-2032223111103331-3000102103002220-2302031111131021-1330001033131320) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--fleet--reference--group-002.md#canonical-3323030303221233-2100022310303231-0221302131302232-3011012213232112-0001210020220121-0302223122202021-3332000223033102-2023203132313122) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-0001132232313221-1301212330111231-0201303232201201-0020323302213312-0210301323111031-3303033113100022-2131211311330102-2310322302233300) |
| `labels` | [labels](resources--fleet--reference--group-001.md#canonical-1110230313122310-3120023020031301-0012300321123230-1111130312121220-1132331332020313-3133022311301020-0003302002002221-0013102012123120) |
| `log_receiver` | [log_receiver](resources--fleet--reference--group-002.md#canonical-3031313201013102-3012032012303320-0312101313130103-3322003301113303-1220323132023022-0301030301323133-2222033021303010-2022122130110203) |
| `log_receiver.name` | [log_receiver.name](resources--fleet--reference--group-002.md#canonical-1130020113021020-3320232131202113-0302012111310003-0010301032313331-2123331221203013-2310211022130210-2303212231301010-3303012113202202) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--fleet--reference--group-002.md#canonical-3203000220210021-2012022203121223-3231300120111030-2311331332031003-0123213133002111-1122201032101230-3010032103021231-3122320031310201) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--fleet--reference--group-002.md#canonical-1112232211002303-3213210020311113-3230102202232032-3131313112232133-1330010131302300-3313312032212033-0302132231000212-0332111013211222) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-2131023030210200-2311112100032332-0203003232113123-3101323130301101-1231112112221012-1033211311300001-2131011123301122-0232232301033323) |
| `name` | [name](resources--fleet--reference--group-001.md#canonical-0121133220020032-0313210013012102-0201133230331030-1203312103132312-0121022001112230-2211220302122300-0201332310211010-2221003010203221) |
| `namespace` | [namespace](resources--fleet--reference--group-001.md#canonical-1120331003331031-2121131021113130-0312013002033031-1212133201023113-3223310113332121-1121203203131010-1233223330331233-2123300303203122) |
| `network_connectors` | [network_connectors](resources--fleet--reference--group-002.md#canonical-2233111321103202-3320001311001223-1302200122123131-1120102313203322-1013031230100002-1200223111020300-0110311031222211-0030300120222110) |
| `network_connectors.kind` | [network_connectors.kind](resources--fleet--reference--group-002.md#canonical-2220233203213002-1010103303011202-0020302231231130-0332332120200001-3030001203330110-2201333303220222-2331012120200023-1333221330230123) |
| `network_connectors.name` | [network_connectors.name](resources--fleet--reference--group-002.md#canonical-1222331120102022-2213033310302010-1313011233223210-1123013312203001-1003010313200211-1110200101102231-3130101130101030-3300202123100332) |
| `network_connectors.namespace` | [network_connectors.namespace](resources--fleet--reference--group-002.md#canonical-0033221133101031-1201020303111012-2030122312223333-2000210102121201-1111111123130132-2301313130220002-2321331110312333-3112011021021220) |
| `network_connectors.tenant` | [network_connectors.tenant](resources--fleet--reference--group-002.md#canonical-3232100332020232-0223211231000021-0132322330000322-0101123321202332-1302031103303010-0101032032132210-2323231122121121-3021000030311100) |
| `network_connectors.uid` | [network_connectors.uid](resources--fleet--reference--group-002.md#canonical-3102022112233220-1000322113230013-0332312313232121-0103231013001200-1211023212000313-2323232300103020-2111202333330322-0220212330102320) |
| `network_firewall` | [network_firewall](resources--fleet--reference--group-002.md#canonical-2320210312033302-0003021301023012-0231131302212333-2022302003133123-1222032202030023-3301111031131101-1210122331032011-2320112313121321) |
| `network_firewall.kind` | [network_firewall.kind](resources--fleet--reference--group-002.md#canonical-0223030201022003-3120030120002220-2201022022220103-0223220033120311-0121213212010021-3001100203102230-1030200003033131-2321320300330320) |
| `network_firewall.name` | [network_firewall.name](resources--fleet--reference--group-002.md#canonical-1021011111130311-2302113223123123-3120133233300333-1311320221333012-1312233201330110-3213202312310301-3323320320113112-3230332100202130) |
| `network_firewall.namespace` | [network_firewall.namespace](resources--fleet--reference--group-002.md#canonical-3012301122232132-3310212310201133-0023123300211213-0200022233111102-1201113132100133-0122113022013111-3023032100001110-1133030001311331) |
| `network_firewall.tenant` | [network_firewall.tenant](resources--fleet--reference--group-002.md#canonical-2312010111200312-1013332103323313-1030102302222122-3323230000221131-1203030230001121-0100033232311223-0100320331210120-3111131221222000) |
| `network_firewall.uid` | [network_firewall.uid](resources--fleet--reference--group-002.md#canonical-0000100210323231-3311001011020032-3322011200100021-3320011102311200-3213011023220103-0232113231202020-2230302313332210-1111320230113121) |
| `no_bond_devices` | [no_bond_devices](resources--fleet--reference--group-002.md#canonical-2233202132112121-1300211023330221-2212133210020112-2021322122322312-0321223001202132-3220130012222211-3210110102122303-1332222011023230) |
| `no_dc_cluster_group` | [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-0300111311003312-2320102111320110-2333221333222221-2321223000332000-0132220211220113-0221111301300211-3031311131013010-3001123212213133) |
| `no_storage_device` | [no_storage_device](resources--fleet--reference--group-002.md#canonical-0220230132023221-0012122010222212-2110320123231220-3000120023300020-3313010120130211-3322003111321031-1011003102230130-0132323201010321) |
| `no_storage_interfaces` | [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-0132211203012133-3201031322322320-1332302222000203-1001303002202201-0323112222133113-3313003302132131-0010231220301220-2322322220001322) |
| `no_storage_static_routes` | [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-3333003330002323-2020312322133222-3003221302331313-3221312221202102-1100230230310222-2202101300132202-0113222211333133-1021333131121331) |
| `operating_system_version` | [operating_system_version](resources--fleet--reference--group-001.md#canonical-3232010320330012-3033321031230233-0101301030010133-1323022102213310-1012211222232133-3221122313112121-2030323103200000-0213213221102222) |
| `outside_virtual_network` | [outside_virtual_network](resources--fleet--reference--group-002.md#canonical-3120120032101112-1320012311032120-3123110212013301-0222000011303300-0020203112032230-3201230233112233-0103200303010013-1233332133233032) |
| `outside_virtual_network.kind` | [outside_virtual_network.kind](resources--fleet--reference--group-002.md#canonical-0103321323202010-1112100230213231-2013222330100131-3100300130311313-2123110110201013-0013230323113200-1000111231213133-2131021032013111) |
| `outside_virtual_network.name` | [outside_virtual_network.name](resources--fleet--reference--group-002.md#canonical-3111030303003033-1030111031000020-0202023121210221-0222210232033030-1330123013202110-2023112113130310-0321010213100132-1200310213311033) |
| `outside_virtual_network.namespace` | [outside_virtual_network.namespace](resources--fleet--reference--group-002.md#canonical-1122113103211203-2003130010301203-1212211233311010-3002233032130233-0030102201132022-2023022231111300-3001233132033331-3131211233030110) |
| `outside_virtual_network.tenant` | [outside_virtual_network.tenant](resources--fleet--reference--group-002.md#canonical-0212012002130323-1011132302233203-2020131211303310-2022201103312101-0020133112000330-2300131020232101-0020022303021000-2212310200333010) |
| `outside_virtual_network.uid` | [outside_virtual_network.uid](resources--fleet--reference--group-002.md#canonical-1012003302101233-1223031302033032-3132122001101032-2312012212303001-1310122212210233-1303120200210121-3033311322021301-0330003121323333) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-3030113113100330-0012133311113321-0021320010121312-3022300023333220-0130331330032012-2302300003220202-3021321212323322-3003213102113110) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-3203210222120212-1213232111221030-1330301330231323-3320223321203311-1210023122323120-1110002201030313-1323331031332233-0222312122231221) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--fleet--reference--group-002.md#canonical-2321122131222032-1311331312232023-1220132003213010-1313331233223101-2301010022103032-3321120112233322-2030223230311232-2111120030010120) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--fleet--reference--group-002.md#canonical-0221131011213303-3333302211113030-0332103013132331-0313211313323002-1312111223120222-1330302333230123-1221022132130011-2131020022022031) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-0102133110100122-1213312131112131-3212313202300102-3010213002211013-1112333132023102-2303102120221133-1312032032313222-1303320333220101) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--fleet--reference--group-002.md#canonical-2323033031003202-0123222000220312-1310020102232010-1321133223313132-3002221002330302-2223310003123311-2211110100030312-3233313331120110) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--fleet--reference--group-002.md#canonical-3100110103313003-2220220000200110-3121111023101023-1311130111113200-1120212313003100-0323210120321110-2000120303323000-2303102232221120) |
| `sriov_interfaces` | [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0122221131232302-1301033322113301-0110223212011211-3233102320110033-0121020202101031-3330230010133021-2322320310031011-0300013133213110) |
| `sriov_interfaces.sriov_interface` | [sriov_interfaces.sriov_interface](resources--fleet--reference--group-002.md#canonical-0111313130111223-0101332233300300-2203102011300122-3303001100021221-1201313331221303-2320100023133330-2300300123213223-0013031000311111) |
| `sriov_interfaces.sriov_interface.interface_name` | [sriov_interfaces.sriov_interface.interface_name](resources--fleet--reference--group-002.md#canonical-2123331310213120-2210230032010112-2320220102201000-1013122020123333-2133301212021301-3302300131120120-2333320113133233-2011003233203113) |
| `sriov_interfaces.sriov_interface.number_of_vfio_vfs` | [sriov_interfaces.sriov_interface.number_of_vfio_vfs](resources--fleet--reference--group-002.md#canonical-2023133122231310-0113222003213021-0313132300113010-3222133301301012-1123330223030202-1202220130201302-2123131332333001-3213231000110110) |
| `sriov_interfaces.sriov_interface.number_of_vfs` | [sriov_interfaces.sriov_interface.number_of_vfs](resources--fleet--reference--group-002.md#canonical-1233020321322311-2312010020020301-2120132130223210-2312002232001100-2331221210303221-3130003303230112-3330131212031312-1023031200111211) |
| `storage_class_list` | [storage_class_list](resources--fleet--reference--group-002.md#canonical-2222300023023313-2122302313220111-2233300300012313-3303221100131111-0112033102321132-1321110322120221-0103222322331100-1100211021322120) |
| `storage_class_list.storage_classes` | [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-1330213321222123-1121113101212033-0130331233330120-2323211313312210-3212020022203233-2102131212001131-0330220113011110-2023323211312113) |
| `storage_class_list.storage_classes.advanced_storage_parameters` | [storage_class_list.storage_classes.advanced_storage_parameters](resources--fleet--reference--group-002.md#canonical-2221030123300233-0110230120011332-1303202011130000-2020112230013002-3221220130102100-0102221332011012-2113210020330311-1113012001022321) |
| `storage_class_list.storage_classes.allow_volume_expansion` | [storage_class_list.storage_classes.allow_volume_expansion](resources--fleet--reference--group-002.md#canonical-2202133111232311-1231331012012123-1121023221003231-2102332300032331-1132210010111111-3333320013013012-1330210010031012-0023121003112131) |
| `storage_class_list.storage_classes.custom_storage` | [storage_class_list.storage_classes.custom_storage](resources--fleet--reference--group-002.md#canonical-3033311331231221-1311001010132332-0211110001021002-2302112231023110-2122322120033213-0223103231220001-0223003330203112-3310232133231010) |
| `storage_class_list.storage_classes.custom_storage.yaml` | [storage_class_list.storage_classes.custom_storage.yaml](resources--fleet--reference--group-002.md#canonical-0012223213023021-2022231023300113-2210222233333302-2300202213333101-1102133000202220-3303213120332332-3021112012210223-1220301023303211) |
| `storage_class_list.storage_classes.default_storage_class` | [storage_class_list.storage_classes.default_storage_class](resources--fleet--reference--group-002.md#canonical-0310131120333103-3131013023012023-3332321000031320-0230220103031110-1122331331303233-2320211131001321-3302202012303030-3111003212101112) |
| `storage_class_list.storage_classes.description_spec` | [storage_class_list.storage_classes.description_spec](resources--fleet--reference--group-002.md#canonical-1120201231001110-3312300231013311-1103123310101213-3023112130023312-1033000201010220-2131110100311320-3233213220021030-1320012310312203) |
| `storage_class_list.storage_classes.hpe_storage` | [storage_class_list.storage_classes.hpe_storage](resources--fleet--reference--group-002.md#canonical-0030320200321221-1122032232023202-1010033100131301-0303232232220311-0120120111310233-1311333110233221-1032103003203013-3313223100132023) |
| `storage_class_list.storage_classes.hpe_storage.allow_mutations` | [storage_class_list.storage_classes.hpe_storage.allow_mutations](resources--fleet--reference--group-002.md#canonical-3021222303033311-0032023011212013-0131011032311111-1012211033111220-0121110212320220-3331320321311220-2111013133113221-2233032231230101) |
| `storage_class_list.storage_classes.hpe_storage.allow_overrides` | [storage_class_list.storage_classes.hpe_storage.allow_overrides](resources--fleet--reference--group-002.md#canonical-1103003000332122-1130002230033202-0200130110002010-2133123131130111-3102010010002133-1233332023123312-3010132233003312-0323210003333100) |
| `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` | [storage_class_list.storage_classes.hpe_storage.dedupe_enabled](resources--fleet--reference--group-002.md#canonical-0330213203012130-2131032223020120-0232111132100010-1013333302020001-0200300013211313-0200300133302132-1013002211223202-1032222331131232) |
| `storage_class_list.storage_classes.hpe_storage.description_spec` | [storage_class_list.storage_classes.hpe_storage.description_spec](resources--fleet--reference--group-002.md#canonical-0212323122033303-1030320320313302-0131212130300330-3223123002300120-3123001121303220-0123010213302023-3013222132213321-3020203200131001) |
| `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` | [storage_class_list.storage_classes.hpe_storage.destroy_on_delete](resources--fleet--reference--group-002.md#canonical-2032311201210020-2101301322313320-0330103020121311-1311021021023000-1011123330312332-0302022323023311-2033032323003223-1020310020313203) |
| `storage_class_list.storage_classes.hpe_storage.encrypted` | [storage_class_list.storage_classes.hpe_storage.encrypted](resources--fleet--reference--group-002.md#canonical-3213233032032202-3222120021103130-1020000100120213-0203032012220101-0023033333301311-0110102123120122-1133210300102221-3132231231121312) |
| `storage_class_list.storage_classes.hpe_storage.folder` | [storage_class_list.storage_classes.hpe_storage.folder](resources--fleet--reference--group-002.md#canonical-0100223302023313-1230212121013111-3020013301330222-3000211003311213-3233111132103221-2310101332022212-1000112022203000-1012321220111313) |
| `storage_class_list.storage_classes.hpe_storage.limit_iops` | [storage_class_list.storage_classes.hpe_storage.limit_iops](resources--fleet--reference--group-002.md#canonical-0331312333303230-3213032021203001-3303331222003121-1211221103231122-0003210111232303-1203323130103202-2331131220233332-2020122331222211) |
| `storage_class_list.storage_classes.hpe_storage.limit_mbps` | [storage_class_list.storage_classes.hpe_storage.limit_mbps](resources--fleet--reference--group-002.md#canonical-0000202201101113-2000113021101311-1312103120020030-1221112332321330-1020013122032001-2223202120311011-0101113002111033-3331022012310011) |
| `storage_class_list.storage_classes.hpe_storage.performance_policy` | [storage_class_list.storage_classes.hpe_storage.performance_policy](resources--fleet--reference--group-002.md#canonical-2312032311313213-2103301311122211-3003002223320220-2113321111222123-0110023132211331-2123022212003132-0232101001233332-2113202313030022) |
| `storage_class_list.storage_classes.hpe_storage.pool` | [storage_class_list.storage_classes.hpe_storage.pool](resources--fleet--reference--group-002.md#canonical-1123120313000032-3010213213201310-0201310123223002-2130203001013203-2113200321303231-1001020332010112-0232220012000023-3323000002030032) |
| `storage_class_list.storage_classes.hpe_storage.protection_template` | [storage_class_list.storage_classes.hpe_storage.protection_template](resources--fleet--reference--group-002.md#canonical-1130323311201032-1120010203311003-1323331001033313-2201120330021223-0312121332110113-0001010113003223-3102232331010231-3123121123013300) |
| `storage_class_list.storage_classes.hpe_storage.secret_name` | [storage_class_list.storage_classes.hpe_storage.secret_name](resources--fleet--reference--group-002.md#canonical-0303123221003200-3100113303333222-3112311000012323-0222213103100310-1201100100133021-1001010210221121-1002012323101012-0310110101333333) |
| `storage_class_list.storage_classes.hpe_storage.secret_namespace` | [storage_class_list.storage_classes.hpe_storage.secret_namespace](resources--fleet--reference--group-002.md#canonical-2021201030233111-0002201221112103-2101222223120030-1220313023023210-0200030331033230-0203023022310311-2112201211212102-0020331232210010) |
| `storage_class_list.storage_classes.hpe_storage.sync_on_detach` | [storage_class_list.storage_classes.hpe_storage.sync_on_detach](resources--fleet--reference--group-002.md#canonical-1221100221101202-0002132220221330-0232110121112122-3212132202120000-2321312103122110-3210013200122132-0323112100110030-2020031100302300) |
| `storage_class_list.storage_classes.hpe_storage.thick` | [storage_class_list.storage_classes.hpe_storage.thick](resources--fleet--reference--group-002.md#canonical-2103130213311223-0331203131333031-2012120002033221-1020313133213322-1033323230020123-1001123233122311-0010201322001133-1111122310101123) |
| `storage_class_list.storage_classes.netapp_trident` | [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-3211323111122200-2022231133310222-2111133212330300-1010113231101330-0021303132000002-3100101322300112-2032232202110332-3231102032113310) |
| `storage_class_list.storage_classes.netapp_trident.selector` | [storage_class_list.storage_classes.netapp_trident.selector](resources--fleet--reference--group-002.md#canonical-1230001331232023-1121200110230223-1320210312010322-2222111303221202-1030032033322312-1213022300230231-0123203120230022-3133132202222111) |
| `storage_class_list.storage_classes.netapp_trident.storage_pools` | [storage_class_list.storage_classes.netapp_trident.storage_pools](resources--fleet--reference--group-002.md#canonical-1323200222023300-3202123201121232-3200302013333131-1202122303320021-0313200002321110-1313123332203113-1103122333132232-1200231333320033) |
| `storage_class_list.storage_classes.pure_service_orchestrator` | [storage_class_list.storage_classes.pure_service_orchestrator](resources--fleet--reference--group-002.md#canonical-2102032013030302-3313002330122221-3211033121302013-0310311210011323-0200201031123010-2233100223203002-1333301133030033-2001102003233021) |
| `storage_class_list.storage_classes.pure_service_orchestrator.backend` | [storage_class_list.storage_classes.pure_service_orchestrator.backend](resources--fleet--reference--group-002.md#canonical-0210010221120302-3002023220030311-0113300021331233-1333323112113322-2210123130321230-1122320202312310-3212230103230113-3032301202032312) |
| `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit](resources--fleet--reference--group-002.md#canonical-1230013231223200-0211121221313133-1122130202323311-0030031320222111-3200232032013311-2022311232030123-1000230010020201-0003222313232120) |
| `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.iops_limit](resources--fleet--reference--group-002.md#canonical-0112023103302112-1303020021023310-2211221203122210-1211032313103023-0130211133203211-2302013020100201-3121003232131030-0311012310302131) |
| `storage_class_list.storage_classes.reclaim_policy` | [storage_class_list.storage_classes.reclaim_policy](resources--fleet--reference--group-002.md#canonical-1221330002110232-1330210010213303-3122131210210101-2323203131323133-2210223112321131-3302203033113122-0303333011323223-1330111211100203) |
| `storage_class_list.storage_classes.storage_class_name` | [storage_class_list.storage_classes.storage_class_name](resources--fleet--reference--group-002.md#canonical-1122312030110112-2123300032021100-1302131230102013-1010202122131123-3030221101211202-2200120003120302-2321320302132103-1322300000230231) |
| `storage_class_list.storage_classes.storage_device` | [storage_class_list.storage_classes.storage_device](resources--fleet--reference--group-002.md#canonical-1011230231022323-1012133323300310-3301103320223322-0031203100122123-2023112220210130-0010032230233322-1221322301321213-0131000113332123) |
| `storage_device_list` | [storage_device_list](resources--fleet--reference--group-002.md#canonical-1220331113303212-2010311203330203-1123010023313010-2333311011310323-2013213130203030-0232022330203123-2033201030211020-3221200223133131) |
| `storage_device_list.storage_devices` | [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-2303022122102011-3003130203301312-2102021200031130-1013321101232010-3022110212312002-1300222232122012-3110113302112132-1103221033211211) |
| `storage_device_list.storage_devices.advanced_advanced_parameters` | [storage_device_list.storage_devices.advanced_advanced_parameters](resources--fleet--reference--group-002.md#canonical-3310031021310201-1130031213203122-2222103030112210-0221123102103022-1031001031020313-1321333202030101-2223223023022022-0123002131313001) |
| `storage_device_list.storage_devices.custom_storage` | [storage_device_list.storage_devices.custom_storage](resources--fleet--reference--group-002.md#canonical-0100120110322223-0112330323012102-0130122221222112-2000033202332103-2211220000120230-1202321200321112-1220322023101100-1233121231010130) |
| `storage_device_list.storage_devices.hpe_storage` | [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-1300123020201210-1210323131311333-3010122233320022-0003003322221212-2133201201131013-2033320331211322-0332300022020123-2323012132313022) |
| `storage_device_list.storage_devices.hpe_storage.api_server_port` | [storage_device_list.storage_devices.hpe_storage.api_server_port](resources--fleet--reference--group-002.md#canonical-0100121002203112-3202321112103103-1303012230123123-1032323131130213-0012031022031020-0232110000113021-3232201211102332-3120202220033001) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-002.md#canonical-1333200000001010-1022131030213331-0211121200013033-1001011211103100-2300121021230232-2202133313323302-1103002233203000-1032302020222310) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-2011203201120113-1011203012002231-3321012033222312-3302200210000033-0101320221330033-1310220100320112-3313103113220210-3121102031033020) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-002.md#canonical-0310223111332103-3122221332110000-1101112020123021-1003212223101213-2222012210003310-1121120200303201-2031233002322332-3013330032133300) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location](resources--fleet--reference--group-002.md#canonical-0221310303202313-2121011001201312-0323032021231120-1231303310202303-0000212233133322-2133230331030213-3030000310233012-2130200100331231) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider](resources--fleet--reference--group-002.md#canonical-3311001322212321-2333002021021222-3322120012331121-0300200100032020-3102113012132020-2302233233033331-0112031130113223-2031302102120003) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--fleet--reference--group-002.md#canonical-2031303221102222-0030111031002030-2333101200112203-2000220011111223-1013231103320302-2220212302122131-0020002230132303-0210001210303021) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref](resources--fleet--reference--group-002.md#canonical-1311031210011103-1312322211013031-2110122020221302-1303332022002303-2222122113131032-1123321300103331-2320031132310001-1302302232112203) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url](resources--fleet--reference--group-002.md#canonical-0023011130010020-0210202013100310-2302312101333331-3012210122330131-3330211000123200-3131033000332313-1331211013220023-2212023302123100) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_user](resources--fleet--reference--group-002.md#canonical-1113023311332300-1111230300333131-2323221112111110-0103131033112032-2210122222300322-1023223200313311-3010102233111331-2020111222300303) |
| `storage_device_list.storage_devices.hpe_storage.password` | [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-002.md#canonical-2131003020012222-1312020222122031-3033303322002020-2203010302123011-0230321123321121-3033331222100311-1330300022110313-0030033002310200) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-1311211111133200-1013113321000123-1121102322020233-1202023132103002-0211322231023231-3112300311030220-3000321101121103-0111312132002321) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-002.md#canonical-1230200211213203-0023103220022320-2012121111222312-3202203202002331-3213332032113330-3112200111232331-0323033022111020-1122232100033103) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location](resources--fleet--reference--group-002.md#canonical-2010013323212311-0312100003202302-2313221011112310-3221130010001122-3331321232001000-0013002322330120-3311002013321112-0332130332112010) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-002.md#canonical-1331332100110312-0032030030111020-2021232333003123-0233203113223031-1333032232123212-2020111333300310-0020130112231033-1230102311132033) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--fleet--reference--group-002.md#canonical-3312223300230222-2221102330132003-3020131323013102-0132313003022123-1232103332312031-0003013222032311-1012223300320022-2031312120310132) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref](resources--fleet--reference--group-002.md#canonical-3122220201313000-0131330200201321-0302111212222202-0301312121203231-2002330300211103-2030011333102333-3101122221223110-1111212300012332) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url](resources--fleet--reference--group-002.md#canonical-1323030021203330-1133231111120013-3020313012220331-3330122100131201-1303332322133232-2220220002322022-2230322100011010-1000202301231102) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` | [storage_device_list.storage_devices.hpe_storage.storage_server_ip_address](resources--fleet--reference--group-002.md#canonical-3130131021220103-2310130023223033-1212220103022231-1112133102113030-2332213330212211-0313231133112020-2322303022023331-1002122021033120) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_name` | [storage_device_list.storage_devices.hpe_storage.storage_server_name](resources--fleet--reference--group-002.md#canonical-1032221020101332-1220231120222122-1123312121002001-1022131000123000-3021310100101310-1101203103020213-2020210021223303-3300013310231100) |
| `storage_device_list.storage_devices.hpe_storage.username` | [storage_device_list.storage_devices.hpe_storage.username](resources--fleet--reference--group-002.md#canonical-2121200030301312-3111212101221202-2323001030322113-3033230130121230-3323321113002213-2200201021313103-3303013001302212-1220303033211111) |
| `storage_device_list.storage_devices.netapp_trident` | [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-3101021010223311-2100233103222002-2031222103120330-0032130300130100-1032231200121110-3313101222102111-2301311133330031-1001330210033033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-1300011332320212-0313000331011031-0302220301103130-2120310231022210-3113120211122333-2233132200230103-0133012320023111-1003312123221132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--fleet--reference--group-002.md#canonical-1021110303213110-3302031121103320-1332312323011213-2212202132032211-0321032102312102-2311210101332031-1000233032220111-3130222013121311) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes](resources--fleet--reference--group-002.md#canonical-0122131302011200-2233200123103300-2303003303103203-1112113321200021-3323333011212210-1030320002323023-3103333102210313-0231120222201103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy](resources--fleet--reference--group-002.md#canonical-1231313321302020-0000021011303123-2330330331320010-1223023221101211-0133213231123203-1131200031300320-3033310031323100-0112203230012230) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name](resources--fleet--reference--group-002.md#canonical-1210031232131302-3211232322012102-0312312001221211-1303202120233112-0220230011320213-3122111223212303-3210013310230200-2023311101200233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate](resources--fleet--reference--group-002.md#canonical-3001312332233233-1101010101222020-2220131133121111-0220001230130021-1011023110132210-2331213312311031-0033300300221311-0000100202000220) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-002.md#canonical-1230113112310231-3223032113113331-1001113333123330-1010323232000312-1233301023032323-3102323330023311-2310212310031311-1033113201201300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-2110132010000223-1001012310202032-1013231232032230-1202010011133322-1020200101303101-0222211103111022-3123010101021331-0200321000110323) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-002.md#canonical-0200303202212032-2320201022103332-2310031120131210-2113231302023112-1222302000123023-1101022303131223-3002011223210222-0101310212332331) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location](resources--fleet--reference--group-002.md#canonical-1003221012303232-1222013221030303-2020113122130211-2322212002002133-1131111002032003-0302123331322131-3033212330231330-0322113020131300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider](resources--fleet--reference--group-002.md#canonical-3110212323130123-3121320311010202-1201103020021211-2333121323312300-2102201223211113-2103331002020330-2003123210121313-1300101120112022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--fleet--reference--group-002.md#canonical-1301333012002230-1211121232220303-1001132330013210-1110312130033102-2111003110212300-2213310312211100-1210130322013100-1303032200332220) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref](resources--fleet--reference--group-002.md#canonical-2212230131002122-2032211110233122-0111313302333023-1012210200303320-3113023123123120-0221220202333122-0132330021020022-2221110303232020) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url](resources--fleet--reference--group-002.md#canonical-3303021030021130-0302233013003223-2020021202013333-2222030310200212-3311333212013013-2032302133233233-0132212230300331-2310000023023131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name](resources--fleet--reference--group-002.md#canonical-3210222313333120-1212302323213103-3133313032220232-3301321131112311-0131030311013101-2201212122100130-0301203031102202-0010102220002212) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip](resources--fleet--reference--group-002.md#canonical-0011310011100021-1012023130231030-2300322111100111-2311310221223212-1232220332133103-2331331030102111-0301313200131123-3302320021212020) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels](resources--fleet--reference--group-002.md#canonical-1131321231021330-0321000230001110-2123300312321330-1302102112312221-0131313231010120-0020130003000031-1010310322112033-0132022123002232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage](resources--fleet--reference--group-002.md#canonical-1212012020302331-1100101013232022-0233323020333023-3322001031010031-2322102212323002-2030032231220110-3101022310322303-3320002213012332) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size](resources--fleet--reference--group-002.md#canonical-2012300211032020-3100103233101013-1112312020230110-2232202023230000-3331130312300221-1031222231120030-1122333202112311-1311303202112303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name](resources--fleet--reference--group-002.md#canonical-3301201200012332-1032312121230011-3010112210201220-1002102302101201-2122102002110231-1013031101220221-1011311032313100-3022100332012303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip](resources--fleet--reference--group-002.md#canonical-0130113221022031-0231200122210132-2030110123103212-2221320021223333-3023300223020032-1213220212331122-3020101001201101-0110330111110102) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options](resources--fleet--reference--group-002.md#canonical-2322331212203101-2032212003322321-2303333321222121-3213010111101213-3103221220001131-1310221300121132-3131331323311222-1000323120210310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-2310201301301300-3100121212331033-2232033100203311-1330302310003123-0232233021333003-2200330123100132-3300122313002212-2100311033121231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1320123220002003-2203232200231032-3301031122212302-3122003300332103-3303331233110203-1302223022222111-0030311012311010-0120002332101312) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-0012222303202103-1030022203232022-1031300000330331-3322130231130130-1011212201110012-3311100003331202-1323301110333310-0112212110323021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-2001103321320003-2233120032302001-0120133101032131-3032020323303122-3122021213013330-1202203022020033-3102100120030100-2211333311110001) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-1321301321301330-0020122313331331-1032310130023223-0110301023103323-1133013313221230-3212331020012300-1311313200223303-3303202300011212) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-3233010122021111-2130130130011012-1321232312010131-3002100011220120-2133022111032011-3021203301111201-0103110110303311-0130302113321132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-0023330102210212-2031110120200133-2030201223100222-0230332112020120-1013120322220231-1033220312320023-0103000200203101-1131230322332032) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-0133202232200021-2122231020032100-3321302331221321-3211120221131301-3301322333210300-3332203210210032-0313102013122232-1332032021323233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region](resources--fleet--reference--group-002.md#canonical-3331213221131132-2023230121203212-2003000331121231-1131000100303122-0203332230231322-2211313103200331-1012110002220312-0000233210031123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0323300303111030-2320210301123301-3031331000201323-2102132131230001-1220023300211123-0002211023213130-1220311130031020-3230300211333321) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels](resources--fleet--reference--group-003.md#canonical-3113032131030330-3211313202123032-3130332011123121-3123100200302130-0300133332323112-1321320130121110-3022110213032121-3133322210232131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-0000132022102330-0231030023212231-2210011210103300-3232131333032210-3120303013311202-3023012311313331-2013301002010120-0032033232033201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-1123313322201132-3201103312320230-2112321023323332-3010202112020231-2230023320323011-0200032231311022-1313231223302002-0201323330301231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-0330202233301311-1330201101233130-2220301010032200-1320002200310331-2313012333031101-0332031113033333-3321200312131200-2211311300002132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-2113330322231332-0002120111033202-0232012320313313-1320030220012233-0010110320120211-1033203000212023-3003031202332313-3112331100022131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-1111301331220233-3312003223022320-1323212032312332-0221033022323022-2221102121311121-3100031312312222-1233221313001301-0113103201310023) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-3100311033130012-1330212011302313-0200201110132231-1220133130000032-3130332230222112-3310301121023133-1220300332232203-2231122220320232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-3003310223321330-1332322213010222-1020201223033002-3233003103223001-0120020231233311-0202101321002001-1020020233221312-1302123333131101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-1320011202312013-2113302320312302-2010221001130331-0033120100033000-1130231311321000-3102210112231002-2012032212022130-0233312020303213) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-1311320330221233-0121331102122101-1333233233330300-2230132222032330-3320023113301320-2003201101103030-0122102131313332-1232222032330132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-0321101321233220-2333320230121301-0113301002332022-2101121111033022-0203110011202231-3110330303303123-2022122131222231-0003022301022130) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-1021021101001221-0200103022202113-3302100233210031-3302013022322301-3131231330323130-3112200233020000-2002202232320211-3210230303322003) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-2323101023330100-0201201012330331-1022232122101132-2133130323121332-2321330031110203-1330332310211000-3012102123232202-1213022303011320) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-3100021310223013-1023112233120302-2210203030020022-3030200301112000-1222010313313022-3333130133030001-3312101130231011-3332300300113000) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-2322020110010321-0033032323201203-2002323010021303-2330031222120100-1300330033230020-3102212133332221-0302310011312332-2331113333130203) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone](resources--fleet--reference--group-003.md#canonical-1111112113201000-1303200213203101-2032032301333011-1311133030021232-2323101133113011-3032022201201201-2111130221100323-3020103122303103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name](resources--fleet--reference--group-002.md#canonical-3030322130013021-2111301300020323-2133223002133323-3201132312233323-3213023110301130-2011320110230132-2031332112331310-1311031120333330) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix](resources--fleet--reference--group-002.md#canonical-0232321133002011-1032223232113020-2012130130032112-3231232020120031-1021011101313130-1332203011210311-1333233012303100-3331303030013213) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm](resources--fleet--reference--group-002.md#canonical-2100310330303223-1020310133303103-0313303300201020-1332121130232322-1111111122320223-3321120302111311-2111232030101112-2002202020311133) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate](resources--fleet--reference--group-002.md#canonical-3320120122022020-3010323023113330-1330001220230111-3323222113112202-2330103233020310-3103030300323011-3012010200313232-2330213320122230) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username](resources--fleet--reference--group-002.md#canonical-3012113013132330-1331210322023311-2031013220130030-1123311013230221-3210313012301333-1020213220003201-3030032222232121-3000022333031310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-0232310013002023-1113101303132333-1223201012331133-2300100203120113-2120003011000200-0120111033032303-3213002011031220-0003323310003012) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-1301310112321123-1103101102100223-3100202111102313-0202312230332100-2131202310321331-2222012023131302-1031000111121200-3300130223310233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-3221220112110001-3210323312300332-3302333312120201-2002213222212033-3212132112031003-2232002033333010-1312213200011310-2331002003110322) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-3313032321030203-1012101202311213-3012030233011233-1122321112330110-1130311331233320-0030300011320130-0001333122033201-2312300201123302) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-3002123320300132-0223331021300220-1210103031121331-2123031323213111-0011011022012122-0212310103013013-1233320313131233-0232112111213103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-1323211230012022-3210030032230001-0121213131022133-0331222030031312-1210123012231023-0220320201230323-1230232023110013-3231022112022120) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-3000210100233320-3030321021200212-2101320103302230-1330101123200012-2003021123001123-2211111001330002-0130300331332130-1223122010322122) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-1211301000102130-2231000200120211-2303313003120220-1120021301230112-0032110300130313-1031332120123030-3132010230120000-2332203302323122) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-1121311121212002-1200120212002203-2011102120031103-3003230220210022-3031132232221101-0330323010110301-0031333103330131-2123212111322232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-1120023203110310-3021011102210320-2102213100322100-3330210232013310-3312112123230333-3332313203220000-3201202133233133-0203332021312010) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-0322200131020031-2002313022011012-3303323230201123-2210023303232110-3000230003033033-2312233331230302-1032230031232220-3101130200313311) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-1132321301130121-2103311321210130-1032011211333013-1002211123312233-0302122032010210-2210313211023200-2000232111103331-0311122212123232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-1301212333212322-2102120103213022-3201002132131123-2120133302011010-2132100320233101-2100202331112211-3123011111112122-2212213303320211) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-0223111201001002-0002122121223333-0020002021122221-0312300022013313-3301031030332211-3212010200321132-3200130031300213-3001322112303330) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-3323210230230230-2010001013033130-1322132130110013-2001110133212220-1213311130133212-3130321202013113-0323121301331013-0123230022003303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate](resources--fleet--reference--group-003.md#canonical-1333220200332203-1102231001200321-2000221232122022-0220130302101023-2200231001302100-0232000312111132-2323131321132010-2132133033011210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-1333023203233113-2220031121323220-1231121203300023-0122013010232003-1220202332213131-1103300033133200-1102230102012002-2033131210130003) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1032101211332033-0323320121220133-0030003312332200-2302232033012231-1211320312213223-0122333220020121-1120010231313222-2211123111103123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-0002330000000333-2100323201331121-3002321101203313-3110102123200333-0020000313303130-3202021233232233-1320003211001021-1200110222210303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-2032213310223333-0013311123223100-3111132102303203-0003322311320032-2030321112030100-1320320130323200-2331310121031331-3330013320122121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-3011122020112110-3020031131323020-0202000100122231-0233121000100330-0133310023323323-2133232223202010-0331030103132223-1231200221330320) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-2111210331213021-1030012213212013-0130210111021323-0011121310030321-0303010302333230-0100302222202033-3102121113230231-1100000232103210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-3131103210331203-2220030230322322-1121010313023300-2012302011020301-3012223303213311-0020331021021002-0213202202010123-2102302102301012) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-1332321023113022-2321210211203213-0313312231212203-2230333131200033-1303333013320213-3113020002212121-2313230130000301-1302213313331211) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name](resources--fleet--reference--group-003.md#canonical-2002111031200132-3311123011032112-0110322030021311-0022332123323013-0011331321131221-2231102003123131-2122230212233012-0110011212102121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip](resources--fleet--reference--group-003.md#canonical-2012111013231313-2100331100002022-1320213312331231-3222113322021232-2110300200223101-1333331120312331-3313231120310200-1322001011321300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name](resources--fleet--reference--group-003.md#canonical-2012011322130221-1123031013002011-2000133311221101-2123111023203221-1323332001233002-2223010010032112-3021000003310000-2111333002013232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels](resources--fleet--reference--group-003.md#canonical-1233203030100212-0313212330132230-2231021101303301-2310213112123011-0033133021101101-2300112212003123-0003120112302133-0101311231131222) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage](resources--fleet--reference--group-003.md#canonical-0131320010213210-1220000033003302-2312203332003003-0310033320002123-1000001013000200-3230032223120022-3033321322222013-1013300112103200) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size](resources--fleet--reference--group-003.md#canonical-3122211201223003-3021303220123210-0331301012310100-3103110210011103-2133220222313213-1300331312220101-2311131102010323-2330002322220121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name](resources--fleet--reference--group-003.md#canonical-0213102023320321-3313120133230223-0212213103112300-2103011130203320-2022033012100132-3122321200202302-2221130122322220-0323302313330002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip](resources--fleet--reference--group-003.md#canonical-3321302332003231-2020020100012231-3102013012200321-3200033233110313-2202220203213011-0031320333333023-2202001110113300-2210223223310110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--fleet--reference--group-003.md#canonical-2312231103301222-2303013202220233-2012103312203311-0302300223213112-1322111133033332-0310121011022233-1331102122201223-0120313021103321) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0311220320002233-2031131013333113-1230200132010013-0120222121311103-3212313300023203-0320022012312101-2320233010131023-2113321121332123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0003322203023300-3012202302303003-1011322203212103-3310323103322303-0113121312302131-3113310101110330-1310000101320233-2103300022331221) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-3030322023203110-1012123230012320-1323010223002002-3311300301131233-0310313202312322-1221112030332123-0020311202111011-0020321231023301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-3123003120331232-2033301223111112-0321030133113031-2231212012322030-3001101302101123-1103033210230203-2012230323123230-0310203333230223) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-1300131113120232-0222132210030312-3023012320330301-3111101130300101-1022202121022013-3213032110331221-0011203330223113-0030301212123312) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-3231023312310010-0111102221002032-2121103001212000-2333130132120032-1223001032131223-2330230220233221-3021000010120311-2003030221203101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-0130220311301111-3201001110102323-3311300222113302-2322132000223120-2330111012311320-3230003321112012-1301322313221213-2231333232022303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-0001230103223201-3103313212232313-3303202101333031-1332331332003120-2231101233330301-2332102301202003-3113223013332031-0323003011302313) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region](resources--fleet--reference--group-003.md#canonical-3113021301202103-1120203201022013-3110131321022332-2321010111201202-0220123202230130-3133102021113222-2333233013003331-3321202212303110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-0211032110321131-3312003313011133-1233331030301322-3312031133133231-0213230312110110-0032012133303202-2321030123303101-2333231131032301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels](resources--fleet--reference--group-003.md#canonical-1031323230331001-1101112001113021-3003130101130201-3212102133311110-3211331013210001-3131231313220311-3011323000230200-2031233023110202) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-2301022122323310-3321013323230120-2303310023223021-0120321200233130-2001120302001031-2010011333030111-2013212312311212-3000313000023301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-1022212002223031-2221112020301210-1303300102113000-1020303122020010-3100003102012010-1031233301131313-1110033022022003-2130313101303201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-0330330022322210-3301101000302032-2132023210033101-0223111032030301-1030330213312101-2232021313233320-2330102021023110-3021223023131222) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-2213021120201112-2003101113230331-0303213102200212-3002333313221013-1123212231202012-2122202023122230-2331112010023211-1210230113103131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-1133003221223230-0303032211031120-3022301113132322-2030233300100020-3303121320031231-1323233311120011-0123023231130030-1111232031201013) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-3013310231230110-0213300133130022-2011011333202031-0123013132133012-1223333022121122-1210300330233101-0323113022030300-2312011313303321) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-1310101232210213-1000222222131132-0133121323130303-0033233011121323-0210133022201023-2201322302232202-2303131002101211-2212100123203302) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-1303130223112010-2020311230033012-1010200302212311-1132131223113010-1000002311211003-0230333331023233-1010210321021131-3321031221021130) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-3122312112000331-2013032020102312-3301120311111221-0323122213002000-0331132232122123-2003222330221131-2333313311213132-3110321122023323) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-1110102000233101-1223012221010300-1222313202211233-1031320301000003-3213331313022303-3122201313201001-1023112303020202-2221100201010332) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-3112103233200230-0313213130201003-2202123221301131-3121223010212031-3201233333013003-2123223212023203-0111113312033301-1012201103313230) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-1103301002033212-2022222310112100-2131310122213301-1032320212232132-0220220222311223-0132112112313211-2132021310211011-1203222302313101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-1320221323223103-1110012330110123-1023130132003333-2223102212222231-2100010222003201-1332213203222210-1202023032120332-1203033030030230) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-1012021000321233-1321020100222120-2001333133203313-3231332221020110-1300313202021020-0101210330030221-2323312100112121-2111332012010101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone](resources--fleet--reference--group-003.md#canonical-2223300130001123-1020221130012330-2221203113311211-2110033211130301-1022321230301010-1032100301011002-0030213103210201-1220030110231322) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name](resources--fleet--reference--group-003.md#canonical-1302220010102231-2313311230120312-0112100020331100-0031021323322033-2010110021211021-2302020303002012-3330323302221112-1300212220332022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix](resources--fleet--reference--group-003.md#canonical-2201002330330333-2002100130001202-1132110133003112-3313102201121221-1002332003103013-3111001332130233-3321230200111013-3100303203111031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm](resources--fleet--reference--group-003.md#canonical-1112213220332103-2331030233202121-3221310322030302-2203033311101133-2103220223012231-3222232222000313-2230203211003112-2022220020030332) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate](resources--fleet--reference--group-003.md#canonical-1030122212211303-1033122101021312-1012213300312023-0001002122203202-3300123202301131-3100003011133013-3321313131100310-2022013213311032) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-1133032223232231-2303221011010102-3232213321303033-0003320130211121-1213101030131133-3111310233100132-0202031331022310-1132230120303303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-3003210202000223-3110123111132321-0210001220333232-1120002022022300-2031300223111011-0310211303323230-3333312113012331-3300210120023030) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-2113232130020120-0223210300032311-3132221222002303-0132322012300200-0131323110111002-1201130222111302-3203332332230022-3320010322120030) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-3032211003021020-1220203130030132-0231120122110112-3031012302030033-2231033103200320-1211110322300222-3320313321001130-2110130112230132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-0333031232023020-2300221312011131-3311310322233330-3313123300202200-1220322120203322-3200321121333320-3212011231100220-0030022011002300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-1112101223312123-0311100031331300-3100300301320122-2321313213312202-3013123313001212-2221030222100221-3133232123233302-2200000233113123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--fleet--reference--group-003.md#canonical-2121012011320310-3203302322321322-3320323000202022-3210323223000022-3303311333003122-0230013331113311-3323220003110213-3103032203021100) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-3103010221321100-0210021113231222-2303033220312212-0001233331300021-2302000112313100-3320323310012201-2113200122222110-0302100010010021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-1031330210211321-2002001031131302-2303323000203302-1332302210103030-2121210100010110-0202233220101212-1003221103212111-0120211320132313) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-2203212113102333-0300233230302032-3020200303301320-1111120221332002-1121330102000322-3303213311130103-1132002311232013-0310303110320301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0120222320012302-2020232212300003-3020323130332312-1131223130221321-1020133103220310-0110120013203122-1031020032222222-1100203201303233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-2001300022010103-2301113102302022-1203013311032231-2102001232100033-3230000012233331-2100210003103301-2110011033001231-2223311311022012) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-0201020020101302-2121100130211133-2212110022221110-2032123110102331-2230212013002200-0313011101023322-1000332301320001-3323023002211022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-2133300000223331-1030211311021131-0210003302231120-2012123220233233-2013303132221331-0030120223322112-2303333133021203-2130320122220013) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--fleet--reference--group-003.md#canonical-0221032110302003-2203100123000322-0113102310010022-3032200301003232-1102201030302323-1312002231022000-2332320031113230-1211313112222221) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-1222023023332333-1133011301320210-3023213132101332-2030200230222032-2222323230322002-3113031001230323-3001320110000331-0100111311312113) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-0231303220332321-3212002132011112-2103003330230313-0312203002111232-3030123332200132-1311232000300002-3112030220033031-2200031131311102) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username](resources--fleet--reference--group-003.md#canonical-2103013301232100-0221031213132010-3123220101022120-2203113032322203-1112313132030322-3112003331200031-0212031321113300-3102321031011013) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username](resources--fleet--reference--group-003.md#canonical-2100321023203102-2230022130112113-3022031232233322-3320201303330133-2001023212101322-0113202323301113-3121233032122022-3110330200112013) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username](resources--fleet--reference--group-003.md#canonical-1022321112232130-0311002130232210-1330122220010222-1333111013033232-3120012303123320-3131322011310001-3021330000220020-3022323320101300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-003.md#canonical-2200223213222001-2013131101001201-0220221003001211-0131132313010112-2131033300030311-3210211013203310-0221033211333312-1031020230330220) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-1233033221100101-0303200212223302-2132120001021223-1231200232003112-1332110220030022-1111212332022202-3321223233033102-2002320232200211) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-3320320012112032-0030112300230030-2112321013123101-2333321200231003-2223222311320020-2212033333020212-1301323232201011-3120200120301201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-1102122032033303-2133123121332312-0021213302202020-1322020301220200-1333122200020220-3212003302003012-1020012302131101-1013001100013200) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-1003221313203203-0131000233301323-1320110301120322-2003020321330111-0302303133000000-1323012221101002-2031232212331331-2101022120221112) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-0322011223022022-2122231232231022-2113122011130013-0003130232302232-3131012012212333-3112130102112310-1000001032101111-1211012110032121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-2310333231000300-2332130023313201-1210201002133133-3333312101312230-1313323200001012-3001003230021112-0312113122301220-2312100020303030) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-1321313120000012-2213101203110002-0010333033021221-2221120232032202-0212123013233001-0133033113221203-0231213032112201-2100230312121130) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-3100000203103031-0331131022003323-3030213023001223-1211023230133200-0233220233201212-2231031303201012-0000131323210222-3023233330302303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-1013120300222302-3132231001320120-3212210101231302-0120213131230323-2300120102012302-1203203302212302-1322021302032223-3301312131013202) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-0333120331121113-0123202130122202-2111302020132003-0031210101222303-2133130113322221-0000002012313012-2111201201112223-0301022200012220) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-1231010021221211-3031113203120211-0222202200311101-3210332030213131-1011031300223231-1022133220003102-1011123302321032-0301312000001233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-3220300000313321-3223310230331233-3213010031032332-2331332203203303-3113302132120231-2200112101330032-3120221301220020-1203030000112031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-0210210300313230-1221233221203300-0213323321030210-1133121200213230-0213231212020112-3123122002133301-0222323031331110-1122102332200212) |
| `storage_device_list.storage_devices.pure_service_orchestrator` | [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0110313201102010-2001230300022320-3333102331213003-3202322213123220-0330311102032233-2011123003100323-3123312222132201-0102023231311000) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1222201323200223-0321321103232112-2101102012212110-2230211100212302-3120201131012133-0122123320122131-1120031330230223-2222131333102002) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-003.md#canonical-1021111023131203-1303203112233301-2002220300323203-0300230103101022-2012023323222203-0330222323321201-3101203223223121-0010130302200230) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt](resources--fleet--reference--group-003.md#canonical-1021210010230023-3110332022313103-0330002032030330-1201301332120303-2101103122320113-1211120320100121-3322012200230220-3103203231312113) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type](resources--fleet--reference--group-003.md#canonical-3021101210023032-0030112230331303-0101031312121010-0200120300211101-2102212101323000-1102303213301030-3333013033010303-0031200001013301) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts](resources--fleet--reference--group-003.md#canonical-3031023332200232-0110013010310002-0002113023012112-3021213302323001-1303203102112121-1110331310112023-1011101303313303-1020211113210020) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments](resources--fleet--reference--group-003.md#canonical-1132301300033232-0001123222301013-0311313132231011-3210112211231032-2211122112231121-3022231201033220-0111133200112021-2120130103132323) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-003.md#canonical-0222320010220211-2123002013303203-3333301122231310-0230300003321001-0202302201110212-2133200332011323-2133311022303221-3103021322023333) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-003.md#canonical-2020310001212022-3022231101313223-0021030012120323-0302231322202311-0200011201322302-1030200300011003-2301112303023011-2222020120131031) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1213302211323001-0300033000030021-2313321101022102-3023100122201031-0213010332101123-0013320013222202-1021030031223233-3201100033233203) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-3010123102230001-2311312122221132-3330323312202313-1213333122103111-0303020031330312-0013221011310301-3232302103033233-2103011110222012) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-0003303203223300-0111032331313023-1111311322010212-2110032033213231-0030222333022321-3210222101102221-0010231323013001-0333130001121203) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-2022102110211320-1003013032202230-2103231311102331-3302331210132020-3331332321000330-0222331130020221-0203230303332331-1113123003013023) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](resources--fleet--reference--group-003.md#canonical-0302301301220231-2002130230331230-0332230330203230-1311211111121333-1202101300202213-0210012130101120-3012202320123011-1322033003302212) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-1200311110032231-3332023103310203-2220030130233111-2032301203230010-0312331031112031-2030100133230210-2331032020110012-2003031210021020) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-2132302021230031-3103130211103231-2033111200331322-2200100221130211-2333320013023011-1221311033220011-2332010310311111-3122130201001100) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels](resources--fleet--reference--group-003.md#canonical-0122012211131120-3020111130200102-1213203032300133-0331122103010303-1223122320130133-0312202213031013-2320132320012002-3031211203120012) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name](resources--fleet--reference--group-003.md#canonical-1003210222212020-0102012303200303-2202312332300003-3331202221111232-3012003023001333-0112310332221303-3213300200003213-3220323013210030) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip](resources--fleet--reference--group-003.md#canonical-2331110210223122-2033232302321203-0123130303121203-3033201113013303-1330103011131011-1232000023212220-2222122313303312-0321220301303003) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout](resources--fleet--reference--group-003.md#canonical-2321312103123003-2121222313033232-3021132110212020-2213112330120232-2112120020010112-0102130303323012-0110312333121031-3110033103012103) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type](resources--fleet--reference--group-003.md#canonical-3113010312330010-1223310023332132-3301111002031231-0222212331222222-2231013120010021-0230020020333013-3220330123203220-3232022203031133) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-003.md#canonical-0023122320012230-0330131303220123-3103002103331211-2133231322302301-3110121323320011-2332132202123020-3033300223100202-1333322310000301) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory](resources--fleet--reference--group-003.md#canonical-0013320101021123-0320312211220000-1132302220011111-2121032130100310-1203010322101233-2231010202303131-1013113300001232-3101212002120211) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules](resources--fleet--reference--group-003.md#canonical-2020221110230312-3020310023221101-3121230311100110-0221130201222213-1202312132331320-2222321123101130-1032332123023330-2212200101130303) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-003.md#canonical-3130013232110200-0123201330102301-1123020122131022-1223233130012333-3001323030121323-0333210213131330-2220112201323123-2331111121123001) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-003.md#canonical-1003112101311133-0022112320333102-0111000201022221-0011303200023232-1033203203300013-2112233032322202-1311100303101011-2322111330133322) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0211300102302120-0302002001203300-3210221310310133-2003003030310223-1031330020233101-1210130203313200-3102110100222210-2133300230011003) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-3333112032131022-3022313202022220-1133033323221303-2233332232233032-3101030233023012-2101113013001222-0021011223312320-0011002223210323) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-3320101031133313-1111110203102032-2103302011111302-1003300211310030-1100211222003311-2223303110220212-0013322330200010-3221110212310031) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-1110303102000301-3023033332010012-2001302111122221-2130202301011230-0103213322221211-2313333121331212-3113223033300100-2111313030333332) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--fleet--reference--group-003.md#canonical-2013312123332222-1003010032313113-0303111133112000-3023211131121323-0211122130000011-3000021213211133-3320121202202112-2010302121321132) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-3012123101310132-0300310112132131-0130312102211220-0323101333032102-3133000201132230-0101010212320010-1131202333332312-1000332303030022) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-1030312111123001-1310310332331001-0322101122103310-2003003121332300-2331210231003013-1231110312233031-1310212312212110-0111001303032122) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels](resources--fleet--reference--group-003.md#canonical-2320012023001332-3021021123023103-2202201031100230-0302332101211000-3023213022300120-3003320133022221-3122300000203323-0113322213023003) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name](resources--fleet--reference--group-003.md#canonical-1301333311311000-0120022021123202-0002231233013310-1232332011312023-2203203202202203-1203313032323111-0231322303231133-2302322230023102) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip](resources--fleet--reference--group-003.md#canonical-1130222232301223-1010210003211211-3230203130201333-2331022332033211-0313113313332110-2111113322213012-3013110103112320-3320131310322112) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name](resources--fleet--reference--group-003.md#canonical-2110312131030113-1301300012011101-3023122011133012-3020133220032212-2011303112132012-1003103322230323-1202303303032311-1031232013023130) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip](resources--fleet--reference--group-003.md#canonical-3220011001133032-0323121013222213-2003312300331232-3312003120132001-3103133330001221-2002202312311222-0120001323213232-3233100030311032) |
| `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` | [storage_device_list.storage_devices.pure_service_orchestrator.cluster_id](resources--fleet--reference--group-003.md#canonical-1331110033133223-2300211312201210-1332001021233302-0122323110132320-3101013322002012-1001120102301233-1120123331333013-3302111032101313) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology](resources--fleet--reference--group-003.md#canonical-1111210320012020-1033020233321323-1131111131221231-3131112033331312-1321322221333323-2033021232211320-1113223021230322-3233011032033031) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology](resources--fleet--reference--group-003.md#canonical-1113301010311301-2202113322220212-1311203220223202-0212122320223120-0202033100323200-2111011232311113-0312022312033303-3333311032203332) |
| `storage_device_list.storage_devices.storage_device` | [storage_device_list.storage_devices.storage_device](resources--fleet--reference--group-002.md#canonical-1303002102023230-1212021212311133-3012312033032013-2102222101303330-1220300031023302-0213130200100133-0320110322222301-1232210033213310) |
| `storage_interface_list` | [storage_interface_list](resources--fleet--reference--group-003.md#canonical-2223313213113003-2010202220123320-3322302112013032-3001322211322310-2322330013330331-1010013302213030-2221310300112203-1322231330200313) |
| `storage_interface_list.interfaces` | [storage_interface_list.interfaces](resources--fleet--reference--group-003.md#canonical-3222302313101310-0000213310101302-1333001013221232-1000321010203200-2230333031200013-3320013322312231-0031030021031210-3131233203313131) |
| `storage_interface_list.interfaces.name` | [storage_interface_list.interfaces.name](resources--fleet--reference--group-003.md#canonical-0001010310110001-1200131311112010-2231313101122232-0322212100003322-3123320023203123-3301022220130213-3033322311320202-3102232200131213) |
| `storage_interface_list.interfaces.namespace` | [storage_interface_list.interfaces.namespace](resources--fleet--reference--group-003.md#canonical-3331312331220012-3303310000313201-0003322220301233-1213002222333210-0023023023021011-2130220012020132-3323310110113311-3012111312100313) |
| `storage_interface_list.interfaces.tenant` | [storage_interface_list.interfaces.tenant](resources--fleet--reference--group-003.md#canonical-1112313210200113-0201231203031213-0311323110130011-1311330212123223-2023110213000013-3222023102113131-3022312230222133-0003220101330101) |
| `storage_static_routes` | [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3300220221311210-2223123102101211-3220330103110130-2013011230213113-0302003112110032-0211100233331131-0000213033313103-2010020102112211) |
| `storage_static_routes.storage_routes` | [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-1113103203120230-0103213211302232-2101121302323232-1122111212313231-2311322311200031-1121301130131233-1130133302322133-3311302231122210) |
| `storage_static_routes.storage_routes.attrs` | [storage_static_routes.storage_routes.attrs](resources--fleet--reference--group-004.md#canonical-2033130301313321-1232220332313121-3311112212031300-2211323031330223-1200021003201131-3310133300223310-2120201103231320-3022031112310003) |
| `storage_static_routes.storage_routes.labels` | [storage_static_routes.storage_routes.labels](resources--fleet--reference--group-004.md#canonical-3200200330113233-0203010303222300-1203213313200323-0002201003200323-2102302130320131-3200001233330020-2303323033120202-2032013003021232) |
| `storage_static_routes.storage_routes.nexthop` | [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-1133203310222332-0201030230322312-3010300302332211-3030020321111100-1000213331231233-3213113130130232-2213202002102121-1300021112121130) |
| `storage_static_routes.storage_routes.nexthop.interface` | [storage_static_routes.storage_routes.nexthop.interface](resources--fleet--reference--group-004.md#canonical-2101122030320222-1011022121032223-2221103102123222-0320321300122312-0211111133123331-0202123303102012-3133231312131131-0210333030032200) |
| `storage_static_routes.storage_routes.nexthop.interface.kind` | [storage_static_routes.storage_routes.nexthop.interface.kind](resources--fleet--reference--group-004.md#canonical-1113331231132223-2330021030122031-0131101120123303-1230000221030001-0001103122332133-0020303311021121-3020111331322212-0012121130302303) |
| `storage_static_routes.storage_routes.nexthop.interface.name` | [storage_static_routes.storage_routes.nexthop.interface.name](resources--fleet--reference--group-004.md#canonical-1333312131013211-0313203020122000-3301102320123301-0132032201020333-0223211322330100-3002023032030030-2200123211033213-1321023132213223) |
| `storage_static_routes.storage_routes.nexthop.interface.namespace` | [storage_static_routes.storage_routes.nexthop.interface.namespace](resources--fleet--reference--group-004.md#canonical-1120121103003330-1121213013223210-1320032132100022-0103131021332331-0022022002211021-2222033102200122-2030120310302221-1322112023232323) |
| `storage_static_routes.storage_routes.nexthop.interface.tenant` | [storage_static_routes.storage_routes.nexthop.interface.tenant](resources--fleet--reference--group-004.md#canonical-1010221012231303-3111210300330220-3230120302321021-3331011112320213-0331201113330212-2113310212112301-1213323211103302-1310121302020020) |
| `storage_static_routes.storage_routes.nexthop.interface.uid` | [storage_static_routes.storage_routes.nexthop.interface.uid](resources--fleet--reference--group-004.md#canonical-1013311001101102-1220313102102201-1111300333221302-0210130310131123-2231201133011131-1102101232131003-0221211032032212-0020023000312301) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address` | [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-1020220223202220-0330131020121100-0332133222131201-2113330101031223-0200301212310132-1130023310213131-1112002021300203-1100001320223111) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-3312122000233102-1302032131331000-1132303333011232-2330202011012201-3313213233332222-2120210220032131-0012030130320302-0013133123100333) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](resources--fleet--reference--group-004.md#canonical-3220103102133331-0223020201122211-0231322223121112-2232333303201211-2123010331323103-1110323203132033-1111101031313110-1112020223121121) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--fleet--reference--group-004.md#canonical-2131013222133111-0000103202032110-1312012023102122-1011211001113302-0013001121333133-3121331020312102-3003333302301111-3320001222000010) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](resources--fleet--reference--group-004.md#canonical-2311322200122000-1103223111300201-1130202332130112-0300020321023100-1202333113230012-3211223300002222-3322000313113211-1021011232331312) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--fleet--reference--group-004.md#canonical-1300111333301232-3323201023033111-3320320023102323-3203023211110120-3000103022133210-0112013322303233-2312332000112102-1123131301302111) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](resources--fleet--reference--group-004.md#canonical-2122033231100233-1101100013211121-2200223211023020-1032033202321101-0313020102023121-2231113130011021-1203332002021222-0210303221103002) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr](resources--fleet--reference--group-004.md#canonical-1012303303212231-1001212213102130-2113203301232132-3020102201203313-3212300023111111-1203123003211130-2221033102032311-0130101311311131) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](resources--fleet--reference--group-004.md#canonical-2320332312310002-2033230031020213-0121233021330300-3220300010333032-1112130000210320-1221313212201232-0320203111030323-1121012333210233) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr](resources--fleet--reference--group-004.md#canonical-0331213132010301-3133333202102211-3003021030202120-0210030330332000-2233302201031323-1200101200023132-0133322010100212-0113122030001032) |
| `storage_static_routes.storage_routes.nexthop.type` | [storage_static_routes.storage_routes.nexthop.type](resources--fleet--reference--group-004.md#canonical-3312013220222102-2000130020023100-1220022211110211-0000003103133000-3321032200002000-2122321300210323-2112021023213313-0312302203312332) |
| `storage_static_routes.storage_routes.subnets` | [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-2110130201003120-3313131233010200-0331032132113333-0013023022000333-3132030321222020-2121121003312322-3023020031111031-1032000302312012) |
| `storage_static_routes.storage_routes.subnets.ipv4` | [storage_static_routes.storage_routes.subnets.ipv4](resources--fleet--reference--group-004.md#canonical-3011302323032013-1312301320133210-0232233100310021-1111110203110220-2322203233013123-2201310020221320-2013130031121130-2233201123003333) |
| `storage_static_routes.storage_routes.subnets.ipv4.plen` | [storage_static_routes.storage_routes.subnets.ipv4.plen](resources--fleet--reference--group-004.md#canonical-1030302012123132-0113032112203100-2220012323302113-1211213122023013-0332202332313301-2133003021302302-0100013123323313-1001013011013011) |
| `storage_static_routes.storage_routes.subnets.ipv4.prefix` | [storage_static_routes.storage_routes.subnets.ipv4.prefix](resources--fleet--reference--group-004.md#canonical-3213201223213112-3221103030011301-2012022031303001-0211302013212103-1112332000102312-0331100323100221-2322012330011223-1110103113023130) |
| `storage_static_routes.storage_routes.subnets.ipv6` | [storage_static_routes.storage_routes.subnets.ipv6](resources--fleet--reference--group-004.md#canonical-2222033102211111-0100010122212323-2331000212000221-2011103211211320-1021123033102030-0030020002003223-1100203030003323-2201001330030331) |
| `storage_static_routes.storage_routes.subnets.ipv6.plen` | [storage_static_routes.storage_routes.subnets.ipv6.plen](resources--fleet--reference--group-004.md#canonical-1202321020130003-1012231131223213-3032033000222222-1022102012331101-1012203103330302-0213013323331032-1331220011311100-3111203113332221) |
| `storage_static_routes.storage_routes.subnets.ipv6.prefix` | [storage_static_routes.storage_routes.subnets.ipv6.prefix](resources--fleet--reference--group-004.md#canonical-3103212331223031-3200111123011211-3213213100303030-2002212031131330-3303220100022330-3310310303122132-1331121331323033-1111030113031013) |
| `timeouts` | [timeouts](resources--fleet--reference--group-004.md#canonical-0101333303030102-2023323213103220-3233013322301023-3332321203031320-3233333023230103-0230313320020131-0332230220222213-3301221323222031) |
| `timeouts.create` | [timeouts.create](resources--fleet--reference--group-004.md#canonical-3322312311221323-0230233310031231-1031033013211122-0023211200003131-0132101030112131-2202022233032230-3333022320022223-0002011110130122) |
| `timeouts.delete` | [timeouts.delete](resources--fleet--reference--group-004.md#canonical-3203133000103123-1222100322131033-1000111213021132-2121011133302120-2032112013132113-1020310210302103-2320112001100022-2323011312311032) |
| `timeouts.read` | [timeouts.read](resources--fleet--reference--group-004.md#canonical-1020230301203111-0003103030020313-1023322320302003-2003303113021220-0301112330213002-1303012123012113-0203331001212110-1111302002202302) |
| `timeouts.update` | [timeouts.update](resources--fleet--reference--group-004.md#canonical-2211201022321033-2330012313032120-0302200302012020-0133102300233231-1113112031323030-2221111031032122-3101323100013101-0200330001112030) |
| `usb_policy` | [usb_policy](resources--fleet--reference--group-004.md#canonical-1122110212330113-3101131220333203-2013203030011212-2310131001313330-1321230331331320-0133123320203112-1102233320013230-2122202121023032) |
| `usb_policy.name` | [usb_policy.name](resources--fleet--reference--group-004.md#canonical-0210031313123110-1013201323233220-3321022023222200-2131332103232323-2203032111130102-3313032333332030-0000332002311233-0320333133001220) |
| `usb_policy.namespace` | [usb_policy.namespace](resources--fleet--reference--group-004.md#canonical-2000203111211212-1122201033323001-3310033331330103-3320330022113132-0300023133131111-0131201313122323-2113011103200123-1002331020122212) |
| `usb_policy.tenant` | [usb_policy.tenant](resources--fleet--reference--group-004.md#canonical-1322013111303300-3212302013200320-3303113031311132-2220211312101300-3130101033211130-2010232231023103-3230111312312322-2223111210320202) |
| `volterra_software_version` | [volterra_software_version](resources--fleet--reference--group-001.md#canonical-1112311233323303-0300320300110001-2131011211002222-2113300313322103-1132133203321220-3302313320200210-0333231020001200-1202231201133302) |

<a id="canonical-0233020232022221-0012000111130132-3331103020110010-1132313232302301-3302110232202113-3033221010233312-1302312102121301-2003110120333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_usb` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- allow_all_usb

<a id="canonical-0011122200232321-0331210222021203-3211310010331331-0313201100121033-0210102200112232-2220223223100313-1001011001120021-3220221212103101"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_usb, deny\_all\_usb, usb\_policy\] Configuration parameter for allow all usb.

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

- [allow_all_usb](resources--fleet--reference--group-001.md#canonical-0011122200232321-0331210222021203-3211310010331331-0313201100121033-0210102200112232-2220223223100313-1001011001120021-3220221212103101)
- [deny_all_usb](resources--fleet--reference--group-002.md#canonical-3031222031130102-0220323323303102-0001201301130320-0310333220032103-2020133202021100-3213120313202223-2013331220210232-1002130131230330)
- [usb_policy](resources--fleet--reference--group-004.md#canonical-1122110212330113-3101131220333203-2013203030011212-2310131001313330-1321230331331320-0133123320203112-1102233320013230-2122202121023032)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_usb = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- blocked_services

<a id="canonical-2300213201220031-0101132121033233-3121011212322211-3303123131131133-1203331103123303-2331223201001111-2121021312131201-1321123122123132"></a>

Type: `"object"`. list nested block, Optional.

Disable node local services on this site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 6,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 6,
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
    "ves.io.schema.rules.repeated.max_items": "6"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "6"
  }
}
```

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333321131221221-0310021300230331-3112310320120221-2330211023000002-3003023110000101-2110203311102133-1130120211112211-1330300131033213"></a>

### Direct properties for `blocked_services`

- [DNS](resources--fleet--reference--group-001.md#canonical-3221332113111220-3010302110013032-3213220312323020-1000113322202101-2022231333313222-1303323113231210-0322133011123230-2012313233132100): complete subsection reference.

<a id="canonical-2213010100023002-1333330311200302-3200201323202200-0210010100210332-1001103132033320-0331312232110230-3113122031033311-0230031130001232"></a>

<a id="canonical-0023133133122310-2132003121112301-0132113102303011-1023223330131322-2313100110232320-3310210111201123-1200121211312132-1023133201133121"></a>

#### `blocked_services.network_type` property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](resources--fleet--reference--group-001.md#canonical-1312203301110103-3130012113230101-2222302103001111-1223002010202130-2211020022130220-0002131113022003-3323311022010330-1232222133330110): complete subsection reference.

- [web_user_interface](resources--fleet--reference--group-001.md#canonical-2330030223021022-3303111302020012-1322230100321132-3313012120100100-2301301031112131-2333010110020231-2330221010013220-1231001021220120): complete subsection reference.

<a id="canonical-3221332113111220-3010302110013032-3213220312323020-1000113322202101-2022231333313222-1303323113231210-0322133011123230-2012313233132100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.dns` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.DNS

<a id="canonical-3001003313321332-3330212121120221-2231001312302113-0130210130101030-1223311221031011-1202322303232120-3103302032023322-0021233022113320"></a>

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
dns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312203301110103-3130012113230101-2222302103001111-1223002010202130-2211020022130220-0002131113022003-3323311022010330-1232222133330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.ssh` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.SSH

<a id="canonical-0302223331101002-0213123231100103-3000030300221332-3300313302033202-0212203220321130-1330103300322313-0321101230120100-0002110202021030"></a>

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
ssh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330030223021022-3303111302020012-1322230100321132-3313012120100100-2301301031112131-2333010110020231-2330221010013220-1231001021220120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.web_user_interface` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.web_user_interface

<a id="canonical-1123031312320312-1220130012132210-0133032010322123-2101323021222330-2023233200223023-2311122321002303-0030311213221031-2003303232021033"></a>

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
web_user_interface = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- bond_device_list

<a id="canonical-3112331033303132-1010323223233021-0001113321020203-0020000321010310-0312210031322111-2212121113112120-0012220002121030-2332000033322102"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
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

- [bond_device_list](resources--fleet--reference--group-001.md#canonical-3112331033303132-1010323223233021-0001113321020203-0020000321010310-0312210031322111-2212121113112120-0012220002121030-2332000033322102)
- [no_bond_devices](resources--fleet--reference--group-002.md#canonical-2233202132112121-1300211023330221-2212133210020112-2021322122322312-0321223001202132-3220130012222211-3210110102122303-1332222011023230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301211031130203-1030321323220031-3210332102301103-2011022210020323-0103122021323120-0312303033130200-0000313300000233-2120112101201122"></a>

### Direct properties for `bond_device_list`

- [bond_devices](resources--fleet--reference--group-001.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113): complete subsection reference.

<a id="canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-001.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- bond_device_list.bond_devices

<a id="canonical-3233230000001102-0311301000203233-2111011133110120-3233103113333020-3021302233123110-2002020302302210-3331023230132110-2230111323213220"></a>

Type: `"object"`. list nested block, Optional.

Bond Devices. List of bond devices.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingListObjectAttributes("active_backup",
    "lacp")}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
bond_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012131200112130-2220133030113232-0310222113211313-0113312033030201-2130120020300030-3313003312002330-2333131232200130-2223210123211223"></a>

### Direct properties for `bond_device_list.bond_devices`

- [active_backup](resources--fleet--reference--group-002.md#canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002): complete subsection reference.

<a id="canonical-1102101300033333-1222302001120331-3001102112133100-1113221021010223-1011033102031323-2001201332011123-3021100212311201-3021223202312132"></a>

<a id="canonical-0000212010022101-2231302230031331-3223213113221011-2323113103320010-1302010130333332-0210010231203312-2203120213211022-1032333223101212"></a>

#### `bond_device_list.bond_devices.devices` property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--fleet--reference--group-002.md#canonical-1232131310312013-1102232002320220-3200203002001023-0003223112323331-1230230223333130-2100022101102310-1212322333003213-1131321102112223): complete subsection reference.

<a id="canonical-3220201301033322-3120000332322333-0213203203003312-3232200020010033-2303012020310131-2302311233300131-0000013010231312-3013130100312323"></a>

<a id="canonical-3013220013022313-0231332320202000-3130210133221031-2312131013122312-0013103323223023-2120220123323132-2300010011023033-2301210010131221"></a>

#### `bond_device_list.bond_devices.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2112131103330123-0303001101002111-1102112103200311-1113220212110022-2303210120033112-0031200132110330-1222020030102232-3102030331313323"></a>

<a id="canonical-3102112213200122-2211130202020203-2320001112102110-2121222011331313-0303232302020300-2320213001211321-2020220110202001-1120111023303303"></a>

#### `bond_device_list.bond_devices.link_up_delay` property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-2133233103031220-1330321001113020-1103312233020121-0100122200001233-1103032120012000-3130012302010322-3313003331032211-2100312203130113"></a>
