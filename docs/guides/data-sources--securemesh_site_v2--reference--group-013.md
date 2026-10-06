---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3102221300010003-2222001231132221-2330100010313201-0130131131132320-1130012211222030-1300303312223031-1122200031300220-0030032113012122"></a>

## `nutanix.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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

- [monitor](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1220032203233122-3000212111302131-0230322231210211-3020210030030131-1000102230112333-1122111322013123-2211001223013202-3003201323213133): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2222301013131301-2321201000103230-0301002131232033-3123101310023021-2322333331323021-3303211321101132-0212201100003121-1121223011310012): complete subsection reference.

<a id="canonical-3033211100021323-3333131002202002-0203122131000103-2113230321212033-0230310302221333-2302200201122020-3210310333100112-0020220231133321"></a>

<a id="canonical-0023010230132110-3213222331302032-2011310233101302-2131320310231211-2332002131200033-2130300223323220-1330330313022031-1130302033031100"></a>

## `nutanix.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3001321302202231-2200301322202302-3132223212133301-0221100230133100-1110312301011122-2112123223221000-1020332331030113-2312012112320200"></a>

<a id="canonical-3131330020230213-0213131013100303-1123133212022321-2110102202320202-2333303202233131-1123101131013032-1122201301022312-0002133022303231"></a>

## `nutanix.not_managed.node_list.interface_list.name` property

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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

- [network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2110120310132301-0333323103202113-0203001122333010-3120002212310310-3301321201101322-1330321201202130-3301023300132031-2320000213313113): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1201132021103012-0320121210130102-0320131023123203-1223103021112230-1212210112010230-0223322113122123-2230110002200012-1133330031212103): complete subsection reference.

<a id="canonical-0213310131211130-1230113213301223-2033231221021133-1211033001321201-1332331031120011-3222110300000233-2121333010322022-1113232303013021"></a>

<a id="canonical-1212302230300322-2000202102032322-1233300003310230-0200232332013000-0311113020323023-1221301103212133-2322311320223222-0213030312033312"></a>

## `nutanix.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3321103030123112-1111002332222312-0001330332230202-1210023031020311-3012220032001012-3202211200231032-3201333232122130-0012331300100203): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122312333000212-1020203202012230-2222133022312222-1000313330102330-1210033002310300-0122320110232122-1002322113323212-1220111120020002): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3121320200010110-3202032231021203-1123033202300131-0321201113233330-0000003322120203-2310101203203111-1233332211300133-0131130000003112): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2212102230003211-3201300121123111-2100131201300322-1130003101002213-1130032130203003-3233330110331212-3022101331031233-2322323020313313): complete subsection reference.

<a id="canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2221102133032101-1123113121032131-3202321333313030-2030210222201222-2201033031131000-0313233223320113-0312321302130232-1233211303221130"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-1332320022323201-1212101301012332-2210200102210322-2211121223012020-0112203302103020-0302021333223310-1102130000021222-1000230113010322"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1031000000211333-1221332320030212-3020123121031020-2133302033310323-0113023011212323-1230220022020002-3112301203121030-0010101023221130): complete subsection reference.

<a id="canonical-0330001021300103-2202231231313130-0033310310130022-2230032011133031-3313223233301021-0130102013112211-2003031131131130-1033011021223223"></a>

<a id="canonical-2120110132013001-0010120023111331-0301013211131232-1330130130323301-1033033322330111-0120330221213330-3220203322031220-1212320020211313"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1300003021233213-1303201221113122-3021131012203213-2130313303131202-0133130100123332-3201113233020330-2232003032210121-1320120302330111): complete subsection reference.

<a id="canonical-3210223010231313-3203232221001132-0023020221100111-1130221302130211-3201322033113130-0102122001302101-3212101110210111-2321322130200303"></a>

<a id="canonical-2221030201030330-2210013220121232-3002320310303322-2133031131201111-2110223131103002-1201330011101102-3112033321011203-0133122303220133"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-0203220311321102-2020033221031123-3223220103300030-3212033132311312-3010022222000130-1212211013302100-0102023213102000-3202002231302302"></a>

<a id="canonical-2023000001333033-0131321233033120-2302200122131233-3101303323023312-3100020332230121-3303210133013103-1201220022101110-1201000201001103"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-2223003102032010-2233212302113300-2211230103120033-3131022200313331-2120130221003230-0133131103013221-3312332100223330-1021230111201211"></a>

<a id="canonical-3200330231302223-1011220103312312-3201101213322103-3321022333032101-0232321122023233-2012130111213302-0222001113101130-2002312003001303"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-1031000000211333-1221332320030212-3020123121031020-2133302033310323-0113023011212323-1230220022020002-3112301203121030-0010101023221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- nutanix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1001013221000312-1030023022122211-2000103020012312-3030321333303031-1000223321031203-3102222132002332-3122110102011310-3122103113010303"></a>

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

<a id="canonical-1300003021233213-1303201221113122-3021131012203213-2130313303131202-0133130100123332-3201113233020330-2232003032210121-1320120302330111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- nutanix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0112231023132002-1300000222322121-2032210203022220-2132222013203203-3221013121222300-3103332221221013-1032333230230213-2011213201231320"></a>

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

<a id="canonical-1221013121330033-3033220311210020-2320010200320311-3332120000222023-1320230322112303-0213033103211313-1111013111331100-3320202002302311"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-3113203022213021-3122131002123322-2103102332202101-0000002132002312-0223121022030033-2212230021200012-0213313233033000-0233101203210220"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-3103231010110021-2023313200103100-1120310120232031-3322130103032011-2120210133322033-0230203202133330-0003331122120101-3020301001222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3232300113010213-3231101023310320-1130331002223112-1103011303200312-2130011033100023-0233202132112112-1312033122320212-1221021102122233"></a>

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

