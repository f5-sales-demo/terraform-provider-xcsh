---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0303123011320312-2012102032323130-3303001110212131-3010000320303012-0333001022310222-3301110123213113-2213211010002212-2103000111233031"></a>

## vmware.not_managed.node_list.interface_list — interface_list / 022001130230 / 2

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121033222131330-3230031233111022-1130200000321220-2103033330111321-1200133323203013-0101023203031120-1133331030312211-0132021130023231"></a>

## Direct properties — interface_list / 022001130230 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300): complete subsection reference.

<a id="canonical-2203300133132012-2213003213211030-3101330111222220-1102120023231130-3013202030002211-3133010202033012-0002210333210312-2112202231112000"></a>

<a id="canonical-1313023333101311-2203032120300322-0222313101232001-1301302310321322-3120230032031221-0332012022122102-2322013222303123-3233233123100203"></a>

## description_spec property — interface_list / 022001130230 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-018.md#canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333): complete subsection reference.

<a id="canonical-2112313301122330-3232030220303032-0330123023311113-2311122302120132-1222022020320303-2010111300030111-1012020010301121-1222112330310201"></a>

<a id="canonical-2113312122003200-1320112001213303-0023010031101310-1231010302203130-3131031321023010-0130033120121212-3213333323122131-3000111000213213"></a>

## is_management property — interface_list / 022001130230 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3033333200311010-1312230121231033-3132123323220120-3222133233200203-1200331001333220-1021021310321102-2120220310120213-3133122121213313"></a>

<a id="canonical-2120303100223223-2201101003032130-3303021212222110-3202003210010020-2310010102331221-1121103220300302-2300103013020020-3101322330213021"></a>

## is_primary property — interface_list / 022001130230 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3221212132330031-1020002003020123-3031002313001132-1031302311203112-2123331022001002-1020330030220120-3320031212211320-3011332022012123"></a>

<a id="canonical-2201122212021200-3023333230103301-0200102202030211-1023330131332020-1001303321013001-0003322310231000-3100022002111130-0330001132330202"></a>

## labels property — interface_list / 022001130230 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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

<a id="canonical-3313330320213333-2031223332231123-3113201013311212-2031013103103220-2320132030002310-2323011330233011-1302311211033302-1222200031322221"></a>

## mtu property — interface_list / 022001130230 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1320001301231202-2021232311323300-2230121012311013-0312311211212003-3022101100111011-2201031010333100-3302202300020103-0330100032000203"></a>

## name property — interface_list / 022001130230 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312): complete subsection reference.

<a id="canonical-1231300102003220-3223200311023233-1013123111100101-3313333311013023-1032232202313123-0332301030311123-3000112123133100-2130311221023300"></a>

<a id="canonical-2211012233131111-2312212320313231-1100323103030312-2030211201011202-0332300112231201-2123323111022120-3312011111103233-1200201131120113"></a>

## priority property — interface_list / 022001130230 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-019.md#canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021): complete subsection reference.

<a id="canonical-1323010121013002-3133122231010012-0030230322100131-0330331103303303-2231123101012021-3123320212132013-0213002211133212-2323121303120012"></a>

