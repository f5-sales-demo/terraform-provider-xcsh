---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1020320202231003-3332013131130130-2231033002133310-2300222120210200-0333030202231123-2301020330103222-2100311023003021-1102103313213203"></a>

## `equinix.not_managed.node_list.interface_list.labels` property

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

- [monitor](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1110021333300030-0200311230113121-3221102202102101-1221123200303331-0301310333011123-3213113100030022-0001230032310102-0230033203321010): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1323323120120101-2321202200232231-0110330110130230-2101130312100033-1302311013303230-2223112202313332-0133121321131211-0222312220120023): complete subsection reference.

<a id="canonical-1223022222311132-3223023311203111-1122001010023031-2211203030033021-2012132131000232-0220120132012011-0010102211221221-1220312011213132"></a>

<a id="canonical-1210003212223032-3110233100302211-2310320030123210-0133303321223231-2113330131120033-2111203322122031-0202032012120130-0322132212202332"></a>

## `equinix.not_managed.node_list.interface_list.mtu` property

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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-2111312310032113-2313322232312210-2101223131110132-1001023332300213-1300013213033230-3022103000132133-0202113100030200-3000300111023013"></a>

<a id="canonical-1021201300120032-3202201020021030-3131301130301112-2223232320103333-1011331220331113-0103322131321300-2111312301021331-0023213322133103"></a>

## `equinix.not_managed.node_list.interface_list.name` property

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2323202132320113-2303111031133130-2231212003203310-1000322102133111-2203222310202311-2231023101001213-3033110323332113-1002311302022133): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012003111230120-0323321130213212-0320111331200320-3001132320101023-3032230221221033-3120213100312111-0021221110303222-1103101211032123): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2003311303223203-0132122220113021-1123303003120112-0110210322102203-1221333322110003-1320133002331222-1200221030301102-3001012013201202): complete subsection reference.

<a id="canonical-1001230123311023-0332100002002102-0233020131331220-3110210311032320-0332303222101023-2311212003201220-3001313030102033-1000230310210033"></a>

<a id="canonical-0311112123322133-1123321132120032-3220232200313112-2002320121231101-1211011021302303-2200101101310321-1122202100120113-1102313013031012"></a>

## `equinix.not_managed.node_list.interface_list.priority` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3123112210222330-2033123122010003-2010212100031303-0021311333122200-3322313001122310-0310103330203323-0213113031311300-2000323231131233): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1311311031113301-1222101320122232-1021032020221301-3001003301023321-2210113020000222-2303021212111020-2322222030302211-3000212001031230): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0030210002121220-2021100123312212-3023201111203123-1131121103123133-3110321221022022-0322301200233333-3200323120130232-1331211103223102): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1333031120021312-0233331213013221-2202103101122103-0333301322230011-0121101130302213-1300112333031203-0320201312300210-3323001323331010): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312211311222033-1000322033022300-2220002201230303-0001132132012310-0133313002031332-1221113303311223-3103230103003310-0000222001311023): complete subsection reference.

<a id="canonical-1313310022002321-1010231013022121-3112323133130003-3302131212211300-0213021031101322-1300201032332332-3320021013112120-3321201031230030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3101132110310033-2311111020133233-1333103023213223-3301113301232001-3013023120101321-2203022331212110-0331222330120322-0230202120330120"></a>

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

<a id="canonical-2302122333122321-3102010313220033-0133131333230013-2011301020120030-1033222320233001-3000310132133202-2000102211321230-0311013130122221"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0000020021233312-2110300223313201-0012032021023323-1311222312201030-3122011110031002-3133112000313303-1330222313012110-3230221211010120): complete subsection reference.

<a id="canonical-2313330302131033-1233222121301233-3333001311223021-1222101033211032-2013012113021221-0100222333213303-2201223213123110-3023013231201113"></a>

<a id="canonical-3131011012303031-3223202311110001-0202102110130323-2113110023130213-0212323333302020-3200210110320123-0020121220010123-0313223033301323"></a>

#### `equinix.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3012100201023311-0001323122111300-2031230003313112-0032131330100030-2013332011111320-0122213202302113-1212002312201122-1201302120012331): complete subsection reference.

<a id="canonical-3331311322020000-0113002301003320-0210210322321132-3002102023031003-1132302011021230-3231211233113112-3200312312101122-0322133001313311"></a>

<a id="canonical-2013012111113102-2100121112231331-1230312101301200-1111131200223002-2303230030200230-0210223310303021-1203022111132221-0230112322332000"></a>

#### `equinix.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0131232321200333-0113102022212133-0120312210321312-3103030100303012-1213221133220001-1231211121222121-2312221131101233-1000333301311301"></a>

<a id="canonical-3031223311313301-1123000332213331-0102322133211333-1011333013211032-1333230000032030-0220322131033122-0111131321202130-1122102122110011"></a>

#### `equinix.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1102032132133212-3000331333331211-1122010301121032-3023001231211031-3120202201101110-3113001020102122-0022112121300003-3103213122212011"></a>

<a id="canonical-2003113200230111-1113302302122201-2032300200311203-1032111030102322-1001131223331112-1213222023300021-3110101231102131-1120223223232201"></a>

#### `equinix.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-0000020021233312-2110300223313201-0012032021023323-1311222312201030-3122011110031002-3133112000313303-1330222313012110-3230221211010120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1313310022002321-1010231013022121-3112323133130003-3302131212211300-0213021031101322-1300201032332332-3320021013112120-3321201031230030)
- equinix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1311000112321123-0012221210021021-0203233003120020-3301210331333223-3231332330201331-3121111000131133-1031000222122202-0102111303211003"></a>

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

<a id="canonical-3012100201023311-0001323122111300-2031230003313112-0032131330100030-2013332011111320-0122213202302113-1212002312201122-1201302120012331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1313310022002321-1010231013022121-3112323133130003-3302131212211300-0213021031101322-1300201032332332-3320021013112120-3321201031230030)
- equinix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0312221202213031-2211231303331121-2030210130201002-0201012011100102-3202232210233113-2101300311121230-2300003201200203-1313323313311132"></a>

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

<a id="canonical-0321233112000303-2202122302320012-3330212202132023-3322332020001011-1301022021312303-2123223312311230-3333223202132111-3131031223331011"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-0000111121111031-2231302200102320-1203030112321303-3111101010321300-1120020200220312-3203330003202101-0032010100333101-2101001130113111"></a>

