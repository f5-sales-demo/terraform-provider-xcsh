---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- vmware.not_managed.node_list.interface_list

<a id="canonical-0310021211332003-0313231311233220-2031032300012221-0202032331000311-1203001010000022-0103323123002001-3123133100310310-0203301322211133"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

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

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303123011320312-2012102032323130-3303001110212131-3010000320303012-0333001022310222-3301110123213113-2213211010002212-2103000111233031"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300): complete subsection reference.

<a id="canonical-2203300133132012-2213003213211030-3101330111222220-1102120023231130-3013202030002211-3133010202033012-0002210333210312-2112202231112000"></a>

<a id="canonical-1121033222131330-3230031233111022-1130200000321220-2103033330111321-1200133323203013-0101023203031120-1133331030312211-0132021130023231"></a>

#### `vmware.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-018.md#canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333): complete subsection reference.

<a id="canonical-2112313301122330-3232030220303032-0330123023311113-2311122302120132-1222022020320303-2010111300030111-1012020010301121-1222112330310201"></a>

<a id="canonical-1313023333101311-2203032120300322-0222313101232001-1301302310321322-3120230032031221-0332012022122102-2322013222303123-3233233123100203"></a>

#### `vmware.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3033333200311010-1312230121231033-3132123323220120-3222133233200203-1200331001333220-1021021310321102-2120220310120213-3133122121213313"></a>

<a id="canonical-2113312122003200-1320112001213303-0023010031101310-1231010302203130-3131031321023010-0130033120121212-3213333323122131-3000111000213213"></a>

#### `vmware.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3221212132330031-1020002003020123-3031002313001132-1031302311203112-2123331022001002-1020330030220120-3320031212211320-3011332022012123"></a>

<a id="canonical-2120303100223223-2201101003032130-3303021212222110-3202003210010020-2310010102331221-1121103220300302-2300103013020020-3101322330213021"></a>

#### `vmware.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

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

- [monitor](resources--securemesh_site_v2--reference--group-018.md#canonical-2123002311101023-1232110121333102-3302333100103032-2301311000103120-3001332221330132-0132113300211131-0012320300112121-1102023200322330): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130): complete subsection reference.

<a id="canonical-1030202202133032-3203222321301022-0212332031103021-1212231210013302-0223202313030221-0201230123011002-1113003011101212-2121322312212123"></a>

<a id="canonical-2201122212021200-3023333230103301-0200102202030211-1023330131332020-1001303321013001-0003322310231000-3100022002111130-0330001132330202"></a>

#### `vmware.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
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

<a id="canonical-2100322201123332-3133122032321203-2030200031310002-3020213030233102-3201213113010321-2201133301122010-0220021122111101-0103330131230222"></a>

<a id="canonical-3313330320213333-2031223332231123-3113201013311212-2031013103103220-2320132030002310-2323011330233011-1302311211033302-1222200031322221"></a>

#### `vmware.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312): complete subsection reference.

<a id="canonical-1231300102003220-3223200311023233-1013123111100101-3313333311013023-1032232202313123-0332301030311123-3000112123133100-2130311221023300"></a>

<a id="canonical-1320001301231202-2021232311323300-2230121012311013-0312311211212003-3022101100111011-2201031010333100-3302202300020103-0330100032000203"></a>

#### `vmware.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-019.md#canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-019.md#canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-019.md#canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021): complete subsection reference.

<a id="canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3111200033322213-2211012132301320-0110011311302321-1120301223032321-0103332211010303-3000003232132123-2021033131031320-1200223003132320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311320131023312-0003212232001110-2202310032101221-0002311222303111-2002012000320011-2212222200130210-2230112210102103-1332002302000022"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-018.md#canonical-2332013311302110-0303120001131020-2132100220020213-2003120022230322-3023213123003311-3011000011203113-2121012131022133-3002312320202120): complete subsection reference.

<a id="canonical-2221102121231220-2030021311212302-2130312330031120-1112012333020020-2020300222001221-2030303101030320-1231322231123022-3033010312300312"></a>

<a id="canonical-0111321332332001-0010032110112023-2130033233312100-1021222231202021-0033003311101133-3123002013221322-3120100231000322-1333110002123131"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](resources--securemesh_site_v2--reference--group-018.md#canonical-1130133021200101-1213230111003311-2121203202030322-1112203302120332-0022323021000213-1300200202211230-3223300321332022-1323111303103210): complete subsection reference.

<a id="canonical-2223023321300321-2120212002033211-1103302231330333-3231331112012203-3333223311303110-0220023020111022-3102132033023213-1322320110221123"></a>

<a id="canonical-0330313220031333-3211200213223023-3331322320030023-3302211123231113-2301202220213010-0233213002131112-0021033321120132-2210313001320013"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-3221233312211021-2231213122112203-2233101332010210-1323132020310320-1112301221031320-2332233231200032-3001031311012231-3112310310030020"></a>

<a id="canonical-3131232203322122-3201003121310020-2333300333212330-0302221022222303-1313320331102321-2031333311030020-1323130132111000-3231310230023301"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-0002212202122230-0022020222222331-2302320322110333-1332120033110010-1300310311233030-1321322001010203-1221221031033302-1311230021130032"></a>

<a id="canonical-1013202210322301-1232301113212133-0000132123211021-3232020212322102-3200130302232333-0102303333122111-0212112201130231-3000031022322001"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-2332013311302110-0303120001131020-2132100220020213-2003120022230322-3023213123003311-3011000011203113-2121012131022133-3002312320202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- vmware.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-3023223301230030-3032111130021220-2203223220123333-1211223310022000-0123313232113113-0212023003310110-3320122133103301-3021130221002103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
active_backup = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130133021200101-1213230111003311-2121203202030322-1112203302120332-0022323021000213-1300200202211230-3223300321332022-1323111303103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- vmware.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1301200233121323-3323220001133031-3311232211112210-3310303013331202-1021330221021011-3210311133330223-3002213313230021-0022232220312013"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001233103101311-0323231120332202-2303012101011313-1011332022330122-0323202001310001-0122000211301310-3221013230200202-1230300102232022"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-1112001100113111-3033230332230101-2101003103310232-0321232323112022-0021311103021333-2233021231003331-0012032211332232-0230110123203333"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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

<a id="canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2223123213031103-1032032333131022-1102213000013112-2000011030000200-0232011123231323-2113130100001112-2122111001132231-2011001320030002"></a>

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
dhcp_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3010122310011333-3211000133301011-0120133123122002-2032001201023013-1133302223113122-3332100311131133-1211230002201013-2130123122132201"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203333311101033-0101203323011131-3223313212330103-3012001103321113-1232212203222230-2001020310203213-1211001301030332-3132112200132133"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-2121311232321012-1033303033320013-3101200321110122-2302122010321320-3021033100120132-2333130310020301-2312323232131312-3333022103112321): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-1032230102021201-2102320121102213-1001212032323203-0331131012113112-1110233121012123-2320011313300010-2031023122000220-3213122213320010): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213): complete subsection reference.