<a id="canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2102032022011120-2223210330220312-2130032102113030-3203123322123001-2130231301120122-2113012022002220-2202303130002112-0322120022110103"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-0123310230212330-0213330113130121-0312122303321020-0031230103023002-1312032231311123-1213221013322003-1003312021010033-0110100030101331"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0130230232002332-3103200103322132-0202030322021033-1212110132333311-0212210211202002-1210303130100231-0003101100232312-2131203110203303): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1212222023033302-0131021331120213-2313231302013010-3132003212103333-3122200233201232-2202011000031012-3202313133302320-1320012313103012): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132): complete subsection reference.

<a id="canonical-0132022212021123-1330130332000211-1301101223330113-3002301103030031-0132122331023001-0203222010321213-0130010121213133-3133321110011102"></a>

<a id="canonical-2101312321131332-2322113223110121-0322013330223230-3333011132101333-3221121030003303-2030333020111010-0103200301103233-2033021110111131"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1132200103101132-0113323130103000-2331003311033322-0233302300013130-1213003111223033-2303303210312010-2313123023303013-3011321203200121"></a>

<a id="canonical-3111132022003202-1233013221330310-3223023333003203-0112100000021233-2323102311310021-1002000233232121-2312232121201233-1130000031330123"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2002323301310213-1213012000313020-2333213003320003-2331030031310311-2233212333111220-1001212101011123-3102030011303221-0130012310003212): complete subsection reference.

<a id="canonical-0130230232002332-3103200103322132-0202030322021033-1212110132333311-0212210211202002-1210303130100231-0003101100232312-2131203110203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0021300122031002-1123001313112223-3030301220231212-3313220313002233-3233002311321013-1121301113103320-3031132020003200-2102020210232331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-1212222023033302-0131021331120213-2313231302013010-3132003212103333-3122200233201232-2202011000031012-3202313133302320-1320012313103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3332310111131201-1221132223220212-2231022111300121-0322032211310321-1110332132132121-2133101322302112-2012302100210000-2231011021310210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3221321122033231-1232332333110201-2201210323310302-3320300020332312-3300232123132030-1132211022332302-0313011102310102-2001113212321321"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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

<a id="canonical-3011300002220323-3111310123023310-0030232133302333-3200103200013300-3333022120000210-0030012232321323-2200311030013332-0023011301011132"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3301110231321221-2213100320021122-2022002212302323-1023031011333323-1102032112023323-3003021323011310-2030022121023131-1211332300112002"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0212121221222302-1020031033203022-1131010200121232-1300222301222233-0010012103133220-0230110022220023-0022003231103012-1330211133313033"></a>

<a id="canonical-1013021220030200-1201030233232102-1222113130220223-2222301321021120-1211003031103123-3302120010132010-3212232203222012-3312002312102113"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1120011021231312-3232333011210302-3202313310121201-3213202303001230-2220113310323332-3012323303320200-0131132303203233-2332202311231212): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3023232223101013-2203301220230331-3023110200233002-1013000300300313-1110110333133121-3220233122302321-3010133000012321-0303112203123232): complete subsection reference.

<a id="canonical-2003020111312230-3223101312122232-1011010213021030-3311221211033011-2103130131101213-3202303013312200-3203011300003001-0133301130320223"></a>

<a id="canonical-0200321223321200-0031011202233001-2323002202013223-3113030002021010-0012211120303122-1311001233332313-1300311212010332-3313122323002030"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3023020131100201-3212011332322002-0103103331022322-0032132312101100-3322121220002030-1200210113232030-0002212232021030-0302303031002102"></a>

<a id="canonical-1333300311001021-1222130000222033-2013232303110222-1003330211133022-0020220313233222-3331122323021220-3000200201133013-0331011102300230"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

- [pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123201321301201-3112230103221232-1023001212201322-0201333330330013-3300332312002102-1133220123331133-1021312221033002-2103011020213102): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1302102333231303-2222300123110202-0100112000321030-3330312132122011-2213323221002200-0121010331113323-1210120313311211-1130323230333233): complete subsection reference.

<a id="canonical-1120011021231312-3232333011210302-3202313310121201-3213202303001230-2220113310323332-3012323303320200-0131132303203233-2332202311231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1312321333200123-2311130233023310-3000010321311002-1230112003302032-0022010203111103-1311321122130231-1321103021312110-1011231013110211"></a>

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

<a id="canonical-3023232223101013-2203301220230331-3023110200233002-1013000300300313-1110110333133121-3220233122302321-3010133000012321-0303112203123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1121110330103133-1333310113133101-0320120222113313-2033032331223312-0001132113010003-0103003232121333-3200023122110101-2130221022111322"></a>

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

<a id="canonical-1123201321301201-3112230103221232-1023001212201322-0201333330330013-3300332312002102-1133220123331133-1021312221033002-2103011020213102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2130022320201312-1300232122112021-3211030300110231-0120231102212130-1201021312332010-1110112211201302-2230210002331232-2331332002212312"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3003110323032321-1221031313311201-1322331203110221-2103210003232011-1013231303201030-2131123122132123-2200203121313221-1022012023011122"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3301123201000033-1102213201101332-3021301220230203-1320112102113232-2023213210133233-3132123202200302-0202112000230330-3121223003023133"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0020213213310220-1221333202233311-3001330211111131-1320301133300101-1111122331131232-1300000221120010-0230213223102030-2310032300331001"></a>

<a id="canonical-2220031231202110-2301103101020112-0223130122011020-0323300232332213-1100001231011321-0001023213120221-1211121333100201-2310211301201323"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-2312210020231220-3331002331303303-1213320331233122-1103200220330122-2311032310101020-1111010321203033-2202121301220232-1001012111021112"></a>

