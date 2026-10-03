---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3133131112002020-2130102231103030-3001321121231032-2033311303012022-1220010002310213-3212333001322022-1030110133110001-0212012100333221"></a>

## openstack.not_managed.node_list.interface_list.static_ip — static_ip / 131202221311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-016.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-016.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.static_ip

<a id="canonical-1033022200100312-1131212220101123-1311023230320222-2111311033110121-0111103202000321-1130210300001223-3110111203333001-2332211200212120"></a>

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

<a id="canonical-0030002320320222-3300001111123212-1100311102212102-2303222120220233-0110302120011003-3303010122220313-2121301120133020-3010330200212122"></a>

## Direct properties — static_ip / 131202221311 / 3

<a id="canonical-1023320000320110-1100023312013200-1000323132202103-1211332302030031-0010233223120331-3313313332100333-2113222230110211-2120121220011111"></a>

<a id="canonical-2010023003231021-2211103112112102-0111200203303131-3003001002100203-3000000032012101-0302122232233230-2221213131211322-2013101131011003"></a>

## default_gw property — static_ip / 131202221311 / 4

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

<a id="canonical-2320202312331301-3012123122200230-2031201020030220-1020102132310320-0332130132322201-2301030030223102-0031332030132030-1201301131221131"></a>

<a id="canonical-2221310311111322-2100133110123231-3321302110031110-0021221233203203-0103103133201312-2310332330122122-3332112113333113-2231113322013321"></a>

## dns_server property — static_ip / 131202221311 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3213303211031123-1033212100220232-1112222321023010-3320233303212013-3021002323320020-1332032012021210-3223121013233210-0213013110101131"></a>

<a id="canonical-2010002202112020-0201212200201013-3201112133310012-3223231133102222-2110200330330210-1111011103313033-1300232211132203-0011323103100123"></a>

## ip_address property — static_ip / 131202221311 / 6

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

<a id="canonical-3210001000010111-0303133121201332-1303320312300100-2223212100213233-1032120203131200-0323223220011211-0130200000202210-0031303323123130"></a>

## Next pages — static_ip / 131202221311 / 7

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303011033102311-2121002323300102-3221130032330101-3100201021231301-3331032133333213-0232303330322310-2330222333000302-1120321311123203"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 112223111033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-016.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-016.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3332220323123022-1031200022030110-1231303222230020-0010210223021323-0032333132312202-2311223210223011-0303121031303013-0211122122203300"></a>

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

<a id="canonical-3211232311232332-1110032132210121-3203131300100010-1200110332022333-0312200302210103-3200322012222322-2202311002112210-0103000001000132"></a>

## Direct properties — static_ipv6_address / 112223111033 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-1322213312132232-1203013222023130-2031122231033203-3120003000321000-1213123021002112-1133013012331102-2010000101223331-2200312331302000): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231): complete subsection reference.

<a id="canonical-2201120332122303-3231201223221222-3020021130201023-2203321303320323-2222032232230123-3210211311032123-0322202121302223-1213322121112233"></a>

## Next pages — static_ipv6_address / 112223111033 / 4

- [openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-1322213312132232-1203013222023130-2031122231033203-3120003000321000-1213123021002112-1133013012331102-2010000101223331-2200312331302000)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1322213312132232-1203013222023130-2031122231033203-3120003000321000-1213123021002112-1133013012331102-2010000101223331-2200312331302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131332223120123-1331310230223021-0120223131022300-2202213220223201-0231211222203322-3102001103202322-1311302023213133-2303012332013102"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 033303302330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-016.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-016.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-017.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-3101312120002222-0110222110021310-0310330231222212-3130110231020101-3120200121113201-0333102202030213-2002111232322230-3102303021030131"></a>

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

<a id="canonical-2020311013121213-0110213201333122-2233121221023131-3111021112101100-3333102133001112-1331011310012011-1211000232210011-3001121320031301"></a>

## Direct properties — cluster_static_ip / 033303302330 / 3

<a id="canonical-3303301300031012-3011012330020213-1232231133001321-1220133023320312-2033033133013200-1311021000022211-1130303001110003-0332221211303313"></a>

<a id="canonical-2102131132121230-2012202213111022-0103103321033211-2132220323031231-1021331331300212-0323222133113332-3130203110223031-1302221202100201"></a>

## interface_ip_map property — cluster_static_ip / 033303302330 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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

<a id="canonical-0003012021001222-3311231022222302-1031223133333112-2221223002011133-1230110130001023-3213123110101330-2030122022301112-3101333130022103"></a>

## Next pages — cluster_static_ip / 033303302330 / 5

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-017.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031310221302023-0030230221001033-0111132200312100-0002000022010002-1022123320210320-0121002013200000-1020300220013020-2012311221313220"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 322121131110 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-016.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-016.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-017.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1030122030003320-3322100330010000-2103322223111332-1322232210221331-1103002133132333-2130120302000313-0233120220103232-2213131300130033"></a>

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

<a id="canonical-2112000031222300-0231203221021132-1313100201002130-1322221033303330-0200331110012212-2212231031133002-0223120223302000-3200331121010011"></a>

## Direct properties — node_static_ip / 322121131110 / 3

<a id="canonical-2313303231101001-3223111110310123-1033310223001301-0002021211101300-1132120110031000-2200232203300032-1123103002123320-3233322023132122"></a>

<a id="canonical-3033030211333203-3201200311132013-2012133113130021-0133210022212211-0120223023022002-0301101213112200-1102113310312231-1122301132100232"></a>

## default_gw property — node_static_ip / 322121131110 / 4

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

<a id="canonical-1012110020311330-0121213003011323-3210103220212231-2033323111222120-2333221000013233-2332322332012120-2132012312230233-0121030210233000"></a>

<a id="canonical-1120213032312020-3120032003213231-0220313101133131-3012200222200320-1110022223211212-1332302323033000-2213223221320222-0023013213231103"></a>

## dns_server property — node_static_ip / 322121131110 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0312231111001013-1001022211313010-2310102201011112-2320211203200200-1201023210031322-0023113321330330-3231103131013302-2103122031133303"></a>

<a id="canonical-0122331112122101-1311231100201322-1111200033233303-2302130003223000-2313110321330211-3323000131000223-3013132100000301-2203322312003303"></a>

## ip_address property — node_static_ip / 322121131110 / 6

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

<a id="canonical-3013022202102122-1203323211330332-3320230121030122-1013333133201331-3323001323223201-2232102232011020-1132213101310113-0222331201120231"></a>

## Next pages — node_static_ip / 322121131110 / 7

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-017.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0020330021133011-0122223222303323-2332322122302031-0103032300320231-1322120221332020-0002202212303101-1102001132023322-2011031201011130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302002311231313-3122330213201221-3003202233311033-3121223201111010-1121201201031222-1100112121330102-0031113123201320-1031210231001310"></a>

## openstack.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 212032321120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-016.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-016.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2100130333300131-3300313233000132-0221033331031333-1000032322313100-0131132022003223-0202302311013122-3102233123200312-2300313131022101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102220103122121-0330123123031120-2311201100211022-3121031332211111-0213003020322012-2200213012002301-2320133002333320-0302200220030211"></a>

## Direct properties — vlan_interface / 212032321120 / 3

<a id="canonical-3132202332331122-2000202131213103-2300310301232103-0221300332210330-1122232311230131-2131030220120321-0002130002013233-2212121102232312"></a>

<a id="canonical-3321002230130321-1310211000101220-1323221231033203-0102103032023113-2232323330310102-0011300332311032-0110200212302211-0131102131020213"></a>

## device property — vlan_interface / 212032321120 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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

<a id="canonical-0333011012111302-1023311230112221-3200132133332113-0311312112211303-0122311202212130-1023333303312000-3102321211323202-3101012010202000"></a>

<a id="canonical-3123233202332032-1003021023102302-0222102303122330-0311303212013213-3133223231200301-0213033233212012-2030111302220011-0302123303200132"></a>

## vlan_id property — vlan_interface / 212032321120 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3300023302223311-0013200230023320-2101300001033211-0321310302100232-0100231122302213-0102302102320111-0110002213011221-1323102330320102"></a>

## Next pages — vlan_interface / 212032321120 / 6

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202023023133233-0022113033303322-1010210101112131-1130120130002003-0312232302211312-2313230221310312-0212311112212223-0330332023321013"></a>

## performance_enhancement_mode — performance_enhancement_mode / 021103101012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- performance_enhancement_mode

<a id="canonical-3222313202323110-0220121122313311-2112010132221321-2303210302002101-3211220111320132-1310211311100123-2122213331100323-0021000321222211"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220123101231102-2103103310103232-1330022331213011-1313211312331010-2323033100121300-3011202333213131-3213232002320332-3330211033100200"></a>

## Direct properties — performance_enhancement_mode / 021103101012 / 3

- [perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320): complete subsection reference.

- [perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312): complete subsection reference.

<a id="canonical-2230213323101320-3101320313312022-3112231213032021-2313021130211121-0220113033300133-0020213100133300-1030200110332022-1320123233323333"></a>

## Next pages — performance_enhancement_mode / 021103101012 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031200131001301-1303132130032200-2333303132002012-3101011102000001-3030122011330011-0303011311210233-0133310010212031-1211300130032122"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 200012133033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2120233022030200-2010331323303213-1223333001220002-3033033032233211-3301210133311332-1212032100300102-3230011020220120-2003313013110201"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222302023201303-1020201131313212-0121013000011000-1221311321003123-2233022232013300-3031200003231133-2323033230031112-3212222102003023"></a>

## Direct properties — perf_mode_l3_enhanced / 200012133033 / 3