<a id="canonical-3221203211001030-1202123020312123-3032332033031031-2122300013211122-3000103133132023-0101213203112021-2220221321203131-2111323321301200"></a>

<a id="canonical-0321020321220211-2132333100100010-3200313110021020-3330202203110200-3120311132002230-2330023011121223-3301221122313202-3231323110013010"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3221221331102023-1300120302033321-1100300102011303-3000133201102102-3323310132122223-0122330232333213-1321330030223021-0133122333123032"></a>

<a id="canonical-2320001000311310-2010123112113232-2023313033012102-2102013311321332-1102213300201230-3023111022213322-3103303033330001-3210200112111323"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-1210321213333210-3220323122002223-3010222323311033-1330113030010301-2221221230300312-2002110033223201-0113030323120331-2222300011331332): complete subsection reference.

<a id="canonical-2121311232321012-1033303033320013-3101200321110122-2302122010321320-3021033100120132-2333130310020301-2312323232131312-3333022103112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1030211011313312-0301102033031103-0302211120121303-1022223301332000-1231320003230123-2202123102312030-2123312221131220-3333100123321332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032230102021201-2102320121102213-1001212032323203-0331131012113112-1110233121012123-2320011313300010-2031023122000220-3213122213320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3320323221311102-1000132300002130-1033113313232102-2103010312332112-0220101300012302-2221231000001322-0233230300321322-2320120311121221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-0112121133232130-3203120100132130-0031310130012200-0332312021123032-2003212130102323-3232213000110201-0003103312220230-2022313120103233"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011122133303232-2111231223333223-0112231312012021-0002330203223030-2011103131112022-2002233001021203-3111220310121032-1013121313031023"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-0022113102010222-0030220103322110-0012303110302133-2222210030213102-1000200100111232-1131210220311323-2322221213013001-0330032220131012"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0223330132033320-3021331233331012-1012233020230233-3210313121333210-0211112332210200-1330031232112012-2332103000311320-1122100103002220"></a>