<a id="canonical-0110123203021003-1221221101311001-3331122210222130-0301000000331322-0032332003201122-1120200320221112-0221132202102123-3022023002310102"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1302102333231303-2222300123110202-0100112000321030-3330312132122011-2213323221002200-0121010331113323-1210120313311211-1130323230333233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3120231213032121-1222120120332023-2123121021111001-1001303022222221-2000303000212301-0202121013211222-1331300320210101-3300322323211010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-2002323301310213-1213012000313020-2333213003320003-2331030031310311-2233212333111220-1001212101011123-3102030011303221-0130012310003212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0003131210203313-3103011320321233-2112103002002020-3113201133210001-2031213031003003-3310313110222311-0231202320111133-3223023233210222"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

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

<a id="canonical-2230232302300211-3021212200332021-3230013003312031-3301233000110322-0033030220321333-3333302010321322-3113001113230032-3223213003100000"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-2233120221233032-3202023002030332-1322021331031210-3103312301331310-1221201103111223-2001132113100202-3101311312031310-0210003203111210"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

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

<a id="canonical-0222200321012331-3320100132003121-1232322313113033-3320300022010221-2033103203122030-2301022032313011-1130313202022011-1320012112122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3010223031230333-0330003011110110-0100320311310000-3322331123130322-1322320132201232-1301022221021311-3112313003231100-2000103223231101"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-2311100233332002-1331101213023301-0111201300312002-2201031211323021-0300003202031203-1033222132131330-2330203023010100-0023311210131133"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-0212003213232120-0132300121133322-2122200100312302-1131101221113202-1101000132232303-2312003030220031-3321022011032010-3100312201023222"></a>

#### `nutanix.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2123312312221011-1122111113010203-3030331003020201-0331022230033000-3031030321303111-2310031100031211-1033221002101232-2232032011320113"></a>

<a id="canonical-3300312331213310-3323132200223302-0100010300331203-2101012111021300-2002320210120031-0123001133331330-1031321221123321-2101312223310021"></a>

#### `nutanix.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0023301133121330-2111033330310133-3123332321030300-2001212223301330-0211120332302200-2330123100201020-0333232202213212-2121212210031111"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

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

<a id="canonical-2011312222313303-1132322131031112-3203230121103102-1202111312023312-0001032213321132-1221202101300011-2320110222301231-2100333123303012"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1132220130211201-2020203133223032-3013010011030101-0202110133313030-1331123321012201-3102222002330021-3310010010013021-3312200131030303): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313): complete subsection reference.

<a id="canonical-1132220130211201-2020203133223032-3013010011030101-0202110133313030-1331123321012201-3102222002330021-3310010010013021-3312200131030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1311100130113202-3200310100311201-1212021033032012-2123310320210311-2122232320211320-2120120230120110-3100131211030212-3303311033203301"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

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

<a id="canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2212211003120200-3020130332203221-1301000011011210-3031120033302321-2010030303110111-3103301100033321-0111312220220030-2300132123312203"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

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

<a id="canonical-1032132213022033-0111303031230101-3320010121311200-1110111103221100-2030121200011011-3032002202212332-0223223211101310-2102210122010103"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320): complete subsection reference.

<a id="canonical-3033011033021201-1221301211022302-3031321301131331-3103011011111130-1113221321120333-0210110103221302-0232111302133123-3120221102221232"></a>

<a id="canonical-3222103230100322-0212221110310122-2000221300212110-2223212223303333-1022122130301101-1102301322310301-3201203200022100-2113131312113031"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001): complete subsection reference.

<a id="canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1002313111132220-0031231321221333-0121002310301120-3002131230101323-2013203201211311-2121301102200023-3322213313010323-2330231033313312"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

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

<a id="canonical-2033100201023011-1303303232101102-3121200211100001-2103330220112201-2323201022131103-1321121122313123-1203133122123322-1033321232101000"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1202330002220320-1011111322231001-0020103300001223-3103131131100213-3323201210321030-3121113112221032-1300202303012333-2013211101311130): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231): complete subsection reference.

<a id="canonical-1202330002220320-1011111322231001-0020103300001223-3103131131100213-3323201210321030-3121113112221032-1300202303012333-2013211101311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1010002132201033-0101031032313201-1330023031011233-3233011211002130-1310020332132203-2310223030021022-0013033002311323-1020302030032230"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-3202333313213300-3210310101102112-3323303021302120-1323111003310131-0133111222120300-2322123032212110-1120223211322021-2121110103312333"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1223333001131003-1113000021103312-1312120311211200-1310310201022310-0021112321021123-2001130213100321-1200210332213222-2202300233003132"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

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

<a id="canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2321201302211132-0310233123203322-2031323122100301-0012100130103313-2100322203311030-2310331232023132-3010303031222310-2120031023303311"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

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

<a id="canonical-0102203122302313-2232103001121131-2030223112232321-2210331133112221-1033331012023121-1203213233321321-0102110123311212-2132110301032102"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2211022103311320-1021330121130211-0123102101331200-2000010113133003-2231201332111132-1310123200120320-3332213302132032-2013320320212302"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123102330302122-0131231032103300-0030202311300000-2113323020212023-3003002022033121-2200230123033301-2123003213030321-1001311102211110): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3323103100033201-0333110223013113-2203322312231021-2223203223000202-2130233320213111-1312233022112121-1312303221223023-3033123013011001): complete subsection reference.

<a id="canonical-1123102330302122-0131231032103300-0030202311300000-2113323020212023-3003002022033121-2200230123033301-2123003213030321-1001311102211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1311220312222333-0313221023311133-2130230102330232-0313010000312200-2133110321221203-0210330311010231-0222311230131230-1300233001303101"></a>

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

<a id="canonical-3323103100033201-0333110223013113-2203322312231021-2223203223000202-2130233320213111-1312233022112121-1312303221223023-3033123013011001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2333323312323333-3203010223223133-3322003203021020-0333023031303310-3112213301202030-0121313303022031-2021113321013020-1223020302203113"></a>

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

