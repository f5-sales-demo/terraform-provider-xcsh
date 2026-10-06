---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- Property reference

<a id="canonical-3220320232020133-2333220110100310-3132000323133212-1002032103213123-3331103203331312-1022300103303112-0311123101001223-2210302300100110"></a>

### Direct properties for `xcsh_fleet`

- [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-1020303212333102-1221132312100331-1230301230330311-0322102012012202-2302112223332023-1131133311222203-0200222121232321-1312210000123110): complete subsection reference.

<a id="canonical-3033011200000022-2001102100300220-1313003302330003-1220312010120321-3300312011330111-3323013312033210-3321003302302212-0213113212102211"></a>

<a id="canonical-1333100103000202-3000003131001010-3332210201321323-2101220333102022-2211100123203010-3023223103122322-2210331020020202-2123013130001312"></a>

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

- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-2003013010222032-0231112013202331-2201001123333132-3030101330330213-0200210113111232-3323123102012112-0310000103030013-1100222131113031): complete subsection reference.

- [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-1211330203133133-0100030312311131-2320023111103033-3211001101101231-2100331200213120-3133220022001023-1132231031111302-0113330202110213): complete subsection reference.

- [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-2112312321330302-1211021232313121-2331211202132011-1132132332222033-2303221021220203-0002022311322203-1311030023321330-2133303232003230): complete subsection reference.

- [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-1011213321333020-0122213232122233-3330203123313132-1102003220000100-0310311103303310-1021233230221210-2033031202202022-0212001013131023): complete subsection reference.