#### `equinix.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321103320312032-1311203003133300-3323210313023033-2012201233013011-0120021033012231-0100111232022332-0000020121332102-3211130202133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2020123123032203-0210031310120102-1331011103200012-1112113101221131-1312311030222320-1121200111230203-3020113302011120-2201001010210202"></a>

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

<a id="canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3323131330103211-1312201012123132-2031131110123123-1021002101120200-1023012321110322-1011200320230321-1211323111112312-2032322201012013"></a>

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

<a id="canonical-3212110302132212-1012101131212223-0112333330033101-3112322131003110-1001203330031122-3210332222102200-0300212313210030-0203132302012031"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1331210223312312-3232322203221023-3011223123100113-3021021033200303-1110311222110113-1330021133001200-3013003010322331-3112311313300030): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2323221321222111-3020112220031110-1222101013001110-1312202322213030-0102110133211112-1322301302033132-0213133101330220-2001013112032002): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322): complete subsection reference.

<a id="canonical-3132310123200322-3023330232130011-0303302132220023-1130220300333131-3021321201021222-2300000033102020-2121010012333013-2010003202110210"></a>

<a id="canonical-2202131133110320-3311203321102313-2123201012110032-2233032030033103-0110121022123100-3311233010002113-1010020233301302-0021030113113011"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1223313233323301-3310101002102211-3203202103213332-1232133011222232-0103301221231311-1230330333332010-0120222102002020-2231112000103102"></a>

<a id="canonical-2122203332332221-1203112103231312-1320230303330131-1010030000012223-1221221213010003-3101230312213021-1030032102232303-1101023032203003"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2132212033021233-0203103322120031-0331022200201233-2120112112130312-0322101122103310-1311203330102123-2321100311122300-1110130322200201): complete subsection reference.

<a id="canonical-1331210223312312-3232322203221023-3011223123100113-3021021033200303-1110311222110113-1330021133001200-3013003010322331-3112311313300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3100130231333122-0000020131013321-3321131020111031-3001033222322332-2031003220133301-2303320332102113-1230110333210002-2112021201002101"></a>

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

<a id="canonical-2323221321222111-3020112220031110-1222101013001110-1312202322213030-0102110133211112-1322301302033132-0213133101330220-2001013112032002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-1020111203113011-0012023333323133-0011232331002002-2033200013011002-3033320000223310-1023231000333121-3303233101310100-3303230010002033"></a>

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

<a id="canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3031020101332112-3133121322232323-2233210102212130-2211113121200223-0101213132212213-0233333122100323-2201233203030012-3310132001010013"></a>

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

<a id="canonical-0212133013313010-3311102012211310-0121303102322003-0300012020333130-3223200020210131-0133333103222232-1312131133320323-2211032203310100"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-1223012321030203-0332032021100231-3133023311331133-0301212210223233-2110333133102331-1021100102233331-3022121201101110-0000131001313023"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2300200110203011-1320313230233113-3222210332001301-0331010312023211-0222122300313013-0012212010032033-0122000003223221-3313302310333001"></a>

<a id="canonical-2011113021220033-2010033122330121-2201112003030102-1133113113011100-1120110103220310-3311013112023003-3333311233033032-2013112310232231"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0210213110333222-2113023330023232-0130111123100213-2113132000011333-3132003230121103-2330203103133332-2121230110322113-3230022332102323): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321011331100303-2102333020213331-2132113002212120-0032100313333232-3220121300310303-1102332222122001-3013132111332223-0211123330321023): complete subsection reference.

<a id="canonical-3013001010022202-1300120220202321-1011311132023332-1101210303311013-1212030132212111-3332103313121003-1020020020011301-1031323102231321"></a>

<a id="canonical-3221323102203312-0132022002110210-1322322313203220-3333012231312222-3002233322023001-2032232221222303-1221020330131210-1120322000213033"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-1031300223311233-3003210113232012-0223221221233332-3003303323230032-3202033133101123-1200302233012302-3113210012222022-1322310001301021"></a>

<a id="canonical-0002001200312310-0110013100100123-1103021113110323-0213112330223322-1203101230103222-3133122211010313-0101211201130330-0012000222113012"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2300020213212210-0113310102103302-1322212031320233-0332001203121210-1321012113032312-2130200201223022-2021232103101021-0122230320112033): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0021022211230223-3303000023122210-2110010132003000-3331222220300220-2002222131232000-0311331230203302-0232110203022132-1300102232103311): complete subsection reference.

<a id="canonical-0210213110333222-2113023330023232-0130111123100213-2113132000011333-3132003230121103-2330203103133332-2121230110322113-3230022332102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1203320222332002-3311323033033003-2310223303101032-0333311112102300-0021111231000000-3213111312122330-2001121302223103-0013320232020320"></a>

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

<a id="canonical-3321011331100303-2102333020213331-2132113002212120-0032100313333232-3220121300310303-1102332222122001-3013132111332223-0211123330321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2321102321003201-2323233320031233-0231211003012331-3213121221133322-0330001330000333-0000110302220213-3022003132313001-2311021022211210"></a>

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

<a id="canonical-2300020213212210-0113310102103302-1322212031320233-0332001203121210-1321012113032312-2130200201223022-2021232103101021-0122230320112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2211331021021010-2300221021110121-0223112010010102-0120201322210301-3112122332032331-3101201101023113-2312021320303202-1113231201021333"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3321232312301200-2110320000101010-2303102012132223-3013330301131012-1133130211331312-0300101013131201-1330011120001100-1202001121012132"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3200202000301003-2123312213022233-0210213111323131-1020030001003302-0201122312212300-2230320301011003-2233132011231232-1010113021203332"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0001222132120232-0131330023032213-1232113302321223-3232113313102003-3221100321011303-2022211132213110-3302030222311001-1020321313231121"></a>

<a id="canonical-3232002010111213-3231321211020203-2020121111300101-2033130230133323-0312201230121210-1121113112220223-3130011002203021-1020133011111103"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-1320330103303200-3012030123200213-1102032021202213-3302320313320131-2212111123220302-1132130110011100-3323323213103301-3302213030202313"></a>