<a id="canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-1120132332102102-3310023223112003-3120110030002010-0333301231203012-3002131311003332-1000221103123210-0321233020201133-1202133122220313"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

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

<a id="canonical-3232023032313121-2111321023001221-1301200320030012-0323313202232031-3330333131331133-1300232333000113-0301022130211002-1103022110032300"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2312222003321111-2311112122300023-0010232033201133-3010023123231220-0032002321233020-3302002013022022-2331303322111310-2032213010300302): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2102232231102122-1023202223033313-0133202203101110-1201011032221322-3123221103120112-0323101123310210-2022223031222302-3123132121110331): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0301110131002322-1010131113300020-1331113322022302-3330022311233021-0031110001032300-2003120322131112-3022031112023013-1132123132010312): complete subsection reference.

<a id="canonical-2300030010221332-3102323231032032-3103010132300013-2121221120213101-1002002331212310-2102210233302310-2120012121330132-1331313200010303"></a>

<a id="canonical-1020211321102213-0333032101203102-1231333303203231-1131211131110121-3131301131200102-3011020023200220-2120130310020133-0013211000220032"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020320223103322-0210012033031112-0002202110033112-0120230101101202-0300310301311203-1212122003021003-0121203332312023-3203231102133010): complete subsection reference.

<a id="canonical-2312222003321111-2311112122300023-0010232033201133-3010023123231220-0032002321233020-3302002013022022-2331303322111310-2032213010300302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2233100212323311-2001233030212232-3130212321100132-2312103300212023-1002011132330021-1021311022200330-3122203333322331-2033131022120021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-2102232231102122-1023202223033313-0133202203101110-1201011032221322-3123221103120112-0323101123310210-2022223031222302-3123132121110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2012212223012312-0333131003221211-2122212122132331-2011000220002223-1201221103031001-3130200301302101-1312223131132232-2032032022012230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-0301110131002322-1010131113300020-1331113322022302-3330022311233021-0031110001032300-2003120322131112-3022031112023013-1132123132010312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0021000330311200-2033013033123012-0110213223101022-2223320101201231-2113323103030103-3020123123112131-0121002031022023-3210012132303222"></a>

Type: `"list"`. Computed.

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

<a id="canonical-0302233023100130-0322123011022111-2001032001130333-1231033303030332-2111312031010011-0201131202330333-3013101021020002-0133300331002321"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-1101002022221033-2101300133022111-0312301001213332-3223212313322103-0010110121211020-0033212003130032-0003323211311313-3330003303000012"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-2011001223023010-2112202020032300-0323103211112220-1122120321320022-2003113102210210-2001011213003312-3020122101011320-3230230312011033"></a>

<a id="canonical-1211130302132233-1301122130002331-3212032211000323-1112233310123203-3022130222122211-0021331332231030-1130333221032211-0011031332132103"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

- [pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1120322312302312-2003112120222013-0200112201100202-3112302012213000-0111022010303210-1003221102111132-2003101031220312-3031022202023012): complete subsection reference.

<a id="canonical-1120322312302312-2003112120222013-0200112201100202-3112302012213000-0111022010303210-1003221102111132-2003101031220312-3031022202023012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0301110131002322-1010131113300020-1331113322022302-3330022311233021-0031110001032300-2003120322131112-3022031112023013-1132123132010312)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0213313030021122-0002010220210121-2001133031211332-0212330031130000-0023302002301011-2130301013203322-2122031011112231-1310113312301231"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2130021022320113-2032021233213113-2233233011010003-1210313031110310-3021011320333012-3222213000112223-0210223333331020-0103100012331202"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0323001213002021-1230311122123020-0130021022331311-3231300302102313-3032030211032302-1133302302300033-1322010221212032-1333123232233112"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2323232113012012-2122212322330323-3232333131322301-2230021331021122-0032300031121332-3010103001230002-3320013021220033-0112111232102111"></a>

<a id="canonical-3321320020323132-0221113333120311-1123331311222123-3002233023212111-3021132303320312-1203211013101113-0130312011002010-2302131130133202"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3020320223103322-0210012033031112-0002202110033112-0120230101101202-0300310301311203-1212122003021003-0121203332312023-3203231102133010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3211230010310211-3300103331303031-0030100210101001-3311312022330001-0210323220310230-1102300022300023-0130313333210231-0021131303312202"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0300211202010212-0131013020300113-2231333220110003-2311111211211002-0322300301032220-1101221113012110-2303310202310233-0002002232223312"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-0131022110102221-1210213203120130-1212321313020333-2313100312120212-1211220220130213-0222010200323130-0122322001011010-2223321013222012"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

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

<a id="canonical-1220032203233122-3000212111302131-0230322231210211-3020210030030131-1000102230112333-1122111322013123-2211001223013202-3003201323213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.monitor

<a id="canonical-1313133033320222-0211203120211123-2011121133120223-2012330121212000-2022112022310031-3033203131032102-0000131311102230-1331100311333023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222301013131301-2321201000103230-0301002131232033-3123101310023021-2322333331323021-3303211321101132-0212201100003121-1121223011310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-2330313331303221-2103233122111110-3231230101200031-0302321222033331-1131023033013301-2121113210303311-1311030013031101-2332023311323203"></a>

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

<a id="canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.network_option

<a id="canonical-3113103112213033-1131213323020123-0001032321332121-2132200210103331-3221103001132000-0133132233001211-1312203002310203-0230101010011202"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

<a id="canonical-3110102020013321-0130110010103020-2223210332023333-2020032130303313-1232021320322132-3213132220221201-2112331011331112-0211331220101302"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1103320012003331-1302032320121301-0010323331010333-2310232331111021-1301313300302302-0001323011001230-2103300013301222-3311033323321313): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3332202022111031-1311002012312023-0102130321311111-2110201333200020-3212300113010331-1130101010110232-3302230222212213-1001310221121103): complete subsection reference.

