---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-0212213212202223-2123222303200222-1320221312320123-2102220113112310-0310200212031303-3211233212230022-3300112310323311-1331032012320112"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client — dhcp_client / 122313222112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client

<a id="canonical-3131110320302101-3321011101312001-1000001330132311-3103313201031213-3333333113123121-0323032323033131-1221023001311022-2323120132030103"></a>

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

<a id="canonical-0220330231332021-2300020313332130-1210221003000002-0110310331010203-0012032022003330-2221111000120122-2202333132303233-0303213113033031"></a>

## Direct properties — dhcp_client / 122313222112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330110321303221-0010021300213031-3230333021311302-3210101112032111-3021032310022332-3102310011101022-3321033022010210-3232300202130033"></a>

## Next pages — dhcp_client / 122313222112 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031203122130111-2303013200221213-3320002312303202-1031201001321120-2113203011010301-3320321200303203-3023201010232012-0331212031002222"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server — dhcp_server / 333012020201 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server

<a id="canonical-2020213012122132-0020230330201033-2232120322311022-3221202201321221-1332300003120013-2230222232103321-1332013322121201-1220231130333101"></a>

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

<a id="canonical-1331232030010120-1330030120211020-2210033132331202-2110232131113032-2130112020023123-3030033312030312-2121232332233332-2101112212102022"></a>

## Direct properties — dhcp_server / 333012020201 / 3