<a id="canonical-3113120112300112-2221022122003102-2302031102022320-2210211320101310-2030300123323033-2320123303111110-0212100033002132-2301102101032210"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0021022211230223-3303000023122210-2110010132003000-3331222220300220-2002222131232000-0311331230203302-0232110203022132-1300102232103311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0010121202130223-3010231131331121-1213322222021122-1332100132030113-3220100332222321-3323113101210000-2201100102111331-2212022210131322)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1320032010300312-3221310110331230-1323002220121213-0030310201111220-1112202332331121-2130202033013300-2110000113110133-2110312330010003"></a>

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

<a id="canonical-2132212033021233-0203103322120031-0331022200201233-2120112112130312-0322101122103310-1311203330102123-2321100311122300-1110130322200201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302)
- equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2020220303030012-0313103021033000-3220031001030203-3213212220000312-3022002121121100-1233121222301213-0011013103122123-1221010123120121"></a>

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

<a id="canonical-3310310332002320-0023020030201133-3130021133311200-0203100001130112-1332112131010033-0102303013003333-0303123313323202-0002001333322322"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-1012031011031021-0211210121020202-1001311123113022-1100003030213013-2022121203100233-2120323320000003-3232312200121000-0120022133223100"></a>

#### `equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-0030022103120212-2330122332133101-3323000322033223-2203313012323233-1122132021130130-2021010032223303-1312302223202331-0031202333110110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-2220132022033110-2200131033021032-2130320033202202-2100201030321312-3322313010123223-2101322130311220-1231120003023212-2003222223013122"></a>

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

<a id="canonical-3203112000132310-1120003330122110-3102102133131222-1110321222332230-0233332033222312-2322320102023330-3021231332101303-1021121000232211"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1031323103331000-0022320113010312-2113310332312313-1103201221223101-3331213100102132-0031332332211123-1302231010331121-3302013001122023"></a>

#### `equinix.not_managed.node_list.interface_list.ethernet_interface.device` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0322322013122100-0321333111210132-1130020102222030-1211101013232232-1213313021201122-3001012132323200-1103323213223212-1321232202332103"></a>

<a id="canonical-2323303001123202-1023131300311212-3212101102123000-0000102103010330-3110323020023232-2023120333002221-0003312020201312-0210312311302232"></a>

#### `equinix.not_managed.node_list.interface_list.ethernet_interface.mac` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0303131330310030-2122221213311030-1133230012322123-1122331213211300-2233330031333101-0232030211302130-0021222031231103-2131220133321003"></a>

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

<a id="canonical-1301023121111132-0200231301112131-0300112201231012-1010123333201302-2100211220130130-2331322203033102-3021101332312332-0302212020310111"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2003203122233200-2000023333131011-1131222302100012-1331000022000323-0022203202210003-3221323222313323-3030100210120231-3321323321203321): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112): complete subsection reference.

<a id="canonical-2003203122233200-2000023333131011-1131222302100012-1331000022000323-0022203202210003-3221323222313323-3030100210120231-3321323321203321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1110320002113222-1213023330313311-1010133200313232-3321233313021120-1013312322222033-0213132100102001-0233032223200210-2212221020103222"></a>

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

<a id="canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0023203232320312-1101111310003222-0323322332131110-3213231131011111-1031221021120300-2120310112232332-3010131310121033-0133200200303321"></a>

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

<a id="canonical-0323220223030333-0123302330103113-3313313321233231-1122210110030331-0113113312223030-2133202313102132-1131113320011110-2202310013222011"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032): complete subsection reference.

<a id="canonical-0120210310003113-0132013210200100-0222201301200312-3310322112312301-3301033012021002-3010000202020013-3023310132120013-3011001001113303"></a>

<a id="canonical-0010222221032302-3323130021131130-1123322332102031-1001111031113031-3123202100000121-2212103022332301-1303113123122010-3203003120110220"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030): complete subsection reference.

<a id="canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-0113100313300120-0222333313212112-3103333113233333-1203032103020123-3100302013231111-1312321121221210-1301311321131212-0002323310022110"></a>

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

<a id="canonical-2123221223232330-1203331233311100-3032221303321110-3301110021022100-1320131333223023-1032010233001122-1120011212102012-0202011102003301"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1301203133012313-3032033110030111-0231222021331233-1303223101102201-2122100232301030-2202101313003001-1322033133211101-0021202303000201): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2211000302002331-1011030322133202-1121322003210330-2023010310212321-0330302202123022-1003230120103322-3033030130300131-0132300303132003): complete subsection reference.

<a id="canonical-1301203133012313-3032033110030111-0231222021331233-1303223101102201-2122100232301030-2202101313003001-1322033133211101-0021202303000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1031220123033030-1303220333021122-1232231130312313-2003111023221220-3123112330003021-2321023101023033-1233032213132112-0010213300133300"></a>

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

<a id="canonical-3230300130221122-3021022233030021-2312232010322220-2021202023033023-3330331022210023-1201211321033110-3222220022320222-3331133101131100"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-2103033333131120-0222111001003220-1122121230122003-3001133220301303-1211321313112223-1110233321330012-3033332030033332-2323332003333100"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-2211000302002331-1011030322133202-1121322003210330-2023010310212321-0330302202123022-1003230120103322-3033030130300131-0132300303132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-3333121023010133-2211103102103022-2310303011010332-2101333021211221-1113231110122231-0100110303022230-3323031122303313-1333200333303213"></a>

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

<a id="canonical-3211232121202013-3203200211321101-2131230331232000-3202003103011302-1011112120120033-2001012301011131-0222122311323012-0230313323311310"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2033023032323223-1110133220012320-3102203322211121-0121331221000021-1330211002123012-3112200003212332-1320232031121002-1133220311132233"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0002230101231233-3123000302302011-1103231202000021-0033220301122321-2011003033301131-2100200321300120-1102122101032132-2133322330302221): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1130210132213322-1001310232123201-2132001312323102-3301113123223030-3200121230012012-2223123301220300-3020131311103232-2303310321011332): complete subsection reference.

<a id="canonical-0002230101231233-3123000302302011-1103231202000021-0033220301122321-2011003033301131-2100200321300120-1102122101032132-2133322330302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2211000302002331-1011030322133202-1121322003210330-2023010310212321-0330302202123022-1003230120103322-3033030130300131-0132300303132003)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3022011100231322-0032331123333003-3232321233011323-1333301222331233-0030233330322113-0320332333221021-3111200231000111-0231122100202021"></a>

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