<a id="canonical-3212311331103311-3332122100211312-1230211012203202-2303120200002322-1131001311101210-3102121102331210-3130130231013221-1203211002223113"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1200322232101113-3020133013201013-0112311111232012-2111001220030230-3303330332002220-3221032013000000-3110102221230132-0320133220101123): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1302230323220113-2220221013331112-3233102001330113-2301131201333230-0220320011123212-2212332010010212-0222033011110312-1020101232330221): complete subsection reference.

<a id="canonical-3012211223303303-1123332331021323-2211030020103331-0230301311200003-1031232100102013-3310033232320232-3213030020102223-3232213301232011"></a>

<a id="canonical-0223323231000233-1033030211232131-0100232020320333-2311002021301012-2300320333030133-0031331030220120-1122013021132101-2113131333211020"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-0200102012301121-3323030300322130-2202000002221123-0002232131023211-0112300031102330-1100203100302132-2000211101300223-2201112132310011"></a>

<a id="canonical-0303213200301323-2233022032200203-2010033120332332-3103102230033130-1103121103213203-3302233202323012-2013120231303313-1133220210110020"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [pools](resources--securemesh_site_v2--reference--group-018.md#canonical-3231123123331121-0230320211331003-2120101323231002-0100301112301133-2000232010212100-3233201013301123-1312102323103301-3033100213230202): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-018.md#canonical-0300311323132320-1112321031303221-0100011310132133-2333201333110232-2333300312233021-3133231120012032-2220110221223331-2001330300200121): complete subsection reference.

<a id="canonical-1200322232101113-3020133013201013-0112311111232012-2111001220030230-3303330332002220-3221032013000000-3110102221230132-0320133220101123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-3031023213110303-3133330023200313-0003030312131011-0231122120301103-1300220320131103-1232010232032323-0310011311133132-2012123201010331"></a>

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
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302230323220113-2220221013331112-3233102001330113-2301131201333230-0220320011123212-2212332010010212-0222033011110312-1020101232330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-3311333003222300-2211313331212102-0123221233130322-2330033213321230-3203332210221232-1022103012212223-3121332132103202-0322010132313003"></a>

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
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231123123331121-0230320211331003-2120101323231002-0100301112301133-2000232010212100-3233201013301123-1312102323103301-3033100213230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-0322212111323323-1111113332331332-3233020220200323-3333201210231023-2213132212133333-2032301100211232-2223233132311332-3130012011303103"></a>

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222011120311002-2322210311012023-3113101103302233-2212002211223203-3132101220212233-3212123212001010-2030213032111133-3122123022311323"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3121131303220311-3230201212120110-0301003101212210-0002003232202022-1323123212210131-0200231002002020-3320120220333201-2032121122333020"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3331012002132030-3131203332313133-1332231330310202-3300003111303222-1031032033310103-2033203333203200-2310221211132320-1203020101301133"></a>

<a id="canonical-3211323212101112-0011210220122321-3221320200310232-2231333222300212-1032011032123220-3131101222221001-2000233002012222-0012233201133123"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0331122123032131-3213021033211333-0032202032020000-3011011111033021-2300000300230023-0130023033313013-1020231010331332-1230002310113230"></a>

<a id="canonical-0320232113323021-3110331203031100-2003213321000010-2330121110111320-1120112101113320-1331213330120310-0320212231111232-2013110200303330"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0300311323132320-1112321031303221-0100011310132133-2333201333110232-2333300312233021-3133231120012032-2220110221223331-2001330300200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0333303210330303-3031312321232322-0031223221232012-3310111023103322-0002032223120213-3303311312030220-1310031210200010-0211000120131313"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210321213333210-3220323122002223-3010222323311033-1330113030010301-2221221230300312-2002110033223201-0113030323120331-2222300011331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0032231133230011-3323300112213100-1301032231123300-2330132213012112-2113023003223121-3330202222000133-2303110111020010-2122201002110020"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112303132010323-0120212022001333-0013231133203320-3121023211202323-0222032311202012-3031101123101112-1001203011302121-2020330203031223"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-3111202331212212-2203313111301200-0232322111331323-3211021303113212-2002223320000130-1103321130220302-2202320013023011-2213301110302113"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3320220232003000-0001122312131223-2323110233332213-0012211031113100-0211111132012023-3031310101000301-0031003121021333-3323202122012101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330023302201013-1111121332221031-3331333312112111-2120032222303023-2010232122333230-1331212112011112-2033311311101200-0120202003132010"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-0212031310221121-3332210103021101-1202321313320312-0131101230101131-0222202321323202-1111332321332222-1113033132113332-1213222200113220"></a>

#### `vmware.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1233123211113312-2333003220131210-2121010100302111-0022100233330302-1133303221203221-3332323330102231-3033302221100323-3110321132013011"></a>

<a id="canonical-3322002030212333-2023013021310001-2211002122203113-3323211110110032-0121233131032202-1301210133013332-2012020332330023-0323011222123313"></a>

#### `vmware.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

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

<a id="canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0101213030122331-3311233210133033-1213101221113012-1001200320300132-2222031331120131-2321321023210213-2301122200130011-3303130120300022"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2232321110131231-3321303301303232-0030120130301123-3220013330133302-2310203310213331-2322133032002103-2002003303020200-1212311230232211"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-018.md#canonical-0023101013331032-0013303312303222-0330102322021012-3212212212000130-2332112031132012-3313102122102330-2123203103230310-3200300202210121): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020): complete subsection reference.

<a id="canonical-0023101013331032-0013303312303222-0330102322021012-3212212212000130-2332112031132012-3313102122102330-2123203103230310-3200300202210121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1120002001312022-2322223112022321-1102120013110131-0331202132311132-2332113333032131-3312312103212032-2330321202032003-1130001111303303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
host = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2003023332230030-0222010123330100-2201012302123110-2030020020112300-0010213320210212-1332210102123312-1202301013322220-3321201231033131"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0011223230232202-0311011030212311-1333231302232110-3201333111010320-1030213100221223-0131221202030112-0231101321301303-3320003012331110"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321): complete subsection reference.