## Next pages — interface_list / 022001130230 / 11

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- [vmware.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-018.md#canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-018.md#canonical-2123002311101023-1232110121333102-3302333100103032-2301311000103120-3001332221330132-0132113300211131-0012320300112121-1102023200322330)
- [vmware.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- [vmware.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112)
- [vmware.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202)
- [vmware.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- [vmware.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-019.md#canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311320131023312-0003212232001110-2202310032101221-0002311222303111-2002012000320011-2212222200130210-2230112210102103-1332002302000022"></a>

## vmware.not_managed.node_list.interface_list.bond_interface — bond_interface / 123220132321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3111200033322213-2211012132301320-0110011311302321-1120301223032321-0103332211010303-3000003232132123-2021033131031320-1200223003132320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0111321332332001-0010032110112023-2130033233312100-1021222231202021-0033003311101133-3123002013221322-3120100231000322-1333110002123131"></a>

## Direct properties — bond_interface / 123220132321 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-018.md#canonical-2332013311302110-0303120001131020-2132100220020213-2003120022230322-3023213123003311-3011000011203113-2121012131022133-3002312320202120): complete subsection reference.

<a id="canonical-2221102121231220-2030021311212302-2130312330031120-1112012333020020-2020300222001221-2030303101030320-1231322231123022-3033010312300312"></a>

<a id="canonical-0330313220031333-3211200213223023-3331322320030023-3302211123231113-2301202220213010-0233213002131112-0021033321120132-2210313001320013"></a>

## devices property — bond_interface / 123220132321 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3131232203322122-3201003121310020-2333300333212330-0302221022222303-1313320331102321-2031333311030020-1323130132111000-3231310230023301"></a>

## link_polling_interval property — bond_interface / 123220132321 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013202210322301-1232301113212133-0000132123211021-3232020212322102-3200130302232333-0102303333122111-0212112201130231-3000031022322001"></a>

## link_up_delay property — bond_interface / 123220132321 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0001221032022112-1003302303111011-3232031012321321-0101011302200121-0100332103231201-1232232021000010-0300031010212101-3121301330033100"></a>

## name property — bond_interface / 123220132321 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2001002310330211-3003332301000020-1231100001102010-1123100332001132-0001013212003113-1010111121102133-0302113001023223-2232001120231102"></a>

## Next pages — bond_interface / 123220132321 / 8

- [vmware.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-018.md#canonical-2332013311302110-0303120001131020-2132100220020213-2003120022230322-3023213123003311-3011000011203113-2121012131022133-3002312320202120)
- [vmware.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-018.md#canonical-1130133021200101-1213230111003311-2121203202030322-1112203302120332-0022323021000213-1300200202211230-3223300321332022-1323111303103210)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2332013311302110-0303120001131020-2132100220020213-2003120022230322-3023213123003311-3011000011203113-2121012131022133-3002312320202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131121001133000-0121102331200322-3033300232010300-1030010231303203-3010203010031313-3302133311232011-0012012310330223-0323322002032002"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 101330011232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- vmware.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-3023223301230030-3032111130021220-2203223220123333-1211223310022000-0123313232113113-0212023003310110-3320122133103301-3021130221002103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-2331003012031233-0030232010221031-2330013212131211-3031011323333121-3122311313112033-2232201320101302-0023132120033002-3321230201233100"></a>

## Direct properties — active_backup / 101330011232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131322013332300-1220001300113111-2300120313300022-2101021020201010-2131010331123120-0100102323002332-1200022132323221-3220032333323011"></a>

## Next pages — active_backup / 101330011232 / 4

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1130133021200101-1213230111003311-2121203202030322-1112203302120332-0022323021000213-1300200202211230-3223300321332022-1323111303103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001233103101311-0323231120332202-2303012101011313-1011332022330122-0323202001310001-0122000211301310-3221013230200202-1230300102232022"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 201311133211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- vmware.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1301200233121323-3323220001133031-3311232211112210-3310303013331202-1021330221021011-3210311133330223-3002213313230021-0022232220312013"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2011211020003130-3001011203323302-0233133130220101-0213120011021022-1112132311203003-2121103001232333-2303231000010231-2003322331000021"></a>

## Direct properties — lacp / 201311133211 / 3

<a id="canonical-1112001100113111-3033230332230101-2101003103310232-0321232323112022-0021311103021333-2233021231003331-0012032211332232-0230110123203333"></a>

<a id="canonical-1030103132332213-0132110331223221-0032103230200001-2102201023011220-0203003231332322-1030030120211003-0201131021301332-3110121132321303"></a>

## rate property — lacp / 201311133211 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0230120233123212-1311002221012211-3330133022023131-2031221233102232-3003023320311021-2210132100001102-1003311332331110-2122032033231012"></a>

## Next pages — lacp / 201311133211 / 5

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331102120203331-1000310313002133-1222112113200121-2110222102231022-2122203230233101-0300010300201131-3103022332021313-2310000310032010"></a>

## vmware.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 132212213112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2223123213031103-1032032333131022-1102213000013112-2000011030000200-0232011123231323-2113130100001112-2122111001132231-2011001320030002"></a>

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

<a id="canonical-1100120331300122-2303202033130323-1332113200113202-0003301232203001-2313323122230210-3200210002222300-3011212301203320-0211111223112032"></a>

## Direct properties — dhcp_client / 132212213112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131302320330102-0331103313132033-1331330111302332-2202111202131203-1202121011232033-1002331032231021-2101223013211233-0301030211110130"></a>

## Next pages — dhcp_client / 132212213112 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203333311101033-0101203323011131-3223313212330103-3012001103321113-1232212203222230-2001020310203213-1211001301030332-3132112200132133"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 033313002213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3010122310011333-3211000133301011-0120133123122002-2032001201023013-1133302223113122-3332100311131133-1211230002201013-2130123122132201"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321020321220211-2132333100100010-3200313110021020-3330202203110200-3120311132002230-2330023011121223-3301221122313202-3231323110013010"></a>

## Direct properties — dhcp_server / 033313002213 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-2121311232321012-1033303033320013-3101200321110122-2302122010321320-3021033100120132-2333130310020301-2312323232131312-3333022103112321): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-1032230102021201-2102320121102213-1001212032323203-0331131012113112-1110233121012123-2320011313300010-2031023122000220-3213122213320010): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213): complete subsection reference.

<a id="canonical-3221203211001030-1202123020312123-3032332033031031-2122300013211122-3000103133132023-0101213203112021-2220221321203131-2111323321301200"></a>

<a id="canonical-2320001000311310-2010123112113232-2023313033012102-2102013311321332-1102213300201230-3023111022213322-3103303033330001-3210200112111323"></a>

## dhcp_option82_tag property — dhcp_server / 033313002213 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3221221331102023-1300120302033321-1100300102011303-3000133201102102-3323310132122223-0122330232333213-1321330030223021-0133122333123032"></a>

<a id="canonical-1230332301030203-1222200302310123-0303212100332010-1322211310131102-2321130132232101-3333021112013130-1230121221310103-1133130330200322"></a>

## fixed_ip_map property — dhcp_server / 033313002213 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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

<a id="canonical-0120221112002132-1302031113003322-2020301011301202-0322100220101301-2331121020302101-0323013110032322-3301320330012021-1312013300030021"></a>

## Next pages — dhcp_server / 033313002213 / 6

- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-2121311232321012-1033303033320013-3101200321110122-2302122010321320-3021033100120132-2333130310020301-2312323232131312-3333022103112321)
- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-1032230102021201-2102320121102213-1001212032323203-0331131012113112-1110233121012123-2320011313300010-2031023122000220-3213122213320010)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- [vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-1210321213333210-3220323122002223-3010222323311033-1330113030010301-2221221230300312-2002110033223201-0113030323120331-2222300011331332)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2121311232321012-1033303033320013-3101200321110122-2302122010321320-3021033100120132-2333130310020301-2312323232131312-3333022103112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031100300012223-1012303022111103-3132121001332121-1301130033123302-0113203303003323-3102022232202131-3111210131110011-1110020323112312"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 111003302001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1030211011313312-0301102033031103-0302211120121303-1022223301332000-1231320003230123-2202123102312030-2123312221131220-3333100123321332"></a>

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

<a id="canonical-2211122300211021-0201233200001303-3033132032333323-3000331310000213-1032202012233202-2313222210221330-2003110110321013-3010212223021110"></a>

## Direct properties — automatic_from_end / 111003302001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120311230112221-2301122220330113-2000022122203003-0220123120301001-0103303302220323-3211110203101102-3202001212010112-0301132102202211"></a>

## Next pages — automatic_from_end / 111003302001 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1032230102021201-2102320121102213-1001212032323203-0331131012113112-1110233121012123-2320011313300010-2031023122000220-3213122213320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111333300202100-2111301223003131-2101131001030112-3210231230332303-1321303032100232-3203210323212130-0122022210132233-0003200233100230"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 120300023203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3320323221311102-1000132300002130-1033113313232102-2103010312332112-0220101300012302-2221231000001322-0233230300321322-2320120311121221"></a>

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

<a id="canonical-2102331122111110-2132202201313201-1312200031000333-2023123223102303-3311210000300011-0002221030012233-3233110213022212-1110200311333300"></a>

## Direct properties — automatic_from_start / 120300023203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123311132030030-1020333233010120-0333003023120133-3303000320033321-2300031011232133-1002221103132300-0230021212332002-2133031031310321"></a>

## Next pages — automatic_from_start / 120300023203 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011122133303232-2111231223333223-0112231312012021-0002330203223030-2011103131112022-2002233001021203-3111220310121032-1013121313031023"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 313021100131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-0112121133232130-3203120100132130-0031310130012200-0332312021123032-2003212130102323-3232213000110201-0003103312220230-2022313120103233"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3212311331103311-3332122100211312-1230211012203202-2303120200002322-1131001311101210-3102121102331210-3130130231013221-1203211002223113"></a>

## Direct properties — dhcp_networks / 313021100131 / 3

<a id="canonical-0022113102010222-0030220103322110-0012303110302133-2222210030213102-1000200100111232-1131210220311323-2322221213013001-0330032220131012"></a>

<a id="canonical-0223323231000233-1033030211232131-0100232020320333-2311002021301012-2300320333030133-0031331030220120-1122013021132101-2113131333211020"></a>

## dgw_address property — dhcp_networks / 313021100131 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303213200301323-2233022032200203-2010033120332332-3103102230033130-1103121103213203-3302233202323012-2013120231303313-1133220210110020"></a>

## dns_address property — dhcp_networks / 313021100131 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1120131110231023-1321211001231321-2122231323133312-1333213312332133-2031300101130103-0233113031323332-0120300332200012-2030102131332210"></a>

## network_prefix property — dhcp_networks / 313021100131 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1111011132220333-3123323000310113-2200012111021101-2112101131103032-3010023221110203-1132100000212031-2203103321103212-2233012010331100"></a>

## pool_settings property — dhcp_networks / 313021100131 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-018.md#canonical-3231123123331121-0230320211331003-2120101323231002-0100301112301133-2000232010212100-3233201013301123-1312102323103301-3033100213230202): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-018.md#canonical-0300311323132320-1112321031303221-0100011310132133-2333201333110232-2333300312233021-3133231120012032-2220110221223331-2001330300200121): complete subsection reference.

<a id="canonical-0320122200313032-1030021032313310-2323213201312010-2010200332103121-0100012320221222-3213230213122323-0111130111110221-1322112220022120"></a>

## Next pages — dhcp_networks / 313021100131 / 8

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1200322232101113-3020133013201013-0112311111232012-2111001220030230-3303330332002220-3221032013000000-3110102221230132-0320133220101123)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1302230323220113-2220221013331112-3233102001330113-2301131201333230-0220320011123212-2212332010010212-0222033011110312-1020101232330221)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-018.md#canonical-3231123123331121-0230320211331003-2120101323231002-0100301112301133-2000232010212100-3233201013301123-1312102323103301-3033100213230202)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-018.md#canonical-0300311323132320-1112321031303221-0100011310132133-2333201333110232-2333300312233021-3133231120012032-2220110221223331-2001330300200121)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1200322232101113-3020133013201013-0112311111232012-2111001220030230-3303330332002220-3221032013000000-3110102221230132-0320133220101123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103010221210230-3222221333321120-2021213013321233-3020031323233330-0113213332031333-1233031221033012-3033113101300212-3031230302001313"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 000122112000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-3031023213110303-3133330023200313-0003030312131011-0231122120301103-1300220320131103-1232010232032323-0310011311133132-2012123201010331"></a>

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

<a id="canonical-2221012321231223-1303012030011220-1032103113113303-0202031120023003-0111212303201322-0001003001130113-0113130303333211-1001302132120230"></a>

## Direct properties — first_address / 000122112000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221230213311211-3122312313002330-3100021233230101-2233111113310031-2202230221313232-2330101020223223-1222113323232012-3030113303331030"></a>

## Next pages — first_address / 000122112000 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1302230323220113-2220221013331112-3233102001330113-2301131201333230-0220320011123212-2212332010010212-0222033011110312-1020101232330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002301333332220-3103300112031013-1230023030033220-2031113321303312-1113121211132133-0231323223133102-0002220122023020-0320120203002001"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 132132223330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-3311333003222300-2211313331212102-0123221233130322-2330033213321230-3203332210221232-1022103012212223-3121332132103202-0322010132313003"></a>

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

<a id="canonical-1211033020333233-2213212032010311-0312111003210221-3120010321220120-0211301233012333-1122233013202133-3313322200232201-0113223223133222"></a>

## Direct properties — last_address / 132132223330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330311333301323-0013223033233132-0110100313223122-2120201203203102-1213120231202311-0302112302203111-2111012230330201-3223312122323111"></a>

## Next pages — last_address / 132132223330 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3231123123331121-0230320211331003-2120101323231002-0100301112301133-2000232010212100-3233201013301123-1312102323103301-3033100213230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222011120311002-2322210311012023-3113101103302233-2212002211223203-3132101220212233-3212123212001010-2030213032111133-3122123022311323"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 332330030033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3211323212101112-0011210220122321-3221320200310232-2231333222300212-1032011032123220-3131101222221001-2000233002012222-0012233201133123"></a>

## Direct properties — pools / 332330030033 / 3

<a id="canonical-3121131303220311-3230201212120110-0301003101212210-0002003232202022-1323123212210131-0200231002002020-3320120220333201-2032121122333020"></a>

<a id="canonical-0320232113323021-3110331203031100-2003213321000010-2330121110111320-1120112101113320-1331213330120310-0320212231111232-2013110200303330"></a>

## end_ip property — pools / 332330030033 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2022233022232303-0211012330113222-3130201103202133-1110320320303133-2000321112321303-1312103003203221-0231023000112212-0001033210231201"></a>

## exclude property — pools / 332330030033 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0331122123032131-3213021033211333-0032202032020000-3011011111033021-2300000300230023-0130023033313013-1020231010331332-1230002310113230"></a>

<a id="canonical-0002300212311233-2022201203223213-3103103023120001-2203101300010200-1112313333311332-3220221132323331-3022212210331120-1223010132103212"></a>

## start_ip property — pools / 332330030033 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1301032213030103-1231302001131302-0023032212303031-2320033022221010-0212311030303220-1302130232100232-2131221131113011-3302210231002023"></a>

## Next pages — pools / 332330030033 / 7

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0300311323132320-1112321031303221-0100011310132133-2333201333110232-2333300312233021-3133231120012032-2220110221223331-2001330300200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031000312131330-2121031122212312-0111322020311332-2212301220133033-3203313220013302-2021032311100020-2300002220132133-3202003232203131"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 301301311102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0333303210330303-3031312321232322-0031223221232012-3310111023103322-0002032223120213-3303311312030220-1310031210200010-0211000120131313"></a>

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

<a id="canonical-2000112203133210-0323131003132321-3302200333311220-3320122200233032-1011212103133212-0333000213033300-2103102332032201-2103020123030000"></a>

## Direct properties — same_as_dgw / 301301311102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332101123023223-0322210222032233-1003111010023010-3002032130311221-3030203032331223-3032312010023230-3012133330321122-3323220321300210"></a>

## Next pages — same_as_dgw / 301301311102 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-3333220303331313-3022100202310132-0313030113030302-3000113031030123-2212311110012221-0111000233231133-3321013223220003-0311221110330213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1210321213333210-3220323122002223-3010222323311033-1330113030010301-2221221230300312-2002110033223201-0113030323120331-2222300011331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112303132010323-0120212022001333-0013231133203320-3121023211202323-0222032311202012-3031101123101112-1001203011302121-2020330203031223"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 030211213223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0032231133230011-3323300112213100-1301032231123300-2330132213012112-2113023003223121-3330202222000133-2303110111020010-2122201002110020"></a>

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

<a id="canonical-2212311321131220-0003021011000231-0221003122221100-1020323232332013-3121012303133333-3103302031020203-2003022003000322-2211013100222212"></a>

## Direct properties — interface_ip_map / 030211213223 / 3

<a id="canonical-3111202331212212-2203313111301200-0232322111331323-3211021303113212-2002223320000130-1103321130220302-2202320013023011-2213301110302113"></a>

<a id="canonical-2322021012220300-2313311013113320-3110223301313223-0020330213031012-3300011321312211-1232300002001111-3223000123323300-1033201330210122"></a>

## interface_ip_map property — interface_ip_map / 030211213223 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

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

<a id="canonical-2133203021122232-0300213201011303-2321033321002323-1033322001322113-0033122121200011-1300130212032313-1113313323211003-3330001221211332"></a>

## Next pages — interface_ip_map / 030211213223 / 5

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330023302201013-1111121332221031-3331333312112111-2120032222303023-2010232122333230-1331212112011112-2033311311101200-0120202003132010"></a>

## vmware.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 332223310123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3320220232003000-0001122312131223-2323110233332213-0012211031113100-0211111132012023-3031310101000301-0031003121021333-3323202122012101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3322002030212333-2023013021310001-2211002122203113-3323211110110032-0121233131032202-1301210133013332-2012020332330023-0323011222123313"></a>

## Direct properties — ethernet_interface / 332223310123 / 3

<a id="canonical-0212031310221121-3332210103021101-1202321313320312-0131101230101131-0222202321323202-1111332321332222-1113033132113332-1213222200113220"></a>

<a id="canonical-1203333233300302-3133020312303030-1013100011223033-2121033233220132-1221302112313203-1222121002322120-2231013201013220-1221130232123211"></a>

## device property — ethernet_interface / 332223310123 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0112232220031101-0202101133320312-1012331122100203-0101102201210112-1333300331012322-1032303213031311-0231301203123100-2321200210100212"></a>

## mac property — ethernet_interface / 332223310123 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3220122132113323-0203312211020110-1031002100003110-1213313303200200-3111132121031012-2032312230121303-2232000233011221-1302101201022113"></a>

## Next pages — ethernet_interface / 332223310123 / 6

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232321110131231-3321303301303232-0030120130301123-3220013330133302-2310203310213331-2322133032002103-2002003303020200-1212311230232211"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 121101010332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0101213030122331-3311233210133033-1213101221113012-1001200320300132-2222031331120131-2321321023210213-2301122200130011-3303130120300022"></a>

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

<a id="canonical-1201131123232003-0102131133200111-2331112203122112-0203312233332231-2031210131130321-2101122001000120-0311132310002111-2113203120211103"></a>

## Direct properties — ipv6_auto_config / 121101010332 / 3

- [host](resources--securemesh_site_v2--reference--group-018.md#canonical-0023101013331032-0013303312303222-0330102322021012-3212212212000130-2332112031132012-3313102122102330-2123203103230310-3200300202210121): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020): complete subsection reference.

<a id="canonical-0221221121323100-0211202131213322-3323002001223033-3323301212010313-1230233213330330-2223320101100322-3011101303010213-2120003133221331"></a>

## Next pages — ipv6_auto_config / 121101010332 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-018.md#canonical-0023101013331032-0013303312303222-0330102322021012-3212212212000130-2332112031132012-3313102122102330-2123203103230310-3200300202210121)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0023101013331032-0013303312303222-0330102322021012-3212212212000130-2332112031132012-3313102122102330-2123203103230310-3200300202210121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221321320123233-1302202000333021-2232031110030021-3021112302300201-0211122121323201-3303303130202301-1320122233021022-2303012301013202"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 232220210021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1120002001312022-2322223112022321-1102120013110131-0331202132311132-2332113333032131-3312312103212032-2330321202032003-1130001111303303"></a>

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

<a id="canonical-1311032223323303-2320102130000232-3122221211013020-0010320122023002-0310123322213202-3101032130303031-3210112321113133-0102013322121130"></a>

## Direct properties — host / 232220210021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012330312232202-2301311311322211-1112212113300313-0311132321300320-0210121030232232-1033230212133301-0110111023121023-1030210112032001"></a>

## Next pages — host / 232220210021 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011223230232202-0311011030212311-1333231302232110-3201333111010320-1030213100221223-0131221202030112-0231101321301303-3320003012331110"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 230321011230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2003023332230030-0222010123330100-2201012302123110-2030020020112300-0010213320210212-1332210102123312-1202301013322220-3321201231033131"></a>

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

<a id="canonical-0033131301333013-2103233000103310-3220021201111021-2210331331233022-3233303323101001-2011002023233330-3332313232311232-2103230121300220"></a>

## Direct properties — router / 230321011230 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321): complete subsection reference.

<a id="canonical-1121132131203003-2220221032211223-3110131033102003-3010001310100231-0033300103033123-1201113202120011-0121200132032221-0313120120100112"></a>

<a id="canonical-3232313111310321-2001223032012022-2210230122213020-2303103031131122-0223213311120323-3023333121313003-2101131233110111-3213000103033212"></a>

## network_prefix property — router / 230321011230 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212032223322301-2302021312123321-1210323131201311-0220010012332311-3201132101321030-1003000212033111-3021100122200030-0302103333200023"></a>

## Next pages — router / 230321011230 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001121232320231-0010131203121313-0033231120301112-0231023201133103-3222200131001322-0002302131132313-0312232020303332-2311102200100112"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 302330013103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3222232220003031-1013012112100023-0220010311002110-3123330120300101-3021000212023310-2200223110103020-2023101212120023-1002201212133122"></a>

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

<a id="canonical-0113021002102003-1220000333011101-3223302311201113-1213031203331300-0132330302220333-2021203210320022-0211031113032120-0012022031300101"></a>

## Direct properties — dns_config / 302330013103 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-018.md#canonical-3333310230323113-0321102303321121-3020312331122002-0202033202222112-1013331320313013-1222011302110313-2011012130122113-3200033232111223): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310): complete subsection reference.

<a id="canonical-0333331312120120-1200033101222031-1003133222111230-1331213103322312-2102330303203111-0231313233112031-2312001001031222-1202202113111112"></a>

## Next pages — dns_config / 302330013103 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-018.md#canonical-3333310230323113-0321102303321121-3020312331122002-0202033202222112-1013331320313013-1222011302110313-2011012130122113-3200033232111223)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3333310230323113-0321102303321121-3020312331122002-0202033202222112-1013331320313013-1222011302110313-2011012130122113-3200033232111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323311332323321-1013322120022012-1112110211113333-1333000201131203-1312012312323331-2121221330121323-2131110100211332-2110231100201222"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 201020230233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3233211323102013-3200321330112031-2210132023221210-3011331120000000-0023311203330122-2102033023300220-3332302313023020-2113221110003211"></a>

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

<a id="canonical-3213120031213302-0311012210320213-0122103201033001-0320131100200330-2302102333232013-0230210001131323-3203211222230122-1003101130032222"></a>

## Direct properties — configured_list / 201020230233 / 3

<a id="canonical-1232233211222010-1233103231212212-2231001221112123-1323232120201300-0101331123113223-0120012323200021-1312200113010133-0031210310130033"></a>

<a id="canonical-3112002012331122-3031013311102230-2323223131120330-3031110312223012-0203101312032212-2020201220110121-0310112101102222-0212123311002122"></a>

## dns_list property — configured_list / 201020230233 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2130210212002332-3123221300133023-1022322210100303-1313110130022023-1021232001020132-0000201123230013-1331002212231222-1133323011120000"></a>

## Next pages — configured_list / 201020230233 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330121010033121-3112102000002311-2310302233121120-0230031133111133-1120212203312132-0031311003030213-0330221100112320-1122131031312121"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 201310110010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2332210323030101-1131320101312333-1112220121011313-3331002013110300-1120010032033310-3002310321220103-1202323311013303-3121100132210213"></a>

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

<a id="canonical-3301310201322111-2022230231212030-0023033130020010-1313022230230333-3101002022102210-2323201032203211-2000022011110013-1120220000213202"></a>

## Direct properties — local_dns / 201310110010 / 3

<a id="canonical-2311231113321120-0012133031212330-2201233013322023-1033101020022120-2302111032132122-3113213303021210-3223211030202332-3033223023322103"></a>

<a id="canonical-2011110111020322-1001210202032123-1301232333302213-2122130310233031-0111020213223212-1211001030121122-2331011101312011-1023210330232201"></a>

## configured_address property — local_dns / 201310110010 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3131121330122112-3311210123102112-1301321230323101-1310100201323133-1222201000231012-0013133132111103-2211302221110122-1313023322311313"></a>

## Next pages — local_dns / 201310110010 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-2331130003302210-1130032133232031-0212033203322032-2320030201313313-2203301200113221-1130011232322232-3322101231333130-0320320122133011)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1002302130301300-1111332313310230-0013031100103102-2033031220320330-0302030000303120-3122302302122032-0120203322300101-2130112100132131)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2331130003302210-1130032133232031-0212033203322032-2320030201313313-2203301200113221-1130011232322232-3322101231333130-0320320122133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123301101303002-0002012033130022-1031223000302123-0203210121231020-3222313232331033-3100202021211112-1210321311001032-2310100010123302"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 133112313222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1213111120021303-1111030113100002-0322003331031102-3133121112310002-3333230321332021-1003202311112000-3013101031123123-0103330013123133"></a>

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

<a id="canonical-3232230231113312-2121020223023223-1213102203132223-2100112220303211-0311102030101303-2220332033003031-3033131001212210-2131020312223311"></a>

## Direct properties — first_address / 133112313222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311212031101202-2112231012323321-2210131022213003-0130213330302221-0320001020020221-2033113310323012-1301223203332102-3013132032100331"></a>

## Next pages — first_address / 133112313222 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1002302130301300-1111332313310230-0013031100103102-2033031220320330-0302030000303120-3122302302122032-0120203322300101-2130112100132131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133320102011301-0212121101032121-3032311213311012-0010210031032332-0011201321201323-3001323220130320-3000233131023223-0131120332302120"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 313211020213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0323301130113122-0111030110010303-3333000310012033-2332301333232002-1133231203311323-1123220123130222-2030022033303223-0100232321130321)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1133002010130331-1212213020102221-3111321231201213-3121133010201033-1110312300203001-3310031300303201-0220202230300031-1332133330121100"></a>

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

<a id="canonical-0301303030001323-0300010000303200-0001233222311123-0033033020011333-2111201303011203-0331330233122221-2103201123330002-3213220313302302"></a>

## Direct properties — last_address / 313211020213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030121120133330-0330211233113232-3230020100011232-2212222033322133-0311031211010112-1101230020022211-0113133000103330-3010031012331310"></a>

## Next pages — last_address / 313211020213 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-3231213012131313-0101221013130330-3310221312111231-2231022123101201-0011303301203213-0021022322130321-2001210020303330-2022303230220310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111312112313303-0322211020200313-0120002100012002-3233222222303022-0113210002331302-2223310102032322-2100320023113223-0222031333220220"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 232203312211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2323220001222210-2320320013112032-1302032200300232-0210011213020230-3112322033001232-1123301312220223-1021231310022113-2213021221213011"></a>

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

<a id="canonical-3311123201313302-3021103233123021-1133130333102102-3322301121202310-2022031121203210-3110011000200110-0333300100232213-1203112200113102"></a>

## Direct properties — stateful / 232203312211 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-0030023302321332-3013200013110130-0103000130321213-3232013112030331-3000122213312310-0023033212200322-1330331212032023-0003312321111222): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-2323211033000320-3021122213122222-3323332032130302-3233201331133201-2032333230122333-0003220123213302-2222112011011012-0310330322303020): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102): complete subsection reference.

<a id="canonical-2101102011312223-1110211013003201-3320022232213331-2321322121110001-1323210121033331-1000120211332230-0311031132102123-2213210033300000"></a>

<a id="canonical-0323221221130031-3012300300031321-1310211123220212-3020101003332300-2322223221031020-1220102222210331-3301010310213113-0130031132232212"></a>

## fixed_ip_map property — stateful / 232203312211 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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

<a id="canonical-3033112100110313-3002110221212020-3310133121011123-1001120100102022-3002300310022313-1022311003003033-2011033233110113-2120313200112131"></a>

## Next pages — stateful / 232203312211 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-0030023302321332-3013200013110130-0103000130321213-3232013112030331-3000122213312310-0023033212200322-1330331212032023-0003312321111222)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-2323211033000320-3021122213122222-3323332032130302-3233201331133201-2032333230122333-0003220123213302-2222112011011012-0310330322303020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-2132002102122130-2322212311223333-2103331033221033-1033131120110320-1313322110200323-2303011212232133-0110120210210332-0031110122011102)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0030023302321332-3013200013110130-0103000130321213-3232013112030331-3000122213312310-0023033212200322-1330331212032023-0003312321111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130223123123220-0132202333303121-3312302011021213-0100233023120332-2033302001001013-1220113123131013-1201311020213133-0312102021122202"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 322221031000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1113133123012200-3320201323023031-3123221223200111-3102211020130022-0020131102030203-0203031330212202-2023002113230312-1010232213232302"></a>

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

<a id="canonical-0032120032210012-3202230031212032-3333230121202113-2302221210030111-3232111230312211-0002302322012012-2211110001222333-2203103212213030"></a>

## Direct properties — automatic_from_end / 322221031000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321131130103212-1023303033221330-2012301232231112-2232113202120211-3202033222031001-3023230120203020-3322201133001302-2303020001331213"></a>

## Next pages — automatic_from_end / 322221031000 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2323211033000320-3021122213122222-3323332032130302-3233201331133201-2032333230122333-0003220123213302-2222112011011012-0310330322303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323202200023313-0302322110220121-0331121133133002-1322101222123200-1312022322302002-3332123223111201-0022201121320223-1203010323223030"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 203211333013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-018.md#canonical-3011301323033332-2031230321130003-2030031331211122-2300011232111012-2020112103030003-2001211211013301-2102133202102123-0210003223320020)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0120300000101302-3202133002113032-1132011002112021-2312132223313100-2302110033210300-2203130331000003-3313033002211010-2301212332132302"></a>

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

<a id="canonical-3103210333013320-2213232330102201-0232012233203330-3211310333103231-3223113203103120-2213210011122213-2002202003101321-1323021130032203"></a>

## Direct properties — automatic_from_start / 203211333013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100000210132102-2030303101212013-1323331111121303-0231020211233133-0121101012101313-2013323130322220-0020130120330201-0233021033120102"></a>

## Next pages — automatic_from_start / 203211333013 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312022323201022-3311120023100313-0030031131001313-3300312011133110-2000211230201002-2132203321100002-3132213212130120-3230132232023300"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 000221312012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2000121033332321-2330303201202231-0022010310033130-1231102131013200-1023200011033030-3101320321030202-2131331122101233-3112003231033233"></a>

## Direct properties — dhcp_networks / 000221312012 / 3

<a id="canonical-3022212112002233-1233220132110130-1312132333123303-2222000120321012-3301320333212303-1300323201222103-0010110120110322-1010211221232220"></a>

<a id="canonical-0002011100233013-2131033232002102-2102221323330211-0111233222302102-1300131233303132-2201011322113013-3102010022232202-2221013332203301"></a>

## network_prefix property — dhcp_networks / 000221312012 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3211132003123111-1211011123111031-1231131332313202-2303020310201130-1112210032002313-1232221022322013-1032213000311332-2033322123221330"></a>

## pool_settings property — dhcp_networks / 000221312012 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-018.md#canonical-0303103101121101-3312332011211321-0110103313100020-3313101203001033-1020221110221100-0332211310113332-3032110310321033-3302021210300122): complete subsection reference.

<a id="canonical-2002033322011021-3023100321303002-3211202032113201-2220213122321102-1232011221130131-0220202333013133-1113130312032120-1310200001010112"></a>

## Next pages — dhcp_networks / 000221312012 / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-018.md#canonical-0303103101121101-3312332011211321-0110103313100020-3313101203001033-1020221110221100-0332211310113332-3032110310321033-3302021210300122)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0303103101121101-3312332011211321-0110103313100020-3313101203001033-1020221110221100-0332211310113332-3032110310321033-3302021210300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330120120130113-2301032103112201-0233033103033231-0200123111011121-3232213321210302-1010132101311032-2322123003200111-2012332010001310"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 312211113231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2233011003131021-2301303032300000-1202230103001232-1331220312022001-1033312302122233-0313231213223000-1033033203303103-2012310323032100"></a>

## Direct properties — pools / 312211113231 / 3

<a id="canonical-2121232203312221-1331321012303132-3100230301223223-2200032200303211-0023020211012003-3001001310211133-3131211203312030-3011311302231133"></a>

<a id="canonical-2202222332111212-1110003221211222-0323011133231013-3203202332230113-2222021223110231-0230021210333020-2023330320203301-3033213203122100"></a>

## end_ip property — pools / 312211113231 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2231022313121221-0112033013312013-3022030103022113-1310230220312010-3130200010122211-0301320113001300-0311130330113300-3100000113203321"></a>

## start_ip property — pools / 312211113231 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3300131022210231-1313021002210111-2001102023033203-2322132011302311-1321212102302021-3310012120012112-3203121210000202-1301322233330133"></a>

## Next pages — pools / 312211113231 / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-1332012330000233-0323120011313321-3220101222002112-0112211102311212-2301221003103201-3133210313333230-3030201210323212-3133202321032102)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2132002102122130-2322212311223333-2103331033221033-1033131120110320-1313322110200323-2303011212232133-0110120210210332-0031110122011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031200322122033-1122033301233000-0013232102322131-1002332322033112-0232203023002313-3303311013332023-3000111223113331-0011111320210022"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 231002131210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

<a id="canonical-0333121323323031-1331013221313002-1301311102021330-3132320312110321-3210211130022301-0311311323021303-1003310011221013-2330203003032233"></a>

## Direct properties — interface_ip_map / 231002131210 / 3

<a id="canonical-2113131002221113-1132231120233121-3223310010101012-1112113123111101-0121201121101132-1333210010113232-2011301222330312-2311330320303033"></a>

<a id="canonical-1011113233233101-0112331233113311-2332113313333212-2200012203111222-1320000000103212-1130233302121013-2220112322023330-2011001311312122"></a>

## interface_ip_map property — interface_ip_map / 231002131210 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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

<a id="canonical-1322221221021001-3002231100333003-3132120021212031-1303222301100123-3100320232332322-2100332231003001-0313331030013213-3311232202322011"></a>

## Next pages — interface_ip_map / 231002131210 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-1330133120032122-1310230131233032-2132113232231132-0200130230121113-2033210203310232-1103301130220130-2323101203333032-1321102312311232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2123002311101023-1232110121333102-3302333100103032-2301311000103120-3001332221330132-0132113300211131-0012320300112121-1102023200322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211001100130231-1201310201110000-0322102123002120-2212320213332123-1232222302001211-1101113122321230-1230331232031232-2212012123011002"></a>

## vmware.not_managed.node_list.interface_list.monitor — monitor / 110222311232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

<a id="canonical-1322013302033131-3202311021120121-1112223002103122-0000221311121313-2120111033212330-0003111103030330-2221102000333120-2103012203000022"></a>

## Direct properties — monitor / 110222311232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330030330123000-2110020322201121-1312321210220203-2332113232003103-2011101211021103-1220222003010012-0030023302102211-0002101103100301"></a>

## Next pages — monitor / 110222311232 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330202020000230-2132302030123012-3332113120321203-0302120003031122-0113112001220300-1130112322022111-3000020222001033-1000201100023122"></a>

## vmware.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 312133201330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1100012213300002-0021010020302331-0103220002022303-2203112110220232-3010232302211203-0122000310123232-1013032012230322-1103203011212332"></a>

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

<a id="canonical-0332122021012220-2322310321302021-3033320231301233-0323313102032203-3103202011003030-3210103022310312-3210321021122010-1321022232231100"></a>

## Direct properties — monitor_disabled / 312133201330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120031133132312-2310121322012233-0101203012120111-1200321121301022-2223021120032331-3233033222221311-3211133331213112-0130103330101110"></a>

## Next pages — monitor_disabled / 312133201330 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311223021320123-0103001333123132-0221212021332202-2301232012313222-2131102030330301-0113123320123230-1022312331123031-1233312133321302"></a>

## vmware.not_managed.node_list.interface_list.network_option — network_option / 222200022311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.network_option

<a id="canonical-0000110010300103-1001310000023333-3230222020231030-0123030022312233-0022330013232221-2330000323011101-0101113022213011-0002013122012123"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210310311113302-3310202302330203-3033321003002203-0011113233133320-1003033331333313-0133200123332303-3002103200230110-3031321230011212"></a>

## Direct properties — network_option / 222200022311 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-018.md#canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-018.md#canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132): complete subsection reference.

<a id="canonical-0310001300313133-0313121233033103-2101100232003202-1222233133231321-0231321323300323-3132132332100121-1033211231303011-2223323122313200"></a>

## Next pages — network_option / 222200022311 / 4

- [vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-018.md#canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120)
- [vmware.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-018.md#canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311122321331022-0003310002330110-3233203223211121-0311322201322023-1322121231110333-3201012200133321-0011203033011121-2320032323133312"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 330212102310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0021332210202012-3323031302101032-3322011323100302-0221032112021230-3012211230033021-0020022010122230-3230030033111202-0310332013033332"></a>

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

<a id="canonical-0213210033330302-0023211101213210-2200323130102132-0112321120111231-3013210223221200-3012221223233222-3101223202333202-3030010102111100"></a>

## Direct properties — site_local_inside_network / 330212102310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101322002113132-3133032102312010-1131311013111103-3303022000023023-2311303231200001-2110032200033112-1332031101002223-3033111113221103"></a>

## Next pages — site_local_inside_network / 330212102310 / 4

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132010312010011-0222312112100311-0301113113303311-2330030310031002-1202131333320231-1222230212111313-2003212111000031-1010222203112030"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 311232232331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- vmware.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0112333032000011-2211320102031221-2222030220310002-0311211231221010-3211220320202003-2032212132102110-0321201232323023-0220213000320313"></a>

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

<a id="canonical-1322122121231111-1102333120021210-3302112103013000-1223301010322130-3003122211010103-2221002101233302-3210101222131013-3023011022330110"></a>

## Direct properties — site_local_network / 311232232331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223111022322100-2113110203123223-3010020121333221-2131311311311033-0102000222021320-3313033322331022-0112230331032022-2031011121320132"></a>

## Next pages — site_local_network / 311232232331 / 4

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320220120031011-2333322012301133-1320330012300110-3311322222210032-0101331121030232-3331011200210233-1333310001220201-0222021101031020"></a>

## vmware.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 021133203101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3220332231331012-2221310003003010-1323111102311003-2131113312211320-2202202332223010-3021302001212230-2031222123212332-3300001003110101"></a>

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
no_ipv4_address = {}
```

<a id="canonical-1013101332212122-2130331212220311-2320233222303232-3012221313323023-0023320223102132-2103313333223231-3100012113021230-3213131032231011"></a>

## Direct properties — no_ipv4_address / 021133203101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300113033303230-3322033201011322-3331333332310101-0330212101303232-2203232223231333-2301222023110130-0113103122333202-0320323122311131"></a>

## Next pages — no_ipv4_address / 021133203101 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123332120312210-3330131021230332-3332322221321001-3313101310212003-0211312113323322-0301112133012021-2033313300010023-2230112320002000"></a>

## vmware.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 310112110200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0303221212303102-3032012032020231-2133232133100023-3003122000331202-1010132332032203-1001300003232003-0022113331213121-0123031002313111"></a>

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

<a id="canonical-1021120212201020-2213013211000022-3023110332302203-2201323023020301-1031233322332233-0222330130011033-1230333203203011-2213212201232222"></a>

## Direct properties — no_ipv6_address / 310112110200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112320232213320-1122320133322123-3020302310110232-2101013333101032-2011133311022112-1121001212002113-2000221212123202-0112001310012021"></a>

## Next pages — no_ipv6_address / 310112110200 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302111202333312-0030120133011203-0123311203100202-3101002220312223-3020021133122023-2212233030110123-3300132310013222-2120320101332201"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 220013223100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1322321123011233-3012311300311111-3102111312131031-3113313031033033-1032022230010022-2302011133010120-1222232131102311-0023321122302233"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-3030321023331223-0001213002101310-0330202020023133-3103310320321002-0032233231310020-1113110313210021-2022111202210020-2111231330203131"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 220013223100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330221023321021-2321222030323200-0102112222311222-1203102213121032-0303023332322213-2101001023230201-0032312220320300-1320033322231032"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 220013223100 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220101211221201-0100023030323111-3320311010300232-2001322203013011-0310112210230103-2303212012113223-1203102313333023-0313201001201202"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 100213100300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0201203123030221-3312023020123121-3102331031320013-2230302121033103-3121031101033101-2210003212130320-3212301020122013-0101113133033033"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-0002101010222031-0120322003332212-3223032101002031-3310332323030102-3213223123131330-1110202211130321-0103220301322332-1201112203002030"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 100213100300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030133230220000-3332313012130030-0110300210332330-0110333231032113-0330010120102310-0223302200000330-2201033011010131-2100310122313203"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 100213100300 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330130301101001-1303333012122231-1033312131300133-3212032313202301-0212131100323021-1013301203123121-1112002233122122-1101302011332211"></a>

## vmware.not_managed.node_list.interface_list.static_ip — static_ip / 301111131213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-0321323012033110-0130002002320113-0023221021200032-2022300001330200-2132301012303011-2120120330223323-0333002032103200-3320111300233213"></a>

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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001212221313313-1033332320100013-1033232111130030-2021121120130312-1103101311131023-3130130332211020-3300202300130021-0030101220103022"></a>

## Direct properties — static_ip / 301111131213 / 3

<a id="canonical-3032310032211133-0121302203132211-0121131332003333-2331310133020222-3003120202331010-1202330122203202-3103331132013103-2301122333122303"></a>

<a id="canonical-0000313012300211-0011301102320303-0030133001333101-1323113013031220-2101211011110123-2213222003201301-3111323303231121-1211323221120112"></a>

## default_gw property — static_ip / 301111131213 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0121013012310121-2300301231320211-0320221222301012-2123033311122032-1030000132302333-0001333311332210-1100310322011113-1000011300032223"></a>

<a id="canonical-2310120032102321-2111211132023210-3303021011322122-0310021222003322-1122120222100232-2302300002211302-2022202121311312-0103313033312111"></a>

## dns_server property — static_ip / 301111131213 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0313222311032030-1133332203322010-0333122321002000-3313231221323320-1322101102300223-0131110201101322-1032032022031302-0232020012300003"></a>

<a id="canonical-1032321012111030-0223131113303012-1300020011003102-0312313230320001-0232002120100122-1310221011322112-2103320122313010-0021223103320331"></a>

## ip_address property — static_ip / 301111131213 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232132201313213-2012320323221102-0333013203021132-0120102120020333-2311132220310320-0010310103332203-2032123231213310-2233130102233301"></a>

## Next pages — static_ip / 301111131213 / 7

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020222103322302-1030123202301033-2101211013212202-2332003221232130-0201010020003331-1003132331322001-1131110033330330-2123233302013332"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 211003202132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-1210012201220010-1122313132100032-3202000232113200-0210133321320020-0111132302331030-0100202232322131-0201213310210013-2302130010302031"></a>

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

<a id="canonical-1231333030001110-3012230210302223-2113101320122030-0223010201312132-1201130320310221-2003310301302113-1233020331231221-1232102133131331"></a>

## Direct properties — static_ipv6_address / 211003202132 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031): complete subsection reference.

<a id="canonical-2021132302323313-0320211221212310-2210312303223101-2033013110200220-2213331312312120-0332222231012030-0232210003112020-2321202233310203"></a>

## Next pages — static_ipv6_address / 211003202132 / 4

- [vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322220031021120-2113221311332011-1201032133331333-1111333302002123-3132311111211023-2101233022102323-1131333101031201-3332331012332012"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 210133301003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1121020030103121-0011012302201011-0111110121000002-1301032122120020-0011211213122232-3002130330100323-1203330320313330-2330321030310112"></a>

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

<a id="canonical-3110013012310031-0111003123202002-0211323023302301-0230200011032332-3231100012313302-0223000031021123-3131332213203301-1210202113212133"></a>

## Direct properties — cluster_static_ip / 210133301003 / 3

<a id="canonical-3133311110032300-3111020202211223-3120312131311010-3301112333301000-3310102011113310-0033311002112321-2101221321231211-0131020321033023"></a>

<a id="canonical-2031013211000321-1210131021233011-1023002012102112-2230213212220111-1230021100021012-2220310102031010-2233020220012101-2201122311013333"></a>

## interface_ip_map property — cluster_static_ip / 210133301003 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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

<a id="canonical-3232032313302110-2021332121313203-2203311322303000-3222210122231102-2132121001301110-2330313330133002-0200311021020101-3113011312201020"></a>

## Next pages — cluster_static_ip / 210133301003 / 5

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