<a id="canonical-1130210132213322-1001310232123201-2132001312323102-3301113123223030-3200121230012012-2223123301220300-3020131311103232-2303310321011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2211000302002331-1011030322133202-1121322003210330-2023010310212321-0330302202123022-1003230120103322-3033030130300131-0132300303132003)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0022023222031012-3332022111113113-0031020320022213-2011122312103131-0011301300213122-3220101312212201-0320221201201220-2231310023302011"></a>

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

<a id="canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-1332330313012111-0203131110233220-2211223230210302-3012103300102022-2222201033011131-3113202233220002-2220032302222323-0332102230003301"></a>

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

<a id="canonical-3320011321111202-1103321232103113-0121301230023223-3131121323033120-2101030102123310-3330101123302010-1133001002322111-2001230131223111"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2030322330122230-1230201101313022-3021002021311120-3201113233011321-0330312310110002-0321222132003333-3030323013311232-3032300102201330): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0103223213122120-2330103111100122-1321312130133203-2003102123213320-2210003102023013-2322231303123321-2030120201331321-3033323020031303): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0121322233000130-2222031122313000-3100122320011021-3331332223320013-1030203220220131-3231113120103310-2331301003212021-0032110322020233): complete subsection reference.

<a id="canonical-2300232331101130-3222320331231120-3303013332233002-0330123110201021-1113223212122202-0231133210033120-2110201113101013-1123022032122102"></a>

<a id="canonical-3221332030130313-2231022321111200-0032232003231303-1220110011313221-2013031332122111-0122120213133120-2012321322303020-2213303103312111"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3133230102203220-1130101223000020-1132331333320120-1101202332111330-0232312112231103-3112123212031001-1210210123333330-3030121313033300): complete subsection reference.

<a id="canonical-2030322330122230-1230201101313022-3021002021311120-3201113233011321-0330312310110002-0321222132003333-3030323013311232-3032300102201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2330001313311001-0320210123032222-0113123100321110-3110120003311300-0320301222320123-1321121023321110-3332201202110002-2012303121212203"></a>

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

<a id="canonical-0103223213122120-2330103111100122-1321312130133203-2003102123213320-2210003102023013-2322231303123321-2030120201331321-3033323020031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1313321300221111-1102200100101230-1002223330303332-0333010210010330-0332221233210113-0210102000223130-3022312320010203-0020002120031002"></a>

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

<a id="canonical-0121322233000130-2222031122313000-3100122320011021-3331332223320013-1030203220220131-3231113120103310-2331301003212021-0032110322020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2322032230122321-2031023200132110-1300201113131110-1023212130210002-1231013223022020-3122120130320123-2132131233312301-2133013313210020"></a>

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

<a id="canonical-0003320203232000-2021100233320023-0101233210322220-1320032322332020-0332311322110200-1302023201120331-1331311230132132-2033320223031133"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-2210303210211111-1312301123233312-1013103110213031-0322201131130010-2003300230133033-1030121103230012-0110013331010123-1320320000223210"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-2321300023133310-0113312013013031-0112201233112231-1232133202212000-0103323101310130-1301330323233011-1201321033230030-1211030033032110"></a>

<a id="canonical-0021130201220011-0123300132310130-1233320210032130-3113211213230002-3111001201330023-1010122001232200-2232220002011103-1003201312013300"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2322002312011030-0201131123320230-2330130210331313-0220322221112303-3230222322220113-0302311310133230-2031022303202231-0131302333233221): complete subsection reference.

<a id="canonical-2322002312011030-0201131123320230-2330130210331313-0220322221112303-3230222322220113-0302311310133230-2031022303202231-0131302333233221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0121322233000130-2222031122313000-3100122320011021-3331332223320013-1030203220220131-3231113120103310-2331301003212021-0032110322020233)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2201313031112100-0110221003123321-3323231222022312-2013301302021231-1023303111001120-1322022002232231-2221123331002132-2230120101031122"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233111322230321-3000010130011123-3303103131102013-0120031030031130-3001001010133123-3113310030231131-3213011300320331-3010300331113132"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-1021001032300203-0332332032133202-3032103233022331-0133332311310333-2120232132012311-1321010210021223-2002101030111330-0313121112202311"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3332010300003230-3011131213232322-3123013320112220-1203030312313002-1111203113231013-2202223331223032-3213211310311012-1101313223003122"></a>

<a id="canonical-1101112203233310-3322330203303133-3330010210113233-2303312222103333-0231201101102203-0311110310112010-1311021131032321-3202213101333031"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3133230102203220-1130101223000020-1132331333320120-1101202332111330-0232312112231103-3112123212031001-1210210123333330-3030121313033300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1200223131322302-0121003322312320-1300021230112222-2301311302333110-2020111202313110-2212102303013312-0203230102310332-2303222101000112)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3321210233311330-2231111020231130-2013003021331211-3133322201133322-2133232220112303-2100330123221001-2311102002132200-1111320201103030)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1112100312131212-2021000130202111-1202231200111002-0332010001122033-0102222220110233-1211113221121101-3111321001012321-0311303022032032"></a>

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

<a id="canonical-2003210020202120-0300313310233330-3230313311101123-2222332121002210-2330002202102120-2033003322230311-2102012110300323-2313000133333031"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2211120121101302-2123312230301100-3100230302210013-0223321001313132-2310131100223220-2311120002330003-2131300110303311-3302220321032332"></a>

#### `equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1110021333300030-0200311230113121-3221102202102101-1221123200303331-0301310333011123-3213113100030022-0001230032310102-0230033203321010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.monitor

<a id="canonical-1220021010230031-3303033123111310-2311203301310320-1231201000002023-3032130013212331-3322130010313022-1130323310002232-1032300311012111"></a>

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

<a id="canonical-1323323120120101-2321202200232231-0110330110130230-2101130312100033-1302311013303230-2223112202313332-0133121321131211-0222312220120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3110313221303000-2020230000313332-2113301232322032-2012000023312312-1133001031131023-1223202103123231-0102031030002231-0230102202111013"></a>

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

<a id="canonical-2323202132320113-2303111031133130-2231212003203310-1000322102133111-2203222310202311-2231023101001213-3033110323332113-1002311302022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.network_option