- [jumbo](resources--securemesh_site_v2--reference--group-017.md#canonical-1112021030022202-0232100133302322-3102000223021212-2321131233012100-3220003032033132-3202333331002300-3120120222132213-2103200001332202): complete subsection reference.

- [no_jumbo](resources--securemesh_site_v2--reference--group-017.md#canonical-2133331031330223-1233231002111321-0301013031120330-0031202100130012-0301322330210123-0132031223223212-2211212012302112-3111223202313311): complete subsection reference.

<a id="canonical-3322322230121332-0313000123300321-1201113121111331-2201020222000132-2310213001010030-2212112010201221-3232121023312022-1122331333032201"></a>

## Next pages — perf_mode_l3_enhanced / 200012133033 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site_v2--reference--group-017.md#canonical-1112021030022202-0232100133302322-3102000223021212-2321131233012100-3220003032033132-3202333331002300-3120120222132213-2103200001332202)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site_v2--reference--group-017.md#canonical-2133331031330223-1233231002111321-0301013031120330-0031202100130012-0301322330210123-0132031223223212-2211212012302112-3111223202313311)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1112021030022202-0232100133302322-3102000223021212-2321131233012100-3220003032033132-3202333331002300-3120120222132213-2103200001332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222223121122201-0211202231002031-0212131031033123-3022103331333211-2222133303021212-0302303011212230-1203033011233212-0300313211203021"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 123232203112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2210130332023122-0011311033212320-3121232113323113-2233230021130203-3110232022321233-0323323133102201-3002101110330111-1023103323012200"></a>

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
jumbo = {}
```

<a id="canonical-2301212200133313-1233032222201001-2111131301132203-2311022021331112-2202012022312011-2122232321013023-2012211310113021-2133321121121202"></a>

## Direct properties — jumbo / 123232203112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000023101033311-1322223123000023-2312203321013122-2122103013111122-2013100112323110-3033302313210210-1102212001133011-2221100013210010"></a>

## Next pages — jumbo / 123232203112 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2133331031330223-1233231002111321-0301013031120330-0031202100130012-0301322330210123-0132031223223212-2211212012302112-3111223202313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010101223033002-3223112300321032-3311310200000312-0200222300102331-2013220222032113-3300231300121213-0001321031221323-2210101021103212"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 121032032100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-1033102131303321-2000000100212310-2303201003313222-2212301021213111-0012133301000312-2112103203212333-0112003132300301-1103212232320212"></a>

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
no_jumbo = {}
```

<a id="canonical-1221212131330010-0320111033202101-3132220000300031-0231110313032222-0333122300120103-1303321332113333-3123003101213311-3010103201013023"></a>

## Direct properties — no_jumbo / 121032032100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231210211303130-1312312101213232-2021322330321031-0120131230130132-1002031313332130-3331131032321301-1211331310310133-3331330023131231"></a>

## Next pages — no_jumbo / 121032032100 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323013220003321-1220013311332132-3211302021332111-2020130212132121-3101321013021313-1303313023031320-2100211210313101-0130102333112210"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 313030122111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-1012202230032201-2331310300300132-3321111301013300-1111111131301002-3132321303200003-1032122310212220-2221233330000211-2210332333221031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312022200221101-3333033100300022-1223332312233321-2032233003333032-1210211130010302-3113012210211032-1111212111303022-2200113220000301"></a>

## Direct properties — perf_mode_l7_enhanced / 313030122111 / 3

- [jumbo_disabled](resources--securemesh_site_v2--reference--group-017.md#canonical-2032323233303102-0031210333001023-1212220111301003-3003133001012120-1302303322121131-3030311211121320-2313000212102310-1213102112001122): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site_v2--reference--group-017.md#canonical-0033222113030132-2031300202112222-0023320013221202-1000232133103101-1210312200013311-0131122203331020-3300022233010021-0201021103003133): complete subsection reference.

<a id="canonical-3303020023220211-2330313030223032-2303103213032213-1132101030302311-1331010030300310-2033110031013123-3302132131221131-0100211312011003"></a>

## Next pages — perf_mode_l7_enhanced / 313030122111 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site_v2--reference--group-017.md#canonical-2032323233303102-0031210333001023-1212220111301003-3003133001012120-1302303322121131-3030311211121320-2313000212102310-1213102112001122)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site_v2--reference--group-017.md#canonical-0033222113030132-2031300202112222-0023320013221202-1000232133103101-1210312200013311-0131122203331020-3300022233010021-0201021103003133)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2032323233303102-0031210333001023-1212220111301003-3003133001012120-1302303322121131-3030311211121320-2313000212102310-1213102112001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013101303121332-2221103201022203-0200221032121000-3120211130230123-1001120230013310-3013312021022001-0302312232001011-2131303101131130"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 110033133212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3111210213011001-0122233221230302-1223203231122302-2130011203231320-3112312120131110-0020133012132311-2120032223333231-2111000200200032"></a>

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
jumbo_disabled = {}
```

<a id="canonical-2230111323313210-0111312131322213-0233302101232021-0002133111220321-3010130031211021-2132300233210130-2222110300320310-0210201120322320"></a>

## Direct properties — jumbo_disabled / 110033133212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302330302033231-2311100013213322-3213330212133332-0113003133123021-1012201111211130-3202310300323000-2220230012110000-3222233120321001"></a>

## Next pages — jumbo_disabled / 110033133212 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0033222113030132-2031300202112222-0023320013221202-1000232133103101-1210312200013311-0131122203331020-3300022233010021-0201021103003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332130103110232-0022010202312310-0301303123310231-2000311122120222-0310013002323211-2233200001130220-0212223301210022-0213211011221333"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 003021212331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1002220121023331-3110130323330202-3111030031130111-1020100132011120-0333333003321031-2120111223223223-0302112132200331-2230011313013203"></a>

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
jumbo_enabled = {}
```

<a id="canonical-0012300201003322-0310032232322113-0000120101213230-1023010300301301-2220322200103130-2133203301213230-1202100212200020-0010310032221322"></a>

## Direct properties — jumbo_enabled / 003021212331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301301120011013-0313132023112223-2003311122011311-2001202030110112-3233031130332310-0323031130101010-2102313132331323-0313202133001223"></a>

## Next pages — jumbo_enabled / 003021212331 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-017.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132333020000033-3121313020113130-1113011323131231-1303130201313202-3331202201110330-3232030020101012-0011123211010231-1301321320123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123000231332122-2100303210002320-1130130223203132-0213000212002223-1123013013222030-0130110032021010-1220232212130210-2332230300231131"></a>

## private_adn — private_adn / 112323200122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- private_adn

<a id="canonical-3001100011332303-2231203223000211-1230200300233313-2012022223033331-1002001201232231-1312011231123311-2020311101200103-0010032010320220"></a>

Type: `"object"`. single nested block, Optional.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Upstream description:

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("private_adn")}
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
private_adn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222332132122132-0303001022332102-3021211331133101-3000231233110023-2213131121313212-0321301030311220-1003202012131311-1031111130031231"></a>

## Direct properties — private_adn / 112323200122 / 3

<a id="canonical-3332013313231212-1022231232201203-2121210130100131-0233132032220133-1333021033100333-2012332313132132-3332331001010123-3333301103230101"></a>

<a id="canonical-0013112111010021-3032012022111021-2102230232032010-0011131101311120-2200130321303103-1112200303020003-2032000310232210-1231102123033030"></a>

## private_adn property — private_adn / 112323200122 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3110202303321200-0030323112200133-2100303112013301-3221030000123113-2200312321331320-1102011031322123-2031033312310311-1221001202021001"></a>

## Next pages — private_adn / 112323200122 / 5

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111010032222333-1030323011220013-2020220312310001-2200110131200201-2321113233221201-3321010012310221-1231220211213102-1202220002322022"></a>

## re_select — re_select / 003021111210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- re_select

<a id="canonical-3101110123032132-0220020232313132-1303330213001120-1201313031301302-1121213131113310-1002022030113211-2022033012223103-3202010003233221"></a>

Type: `"object"`. single nested block, Optional.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("geo_proximity",
    "specific_geography"),
  validators.ConflictingObjectAttributes("geo_proximity",
    "specific_re"),
  validators.ConflictingObjectAttributes("specific_geography",
    "specific_re")}
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
  "x-ves-oneof-field-re_selection_choice": "[\"geo_proximity\", \"specific_geography\", \"specific_re\"]"
}
```

Terraform syntax:

```terraform
re_select {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213132331320010-2302311302122022-2030211200100213-2220030023233211-2212202031111002-2111110123102030-1122022211111122-1233111123230020"></a>

## Direct properties — re_select / 003021111210 / 3

- [geo_proximity](resources--securemesh_site_v2--reference--group-017.md#canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100): complete subsection reference.

<a id="canonical-1110021000312201-3310112332202210-2203130220120113-2020321232223022-2320113331312121-0310322130302231-0122333131013223-0312223032313330"></a>

<a id="canonical-2331103230011102-2202331331110031-0331122010233310-1310030102301032-1223202112110000-0101201331100311-0312012203033223-0003223320032010"></a>

## specific_geography property — re_select / 003021111210 / 4

Type: `"string"`. Optional.

Geographic selection for the site's Regional Edge connections.

- [specific_re](resources--securemesh_site_v2--reference--group-017.md#canonical-1211312003122330-1331202120220330-1302000112320001-0232023111130032-0221001111330011-0220021320230210-3033030101123302-3320121132333331): complete subsection reference.

<a id="canonical-3202102203310211-2013110333103322-2320220232212003-3101200312203202-0130031321030320-2131012010333021-3212322132300333-2131122310102220"></a>

## Next pages — re_select / 003021111210 / 5

- [re_select.geo_proximity](resources--securemesh_site_v2--reference--group-017.md#canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100)
- [re_select.specific_re](resources--securemesh_site_v2--reference--group-017.md#canonical-1211312003122330-1331202120220330-1302000112320001-0232023111130032-0221001111330011-0220021320230210-3033030101123302-3320121132333331)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013102030211001-0001130132012023-3023203110212003-1012233200010102-2020122313033203-0012111313010100-1202101333212113-0232103211112303"></a>

## re_select.geo_proximity — geo_proximity / 332303210221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [re_select](resources--securemesh_site_v2--reference--group-017.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- re_select.geo_proximity

<a id="canonical-3101000013030220-2131313232331012-3200020122110230-0013110120023003-3121301102020230-0130012322300332-1332321222231030-0110311332100102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for geo proximity.

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
geo_proximity = {}
```

<a id="canonical-2211323021303233-3120323231013322-2311200010131010-1232023202221330-0231032020211211-3033113221203233-3321203231200011-0112111232010033"></a>

## Direct properties — geo_proximity / 332303210221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103311000330330-3202100011033111-2031222122310202-1110131123223201-1210300033301201-0230232103120230-1332031302302321-2113313100310022"></a>

## Next pages — geo_proximity / 332303210221 / 4

- [re_select](resources--securemesh_site_v2--reference--group-017.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1211312003122330-1331202120220330-1302000112320001-0232023111130032-0221001111330011-0220021320230210-3033030101123302-3320121132333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023302120203120-3202200203110131-1210131100101213-2223312032230223-0220331322330103-1301100010330002-2230223213323131-3332301120122310"></a>

## re_select.specific_re — specific_re / 133211222200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [re_select](resources--securemesh_site_v2--reference--group-017.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- re_select.specific_re

<a id="canonical-3223112100302232-0223330113130200-3020201331033102-3301022011320213-2233222113210102-3221320313123030-3220102201232020-0032211213013103"></a>

Type: `"object"`. single nested block, Optional.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

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
specific_re {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132003010303102-1030011211332133-0332100321321001-2000031210302013-0123211110132223-0112023102130111-3022012121301001-0101120030031201"></a>

## Direct properties — specific_re / 133211222200 / 3

<a id="canonical-0331201030103001-1330202110331302-0222123113011303-1233223331112031-2001113321122301-1032321001221110-1110100022033113-1033123333031022"></a>

<a id="canonical-0333232032120000-0301021002320101-3330013002303121-1321213213011322-3322001301002120-3201323013223003-2120223200213100-3221211031021311"></a>

## backup_re property — specific_re / 133211222200 / 4

Type: `"string"`. Optional.

Select backup RE for this site, cannot be the same as Primary RE.

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

<a id="canonical-1303002323130123-3012331323333310-2213203302023132-2233000212003230-3001202012132122-0011311001200101-0210111121300021-3103302213102331"></a>

<a id="canonical-1332320310322122-0123033023232023-1312233201233302-1310301111113031-0132313332223100-3111031003232103-1011002231230013-2321222300020102"></a>

## primary_re property — specific_re / 133211222200 / 5

Type: `"string"`. Optional.

Primary RE Geography. Select primary RE for this site.

Upstream description:

Select primary RE for this site.

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

<a id="canonical-3131312231230110-2300023003130020-1002201230200113-3003101200100032-0230033322322122-2330220012032102-3023312202112122-0132033032210030"></a>

## Next pages — specific_re / 133211222200 / 6

- [re_select](resources--securemesh_site_v2--reference--group-017.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221020121132032-0111122133233320-2201110313233000-1322000000033023-0011301013100320-2302123230010320-3020020311102002-0200232031312012"></a>

## segment_vrf — segment_vrf / 020301101321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- segment_vrf

<a id="canonical-3230021233323013-2212321323223310-3232303201223222-0011210110010230-0233321201312301-3332031300130202-0022120223300213-0221022133112223"></a>

Type: `"object"`. list nested block, Optional.

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Upstream description:

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
segment_vrf {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212122322302111-2320301213130123-1231123033131033-0023123123103212-1321013313210132-0312200331013020-1230222033321300-2333122102303102"></a>

## Direct properties — segment_vrf / 020301101321 / 3

- [segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021): complete subsection reference.

- [segment_network](resources--securemesh_site_v2--reference--group-017.md#canonical-3110213102100123-2213032002131131-2001112022032210-1133200211232232-0332033311333113-3210132330223010-0302230001102032-0313121312022203): complete subsection reference.

<a id="canonical-2111100032320033-3011013222113113-1211332220133013-1102301203132311-3303223211112221-1011302310223100-0002311110110302-2131230012020220"></a>

## Next pages — segment_vrf / 020301101321 / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_network](resources--securemesh_site_v2--reference--group-017.md#canonical-3110213102100123-2213032002131131-2001112022032210-1133200211232232-0332033311333113-3210132330223010-0302230001102032-0313121312022203)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122010003022320-1212220210112003-1122311121101100-0310331023033202-3223221213312012-0210103223301020-2011211312131132-1022223032202033"></a>

## segment_vrf.segment_config — segment_config / 203131233130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- segment_vrf.segment_config

<a id="canonical-1121321000213231-2132012020102021-0011333001013023-0121310021021100-3311123331331231-1030310311010003-1030133002003332-3022011233330320"></a>

Type: `"object"`. single nested block, Optional.

Segment Network Configuration. Segment Network Configuration.

Upstream description:

Segment Network Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
segment_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213002331202222-2013030322033103-3301321123000001-1213231001212002-3230113032300020-3030111333100111-2211302312213303-1033322121303123"></a>

## Direct properties — segment_config / 203131233130 / 3

<a id="canonical-1000220003133312-2202300201020122-3032233113120302-1012300132131221-1233310103101230-2310323222200100-1202120020211331-0301012202231313"></a>

<a id="canonical-2100112210113001-0122302033213200-0110123222201020-0100023220320001-1312002032212103-1232300022101201-0130100101030000-1011300101032031"></a>

## nameserver property — segment_config / 203131233130 / 4

Type: `"string"`. Optional.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-1330032102221030-2100130132333313-2133103331222101-0010033002132332-3100131030120203-2012301200112221-3122103032223220-1000131121111122): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-1122111130130121-1310322230113120-3320103131022213-2330113120211023-2213120101110331-0110301202123320-3210123023220122-3121012032012303): complete subsection reference.

<a id="canonical-0020121300313303-2110013112123120-1010221222031113-1213220021110310-2321013213320103-2203213010003212-3032121133132210-2331201231200130"></a>

<a id="canonical-1111321330030113-2123200320223233-3200012131020211-2001010130032130-3021212000231033-1213030122031231-2031312311133320-3333320133221000"></a>

## secondary_nameserver property — segment_config / 203131233130 / 5

Type: `"string"`. Optional.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022): complete subsection reference.

<a id="canonical-0323133103303330-1233223303031103-2030112110200222-3223112013003203-0301212220301320-3321203101032333-0231212023002003-2003132220213312"></a>

## Next pages — segment_config / 203131233130 / 6

- [segment_vrf.segment_config.no_static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-1330032102221030-2100130132333313-2133103331222101-0010033002132332-3100131030120203-2012301200112221-3122103032223220-1000131121111122)
- [segment_vrf.segment_config.no_v6_static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-1122111130130121-1310322230113120-3320103131022213-2330113120211023-2213120101110331-0110301202123320-3210123023220122-3121012032012303)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330032102221030-2100130132333313-2133103331222101-0010033002132332-3100131030120203-2012301200112221-3122103032223220-1000131121111122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210131120001012-1001302233210323-2202221132012012-3030212132033330-0200321311131300-2220220110201310-3122321020230001-3202302131012212"></a>

## segment_vrf.segment_config.no_static_routes — no_static_routes / 223220331332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.no_static_routes

<a id="canonical-2312001031300222-1013111310221000-2011023223032021-0310023131113221-3322301220033301-2022210113332103-1000202023233132-1311312333121110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-3310211231220030-3222233123300020-1012203012021103-3323301002310222-2003012213321201-2212122020111003-0003130121123303-1012302101230112"></a>

## Direct properties — no_static_routes / 223220331332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113230212012221-1033320111331302-3233110132102002-3131201222113303-3333202210211302-0200321013132330-1201320132203212-0103303123200132"></a>

## Next pages — no_static_routes / 223220331332 / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1122111130130121-1310322230113120-3320103131022213-2330113120211023-2213120101110331-0110301202123320-3210123023220122-3121012032012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102011222202011-3100223320210300-1220213233132123-2301000103221031-0333302012221223-3002312301023110-1113312131300331-0333203130222303"></a>

## segment_vrf.segment_config.no_v6_static_routes — no_v6_static_routes / 111010333003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.no_v6_static_routes

<a id="canonical-3120210021311230-1112123322101302-1233203222233211-1120002123310212-3203300023033333-2222332110232020-2011100233200233-2321123202212013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

<a id="canonical-0220322123213302-2323203311201230-3300100111203313-1111312100222111-0222200312311233-3121320203222030-1231112303333021-3332221021220030"></a>

## Direct properties — no_v6_static_routes / 111010333003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232231231231121-3200123310002112-0012232101302311-1013022200110110-2331002131323233-0132021222331012-3132312210220300-3323122312322200"></a>

## Next pages — no_v6_static_routes / 111010333003 / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020320011231330-2301231132223210-2312210330300020-1201333322110222-2123223123122130-3313221231022012-3223111130313023-3111320130232112"></a>

## segment_vrf.segment_config.static_routes — static_routes / 133220121323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.static_routes

<a id="canonical-3102132021022121-2112330102023031-3303212003111203-1022313103120323-2202213030000330-3201223130230133-3130322120020322-1103202103000332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003020302311033-3032132231222102-2013310103111121-1320232300021112-2223302013223232-0123121121100332-1131313300312321-0320310020223213"></a>

## Direct properties — static_routes / 133220121323 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311): complete subsection reference.

<a id="canonical-2103222022011012-1300231130312210-0103201232013213-3222303322321013-1122332112202331-0112000132230330-3021132312033300-0321121133223212"></a>

## Next pages — static_routes / 133220121323 / 4

- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013213112201122-3020023013012221-3113323331001111-2001120122300100-1023220130312003-2232302010200121-3123301223120121-1303213120131013"></a>

## segment_vrf.segment_config.static_routes.static_routes — static_routes / 121332331212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- segment_vrf.segment_config.static_routes.static_routes

<a id="canonical-2300223320301032-1212311322133003-1301033111133331-0011012211032223-3103022231332030-0201331201010131-1031332110230300-0031211323331113"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
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

<a id="canonical-3010111313213230-0103023000022013-1203222023010120-1333210302331113-2113030120013332-1120223022213310-1020320112100122-2333100212322213"></a>

## Direct properties — static_routes / 121332331212 / 3

<a id="canonical-3013111120011010-2333333033002031-0022023221131002-3202311023120021-2331230321211100-3123300232303112-2321212122022300-1313130112321301"></a>

<a id="canonical-0120232023231111-0220132303013331-1030303313112030-1301020010120323-3013221111031332-1120210303321300-2131212002303311-0202120000022331"></a>

## attrs property — static_routes / 121332331212 / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-3212121032131230-3022010321123000-0312130021130002-2133123102020332-3311233020100301-3132031121022320-2102010111311230-0101000301331130): complete subsection reference.

<a id="canonical-2130301332322320-2003222322033122-0210301232223301-0121231231102202-0132320331123203-3221310132101011-3130230130322102-2233010031103030"></a>

<a id="canonical-3132112330111103-3203100310232301-2223210223232021-3233300023101312-0020103300033203-2313010310211323-0330302012330010-2020111011210302"></a>

## ip_address property — static_routes / 121332331212 / 5

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

<a id="canonical-2213131301232031-2133021313210212-0310001311313323-3201203013331232-1302320002012133-0003122110131301-3201212203320101-0321323222021030"></a>

<a id="canonical-0232001233301123-0212022212320202-2100233012230131-3103023332110000-3202133321332303-1020013310220200-2022301202113323-3310300211003101"></a>

## ip_prefixes property — static_routes / 121332331212 / 6

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

- [node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023): complete subsection reference.

<a id="canonical-0210332311232023-1230111211300321-3320100303122320-3103331210102200-0113033013322000-2133310113212212-3032220123101033-1222023331211011"></a>

## Next pages — static_routes / 121332331212 / 7

- [segment_vrf.segment_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-3212121032131230-3022010321123000-0312130021130002-2133123102020332-3311233020100301-3132031121022320-2102010111311230-0101000301331130)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3212121032131230-3022010321123000-0312130021130002-2133123102020332-3311233020100301-3132031121022320-2102010111311230-0101000301331130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203332100213000-1121210022011122-3033333310030300-3200211111131132-2203022022212232-3230131202203100-2120202021310210-1120021103223120"></a>

## segment_vrf.segment_config.static_routes.static_routes.default_gateway — default_gateway / 311012320032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="canonical-0213122212210021-2311112221232101-3012331113102302-2212032011331320-1012330121123133-1031120212211102-1322033311023323-0231233111030202"></a>

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

<a id="canonical-2113323010103231-2231231212320300-2000212301103013-3301100101113211-3202123302313013-3001322013311231-3210212023022323-2000010101332220"></a>

## Direct properties — default_gateway / 311012320032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301322100101003-0030213320220310-1232222213100203-1302311000033202-3132122232121011-1121123111132123-3202330123102101-0301103310302130"></a>

## Next pages — default_gateway / 311012320032 / 4

- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201022310112112-2321211033201000-1010301321023122-2213010132213133-2332301001222311-2020332123103310-0101000231001210-1333321212313201"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface — node_interface / 032121230120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="canonical-1320203221023300-0021031131002102-1021311003003020-3323323011211201-1021121210020133-1131111201132123-0032330123022032-1222301301331310"></a>

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

<a id="canonical-0323020101121323-1312020221003222-3121023110011202-2030013120331223-1032303312233113-0201100111331010-3103322220121121-1332030203322233"></a>

## Direct properties — node_interface / 032121230120 / 3

- [list](resources--securemesh_site_v2--reference--group-017.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102): complete subsection reference.

<a id="canonical-3320010333022311-1013231213101010-3302112213321003-0223223230322132-1013101101332100-2330301101233023-2123202021100223-3210102023131302"></a>

## Next pages — node_interface / 032121230120 / 4

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220101330102133-0112312201232321-2301131001221211-1120321201010312-1311033202031102-2022123012321330-0323331222323333-0231013320033120"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list — list / 133303100313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="canonical-2103303232131212-3003030331112110-3300020230202223-0111330033021210-3101131302200201-1023102032001201-3022232021200213-3212333031002210"></a>

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

<a id="canonical-3210121332110330-3211220232220012-1023030331010120-0330333132121111-0033120131202133-1310103323302330-2321330003300003-2103312310232333"></a>

## Direct properties — list / 133303100313 / 3

- [interface](resources--securemesh_site_v2--reference--group-017.md#canonical-1123232233033330-1210330312203023-0122110200133222-0312023012313302-3333123331132221-2031133013030111-3100211011203313-3132310132103032): complete subsection reference.

<a id="canonical-3101000120200330-3201002301032330-1303323012133011-0110322021333030-0003131031000122-0203313311030201-0022121232230331-2130023222310212"></a>

<a id="canonical-1123100323233032-3121012133021003-1132331221323203-2330020123032023-1332102122113231-3200221100230200-1303311311121232-0001120200223212"></a>

## node property — list / 133303100313 / 4

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

<a id="canonical-1013021003022103-2320232333310221-2030123000130121-1331313320033023-0123213103032133-1333212111003321-2023111102013323-3101013200302300"></a>

## Next pages — list / 133303100313 / 5

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-017.md#canonical-1123232233033330-1210330312203023-0122110200133222-0312023012313302-3333123331132221-2031133013030111-3100211011203313-3132310132103032)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1123232233033330-1210330312203023-0122110200133222-0312023012313302-3333123331132221-2031133013030111-3100211011203313-3132310132103032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121032302321021-0020001221223020-3000203023132300-1200010231022321-3310011210000311-0302300332311121-3201332131103212-0313121213011002"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface — interface / 333133030120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3022211220121201-2022013101133301-1233201333313120-0121211230012123-0120220333030001-2332232210030033-3010212323113030-3302220320233020"></a>

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

<a id="canonical-1131031231200333-3332013223211102-0110000023023010-1201013113112123-2001000332110302-2002003100020233-2231201112233310-2330231333332210"></a>

## Direct properties — interface / 333133030120 / 3

<a id="canonical-3030031033011122-0331231131221032-1220311000131130-2213201220020030-2131330003110212-3310232331201032-0122330332222322-1031010003123233"></a>

<a id="canonical-2030303323111110-2000003231133010-0320132003233232-1310332032113130-1010100023030110-3030302133130033-1010212100120122-2320113302023330"></a>

## kind property — interface / 333133030120 / 4

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

<a id="canonical-3211232013022131-1130010212233230-2111032110202120-3112320221221020-3021202111300301-1303202231113000-2320330332230232-3100000213302212"></a>

<a id="canonical-0103330333123203-1331300033332002-3110302223301311-3230321022101010-0113113201013301-1322310201020332-2321011011213023-2021002232032233"></a>

## name property — interface / 333133030120 / 5

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

<a id="canonical-1202233113223212-1211203112100200-1110202232130030-0323112112231330-3322200100100103-0020020102121120-0013331120313013-2001213211331023"></a>

<a id="canonical-1013303321302200-2000022101030022-2112010230133300-3311200002200123-0301133032301132-2122313201113220-2321002301331021-0012321021231032"></a>

## namespace property — interface / 333133030120 / 6

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

<a id="canonical-1300321312233111-0230022131003202-2302131001302212-2011021310111001-3010321333102323-1303103322333213-0322230322321033-1110002300033010"></a>

<a id="canonical-1211122123312011-2010221102101232-3132300320022312-0330300302112323-0111021302132032-1133322033313013-3013212223313320-3113200011233313"></a>

## tenant property — interface / 333133030120 / 7

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

<a id="canonical-1011231133313212-2112201102222001-3321011131123333-0222112011221013-3313221333010202-3213221303302331-2103210211000220-0332110101121101"></a>

<a id="canonical-2132201010110123-2212103320301231-2231020332301133-3012322211303000-1130320130313020-0330122333000031-3013233320213311-2301330033130101"></a>

## uid property — interface / 333133030120 / 8

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

<a id="canonical-0302321001311001-0321112212312332-1101033031223330-1321031223232123-0012010301003102-1001311301223012-1123320220022232-0223222103103301"></a>

## Next pages — interface / 333133030120 / 9

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011212201110311-2012320003022021-2113220033101312-2302011120213021-1232201102301301-3332100030223013-2330201003320003-0213032302313333"></a>

## segment_vrf.segment_config.static_v6_routes — static_v6_routes / 101303020303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.static_v6_routes

<a id="canonical-0132220013132210-3203210302102232-3223110310330213-2010032111133012-1003320321111023-3032032123131122-0012231032103221-2302013111211211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000211321331230-2030102023033121-1012310213010131-1120223102333321-3333210121230021-1213223020131122-3120200310201200-1232012032123010"></a>

## Direct properties — static_v6_routes / 101303020303 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312): complete subsection reference.

<a id="canonical-2011213001030010-3202200131023233-1200021330223331-0223332131023122-1003031203030133-2111330131100222-0003330132222203-0121120020201323"></a>

## Next pages — static_v6_routes / 101303020303 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120310233000021-1013321323222033-3103231112302200-2203121311303001-1110313032010302-2331200233122301-3021321200201320-0323000302210303"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes — static_routes / 032103212332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- segment_vrf.segment_config.static_v6_routes.static_routes

<a id="canonical-3311300003011333-3111112102311220-2123110203213122-3131002303000132-2131121313310213-1021132231012333-2222002310122321-0201121302120322"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
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

<a id="canonical-0112031100201211-2212332300001022-2111012120010321-0233013033013330-3031020300112201-0323022103010010-1321123100233200-1130222210022232"></a>

## Direct properties — static_routes / 032103212332 / 3

<a id="canonical-3033211210220223-1102002002301320-0033213121233201-0123321032002212-2031110212231110-1122113211000322-3030322003013011-1330330110001212"></a>

<a id="canonical-3013023122121220-3131000010212123-3030102133031021-3232132300131123-3232031221033023-0230230320211231-2020033110232300-3033200022011022"></a>

## attrs property — static_routes / 032103212332 / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-3121313130023221-0312013200302303-1222032003133200-3322033321030132-1110222032201022-1121030002103112-3002320101032212-3003013021312220): complete subsection reference.

<a id="canonical-3333211130212322-3213212300320030-3333232313103100-2230313230200012-0012010003011220-0131113021100021-1020230103302333-3302233212300003"></a>

<a id="canonical-3233100233211300-1012131101021131-1120223130113311-0111031201132322-3223012333122322-2121321232131310-2022223121031022-3012012333133222"></a>

## ip_address property — static_routes / 032103212332 / 5

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0121301231321002-2202010200100000-2320003321002031-1330033322123122-1032202112323013-2001233111132320-0113111330031022-1323301200220013"></a>

<a id="canonical-3330203013210030-1332131120021000-3221321233112222-0321212133231303-0103022302300212-1300012313112001-3133200222101331-3103011331222121"></a>

## ip_prefixes property — static_routes / 032103212332 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233): complete subsection reference.

<a id="canonical-1022021110000013-1022012001030133-2132032020320030-3130312113122301-3112322210103030-0033102211120312-1322210021310303-2332220230303331"></a>

## Next pages — static_routes / 032103212332 / 7

- [segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-3121313130023221-0312013200302303-1222032003133200-3322033321030132-1110222032201022-1121030002103112-3002320101032212-3003013021312220)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3121313130023221-0312013200302303-1222032003133200-3322033321030132-1110222032201022-1121030002103112-3002320101032212-3003013021312220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111323133320002-2301310221033201-3223002110010030-3223200303100202-1323320220303103-1122132121311200-1111313033031221-1221100112002223"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway — default_gateway / 302302301332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1333333030133003-2122003000031303-0022221333000000-2202223322321300-2303010012331111-0301133333330300-0232201123011310-2203010102320033"></a>

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

<a id="canonical-3133100113013210-1320030100231323-0000321013110130-2123331303301131-2121201022031301-3331323202012200-3010122332303102-2322202012301012"></a>

## Direct properties — default_gateway / 302302301332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331021003332022-3010330211002213-1211220313011323-3112131121323213-1100313330112222-2122003023322202-2323300112113030-3101311013323111"></a>

## Next pages — default_gateway / 302302301332 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222021231211321-0030333020100021-3000232112213233-1311332110222023-3321131220100122-3133212031002211-1210030201130100-1222122022003301"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface — node_interface / 322331020311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

<a id="canonical-3213113113333303-2312330123322211-0331033312023330-0312203333022300-2212103210002202-1012212233200123-3302021313231021-3302122321010201"></a>

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

<a id="canonical-1110002000323112-2121100202022211-2221023232032301-3200220022103322-1132132210303201-3100102100232331-0302131030021211-2301223023031023"></a>

## Direct properties — node_interface / 322331020311 / 3

- [list](resources--securemesh_site_v2--reference--group-017.md#canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031): complete subsection reference.

<a id="canonical-0203333311231321-0133011322332222-1311211131130223-1331201021022323-2200333311303332-0321321221330000-0323132331200212-1132303320110331"></a>

## Next pages — node_interface / 322331020311 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120103210223322-2103113313030320-3023220333323022-3321031033233303-2100013013130210-2323010002121033-2233102323323210-3013223201131212"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list — list / 113133331122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2023130321210222-1012031211220011-1322033310333131-1123200103202121-0113022331313002-1022333022321133-2012301213101010-3331202012002003"></a>

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

<a id="canonical-1233331113120313-3131013232320310-3300330313003110-2131202233311300-0021033220002030-2032122121002033-0322112323320100-0311200313213123"></a>

## Direct properties — list / 113133331122 / 3

- [interface](resources--securemesh_site_v2--reference--group-017.md#canonical-0203013210220203-3011210203020001-2303302000233320-1320323230302020-3330201113321322-2033213313131303-1011320032330120-2302322033112212): complete subsection reference.

<a id="canonical-3121100230233331-2311122013110013-3122103200332133-3012011121310232-0213320313231130-3120313131202200-1013200332000131-1102211221033101"></a>

<a id="canonical-0123231131213331-0031202023032121-1002203211233203-0213222102003032-1302133021130031-3133003210231323-2210113023100330-3111332212121323"></a>

## node property — list / 113133331122 / 4

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

<a id="canonical-0331001333120133-0301330001112011-1213110100112233-1113121133212003-1113221233221230-3010131322033003-1221130333332332-2212011022331100"></a>

## Next pages — list / 113133331122 / 5

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-017.md#canonical-0203013210220203-3011210203020001-2303302000233320-1320323230302020-3330201113321322-2033213313131303-1011320032330120-2302322033112212)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0203013210220203-3011210203020001-2303302000233320-1320323230302020-3330201113321322-2033213313131303-1011320032330120-2302322033112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321012300003202-0012330330102032-2112022120010203-3110032321011331-3201021130022323-2233201122110031-1002022333002211-3031102213330033"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 030132121033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-017.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3322023112000201-2313113111201310-0313001322311311-0120003322112233-3303130321322200-0000300211203011-0112213210021211-1310320233302032"></a>

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

<a id="canonical-1310022131121032-0203311121212020-0021020203131230-0311112132212032-0222330232210120-2110230323013112-1033023302210320-3210100120013021"></a>

## Direct properties — interface / 030132121033 / 3

<a id="canonical-2100233331311301-0110333311333212-2202323121012221-2213333311112022-2332201321002210-1231330033331010-3122020010233003-0011202113011010"></a>

<a id="canonical-0110202200133131-2030311322112213-3030101022113221-1310032122000203-0013230200203232-1010033133223310-1220130012330111-1233132111121203"></a>

## kind property — interface / 030132121033 / 4

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

<a id="canonical-1010101320213201-3331012100013031-3210303031020120-2322013231132032-3023023102301222-0233203222121100-1002230231112031-0202302213313120"></a>

<a id="canonical-1303001103002331-2120131231232011-2030033231333301-2331320322022021-3322123313103101-1330021202011031-3001333233311130-1321013123101213"></a>

## name property — interface / 030132121033 / 5

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

<a id="canonical-1032130221022210-2112002123323033-0210102321030320-0300213132022121-3203002112101303-1220112222121230-1331202212022330-1131230001322032"></a>

<a id="canonical-3202202222232232-3103033320031210-3330300233111012-2213133013223300-0311213210332310-2033310133102122-3100122322331103-1302100123220121"></a>

## namespace property — interface / 030132121033 / 6

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

<a id="canonical-1103321233213300-0122232333023303-1101310320111032-0300010210103003-3213021331010120-1113012231032100-2132221012223201-2310223010020321"></a>

<a id="canonical-0020112211231033-2222131033331313-2311023303121331-0030110333123131-1120330120011123-2312213000012330-1202033010331221-2000230130131133"></a>

## tenant property — interface / 030132121033 / 7

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

<a id="canonical-3330031110022033-0303011021331033-2122103033212232-2113000120323110-0310132130020233-2223203203211121-1220100013132123-3021000131323330"></a>

<a id="canonical-2310010323210203-2332112212310102-3220000000023210-1220323110003030-2111111333023133-3231111132212112-1130013113023313-2332301311212002"></a>

## uid property — interface / 030132121033 / 8

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

<a id="canonical-3320012102333311-1232330120123213-0221330010011023-1211231031201300-2100111130121130-0031123312211201-2213200211121102-0310330120003001"></a>

## Next pages — interface / 030132121033 / 9

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3110213102100123-2213032002131131-2001112022032210-1133200211232232-0332033311333113-3210132330223010-0302230001102032-0313121312022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203120113133121-1122321120021231-3230101132221210-2320332001023103-0121232330112232-3000300212200322-0211003113111001-3231221230103012"></a>

## segment_vrf.segment_network — segment_network / 030222231301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- segment_vrf.segment_network

<a id="canonical-3310010003302231-0220201010012300-3320220012120031-3000121322122220-2012220131121331-3103030203003100-2300222032222303-0110330120011031"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name for public API and Uid for private API This type of
reference is called direct because the relation is explicit and concrete (as opposed to selector..

Upstream description:

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

Terraform syntax:

```terraform
segment_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232100032123003-2200213333122112-1311111120200330-2133020100002232-3101100302121310-1102121022202302-3013221113211213-3202122021131211"></a>

## Direct properties — segment_network / 030222231301 / 3

<a id="canonical-2033202132013303-1013202332031231-1321222322111002-3300033213122211-2322230023233210-3321020032003313-2001313101302020-2103101312021221"></a>

<a id="canonical-0011332003033231-3332123221331110-2322301123101211-1323311310023021-2331330200223130-2320120303112233-3220111011220210-1303020222002011"></a>

## kind property — segment_network / 030222231301 / 4

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

<a id="canonical-0321313113030030-3010203112321201-3231223030331300-2223030312122210-1213113111102003-1300003311302200-2003201303002333-1202231302310230"></a>

<a id="canonical-0200213011301001-0133202100103222-2330310122303010-2030021202312002-2000032100022320-1302131311031002-3323112212100300-3233213000230322"></a>

## name property — segment_network / 030222231301 / 5

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

<a id="canonical-3101310313100030-1203200302112323-3211001202322101-3032113120333230-3120222211133021-1132023202112201-3001212200130323-1232031003221230"></a>

<a id="canonical-0321201012330321-1013313020100332-0320121003021102-2233321031133211-1000003223120231-3010211101032121-1033032013103333-2103102023012302"></a>

## namespace property — segment_network / 030222231301 / 6

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

<a id="canonical-2210221032033312-1101011031200233-3303110200001122-2021231123111023-3322020012101220-2332110031213112-1001011122320003-3210323011123332"></a>

<a id="canonical-2203033202010312-3022020213113033-2023303130111023-3320023210110011-3302320203133220-3102101132302002-3330230322103000-3213122011313230"></a>

## tenant property — segment_network / 030222231301 / 7

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

<a id="canonical-2120020210210311-2232103303331021-1233100313021232-0233013122230231-2031010113111232-3121232103100331-3310122010022133-0131011020310232"></a>

<a id="canonical-3232010312021200-2213133022012020-1301021100130101-1312220120200123-0030213031223333-0201302113033210-3230101332301201-0032333013232110"></a>

## uid property — segment_network / 030222231301 / 8

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

<a id="canonical-1321232023300331-0131320233011331-2302300113302021-1030103001321331-1311201002201202-3002203133121302-2111003111311311-3332000121130332"></a>

## Next pages — segment_network / 030222231301 / 9

- [segment_vrf](resources--securemesh_site_v2--reference--group-017.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122203213000110-1030200021232110-2212000032322002-0020221023122303-2023022130303313-0000133020331310-0321310120002102-3301030203022101"></a>

## site_mesh_group_on_slo — site_mesh_group_on_slo / 121233202331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- site_mesh_group_on_slo

<a id="canonical-3212222333120111-2230203121211023-0313133130013232-2310301323102000-1222112001023100-3302312232123100-3211022311121112-3323310000112231"></a>

Type: `"object"`. single nested block, Optional.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_site_mesh_group",
    "site_mesh_group"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
site_mesh_group_on_slo {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131302013012332-1200322313011323-2322303132122030-2223021103300023-3320001203122202-2022303311210211-3111000311203023-3111323031003031"></a>

## Direct properties — site_mesh_group_on_slo / 121233202331 / 3

- [no_site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-3213012221100220-3100021321210121-2201100000112121-1010010332201323-2131222301213210-0230011302111310-2203212022110131-3230021003021110): complete subsection reference.

- [site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-0221223233130001-2133301130312131-3202122030310102-1130300022002301-1220302213132021-3013131201321033-0313221122110012-2023311312012002): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-0132123210131012-0313113211220131-2333303223213133-3101032222221100-3003000133203332-1102020013223020-3000112020112201-0320222130320130): complete subsection reference.

<a id="canonical-1013220213202130-3123212001320101-1312220111332031-2202332033213133-0131333311131323-3331231332121031-0122002203330312-3201220121222112"></a>

## Next pages — site_mesh_group_on_slo / 121233202331 / 4

- [site_mesh_group_on_slo.no_site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-3213012221100220-3100021321210121-2201100000112121-1010010332201323-2131222301213210-0230011302111310-2203212022110131-3230021003021110)
- [site_mesh_group_on_slo.site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-0221223233130001-2133301130312131-3202122030310102-1130300022002301-1220302213132021-3013131201321033-0313221122110012-2023311312012002)
- [site_mesh_group_on_slo.sm_connection_public_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-0132123210131012-0313113211220131-2333303223213133-3101032222221100-3003000133203332-1102020013223020-3000112020112201-0320222130320130)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3213012221100220-3100021321210121-2201100000112121-1010010332201323-2131222301213210-0230011302111310-2203212022110131-3230021003021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300011020210201-1023203232111323-0311121303310131-1120323013023320-0001202012121303-3313031113032022-0100032322131023-1123221300111113"></a>

## site_mesh_group_on_slo.no_site_mesh_group — no_site_mesh_group / 123031320302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-3012320000213211-2212320133121022-2221133301312211-3332132310100323-0111300311010302-0303111301210210-0133002112233011-2310311013222031"></a>

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
no_site_mesh_group = {}
```

<a id="canonical-3110322221221223-1313010300313332-2013030233202211-1212130200030331-0330311112132101-1103031221312113-0212202223030203-3232122300111003"></a>

## Direct properties — no_site_mesh_group / 123031320302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300302323211021-1130120223123213-3001132201021220-0101122223031303-0311013033221123-0201233011030321-2223001302201300-3213131333011301"></a>

## Next pages — no_site_mesh_group / 123031320302 / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0221223233130001-2133301130312131-3202122030310102-1130300022002301-1220302213132021-3013131201321033-0313221122110012-2023311312012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121130221133112-0303233222200030-0031012010022211-2320121302301333-1120230230230032-2313122120203203-3213211220101222-1323131010311313"></a>

## site_mesh_group_on_slo.site_mesh_group — site_mesh_group / 213302333302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-3203031101211003-3103032221321300-2122232301113121-3133013022323300-1213213302220000-1203031332000001-1112310133312033-2231210233022121"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
site_mesh_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121003010030223-3202232011201232-0222203001130223-3031033000103120-2100002031212200-1021103331201130-1020313220231101-2013130233301321"></a>

## Direct properties — site_mesh_group / 213302333302 / 3

<a id="canonical-2223102120030000-0131103201300122-1033300000031313-0202212301101212-1101123130320021-3022012032221011-3132121000331222-1033220010201002"></a>

<a id="canonical-1211200111011301-1321030111233222-1220122002032103-2131220010023122-0220221230001330-3303223230221230-0223321220203113-3210003100332202"></a>

## name property — site_mesh_group / 213302333302 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2121221023121303-0122111123030112-2122020210100030-1011013100101101-1000213213310013-0330122330213320-0123330002202211-2222221211203320"></a>

<a id="canonical-2311322030322102-2032031130021133-1303211100102303-3000323311001321-3221310021011010-1011021112231112-3300101032300303-3223231203123320"></a>

## namespace property — site_mesh_group / 213302333302 / 5

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
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0211103112030103-1111230213022322-0301321301101312-2111110122120121-3032033211220021-3322031320200321-0333111131303310-1201330002211002"></a>

<a id="canonical-2002322023333332-0323002031132113-2312120030300031-0303010300311011-1023211223120231-1122231310210230-0131020223021211-2001102012012111"></a>

## tenant property — site_mesh_group / 213302333302 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3103103121213213-3012020323130311-2013302210023020-3303033221320022-3130103223132233-3113130232212020-2003131003131003-3202232303031233"></a>

## Next pages — site_mesh_group / 213302333302 / 7

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003102313122130-2320022110003001-3313000212203130-3033032120021200-0002121312331020-3223102223233123-1103123003231323-2210122132022013"></a>

## site_mesh_group_on_slo.sm_connection_public_ip — sm_connection_public_ip / 000131112122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-1320133300201220-1123311331301032-0102212203331001-2301323323000101-3211130111200033-0131303000303111-2223310312130131-1102102111032333"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-2331102213232022-3202031231221302-0121012120131321-3223102330323022-3221323020121123-1131003223111012-3103031210112102-1020200313030233"></a>

## Direct properties — sm_connection_public_ip / 000131112122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033002101033303-1130233103001311-0231330323011003-1102112011330231-0222003032313321-1333010100031232-3002103002221330-2230010213300001"></a>

## Next pages — sm_connection_public_ip / 000131112122 / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132123210131012-0313113211220131-2333303223213133-3101032222221100-3003000133203332-1102020013223020-3000112020112201-0320222130320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112102223232012-0021230002300103-3322303002230123-2101122322230011-1230011311032121-0031201330000231-2302103002323223-1023210113023333"></a>

## site_mesh_group_on_slo.sm_connection_pvt_ip — sm_connection_pvt_ip / 221320103031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-3202013112113133-2210320113101033-0220010013032102-2330210232322323-1211000232032301-0031130021122002-3223313210323102-2021222332203323"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-2323200132200310-1331323333303131-2212211330212013-3112132202122320-0023020332031133-2331210321013320-0201021022101030-2323133222010030"></a>

## Direct properties — sm_connection_pvt_ip / 221320103031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203132100030023-0232123022310232-1203203012033222-0101110101222031-1203200030223101-0313012130020002-0001130210201313-0001000103120110"></a>

## Next pages — sm_connection_pvt_ip / 221320103031 / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112321100201220-3003031031223103-2112331001020123-1330221011301012-0110212011330033-1131030213121312-1320030121313200-2130331322301031"></a>

## software_settings — software_settings / 101203232033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- software_settings

<a id="canonical-2021303230011013-3122113030331033-1131000201101130-1223003121320230-0201302021311101-2331313102113102-1032022110212020-0233323113321131"></a>

Type: `"object"`. single nested block, Optional.

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created. This block is a create-only,
write-only input; changing it replaces the resource, and refresh preserves the configured value
without claiming XC observed it.

Upstream description:

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created.

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
software_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130111121000103-2001221023012113-0313312130022322-1233212023020321-1032320120000013-1233132003221332-1200132012320210-1201013213232011"></a>

## Direct properties — software_settings / 101203232033 / 3

- [os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302): complete subsection reference.

- [sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300): complete subsection reference.

- [waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031): complete subsection reference.

<a id="canonical-2112022112201213-0231021113031333-2322002301120022-1203201110022332-2331112021231333-1011231200330121-0212001221221123-0211321111122213"></a>

## Next pages — software_settings / 101203232033 / 4

- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302)
- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233021331102033-2301001132012203-1011112101302013-2231002323013311-0101323201100202-1311020301021112-2000031321022233-0322023220222203"></a>

## software_settings.os — os / 300122102223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.os

<a id="canonical-1310220101111131-2102110111201321-3111032230233210-2332211103222102-2023133321023231-1230021221222323-0332120102122103-2113302330301320"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101031110311132-2200210102311223-2200311103021033-0023130201101310-3320033003223032-0020100101001312-0202212032000203-0111021212103303"></a>

## Direct properties — os / 300122102223 / 3

- [default_os_version](resources--securemesh_site_v2--reference--group-017.md#canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232): complete subsection reference.

<a id="canonical-3110011123301322-3023331103222310-2201322123101120-1012002231201022-1221220333012123-3010012213133323-0123020021100230-1001213202111013"></a>

<a id="canonical-3001132313330130-2103320301210230-3002311131313233-0300131213121330-0203012333013000-0120200303130322-3001322232323000-2002132021230202"></a>

## operating_system_version property — os / 300122102223 / 4

Type: `"string"`. Optional, Sensitive.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0012222111011230-3310100212110022-1332002132113123-0003330112200131-2313020313121301-1303133033231031-3001301200001311-2030001301231100"></a>

## Next pages — os / 300122102223 / 5

- [software_settings.os.default_os_version](resources--securemesh_site_v2--reference--group-017.md#canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132313132130222-2022010120302023-0103023321032321-2221101022032111-2120333031110213-3003201133221221-0103033033030200-3123133231333301"></a>

## software_settings.os.default_os_version — default_os_version / 120102303030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302)
- software_settings.os.default_os_version

<a id="canonical-3111110002212210-0102322331210330-1001201020332230-2031031302203002-1103221120211323-0313103112310301-2333301331323322-0133130222221312"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
default_os_version = {}
```

<a id="canonical-0202200333131102-0323331121010212-3302002031132311-3302203312331131-3323332320032300-0330023301212102-0333033133302020-0120213023201322"></a>

## Direct properties — default_os_version / 120102303030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300130223211232-3112203011131022-2122031213001000-1030131103012102-1122313113112003-1103120123201233-1000320113122311-0021200220003303"></a>

## Next pages — default_os_version / 120102303030 / 4

- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130123111221101-1330322023212120-1123113330221033-0000120011000333-2110013030001031-2300203001200231-3111221131233102-2321320332113132"></a>

## software_settings.sw — sw / 300021000330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.sw

<a id="canonical-1110113213023313-1102311132110301-3323021130312011-2100333211122331-2120033223230133-1031233211311033-1120232111230220-2113330201012231"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223000331030221-1211001213031230-3120122101002202-2230033232010133-0031030010323212-2201212023230223-1301123230330102-0022123122302122"></a>

## Direct properties — sw / 300021000330 / 3

- [default_sw_version](resources--securemesh_site_v2--reference--group-017.md#canonical-2321301203002322-0331112232333333-3333211033120312-2001312211012223-1020013013113030-0021212302302320-0122320312112021-0031302110201003): complete subsection reference.

<a id="canonical-2210323000122022-3231320233210302-1023302211020032-2131233123030233-1220000202132122-2031103030233111-0033223301101020-1103222333232202"></a>

<a id="canonical-1333131212220210-2302002011121002-0000230220011111-1303012100103223-0000313231311021-1001322301110012-0333110221112020-1013022132302332"></a>

## volterra_software_version property — sw / 300021000330 / 4

Type: `"string"`. Optional, Sensitive.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3233322211210330-3031110321131321-3203210132020022-2133332211333003-3021012121011230-2230123303301022-3013203320011013-1102102213313021"></a>

## Next pages — sw / 300021000330 / 5

- [software_settings.sw.default_sw_version](resources--securemesh_site_v2--reference--group-017.md#canonical-2321301203002322-0331112232333333-3333211033120312-2001312211012223-1020013013113030-0021212302302320-0122320312112021-0031302110201003)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2321301203002322-0331112232333333-3333211033120312-2001312211012223-1020013013113030-0021212302302320-0122320312112021-0031302110201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003332031101021-1301110332310212-1112222030301333-0120031120233301-3303022201023230-2123110002323120-0233011210033033-0232023203132212"></a>

## software_settings.sw.default_sw_version — default_sw_version / 323200332112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300)
- software_settings.sw.default_sw_version

<a id="canonical-3203220002133302-1220333322120121-2112012331302320-2200100123110230-0121033011123310-2120313011200210-2110132302330203-3032220203131333"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
default_sw_version = {}
```

<a id="canonical-2311030031213101-0203213000320130-3211320200110313-3113131200301033-2330022103020223-3300111030023113-0302210111011031-1131011031230303"></a>

## Direct properties — default_sw_version / 323200332112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320131032320231-2123102003122033-3101133120331011-1320230122301333-0233111222210220-3000110223211300-0011001112213332-3301102310113031"></a>

## Next pages — default_sw_version / 323200332112 / 4

- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210233033330120-1322212123000110-2201330120211100-3222102213011203-3210002020332110-2222332222120202-1330110301103231-1323033332012123"></a>

## software_settings.waf_signatures — waf_signatures / 212122210203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.waf_signatures

<a id="canonical-2123333211201220-1021203012100232-0221321230021311-3110233330322110-3200332000212330-1013002101200320-2333132212230221-2013331233320202"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230231213001033-1032203333230311-3311313200111023-1333332201033100-3100201101232122-2212312103321320-2023103132023211-2213322311320202"></a>

## Direct properties — waf_signatures / 212122210203 / 3

- [automatic](resources--securemesh_site_v2--reference--group-017.md#canonical-1321112103302132-0311330003120012-3202222023122102-2222222333321323-2313231200132032-0001101331002033-1112022212311112-3010310111203233): complete subsection reference.

- [manual](resources--securemesh_site_v2--reference--group-017.md#canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023): complete subsection reference.

<a id="canonical-0131223322233130-1322031233212022-3200121132331320-3233300230312000-1131321332100301-0131011132321023-1120231202230222-1121310331202303"></a>

## Next pages — waf_signatures / 212122210203 / 4

- [software_settings.waf_signatures.automatic](resources--securemesh_site_v2--reference--group-017.md#canonical-1321112103302132-0311330003120012-3202222023122102-2222222333321323-2313231200132032-0001101331002033-1112022212311112-3010310111203233)
- [software_settings.waf_signatures.manual](resources--securemesh_site_v2--reference--group-017.md#canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1321112103302132-0311330003120012-3202222023122102-2222222333321323-2313231200132032-0001101331002033-1112022212311112-3010310111203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012312302101311-3103311213321122-1321321131230013-2101123123103023-2300221222300333-0223130133210330-0222100211002213-3011110120213210"></a>

## software_settings.waf_signatures.automatic — automatic / 230112113013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- software_settings.waf_signatures.automatic

<a id="canonical-0212033110121121-2221203031112002-2220323311203110-0030030011131000-3111101323112121-0232233032330133-3310101210113132-0223222330203000"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
automatic = {}
```

<a id="canonical-2010102232220323-0012313223311020-1120003033322102-0331201230111312-2110120312033110-1130311030121332-3223131203002112-2103120110010112"></a>

## Direct properties — automatic / 230112113013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000113002230300-1203102112013112-3222012321103113-0000321223001010-1201200220200300-0012020022323001-1230212101133103-0322013023322021"></a>

## Next pages — automatic / 230112113013 / 4

- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312011231220302-1212233233212032-1203203011010202-2012211322013231-2230200003121233-0112033323032023-0110000211002111-0112021011203310"></a>

## software_settings.waf_signatures.manual — manual / 103111131001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- software_settings.waf_signatures.manual

<a id="canonical-3221332221221101-0102121321103320-3211111203002301-0020222103103013-0212322131112023-2102311100231210-3223200221222230-3321213001123032"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
manual = {}
```

<a id="canonical-3300111022111030-0223221200230203-0211313301202232-2023032211220202-1001030113311213-3102210232323021-2233311121002212-0323333112202213"></a>

## Direct properties — manual / 103111131001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313331101302133-0012223120320303-2001012000333312-2110213001030122-0233233131211320-0322003320233221-1210101111112213-0002230200123033"></a>

## Next pages — manual / 103111131001 / 4

- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2131133220012220-0222312200312320-2332131230210133-2323001331232302-0312012222021021-0213223112001210-2110313123231232-3312321032110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021321101013023-0113201331010321-1223313213301203-1102122020330113-1301201111333122-0100102230121200-1102022310221121-2123111100031331"></a>

## timeouts — timeouts / 322320321212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- timeouts

<a id="canonical-0123232131001300-0011323332301222-1331300031113011-1313301030310003-1101033221200021-3320311312021033-0002021220231213-3010101123210223"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210201022013022-3010101101010213-0321203021022011-0303021113330030-2223120120201123-0220122021311112-1320031202301311-2100203010002121"></a>

## Direct properties — timeouts / 322320321212 / 3

<a id="canonical-1110023132023112-1100103232120003-2233121122233201-3102321030123121-2111201002203301-0010022121120230-1321131213311030-0102133312300210"></a>

<a id="canonical-2121232011321202-2203023312230112-3303230112222221-1230213322110100-3310010221010320-2100323100130101-3223312013132232-1113332030013223"></a>

## create property — timeouts / 322320321212 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3130133011022332-2230123123000100-2021330331212132-1122103310110301-1102211033201131-1022213332002323-3032001030113000-2101002021101032"></a>

<a id="canonical-2203113031000212-0112123303133102-1223122313012112-3113313200030100-2121231321120220-1302230133321100-0120032323022320-0122330223201332"></a>

## delete property — timeouts / 322320321212 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2213013133300101-2020313200320200-3012022211113211-2233211130000331-2213130010313222-2230123120212320-3231211022000323-1002223123231322"></a>

<a id="canonical-0132231103132232-3122232312322012-2302013003030012-1210320221330230-2000032033023003-3003132313233103-2031010100321102-2211110300112103"></a>

## read property — timeouts / 322320321212 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0003313021213233-2220003330111102-0010200221002312-3332230211002222-1001101301312301-2030130303103123-2203003122030000-2303003002323021"></a>

<a id="canonical-3030130203203002-2032020321032231-2112122021211220-1203313333132202-0233331331333011-3203210013330102-0010331002203313-3202220002133230"></a>

## update property — timeouts / 322320321212 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1120230002220032-1013203123113332-2333122120102031-1202032303303110-2132023032110022-3112332032123010-3120110023203123-0020110222010123"></a>

## Next pages — timeouts / 322320321212 / 8

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122103030011130-1130031223203121-2310123120110211-3103203033131122-3303003100310103-2001220130103132-1220332013012012-1220133013020132"></a>

## upgrade_settings — upgrade_settings / 020213210320 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- upgrade_settings

<a id="canonical-1012300130023020-3121103011122011-3200231111011031-1132302032012101-1022231300310311-1322000133202211-3332113131103223-3020202230200002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for upgrade settings.

Upstream description:

Specify how a site will be upgraded.

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
upgrade_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323132310222330-3020220110313010-1330231211233200-1310131203011302-3120221101230113-3111003310211012-1333313012233011-2100310110301323"></a>

## Direct properties — upgrade_settings / 020213210320 / 3

- [kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300): complete subsection reference.

<a id="canonical-2213120032030313-2213212222022121-1300303113230013-0001231302113213-3103231133113130-0233330311331331-3212112130302233-0323011012222311"></a>

## Next pages — upgrade_settings / 020213210320 / 4

- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101321200011100-1112313033210132-2002311330030222-2231031201011003-0200110130023201-3231111111102333-3212102122231321-3220231320113002"></a>

## upgrade_settings.kubernetes_upgrade_drain — kubernetes_upgrade_drain / 320012020030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- upgrade_settings.kubernetes_upgrade_drain

<a id="canonical-3220131333120013-2202322022131000-1000010222103213-2013222300320010-2122210011001020-3001103121202113-0103222211322232-1121002031102322"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210210303132220-1233300022001301-1030133321001323-3112313231011303-1212333032231113-1023232132133122-0032123131022102-1320320302033233"></a>

## Direct properties — kubernetes_upgrade_drain / 320012020030 / 3

- [disable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3222130313222010-2113200000020330-2312232021233113-1330033220323210-2200321101300231-3032233121210023-1120230112112230-0000123202210303): complete subsection reference.

- [enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331): complete subsection reference.

<a id="canonical-2222002000233332-3223100103212312-3322131211220131-2132020200132132-0102033213230103-0113233130322202-1102231202230320-3133110322200030"></a>

## Next pages — kubernetes_upgrade_drain / 320012020030 / 4

- [upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3222130313222010-2113200000020330-2312232021233113-1330033220323210-2200321101300231-3032233121210023-1120230112112230-0000123202210303)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3222130313222010-2113200000020330-2312232021233113-1330033220323210-2200321101300231-3032233121210023-1120230112112230-0000123202210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222110020213203-0232323010123321-2113233123100030-3311113321100322-1100003000310332-2200013031303313-0033223221220301-0120133320212023"></a>

## upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 030001311332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2321012113022223-1013233322221032-2113023113030230-2001203212122013-2011023203033332-0222230313023003-1203131123300310-2120212331312303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-3300212033211301-3312331112122001-1200103301330330-0010212010030113-0023112110310302-1123010003332011-3100332312222230-3313102333332011"></a>

## Direct properties — disable_upgrade_drain / 030001311332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201223320102220-1221113200110133-3330120303323011-2221211302310223-3113331133330321-3030131301013232-3000100223303020-3113201132032111"></a>

## Next pages — disable_upgrade_drain / 030001311332 / 4

- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133031003022130-3230200220200113-2330220221333030-2333221222232301-3203113332112223-1302303123103101-3310331200132222-0203332033230330"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 011022213121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-1230100330002323-0110122001332122-1021203213231021-3221200310110203-1033300111021200-0012331102201322-1221222030211313-0331001023233011"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300030022123013-0302213232301130-2200232312332323-0011123223011200-0303021232231011-2330122120132201-3332120110100121-0230300303301200"></a>

## Direct properties — enable_upgrade_drain / 011022213121 / 3

- [disable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202): complete subsection reference.

<a id="canonical-0301130333233111-3220210110033323-1111202321111333-0031212120321321-3021121210313002-0211300001100023-3020301101123012-1130123320001132"></a>

<a id="canonical-3030323100203231-2331330210332002-0211130222122210-3020030121132113-1030221120121231-0001203102300022-2322000131123013-1301222122000330"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 011022213121 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2032210002330002-1330002321132010-2012221131221013-1303010103223013-3331213222211230-3333021232103212-2020101321123123-0202030321330320"></a>

<a id="canonical-3332212003212102-3003032232202102-0313332211300031-0130112022000133-3230012300030002-3022221233111122-0300131012230123-3233013112023101"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 011022213121 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-3320032100232312-3031311210022132-0222313022022211-0331213102230030-0300103210203120-3100122111030010-0031300323002322-3101233312001331"></a>

<a id="canonical-3212030001332301-2212310032130013-1220212203321012-0320332202303203-1313321311223333-0230122210113110-3232302211220031-3223213203003213"></a>

## drain_node_timeout property — enable_upgrade_drain / 011022213121 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130101021010320-3202311123201221-3231221223003212-1033030023012101-2021121332113111-2233131230332330-3131002210023130-0321000131233100): complete subsection reference.

<a id="canonical-3332311321201120-3022223300303220-0112200213321131-0023120011133213-3312212213210332-1131323233322213-2332211022133110-3321311200201201"></a>

## Next pages — enable_upgrade_drain / 011022213121 / 7

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130101021010320-3202311123201221-3231221223003212-1033030023012101-2021121332113111-2233131230332330-3131002210023130-0321000131233100)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003211021111312-1223001131001333-0323011312221231-0201203233201011-0033200032303100-1222213300012310-2110323333201111-3303200301130212"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 021112312021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-3102022031121003-1323013230331033-2222213102030110-2003103132303331-0312302233301213-1113000221111210-3201010320331001-0021110130020230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-1100220130130020-1122023013311300-2301003211112321-1330032313002213-1201120330012033-0300010020013301-3103203320210310-0002012021031133"></a>

## Direct properties — disable_vega_upgrade_mode / 021112312021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222203301303031-1102303301003002-2333102031132222-2333130112201011-0331001012030112-1203102231300213-3001013211130120-3013202130300320"></a>

## Next pages — disable_vega_upgrade_mode / 021112312021 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2130101021010320-3202311123201221-3231221223003212-1033030023012101-2021121332113111-2233131230332330-3131002210023130-0321000131233100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123331230202233-1110031020131100-3123003032111111-1220331221031202-0322323022333020-2202020332133201-2210212032312333-0031002022030131"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 113312120121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-1013033301121202-1312030331131013-1222303323220300-0002030002232133-0131001133211203-0322110002310211-3132322030130112-1233320312203302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-2232113311233333-2221323201023312-1032310111211203-2032220212102123-0302001213323130-1220230111033101-3333300131112212-0233321200013200"></a>

## Direct properties — enable_vega_upgrade_mode / 113312120121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020331112202033-0013033012233011-2112211302230200-2030023111120212-2311033311121202-1101110020212201-2202202010230300-1212113130002211"></a>

## Next pages — enable_vega_upgrade_mode / 113312120121 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201033311323312-0112303311222322-1013233210001021-3303231332110320-2122111213120300-3233202232022202-0102310321012232-3031310122312312"></a>

## vmware — vmware / 302220032103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- vmware

<a id="canonical-0013121331333101-0110031232203120-1332123300101323-0213331300203303-3102322233201100-0101131220111322-2213230000203031-1133112232132230"></a>

Type: `"object"`. single nested block, Optional.

VMware Provider Type. VMware Provider Type.

Upstream description:

VMware Provider Type.

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

Terraform syntax:

```terraform
vmware {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212103331333212-0101231020101110-1310023312103311-1110222300322333-2301130230230103-1023200233203312-2023201001002120-3313301001232030"></a>

## Direct properties — vmware / 302220032103 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133): complete subsection reference.

<a id="canonical-3323031103333013-0123313303121113-1100120322001333-0230100310033101-0003212133301120-1200313323011323-0001101201310023-3331011112021010"></a>

## Next pages — vmware / 302220032103 / 4

- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232110020133112-1030322100231002-0323310122210123-3122010001310110-3000330202000213-3322113111123221-1321003100113211-1021321231220111"></a>

## vmware.not_managed — not_managed / 111302311112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- vmware.not_managed

<a id="canonical-3212320032322213-2211221212101101-0030112233232122-3121231223331210-1111233221202320-1300010012113133-0321233132100213-3211203131021331"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

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

Terraform syntax:

```terraform
not_managed {
  # Configure direct properties listed below.
}
```