<a id="canonical-1103320012003331-1302032320121301-0010323331010333-2310232331111021-1301313300302302-0001323011001230-2103300013301222-3311033323321313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323)
- nutanix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1002233012002010-2323233112013113-2030111133322100-3220232232221000-0100320032320210-1023020102022013-1000010002021220-2020101133010302"></a>

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

<a id="canonical-3332202022111031-1311002012312023-0102130321311111-2110201333200020-3212300113010331-1130101010110232-3302230222212213-1001310221121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323)
- nutanix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2121000132121222-3200220013110110-1310330032030002-1311103100121111-2323030033102130-2212000222312201-0313311132201221-1122312130332200"></a>

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

<a id="canonical-2110120310132301-0333323103202113-0203001122333010-3120002212310310-3301321201101322-1330321201202130-3301023300132031-2320000213313113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-1312000330012011-0311033122200311-0133210020030013-0222133021131210-2131311201302012-0133103012020302-3201222010101333-2322111031302010"></a>

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

<a id="canonical-1201132021103012-0320121210130102-0320131023123203-1223103021112230-1212210112010230-0223322113122123-2230110002200012-1133330031212103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2110122112002323-0022202030212112-3223120301312203-0110231210212322-0120003101122313-2320320200103322-2101020200211212-0121232033131112"></a>

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

<a id="canonical-3321103030123112-1111002332222312-0001330332230202-1210023031020311-3012220032001012-3202211200231032-3201333232122130-0012331300100203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1222121121023203-2333022321011101-3111002323200030-0311230301030322-1300223301231120-0222202122001023-1330000103210213-1031223221320032"></a>

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

<a id="canonical-3122312333000212-1020203202012230-2222133022312222-1000313330102330-1210033002310300-0122320110232122-1002322113323212-1220111120020002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1210110012132311-2302021120303210-1221312102122311-1011231310230232-0300032331132012-1303122322123133-0011333111102233-1031332331002002"></a>

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

<a id="canonical-3121320200010110-3202032231021203-1123033202300131-0321201113233330-0000003322120203-2310101203203111-1233332211300133-0131130000003112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.static_ip

<a id="canonical-2203000110220030-0022203010302112-2312111031101103-2130332031300010-3333220301122012-1113321033023211-0210133002120102-2133002012211032"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-2301002113013033-2023222220202210-0331002022201001-1002310300230232-0110013220322211-3312012221021223-2211110320023123-1321313101003220"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ip`

<a id="canonical-2323113133131322-2312331032023021-2313201330300003-0000210112102131-1121022210100130-0131033010122110-1021303120032332-3103010133212023"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3310122311022023-0211322001110031-1120120200222332-1220300300100121-3302031011322232-1011323000112031-0010023103033232-2003032322032113"></a>

<a id="canonical-3330323300123232-1103102302330003-2133312012211021-2203110112211111-3233113222033032-1321133030012331-3001232122123311-3110113111132032"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2213013122023122-2001312332022013-2111212130331333-1022121300322012-2330001201321122-3332213110012111-2300023021223031-0333110030221100"></a>

<a id="canonical-2300121321303111-2102132033200200-1320031002301103-1200030310321011-1320000331200122-3201201322030002-0223120333132213-1213003000102121"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.ip_address` property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3022311201120113-3221133323310122-3211221311023222-0033030311002210-3132200331011113-1030210111202210-2033133233200301-1031111222331333"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

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

<a id="canonical-0201232303010330-1232322102210021-0011220302123213-2300222233320312-3332300023223013-0131133112002121-3012000213123011-2220310113310130"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0011312220130223-1111323111013110-0123003120330022-2123130200121301-3123330102223131-2232211010003002-0303201201133312-2011323230330122): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3110112010232210-2120011131230301-2220302120301211-2232003100113011-1101300332230032-1122133200120230-2003102120302132-2111222322222030): complete subsection reference.

<a id="canonical-0011312220130223-1111323111013110-0123003120330022-2123130200121301-3123330102223131-2232211010003002-0303201201133312-2011323230330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1002202201113333-2003321200212101-1332221302202302-3112320220332021-0200033312232313-0003022020213320-2300323000032013-2132303212203111"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3132300222203200-1331013231001130-0020102332331320-3220032100313333-0301111002021313-0003120103201121-2221120002123022-3331022330301121"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-3131002132023110-1133310331301300-0330332110133221-1332103133010222-2320301022300222-2132012231233121-0103113010122123-1000133002123110"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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

<a id="canonical-3110112010232210-2120011131230301-2220302120301211-2232003100113011-1101300332230032-1122133200120230-2003102120302132-2111222322222030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2123333330033332-3303123211020020-0000030230021322-0022003211321113-3132320211223320-1201303000311333-3221202103233221-0330212001300113"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-0000201002321133-1001110023201102-1020113230012101-2010202110123213-1331222233211130-0121112202223101-0320010003032001-0121110321312012"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-0113010320020330-3203221022122221-1220303212020101-0221000113211133-3213213223213232-2202200100121202-1210333021300021-1013202130302223"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2333120313201303-1333322102121132-3022233212011111-1332013021102332-3312111211301113-1123331110332102-1203203321033322-0211221233320022"></a>

<a id="canonical-2311220131323020-0120211323120002-0223001101333131-2030210203222120-3010313221121100-1302311012003301-3310233310221123-2103223223202332"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1232200012302110-1003020033300202-2113332012211032-3002123320313102-2012130103220012-0331320101000233-0012021002021323-2221310133321100"></a>

<a id="canonical-1001200221210223-2002212031003121-1133220303021121-2021321000021012-0212313010120301-2310030333021132-2232102303233001-3303211332301003"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2212102230003211-3201300121123111-2100131201300322-1130003101002213-1130032130203003-3233330110331212-3022101331031233-2322323020313313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3033201102211021-1133131032100111-0210313012030010-1201203333011211-2000222120303122-0010113210331333-2223023212233123-3213103323331103"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

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

<a id="canonical-1332022203330221-0201223130232023-3022200211221333-2012113010012030-0122002133212011-2302203132120130-1232110020323303-2111000111222330"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1113101201110330-2130123312122302-2232132300321120-0020323022113203-0331201203022333-2320102312200233-1232021330021133-1233203320110323"></a>

#### `nutanix.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2030232230212002-3330300230302232-3313330332312333-2013313331323120-1211211311211203-1331331032011013-2312212202212311-1100300021123022"></a>