<a id="canonical-1300223003120233-0112001230123213-0233211002101211-2113233303213331-0103020220020102-2022122100123322-2132330203311321-2002230102003210"></a>

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

<a id="canonical-1120221223231332-2303232002200130-0022231102121030-1113132310132211-1120320021013023-1011330023210101-0103003011223121-1210030133121001"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2122122222311221-1221200200022011-2203003330130012-2111031033031123-2222012100103000-0233202021212331-2113010011012123-2100300111110320): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1301331213310322-1232132113200223-2003302232232023-0032212302333303-1030121313102030-0011010100031330-3001301133122013-2333120110032112): complete subsection reference.

<a id="canonical-2122122222311221-1221200200022011-2203003330130012-2111031033031123-2222012100103000-0233202021212331-2113010011012123-2100300111110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2323202132320113-2303111031133130-2231212003203310-1000322102133111-2203222310202311-2231023101001213-3033110323332113-1002311302022133)
- equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0110013203203023-2233213302123321-1031221203000333-3103213031100023-2233010032122203-3322330311211233-3202323321302322-3011322030130313"></a>

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

<a id="canonical-1301331213310322-1232132113200223-2003302232232023-0032212302333303-1030121313102030-0011010100031330-3001301133122013-2333120110032112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2323202132320113-2303111031133130-2231212003203310-1000322102133111-2203222310202311-2231023101001213-3033110323332113-1002311302022133)
- equinix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3120302220210220-0133322321132113-3313031330113200-0021233220001030-0323121000102323-0331330333000212-3222310113313112-3112302011131200"></a>

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

<a id="canonical-2012003111230120-0323321130213212-0320111331200320-3001132320101023-3032230221221033-3120213100312111-0021221110303222-1103101211032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3000300301300031-1011312323223312-0203313113232321-0123111010122232-0230303002200133-3103333020021331-0021122120232331-1122333023210121"></a>

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

<a id="canonical-2003311303223203-0132122220113021-1123303003120112-0110210322102203-1221333322110003-1320133002331222-1200221030301102-3001012013201202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3100033030211320-0030321131233011-0012101301303231-3221221032020120-1033003312033301-2013123322112012-2300200122113200-0320103331013303"></a>

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

<a id="canonical-3123112210222330-2033123122010003-2010212100031303-0021311333122200-3322313001122310-0310103330203323-0213113031311300-2000323231131233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3112323111223213-2220102220210131-2010202200033301-2132201221213212-0012212003103033-3110031310321031-3212311003310310-0031201222013303"></a>

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

<a id="canonical-1311311031113301-1222101320122232-1021032020221301-3001003301023321-2210113020000222-2303021212111020-2322222030302211-3000212001031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0122022220031331-2230202133303013-1201203020101231-1010030013031222-0013203113202111-3130021331312323-1133301031133302-3200301332021312"></a>

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

<a id="canonical-0030210002121220-2021100123312212-3023201111203123-1131121103123133-3110321221022022-0322301200233333-3200323120130232-1331211103223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.static_ip

<a id="canonical-2221101331312312-1221122110223330-1003022121100230-3212300010100310-1311023010031223-1300122212130330-1013003203332133-3100023123002030"></a>

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

<a id="canonical-2030321102212013-0312120011123200-3030321033202130-3123112021321012-0020121223131110-1123121311233101-0333232113302103-2312221330202123"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.static_ip`

<a id="canonical-0032323321322120-3300333213123201-1303320323231312-2023202020222222-0123033230000212-2012003120102021-0202323331231230-0311310332333033"></a>

#### `equinix.not_managed.node_list.interface_list.static_ip.default_gw` property

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3313312100003213-1111230112233330-3210111330121230-0310310120011131-0322032310213310-2113022230021023-1130232133120022-0322223203201110"></a>

<a id="canonical-0331131010320200-0111212022011311-1032311020010211-2110223222030320-3113301112120302-3301323313120213-0322032113331310-3001102310321031"></a>

#### `equinix.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3121220120011000-0003203220011211-0130301223103300-3030133121000313-0120021110100232-1132112002201203-3301303311320023-0220203021102010"></a>

<a id="canonical-0011202201301011-1302232330030332-0133203200000113-2313300232110101-1211301321030003-1103013130232310-1133301033233320-0020100320130220"></a>

#### `equinix.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1333031120021312-0233331213013221-2202103101122103-0333301322230011-0121101130302213-1300112333031203-0320201312300210-3323001323331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2130021321001221-3322322233110033-1300133301323020-3000222302210331-3210012301131322-0011302232313011-2003000010033101-0231130233023332"></a>

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

<a id="canonical-3030222222123331-3130313003210120-3012312103300311-1333232113200222-1303303223200320-2001313121101001-2131001230010033-1002232322212330"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3233011203200203-3213210220020312-1202020032210003-1211003303230123-3333312210112301-0101132033022112-3221310030131031-2130300330122230): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3331101113220000-0113020012133023-1131101313020200-3113013122201233-0200132232022031-1331003112200300-3211201113130302-3101302333312133): complete subsection reference.

<a id="canonical-3233011203200203-3213210220020312-1202020032210003-1211003303230123-3333312210112301-0101132033022112-3221310030131031-2130300330122230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1333031120021312-0233331213013221-2202103101122103-0333301322230011-0121101130302213-1300112333031203-0320201312300210-3323001323331010)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1001023220020130-2130013232231221-2201013223110133-3001312233020303-0123213303231313-2333122001120122-2231323123300213-1032021212302000"></a>

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

<a id="canonical-1321322200331032-2010332333102311-1322100013010032-0203030320221013-1132120022123022-2313011101011020-2010203102221303-2003201303010003"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-1130223230031012-1031312132331122-1032113111020200-3301100300321012-3003200032310211-1301013120001001-3013303102301322-2301202020310210"></a>

#### `equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3331101113220000-0113020012133023-1131101313020200-3113013122201233-0200132232022031-1331003112200300-3211201113130302-3101302333312133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1333031120021312-0233331213013221-2202103101122103-0333301322230011-0121101130302213-1300112333031203-0320201312300210-3323001323331010)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2322313310010213-2200010223320131-2010212203032123-0311322323320113-3101102021231011-3230003020102130-2210002032203002-3103312211201222"></a>

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