- [automatic_from_end](resources--voltstack_site--reference--group-004.md#canonical-2132011012203332-3103112000123231-2321210231110110-2330102321202221-1312313203102223-3103110233101231-2031130123100103-3000300311130123): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--reference--group-004.md#canonical-3102113302120103-2231023230230031-2210320322313112-2202013310323203-0223230013213101-3003300303232021-3103101022002130-1232030311301121): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132): complete subsection reference.

<a id="canonical-2230131313202001-2010132323122230-3233002033210132-1221202033123310-3233202131021013-1332011022001220-2221011102322230-2231222131211011"></a>

<a id="canonical-3302113232212230-1032130130321202-2103312032120113-2131213010100201-2023013132130030-0000201311231110-1303031331310330-2123212103011122"></a>

## dhcp_option82_tag property — dhcp_server / 333012020201 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-1202322300030202-2202301201323313-1110000320113210-0003133102323121-0031033103133230-0030211220121233-1121131022112022-0220310020102132"></a>

<a id="canonical-1203333220302121-2130212033110320-3133300310222212-3000332211002322-3223103121311201-1030101310003020-0101200033330112-1013102230022212"></a>

## fixed_ip_map property — dhcp_server / 333012020201 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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

- [interface_ip_map](resources--voltstack_site--reference--group-004.md#canonical-0122313020313103-2333232220202201-2302222320312331-1210220323022212-3001232022133133-3333322123333003-2010201112213113-1312122323103202): complete subsection reference.

<a id="canonical-1112023230223120-2201101110001330-0100212303210222-3102120200302033-0321133222333121-3110311113212212-0233133021000033-1321130321001030"></a>

## Next pages — dhcp_server / 333012020201 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--voltstack_site--reference--group-004.md#canonical-2132011012203332-3103112000123231-2321210231110110-2330102321202221-1312313203102223-3103110233101231-2031130123100103-3000300311130123)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--voltstack_site--reference--group-004.md#canonical-3102113302120103-2231023230230031-2210320322313112-2202013310323203-0223230013213101-3003300303232021-3103101022002130-1232030311301121)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--voltstack_site--reference--group-004.md#canonical-0122313020313103-2333232220202201-2302222320312331-1210220323022212-3001232022133133-3333322123333003-2010201112213113-1312122323103202)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2132011012203332-3103112000123231-2321210231110110-2330102321202221-1312313203102223-3103110233101231-2031130123100103-3000300311130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210003133132322-1300213303111310-3130230022030132-3031201233322001-1100132201012212-2121022202122123-0113032132021010-3333203311113230"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end — automatic_from_end / 132100223223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-1301320331302331-2121331200101312-1230130110030332-1213113123201310-3210323320130030-2232013220230002-1133313232011332-0023031322002122"></a>

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

<a id="canonical-1113111331212213-2331021232210323-2031232322123030-1200121112031303-2202031101232201-1131312212121030-3310310232032213-0320222201200003"></a>

## Direct properties — automatic_from_end / 132100223223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010221322022331-3121300311133311-2132323033232000-0203112110102031-3211202310013122-3033012021333122-1300221223222011-0333022012321302"></a>

## Next pages — automatic_from_end / 132100223223 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3102113302120103-2231023230230031-2210320322313112-2202013310323203-0223230013213101-3003300303232021-3103101022002130-1232030311301121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221130333232033-3210332303012103-1231203220111012-1330030030101011-3203332232002331-3211212023330131-2320212323302332-2221311203013201"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start — automatic_from_start / 122113032312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-1010211211201002-2321110033200333-1301220120322202-1021320312133233-2200021101313130-0211210323131013-2023020031123123-3323310133032313"></a>

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

<a id="canonical-2031232011002131-1133011231120332-0101000003113201-0112110222133010-1221312330120303-1033111222000021-3031231313303021-0133232222222102"></a>

## Direct properties — automatic_from_start / 122113032312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120111131133323-1133013332100013-2230131123133113-0202301021001021-3120011031102103-3030203133100002-2210301321111222-0211233213032122"></a>

## Next pages — automatic_from_start / 122113032312 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000220100223300-2321221203112302-2223111312231013-3232130022012101-1330312230023031-1310002211002231-3111233131020321-0230100311120222"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks — dhcp_networks / 021022132211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-3131032011310303-3011320211223030-3202223023131312-0201333120312102-3021121013012110-1012000002112332-2001112213031211-3032201220133303"></a>

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

<a id="canonical-2202333302012323-0322032011310033-3332120203030020-0200030210020202-3023103320100200-1002110120321121-1121232102202002-0300103000221112"></a>

## Direct properties — dhcp_networks / 021022132211 / 3

<a id="canonical-3330231123120222-1023133322201323-1211212130312232-2330121113222310-3121202230000120-0131220210213112-3221100012110123-0323201202022002"></a>

<a id="canonical-1221110030112132-0303222110103110-0301320223011113-1032213123211330-1113132102221020-1132120103123032-2131030321013300-2313030123210031"></a>

## dgw_address property — dhcp_networks / 021022132211 / 4

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

<a id="canonical-3312033321120212-1203230220003211-2213221332110110-2122001301023011-1232101330013202-1222310323321300-0210012211320013-0321313211011213"></a>

<a id="canonical-3121010131032001-0310331333200200-0111300223200121-1032231120132032-2121201032012232-1031100113330322-1012103331313012-3113222330011133"></a>

## dns_address property — dhcp_networks / 021022132211 / 5

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

- [first_address](resources--voltstack_site--reference--group-004.md#canonical-0102013102313200-0311201201133132-2201302132011121-1232313312133013-0032023113201221-3133332022003301-1320020013202200-3201213131011333): complete subsection reference.

- [last_address](resources--voltstack_site--reference--group-004.md#canonical-0121021002112230-1310121212330030-2322203122300222-3322312100311133-0300111023201021-2121321331021321-1003102222332111-2312113323230202): complete subsection reference.

<a id="canonical-0230130121002312-1201023101320132-0322133012212122-0220201112330333-2313113303103330-3020310110312211-1012222000011102-1312101200020230"></a>

<a id="canonical-3102102000133130-3120023212102100-2033000321233233-1303130021220221-2201002032033200-3321332233322210-0132303201013013-0223121211010231"></a>

## network_prefix property — dhcp_networks / 021022132211 / 6

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

<a id="canonical-3312222202031331-2211310333103110-2020301003301123-2011330300330233-3212321322310313-0032001230020122-1223011010111300-3201011213313233"></a>

<a id="canonical-1003103021031231-2133312112322030-0230330221132122-1213133210122312-2103023323003031-0111211200330130-2132132003102313-2331311102002112"></a>

## pool_settings property — dhcp_networks / 021022132211 / 7

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

- [pools](resources--voltstack_site--reference--group-004.md#canonical-2000331322020122-0022010231000132-2323132113300212-2321330001223120-0002123311330321-1321130313212103-0031310123100320-0332303332010011): complete subsection reference.

- [same_as_dgw](resources--voltstack_site--reference--group-004.md#canonical-1311032113123303-3133203331301220-3302311331100312-1202130230001010-3101110003231313-1223230023232213-2123200311112121-0213030133221012): complete subsection reference.

<a id="canonical-2000111213122030-1131301012312201-3023213110323120-3110222101213102-0212320310122230-2230103210300100-2011303110032221-1201211313230330"></a>

## Next pages — dhcp_networks / 021022132211 / 8

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--voltstack_site--reference--group-004.md#canonical-0102013102313200-0311201201133132-2201302132011121-1232313312133013-0032023113201221-3133332022003301-1320020013202200-3201213131011333)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--voltstack_site--reference--group-004.md#canonical-0121021002112230-1310121212330030-2322203122300222-3322312100311133-0300111023201021-2121321331021321-1003102222332111-2312113323230202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](resources--voltstack_site--reference--group-004.md#canonical-2000331322020122-0022010231000132-2323132113300212-2321330001223120-0002123311330321-1321130313212103-0031310123100320-0332303332010011)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--voltstack_site--reference--group-004.md#canonical-1311032113123303-3133203331301220-3302311331100312-1202130230001010-3101110003231313-1223230023232213-2123200311112121-0213030133221012)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0102013102313200-0311201201133132-2201302132011121-1232313312133013-0032023113201221-3133332022003301-1320020013202200-3201213131011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202133120031321-2013111322113303-0033121211111301-1032021311231010-1203202330020120-3300122332213223-1022232301021003-1102202332013010"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address — first_address / 033330212201 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-3103303213021131-3303200300210021-3100100300312202-0210113012312331-2123012103102012-3313132203302300-1323013312123213-0303113323231133"></a>

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

<a id="canonical-3233223001230300-0331032302212112-0310002030000003-2020102301021011-3102320323131201-1212110102332233-1232331310000321-0302231003231210"></a>

## Direct properties — first_address / 033330212201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011223023131233-2033012213222213-3320121110003220-3011231123201321-3313102021211221-1323000311133012-1210320231233032-1130221223333202"></a>

## Next pages — first_address / 033330212201 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0121021002112230-1310121212330030-2322203122300222-3322312100311133-0300111023201021-2121321331021321-1003102222332111-2312113323230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023332012020321-2000230230201321-3310233222231320-0200210221221032-3002130101010220-0023121100001020-1101123032221301-0110022111211232"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address — last_address / 231221031200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-2033132311213301-1113033333302030-2211330323000120-0323300133120120-3330002322331010-1123310100333223-3032021201132300-1021023120310111"></a>

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

<a id="canonical-2221030301212012-3220312230310200-0021010031213233-3132120202100330-1201000200120111-0103030123032012-0312030133103232-3301332222002213"></a>

## Direct properties — last_address / 231221031200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102312331000021-1122101003200020-3123123231332200-1221013001110023-3313212331213021-0220210313123133-3132112002221320-1231233200012223"></a>

## Next pages — last_address / 231221031200 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2000331322020122-0022010231000132-2323132113300212-2321330001223120-0002123311330321-1321130313212103-0031310123100320-0332303332010011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332030231030323-2331232322332131-0011302002003211-1211322303002133-1222223001311231-3110002200100300-3321313103210332-2003003013030011"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools — pools / 033321022330 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-2032213233030222-1320231101211133-2111312011033220-0332200132031123-0123100233220003-2330122322212011-0022032202102103-3101133321113332"></a>

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

<a id="canonical-1202303122122022-1333211133312310-1133231133011103-0032231121102133-3220021200110233-3211102132030121-3022102021132210-1033321303330132"></a>

## Direct properties — pools / 033321022330 / 3

<a id="canonical-0021003312321200-3110113102302120-2103123223220003-0202133113111230-2122231022032211-3021321121012012-3212032111323221-2333103233301122"></a>

<a id="canonical-0330001131311202-2132202101203322-3210232300120310-1211032213302133-3300030231112113-3110003233130203-3020333230231110-3031032323112233"></a>

## end_ip property — pools / 033321022330 / 4

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

<a id="canonical-0310013210122002-1022121220111231-1210033223102030-2300300102013333-3301032222222032-1221313020332121-0002220300230332-0311311200011133"></a>

<a id="canonical-2330302113122322-3213032112232021-3100200011210030-1133301212030000-3331213313333333-1103133332022313-0231133323101023-3202331311022120"></a>

## exclude property — pools / 033321022330 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0210110222131021-0113132333023231-0310212201123323-1202021222320213-3131302331322132-1331102033203122-3303131312312031-1332301323321032"></a>

<a id="canonical-2110113002200302-3212221333322200-3320100230300101-2322301100320123-1223332033031130-3120100001233021-3203230133223110-2312012030220313"></a>

## start_ip property — pools / 033321022330 / 6

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

<a id="canonical-2222103300002121-1130012303211210-3210233320112213-3131232002200131-3113102213130101-3300333211030100-3122022220033222-3100121212220011"></a>

## Next pages — pools / 033321022330 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1311032113123303-3133203331301220-3302311331100312-1202130230001010-3101110003231313-1223230023232213-2123200311112121-0213030133221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013030031201000-3213230311223210-3330021330000112-0131120323201232-0203232122322003-0101232210201020-1000201213212333-1230320021302101"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 331232030113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0110000201301313-2102322101110010-3123010120231013-1023330302203013-3221211022001313-3001211313333011-3033222100231100-1022310121130123"></a>

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

<a id="canonical-0301012120303233-2213033220123003-1200231100230310-0230011111230000-3000022330211012-2201232000330023-0031302121011213-3011133223033322"></a>

## Direct properties — same_as_dgw / 331232030113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003030232210013-3110100310013230-1000313303323001-3102202230311213-3201213213133232-0310303300312331-0302202113012212-0123322203200130"></a>

## Next pages — same_as_dgw / 331232030113 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0122313020313103-2333232220202201-2302222320312331-1210220323022212-3001232022133133-3333322123333003-2010201112213113-1312122323103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123111301022210-0011310000123321-1132003210322031-2013300223030313-3212121022302321-2212321032223230-3020001231213301-0011011131032300"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map — interface_ip_map / 322211020312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-1120233020110132-3133003210000113-1312233200032313-1003003210232312-1023002231110233-2232002031331213-2201322022201130-2123221131131330"></a>

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

<a id="canonical-0203121230332132-3033301101303023-1301202132332110-1311031113231013-1333212232323333-2023100130310203-3211321230312222-0002303212101211"></a>

## Direct properties — interface_ip_map / 322211020312 / 3

<a id="canonical-3001122131310013-1110112033302233-1012033200131130-2030001123133233-2310032131233233-2023232313000203-2112210312322110-2112310103002211"></a>

<a id="canonical-0122102103112013-0323202132232003-3003303002012203-3202120332333013-2213302221312323-2011101030111213-1221030310203133-2331333330220000"></a>

## interface_ip_map property — interface_ip_map / 322211020312 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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

<a id="canonical-2020013223130100-1102033220330311-1010102032202103-1033010221130331-3032001023231230-0100011120132232-3330033122212300-0111311033103023"></a>

## Next pages — interface_ip_map / 322211020312 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--voltstack_site--reference--group-004.md#canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131120332220120-2203102111013020-2130211203133302-3110231100022001-2003111320132230-3010322221100102-2310033030332322-3031110000010031"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config — ipv6_auto_config / 332110121332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

<a id="canonical-3313020123002310-2111130030220201-2002110032012320-3032010113322102-3333133223023031-1000111213101123-0112122200032100-0200031302111203"></a>

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

<a id="canonical-2331221301022323-0321203311330200-3203321030131123-0120333213011232-2032222231200131-3212012321003230-3233330013110212-1301232013213230"></a>

## Direct properties — ipv6_auto_config / 332110121332 / 3

- [host](resources--voltstack_site--reference--group-004.md#canonical-3032020222322110-0033232111020102-1130322112013230-0012202002222222-3321220010002032-2332213320331210-0120220113223020-0203110213131023): complete subsection reference.

- [router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202): complete subsection reference.

<a id="canonical-2322110313103012-0231111110011332-3132132230230320-2011212222033233-0132123320003013-1301101222123230-1100230111032213-3212021120113013"></a>

## Next pages — ipv6_auto_config / 332110121332 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](resources--voltstack_site--reference--group-004.md#canonical-3032020222322110-0033232111020102-1130322112013230-0012202002222222-3321220010002032-2332213320331210-0120220113223020-0203110213131023)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3032020222322110-0033232111020102-1130322112013230-0012202002222222-3321220010002032-2332213320331210-0120220113223020-0203110213131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121031303103103-1023101313021121-2320103213103312-0031022102000230-2232120012002312-1123301033112032-2320201011030321-1221131222301100"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host — host / 322032311332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host

<a id="canonical-0020310102022111-1032113033300122-2331202002113331-0020211201232101-3302131002000111-3311122222003002-0100330103203003-0102031120231312"></a>

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

<a id="canonical-2232322021110011-0001313203112331-3203032332211131-1002022031131133-0010112230021302-3030312013131000-3020010210121300-1030032120000002"></a>

## Direct properties — host / 322032311332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101303020113301-0321220203300120-1303131000311231-2222220233112230-1320032131231312-3130230031121231-0132232103032233-1123112113223200"></a>

## Next pages — host / 322032311332 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312300010130013-1031332301313221-1202223033113230-1312100010000000-1200303112331203-2021120112100013-1203332313132123-2101300300030002"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router — router / 021331001202 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

<a id="canonical-1023212330300132-2120013230013212-1312133311301130-0010331000320022-3132100322203123-1210320302122133-1311311013312321-0011311023230133"></a>

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

<a id="canonical-0331321110033323-0001322310100013-3012022230001302-3233101333013013-1221231212230312-0232223112012231-1202320030202112-1310313123222121"></a>

## Direct properties — router / 021331001202 / 3

- [dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133): complete subsection reference.

<a id="canonical-1233130100001002-2013020323100300-1203200030311033-3331113122031112-2132323133213030-1123101223030112-2300213132123333-3130311331033210"></a>

<a id="canonical-1301122320100011-1011032213121221-2130201301331101-0303333223030033-0111301330122311-3110330213202302-0233232120211302-3011311220320302"></a>

## network_prefix property — router / 021331001202 / 4

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

- [stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210): complete subsection reference.

<a id="canonical-2112132110001230-0010233333200322-3123321332001010-3103202203233313-0013101312302301-3221321023321223-0311212301330031-1221022130032133"></a>

## Next pages — router / 021331001202 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320310311033310-1221101133010012-0031232333321113-3130032002310032-1222023312013001-0210323200220310-1132102320022133-2010030301330322"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config — dns_config / 023220311120 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-1102101312312111-2211331202032013-0001003002211303-1123332031023133-2332302013112113-3123330211022221-1232122201022233-3120123022112233"></a>

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

<a id="canonical-2112332011301010-1333122312100103-2103233223212311-2203303223131301-1210330001323332-3031202331123121-0010212212210132-1210212212032332"></a>

## Direct properties — dns_config / 023220311120 / 3

- [configured_list](resources--voltstack_site--reference--group-004.md#canonical-3002221203213202-2003303112310110-3011102232032230-0121112301222020-0203023113131100-3121322023230130-3301312300230233-1211012233310331): complete subsection reference.

- [local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110): complete subsection reference.

<a id="canonical-1301033100332013-1321210332012121-0311220300300320-0020310301311213-0221101233322311-1231031320031313-3330032302331332-0003212310330202"></a>

## Next pages — dns_config / 023220311120 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--voltstack_site--reference--group-004.md#canonical-3002221203213202-2003303112310110-3011102232032230-0121112301222020-0203023113131100-3121322023230130-3301312300230233-1211012233310331)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3002221203213202-2003303112310110-3011102232032230-0121112301222020-0203023113131100-3121322023230130-3301312300230233-1211012233310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200132202212302-3123230123001130-3200312310002311-0031330023311010-1303133321231302-1223123221222012-3122020100202231-3103021012111013"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 023003130223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1110002310322123-0132120221303310-3021303000132111-3332020322233210-1110201031103103-3300120330111121-3203102022101332-3121020221323323"></a>

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

<a id="canonical-2312202233130020-0020333311203231-1103213332221021-1330121233010333-1333310313230203-1133300023232001-1023020010213100-0010221011023310"></a>

## Direct properties — configured_list / 023003130223 / 3

<a id="canonical-2331200211321121-1002233311023333-0222022123210033-1111232001112103-3001203332201003-2101120200303033-0233112112331103-1211333231123002"></a>

<a id="canonical-1013111201102012-2002320312131023-0302010233302233-0112020310102133-1220222320222133-2201133003231030-3100022222311331-1301321111223010"></a>

## dns_list property — configured_list / 023003130223 / 4

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

<a id="canonical-2003310132212310-3131322001331013-2102032203223032-3210223010323333-2133020331313200-2101313000020310-0310232212213023-0121332102210203"></a>

## Next pages — configured_list / 023003130223 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102022103121021-3110332211203330-1122120202020120-3230110202010131-1102112222002031-0031130312013103-0123122211333111-0220320131301213"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 021332233323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1320003011010002-1001113021300123-2020221221123102-3023302302220212-3130300330303121-2203112300233110-3013201013330100-2101203013212123"></a>

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

<a id="canonical-1010222101310301-1203223131131311-2112312121121203-0321231232313122-1331100211112330-1302220123022013-2233023012303133-3003303123231033"></a>

## Direct properties — local_dns / 021332233323 / 3

<a id="canonical-1212030110223011-2102323020222210-2230200210122303-0123121101202102-0120331211333023-3031121022011232-2332033223230000-2232001110031000"></a>

<a id="canonical-2012001300011211-1011303133003133-3230213121320113-3300101212002022-1233023203210303-3021201332131300-2212021123331132-2303030123232130"></a>

## configured_address property — local_dns / 021332233323 / 4

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

- [first_address](resources--voltstack_site--reference--group-004.md#canonical-0330222212332231-1321130010211312-1032323203123322-2022121122300210-1032200312002303-0321023131212302-0100211330033230-3213110030021220): complete subsection reference.

- [last_address](resources--voltstack_site--reference--group-004.md#canonical-1001313102110202-1322301130020300-0231013301020012-3123220111202123-0133233311232102-2131102212221003-0002033331311223-1303221231331200): complete subsection reference.

<a id="canonical-2110102332313013-3010013200303022-3331030211130120-1020302201113011-3031121220231032-0021320100313010-2211201220321333-3230211011211302"></a>

## Next pages — local_dns / 021332233323 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--voltstack_site--reference--group-004.md#canonical-0330222212332231-1321130010211312-1032323203123322-2022121122300210-1032200312002303-0321023131212302-0100211330033230-3213110030021220)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--voltstack_site--reference--group-004.md#canonical-1001313102110202-1322301130020300-0231013301020012-3123220111202123-0133233311232102-2131102212221003-0002033331311223-1303221231331200)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0330222212332231-1321130010211312-1032323203123322-2022121122300210-1032200312002303-0321023131212302-0100211330033230-3213110030021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110222031210223-0120233112012220-1000131201331113-0330310331100200-2022010021002220-2021121313001300-0000310231230322-0133002211332132"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 020022131111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3323330302003332-2112123230123312-1001221101010200-2321123030130303-2132232103030012-0031301312210002-1222120331213220-3200312133302302"></a>

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

<a id="canonical-2302201210121033-0113010011020302-3222101331131123-1131321133132310-1101303023013303-0230210203110121-3300011130000301-3101031321121311"></a>

## Direct properties — first_address / 020022131111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310112111331213-3020213231122313-2122330331230010-0323033333120003-3032100212030301-2220201113231113-1113020130322223-0223102232320311"></a>

## Next pages — first_address / 020022131111 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1001313102110202-1322301130020300-0231013301020012-3123220111202123-0133233311232102-2131102212221003-0002033331311223-1303221231331200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030231000232033-1232131201330013-3321223010030300-3213323123320023-1022013101032301-0200211331330223-2211302310201321-1201213132113031"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 130330000132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-004.md#canonical-1210013300021311-2113221213310333-0100220013100221-2212222201333303-3202103120213113-3012303321123120-3120110110221021-2012100211010133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0331232033100233-1022130301213201-0002033311202130-0301211103303210-0311202210130330-0120232013302332-2023130320320201-2020232233201001"></a>

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

<a id="canonical-3133122330021332-3331302112022032-3101130013303231-0211221203303221-1310110220233030-3300130033210311-0120100021231330-2312201020121301"></a>

## Direct properties — last_address / 130330000132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221032103321201-2320331011102331-1032200010203232-0010223200313102-0310022332033232-0122013130100110-0320230213223001-0233113031120113"></a>

## Next pages — last_address / 130330000132 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-004.md#canonical-2100132003202213-2122102233011122-2011102011020333-1002011332320130-0322121012032220-2300210322120103-2331211021010123-3121031120200110)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133030213012001-3021123313032231-1012013200131032-0332131313013232-3231010011102123-0200010121030110-2110213010203000-1031201303301232"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful — stateful / 222232231221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-2300033211022012-3112233320123310-0033130012212011-2023021213031100-1110023310020023-1003213303201010-3012333012232223-3332110301010200"></a>

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

<a id="canonical-3000303212233131-0032020033323301-0133022313113022-3010002003112033-3120001031102232-3010301310302312-1331010121010200-3130302311032132"></a>

## Direct properties — stateful / 222232231221 / 3

- [automatic_from_end](resources--voltstack_site--reference--group-004.md#canonical-3013012012313010-1120013202332223-1022031211330323-0212120303301301-2200120321213201-0103000212032311-0023121232303001-3200100332331220): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--reference--group-004.md#canonical-3220221010330230-3320313303023221-2032000103021302-2112300203021312-0220313210112003-2020302303200033-0123230203211301-0101210011301113): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-0322102113301303-2220022130213030-3223310222103230-3022033021322330-0020202313302230-1000231111333031-1112113102203301-0321111121032113): complete subsection reference.

<a id="canonical-1012311230113033-2110133300201001-1113000123220321-3000132121113023-3203303210230303-0032300312031000-2013332003003220-1212331120023203"></a>

<a id="canonical-0210102012333121-2312033223301112-0123332201213330-0010221313112112-0021110321200030-0100103102100133-0303133232210330-2111311301300101"></a>

## fixed_ip_map property — stateful / 222232231221 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--voltstack_site--reference--group-004.md#canonical-3210211133211011-0221123101320311-0210023223111320-0302123021011032-2222211331233002-3222133322102131-2301110011133023-1030002310203321): complete subsection reference.

<a id="canonical-3010103200333231-0133103200210232-0311101031011012-0112000123302022-3110100100123312-0322230011020322-0033002311312111-2003223333232122"></a>

## Next pages — stateful / 222232231221 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--voltstack_site--reference--group-004.md#canonical-3013012012313010-1120013202332223-1022031211330323-0212120303301301-2200120321213201-0103000212032311-0023121232303001-3200100332331220)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--voltstack_site--reference--group-004.md#canonical-3220221010330230-3320313303023221-2032000103021302-2112300203021312-0220313210112003-2020302303200033-0123230203211301-0101210011301113)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-0322102113301303-2220022130213030-3223310222103230-3022033021322330-0020202313302230-1000231111333031-1112113102203301-0321111121032113)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--voltstack_site--reference--group-004.md#canonical-3210211133211011-0221123101320311-0210023223111320-0302123021011032-2222211331233002-3222133322102131-2301110011133023-1030002310203321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3013012012313010-1120013202332223-1022031211330323-0212120303301301-2200120321213201-0103000212032311-0023121232303001-3200100332331220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003320333011101-3031233210212033-3223033310321301-3220311202231011-1112311323013231-0202220113033210-1001202030110210-0022123030132323"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 121100130133 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1011120320112123-3322221003333312-2211121303022232-1000101032020132-3212003132313121-3221132300303012-1230023330303113-1021301003132301"></a>

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

<a id="canonical-1022323330111310-2202201301211130-0332133102001323-2133320231300000-3233203220100010-0013322231001113-1112020203011333-3101131322012211"></a>

## Direct properties — automatic_from_end / 121100130133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301211122311013-2130121311323100-1330121012120121-3112131331301003-3130320231132220-1033230300302333-3302112100331112-0130303232110101"></a>

## Next pages — automatic_from_end / 121100130133 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3220221010330230-3320313303023221-2032000103021302-2112300203021312-0220313210112003-2020302303200033-0123230203211301-0101210011301113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121013121231333-1303123110132202-0122333010032023-2313002313232032-1101202000200002-3031110032233330-1212031330213333-0020213123301230"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 212233320203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2210201232201221-1200003313130010-3102221023123313-3230003101230101-2223302231020230-3211332311201332-3202222222323132-2002312200033030"></a>

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

<a id="canonical-2001123113110202-2132210111300311-3302133233022110-3130333231323002-2120011211331111-3203032211113320-3330002220003132-2202022033102222"></a>

## Direct properties — automatic_from_start / 212233320203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222311112021032-3230013032003212-1100331123300313-1332023221201310-3122323033233220-2212313303021100-2320030011301122-1231132301312321"></a>

## Next pages — automatic_from_start / 212233320203 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0322102113301303-2220022130213030-3223310222103230-3022033021322330-0020202313302230-1000231111333031-1112113102203301-0321111121032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201332232031000-1331112100222201-3102133032222332-1131133301222023-2110322112010112-0131013031003320-3103310200210200-0021112221233012"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 312103210022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0311332322022130-2131311213012300-0322201301203013-0001313220330003-0100320310110201-2310233210222330-0020121000001322-0022110122012330"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

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

<a id="canonical-3301213023333333-2112303131022302-3132133213221322-2120100121111113-0001132302013133-2300122110202231-1223233131011302-0230111321013013"></a>

## Direct properties — dhcp_networks / 312103210022 / 3

<a id="canonical-3021200131000033-1103021201103331-1322020310202013-1100333312132312-3212213033102233-0233112013302120-0220113023013112-2202113113111223"></a>

<a id="canonical-2022223213311012-2002323011132213-0200211302111023-3013201112101010-3121322212132223-3012133011330002-2333032210332132-1132100300111223"></a>

## network_prefix property — dhcp_networks / 312103210022 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-2111001031010303-2020223133302012-0020110303112002-3303031023221302-3211130323302221-0001331202301310-2023211022131331-1321131232331133"></a>

<a id="canonical-0002231133311010-1130013301120111-1332112023021031-3230023312233220-3031103111030010-1020200111210233-3332000122230223-0121301022010211"></a>

## pool_settings property — dhcp_networks / 312103210022 / 5

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

- [pools](resources--voltstack_site--reference--group-004.md#canonical-3213301131232323-3303331132010020-0220111030122023-0012322333122110-3100322320300112-0323210121321033-0233130313323301-1103130103312122): complete subsection reference.

<a id="canonical-2102311013113102-1330310132332132-0023231023230000-3011031020003132-2312020113002031-3120003010220210-2131212300133231-2120300100322320"></a>

## Next pages — dhcp_networks / 312103210022 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--voltstack_site--reference--group-004.md#canonical-3213301131232323-3303331132010020-0220111030122023-0012322333122110-3100322320300112-0323210121321033-0233130313323301-1103130103312122)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3213301131232323-3303331132010020-0220111030122023-0012322333122110-3100322320300112-0323210121321033-0233130313323301-1103130103312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011320102223213-2322333030333012-3111111000232113-1130221312300210-3321301002100120-0111122302131311-3000102101110132-0111302222310013"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 112230231110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-0322102113301303-2220022130213030-3223310222103230-3022033021322330-0020202313302230-1000231111333031-1112113102203301-0321111121032113)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1001330201222200-2322102221200031-0212021003322130-3100212110233001-0320211113000332-3213301331212011-1201131000113212-2033301220012200"></a>

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

<a id="canonical-2102003020123122-1023202313231202-3121223000213010-2322031231320032-2220301332333231-0112030130223221-2211331011133101-3213033301130210"></a>

## Direct properties — pools / 112230231110 / 3

<a id="canonical-2312313212102003-0310211302303301-1023233322222022-2112210202310312-2201132103313303-1111210013311020-0223110213111103-0333222032021333"></a>

<a id="canonical-0321132031112033-0030332012320030-1122310031121122-1102130033001322-3332123023332210-2131220133111023-3031230210023111-2000230333311012"></a>

## end_ip property — pools / 112230231110 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-0120010312321133-2003122022113121-0111323111130023-0021212320023310-1002133001103032-3111211312103231-2003123101221221-0321310222321122"></a>

<a id="canonical-2310222132302310-1312020111113313-0130022133301233-1210132031012131-3323031121302200-0303333301321221-1013330002122321-3032121133110313"></a>

## start_ip property — pools / 112230231110 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-2202211112121233-2011232113210213-0322002320021132-2221032013023312-1120101121221210-3321310231220303-2131133211100133-2302111311223203"></a>

## Next pages — pools / 112230231110 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-004.md#canonical-0322102113301303-2220022130213030-3223310222103230-3022033021322330-0020202313302230-1000231111333031-1112113102203301-0321111121032113)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3210211133211011-0221123101320311-0210023223111320-0302123021011032-2222211331233002-3222133322102131-2301110011133023-1030002310203321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201320230020023-3303213212202202-3032203033100320-1131132232000001-3111121110021301-2020331300101012-0031033003322020-0311230333323232"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 220230322133 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--reference--group-004.md#canonical-3313230013020222-3121031111320321-3203331121200111-2003112103021213-1223332311110330-2330110133232130-0030221103221122-3202302120132221)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-004.md#canonical-0000131230110121-0111313101000300-2110110321102311-0000110033332210-3203330122231032-0303210203313202-1323103031332122-2022003032231202)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3123321222121313-2011303212311212-0020230033311031-1011231120120132-1332221011322201-1111132122311002-3200203023231002-2230120310011103"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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

<a id="canonical-3231113213213232-1112232133223320-2111212100211003-2031021333031122-1332020100220030-1303112312020003-0100232311232131-1010113312011100"></a>

## Direct properties — interface_ip_map / 220230322133 / 3

<a id="canonical-1313222123021111-0100120303120323-1211003330303311-3103323301011101-1013013033220100-2310132132103121-3020003210301013-0010031100230020"></a>

<a id="canonical-2110102331212203-3132333110021202-0320322201223213-0203333311011332-0003000311000013-1230233030202001-0223111033131012-3121030000231300"></a>

## interface_ip_map property — interface_ip_map / 220230322133 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-1200303331000003-2021310333113213-1023311322322010-1333032231011201-3211212210001100-2121023231211110-1020210332300330-2031000300312330"></a>

## Next pages — interface_ip_map / 220230322133 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-004.md#canonical-3000210213223003-1131101123233223-3323133123232330-3223313122222222-2111131213110033-2122131012200331-0220231103033231-1302003210120210)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0111310011222331-1003110232112333-1111322131031333-1000030102132212-2323033002112101-0013022000230213-0133330230231230-2010232223022000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310202231222322-2021313223110011-0330022020333323-2020130212301302-2132320020113303-3123110012123301-0032033131201123-2121100200220222"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.is_primary — is_primary / 210121120001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.is_primary

<a id="canonical-0132030310212003-0013213130233020-3110100110110112-3032000322123001-1121333232222103-0220202311321000-3022212231103322-2233313301011322"></a>

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

<a id="canonical-0010232223011212-0120303033302233-2101030330123000-2222002110030301-0133120200112003-2233221120021322-0322101101101333-0010123210112200"></a>

## Direct properties — is_primary / 210121120001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122103321010010-3223020002201100-0203201011121302-2322123303213230-3233110331211133-1020210020303310-2330312122220023-1203332132322222"></a>

## Next pages — is_primary / 210121120001 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2200002212330130-0023011213233003-1032231011130223-1332003033332223-0202313002123033-2012123023320012-2213013330212203-0213131311223203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301020303212112-3020033312121121-1203233331323200-1123212001003323-1302230301020120-3003101030203222-1322231232310220-2222003111023320"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor — monitor / 130222011330 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor

<a id="canonical-3123011232201132-3220123200130101-0000230003321233-3223132110022330-0030101230213301-1223103203203203-3002131100230330-2221330120212333"></a>

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

<a id="canonical-2332212120332212-1020222133103022-3231102001001130-0303110010212001-3002020131302010-3020102303233101-2312233011033332-3222322203031222"></a>

## Direct properties — monitor / 130222011330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000012322203123-0011032012333003-0333011323202103-2031210021123030-2330212200200301-0023030033203321-3121022000210133-2100332302030223"></a>

## Next pages — monitor / 130222011330 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2320312302030103-1200223333300313-3333313233333112-1223223220212311-2323220313300023-2002111013130012-0202020232101202-3100120012120220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101232030023211-0130020103320333-1220112311101103-2201000202301101-2012312200313303-1123232132212212-0001313102320032-1330031102321223"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled — monitor_disabled / 302033113101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled

<a id="canonical-1230200323310130-3002233210120132-3030301033213011-2220233330100303-1300121203103211-3320030013031203-0113320033303121-2300303030101020"></a>

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

<a id="canonical-3033210231030001-1202220332302302-1200221002003333-3202213120102311-2110010102310303-2212123222133012-0033113203113321-2331313212131220"></a>

## Direct properties — monitor_disabled / 302033113101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131323010132030-2302130130213031-1302123032032221-1123021212213123-3131123000130230-2312233022202121-0110333203333103-0110311012001321"></a>

## Next pages — monitor_disabled / 302033113101 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3113113332212111-1131332022103233-0332230011303101-1022222201222013-0002200113013330-3201011131003300-1222201033030030-3313333112311003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022120120003003-2312012022102121-0103012032233113-2223023323201231-2133001203003231-1321200131020231-1022012132030002-3231222021301201"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address — no_ipv6_address / 131011221030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address

<a id="canonical-2121021222102203-3031011133102233-2133100300101033-0020032310203310-2011211131200133-0003130321220223-3113333311303021-1233202100101030"></a>

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
no_ipv6_address = {}
```

<a id="canonical-2000122232211020-3202020022033031-1231122032102031-0202302230221322-3323230302111033-3132230210030113-2301313221323300-0322320303013320"></a>

## Direct properties — no_ipv6_address / 131011221030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120210003221111-2123033330011123-2110222010321111-3002221133010322-0010111100121112-1111121223111332-1033301010331333-2010232300022233"></a>

## Next pages — no_ipv6_address / 131011221030 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2012213203133131-3033132012322120-2332130232300221-2200022322011332-2212023330222110-3300321023230131-0330031310321100-1033330321001211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320012011131131-2103021030011301-2123031002002200-3310003112333030-2312203133132023-1223310300321233-1302013231220301-2122323112203331"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.not_primary — not_primary / 013301200223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.not_primary

<a id="canonical-2121211103023301-3312102213312211-1212032022021200-1332223300022331-0321003003301212-3213213111211200-3020032313330111-2212221302011032"></a>

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

<a id="canonical-3230032021123310-0330002203333220-0312021321332323-3313133332101320-0033200131113032-0200000300320303-3013320033310332-1133233322131122"></a>

## Direct properties — not_primary / 013301200223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110312020133222-1111301130333003-3322311302032332-3001131020330321-3211023101112111-2311303212032322-3202222003313033-3313220033313323"></a>

## Next pages — not_primary / 013301200223 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1001213003222012-1032203333022030-2121223132313333-3221303033102330-0333233112231113-3020110023133301-1013300121031211-0302032201233031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021203212211332-2000230003232021-2203331122232223-2232222233322011-3101300320001213-1010002133321111-3001320010201311-3033321100212021"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network — site_local_inside_network / 211311102222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network

<a id="canonical-1102101222210233-0301232333312122-1002001211212120-2221002010110103-2310110220130212-3021323030021323-3301133000011230-0232120011130231"></a>

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
site_local_inside_network = {}
```

<a id="canonical-1000031221102103-0021130021120023-0103223001202031-0231032222131132-3230233201012311-2130102021103302-3110232313322302-1333313123031220"></a>

## Direct properties — site_local_inside_network / 211311102222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221211021101223-0122132222123002-2330120031000323-1010013002322022-3103122212011303-2332031203200102-3023123203200200-0302210321113002"></a>

## Next pages — site_local_inside_network / 211311102222 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3013032210221100-1031000310320031-1113200103113322-1213221221021203-0033112130020223-3231113110023110-0311233222331220-0012032030203331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013331032020300-0222301013230330-1232113323121330-2200302013323232-0321030310010033-2130210131200232-0120303302130320-2112231012330332"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network — site_local_network / 311231111210 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network

<a id="canonical-0010120311033202-1330222033230030-3312001302322233-0130120321232112-0320202333000013-2222033331022320-2332023110023000-1312320311111121"></a>

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
site_local_network = {}
```

<a id="canonical-2312023230202022-1210102312310320-2132312213133032-2320100210231112-2302222030222032-1033130032112320-2033122010301331-2213321300102032"></a>

## Direct properties — site_local_network / 311231111210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120112323022110-0132003032120301-2001101201020022-3332232130322133-0021331012000122-2022002311033332-2123232322323210-0031110010020313"></a>

## Next pages — site_local_network / 311231111210 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3003122101211022-3002131210021013-3230110200322123-1212200332100101-2303033200310002-0301223131030300-3012213210322112-1300111321131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331311000232111-3213120122021232-1202111133303103-1321323231101012-3222021110333312-1212313003301010-3221232302122133-1011202331013213"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip — static_ip / 231333021213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="canonical-0200301022222132-2311320013212130-3312231222020113-0011302311230220-3220223121331123-3132011101133002-3120231110313032-3222112013101203"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122201213212322-0021210203111111-0233110132010213-1333133033113313-2323333133031212-3323102002212303-0332203010031013-2231102320031213"></a>

## Direct properties — static_ip / 231333021213 / 3

- [cluster_static_ip](resources--voltstack_site--reference--group-004.md#canonical-1010221012022122-0211022030210132-2231211013011001-2220313101013302-2333002003222300-2233020202211301-1131223112221123-3000131121233120): complete subsection reference.

- [node_static_ip](resources--voltstack_site--reference--group-004.md#canonical-3203331112332031-0133021222122322-2010110323030300-3121103013232232-2033211002031320-3133223001101213-1322021021213312-1122112223313132): complete subsection reference.

<a id="canonical-3131202012330120-2013232213032001-3332130300313122-1220113220000110-0030110123000021-2312002131012222-1030010033112003-2030301022323001"></a>

## Next pages — static_ip / 231333021213 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](resources--voltstack_site--reference--group-004.md#canonical-1010221012022122-0211022030210132-2231211013011001-2220313101013302-2333002003222300-2233020202211301-1131223112221123-3000131121233120)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](resources--voltstack_site--reference--group-004.md#canonical-3203331112332031-0133021222122322-2010110323030300-3121103013232232-2033211002031320-3133223001101213-1322021021213312-1122112223313132)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1010221012022122-0211022030210132-2231211013011001-2220313101013302-2333002003222300-2233020202211301-1131223112221123-3000131121233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031313122212331-2110112111112022-1231233103013202-2130222133212212-1322222301213200-3322320133333031-2123110213112301-1030220311133322"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip — cluster_static_ip / 232022333332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-3003122101211022-3002131210021013-3230110200322123-1212200332100101-2303033200310002-0301223131030300-3012213210322112-1300111321131313)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-1200021320023333-1132222120310223-2121302033321132-1021210112211133-1032122233332300-3300032101131201-1301031320000000-3332200113033332"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132302300100013-1111200111123122-3010121330200200-3113022103020230-0212203330300320-3110111101210200-1032310003313010-0310131333022112"></a>

## Direct properties — cluster_static_ip / 232022333332 / 3

<a id="canonical-0022333121332300-2223000002131112-0322311220200113-3013300122111233-1032100312230113-3211011132023301-2321321200313303-1210202020123012"></a>

<a id="canonical-0231320122200300-3323123032213210-2221200031223313-1302222223213332-2321303232032203-0303210120030320-3331001021013013-1230333210012212"></a>

## interface_ip_map property — cluster_static_ip / 232022333332 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-0033233103022230-2320131332130032-3113311232130311-2022032301113032-2021330122031011-1130322303203110-2122202012033103-3300301203322212"></a>

## Next pages — cluster_static_ip / 232022333332 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-3003122101211022-3002131210021013-3230110200322123-1212200332100101-2303033200310002-0301223131030300-3012213210322112-1300111321131313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3203331112332031-0133021222122322-2010110323030300-3121103013232232-2033211002031320-3133223001101213-1322021021213312-1122112223313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131112003012310-3132212113332211-1012333221130021-0331232101023320-3120233131300002-0133113132203332-2133320133022133-0100013201211223"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip — node_static_ip / 013012001233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-3003122101211022-3002131210021013-3230110200322123-1212200332100101-2303033200310002-0301223131030300-3012213210322112-1300111321131313)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip

<a id="canonical-3210122221311020-0331133333320022-3100323013130332-3122013131310310-0131100011122300-3201101200311111-2001230031313103-1003211202201322"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312213130320331-2100101232100113-0232113020210310-2100121331000330-1310122133030021-1030001302113221-1032222203031332-0300001020211113"></a>

## Direct properties — node_static_ip / 013012001233 / 3

<a id="canonical-1330120333312113-1013103313302233-0022303321000321-3300100103311130-1001211113002221-1200311230213212-3003032203313023-0120032032302012"></a>

<a id="canonical-0303223332131312-1112221010131322-0212013012013303-3023211033310332-1013330333203100-2030032023332220-1003301032030122-3032213222130000"></a>

## default_gw property — node_static_ip / 013012001233 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3002213303200110-2321232011012010-0110121130031001-0233020210131332-2202013012131211-0322333000003031-0110123322101202-0023313002130211"></a>

<a id="canonical-0213321010230113-2032102323223201-3211002310130013-2301013310113320-2311100013031331-0023311221102032-3111103030111311-0023010020021033"></a>

## dns_server property — node_static_ip / 013012001233 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3132223012301331-3111311330223001-0332223220100333-2232103120102023-1230303023320303-0111313322223313-1110331312322233-1301121322231232"></a>

<a id="canonical-2320030300032322-3022212103221021-0301223301102122-2233130120010031-0200330030311133-2211000111331200-0110030030201223-2031002010110333"></a>

## ip_address property — node_static_ip / 013012001233 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-3123020021032321-2110020000123001-1112021020001013-3003030330130132-2201210330101131-1131130113130212-0123220021301221-0002120030133323"></a>

## Next pages — node_static_ip / 013012001233 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-3003122101211022-3002131210021013-3230110200322123-1212200332100101-2303033200310002-0301223131030300-3012213210322112-1300111321131313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3131302103210300-3223212220131130-2131321000320031-3020301122230303-2032101033030110-1301133233113112-1112232030222121-2000010131031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001300020002113-1201023233030302-1221111323102122-1310302013113012-3100211223211022-1033330330103320-2232031120133333-3112323302331331"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address — static_ipv6_address / 231330311100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

<a id="canonical-2330220200202233-2302202303231003-1213031130130123-2333003123031330-3013000212031132-3300013000220300-1302032021322122-0131310112320230"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320301112012121-0132123121113100-2110312233210012-2010220103010010-0320312022031031-2113132110122302-1031210120031023-0233323333010100"></a>

## Direct properties — static_ipv6_address / 231330311100 / 3

- [cluster_static_ip](resources--voltstack_site--reference--group-004.md#canonical-0310311332303122-2031110022021121-2221313212000232-0221002201032321-2331013300203133-1133102012123021-2322323012102213-0121222131021211): complete subsection reference.

- [node_static_ip](resources--voltstack_site--reference--group-004.md#canonical-1211122032003303-1220010121220031-2321101121120222-0322330331321303-3133023312123133-1023131320221323-1231022300202212-2200013231222201): complete subsection reference.

<a id="canonical-2111321201203221-1331322302121210-3131313102021121-1022013002012210-1221303033110030-2023031132322032-0023210232201322-0203323021103332"></a>

## Next pages — static_ipv6_address / 231330311100 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--voltstack_site--reference--group-004.md#canonical-0310311332303122-2031110022021121-2221313212000232-0221002201032321-2331013300203133-1133102012123021-2322323012102213-0121222131021211)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--voltstack_site--reference--group-004.md#canonical-1211122032003303-1220010121220031-2321101121120222-0322330331321303-3133023312123133-1023131320221323-1231022300202212-2200013231222201)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0310311332303122-2031110022021121-2221313212000232-0221002201032321-2331013300203133-1133102012123021-2322323012102213-0121222131021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113231331302102-3210123100112232-2100221303311333-3323232302033021-0222212112331111-3032100221210112-1110120303232203-3033211203110032"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip — cluster_static_ip / 212132231021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--voltstack_site--reference--group-004.md#canonical-3131302103210300-3223212220131130-2131321000320031-3020301122230303-2032101033030110-1301133233113112-1112232030222121-2000010131031113)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-1103121010201212-1013303301213131-3003002122200122-3001023212331311-0022220013222231-0002202313212313-0232020313320232-1033232133131200"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032031123023302-3312112011201113-2121031312033120-3313223032031331-3201302203020323-3032001021133020-3030000131231223-2120112020121331"></a>

## Direct properties — cluster_static_ip / 212132231021 / 3

<a id="canonical-3103201102121333-3232121301203033-1321021311003011-1332312320103332-3223003000320220-0011302020312222-3121212310211330-0132113010231311"></a>

<a id="canonical-0131123001130310-2020312301212302-3311113233131132-0112220000311030-2001020301020223-1113212203030321-1000311211101110-3220230213222332"></a>

## interface_ip_map property — cluster_static_ip / 212132231021 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-0131231130303320-0000310222312111-1200200330101032-2320002001033331-0012022020303100-2010222132232332-0100110113213100-3031322212001110"></a>

## Next pages — cluster_static_ip / 212132231021 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--voltstack_site--reference--group-004.md#canonical-3131302103210300-3223212220131130-2131321000320031-3020301122230303-2032101033030110-1301133233113112-1112232030222121-2000010131031113)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1211122032003303-1220010121220031-2321101121120222-0322330331321303-3133023312123133-1023131320221323-1231022300202212-2200013231222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230303112321300-1303132110230101-1322102120133332-1101023002103302-3223223203311101-3031112200320102-0113122113231130-0210200103131211"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip — node_static_ip / 321333122323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--voltstack_site--reference--group-004.md#canonical-3131302103210300-3223212220131130-2131321000320031-3020301122230303-2032101033030110-1301133233113112-1112232030222121-2000010131031113)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-2123112112013122-2012333111001031-3122202310131232-2222132312210312-3201020100223133-2113213310220213-2120113301323102-0213312312120103"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033312212110322-3221311201113031-3311310233103332-3323221230032212-3220222332320213-1223131001321221-0331102210101201-2203312031312210"></a>

## Direct properties — node_static_ip / 321333122323 / 3

<a id="canonical-2310030332130130-3113123332001323-1200312220031022-3123101302232022-1223001211331233-3023013112121331-2321010321300211-2003320100001220"></a>

<a id="canonical-0202312102013312-1303200313300233-2020300032102030-2033231312330212-2213202132032012-1321010010313002-2223223300112221-0011323222110113"></a>

## default_gw property — node_static_ip / 321333122323 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3011133331230113-3300330321121331-1021101101112113-1232212311132321-2221303200331311-0032002023323313-3131133201033033-2330212110021120"></a>

<a id="canonical-2222110021101033-3020121000212301-0323103002210133-2331322023033312-3323331231013200-0321200333321233-2002120020332032-3232230312210213"></a>

## dns_server property — node_static_ip / 321333122323 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0201020110332230-2130012103300202-0203133233333012-1113101020123122-0330130232013223-3011313313210123-0203112103232213-1333110031033002"></a>

<a id="canonical-2013331033322201-3213131331010120-0233311002203100-1000230212221223-2200022222111213-1032130221311001-1112003312002103-2233233020200113"></a>

## ip_address property — node_static_ip / 321333122323 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-2133233103121113-2213302310021212-1033022100103200-3213110232012310-2320322103223130-1011300103331133-0030333102103220-2302303030000200"></a>

## Next pages — node_static_ip / 321333122323 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--voltstack_site--reference--group-004.md#canonical-3131302103210300-3223212220131130-2131321000320031-3020301122230303-2032101033030110-1301133233113112-1112232030222121-2000010131031113)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0331003030121011-2323223300133020-3021101100023000-3101023033233233-1220023302312020-1132210330320101-2023300110032112-0102321032213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211113031330111-3021332023133310-0321122331103321-2020233101213333-1020302233002222-0010332110120332-3130230201010221-2032232111100003"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.storage_network — storage_network / 111200103221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.storage_network

<a id="canonical-3003211020331120-1202030213331131-1311022030213333-0013223013110031-1301112011123213-2010112130212131-0133333122013033-2331202132200000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for storage network.

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
storage_network = {}
```

<a id="canonical-3121013302332013-0102231301333030-0200322323321003-0202021021110313-2000333023201330-1013302331210130-0021310313111032-3303321123033202"></a>

## Direct properties — storage_network / 111200103221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113123232300233-2203103033231101-2021020312102301-1123132221103132-0122211321231002-2130010023121023-0321312013231221-1123103301301110"></a>

## Next pages — storage_network / 111200103221 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0233023300003331-1230012301330210-2012233013120233-2213302310231030-0120213132103130-0213021330331222-0301110013021133-1311211323112113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013231310012302-2133330102310202-1320013133000201-2022321111021323-3221223232323212-1331323203332203-2133311323232122-0013233123300130"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.untagged — untagged / 002001002020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- custom_network_config.interface_list.interfaces.ethernet_interface.untagged

<a id="canonical-2221021003112302-1113120332132121-0231221221110313-0200301112311103-1231000312111010-2031111312332322-1102101210110230-1310020111132033"></a>

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
untagged = {}
```

<a id="canonical-3213032230320302-2020123310232001-2023230000131212-2102301221211220-2220113233233211-3323002311031131-1000332302312233-2032311011020003"></a>

## Direct properties — untagged / 002001002020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322020020031323-2221210110232002-2300132201231110-1231100012122222-3131312201013033-3123332031023111-0103302100330031-2221332000103223"></a>

## Next pages — untagged / 002001002020 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--reference--group-003.md#canonical-3311111210321031-0311302100211031-0230203130311233-1002132031121013-2012301010333203-1321300101222010-1102233121233000-1033021313320201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3013210122123320-2001013331000233-2200032121012131-0000123020201321-2110212323312113-2203110103101133-1333313000202120-2003202312221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023002113312333-2213302000110203-1100330031203302-1013313010232103-2012233113100203-3303002323231131-2232310301231121-1232201021302223"></a>

## custom_network_config.interface_list.interfaces.labels — labels / 100113113022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- custom_network_config.interface_list.interfaces.labels

<a id="canonical-0321320111330210-1303311020210121-3100312001200210-1133330200322032-0201133212303212-0333013022320112-1023313011000233-0212210012303103"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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
labels {}
```

<a id="canonical-1003013112000110-0023333132131333-1311311320233311-2132320231221202-3220132332320303-1333322203331101-3221033122012311-3000330001103222"></a>

## Direct properties — labels / 100113113022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220321220111230-3222112232301113-0111113011031103-3121032230110212-2201233303332201-1321330032110313-2001202301311130-0011221113110303"></a>

## Next pages — labels / 100113113022 / 4

- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033333001200120-2202310003032112-0112231303033212-0132003330133103-0331002030203112-3313220301011211-3303022020122322-3100322010313333"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface — tunnel_interface / 323030220221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- custom_network_config.interface_list.interfaces.tunnel_interface

<a id="canonical-3201230022102220-2312100200001230-0233021332110101-0231313333101331-2113033121302122-3010303231123303-2313023223322323-2110003101032011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

Terraform syntax:

```terraform
tunnel_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121201331202100-2233300222030132-2022103323331000-0000100032011001-3013210221222010-3123132031323200-0023022030102121-1213313201311031"></a>

## Direct properties — tunnel_interface / 323030220221 / 3

<a id="canonical-0303013331000102-0130030013020120-1121200310332310-2112023322201123-2022003122232323-0022101110332213-1212113223110002-2231310200121101"></a>

<a id="canonical-2110321231103012-3022023331223100-0203132212222130-3321201102003012-3211220131113133-2131131122202232-1220113212111212-1300320012220110"></a>

## mtu property — tunnel_interface / 323030220221 / 4

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

<a id="canonical-0102200032120302-3022013111030113-1012231132121210-1100210103001212-2213212333130231-3331121112331301-0230233030300002-0203012203032332"></a>

<a id="canonical-0103300333321102-2131311232231111-1232131312230032-3200221320103101-1222202033032031-0113013333113100-3221030301033010-0333022020201202"></a>

## node property — tunnel_interface / 323030220221 / 5

Type: `"string"`. Optional.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

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

<a id="canonical-0211020121313122-3110113223211001-0323300210332211-0120210231233012-2032131103123011-2322003120212020-2221221223200200-1230120020200101"></a>

<a id="canonical-1310201210211023-2303223013123103-3002300113102011-1132212103100213-2313112223113200-3113020000222111-3320101032022020-3133102232000111"></a>

## priority property — tunnel_interface / 323030220221 / 6

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

- [site_local_inside_network](resources--voltstack_site--reference--group-004.md#canonical-1202311203203210-3222212223220331-0332233310332222-0223321123321022-0022201001230130-1200331323011330-0231330212120301-3213220120101310): complete subsection reference.

- [site_local_network](resources--voltstack_site--reference--group-004.md#canonical-3223132132112203-2101021030330202-2120033120303233-1202210011222201-1102030003321330-1123203111233021-0201330131011332-2111222332131100): complete subsection reference.

- [static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012): complete subsection reference.

- [tunnel](resources--voltstack_site--reference--group-005.md#canonical-0220322301001222-3022112332322211-2232333231020032-0201320212231012-3331231230023133-1032013112021312-1002130020322021-3220133323211300): complete subsection reference.

<a id="canonical-3312321303230103-1321002232303032-1012102201013003-3113321012100232-0010230321312312-3011113310121133-3230200312000223-3010200230213010"></a>

## Next pages — tunnel_interface / 323030220221 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network](resources--voltstack_site--reference--group-004.md#canonical-1202311203203210-3222212223220331-0332233310332222-0223321123321022-0022201001230130-1200331323011330-0231330212120301-3213220120101310)
- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network](resources--voltstack_site--reference--group-004.md#canonical-3223132132112203-2101021030330202-2120033120303233-1202210011222201-1102030003321330-1123203111233021-0201330131011332-2111222332131100)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012)
- [custom_network_config.interface_list.interfaces.tunnel_interface.tunnel](resources--voltstack_site--reference--group-005.md#canonical-0220322301001222-3022112332322211-2232333231020032-0201320212231012-3331231230023133-1032013112021312-1002130020322021-3220133323211300)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1202311203203210-3222212223220331-0332233310332222-0223321123321022-0022201001230130-1200331323011330-0231330212120301-3213220120101310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322222130220321-1123210332102221-0213220000330332-0113111300101123-2023111030211203-2000031123302230-2123221330000213-3310313223221300"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network — site_local_inside_network / 120130100102 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network

<a id="canonical-1220330100322303-3031310012120011-0030121312110323-3023032203033020-1023121001132002-0133231012300010-0333112210211133-2202231113211220"></a>

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
site_local_inside_network = {}
```

<a id="canonical-2113233221002111-2213221122123231-1130002302331113-3333110220212100-0100121210210003-3112220302233033-3201200012332310-3200020133223003"></a>

## Direct properties — site_local_inside_network / 120130100102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311323022213130-0312323121112233-0333312133223131-3323303330331030-3031332313022023-3322010302232222-2033030112302321-3001132130102332"></a>

## Next pages — site_local_inside_network / 120130100102 / 4

- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3223132132112203-2101021030330202-2120033120303233-1202210011222201-1102030003321330-1123203111233021-0201330131011332-2111222332131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321102322020000-1122032221032211-3212012002130022-3323133122001303-1303103130333201-1303132130330302-2013211211232223-2001120300003013"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network — site_local_network / 032331001133 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network

<a id="canonical-1310310232132222-2213313120212031-0221213300102003-0303110212320100-0000123031023302-3101101310003233-0021213233210103-0333101033110133"></a>

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
site_local_network = {}
```

<a id="canonical-3331320000233001-2020111331131300-0110203100031201-3230032013312021-1312313022310002-0101122220110113-0031220332323311-3112013310222322"></a>

## Direct properties — site_local_network / 032331001133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223301333223013-3100011223223023-0231000131033232-2333321131100022-0123311313232303-0232130331231123-2322121223013033-3122300033302011"></a>

## Next pages — site_local_network / 032331001133 / 4

- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132021100300101-3020222003230012-1232230323222221-1222230011301222-1233223013201002-2331202311012121-3300330120330002-3232100101122201"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip — static_ip / 133121302101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip

<a id="canonical-1231130321311113-0230201122011103-3332302223212201-3002233210032010-0300233030310000-0300320100021321-2301113312101120-1112313113323322"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```