<a id="canonical-1121132131203003-2220221032211223-3110131033102003-3010001310100231-0033300103033123-1201113202120011-0121200132032221-0313120120100112"></a>

<a id="canonical-0033131301333013-2103233000103310-3220021201111021-2210331331233022-3233303323101001-2011002023233330-3332313232311232-2103230121300220"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232): complete subsection reference.

<a id="canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3222232220003031-1013012112100023-0220010311002110-3123330120300101-3021000212023310-2200223110103020-2023101212120023-1002201212133122"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1001121232320231-0010131203121313-0033231120301112-0231023201133103-3222200131001322-0002302131132313-0312232020303332-2311102200100112"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-018.md#canonical-3333310230323113-0321102303321121-3020312331122002-0202033202222112-1013331320313013-1222011302110313-2011012130122113-3200033232111223): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310): complete subsection reference.

<a id="canonical-3333310230323113-0321102303321121-3020312331122002-0202033202222112-1013331320313013-1222011302110313-2011012130122113-3200033232111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3233211323102013-3200321330112031-2210132023221210-3011331120000000-0023311203330122-2102033023300220-3332302313023020-2113221110003211"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0323311332323321-1013322120022012-1112110211113333-1333000201131203-1312012312323331-2121221330121323-2131110100211332-2110231100201222"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1232233211222010-1233103231212212-2231001221112123-1323232120201300-0101331123113223-0120012323200021-1312200113010133-0031210310130033"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2332210323030101-1131320101312333-1112220121011313-3331002013110300-1120010032033310-3002310321220103-1202323311013303-3121100132210213"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0330121010033121-3112102000002311-2310302233121120-0230031133111133-1120212203312132-0031311003030213-0330221100112320-1122131031312121"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2311231113321120-0012133031212330-2201233013322023-1033101020022120-2302111032132122-3113213303021210-3223211030202332-3033223023322103"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-2331130003302210-1130032133232031-0212033203322032-2320030201313313-2203301200113221-1130011232322232-3322101231333130-0320320122133011): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1002302130301300-1111332313310230-0013031100103102-2033031220320330-0302030000303120-3122302302122032-0120203322300101-2130112100132131): complete subsection reference.

<a id="canonical-2331130003302210-1130032133232031-0212033203322032-2320030201313313-2203301200113221-1130011232322232-3322101231333130-0320320122133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1213111120021303-1111030113100002-0322003331031102-3133121112310002-3333230321332021-1003202311112000-3013101031123123-0103330013123133"></a>

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
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002302130301300-1111332313310230-0013031100103102-2033031220320330-0302030000303120-3122302302122032-0120203322300101-2130112100132131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1133002010130331-1212213020102221-3111321231201213-3121133010201033-1110312300203001-3310031300303201-0220202230300031-1332133330121100"></a>

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
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2323220001222210-2320320013112032-1302032200300232-0210011213020230-3112322033001232-1123301312220223-1021231310022113-2213021221213011"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2111312112313303-0322211020200313-0120002100012002-3233222222303022-0113210002331302-2223310102032322-2100320023113223-0222031333220220"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-0030023302321332-3013200013110130-0103000130321213-3232013112030331-3000122213312310-0023033212200322-1330331212032023-0003312321111222): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-2323211033000320-3021122213122222-3323332032130302-3233201331133201-2032333230122333-0003220123213302-2222112011011012-0310330322303020): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102): complete subsection reference.