<a id="canonical-1020120210232022-0221111333122112-3331122221202230-3232030031130123-3310000012001230-0303022202033131-2001133032232223-2201200321201120"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3100032222001203-1131311012002303-0111103002123332-3133331011220020-1131313230220123-2131201313322113-2033223002212130-0123331211000302"></a>

#### `equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1333221101023310-3032000203220002-3331230012203102-2220100302203322-0021121021312001-2112313313112112-1200303312033110-0101000012322010"></a>

<a id="canonical-0111132201000132-1330121131211023-3212333310113211-0020220003323320-2022000323030321-3123220330132213-0132021002111231-1133223000021133"></a>

#### `equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3000211103032230-3312020203200301-2203200010033322-1020230211131003-1200223212201313-1301022332301130-2010303013300122-0223121013133102"></a>

<a id="canonical-3030123222231300-3030201221103112-0312022103131021-1130311003120222-0312002033213212-0302220103331033-2102112330223102-3010130113111203"></a>

#### `equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0312211311222033-1000322033022300-2220002201230303-0001132132012310-0133313002031332-1221113303311223-3103230103003310-0000222001311023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213)
- equinix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2312332202013031-1321301001100232-0032210221131233-2103311232110232-0213110122013321-0231233320332301-0123320232311310-1211211002310230"></a>

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

<a id="canonical-1200301321212230-2010301221321023-3321223333312321-0031201223023031-3311223110122231-0300203003120312-1003200222210202-0111123210121330"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1001231033202000-2331022003300302-0203211223213021-2223121122232221-0032333223301313-2021202233313121-2200312313320203-3012202123001012"></a>

#### `equinix.not_managed.node_list.interface_list.vlan_interface.device` property

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

<a id="canonical-1231132133100203-3022230311201023-0111020121023023-2030302210232311-2132111102022333-1333123200312103-1332332211100310-0302222100231120"></a>

<a id="canonical-3100022222233200-0101022222332102-0122013012001323-3011112033231322-3112013102300332-0213200123121002-3021121202330013-2331131303011310"></a>

#### `equinix.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300101301002031-2100011030120022-0123330033310013-2200013131331122-0300010013002303-2111121113111031-2300322332300313-1313112003030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_proxy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- f5_proxy

<a id="canonical-3110032333122303-2020323022323023-0112001020332221-2332032211032313-2122121132011312-1332331102200233-1030311312120103-0103320003010200"></a>

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

<a id="canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- gcp

<a id="canonical-3223223311002203-0300112311031232-2111323121322120-3121012122233022-3003003332012003-0313031111330330-3321123233012013-0312110200331122"></a>

Type: `"single"`. Computed.

GCP Provider Type. GCP Provider Type.

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

<a id="canonical-0323222211213332-2303101112122133-1331331123002123-1310201131101313-0111203322311002-1323110011010101-3130312320021133-3330102332023220"></a>

### Direct properties for `gcp`

- [not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322): complete subsection reference.

<a id="canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- gcp.not_managed

<a id="canonical-2031113210123133-2003231132321330-2011011201220010-3002133100013233-3132331213230211-0023213131132321-0233103230003322-3112332323100000"></a>

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

<a id="canonical-2100030320321130-0110132020321110-2103211221120130-0331230033233303-2320311323031201-3302111313011223-2111222131033311-2001310112113121"></a>

### Direct properties for `gcp.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131): complete subsection reference.

<a id="canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- gcp.not_managed.node_list

<a id="canonical-1113120201332130-3013302333012123-0012001010100333-0103223030203322-3112123103003011-2011231220312023-2313101112313313-3322221203023101"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0223000301021121-0220322133123223-3312213101333301-2322330003222132-1232002230100001-3021213022110132-0101201012112133-2133332131201032"></a>

### Direct properties for `gcp.not_managed.node_list`

<a id="canonical-3312123322210032-1220203112002323-0310332133302331-2032200211211302-3131232220011230-0121303130333011-1231001013131330-0103200020002101"></a>

#### `gcp.not_managed.node_list.hostname` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023): complete subsection reference.

<a id="canonical-1031321122002321-2010030212132211-2020013231131203-1320000022300320-1110033210130201-3023003003200313-0023200130210221-1120320302233102"></a>

<a id="canonical-0132030111222313-1310201032303030-3013200122321212-2200332122321130-2303310113132131-1133030002120332-3210300213023102-0123233021011332"></a>

#### `gcp.not_managed.node_list.public_ip` property

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3213321100132233-2032322232203303-2212323321032022-3023112010133001-1310122101221212-1100303023313020-0300301203332001-2132032302032112"></a>

<a id="canonical-2313122012303101-2321211313101211-2202320320020303-3103011302001313-1101102323312031-0331302122113110-3001011100311201-3331122032110022"></a>

#### `gcp.not_managed.node_list.type` property

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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- gcp.not_managed.node_list.interface_list

<a id="canonical-0012003100113300-2112020221003031-2313113113110210-3300231300022010-2321333122011022-2312231303130300-0230332332030213-3023113020013222"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023310000321032-3222001212110031-0233323230030313-3301013221212001-2101202011100213-2303103321103310-2200333031132031-2202202223212032"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0313213013000020-3212233120330022-0022032310201302-1200122220033331-0200232022020130-1001032201301022-1122123222223321-3213110320331121): complete subsection reference.

<a id="canonical-2102130232120001-1122201000010213-2223310302223121-3233111313313232-2310122131221303-2120100030232301-1022003032031021-1110013211033232"></a>

<a id="canonical-2330332020113032-0002321123303230-2312021323231213-3212330113210333-2133113001103121-3031213302223333-3313123000021212-0002221012212013"></a>

#### `gcp.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2313113122303132-2031000233221130-2300100111113033-1210001112100130-1330002131313031-0021000101233322-2001202132031133-0000210231302222): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0331121103102220-0200021313232020-0300210020121110-1113121123231130-3103003330213222-1200032230223213-3011321232003203-1100030113301203): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3131100232112021-3320222221313302-0030030333211322-1313203203122230-2210131002031131-0120332000231110-2311231030102020-3323010302113203): complete subsection reference.

<a id="canonical-2112320203233132-1301321223010113-3130011130122031-0212311232011003-1102011322201103-2121001101201001-3301230331201232-1131022213023021"></a>

<a id="canonical-2021201031320202-2111221022210331-1010232131113312-1301011131323330-1333202033301103-0233330010121323-1130312333331020-2332323313221122"></a>