<a id="canonical-0202223231330030-1110013313213132-1231313010123333-2111001110313321-2212111222303132-1322231330301130-3320000201210030-1210201323003222"></a>

#### `nutanix.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
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

<a id="canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- oci

<a id="canonical-3003032322223011-3031301022113113-2222200133031102-0210110230321023-3132021121030013-3131321100232320-2021333131111101-1213023020030310"></a>

Type: `"single"`. Computed.

OCI Provider Type. OCI Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

<a id="canonical-1003203213301333-3021133222102203-3303120200300132-0320122332320120-2103211031210110-0030302121313003-3303322212331300-3200321122131300"></a>

### Direct properties for `oci`

- [not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203): complete subsection reference.

<a id="canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- oci.not_managed

<a id="canonical-2030210120110033-1000332031330011-3311220030302022-1101311211131310-0213322220111300-3133200132113211-0120323303010030-0013112302303010"></a>

Type: `"single"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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

<a id="canonical-2213303230130210-1232123000001230-3100000320110022-1002000221223130-2232020113333132-0103013122322313-3322223210130010-1231033220010012"></a>

### Direct properties for `oci.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001): complete subsection reference.

<a id="canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- oci.not_managed.node_list

<a id="canonical-3110013010230103-2013313213033313-2112133133113003-3322213122302231-3223011200010333-0003211121311003-0022003000032203-1301332213223302"></a>

Type: `"list"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1221002112202001-1101021300123022-3212311132100032-1032101230112313-1120310120023232-0023002030123120-2311211130222112-2233103003032033"></a>

### Direct properties for `oci.not_managed.node_list`

<a id="canonical-2131203211133011-3202201000012122-2030332022302122-0132302222012232-3001100100203311-2222111213010331-2010201311033322-1210003111201003"></a>

#### `oci.not_managed.node_list.hostname` property

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021): complete subsection reference.

<a id="canonical-0200310200320322-2100032312212123-3112001332232321-1211100320000212-3022330223332122-1033310300333332-2233111012123023-1032011202302131"></a>

<a id="canonical-0332032131233102-2110112210310122-0203101101103101-0221133103211132-3210303322300330-0231201300100133-2221013010012322-0211000103132311"></a>

#### `oci.not_managed.node_list.public_ip` property

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

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

<a id="canonical-1012012110330130-3330112313001321-0021021033203031-3310112321200032-3210020001331202-3211210220220013-0230000103310333-0313230002131101"></a>

<a id="canonical-0322012003013121-3231211012333133-3132232222233200-0122120313203011-0213001332312303-3030030120322033-2332331120311303-3320220201020233"></a>

#### `oci.not_managed.node_list.type` property

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- oci.not_managed.node_list.interface_list

<a id="canonical-0303320212020013-3302223211023030-0302023201103213-3010230012000113-0112121130101303-0013012013012212-2133010313002112-3021212233122123"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3011112020301321-3130210312103312-3320133231103201-2233103202230001-2020000232032333-0001322310331131-0302102212011031-1033313113202311"></a>

### Direct properties for `oci.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3102311321011003-1130113132323220-0222031322212102-0230111001200300-3203320302233030-1003021220223303-1330201303230212-3011122212022133): complete subsection reference.

<a id="canonical-2211033002002031-3313212220213112-2311330020211330-0311102333330222-2031221022221222-1021021201001102-3020213003101211-0112211023001023"></a>

<a id="canonical-0321332120221133-1223110322010303-2222220203022301-2102002312222201-3223220232310121-3131023202123123-0132221302022132-1313131122331123"></a>

#### `oci.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1010232112200012-0020132300033301-1010210232122120-1013210220111222-3020130311301220-3203001223111101-3122112002312311-3202032320212330): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3102331120201201-1123013131203130-3312013331023303-0211200301310123-2023223010100311-1010212202222313-1121322033003112-1002233230032303): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3320211032120332-0302310221030002-1102200232110213-0212332332233023-0321123011211100-3212222010331001-1121303010330333-2013133000103332): complete subsection reference.

<a id="canonical-2111132102220323-1033313102302222-0101333212203330-2101002102312330-2112032211300113-1120121300222200-3230223300310311-0001220133033033"></a>

<a id="canonical-2000031232132101-2130001103133121-2313122302302013-1100010333022230-2012110321221020-0210121200210011-3333222300113103-2212013202031231"></a>

#### `oci.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0000232121010211-1231130103301101-2222100112300210-3302131000013331-0310012132023331-3012333201223021-2301000011302000-2113211101100300"></a>

<a id="canonical-3111320110222122-3023112032002110-1023031211002131-0100131230001020-2230203210220130-2221221330021031-1220202012211023-1110233102303221"></a>

#### `oci.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-0211232200032030-2023121211202023-3202022133101202-2112322213023132-3102021222312312-0121302123031122-2020311212322300-0232322130313131"></a>

<a id="canonical-1111201211011330-2110231101303333-2100020303113031-0032230203100122-1302203301111102-2321200000000332-0100111303020011-0010310211200212"></a>

#### `oci.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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