<a id="canonical-2101102011312223-1110211013003201-3320022232213331-2321322121110001-1323210121033331-1000120211332230-0311031132102123-2213210033300000"></a>

<a id="canonical-3311123201313302-3021103233123021-1133130333102102-3322301121202310-2022031121203210-3110011000200110-0333300100232213-1203112200113102"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-2132002102122130-2322212311223333-2103331033221033-1033131120110320-1313322110200323-2303011212232133-0110120210210332-0031110122011102): complete subsection reference.

<a id="canonical-0030023302321332-3013200013110130-0103000130321213-3232013112030331-3000122213312310-0023033212200322-1330331212032023-0003312321111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1113133123012200-3320201323023031-3123221223200111-3102211020130022-0020131102030203-0203031330212202-2023002113230312-1010232213232302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323211033000320-3021122213122222-3323332032130302-3233201331133201-2032333230122333-0003220123213302-2222112011011012-0310330322303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0120300000101302-3202133002113032-1132011002112021-2312132223313100-2302110033210300-2203130331000003-3313033002211010-2301212332132302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0113111033113303-2011101301222322-3321133232203002-0330200003033231-3103032020000100-3323330100133011-0223013100120233-3322320212320333"></a>

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312022323201022-3311120023100313-0030031131001313-3300312011133110-2000211230201002-2132203321100002-3132213212130120-3230132232023300"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-3022212112002233-1233220132110130-1312132333123303-2222000120321012-3301320333212303-1300323201222103-0010110120110322-1010211221232220"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-1101131220211003-0010032112210231-0301322121332101-1030203011333012-3012300333332131-3310221322023031-2213320033210133-0132000013201121"></a>

<a id="canonical-2000121033332321-2330303201202231-0022010310033130-1231102131013200-1023200011033030-3101320321030202-2131331122101233-3112003231033233"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [pools](resources--securemesh_site_v2--reference--group-018.md#canonical-0303103101121101-3312332011211321-0110103313100020-3313101203001033-1020221110221100-0332211310113332-3032110310321033-3302021210300122): complete subsection reference.

<a id="canonical-0303103101121101-3312332011211321-0110103313100020-3313101203001033-1020221110221100-0332211310113332-3032110310321033-3302021210300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1200102132212230-0000130232332102-3211301323102210-0210203330330312-1000320223122022-2301210202332010-0302010310131103-0132222001030310"></a>

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330120120130113-2301032103112201-0233033103033231-0200123111011121-3232213321210302-1010132101311032-2322123003200111-2012332010001310"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-2121232203312221-1331321012303132-3100230301223223-2200032200303211-0023020211012003-3001001310211133-3131211203312030-3011311302231133"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0311022311231112-2010002120212031-0300103103000322-3221203002210001-1113312222201331-3323312103231112-0033302213231101-3031003131201020"></a>

<a id="canonical-2233011003131021-2301303032300000-1202230103001232-1331220312022001-1033312302122233-0313231213223000-1033033203303103-2012310323032100"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2132002102122130-2322212311223333-2103331033221033-1033131120110320-1313322110200323-2303011212232133-0110120210210332-0031110122011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3111221223013100-2123120032122112-0313220032122300-3231123102332233-3130130332221013-2233020022021012-0320022002211312-3010023113110003"></a>

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

<a id="canonical-1031200322122033-1122033301233000-0013232102322131-1002332322033112-0232203023002313-3303311013332023-3000111223113331-0011111320210022"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2113131002221113-1132231120233121-3223310010101012-1112113123111101-0121201121101132-1333210010113232-2011301222330312-2311330320303033"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2123002311101023-1232110121333102-3302333100103032-2301311000103120-3001332221330132-0132113300211131-0012320300112121-1102023200322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.monitor

<a id="canonical-1112320033111031-0222023033002122-0122230031203010-0030213010203032-2303200032133030-0303333112100111-3002101002011113-1003332211231331"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1100012213300002-0021010020302331-0103220002022303-2203112110220232-3010232302211203-0122000310123232-1013032012230322-1103203011212332"></a>

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
monitor_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