#### `gcp.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2012131221131033-2310233103101332-1333100031121210-0012002122300200-3000032130020102-0100122123003230-3200033322032120-1121330320110132"></a>

<a id="canonical-0223310131012110-2132031220100112-0231021033111110-1002000031023013-0133211200030312-1200013113022123-1322011211033202-3103311010311310"></a>

#### `gcp.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3220001300332230-2312101021200003-1102300032002112-3212321233302022-1310100323201322-1201112013222212-2121323220211211-1131310320333122"></a>

<a id="canonical-0123233000030030-2133330332300232-2223020233221311-1012113130301312-0133211322131002-0000222311013222-0213200113311023-3030312002113120"></a>

#### `gcp.not_managed.node_list.interface_list.labels` property

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

- [monitor](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3330020203333123-2000110321011100-2320111210323023-1111102102310012-2211012130203010-1202113121232333-3021020103011231-3210211020020110): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-010.md#canonical-2302100330011231-2331130202323002-1323200333331223-3032112122102221-0222113213231231-3232211022020012-0220302012131002-1310113120101230): complete subsection reference.

<a id="canonical-3011220023333330-0101332212132013-3021232212200110-2311230220332230-3230330220102212-1003111313333313-1221313022300022-1302002030202322"></a>

<a id="canonical-2113203022320033-3013022002230331-0031131310012323-1130121231223221-0110303332221102-3012311313103311-0220033112213220-2031123102100313"></a>

#### `gcp.not_managed.node_list.interface_list.mtu` property

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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-2001030211133231-1233203212112220-0230020031232231-2321010132321003-1001220121223011-3031323322112033-2103330100031011-1201301321320121"></a>

<a id="canonical-2233031113223002-0310301003332012-2032313012112332-2020322331003131-3020222322000200-2200201201312131-3310331030303123-0132302322111121"></a>

#### `gcp.not_managed.node_list.interface_list.name` property

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3113103003220021-2132203222033221-3311010023111202-2110021332101233-2133110331322213-1023313321233313-2002332332120130-0312320112010121): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-010.md#canonical-2112023321131113-3233101211332230-0132220011002322-1011133103020301-3110003120032311-3210232230313132-2100013113213313-0301010110031100): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1302210203130003-1313223320200122-2331122322113321-2301221303121211-2322032132213332-0333013031132120-3233000200330302-2001122111010303): complete subsection reference.

<a id="canonical-0101020202303301-3210230320023202-1022313123120033-0213031111320131-3330213303312111-0000000030222010-3333220322103123-1200202030202230"></a>

<a id="canonical-3021110222102010-3102122001000330-2231023223003030-1312321221023122-2030302330203321-1031313302112203-1132011101230301-2010222323221112"></a>

#### `gcp.not_managed.node_list.interface_list.priority` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010103133231021-0321113101113112-0200033221200201-1213231202202211-3003132202111230-1000212102123000-1020003131202322-1320333021011123): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3310211011233213-2103000113120311-1000113103201310-1312103300102323-0220220033100123-1220112221302312-1003112012323201-3030301002303230): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0210101222022030-1123133310231331-2112031011033320-2221300200103131-2322333202321102-1233310101232213-3313232221212001-1332101300222002): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0322012311312010-1323301303002130-3212233121032001-3102320012012300-2000110102100002-2031002123020303-0312033132011311-2011122233321022): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3200123333213120-0333112130222233-3130303210201320-2312023000112100-2001010330310222-3101030232221200-1223001110120032-0201320101012013): complete subsection reference.

<a id="canonical-0313213013000020-3212233120330022-0022032310201302-1200122220033331-0200232022020130-1001032201301022-1122123222223321-3213110320331121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- gcp.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2023303001300112-1003132303113120-2212131312131113-2323331230013213-3332001012002001-2030331113002203-1233130131322221-1032030011310230"></a>

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

<a id="canonical-3123323111023002-3230301123032001-3210330310011223-2312331132100122-3301323122301002-0202013133012003-1110011321221023-3220121310100320"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0333332331103310-2332002120132202-1211233301313023-2030323120003210-1132131003013132-3332011112113232-2003002223101000-1101013320022020): complete subsection reference.

<a id="canonical-0000000022333211-0001232000101220-3020302013121112-0222030113212101-1211022322213200-2123200313302123-1310331132333301-1221022130021322"></a>

<a id="canonical-3201322113310201-3333222011010323-1312310122111123-1112313023222230-0101210101300121-3230121231232330-1212211323133133-1310311302201213"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3122310201111311-3031030132132122-3303330320002122-0200302013132231-3020221110323301-2110111322012311-0001303131320203-2302131131211112): complete subsection reference.

<a id="canonical-1113311133020110-1320133102003120-1231223333332110-2130113013232312-1313030111022131-3120131033010113-1110013133033233-1223232312122211"></a>

<a id="canonical-2233120112011212-1120330130033133-2221312322132130-1203322312331101-1011131222313121-2322220233333032-0311033023020120-1210323011310102"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210231310321212-2303320321210330-1321012101332322-1000230023002223-2102231001330031-2133100031030200-2113321122300300-1321011222022132"></a>

<a id="canonical-3032013213330310-2113011033022310-2223012133302013-0101323100203303-0212121000102310-0312202200021201-0033222030120323-1010030222001203"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1220003302032333-1310103320211113-3221000300333220-3320031101100233-0023210213002321-0102321310111321-0323211303021001-1131103012120023"></a>

<a id="canonical-1230123330310021-0133233211101202-3021031032100122-0031200322202113-3311010002321020-0223203102020313-0031123032022221-0323320202010233"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-0333332331103310-2332002120132202-1211233301313023-2030323120003210-1132131003013132-3332011112113232-2003002223101000-1101013320022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0313213013000020-3212233120330022-0022032310201302-1200122220033331-0200232022020130-1001032201301022-1122123222223321-3213110320331121)
- gcp.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1333212202132012-3203001110030120-1111222221010201-0311231033132013-3302000211021130-3320002121120031-2231020231300233-1032311301130023"></a>

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

<a id="canonical-3122310201111311-3031030132132122-3303330320002122-0200302013132231-3020221110323301-2110111322012311-0001303131320203-2302131131211112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0313213013000020-3212233120330022-0022032310201302-1200122220033331-0200232022020130-1001032201301022-1122123222223321-3213110320331121)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0001000202210330-3123100200021313-0232211213131031-2230131030013213-3210023223200330-1100131213020031-2030300013331100-3203012031211323"></a>

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