- [monitor](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0131012311131110-1331123203232210-3112033111100111-0101012130203223-0322033202111301-3033020032131131-3220010330021301-2310032131031100): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2102330113330313-3233303233223023-1132332312122020-1320201202120212-3312322222202123-1212300130222011-2100300003310323-1020222320003330): complete subsection reference.

<a id="canonical-0213020220101321-2112122112302331-0222200220222020-1113212211033021-3222000332212010-2022030203033202-0130202010113332-3022022201220001"></a>

<a id="canonical-0223230010300121-0033212000100232-2221200103300221-1303100222332230-2031220102231001-2112120302023031-3301000021232032-3331203011231230"></a>

#### `oci.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-1312311312112220-3232300223330303-2020000222033220-3203030330031323-0111000103020221-3110023031302130-1031022310032201-3010201121232100"></a>

<a id="canonical-2233320320023102-2311301001232200-1210000303120313-2123023300000233-1003313033333110-3131121300132311-2010102303101221-3313120320012332"></a>

#### `oci.not_managed.node_list.interface_list.name` property

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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

- [network_option](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3331203023301320-0301201003322220-3020122231131221-1311333031121132-1010213001311233-0323312210033301-3112121220312313-0202012230302203): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0201023232300210-0231211110003311-3033313233023311-0133331123222210-3032323301303021-3013023111220321-3100130321101303-1212100032013232): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000212321023123-0023110130113013-0202313132312033-3231210221023330-0023031133333013-0302122211032302-1231000310321122-2322211332312312): complete subsection reference.

<a id="canonical-3023012021132030-0303112003031232-2110001012313121-2101132031313111-1132330330313011-3031102032322010-1313111121133302-2210031103010232"></a>

<a id="canonical-0010100331000131-1232301122312213-0120322020122101-2012002233310030-0013202123030101-1210203012222103-0332321021330122-2100333301313030"></a>

#### `oci.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2230102222111331-3321002013100012-3323133132032000-3213021322131230-2223303021322111-1201033100111231-1131121100030121-3330010033131223): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3300223210132322-1233300121131021-0310331023003003-3122310313033032-1031100123032232-1010211012313220-0132103132011132-1100132332232302): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0223213120323220-0313002033232320-3003203021121023-0102233303311133-1012010231032121-1122101331331212-0131322020322300-2113210333101022): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1302133111102333-2202123032021013-3232133131130303-2122321122230020-0323023231230011-1001111332212221-3310312021020222-1323313120221233): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0030022001111011-3123002331100000-0230013322122011-2302333132120200-1113230003201332-3322133332131333-0000311223022113-3011033033302112): complete subsection reference.

<a id="canonical-3102311321011003-1130113132323220-0222031322212102-0230111001200300-3203320302233030-1003021220223303-1330201303230212-3011122212022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- oci.not_managed.node_list.interface_list.bond_interface

<a id="canonical-1330230320023211-2313001121322300-1313133302112131-2102030122102323-3233203200103230-2122303212332100-3310300223132100-2231333202131212"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-1301020301322132-0330012201322110-2110100201032311-2121102202201302-3030332203131222-3013331032131030-0233113020033301-0332002213122000"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0113232322221120-0030032020020010-2031322022111303-0311222001113203-3330022032132022-2010121113030322-3300220201221131-2012102212231123): complete subsection reference.

<a id="canonical-1100221121102333-0023102012300230-0121210232222223-2213122131210232-2302030001213312-1022130330002103-2303010020203121-3320320302321201"></a>

<a id="canonical-2311203130311233-1212021130230132-0220112102100203-3003332132320000-3213113333203010-1211120023221033-1231112101112300-2012300022313230"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210313101300301-3002210100313300-3132100321322322-2132111233322122-3230301103203122-1021103102123332-0011223202210020-1022010231032331): complete subsection reference.

<a id="canonical-2110032200032223-3030030122330113-3221130022311022-1011302021213020-3130233001032203-0203132220333300-2301202331112222-1303022011310332"></a>

<a id="canonical-2301023113131110-3012031122231012-3103203020010310-2310332000230303-2110113011220212-3221211120232201-0121101311000031-0112012100101023"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-0120320133010011-3203131013123302-3200301103112222-0213232333212313-3010030232200311-1202031320122100-2112323300312021-3221011201210300"></a>

<a id="canonical-2333301111202121-1130232222213012-0110122230201200-0231232022230211-1001112033232310-0200212202000230-0023110103332220-1123022010331320"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-0031213231121030-1000333020012320-3113330213202223-1212120300010232-1112223010222301-3200323030330003-2202011330230120-2203100102013030"></a>

<a id="canonical-3022112321120110-3012121232313100-1022333323220122-2323000232120100-3213031121323101-3112020232101303-1020000132331103-2120012321111122"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-0113232322221120-0030032020020010-2031322022111303-0311222001113203-3330022032132022-2010121113030322-3300220201221131-2012102212231123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3102311321011003-1130113132323220-0222031322212102-0230111001200300-3203320302233030-1003021220223303-1330201303230212-3011122212022133)
- oci.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1211112133100321-3330311021310300-1123330203012111-2110023301321313-0331032020113222-2102213302001222-0231312333312131-0210203030311030"></a>

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

<a id="canonical-2210313101300301-3002210100313300-3132100321322322-2132111233322122-3230301103203122-1021103102123332-0011223202210020-1022010231032331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3102311321011003-1130113132323220-0222031322212102-0230111001200300-3203320302233030-1003021220223303-1330201303230212-3011122212022133)
- oci.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1113310210333032-2031112233301022-1313033332303123-1220122120003331-2323212302012100-1101211301313101-2213031232120202-0331011222101211"></a>

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