- [default_config](data-sources--fleet--reference--group-002.md#canonical-1232232122223230-0300310232023033-2100332200122130-0211200101321332-0000311133222122-0230121123110011-0201012012212022-1221302112223212): complete subsection reference.

- [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-1033001023033030-3300110122022001-2102101121002010-0102011111321302-3323110310023332-2112010011212302-1110320322001132-3102320123230202): complete subsection reference.

- [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-2130302203111131-3121312201232200-1002223322032001-3203202102210303-3010121002122122-3101131223332203-1300222311301312-0311213113232021): complete subsection reference.

- [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-0022132120103110-3011033030320022-2121100111221111-1030322032011211-2033202230113022-0112303030121313-2001001200312221-3320303123030211): complete subsection reference.

<a id="canonical-3000120212110210-2320001003331113-3332332310111033-1221111011001013-3022303211020311-1302330031213320-0213022302321033-1223221223122112"></a>

<a id="canonical-1313212131022131-0201321003100021-3101030000131312-2032022020310302-2203230333202030-0131302012011223-3113103020132022-1322033323311331"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Fleet.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [device_list](data-sources--fleet--reference--group-002.md#canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323): complete subsection reference.

- [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-3130311131301002-1132033103331133-1202332201233213-1132213322011332-1011320123020311-1201320333323002-1310223001230020-3202103331012201): complete subsection reference.

- [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-3010223011012221-3311031211012022-3113130032121323-1100233121200201-0022032322131333-2221122020212113-3102311212332322-0230212232301232): complete subsection reference.

- [disable_vm](data-sources--fleet--reference--group-002.md#canonical-2330100011232133-1230010220110222-3001003032121133-1003303230211323-1100022022303001-2010203310213123-3103022122231203-2212333211200002): complete subsection reference.

<a id="canonical-2023212313210220-3321102222330012-1021213102122221-3320310033131011-3022020113321201-0033210322333201-3203333212231333-2113211030302023"></a>

<a id="canonical-2112233223032033-1322330201321201-0233132212122232-1130211201110102-0113301323000130-2100102221130131-3332111001222120-1030013103331322"></a>

#### `enable_default_fleet_config_download` property

Type: `"bool"`. Computed.

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

- [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-0310032200331023-2331221302202203-1333300101031103-0203100200300220-1121123020003110-1003110222120021-1121322210111332-3303112011011320): complete subsection reference.

- [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-3203031033330322-1321112310121322-1323220113000022-0203131210113200-3313302013102112-3331202022212031-1303313212001010-2303130111320221): complete subsection reference.

- [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-1031101321301121-1232330032312323-3013110201310103-1201113000102302-0300000320123233-2021130001313132-3302020232203232-0020323013122323): complete subsection reference.

- [enable_vm](data-sources--fleet--reference--group-002.md#canonical-1113110323023002-3202020020033032-2102002213011000-1322112033032223-2201001310100103-1213132303120023-2331011300313031-1330120012301000): complete subsection reference.

<a id="canonical-1100313101112122-3220013132031222-2332230113102201-3021112100021300-0220010003201113-3103120000323023-3022200333223102-1232222203300312"></a>

<a id="canonical-2222110320122303-0110003300222123-3310330312313220-3202020032233101-3213012212222320-0113101130120122-1232201001011121-0201222113323122"></a>

#### `fleet_label` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2102011310101023-0211302120131210-1003222320032033-2100211331201133-3100322121103022-0003000302303120-3133123330012101-2021103123303221"></a>

<a id="canonical-3003123210303011-1002012233230011-1310102201201311-2230032321013132-2030221200130002-0131321333002202-3233201212133113-3102123202121113"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [inside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-3110033202301221-0230011123210012-2230320123132312-3130002012323300-0300203022111202-3130003001220313-3021313002211102-2012002112213210): complete subsection reference.

- [interface_list](data-sources--fleet--reference--group-002.md#canonical-2331132030231322-2323132333102123-2223312302310133-2010010233221210-0131002013222232-2213320220213201-2023231030021312-1122032020022121): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202): complete subsection reference.

<a id="canonical-3133011301321311-0331102101200233-1300202110331120-1132232001312003-0300300033212231-2101130122320222-3202021023220333-1300022121231133"></a>

<a id="canonical-2020102211210002-0011003001120103-3331303131323103-1210031132100132-1201210031331111-2111322022132303-2022322203221103-2200321303000130"></a>

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

- [log_receiver](data-sources--fleet--reference--group-002.md#canonical-1201301010102220-2100130021022310-3332321023231220-2333030020220310-0033013130102022-0331111133203220-2213203023002022-0313201111201132): complete subsection reference.

- [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-2021310223321313-2033131211212111-0113103022232333-2010100020021011-3113131010131303-2321233323330001-0200120122003231-3232132212231033): complete subsection reference.

<a id="canonical-1312303232112030-1301211030020013-2221020211000210-0030322002231021-0233202013322133-1033222020132023-0133202301102231-0323201011132033"></a>

<a id="canonical-2033323310121201-3221310132203021-2001212123123123-1103330331002020-2102231113332203-0200221102212100-0200322000132322-0221202213300031"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Fleet.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3202033333302011-3312300000333232-2002203210021103-0021230330330300-3310013230230232-0110120312200232-1322110223313000-0030211220330113"></a>

<a id="canonical-3100322111323023-1011002202233200-2331132323020122-2301333110032332-1000112033330232-1332013213303330-0022111202112220-3201220011110030"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Fleet exists.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [network_connectors](data-sources--fleet--reference--group-002.md#canonical-3123211022312221-3200100111011110-1201100232001010-1230100331220302-0013301221331213-2220233203231132-0121230201000122-1132010113102123): complete subsection reference.

- [network_firewall](data-sources--fleet--reference--group-002.md#canonical-3321202220211332-1233100100233320-2030101201021220-1310201200230102-1011003120311131-2130232230230210-1022110013030131-2000202132233333): complete subsection reference.

- [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-3231300102332313-2132132030313103-2122021022212031-2022211230213122-2030000022032311-3110302202102103-0002101133113121-0021030121221122): complete subsection reference.

- [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-0221310033011230-1200213332232320-3131023300031210-3031133012130331-3132110111231213-0210320332210001-3122322022310032-1303302321010011): complete subsection reference.

- [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-1232113232310321-1222130032123323-3333333031011322-2101300132000023-0210131112003320-0023322310102202-2231301313303232-3221202133012133): complete subsection reference.

- [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-1301230210330122-2303332032010312-0200323323023201-0033133101310013-0102213033032231-1113332131230030-1221133003003300-1310300203130100): complete subsection reference.

- [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-3330102102110003-0233200221011033-0302123300312033-1211200102221301-1102322220013301-0112121232210100-2232012022332013-3230220321322303): complete subsection reference.

<a id="canonical-0032133311300220-3112212221113122-2001121122021010-1302311010033000-2300130131013210-2223033123112333-2302222230222220-1113012310100002"></a>

<a id="canonical-3030300123322210-3300021310133122-1331311212230021-2023301000030321-1312133321033030-3000222221012022-3020010103213313-1031101022001301"></a>

#### `operating_system_version` property

Type: `"string"`. Computed.

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [outside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-3322003033013013-1301210121021213-1130031120131113-3200101322001310-3222130102023200-3232303313121331-2133321132323111-0011301020200212): complete subsection reference.

- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211): complete subsection reference.

- [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-1132110120003111-0323103313100123-2113310212200301-0133232320111330-0030220003213233-1113311012021102-1212011223322232-3112201313103010): complete subsection reference.

- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121): complete subsection reference.

- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021): complete subsection reference.

- [storage_interface_list](data-sources--fleet--reference--group-003.md#canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200): complete subsection reference.

- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300): complete subsection reference.

- [usb_policy](data-sources--fleet--reference--group-004.md#canonical-2132300120030302-2100113132002202-2230322111111011-2301031030022032-0303102003030130-2111000030223010-1011003201131323-2230112300300113): complete subsection reference.

<a id="canonical-1010001010110200-3120320132311103-3010311202213130-1300100213231221-2003030232001232-1301003332011103-2310033011213010-3203301123121122"></a>

<a id="canonical-1020303130233110-3120031131111033-0003013323132020-2110000000333312-0131033321131202-3011322113201333-3120030211232231-2033130011132212"></a>

#### `volterra_software_version` property

Type: `"string"`. Computed.

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0103013012003112-0131231031132133-0213223021112133-0233331011003230-2133100333110230-1301302332010333-1013111131321211-1310123122022321"></a>

### All schema paths for `xcsh_fleet`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_usb` | [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-2021101313013013-1033323101021333-0310311232320022-2131030212320211-3310303321320310-0100213023130322-0110302012202312-2123232131300330) |
| `annotations` | [annotations](data-sources--fleet--reference--group-001.md#canonical-3033011200000022-2001102100300220-1313003302330003-1220312010120321-3300312011330111-3323013312033210-3321003302302212-0213113212102211) |
| `blocked_services` | [blocked_services](data-sources--fleet--reference--group-001.md#canonical-2130211202002331-3222231301111321-3133110030013133-0102103110210231-1002002013322330-0303000031102103-0220121201202323-2312202030201211) |
| `blocked_services.dns` | [blocked_services.dns](data-sources--fleet--reference--group-001.md#canonical-0311023200202021-2122303303322122-0223032021231303-0020022102323232-1112332312121100-2021100321321112-1133133301200333-3002221113222132) |
| `blocked_services.network_type` | [blocked_services.network_type](data-sources--fleet--reference--group-001.md#canonical-2120311323121000-3313300213022212-3312003033123302-2013013112300301-3110123112001311-0332223323213331-0231210010330001-1003311323232310) |
| `blocked_services.ssh` | [blocked_services.ssh](data-sources--fleet--reference--group-001.md#canonical-2301010330320113-0102212222001110-2312002123133232-3122332313313131-3210030321302213-0312311131332111-3202301203131113-1220301323303330) |
| `blocked_services.web_user_interface` | [blocked_services.web_user_interface](data-sources--fleet--reference--group-001.md#canonical-1232300013003002-1203201233323332-0223021212323012-0232201003203202-1312011032312122-2322032020100302-2031131023033033-1202000200303002) |
| `bond_device_list` | [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-2111330100223331-3300130221220232-3123203101223113-2231113033200102-3323120223131223-0110320213233022-3223001020113031-1112013312012301) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--fleet--reference--group-001.md#canonical-1032210001011102-0030333202303211-0030310021310013-3133303332013012-3312332220230231-0231331003130131-3233122123203020-0230223201022102) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--fleet--reference--group-001.md#canonical-0303012002032201-3333211033102202-0121120211203330-0122310221211203-3000020210202031-0000220011122030-3232222101232223-3102312011032300) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--fleet--reference--group-001.md#canonical-0003213311031122-3211112011321113-0002303130011102-0231023321321111-2320012210012101-0312003001003133-0013233010021230-0313000200323222) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--fleet--reference--group-001.md#canonical-3131010200302002-2330001123210203-3113213033233302-2212323333120111-2211203332021032-2333123010022321-2201312212330012-2232122213102031) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--fleet--reference--group-002.md#canonical-1323023103302121-0032331200203312-2332003031022310-0330120232202200-3201112223202103-1312211203131312-3011130130010020-1000302021202302) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--fleet--reference--group-001.md#canonical-1301001010330210-1102301301130332-1232120220212133-0012301121323312-3213020120313011-1122012313111203-3332201121200220-0122211102003131) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--fleet--reference--group-001.md#canonical-1122332121001121-3310101112213322-2012031300010211-1110201221221330-3031023031010211-1023130120132311-2000030200031203-0111223133121010) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--fleet--reference--group-001.md#canonical-2231223333021020-0232311320013010-0223231323313232-0231213112122023-1003232130223012-1213112101131030-3210202300032200-1230112002022131) |
| `dc_cluster_group` | [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-1323102031301102-2302023121333302-3003033313301001-3200222232333112-3221033022320300-2213112303112112-3111100331223113-3022013100211010) |
| `dc_cluster_group.name` | [dc_cluster_group.name](data-sources--fleet--reference--group-002.md#canonical-3101200231322220-0030110100012122-1122030302023033-2100213222203322-3013110111233233-0030013331212201-2013220320021112-2211001233110202) |
| `dc_cluster_group.namespace` | [dc_cluster_group.namespace](data-sources--fleet--reference--group-002.md#canonical-1111002223033100-3010002312233113-1101011133233312-3311322232030220-1230333231031011-3103100222030202-3221230220300321-0133122311311121) |
| `dc_cluster_group.tenant` | [dc_cluster_group.tenant](data-sources--fleet--reference--group-002.md#canonical-3200332321310101-1113002021233023-3220132123323222-2110110130020113-3221211221011221-1011132331321031-1212013120021113-3033333033223311) |
| `dc_cluster_group_inside` | [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-0320130313200120-1222323131231112-2121322303301022-1322103111131230-1212313323223121-0132201113112130-0230132121302013-3023123020233011) |
| `dc_cluster_group_inside.name` | [dc_cluster_group_inside.name](data-sources--fleet--reference--group-002.md#canonical-3220102313322131-2231121010232020-2131232311230230-1221120000212310-3001322301212300-2100222002022212-3022230203131020-3021120101133023) |
| `dc_cluster_group_inside.namespace` | [dc_cluster_group_inside.namespace](data-sources--fleet--reference--group-002.md#canonical-0110332101221300-3103210103011130-0202021120021002-2122213331232223-0001312201021210-2001032102103311-0222231100133011-3032320301103030) |
| `dc_cluster_group_inside.tenant` | [dc_cluster_group_inside.tenant](data-sources--fleet--reference--group-002.md#canonical-0131233301311000-1023230333031322-2020212213332313-2131013121300000-1312222200223120-2300213320213113-0310300100232020-0200102313012301) |
| `default_config` | [default_config](data-sources--fleet--reference--group-002.md#canonical-3123101121002222-0023112111233223-2020310133003302-2100022113000013-2020323002023330-0112103301320002-1313303103032320-1010021113303113) |
| `default_sriov_interface` | [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-1331130211111001-0122132322111332-0020302032000110-2211333211111331-3220132310201330-0330113111110013-2203120013110032-0101011122133233) |
| `default_storage_class` | [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-1230313022311013-2033130033232303-3113323201222203-3110331212320020-2102211123032221-1303110110332232-2222211212001213-0322312320120321) |
| `deny_all_usb` | [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-0323201230021333-3213210223102010-1222323310230210-1003010300303302-3103023120303323-3331021111202320-1202222300110102-3301322010230232) |
| `description` | [description](data-sources--fleet--reference--group-001.md#canonical-3000120212110210-2320001003331113-3332332310111033-1221111011001013-3022303211020311-1302330031213320-0213022302321033-1223221223122112) |
| `device_list` | [device_list](data-sources--fleet--reference--group-002.md#canonical-2011300210203112-1212221211332200-3213022223032230-0000021313221110-2102020102331131-0130001302313130-3211210123310210-1201013002123013) |
| `device_list.devices` | [device_list.devices](data-sources--fleet--reference--group-002.md#canonical-2202032110211321-2300103112313321-1023023022323021-0301213113021332-0022001330301103-2100130231330300-3312322202100211-0020321311112312) |
| `device_list.devices.name` | [device_list.devices.name](data-sources--fleet--reference--group-002.md#canonical-3200123031113200-1001001132222210-0211213012310112-0212012312130010-2311023332120120-0303103022313223-0231121021232122-1331313332020010) |
| `device_list.devices.network_device` | [device_list.devices.network_device](data-sources--fleet--reference--group-002.md#canonical-0320300202003032-0302232030302230-3203121322110123-3122030332100130-1222123031130121-0013121130321213-2102120313103002-1102203310210112) |
| `device_list.devices.network_device.interface` | [device_list.devices.network_device.interface](data-sources--fleet--reference--group-002.md#canonical-2112111333200221-1223010122303313-3013000020032101-2023030310011002-0231012111120132-0100001010232313-1130022103330030-0002203321303121) |
| `device_list.devices.network_device.interface.kind` | [device_list.devices.network_device.interface.kind](data-sources--fleet--reference--group-002.md#canonical-1303321303201323-2333033320230121-1211220330021113-3023101221212013-0201022033202120-2011033333212200-3121220230132011-2032132132030230) |
| `device_list.devices.network_device.interface.name` | [device_list.devices.network_device.interface.name](data-sources--fleet--reference--group-002.md#canonical-2303302201213301-3121130133112232-1322311333103233-2011303323001203-2112201313102110-1302020301001021-0220021101110131-2123221223023121) |
| `device_list.devices.network_device.interface.namespace` | [device_list.devices.network_device.interface.namespace](data-sources--fleet--reference--group-002.md#canonical-1000202303113302-3223233313101003-2303320020211113-0310322322000130-3332222310331303-2220320200300112-2202003102213332-0312300303001003) |
| `device_list.devices.network_device.interface.tenant` | [device_list.devices.network_device.interface.tenant](data-sources--fleet--reference--group-002.md#canonical-0310133222223303-2113312030313113-0333011003031031-1331331131312313-0332112123023002-0333122231210112-2121331002310322-1300112301202001) |
| `device_list.devices.network_device.interface.uid` | [device_list.devices.network_device.interface.uid](data-sources--fleet--reference--group-002.md#canonical-3130200202230302-2310030222231212-3010331322120120-1330111022303231-0302300012113112-0032220230211112-1223130000330032-1033001311111200) |
| `device_list.devices.network_device.use` | [device_list.devices.network_device.use](data-sources--fleet--reference--group-002.md#canonical-0030022301003213-2210012012300032-0012010111313200-3113330333311301-3321230222312121-2121122130132112-1011330210133233-3120100233121313) |
| `device_list.devices.owner` | [device_list.devices.owner](data-sources--fleet--reference--group-002.md#canonical-1323233211030332-3223103001111232-0110100113331220-2213013202023302-1032330232322113-1123332010132313-1312333202021213-1213021001330020) |
| `disable_gpu` | [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-0321031311123031-0000123131131113-0112033002330000-3030303003030323-0131123113120032-0201023322312023-2320333330101233-1031103000101001) |
| `disable_log_anonymization` | [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-2101033032331320-3200303222032210-3030311223222310-2131323332233010-0212331102200023-1133032132322221-0222012132120323-3213021210122012) |
| `disable_vm` | [disable_vm](data-sources--fleet--reference--group-002.md#canonical-2012021110001110-0133232210322213-1323223201321211-1113310113003003-1110012031013001-0213012103310230-3003010202013132-0220222112202110) |
| `enable_default_fleet_config_download` | [enable_default_fleet_config_download](data-sources--fleet--reference--group-001.md#canonical-2023212313210220-3321102222330012-1021213102122221-3320310033131011-3022020113321201-0033210322333201-3203333212231333-2113211030302023) |
| `enable_gpu` | [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-1123320121201021-1123212000311132-2002023122101311-1000030320110023-2002330003123200-3000013212011110-2032112333331201-0333011323201321) |
| `enable_log_anonymization` | [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-1321321032230021-2202033011023221-0301122312020021-0212301302110210-0312012333032201-3112020103331300-3103200131203232-3322213303200103) |
| `enable_vgpu` | [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-2300303031121201-1301011321012331-1113323000210020-2311010232021132-3020133331232123-0001321233221331-1023303300013233-3111322220301110) |
| `enable_vgpu.feature_type` | [enable_vgpu.feature_type](data-sources--fleet--reference--group-002.md#canonical-0231112300102121-0311233222122022-1223300130111333-2132213131203203-3203201222032132-2113120013013333-2313332110100213-1112331322321231) |
| `enable_vgpu.server_address` | [enable_vgpu.server_address](data-sources--fleet--reference--group-002.md#canonical-3002211211002022-1033300022020022-0313030130120011-2201100332233330-0333330100130222-2120201113002313-1031321231322212-3102213301223111) |
| `enable_vgpu.server_port` | [enable_vgpu.server_port](data-sources--fleet--reference--group-002.md#canonical-0210323130301000-1233222212021120-1300101102030012-0020221311020213-2203332223013120-3021300121330112-3113210022102001-0030021331003331) |
| `enable_vm` | [enable_vm](data-sources--fleet--reference--group-002.md#canonical-2101023113213303-3312233120303332-3230223210222010-2023322002211301-1021123102120221-2121123331321201-0202011100031023-1123001012210033) |
| `fleet_label` | [fleet_label](data-sources--fleet--reference--group-001.md#canonical-1100313101112122-3220013132031222-2332230113102201-3021112100021300-0220010003201113-3103120000323023-3022200333223102-1232222203300312) |
| `id` | [ID](data-sources--fleet--reference--group-001.md#canonical-2102011310101023-0211302120131210-1003222320032033-2100211331201133-3100322121103022-0003000302303120-3133123330012101-2021103123303221) |
| `inside_virtual_network` | [inside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-0301321331200333-2232120232101122-0033233022103113-1330120121333023-2132303111331303-1320001332021323-3102030220223203-1333303103001223) |
| `inside_virtual_network.kind` | [inside_virtual_network.kind](data-sources--fleet--reference--group-002.md#canonical-3210320300100322-2120300210100322-2300331211122332-3101300112302323-0010220022133000-1312022322002231-1121000330312312-3223102003221103) |
| `inside_virtual_network.name` | [inside_virtual_network.name](data-sources--fleet--reference--group-002.md#canonical-0100211013233113-1120313312223202-3003331131100333-3102231202123001-2033000112332121-2113202002200010-1031200211020323-3230033321202001) |
| `inside_virtual_network.namespace` | [inside_virtual_network.namespace](data-sources--fleet--reference--group-002.md#canonical-2310313213303020-0111032110001131-3111232323102123-2212301130130103-1201313223303113-2100133003131133-3231223312210232-1012112313033210) |
| `inside_virtual_network.tenant` | [inside_virtual_network.tenant](data-sources--fleet--reference--group-002.md#canonical-1002031031213131-1002131221033011-1100313133320323-3231312303203300-0202312120113201-1213101011103300-2001000220303121-0300023133123020) |
| `inside_virtual_network.uid` | [inside_virtual_network.uid](data-sources--fleet--reference--group-002.md#canonical-3213233233302220-3131003223333020-2300100231100332-0200312212021030-3122112331032011-1010233302123310-3331331220301312-3223031322103202) |
| `interface_list` | [interface_list](data-sources--fleet--reference--group-002.md#canonical-1121100011030011-0112201021211330-0232123013200213-0013132120303311-0211002302121000-0323003013300122-0112213121000233-1113331010200030) |
| `interface_list.interfaces` | [interface_list.interfaces](data-sources--fleet--reference--group-002.md#canonical-0233013330202230-2103203221230201-0120300120203130-0222113132330103-1123010103030010-2032032002331020-1212112301010123-0023101023102203) |
| `interface_list.interfaces.name` | [interface_list.interfaces.name](data-sources--fleet--reference--group-002.md#canonical-2000303113030331-3223221011030333-2230011013131211-2032121023202333-1010003303031311-0202132112031013-1302110001031021-2232312220223203) |
| `interface_list.interfaces.namespace` | [interface_list.interfaces.namespace](data-sources--fleet--reference--group-002.md#canonical-0103020003230333-0223332322130110-2131322032112230-3120221313332333-3331323013121001-2023011303210102-2021332031310021-3331210223131120) |
| `interface_list.interfaces.tenant` | [interface_list.interfaces.tenant](data-sources--fleet--reference--group-002.md#canonical-2301320300313203-0113301131323200-0002100010131113-1301100022031323-0222013311113202-2112110112313231-1101032222202331-3303312330010330) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-0033222103311202-1301311123132020-2120212211121213-0103232230223330-3021301001000230-1233213322031220-1031203003011122-3203230212301313) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-3312032011301300-0230332031201323-2320000112331313-3130302010202030-2010101333230101-2300010300130223-0110011123120031-0031100030211230) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2211203311302020-0000022320122023-0330103110112000-0100122032122002-3100111212112231-3210301321320311-0320202333021122-2210231333322332) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-2200201312002301-1033333102111213-0122100123013331-0032203211132232-1301021021323112-3022020031312123-3210020010012313-2011001012102201) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--fleet--reference--group-002.md#canonical-0312010020122010-0010003100311002-3201232121200311-3303210130131210-3132012233300011-2003202133301232-1232010120330223-1130103330032332) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--fleet--reference--group-002.md#canonical-1133010322313123-3330111012330232-3033212213102002-2002302113202110-2112032311212130-2023110033102013-0010000020110013-3322300331301311) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--fleet--reference--group-002.md#canonical-1213301301020331-3222231122211030-0122111132232101-2030103123233033-2130303021133333-2101002210220020-3022000012111023-1303321201330130) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-0200103012003232-3132222012002223-3010003302331002-1032132233120232-1302331011331100-2231132030003130-1323133202000022-1233002201110033) |
| `labels` | [labels](data-sources--fleet--reference--group-001.md#canonical-3133011301321311-0331102101200233-1300202110331120-1132232001312003-0300300033212231-2101130122320222-3202021023220333-1300022121231133) |
| `log_receiver` | [log_receiver](data-sources--fleet--reference--group-002.md#canonical-3312133232332200-3131012330233221-3302321110021211-1233220200313031-3223312202102213-1303320202103233-0322132222132031-2331020112313010) |
| `log_receiver.name` | [log_receiver.name](data-sources--fleet--reference--group-002.md#canonical-1032222211330222-3011012102002122-0220023303120323-0310133301230113-2130223032113330-1002113130312112-2132223003313000-0321313032301300) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--fleet--reference--group-002.md#canonical-0113331100013321-0013133013311333-0001120223133120-1231032230323033-1100101322202023-0132010200332332-0120121311032111-3333310110130101) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--fleet--reference--group-002.md#canonical-0133310222310121-0313010303230000-2000322302313310-0302333002232031-1300103031021021-3210100031130213-1213233220102020-0223102303113101) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-0110203111011030-1320123113232221-3030201100112202-2203312312332211-2322203331001002-3133002000011332-2313302313333121-0303133230302231) |
| `name` | [name](data-sources--fleet--reference--group-001.md#canonical-1312303232112030-1301211030020013-2221020211000210-0030322002231021-0233202013322133-1033222020132023-0133202301102231-0323201011132033) |
| `namespace` | [namespace](data-sources--fleet--reference--group-001.md#canonical-3202033333302011-3312300000333232-2002203210021103-0021230330330300-3310013230230232-0110120312200232-1322110223313000-0030211220330113) |
| `network_connectors` | [network_connectors](data-sources--fleet--reference--group-002.md#canonical-3203020222222130-2112202333112110-3013222300132220-3200332010130130-1010201232331313-1223033303322210-2133003223201003-1311200030303102) |
| `network_connectors.kind` | [network_connectors.kind](data-sources--fleet--reference--group-002.md#canonical-1000031223220020-2220033010023220-2031011032110322-2302313301122202-1131132030003301-0132220002010203-2112312300211031-2302003211202000) |
| `network_connectors.name` | [network_connectors.name](data-sources--fleet--reference--group-002.md#canonical-1010221313033123-0321133003001020-0020330131213311-0212010131101312-1202131230203023-0023131223123212-3120032320321033-0333100332113213) |
| `network_connectors.namespace` | [network_connectors.namespace](data-sources--fleet--reference--group-002.md#canonical-2100312301231223-1201211313112103-2201022131133230-3112230311032222-2332301300232200-3111333200120312-1310000130122011-2011113201113102) |
| `network_connectors.tenant` | [network_connectors.tenant](data-sources--fleet--reference--group-002.md#canonical-1110201230233201-2101302110023223-2003000102220321-3010031331010002-2022222123232112-2112323201311220-3100313321322100-2032002311210131) |
| `network_connectors.uid` | [network_connectors.uid](data-sources--fleet--reference--group-002.md#canonical-0222101321313132-0003203002010120-1231023011023113-2321320133021003-1110301212002212-3323303330000033-3113131012000311-1000302322133112) |
| `network_firewall` | [network_firewall](data-sources--fleet--reference--group-002.md#canonical-0323002023133100-3213113122023212-2210101100101232-2330032000112331-0032023223132312-0212032011021100-0302010331303301-3123210322012132) |
| `network_firewall.kind` | [network_firewall.kind](data-sources--fleet--reference--group-002.md#canonical-2132312121120231-2003032200221021-3122303302112322-0223321330030132-2131001131133330-1003103320221020-2110233311222023-1103323023012131) |
| `network_firewall.name` | [network_firewall.name](data-sources--fleet--reference--group-002.md#canonical-3101321320100333-2301303130233210-3000310230232100-2202333232210111-3000231233120233-3101322012123131-2032331123000222-1211133123203001) |
| `network_firewall.namespace` | [network_firewall.namespace](data-sources--fleet--reference--group-002.md#canonical-0303202021223113-3301222201331203-1102210023002322-0223121300133220-2233013020023231-2313210102032323-2221000312230331-1222102312133033) |
| `network_firewall.tenant` | [network_firewall.tenant](data-sources--fleet--reference--group-002.md#canonical-1212112012320012-0031021310112321-1220303010200232-0332330331212221-3321121012113121-0313322001110221-3023302322122122-1033101113220003) |
| `network_firewall.uid` | [network_firewall.uid](data-sources--fleet--reference--group-002.md#canonical-2333222113120100-1003003203333221-3100333133000002-2123022033013222-0310330333202131-1131001120210013-1133323132210230-2322100121330002) |
| `no_bond_devices` | [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-3023231313233222-2123102110130331-0131021311233301-0232210010022131-2100101001012022-1230233223123212-1231312301132111-2110320121201112) |
| `no_dc_cluster_group` | [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-1100023233033321-1031202011312123-3220000313212323-3201320320113320-0301111211231330-3010123121312200-2311202132022322-1301230111210222) |
| `no_storage_device` | [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-0212022003202000-2210221312322100-1023231001131111-2220322301211232-3220222003032201-0333231023102212-1121310020200021-1313031112120012) |
| `no_storage_interfaces` | [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-0322332222120213-3002300233031103-1211310310112200-1322030123230220-2013300132123111-3031333113133123-2303321030030013-3323013312300231) |
| `no_storage_static_routes` | [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-1111022332203132-2200022301101201-3201120013303101-1330033232011220-1301021003311330-0023231001020021-0133211032131311-3130330322013231) |
| `operating_system_version` | [operating_system_version](data-sources--fleet--reference--group-001.md#canonical-0032133311300220-3112212221113122-2001121122021010-1302311010033000-2300130131013210-2223033123112333-2302222230222220-1113012310100002) |
| `outside_virtual_network` | [outside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-1100332213023013-3301211321203301-2132230213220031-1210123100313003-1020033101303102-3213330320103023-1113011213230332-2002002330323102) |
| `outside_virtual_network.kind` | [outside_virtual_network.kind](data-sources--fleet--reference--group-002.md#canonical-1032333030013023-2023301001010302-1130122311112233-2303013200030202-3002002301213323-1131330102202212-3223130311012321-1012111120300013) |
| `outside_virtual_network.name` | [outside_virtual_network.name](data-sources--fleet--reference--group-002.md#canonical-2133201200110023-2330213000202333-0001101301233021-3213021033103112-2200012321030010-3112102101013103-0000120033030131-2130213021231222) |
| `outside_virtual_network.namespace` | [outside_virtual_network.namespace](data-sources--fleet--reference--group-002.md#canonical-3223323110132233-1101221021131211-2011033303201020-2311221133020300-0311203332300020-2101301211031222-3301302323320323-0331032211000123) |
| `outside_virtual_network.tenant` | [outside_virtual_network.tenant](data-sources--fleet--reference--group-002.md#canonical-3203332013333021-3020301220131033-2320133132321200-2300020002033210-0112313021111323-0020233202333230-3320320331313330-2102103122123222) |
| `outside_virtual_network.uid` | [outside_virtual_network.uid](data-sources--fleet--reference--group-002.md#canonical-3233120331332313-2113223113332212-2333023301231001-2212023012301000-3122123212033001-1212313332030210-1331023330212133-3330133230223032) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-2111111130321311-2300101203011331-3220311012120031-1221113012031202-2312313021300030-2122102313312112-1002310021312323-2201201301223020) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--fleet--reference--group-002.md#canonical-1210123332001023-0011032303002121-1033032132012021-1330022010331011-3300130002101100-1122130331103132-2123012133310233-1031320311233200) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--fleet--reference--group-002.md#canonical-2130031300102213-2233313012101201-1213102112212103-2201122033300030-1033223103010320-2222110003021331-3301033002103112-0133122013310030) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--fleet--reference--group-002.md#canonical-3303003031332303-1112210012311132-0113313322002101-0101123232303302-0101203110332222-1320230323220320-0112002132223230-1003322323303101) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--fleet--reference--group-002.md#canonical-0031303022113010-0022213303333213-0311232230023023-2233021113131110-1330032322002131-1133102221212133-1310230212013221-1303010131120213) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--fleet--reference--group-002.md#canonical-0130110200111003-2302002032101102-2113310203202112-1300200331103323-2222202030022323-2033213021013003-3311231013102231-0132332303011202) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--fleet--reference--group-002.md#canonical-1230221021223321-1330010222210131-2211103001211121-2331130110302010-3131331200122303-2113033211011232-1022122101230230-3132020211011201) |
| `sriov_interfaces` | [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-2123333323013300-0110133102200320-1232210230303220-3011003312322200-3022100322000212-0311121301321002-2220103203200120-2223000320220210) |
| `sriov_interfaces.sriov_interface` | [sriov_interfaces.sriov_interface](data-sources--fleet--reference--group-002.md#canonical-0112222100133120-2030101303200012-2113333303331123-3112021310103000-1112323121222122-1102121021232010-0233313133101031-1121312122100333) |
| `sriov_interfaces.sriov_interface.interface_name` | [sriov_interfaces.sriov_interface.interface_name](data-sources--fleet--reference--group-002.md#canonical-0310221020023223-1121323131312002-3311112323221210-3130133123211021-3112333011011210-0313222012201213-3303001232112130-3221320232313130) |
| `sriov_interfaces.sriov_interface.number_of_vfio_vfs` | [sriov_interfaces.sriov_interface.number_of_vfio_vfs](data-sources--fleet--reference--group-002.md#canonical-2111000223300022-3331212201322303-3232313303030113-1302111011210002-3103210130231311-2232323120132302-3120220233300012-2310200223332313) |
| `sriov_interfaces.sriov_interface.number_of_vfs` | [sriov_interfaces.sriov_interface.number_of_vfs](data-sources--fleet--reference--group-002.md#canonical-1023102302001031-3003122132310132-2120330033332332-3123122033313330-0300310112333000-2201330202213321-0303311013100022-3312303223020120) |
| `storage_class_list` | [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-1012302103312022-3331132220002110-1120220121220023-0000311103003122-0111002031331211-1000002000021111-3101030221022312-0130130013123023) |
| `storage_class_list.storage_classes` | [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-1123301131200103-0012020101121230-0233123223010203-3011313030110331-1032003223301131-3321210232002001-0133010001120330-0012211311302133) |
| `storage_class_list.storage_classes.advanced_storage_parameters` | [storage_class_list.storage_classes.advanced_storage_parameters](data-sources--fleet--reference--group-002.md#canonical-0032120033220023-2223220123322220-3302210001003130-1312220332100020-3301132223111033-1021321113123013-0310321111211010-2220220311231313) |
| `storage_class_list.storage_classes.allow_volume_expansion` | [storage_class_list.storage_classes.allow_volume_expansion](data-sources--fleet--reference--group-002.md#canonical-2131112102201023-3231101321131121-3033021331102121-3131011130322303-3230012330033123-2000231202333303-1231302020001013-2210110331100020) |
| `storage_class_list.storage_classes.custom_storage` | [storage_class_list.storage_classes.custom_storage](data-sources--fleet--reference--group-002.md#canonical-1011032211012230-3003330123101032-2221332003020331-1332102020103102-0212032031133333-2311211120122330-3302123002003123-2322311321023313) |
| `storage_class_list.storage_classes.custom_storage.yaml` | [storage_class_list.storage_classes.custom_storage.yaml](data-sources--fleet--reference--group-002.md#canonical-3223313220020302-1233232332211031-1100101110320331-0112332103110001-0233322200311130-1120002312212310-1333130301013122-3320203321331030) |
| `storage_class_list.storage_classes.default_storage_class` | [storage_class_list.storage_classes.default_storage_class](data-sources--fleet--reference--group-002.md#canonical-0302113031200312-2220302113220101-3001101120212003-1112323122001300-2132032011202032-2223002201303332-1202022330312011-3132222011032312) |
| `storage_class_list.storage_classes.description_spec` | [storage_class_list.storage_classes.description_spec](data-sources--fleet--reference--group-002.md#canonical-3121203321320310-3313020133321013-2100133002223210-1202023211310331-2200013013320132-2030231301021110-0131330331320300-3120323030030220) |
| `storage_class_list.storage_classes.hpe_storage` | [storage_class_list.storage_classes.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-1033320212320320-1323231013103122-2211011120321203-2102220331113233-2131132020021312-2312101132100223-2013022220203132-2123100130121120) |
| `storage_class_list.storage_classes.hpe_storage.allow_mutations` | [storage_class_list.storage_classes.hpe_storage.allow_mutations](data-sources--fleet--reference--group-002.md#canonical-1022130100122133-0032132302222132-3120210100110132-0302201100331031-2312031130120011-2200210223320320-0101231000322002-0011103211302210) |
| `storage_class_list.storage_classes.hpe_storage.allow_overrides` | [storage_class_list.storage_classes.hpe_storage.allow_overrides](data-sources--fleet--reference--group-002.md#canonical-0110300302201122-0130112011330220-2113312010001321-3132311102023123-2121101121222330-2102320031130011-1333231122301321-3012112200330322) |
| `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` | [storage_class_list.storage_classes.hpe_storage.dedupe_enabled](data-sources--fleet--reference--group-002.md#canonical-1310011022323021-0020011200310200-2012103122030303-0203223113111300-0211123331112302-1133020133200211-3221003122203212-2230122310301223) |
| `storage_class_list.storage_classes.hpe_storage.description_spec` | [storage_class_list.storage_classes.hpe_storage.description_spec](data-sources--fleet--reference--group-002.md#canonical-3202300032023110-0203313012000301-3321013002023003-0301003200013130-0333032013233212-2102322130302300-0111210213231131-3132132102103110) |
| `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` | [storage_class_list.storage_classes.hpe_storage.destroy_on_delete](data-sources--fleet--reference--group-002.md#canonical-1230300330313101-1030221330022300-2222113311332131-0332023311223133-0030022310020212-2131202330333121-0101031100112011-3222033311020020) |
| `storage_class_list.storage_classes.hpe_storage.encrypted` | [storage_class_list.storage_classes.hpe_storage.encrypted](data-sources--fleet--reference--group-002.md#canonical-0121200001130232-2101032303201123-3033231030320023-2001313112320222-2021103203102033-1332221001230022-1321211110011033-2103330022121311) |
| `storage_class_list.storage_classes.hpe_storage.folder` | [storage_class_list.storage_classes.hpe_storage.folder](data-sources--fleet--reference--group-002.md#canonical-2000321213313102-0132330302330331-3110112323111121-3203232233011012-0100321113002030-2022232132000230-3231022022303022-0331332223022012) |
| `storage_class_list.storage_classes.hpe_storage.limit_iops` | [storage_class_list.storage_classes.hpe_storage.limit_iops](data-sources--fleet--reference--group-002.md#canonical-0131130210231232-1011332332223221-0011031101133103-2011031323332322-2133321211020321-0132312110033123-0320213212311213-2033230030322310) |
| `storage_class_list.storage_classes.hpe_storage.limit_mbps` | [storage_class_list.storage_classes.hpe_storage.limit_mbps](data-sources--fleet--reference--group-002.md#canonical-3333003020323302-0131031012023331-0113112203003121-3122021031003211-3311132312122113-1103320332023001-3012332300032012-3002103001332330) |
| `storage_class_list.storage_classes.hpe_storage.performance_policy` | [storage_class_list.storage_classes.hpe_storage.performance_policy](data-sources--fleet--reference--group-002.md#canonical-1100202222021302-1303032030010312-1331020310120022-0120322123010200-3213230132201010-2100233210120322-3102122033302233-3322023310233232) |
| `storage_class_list.storage_classes.hpe_storage.pool` | [storage_class_list.storage_classes.hpe_storage.pool](data-sources--fleet--reference--group-002.md#canonical-2023021101233313-0132221303122120-2300313131333113-1121323112310201-2233012022222120-1123200000032320-0203012221321030-2123223211031112) |
| `storage_class_list.storage_classes.hpe_storage.protection_template` | [storage_class_list.storage_classes.hpe_storage.protection_template](data-sources--fleet--reference--group-002.md#canonical-2232230020221112-3131101330212320-3232312300002022-0131201202311211-3312202201001203-0223232213313330-1131233020101003-3321311211201211) |
| `storage_class_list.storage_classes.hpe_storage.secret_name` | [storage_class_list.storage_classes.hpe_storage.secret_name](data-sources--fleet--reference--group-002.md#canonical-1202220133122332-0210233223020133-2330120123212203-2012202312202311-1210231013000032-1110003121021013-1123113311220321-3223102330223022) |
| `storage_class_list.storage_classes.hpe_storage.secret_namespace` | [storage_class_list.storage_classes.hpe_storage.secret_namespace](data-sources--fleet--reference--group-002.md#canonical-1120210030131033-0202303200113323-3202301031003212-0221322202223100-2331202122111331-0323033310123210-3301332323021302-2131203021013000) |
| `storage_class_list.storage_classes.hpe_storage.sync_on_detach` | [storage_class_list.storage_classes.hpe_storage.sync_on_detach](data-sources--fleet--reference--group-002.md#canonical-1003133213102210-2000323123130030-0123001101220312-0132230032320012-1333130110133330-1102013323113201-2300103312002130-0111112122221203) |
| `storage_class_list.storage_classes.hpe_storage.thick` | [storage_class_list.storage_classes.hpe_storage.thick](data-sources--fleet--reference--group-002.md#canonical-2320322113303121-0323223311100012-2032012003302203-3202201200313030-1031102211001200-3030022012310010-1110233131212001-1003213332321012) |
| `storage_class_list.storage_classes.netapp_trident` | [storage_class_list.storage_classes.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-1010132203332021-2231322223023212-2021203320331211-0123330220103120-1232011232011221-3233123112011200-1221311123000212-3220301033330112) |
| `storage_class_list.storage_classes.netapp_trident.selector` | [storage_class_list.storage_classes.netapp_trident.selector](data-sources--fleet--reference--group-002.md#canonical-1210032121301031-0233321201120230-1200222111312013-0030030332100322-3212012233222300-1032111130223131-2130223031221031-1011332021232121) |
| `storage_class_list.storage_classes.netapp_trident.storage_pools` | [storage_class_list.storage_classes.netapp_trident.storage_pools](data-sources--fleet--reference--group-002.md#canonical-0103300203220011-0203210310021323-3230323320322211-1230032233330212-2010331133132313-2130123231330201-1320113010122310-3123212301210220) |
| `storage_class_list.storage_classes.pure_service_orchestrator` | [storage_class_list.storage_classes.pure_service_orchestrator](data-sources--fleet--reference--group-002.md#canonical-0212332120132020-0223223111102131-0311301300333231-0303221111021331-3311132212301322-3120311012031201-0313132021133101-2102112311313100) |
| `storage_class_list.storage_classes.pure_service_orchestrator.backend` | [storage_class_list.storage_classes.pure_service_orchestrator.backend](data-sources--fleet--reference--group-002.md#canonical-1300022322110132-0311231320020103-2010121223000322-0011232201320102-1023333110001320-1302321311123300-0010322022111133-2201310230331332) |
| `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit](data-sources--fleet--reference--group-002.md#canonical-3300130132110032-0011122313100011-2121123010300322-1320330301101122-0100000210130311-1311200120302133-1030213111110310-2303001222133021) |
| `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.iops_limit](data-sources--fleet--reference--group-002.md#canonical-3303230333030000-3310303011032122-3123130102333210-1121202013302231-1000320111123111-0021300033002031-1331031320000300-0002323332233011) |
| `storage_class_list.storage_classes.reclaim_policy` | [storage_class_list.storage_classes.reclaim_policy](data-sources--fleet--reference--group-002.md#canonical-0332233100132100-3313211102031202-1112201310330102-0022222000222310-0232210221321102-1213002013033322-0131230211001303-0323303222221311) |
| `storage_class_list.storage_classes.storage_class_name` | [storage_class_list.storage_classes.storage_class_name](data-sources--fleet--reference--group-002.md#canonical-0113021022011012-2301301010330021-2010030022323311-3203122001132300-3211310030221033-1212102113021233-0302130322323220-0120130330200020) |
| `storage_class_list.storage_classes.storage_device` | [storage_class_list.storage_classes.storage_device](data-sources--fleet--reference--group-002.md#canonical-0222132210301222-1001230113113223-2111213320101320-0333330123022132-2200112322132221-0212301120313203-1121120010312233-0332231212202112) |
| `storage_device_list` | [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-0101313010021002-1132021231301111-1333230303120031-2300003010103003-0000212332000030-2023011303230331-3130120302303021-1310301012210220) |
| `storage_device_list.storage_devices` | [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2332132030000132-3131331222133020-1110320023312230-0200222110002313-2113121002302103-0330133100123100-2100103222312003-1002113031210330) |
| `storage_device_list.storage_devices.advanced_advanced_parameters` | [storage_device_list.storage_devices.advanced_advanced_parameters](data-sources--fleet--reference--group-002.md#canonical-0103123113203113-3301021330112121-0221100230210030-1002312300202303-1113233003300321-3213123122130032-0333103000031031-0333103312021010) |
| `storage_device_list.storage_devices.custom_storage` | [storage_device_list.storage_devices.custom_storage](data-sources--fleet--reference--group-002.md#canonical-2033333111100031-0322201230313331-2103013012013113-1333033103312210-1123011002330201-0221030302130102-3120322000132033-3300302230021112) |
| `storage_device_list.storage_devices.hpe_storage` | [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-1212112201320032-3131303222211210-1233131001130100-1110001310233102-2232031303300033-2321102313222101-0112011000331022-3321320102323221) |
| `storage_device_list.storage_devices.hpe_storage.api_server_port` | [storage_device_list.storage_devices.hpe_storage.api_server_port](data-sources--fleet--reference--group-002.md#canonical-2113032022331123-3121132312013300-2100213212300022-1200013130113330-0031302223110322-2230213220201123-0303122230312220-1201101323111213) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-3002300233200011-1312213130100100-0232313112203300-2322011203133130-0211323013020113-0310103001321022-1112000003032201-2013223333322232) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-2200102010012102-3010112120320011-2301020213233100-2001031132131231-0231033312202200-0301323321020001-0102031012133103-0112303331001112) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-002.md#canonical-2230321312010112-1123113310200301-0321232033231221-3210000023212000-2200210103300302-3202303131301120-2012122202312322-0103031120220210) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location](data-sources--fleet--reference--group-002.md#canonical-3311021013020013-0012130313112002-1303311121022223-1213021130111112-3012310201233333-2133330320202120-0132311300012033-2332310002001011) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-002.md#canonical-0330322222223333-3003221111202133-1122230220231133-3002202031102032-0113310113001003-2300013021332132-1221033310311322-0300221203332030) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-3211331323113323-1330100002010213-0332211113233321-1223230302031003-1202033221022000-2313033120020202-2020110103330323-3031120321211021) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-002.md#canonical-0210231030032132-0223203201003202-2233213313222302-2222023011331312-1023032100113211-0310233123311023-1232212111010010-0033303303321203) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url](data-sources--fleet--reference--group-002.md#canonical-2301121123110031-2113211011232000-3111330203211001-0123301300122330-0003310012112112-0110111010223303-3133200232132302-2030232200101323) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_user](data-sources--fleet--reference--group-002.md#canonical-0103200223000220-0331022103311121-3123010030033321-1020223100323331-3232132232112030-3301130310202332-0130200213032122-1200102022230233) |
| `storage_device_list.storage_devices.hpe_storage.password` | [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-002.md#canonical-3022311202213013-0111211232111011-0020331030232013-2120323021100020-2021201033230011-3021113112111000-2021020310302203-0301123201132133) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-3320020032311103-3213233232313221-3330123202331211-1031232003012013-3032112103330110-1231002223302100-2323103131223302-1223023011021300) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-002.md#canonical-1121212011333300-3332211131221132-2120320030021303-0230321023023200-3230101310131212-1331032231112012-1023022301123030-3200132100122322) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location](data-sources--fleet--reference--group-002.md#canonical-2122101103100010-3011123211303121-2200132303232310-1032123211213331-3201303312002320-3330132112332320-0001320301103321-1213311022110200) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-002.md#canonical-3302011003230232-2331023023303110-3221322311211002-0300130120011320-1231321330100130-1203203113210011-3133121012201012-1311332302303130) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-3000312121032023-3210130231331133-0033213103303301-0011323122303113-3003021012300000-2031213023013232-0320113132120331-0111213313003323) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-002.md#canonical-1020302012322333-3121201332233132-0001323120030131-3002203233323322-2100110313121102-1100010213123330-3031011030230121-0231122001110120) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url](data-sources--fleet--reference--group-002.md#canonical-3132233123232231-2020322101010020-1103312120011203-3322330330220202-1331321322322012-1001232302322100-2212310211123330-1233232000101011) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` | [storage_device_list.storage_devices.hpe_storage.storage_server_ip_address](data-sources--fleet--reference--group-002.md#canonical-2330312111310331-0223302321131011-1023231020313012-1133002323300110-1122101312331200-3223210202101120-2021030233003013-2132001013313112) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_name` | [storage_device_list.storage_devices.hpe_storage.storage_server_name](data-sources--fleet--reference--group-002.md#canonical-1220002021111231-0031223333310000-1200032100200232-2300331102021131-3301131102001101-0111101202113003-0133222033113130-1101322020101033) |
| `storage_device_list.storage_devices.hpe_storage.username` | [storage_device_list.storage_devices.hpe_storage.username](data-sources--fleet--reference--group-002.md#canonical-0210031030002121-2031200113300102-2032121230132330-0112121111031321-1000322110223032-1333200011232300-2100113303200201-1120101233131121) |
| `storage_device_list.storage_devices.netapp_trident` | [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3212322031313002-0000103303323330-0201110200213003-0303311233302032-2033023220010002-1102111001222103-1310010210112302-1100231333201300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-3233032001002013-3021100023130033-2100003322311201-2203310113201001-3033030102202132-3321311003001223-3130331230000323-2312111022010011) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--fleet--reference--group-002.md#canonical-0022021131213113-0121021030013132-0221122233332210-2001200320013331-0212121130300121-3332023312033233-2220321021010203-3013013210223102) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes](data-sources--fleet--reference--group-002.md#canonical-0201223101123222-1202333220121203-3232303023011121-2122013033113003-2021032010030002-0133012131013210-1312022003110232-0133313110100110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy](data-sources--fleet--reference--group-002.md#canonical-3311021333011023-1313302203121110-0231211011221321-2202320120020030-3322121330110331-3000203122312210-1312112302000000-1301333003313022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name](data-sources--fleet--reference--group-002.md#canonical-2101013322032001-2111233131211122-2122232203003020-3310111113112320-1323100332111232-2300112231213101-2211103221110302-1232313020120213) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate](data-sources--fleet--reference--group-002.md#canonical-2103232230102223-3323030322231200-1221320222113322-2023003022323121-0123030120020203-2233331121302110-2310102011310211-0213021311311103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-002.md#canonical-3101033301333120-2110302232130213-0201321203213322-1123313312020232-2312120120131001-1213131231233120-3222221121100113-0203332022300002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-3232222332110310-0311213202021101-0301110302122331-3033101031021033-1331101132212330-1122010322001030-2332123223120310-2121001123100112) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-002.md#canonical-1113011332230003-0101021320122313-0332011133010102-0000332100002300-3013311023122312-1011131323001120-3021100233201120-1320003100322021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location](data-sources--fleet--reference--group-002.md#canonical-0301002221330110-3302011030032031-2132313110130202-2103320123121211-0012332012220302-0221111222123200-2222131303213131-3032030020030022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-002.md#canonical-3100012231103323-3320312112013330-0130221220221231-3333103233121103-2013300123030021-2231012011201111-2203023333232320-2331321233023020) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-0110122013122001-3130103312010112-2310311132223230-1220030100313323-0223300111121000-1121113022220112-1300110003113121-2021133111233333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--reference--group-002.md#canonical-2120012311222112-0020313220110200-2122130123030221-2132200101303003-2301212202013202-2203333310130312-1031102110033212-0010023203032311) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url](data-sources--fleet--reference--group-002.md#canonical-1202012022133132-3331023301021132-1130332333233021-3231331213303023-2022201310100312-1312000231322302-1133211110320222-2033103231323111) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name](data-sources--fleet--reference--group-002.md#canonical-3201310203302233-1310321113133113-3220230230131001-1323133101132221-1101223120032303-3232223132320012-1133311011203120-2021021131200121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip](data-sources--fleet--reference--group-002.md#canonical-1300031232220123-2121120221113122-0100131301001201-3233203200203222-2132300030031322-1110101212200211-2112001303112210-3201111100211210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels](data-sources--fleet--reference--group-002.md#canonical-3112233010300131-3311231210010230-3321020232322120-3002100133233330-0010130122031233-2132333033322020-0203333103333101-0110330022112000) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage](data-sources--fleet--reference--group-002.md#canonical-1200312103123212-2013113231211003-1032211132231000-2223232100023021-1310200200111021-3011230032231100-1223313202000131-1301113100110131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size](data-sources--fleet--reference--group-002.md#canonical-2002322111030132-0312100202112100-0333302021010111-3332103310123103-2320201302201233-0021032131100032-2212331323112120-2121010030133123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name](data-sources--fleet--reference--group-002.md#canonical-1331031333022103-2130301132322311-0022011222311023-0013203312223030-1133302231010320-2212001202321221-1200032222220021-2012223323332023) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip](data-sources--fleet--reference--group-002.md#canonical-1011201032102120-3333111222212223-2113030202133313-3310311310212011-3323201333203103-0103320133001002-3002031203003321-1030333131012233) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options](data-sources--fleet--reference--group-002.md#canonical-0233011322112312-1200231111020212-3313213022310210-1203203122002003-3302213232333303-3330132312330322-0031310221312321-3223201033130311) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-002.md#canonical-0132202332301213-1211220100201111-3230100100111111-2020130003211121-2210003212200321-1032201202230332-0222201302310333-1022201101302333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-1200002210003303-2130111011102030-0212312132101323-2011331010232132-0101300121102223-1331322222301211-1000223300021112-3013202110201232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-002.md#canonical-0100113300020123-0100110013011301-3120213311111022-3133113232100213-1011311002121301-3200101130032130-3323113000131310-3012010110203310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location](data-sources--fleet--reference--group-002.md#canonical-0210030313301320-2002111031101330-0323200101121213-3000203212210213-1010113023221202-0301321303123132-0021200201222002-3302303032212303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-002.md#canonical-2223110010330200-1100202321121202-3130012333133333-2223121303123032-2032300010310231-1022300121110210-3322301101312001-2201321013220302) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-0111321322003023-2003101110132012-1020311222322113-0023233131322123-1033003033033120-2222312032031333-0020132101133010-1030112320321033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-002.md#canonical-3032211231103322-1122130200212330-1000102211010010-2223001221300020-1101122220201322-0331330203202022-3331312230002120-3300311023302221) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url](data-sources--fleet--reference--group-002.md#canonical-2203002033230233-2132033313220301-2121202301113222-2132233111133222-2232111130120130-2030321310011122-3321030120201101-0321331310132022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region](data-sources--fleet--reference--group-002.md#canonical-0303213303210031-0212122212133211-1201332213210201-0201203231300110-3203231311012210-2102330313322000-1200130200322313-1130111223323011) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-002.md#canonical-1022312031011321-1112330310300202-1221302332301303-1322221323332033-3120203020001121-3120320311023210-2031222121231101-3003213333303313) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels](data-sources--fleet--reference--group-002.md#canonical-0101032211001230-0130131011011021-2111232232033300-2212322033102233-0321323331032310-2013021133221113-2321102210130021-0021033123133030) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-002.md#canonical-2222002033211122-2322213312112112-1212313101200013-1303033333133103-1103301233211120-3100031131002033-3100313211110333-2300111231100331) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-002.md#canonical-0121302320001102-2021300120212030-3232232211113020-3023120010331133-2122112321210032-0210110013230330-3323212202132323-2230010310120022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption](data-sources--fleet--reference--group-002.md#canonical-1132133313322022-3001320213301011-3332122233212030-3110000010213223-2133030313031231-3313110123103231-3121020203001223-3220122111002033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy](data-sources--fleet--reference--group-002.md#canonical-3100211330213101-0233201122303301-0130022000113311-3303012320122001-0001211130120130-0123311223001302-2112113322101212-1000301133002033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-0332332220231030-0323322002200213-1231110312111111-0121233021032223-1110300032110103-3303230002322032-0330230022121001-1012200231021333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy](data-sources--fleet--reference--group-002.md#canonical-1132012000333113-3321133321101222-1202123300330020-0310012332001202-1000310003122023-0211100030311113-1010221003131201-0013020212000023) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style](data-sources--fleet--reference--group-002.md#canonical-2020131220322220-1322011001002223-2313100103112230-1122130101120220-3132130003200012-3201232322020110-1110231111333032-0132021223020332) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-2323110203001112-2012323102131010-0301302021133110-0112111313120201-2233223101232301-3321203000113031-0101112233231302-1033023330300301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-1303122323302033-3131333030232321-2313002231333131-3311131001220301-1210110231203220-1301012230123221-2231000202122122-0303032002230331) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-1312131131123032-2020120012332023-3102333320300113-0320221030020130-2012111112110103-2020313121010311-1031310022310111-0302223221110023) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-3131013312012312-1221200322320120-2103233022302222-0330201200313313-0322313102201012-3331202100112033-1202232021301002-1010221033233211) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-2222231031120220-3323023000012211-2310110332110121-3110130203033022-3113023330310310-0230121012100103-1223312211323302-3131220302202021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-2003323003031210-3331110123011303-0333223000000320-3123133313300202-3102212001210132-2330102220330023-0210032300113001-1011213222202033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-2331313110233112-1333312133031210-2332203033231031-2120132011321132-0130231222231221-2010210021031212-1333212213203222-3210222231310203) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone](data-sources--fleet--reference--group-002.md#canonical-2201031100230100-0021012102011032-3231200133001101-0333231311131001-0013002132101011-1020300301313120-0120000203321310-0021320030100002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name](data-sources--fleet--reference--group-002.md#canonical-3033301213032113-2330323113102132-1000203032311032-2323202020233120-1010203222202022-2021103123120212-2100323331303323-0021201011210132) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix](data-sources--fleet--reference--group-002.md#canonical-3020113222210123-0110200213132011-2223211130012322-2012122101330310-3113333132213000-2010230202230113-2132002220032311-3311122111210300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm](data-sources--fleet--reference--group-002.md#canonical-0013110023120213-1130221032313331-2120200112300133-0133222300011102-0102233332313311-3011303212201110-2133230222313002-1101000330221321) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate](data-sources--fleet--reference--group-002.md#canonical-2023303332033200-3230202031222033-3211200213300120-2133011200121013-2003130003031133-1213323202210112-2111330020213302-2000031223311021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username](data-sources--fleet--reference--group-002.md#canonical-3300131302300022-3302103033201332-2222303322201123-3010020033220222-1300332321122311-3030112010203203-2031222223013202-2132032023303120) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-3120220122313122-3321001102020332-2132201023202013-1001003321001101-2333303320313313-3102211332313313-3031321010311233-2122233011223210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-2223031010002302-3101233322331303-3301000120232131-2110231230033001-1132130011332202-0320001313233201-2110300210201133-3200122100132003) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-3110212201321300-0333001203312101-2013211131100310-0311233112002320-2012011200131223-1123001112100212-1003201112320003-1033123020110010) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-3000221123203222-3300232012131211-0233002213021312-2333111003222213-2230223123232122-1202100011113032-0012301301021111-1120131100021221) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-2021001022300031-3202100022231310-0300011112223000-3013303023101202-3120133011203220-1232312001320012-1101010021010320-1223000203133202) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-2213300111320222-2003130111233203-0200002221303123-2101201300102232-2022311333321233-2101331132101320-3330022310311023-0233122032330303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-1332011221032211-3112212301230301-0032113121023101-3102202022302112-2313110212013233-0122213221033310-2222220231013232-1101123332013023) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-1333313030112300-0123000330132300-3313022130311233-2111333223003130-1120033330021323-2301212323320322-2202303321301112-3201213003200033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-2012221032212021-3111203233312123-3230033223333031-2001212222230220-1131233321223300-1120012202320210-0000001023221022-1233313020110303) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-3210113120113130-1302003203220102-1301322313023200-1203210302320230-1000331222221113-0202222231021122-0121110203013100-2030230202000222) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-0030331122233001-3012301030303032-0202131032313222-3313322111323201-1031132322131310-3012132302323103-2021231030232021-2113102323323322) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-1322331331203332-3023110133211300-3131212012101011-2223322332121320-3112030321111222-3321300001011101-1102130011203213-3031120013201031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-0323011110310122-1331301002221123-3020113030121221-0003330131212322-0003103102133031-0333020212203002-2333201120311323-3220230211113002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-2203210122321211-2332220031010113-3222322310010033-1021322210033003-0133121332112012-3030111213321120-3212103032012001-3312231312101323) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0312112002310313-0011230320033000-0221223220022300-0123133033010010-0331311100212213-2001331113201230-3132333113002002-0311003130300301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate](data-sources--fleet--reference--group-003.md#canonical-1103102323312322-1132320211201101-0100201201230033-3003012001303301-1121023102010221-2013113032121030-2122102102120222-1311111311322002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-3031003302120210-1112103033332012-2322330201031200-2232132110032223-0033200211101111-3100223210012331-0233223020222332-1103232300111231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3200233233132022-2111132032132010-2303203113011311-3101221233200013-2100121322133213-0210030311010220-2002222103030121-1113323122100031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-1110023111312112-1231021123032003-1110211030300021-1230232130122322-1331132003000320-2210131221201221-3012133233002220-3203203121010223) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-0003321100231310-2010123033102232-3132333113201000-1312223331300120-1033311313322221-1303120021223003-1300102132220210-3101001201310121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-1201000000021200-0331323302302310-1003021011230300-2110012333013200-2300111332112113-1132032111001232-0232231221321130-0123333033300031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-1323002212231221-3031200330111102-1201032232101100-0132302300021133-3201311113123001-2331222030313221-0201031310312331-1131230222322201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-2103313223212303-1112220332121103-3003133033030230-3030302132003320-3121100112312031-1033230013310333-0123301002022112-3133132211303100) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-0232202332113011-1113232303133010-1201101302202131-3113031120131213-3133313220323232-1320133020022231-2201231310322021-1121021222230021) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-1223011231032031-3320201301010320-1212323222010312-3100102111303203-0111130310121120-1213013121213201-3132310111210203-2022213301201310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip](data-sources--fleet--reference--group-003.md#canonical-1213302220302210-2213013233012032-0033220221130121-0330121333331331-0302201300211201-1112011013213202-3030311030323131-0112212103200302) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name](data-sources--fleet--reference--group-003.md#canonical-2010323130231333-1023302320210111-2123330310131313-3333011123033023-1133133031201000-2321203301301021-2103022231110210-1321030331211003) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels](data-sources--fleet--reference--group-003.md#canonical-0222003121002303-1132113230233322-2101231313301301-2013213121330333-0231003300112203-0302210120022030-3020103211321322-1110032332121003) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage](data-sources--fleet--reference--group-003.md#canonical-3222131122000132-2121230300303330-1122203111112130-2331232021333331-3122100301113323-0000302131130221-2222002222211300-1022303212302231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size](data-sources--fleet--reference--group-003.md#canonical-0133123200301012-0120230111120202-0030200213020320-2132121010212232-1230210211113322-0311332323031013-1313320102131101-0312111031102122) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-3310233010320233-2100333322313023-3220103023333232-2112103123313330-1023022301221002-2233022002021122-3200101010322033-3020201001323031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip](data-sources--fleet--reference--group-003.md#canonical-1000220103221121-0321233031310222-3211232303013303-1020032031132101-1011330302010231-3032103023031321-3202301033110210-0103301113310032) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--fleet--reference--group-003.md#canonical-0310031112220200-0123213310312233-3212003112311332-2030013013221131-1302002123030320-3021012122132210-3301133303230322-1013013303310122) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-0310232031000322-2020333220332233-3102000221213313-3330031001223123-0013203031031021-1002202203222203-0221000322201112-2223022211013200) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-1011032010132111-1210033133110121-1211320232203110-3101201121202321-0102231300121111-0120310322130210-2210000002030013-1023222120222301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-3021103221322131-1213221210233200-0030213231113000-2311013000332121-2131311002003011-2103211010220220-2222311310110100-1111311302301102) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-0222222312313033-3133322100001012-1231132131002202-1013021310210311-3021112311102221-2033333010311201-1120013323231001-2213312332021123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-1311121320101323-2023130201222003-1231113013033131-2200120011331231-2331331110122302-0100022032123312-0020003332313002-0200213003031333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3231010021101212-2330203013213232-0333103330222012-1312023210230211-2221020033302012-0121121131222112-1013322011201003-2301211312132103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-3302311001112122-0023202123033212-3231102203230031-1232032102132122-3210030220020013-0323103131303101-1131320333133010-1231121320313321) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-1002323110330003-1303310332322020-1312023031110032-1020230110223320-0212312120311031-2200132223213012-2120220033102100-0023101031101302) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region](data-sources--fleet--reference--group-003.md#canonical-2020022010103223-2021130102210223-3133323031212321-3013332032332021-0013230203210011-1010302103032132-2220313301332002-3000002101331313) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0310312313110012-0202213323202030-3012112021002203-2301311202031310-2103221132110001-0023130211002111-0210010331321201-1200210032113110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels](data-sources--fleet--reference--group-003.md#canonical-0132220222322321-2032001131231003-1212230231132231-3302231220001233-3313323222000001-1211131031132002-2303313112002232-2123023323200031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-2203103131332010-2222303201333330-3323333211012333-1320201310300112-3120133211123031-2122033311121220-2001122113131130-3311303210030310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-2013320102233201-3300111130201133-0011223320001200-0010210032310022-2321000111002213-2232022010133321-0312331301210202-1132023010030333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-3330000302201033-1120323012301320-3302033320011131-3220113032320022-2310123013203023-2102010002303032-3113031231122213-2032122022110203) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-2101020230122302-3200133102212201-2011220101220020-2300330233212323-2203132311010020-0203002210202303-2002222113310032-1330123110313022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-2132313213220021-3232233203203002-1111201021233232-3200200321100203-3333021312131130-3132023011210130-0201023132223020-0213202033021210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-3010230100021132-0200112020110321-0002330300103020-2313223300322312-2011320003031203-0332012320210031-3032222303311300-3301112211003130) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-0221000231020020-1332131130032312-2330311131020220-1231102320021312-1112310200203312-0320200232131301-0321033201202230-3221201021220012) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-1123232011332130-0011201032221000-0211223223223103-1120202213011013-0331203301010112-2223011021023321-2213023232202223-1330211323001300) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-1102211102123332-1112230032312300-1313233212211231-1212133131200020-0333113120301323-1122122222132031-2210020331323023-0011312223223201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-2002300030332321-3033310313110003-2002101330113031-2311001310032112-2003003201021131-2100203203012100-0303210220011123-2210211003101101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-3002023203302231-2221001032023310-2211222202231023-3020032100330233-3030003201330023-3101130302023023-2202210130311310-3030000332031022) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-2320313331131302-1112030310210210-0230101233021223-2211101021030133-3101033331211320-2011033131232112-3302221120330201-3001031212120113) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-1030303301321303-0302111101300323-1021221100110002-2030111031101201-0322331311323320-2031313011331000-3302203021211212-2301132032231002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-2012300231311000-3021200322013233-2313312033023211-2130003133132132-2012112312011210-0202102231223003-1123213130300332-1202111130113001) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone](data-sources--fleet--reference--group-003.md#canonical-3010122312301113-3020310312300202-3033213301101001-0200202230333111-1303333323302230-1332000233331021-3222213231220222-1300022312133010) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name](data-sources--fleet--reference--group-003.md#canonical-0211322203302201-2300121302203201-3133221110121030-0033122031303030-2031031031223231-1002210021211121-3032200221120123-3001310123301101) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix](data-sources--fleet--reference--group-003.md#canonical-0322122121233033-3001031213030200-3002030113102013-3110230203230110-0231000221210030-1321222321220313-3201003112211322-1221133301322123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm](data-sources--fleet--reference--group-003.md#canonical-0133200121312230-3113011211332221-1102303012011010-2233221332110000-2313021102100330-2213210012012211-2313021120322331-2111132110313330) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate](data-sources--fleet--reference--group-003.md#canonical-3312110300012211-1112031120233122-1103331132210021-1133011111131232-1212110011221330-3311233220113312-0121232102130320-3200131222013010) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3022032202232031-2330032132010002-2220122221011220-2323310200322320-2323222110012203-3302020322231021-3023011223311111-2203231010130110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-1121122311111333-2010100230213332-2302011132320211-3220012023210112-1010310010132010-0131110031200101-1020110123220331-2021130301303103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-1210312020222122-0311102012211233-0021003031210230-0131220221032203-1113101200321321-2112312121030212-0112300123322021-0130212223322123) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-1120220122032302-3032231022031201-2013133023021022-0020021013132013-2012013320112300-3310232103032333-2321010330013200-2131123232332222) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-1301222302320302-1033103201301010-0013032000032113-2333132021002022-2110311221301021-2123130322322300-3333132132331232-2010121002220231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-3133211210222110-1303113032323320-3021121023221203-2212300223123323-0323311023313213-1220312210223000-0320210331013030-2232200130220230) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-2223020212022331-3131033330011233-3232030101122132-1333030132122021-0031133120030033-2012201301110021-0012122110021011-3003030022302031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-3113003011211123-0330012032002031-3113301030003211-3230130333202022-0030123030102020-2123031311212233-3112300212311101-0131001120133333) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-0221013221131232-1313102331122231-2322101201212133-0320331010130130-3002330132033011-3130022311212231-1200112123123321-0310212123102121) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2111032202300332-0202103323100031-0223020330201333-1011031112200332-0333131202301203-0133301002033001-3102120001132110-3122312230312033) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2311100313330320-2301332312000222-0132012020012220-3312112032210300-0311211333202033-1123022020032013-0321320230230133-0133210020313031) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-2211320130111310-1113210033212101-3030102032300111-0131130203300231-3111210302033221-0211213112213233-0123001102223320-3113203303300310) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-1012211210321310-1321131033011022-0300322113231233-3210331033023011-2001221103330130-1213022100232211-1020321310331111-3032231012313210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-3310131330110201-1330211300211302-0012022130321133-2011213013112032-2003333211310323-1122330203113230-3222200301023030-1013022113011332) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-1011000022331322-2202333231231210-0330313312000132-3313012122010202-3233212323000010-3223133011120330-1322222332322100-1331230010220312) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-3130210001220331-0233213130111131-0012120132323313-0311211332030320-1100220033012120-1020303113232212-3311233132001000-2320202111032102) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-2013110222210210-1213300101120031-2233101311031203-0110310031232100-1311030201012233-0123330130010330-3311003321113101-3233110103031131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username](data-sources--fleet--reference--group-003.md#canonical-2131322311223003-1221000333013200-3013031110012322-3213230321322110-3231130002311102-0000331300131010-3331222203103330-1100210003121301) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username](data-sources--fleet--reference--group-003.md#canonical-3133233032230213-1222110300230203-0123211033313023-1120202133101201-3021103202131313-0313201022132200-2001122200012300-0311003232302201) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username](data-sources--fleet--reference--group-003.md#canonical-2033003311022231-2100213321002033-3013221033210111-3310100003012033-2001012002010201-1010200313133320-0112303303223033-1302012333123013) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1332120030303100-1011030102300031-1320033110110031-2221101323321313-1001122120023202-1122303033122221-2011320210112110-3301302212233320) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-0223312230221031-2223213122223030-3201123231233102-1211211210013000-0000101231233033-1122112233103021-0211132103033023-3030330123123131) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-3011310020100202-3001033230102321-3010133231100130-2120311321122211-2001003321333132-0030130133301023-2023310113301230-1123311130320231) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-1331313020000122-2332221302010221-0120000130322233-2112010230022122-0222233113312112-3331123001232100-1033222001000120-0131002032012223) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-1103111130302023-0013203012111220-0020223120232110-0230113320122320-0310232321323221-3302022113031001-1123232211313310-3113322022133202) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-1001113023012001-3131000323101332-2012021220333012-0003122123313331-0303213333202120-2031203313001111-1320221032022110-0310231121103212) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-2323313100102031-3223022212130221-3323123312200312-0100010032031300-1021202031130002-2033120300302300-1100133201212101-3011101000332110) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-0303201022200003-3011211131102101-1332033220300030-3331223031121001-2211311003130023-3332233122023222-3030201033101011-2320130020313210) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-2330002311031200-1323302032003100-3221003230330001-3231333232100011-0213330222321311-3030110000231210-0223011003330213-3321020023303103) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-0010203023233122-0111302032300213-1300023003311020-0013222202311101-2312322231222022-1203011221233230-3120032300110000-1101230000031130) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-3032231002021200-2033111103132301-2331300103000022-3202013221031012-3300211032033002-2303300120022223-3131133023310111-0001221122132232) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-0132231321211230-3002222301200020-1221220123322223-0113001131301220-2113312122022311-2213200320001023-0230013133022022-3211131033002002) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-2002101112132233-2132010020221320-1032230122321301-1231312010330200-1333220131012313-1310020323212332-3001333210011011-3123031102202001) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-3333032202313302-3232311230323102-0130002213122033-1002321012023232-3233102321222120-1230131003203310-0220111330012033-3033031132022130) |
| `storage_device_list.storage_devices.pure_service_orchestrator` | [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-3103200021113021-2113212020022021-3033203132002300-2030331010033303-1110331103303232-2221010220120101-2001201013001121-1002013130103233) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-1220133332301021-2220302313013310-2232231112111331-3331213003133231-0002230133132323-0301013231023201-3323320200020303-0002121312333132) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-003.md#canonical-0302201022230330-0012011321221330-0120310021103012-0300303211221032-0002221201013200-0302113310212312-0101323330202020-2223032211113330) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt](data-sources--fleet--reference--group-003.md#canonical-2131223011212022-0110323200102202-3102103122033213-1220231130221110-0000320233110000-0232013123213010-0031221003021212-0011211101013011) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type](data-sources--fleet--reference--group-003.md#canonical-1130112130001332-3011220300332020-3031033330012131-3111221300123102-0111120122221201-3010230320223201-3310302032111100-3223203232330311) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts](data-sources--fleet--reference--group-003.md#canonical-2111332201122310-1111021121110001-1020230131023332-0111303001130013-1321332210011010-3011021231333313-2330002203310222-2233032111212102) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments](data-sources--fleet--reference--group-003.md#canonical-1312021032030033-0330322302300111-1231130221323323-1203030311310311-1310320211220302-0201200231101221-3301133210231120-3012120331002102) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-003.md#canonical-3103113323013021-2013123110001301-0032021233130100-0021032201110132-2311200320133212-3200330121131233-0010210220220301-1222312111012121) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-003.md#canonical-0312220003010031-3202200103111110-2313231221013322-0223033121020023-2122311002211321-1000000122320110-1331102130132220-3302030133210122) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0303332331030102-3213230100301230-2102130323020100-1223212223021120-3213212310022210-0133322103031123-1230212210233232-1102330200231000) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-1301213221022011-2231221221103202-3000112200213012-2231000223301013-1021212033020000-0233132122102103-3112003130032300-1031203031001113) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-1120002232121121-0020133330312102-1233002123022023-2222100012031210-3333012000012201-2312103322322322-2001001013303301-0320031321323322) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-1112231322311101-0130133230303320-1202231133221001-0131233031230031-3112310032112301-0322201122320222-3112130212301223-2021001220103100) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0021013322031303-1233222001012201-1212301210302101-2132101033133303-2030233032122331-1202122300122130-3001002022322231-1131000320223323) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-0233021131130012-1003022133313212-2323213030011331-0103220221112132-2101301100023102-0300300203023310-0012332223131302-2031201303321010) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-3231022313133212-2310231300232001-3103202220310300-3110013003310323-1122322232300132-2000030202322112-0312013133130130-1022013302120000) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels](data-sources--fleet--reference--group-003.md#canonical-1112222310111031-2023132301131211-2021102110323103-2132021132303230-0011321132112203-0332222000221312-0313301331033311-3031333302322022) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name](data-sources--fleet--reference--group-003.md#canonical-0200001223210200-3012103103032021-1001302330322103-2310112332122211-0213121022001203-0131313323330230-1023131312323120-1130002320320110) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip](data-sources--fleet--reference--group-003.md#canonical-1002222202122210-1022200331213131-0330331202330031-2101113133230232-3002120313020201-1022032203033121-3210010302310020-2012130033302321) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout](data-sources--fleet--reference--group-003.md#canonical-3112101121232013-1311320131013102-1213323211210312-3003330003321133-1023101001021133-1032312121322322-2220212123121302-2303030132133312) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type](data-sources--fleet--reference--group-003.md#canonical-0302122123033100-0300322222122110-0131213232020222-0100030303311300-2100323000202203-0301101023301203-3112320301113213-1000301301033213) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-003.md#canonical-3301212300021000-3022111311233332-3222331022202022-3011032333032301-1301213320333212-3331122330001310-3311232031032221-1313100121032322) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory](data-sources--fleet--reference--group-003.md#canonical-0023111222020120-2130321133020000-1031123032323023-2302302220113121-2102102203132201-1202302320212330-1020122100230011-0030123122231300) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules](data-sources--fleet--reference--group-003.md#canonical-0233020311321021-3333233212201203-0110201322310133-1232233133231221-1013131132302333-0333213300033002-0101200203201323-0231123121011010) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-003.md#canonical-3113322203312112-2320103102230022-3101033023320312-2123121313032122-2320123021333000-2221033111110100-2303310202201022-2321230333121021) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-003.md#canonical-0021033303010333-2212110030332110-1021101310013202-3130011001030221-1230101032033020-2010231023223021-1331300022303313-2002202112102001) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-1112133021203233-0120323301013121-2310320323113212-1213322002201232-1201210133330322-3031022123330112-3200321232300102-1322202130011321) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-3200121032302100-1223033021102330-0212023000201132-2003102322122100-0121002101021123-0121131122210223-2000001100131110-3020130000303330) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-2033022230311112-3332001132303112-0011201221103031-1212103232021201-1031313223211322-0231311213203213-1010110221330213-2213131120022331) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-3333203001133002-1213102111310013-1302100000233223-2103221320222321-1201322011130331-1132032220331213-2012211200100232-0031323033000032) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3003220022210311-0331322102200122-2321001333331321-2301120010110122-0021231013302303-3230323001313012-3210201221000130-2201031121321213) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-0121230202312130-2223021112300133-1102032031212112-0322322113100211-1001030211122201-2011321022333332-2003312022331102-3330021111013212) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-0230221020030020-3311113313312211-0101000331332133-3223033303212102-1311303002131123-2111313230332302-0132312032231113-2203102321221102) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels](data-sources--fleet--reference--group-003.md#canonical-2333020233221223-1203231313233301-0011030002121232-1321030012230003-0212131322032202-3203121132213300-3333002320200322-0111332122332113) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name](data-sources--fleet--reference--group-003.md#canonical-0122232110112012-3300332203201123-3010320001220022-2132312332032022-3203111132202033-0113021022002101-2230213123302100-3110131023000101) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip](data-sources--fleet--reference--group-003.md#canonical-3101133123322320-3112101220012312-0100301313203222-0313121231233223-1132123032233133-2130031233203210-0120301230333002-3013022301113032) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name](data-sources--fleet--reference--group-003.md#canonical-3131003020012111-0320223201311102-2313012303001211-2322001033221203-3311332031133023-3010221210303013-2020220111103120-2221022020202032) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip](data-sources--fleet--reference--group-003.md#canonical-1232001313211322-3320331330033020-2121120223221232-0133312101320332-2033032303030333-0023311101013021-0300012310331231-1023033220223132) |
| `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` | [storage_device_list.storage_devices.pure_service_orchestrator.cluster_id](data-sources--fleet--reference--group-003.md#canonical-2101031202010310-1031221221321122-0112033111020322-1313301032211103-1233131223333332-0202020103320112-0202332120012200-2220202203231332) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology](data-sources--fleet--reference--group-003.md#canonical-2100311003310111-1001222033122302-2211102021322333-1223100012322100-0233230012221110-3002332033331023-3321111002210030-3330021013101222) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology](data-sources--fleet--reference--group-003.md#canonical-1323013031203322-3200122311011320-1310311133223331-2232312032311312-1122300022321220-3123113032020031-2122031111110210-1303021330101313) |
| `storage_device_list.storage_devices.storage_device` | [storage_device_list.storage_devices.storage_device](data-sources--fleet--reference--group-002.md#canonical-1133223211122022-1310133110233323-0100211113310010-3130303310132221-2130102103200102-2000122010113011-1030010012210101-2103031201022020) |
| `storage_interface_list` | [storage_interface_list](data-sources--fleet--reference--group-003.md#canonical-2121301312201313-2003303021323200-1113203311322120-1012103002103230-2013220110320003-0220331131210223-0321010211132303-2320301002023322) |
| `storage_interface_list.interfaces` | [storage_interface_list.interfaces](data-sources--fleet--reference--group-003.md#canonical-1122201013213312-1213321213321112-2333130011001003-0003030300101031-2312332033001313-3003313032230033-1202330033301311-2111302313321313) |
| `storage_interface_list.interfaces.name` | [storage_interface_list.interfaces.name](data-sources--fleet--reference--group-003.md#canonical-0212332102111200-3320032000321203-1202332311211102-0201332310320001-2123322132033202-0311313020113102-2231302303222121-1211123330131031) |
| `storage_interface_list.interfaces.namespace` | [storage_interface_list.interfaces.namespace](data-sources--fleet--reference--group-003.md#canonical-3301020201020100-2010220120303110-0100133332112211-2322233313332201-2310030111130300-2030312131130013-2202002101112113-0302100130332202) |
| `storage_interface_list.interfaces.tenant` | [storage_interface_list.interfaces.tenant](data-sources--fleet--reference--group-003.md#canonical-3101313223200101-0132311203200313-1322123133110022-3230123003300203-2312213320223020-0123013210323210-3233330203023100-3222222230220121) |
| `storage_static_routes` | [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-3012331233311111-0123111321131231-3213032311022211-3000231010022111-1001031021200012-0323022200021110-3333332210321120-3330100233233210) |
| `storage_static_routes.storage_routes` | [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-0002121312211212-2131010023231311-2303033332000203-0203113331203322-3113121222031102-0211010112002032-0330023132100311-1113010230012303) |
| `storage_static_routes.storage_routes.attrs` | [storage_static_routes.storage_routes.attrs](data-sources--fleet--reference--group-003.md#canonical-3310113203102321-0220331313021223-1210113113222021-0310011211133303-0111110313211010-0032313110102321-3010010021213202-2333111201101201) |
| `storage_static_routes.storage_routes.labels` | [storage_static_routes.storage_routes.labels](data-sources--fleet--reference--group-003.md#canonical-3103023220123112-1331103110032301-1100200232101013-3201303231110003-3013312332030102-3302312322303003-1232210221131112-0212323010101333) |
| `storage_static_routes.storage_routes.nexthop` | [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-1332000133212021-1023230000201303-0003013230221303-0220302202132323-3232023313031201-2332132010133323-3301122210110313-3101120330300132) |
| `storage_static_routes.storage_routes.nexthop.interface` | [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--reference--group-003.md#canonical-1101320030003212-2103123221332333-0010002001331313-1232300103011123-1223231022210033-2010030223001133-3203133330300112-1112003030322011) |
| `storage_static_routes.storage_routes.nexthop.interface.kind` | [storage_static_routes.storage_routes.nexthop.interface.kind](data-sources--fleet--reference--group-003.md#canonical-0221330321302132-3103132202133002-2123302200102120-1201032132113300-2230333012002331-1033033010200322-0021210303103330-2011101200020230) |
| `storage_static_routes.storage_routes.nexthop.interface.name` | [storage_static_routes.storage_routes.nexthop.interface.name](data-sources--fleet--reference--group-003.md#canonical-2212023012122320-3321321103022033-1201010103312111-0331321313230123-1012322300122323-0302112321100020-0002130110323230-2012022121320113) |
| `storage_static_routes.storage_routes.nexthop.interface.namespace` | [storage_static_routes.storage_routes.nexthop.interface.namespace](data-sources--fleet--reference--group-003.md#canonical-1030310330232201-2233001002222303-2103102022310111-1311113102221330-0031330210101032-3223300030102220-2133220131110133-2211032323122200) |
| `storage_static_routes.storage_routes.nexthop.interface.tenant` | [storage_static_routes.storage_routes.nexthop.interface.tenant](data-sources--fleet--reference--group-003.md#canonical-1021030213203231-2301210022200121-3320002231001133-2101212100303013-0201113232130320-0023220100223220-1133002213000322-2022132303302202) |
| `storage_static_routes.storage_routes.nexthop.interface.uid` | [storage_static_routes.storage_routes.nexthop.interface.uid](data-sources--fleet--reference--group-003.md#canonical-1001311231020131-1013001312320310-0013210013303131-0001300123233331-3023023102332011-1223101103222230-1010102020011321-2101323011120121) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address` | [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0323212123132300-1112331232012132-3220213323231021-3003031110230301-0232111220311022-3123030133110010-2112312012122202-3121103331200021) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-003.md#canonical-1103200233003002-1221133113200013-2203032010230202-1301330010003110-1130233112112303-3311033311031021-1223022331201021-2210133303232312) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](data-sources--fleet--reference--group-003.md#canonical-3011310122320321-2331020130230233-2322331013201032-2302311201113132-3030132023113010-1032003010212323-3210313113330220-0302321003223023) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--fleet--reference--group-003.md#canonical-3030312222002332-2211112333101022-0322122013122221-3102032202132012-0120010001132030-0200313331123031-0200023233311101-2302132032023210) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](data-sources--fleet--reference--group-003.md#canonical-2110212021130032-1202110122000102-2022332332133233-1021113013101033-0003310011220110-0300000112312130-1131012212223231-2010220331210113) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--fleet--reference--group-003.md#canonical-0031332122230123-2320102032122203-2003313022030331-1120323013330312-2023023211113331-1121100332012200-3000230100211031-1011302210221012) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--reference--group-003.md#canonical-0110331321002230-2010332312131131-2130132323111022-0211020101231312-1310131222211333-2131211122133231-1030010113211303-1111302000302131) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr](data-sources--fleet--reference--group-003.md#canonical-3013233232200232-3131103301210110-3200030331021322-0210323303213331-0201202013202110-1212001331011012-3300010111111203-3131223130120313) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--reference--group-003.md#canonical-2220130211322320-0302002013002220-1322021313011120-0103321002032111-1021101300231010-1302203202301020-3222203303312210-2031333123230212) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr](data-sources--fleet--reference--group-003.md#canonical-2300133101021302-1202013020231010-0113210321032102-2003313301230211-2320211110230220-3202012010223001-2220033111010303-3233131032032232) |
| `storage_static_routes.storage_routes.nexthop.type` | [storage_static_routes.storage_routes.nexthop.type](data-sources--fleet--reference--group-003.md#canonical-1130201332320032-2211130120201322-1313001230111021-1213111203222323-3030233222333221-3213301320002103-1002302121200331-2111221210121322) |
| `storage_static_routes.storage_routes.subnets` | [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-0121220231232033-2201120120322120-2020231221003131-1112312103331211-1033023301303323-0003123032310111-0003201333213233-1113303033221230) |
| `storage_static_routes.storage_routes.subnets.ipv4` | [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--reference--group-004.md#canonical-3111021123013120-1113020100122333-0331321321332011-1130023020312203-3021313031121130-0120232323121310-2323100221032020-3213122112301233) |
| `storage_static_routes.storage_routes.subnets.ipv4.plen` | [storage_static_routes.storage_routes.subnets.ipv4.plen](data-sources--fleet--reference--group-004.md#canonical-1010201232103233-3212022220013201-1002123030101122-0302001220102323-0102322310321333-1021012131133233-2330301101201121-1301020122301222) |
| `storage_static_routes.storage_routes.subnets.ipv4.prefix` | [storage_static_routes.storage_routes.subnets.ipv4.prefix](data-sources--fleet--reference--group-004.md#canonical-3330122023022133-3021202132103230-0202020213021100-1330111313110213-1222233312330011-2021023332121002-0232120230103223-3222203120111201) |
| `storage_static_routes.storage_routes.subnets.ipv6` | [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--reference--group-004.md#canonical-1330332231122313-3220022201313212-3332030102322033-3021113101203313-2003130320102221-0320322202003210-3203013113103010-3333033111020013) |
| `storage_static_routes.storage_routes.subnets.ipv6.plen` | [storage_static_routes.storage_routes.subnets.ipv6.plen](data-sources--fleet--reference--group-004.md#canonical-2232103210312231-1221103313021222-0123100131001300-1200212030320313-1001231022002031-0320022201220110-2131100210201001-1313203333010000) |
| `storage_static_routes.storage_routes.subnets.ipv6.prefix` | [storage_static_routes.storage_routes.subnets.ipv6.prefix](data-sources--fleet--reference--group-004.md#canonical-1112320312302203-0021021011033332-0032012322211232-2012212033000321-3202031300201103-0032002230032012-1022200231333022-2021310011201221) |
| `usb_policy` | [usb_policy](data-sources--fleet--reference--group-004.md#canonical-0031113210110330-1320123122203030-3102000002211020-0300002012223023-3100323131030300-1303202220232201-2302030110112130-2113203111101330) |
| `usb_policy.name` | [usb_policy.name](data-sources--fleet--reference--group-004.md#canonical-0301221010202211-0132101021111021-1312222323232010-2212232001211122-3211332100232100-1323120210013011-3020222032301300-3132112021201001) |
| `usb_policy.namespace` | [usb_policy.namespace](data-sources--fleet--reference--group-004.md#canonical-0003220322203031-2113233001023022-2013102230031303-0031013310022321-0122032333010211-1010010131111121-0333333203300100-2320010023301322) |
| `usb_policy.tenant` | [usb_policy.tenant](data-sources--fleet--reference--group-004.md#canonical-2011103303103012-0131103303222023-1331321033102010-0030230212112233-0301102001302101-3113023320012123-0322223321300332-1232203200023021) |
| `volterra_software_version` | [volterra_software_version](data-sources--fleet--reference--group-001.md#canonical-1010001010110200-3120320132311103-3010311202213130-1300100213231221-2003030232001232-1301003332011103-2310033011213010-3203301123121122) |

<a id="canonical-1020303212333102-1221132312100331-1230301230330311-0322102012012202-2302112223332023-1131133311222203-0200222121232321-1312210000123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_usb` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- allow_all_usb

<a id="canonical-2021101313013013-1033323101021333-0310311232320022-2131030212320211-3310303321320310-0100213023130322-0110302012202312-2123232131300330"></a>

Type: `["object", {}]`. Computed.

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

- [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-2021101313013013-1033323101021333-0310311232320022-2131030212320211-3310303321320310-0100213023130322-0110302012202312-2123232131300330)
- [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-0323201230021333-3213210223102010-1222323310230210-1003010300303302-3103023120303323-3331021111202320-1202222300110102-3301322010230232)
- [usb_policy](data-sources--fleet--reference--group-004.md#canonical-0031113210110330-1320123122203030-3102000002211020-0300002012223023-3100323131030300-1303202220232201-2302030110112130-2113203111101330)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003013010222032-0231112013202331-2201001123333132-3030101330330213-0200210113111232-3323123102012112-0310000103030013-1100222131113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- blocked_services

<a id="canonical-2130211202002331-3222231301111321-3133110030013133-0102103110210231-1002002013322330-0303000031102103-0220121201202323-2312202030201211"></a>

Type: `"list"`. Computed.

Disable node local services on this site.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1320122211003211-1233020100302322-3031230100010333-2330122102203231-2322310311202033-1002033303111211-1100002220311221-3213020113030231"></a>

### Direct properties for `blocked_services`

- [DNS](data-sources--fleet--reference--group-001.md#canonical-2020131323013200-0022030321000000-1103112001121002-2233320223300313-0010211002131013-0020000122233112-3221030031332100-1001113021203113): complete subsection reference.

<a id="canonical-2120311323121000-3313300213022212-3312003033123302-2013013112300301-3110123112001311-0332223323213331-0231210010330001-1003311323232310"></a>

<a id="canonical-2321033213102031-0102110021000112-0202221212020012-1102213002013030-3021002032221313-0133302133322023-3222113103212123-2223112013213020"></a>

#### `blocked_services.network_type` property

Type: `"string"`. Computed.

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

- [SSH](data-sources--fleet--reference--group-001.md#canonical-0010321313313123-1232322310312033-3202300133302310-0103210121331230-2113322030013310-0312101302001320-3311211021023121-2033032200231213): complete subsection reference.

- [web_user_interface](data-sources--fleet--reference--group-001.md#canonical-2012003100030131-1012333011300111-3230002230002232-0221030100230232-0031202121232202-0012001120113311-3223332301102022-1012100311101001): complete subsection reference.

<a id="canonical-2020131323013200-0022030321000000-1103112001121002-2233320223300313-0010211002131013-0020000122233112-3221030031332100-1001113021203113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.dns` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-2003013010222032-0231112013202331-2201001123333132-3030101330330213-0200210113111232-3323123102012112-0310000103030013-1100222131113031)
- blocked_services.DNS

<a id="canonical-0311023200202021-2122303303322122-0223032021231303-0020022102323232-1112332312121100-2021100321321112-1133133301200333-3002221113222132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010321313313123-1232322310312033-3202300133302310-0103210121331230-2113322030013310-0312101302001320-3311211021023121-2033032200231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.ssh` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-2003013010222032-0231112013202331-2201001123333132-3030101330330213-0200210113111232-3323123102012112-0310000103030013-1100222131113031)
- blocked_services.SSH

<a id="canonical-2301010330320113-0102212222001110-2312002123133232-3122332313313131-3210030321302213-0312311131332111-3202301203131113-1220301323303330"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012003100030131-1012333011300111-3230002230002232-0221030100230232-0031202121232202-0012001120113311-3223332301102022-1012100311101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.web_user_interface` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-2003013010222032-0231112013202331-2201001123333132-3030101330330213-0200210113111232-3323123102012112-0310000103030013-1100222131113031)
- blocked_services.web_user_interface

<a id="canonical-1232300013003002-1203201233323332-0223021212323012-0232201003203202-1312011032312122-2322032020100302-2031131023033033-1202000200303002"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211330203133133-0100030312311131-2320023111103033-3211001101101231-2100331200213120-3133220022001023-1132231031111302-0113330202110213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- bond_device_list

<a id="canonical-2111330100223331-3300130221220232-3123203101223113-2231113033200102-3323120223131223-0110320213233022-3223001020113031-1112013312012301"></a>

Type: `"single"`. Computed.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

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

- [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-2111330100223331-3300130221220232-3123203101223113-2231113033200102-3323120223131223-0110320213233022-3223001020113031-1112013312012301)
- [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-3023231313233222-2123102110130331-0131021311233301-0232210010022131-2100101001012022-1230233223123212-1231312301132111-2110320121201112)

Select alternatives according to the provider validators above.

<a id="canonical-3012323110032012-3110111311102200-0122001202201213-2313211200220330-1030312012102010-3011200002330103-1201121100121330-2313223033032311"></a>

### Direct properties for `bond_device_list`

- [bond_devices](data-sources--fleet--reference--group-001.md#canonical-3123232301320012-1320101211002101-1320321033221102-0222320002231010-3302020312102031-3201121211312332-2032022112202201-3030031333001001): complete subsection reference.

<a id="canonical-3123232301320012-1320101211002101-1320321033221102-0222320002231010-3302020312102031-3201121211312332-2032022112202201-3030031333001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-1211330203133133-0100030312311131-2320023111103033-3211001101101231-2100331200213120-3133220022001023-1132231031111302-0113330202110213)
- bond_device_list.bond_devices

<a id="canonical-1032210001011102-0030333202303211-0030310021310013-3133303332013012-3312332220230231-0231331003130131-3233122123203020-0230223201022102"></a>

Type: `"list"`. Computed.

Bond Devices. List of bond devices.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3200211113030022-2312223301113102-0121302220122330-0302203203121131-3111123300103130-2231031200222122-1230003112000230-1010313202112003"></a>

### Direct properties for `bond_device_list.bond_devices`

- [active_backup](data-sources--fleet--reference--group-001.md#canonical-0112030301033001-1231312003110030-3001331003122311-2130300031033002-3023331132011130-1203113000200320-0002000022011201-0321311211032132): complete subsection reference.

<a id="canonical-0003213311031122-3211112011321113-0002303130011102-0231023321321111-2320012210012101-0312003001003133-0013233010021230-0313000200323222"></a>

<a id="canonical-1233020132102303-1311113113123320-0221100231312021-2210212101320113-1010320320331212-3100213030210233-0302001211333201-0232113123023323"></a>

#### `bond_device_list.bond_devices.devices` property

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [lacp](data-sources--fleet--reference--group-001.md#canonical-1310203023031223-0020020212201211-0320233131200022-2101112313111021-2212233112110130-2230003120030113-1200121301220330-0213011022032322): complete subsection reference.

<a id="canonical-1301001010330210-1102301301130332-1232120220212133-0012301121323312-3213020120313011-1122012313111203-3332201121200220-0122211102003131"></a>

<a id="canonical-0113320302113301-1303232330113300-2011300202111231-3330323022133222-2100301011011330-2232200232303013-0210233323322012-2123013112233233"></a>

#### `bond_device_list.bond_devices.link_polling_interval` property

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1122332121001121-3310101112213322-2012031300010211-1110201221221330-3031023031010211-1023130120132311-2000030200031203-0111223133121010"></a>

<a id="canonical-2022131301011102-2301013110302001-1212112131321001-1230302131330301-3312031330120030-3332330310222110-2001211211001033-0332122302000100"></a>

#### `bond_device_list.bond_devices.link_up_delay` property

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2231223333021020-0232311320013010-0223231323313232-0231213112122023-1003232130223012-1213112101131030-3210202300032200-1230112002022131"></a>

<a id="canonical-3320013001313000-2320311302131313-0333020131101311-0201122130303201-3233112303003232-0221313020320323-1301302213111012-0123221110213300"></a>

#### `bond_device_list.bond_devices.name` property

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0112030301033001-1231312003110030-3001331003122311-2130300031033002-3023331132011130-1203113000200320-0002000022011201-0321311211032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices.active_backup` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-1211330203133133-0100030312311131-2320023111103033-3211001101101231-2100331200213120-3133220022001023-1132231031111302-0113330202110213)
- [bond_device_list.bond_devices](data-sources--fleet--reference--group-001.md#canonical-3123232301320012-1320101211002101-1320321033221102-0222320002231010-3302020312102031-3201121211312332-2032022112202201-3030031333001001)
- bond_device_list.bond_devices.active_backup

<a id="canonical-0303012002032201-3333211033102202-0121120211203330-0122310221211203-3000020210202031-0000220011122030-3232222101232223-3102312011032300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310203023031223-0020020212201211-0320233131200022-2101112313111021-2212233112110130-2230003120030113-1200121301220330-0213011022032322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices.lacp` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [bond_device_list](data-sources--fleet--reference--group-001.md#canonical-1211330203133133-0100030312311131-2320023111103033-3211001101101231-2100331200213120-3133220022001023-1132231031111302-0113330202110213)
- [bond_device_list.bond_devices](data-sources--fleet--reference--group-001.md#canonical-3123232301320012-1320101211002101-1320321033221102-0222320002231010-3302020312102031-3201121211312332-2032022112202201-3030031333001001)
- bond_device_list.bond_devices.lacp

<a id="canonical-3131010200302002-2330001123210203-3113213033233302-2212323333120111-2211203332021032-2333123010022321-2201312212330012-2232122213102031"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

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