<a id="canonical-3303301132111231-1033001120310232-1200220211132110-1323302303331203-0030111022311231-3212122231231010-0320222020011132-3200023111332032"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-0201211312120301-2332233322331312-0321000203032311-3130132321310202-0200010122002012-3021030321221231-1011210332301322-3030222201222333"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313113122303132-2031000233221130-2300100111113033-1210001112100130-1330002131313031-0021000101233322-2001202132031133-0000210231302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- gcp.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2212231333213301-2122012123311111-3223301303311323-1302030121311302-0101112012201001-2203013333030110-0323031032310020-3133133101323000"></a>

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

<a id="canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- gcp.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3213232302113111-1201113311023103-0301211210110220-0203123323012212-3111300133222300-3033323002021300-3213022200122002-2333120103210010"></a>

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

<a id="canonical-0130111132013211-1020102022302201-3221020011221213-3122011320320212-3222132311002222-0321012321330333-2113233102203132-3212000203132122"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3201333011331103-0000231130230223-0102222021031211-3211011033200213-2123031031102023-1333113021110331-2300031030033110-0333020121300203): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0021312211011122-3122111203113301-1230321110211301-0002303210203020-1021311331311202-0103200321311323-3331001220021322-0023132300111001): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1221200002320313-0333310020000023-2213020013121221-0303213213231232-1012112000122330-2003131133201233-1211200232320100-1230033313103310): complete subsection reference.

<a id="canonical-0100222321301111-2023232011031321-1232022003311223-0200220303332231-0013000200111323-3220123100222002-0123031203010223-0220131231312031"></a>

<a id="canonical-3313231103231210-1121332300332233-1220303330003121-1210000130323311-1310032102333011-0023302332221003-1121202003202131-3033121013301110"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-2111321222131232-2330120210222300-2103312133220123-1100031001333003-0010330310012330-1321133001302100-1131310110312201-3203211301003311"></a>

<a id="canonical-1332301320333001-2333232222033331-1130031302333130-0231313203131302-0303102111002200-3212010300022213-2013223232010020-2213321033323013"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1000312330122200-3311320010221333-2123323121100000-3321333013122320-0101230323003030-3121332001013232-0213223022103220-3000323231032333): complete subsection reference.

<a id="canonical-3201333011331103-0000231130230223-0102222021031211-3211011033200213-2123031031102023-1333113021110331-2300031030033110-0333020121300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1132101113300323-0013323230111300-3203200223021212-3102020000211030-2233201232022232-0223013022203000-1000131320211122-3121121232010321"></a>

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

<a id="canonical-0021312211011122-3122111203113301-1230321110211301-0002303210203020-1021311331311202-0103200321311323-3331001220021322-0023132300111001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-2332023102332003-3012033313213113-3011020233020302-3222031123201030-0300112202220203-2003210101321223-2110101212112002-0122121010331212"></a>

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

<a id="canonical-1221200002320313-0333310020000023-2213020013121221-0303213213231232-1012112000122330-2003131133201233-1211200232320100-1230033313103310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2011111201103102-0201132210121232-1101022303100300-1300333102231133-0011020332031100-1222031212220200-1201030200101103-3013132212312101"></a>

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

<a id="canonical-1100211032031003-1221301012321132-3122111003130033-0310330123113121-1021212102212022-2110031220232320-1101221231322010-1300310002202220"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-0123221000110010-3022303230222220-0232322201103103-1102210200310331-2202010233113120-1333313300301120-1123132023303000-2020303233032233"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2100132100323223-3021113332212312-2030201033013333-0010010033232122-2332113122322310-0010130001030111-0023210322213201-3002222021321302"></a>

<a id="canonical-1101031102100231-2123021320201201-0200121101033320-1232201003230010-3321001133112332-0333202212133331-1330131212112032-0200321021233032"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2031300211130201-1230003130100312-2303313330131023-0231202222210122-1012101103301120-2013023310203331-3313301310132013-1321210103312111): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1311201301223030-3201102120220223-0320332201303130-0102231201012223-1320320331320210-3232122033112231-3221033310122010-1232332123222321): complete subsection reference.

<a id="canonical-1303303330110231-3233121221201300-2201223212022321-3013121133223202-2023123220131131-2130133031130200-1102122101301323-2021102310220113"></a>

<a id="canonical-3303303021012121-2012322033312123-1021203130231002-3031131201023131-2120030110001121-0331000132110230-0323222010133332-0003202011233201"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3000222113213021-2323203200320133-2331220210201122-3321003203023102-1320103003023033-3212321120131123-2011200113330320-1333100203302133"></a>

<a id="canonical-3113230003220312-1333121112032121-1300312110202230-0210303331301231-1002103213312311-3302203101222003-3302131231210211-1223231321130010"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-010.md#canonical-2233223022010012-0011312111313323-2031211100023100-1231231002321000-1103030021010020-2222023003023321-1122200312000011-3203002303221120): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1231233010231123-1010111110301201-2202221322012022-3122020210113030-3330312200010033-0232300320011130-3232201321200110-2003230301322311): complete subsection reference.

<a id="canonical-2031300211130201-1230003130100312-2303313330131023-0231202222210122-1012101103301120-2013023310203331-3313301310132013-1321210103312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321)
- [gcp.not_managed](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2012131332232111-2113201200331003-1221123131103130-1302230113011200-0013301130220123-2032031013003232-2132320102233210-0031111123012322)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0022330113111022-3233132311302032-2122100301302032-2232023203323331-1103222110203020-2122120333120112-1332330120002321-3223033213101131)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2333220033211201-1212213220202210-0021020220322033-0113011200210001-2210202331222320-2200002310202001-1011110303010332-3303030213312023)
- [gcp.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0131222023132311-0331103101113321-2222211303031301-3210003321211320-2103231200020212-0023230011003131-0023311131223123-1300133021010021)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1221200002320313-0333310020000023-2213020013121221-0303213213231232-1012112000122330-2003131133201233-1211200232320100-1230033313103310)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2232103120122023-0303122103122031-1131103110102300-2023020102010102-3123333300012013-1313010013102030-2333131003012231-3122211032213121"></a>

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