<a id="canonical-1100122331021200-3220221320201001-0112020332013103-3230202320220322-2321321032301330-2030103031320103-3333203122320111-3112303012330132"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-3233012010320101-0010203032232030-3020002032210303-0211300133121032-1121211000132133-2331000010233202-0010001103220012-1221311131103220"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1010232112200012-0020132300033301-1010210232122120-1013210220111222-3020130311301220-3203001223111101-3122112002312311-3202032320212330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- oci.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-1001231330013023-3020000311212322-2322333320323232-2122031212222100-0222133103312031-0102322312121310-3001332111130120-0312021123322010"></a>

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

<a id="canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- oci.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-0300222222123220-2303113201000032-0121012123001212-3311331301303313-2202313022213333-0211333011113032-2103213321210110-3320031203003101"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-1212300102002100-1120212033100021-3132112203331023-3333311230323111-1002300212000122-2000320121123113-2121120201022302-3233033301030233"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1002121101103233-3030310200330223-3213203112330103-0222321031131130-2020331121010022-3020231330012112-1321100303330320-3031310212020002): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2120211133331312-2311322313310103-2103330233303230-3103132213123000-2011202303133030-0200203032221130-3110133001003023-3233200100320220): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0133223332301003-2030321302323011-3003001000101210-2032222220033130-0001323122001132-2302320101123213-0033233231033331-1003033112033320): complete subsection reference.

<a id="canonical-2201103032323103-3323000221131303-0223121323011302-2320001010010223-1130022233313231-1203101323202223-1221121011002122-3010302020110001"></a>

<a id="canonical-2000222131122111-1032003123322230-2230220322331302-1222023133230022-1321201000033122-0203310301210322-3030231032322022-0030103120322200"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-0211301120200122-2333111200031030-3011232233310320-2100013121233203-1323102003222221-0101202333112323-2031020131220110-1120132101012303"></a>

<a id="canonical-0310320313032022-2321230133000202-1021021233321012-3202330032030221-2311101030321020-0032111011012222-1303210330303000-0003323323021021"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1113021223301303-1110103200323112-0213221010122323-2131133131230302-0021233121032130-3010231031223210-1211031120101130-0013313120113332): complete subsection reference.

<a id="canonical-1002121101103233-3030310200330223-3213203112330103-0222321031131130-2020331121010022-3020231330012112-1321100303330320-3031310212020002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3101302233102012-3211301020123101-2223333332131123-2121322133302203-3201010103101302-1320000022230323-3301210212022113-1113020130122211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-2120211133331312-2311322313310103-2103330233303230-3103132213123000-2011202303133030-0200203032221130-3110133001003023-3233200100320220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-1222232312011210-2002120330033321-2320011112000312-0103021310002102-3101120132013011-3101002322321030-0310101011312312-3133130223032232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-0133223332301003-2030321302323011-3003001000101210-2032222220033130-0001323122001132-2302320101123213-0033233231033331-1003033112033320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1130302030020220-2033223302230221-3123022130112123-3201213210131012-0201200330332333-0202203032201303-3231211212130131-0101102101213312"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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

<a id="canonical-3112230030001013-1023101313222203-0323113101300302-2102022233323332-3212001013331211-3211101321333332-3333202011132013-1223131133333121"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3300013010002323-2102021232012100-2220123212310033-1113221001210211-1012203230111311-0210033332010301-3120212311120211-1100111123100010"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3132002202321100-1203131300131012-2201201100302201-0032100111033323-3311232311023131-2000100022320310-1132303311100322-1010313131202003"></a>

<a id="canonical-0010011132000132-1103112331013311-2133000221211312-2323222021130032-2001132102222023-3331313103131230-1311310323102332-1030303323130320"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0321000221220330-3323331232122001-1202130210013221-2012301332112331-1210113323023213-0130132233301223-3302330120203122-1103031023211223): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1322311122112000-2111322230312113-1232100311101123-1312221210000113-3313010102033232-2203011022112332-0303020322120010-3122330232021322): complete subsection reference.

<a id="canonical-3331001231330222-0023233130021303-2031000312211121-1023232113123200-0120221312021032-2313220121222021-1112130300033303-3033232031230210"></a>

<a id="canonical-0111331303122013-3111321101211132-1112202333300332-2313111322321121-1300022002203321-1121112111020030-1131123303111033-3031303113312102"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0111123223212132-1302310031112110-0322120022002010-1221110333310102-2032101210122330-3133010302033303-0232023221130300-3211223331230311"></a>

<a id="canonical-1323000310012223-0022220322020312-3322203321201002-1331100311120230-1233100101330010-1323231230330210-1221332203000103-2123300212110020"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

- [pools](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3112321201101130-1303233102321111-3121130322001330-2110321223000012-2333003000031311-2303201123301332-0323021131301300-1221010111120311): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2030101003300113-2111210023230122-2212000212223133-2110031013333110-3200222313211213-3010122333231121-2200030010020330-1330200222133000): complete subsection reference.

<a id="canonical-0321000221220330-3323331232122001-1202130210013221-2012301332112331-1210113323023213-0130132233301223-3302330120203122-1103031023211223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0101101120123001-1031013033211013-3003321022102301-0110302223020110-3213322322230302-2123120223212011-0322302200030221-1032031303003203)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1013213101303111-1013332032302331-1222120112032112-0020033303133311-3001020332002123-3222310132000222-0320300100102101-3022311322300001)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223020323302020-2321000323211213-0321233020023301-2101321101320120-2012210033020133-1021201023212131-1010102011033332-3221111112110021)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2302002032331313-0310210031231000-2021023123111023-1212120303102001-2220302221330121-3203211002013231-0212323002021100-3010100112021000)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0133223332301003-2030321302323011-3003001000101210-2032222220033130-0001323122001132-2302320101123213-0033233231033331-1003033112033320)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-3103003312310103-2232211003332312-3112032313120022-0021112322331103-2031020301133000-1130320002010001-3203101323311313-0201311233022131"></a>

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
