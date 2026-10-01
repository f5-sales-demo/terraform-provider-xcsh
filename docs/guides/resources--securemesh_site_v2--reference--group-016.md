---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3110200012013330-1011233331120110-0332132113333131-2132210103030002-2022000010033312-3302230323313203-1100322233323033-0122020102103311"></a>

## openstack.not_managed.node_list — node_list / 200332002323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- openstack.not_managed.node_list

<a id="canonical-2110022221221312-0101320313103300-1233113013000200-3001111123002100-3011100220210120-3203003322222323-0230131313323312-1212313120010310"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

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
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120202311233133-3131320202310313-1211221123033031-1133212231333102-1320000300022132-3132000202112333-3001030323223021-0201023230120132"></a>

## Direct properties — node_list / 200332002323 / 3

<a id="canonical-1322331111302103-0221112022330010-0300023333311030-3232311021203333-0210130220302233-1320123220211103-3303021332211313-3133030331312120"></a>

<a id="canonical-0132322301201112-3010012113222333-1301332103021123-1012221310230003-1232013023122112-2230110121212210-1022201332302030-0233311131111103"></a>

## hostname property — node_list / 200332002323 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132): complete subsection reference.

<a id="canonical-1110031220332222-1012123312232302-2302103110320033-3102130002221110-2011102213302023-2332130101022201-3123010123213101-3323202213303022"></a>

<a id="canonical-3132003131110223-0013210302223303-3233010311113010-0202203130303232-0230013130030100-0133222332212201-1220111033111212-3122220121330313"></a>

## public_ip property — node_list / 200332002323 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2123221013120333-0021223130221310-0331203303231111-2300102013303223-3233032100012011-3111102032220300-1120201320001233-0200023221212332"></a>

<a id="canonical-3232030111311120-1120023212223011-0001023300302302-2020303002211300-3103220323002201-0130013312103120-3220313122313130-0111131122220331"></a>

## type property — node_list / 200332002323 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-1221110212022233-1320113030320211-1302001131012201-0300121030031033-0202221111102221-1030323313230100-3323000232103032-2233303022230330"></a>

## Next pages — node_list / 200332002323 / 7

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301311331032330-2233000312013010-0023222002312223-0231111020032032-1201230332202201-3321320230033021-2302223203230002-0201103212312101"></a>

## openstack.not_managed.node_list.interface_list — interface_list / 020122232032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- openstack.not_managed.node_list.interface_list

<a id="canonical-0310323121211222-2103231002130122-2220220203200233-0211001023320111-0322000121121010-3303202103333120-3023232310211223-2012120220221013"></a>

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

<a id="canonical-0021011111310312-2001311133332133-2123321103301210-2132200100101303-0110300133100022-0332302211111300-1112112011133233-3222101331302003"></a>

## Direct properties — interface_list / 020122232032 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133): complete subsection reference.

<a id="canonical-2113111103022211-2133211022301112-1000302202312010-1301220202021100-1012010301333131-2212101321122233-2100223302300323-1103112320321123"></a>

<a id="canonical-1202212101100021-1321122223132201-2322200233333320-3031101112123211-0212333211133113-2011122312301223-3321133032101203-0230000103010312"></a>

## description_spec property — interface_list / 020122232032 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-016.md#canonical-3230113312212313-1330233220212100-3203311232020130-3012012231200201-3230331132330213-1320323302203122-0310130112001121-3002210101233322): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-1330203230212300-2103311130233331-1103133013123230-3202300010230033-3332211103233312-2123113330001101-0223131232213221-1022112130113323): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113): complete subsection reference.

<a id="canonical-3210023300301021-0303013233213222-3321002033232231-1320332111220011-0230311303030121-0021002232222303-2211310033312202-2300301332313003"></a>

<a id="canonical-2322013323312023-1033322111001323-0222102310213122-2212132320113012-0032020202312310-0101223211221002-0301013111201230-1030211203320232"></a>

## is_management property — interface_list / 020122232032 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2102003002112121-3121021002202213-3332001322031311-0323022321120112-0223212230032300-1023311011312201-0122231121011020-2321001300013202"></a>

<a id="canonical-0030023031231312-1220123232311110-1212003323023210-1303303322022320-3323332013202121-1331102032012032-3031322100333231-3200022223113123"></a>

## is_primary property — interface_list / 020122232032 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1112101020300012-0103023203230210-3112311112232131-2223331212312313-1330132031021302-3211212332101030-0001233213100120-1101322110012231"></a>

<a id="canonical-0101111030131102-2333031113112022-2101021120010111-1012023222313210-1322321322223133-3003203103213201-2320202011230123-1132031312301013"></a>

## labels property — interface_list / 020122232032 / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-016.md#canonical-3101113233000123-3211233033203031-3230112121131320-3021032013030320-2211023113300001-0002133320033300-3100233023020320-2021301303200202): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-3222123322300001-0110133122012110-0111313122300121-0010031303320211-3130112222100330-1130133233002013-2230032311102211-0211232000012200): complete subsection reference.

<a id="canonical-2010220332112102-2003332213113202-3002133231210110-0110302131103002-0022012212230321-3323031022330033-0003310102133023-3123100031313201"></a>

<a id="canonical-3121222133130230-3103033123100110-2022131113121031-1113300120202123-0101011120132022-2002103200122021-2013100310302133-1023323320110031"></a>

## mtu property — interface_list / 020122232032 / 8

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

<a id="canonical-0333102331110023-3131123113022320-0320110323310103-0333000032132102-0200111311310223-0203322230302031-1323133131200000-0322203133212202"></a>

<a id="canonical-2033123333313233-2202003231003002-3121330020012320-3110112220010202-0132101013112303-1310230333113123-3032333223123103-3213001003000111"></a>

## name property — interface_list / 020122232032 / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2332311111121233-1023200032232001-3122133030130130-1123313102323230-0200221122211002-3230133312201011-0122132200032030-0013213232231232): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2110313201023123-2310200020000232-1311232001030303-2222100103031323-0211313131023030-2300033103012320-1233022303303331-0011200211020330): complete subsection reference.

<a id="canonical-3033101202223200-1110333202033113-1000013321203121-2003312312330012-0010000233111230-0110122123012102-0323220322310300-1110101323321013"></a>

<a id="canonical-3302331133000022-0330333233002003-2123301031002232-0333300122132131-2213200232133122-0130011321113200-1122131110123112-0010003022212001"></a>

## priority property — interface_list / 020122232032 / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0210012301212222-0111212112333310-1312311313231131-0323100032010230-0030031102013120-1311123222203330-0103003211100302-2202220211100133): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-2031330202202022-3331000031001100-1200313120202122-3313120233311233-3322331012221303-0023120222112321-1130321210023233-3011313203321011): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-2112231302320012-3123103320233130-1022202103132302-0121321121313211-3111011202010200-2222311200010133-3021032133223331-0323330130002321): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-0020330021133011-0122223222303323-2332322122302031-0103032300320231-1322120221332020-0002202212303101-1102001132023322-2011031201011130): complete subsection reference.

<a id="canonical-3030200330201321-0130220323210023-0111122112011333-3210202201201222-2322311303301110-3030323022302301-0112032032011101-3332300320331200"></a>

## Next pages — interface_list / 020122232032 / 11

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- [openstack.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-016.md#canonical-3230113312212313-1330233220212100-3203311232020130-3012012231200201-3230331132330213-1320323302203122-0310130112001121-3002210101233322)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-1330203230212300-2103311130233331-1103133013123230-3202300010230033-3332211103233312-2123113330001101-0223131232213221-1022112130113323)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-016.md#canonical-3101113233000123-3211233033203031-3230112121131320-3021032013030320-2211023113300001-0002133320033300-3100233023020320-2021301303200202)
- [openstack.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-3222123322300001-0110133122012110-0111313122300121-0010031303320211-3130112222100330-1130133233002013-2230032311102211-0211232000012200)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- [openstack.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2332311111121233-1023200032232001-3122133030130130-1123313102323230-0200221122211002-3230133312201011-0122132200032030-0013213232231232)
- [openstack.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2110313201023123-2310200020000232-1311232001030303-2222100103031323-0211313131023030-2300033103012320-1233022303303331-0011200211020330)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0210012301212222-0111212112333310-1312311313231131-0323100032010230-0030031102013120-1311123222203330-0103003211100302-2202220211100133)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-2031330202202022-3331000031001100-1200313120202122-3313120233311233-3322331012221303-0023120222112321-1130321210023233-3011313203321011)
- [openstack.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-2112231302320012-3123103320233130-1022202103132302-0121321121313211-3111011202010200-2222311200010133-3021032133223331-0323330130002321)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- [openstack.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-0020330021133011-0122223222303323-2332322122302031-0103032300320231-1322120221332020-0002202212303101-1102001132023322-2011031201011130)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213323313201110-0201331230331232-0231221330110100-0002021323032211-3122232012211213-3122132332021212-2001102030320300-0202132302022312"></a>

## openstack.not_managed.node_list.interface_list.bond_interface — bond_interface / 312112223211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.bond_interface

<a id="canonical-1101112320333211-1213301001313000-0100202023313011-1121202113133320-3203333021212003-1232311120032021-2020133222303211-1132300133202211"></a>

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

<a id="canonical-0030022130122210-0121200103001330-2302312313102323-0021123102212002-2301103332331122-1011012220001120-2320303001203201-2302332132013100"></a>

## Direct properties — bond_interface / 312112223211 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-016.md#canonical-0313313111010312-1332003030032011-3332132101103332-0030123133323310-0100012132121011-0121233023100132-1102032120312012-1021002230013121): complete subsection reference.

<a id="canonical-1200121322232002-0012131332032231-2203132031023323-0221323331010322-3112222100302331-1331100031013303-0222021122112201-2101001201010013"></a>

<a id="canonical-3112313021221031-1311101222201313-3012222302033102-3313131003232023-3003033200000321-0021003123101102-3112130022202103-0002231030132303"></a>

## devices property — bond_interface / 312112223211 / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-016.md#canonical-3102223220023321-0111220332033310-1212222201122120-0302303033312332-0302021232311303-3210110002003212-0133322031022003-0032123012223003): complete subsection reference.

<a id="canonical-2113102110120322-0011023310220310-2233112122301130-1110033003020033-3022331322322220-1013321202030311-0101330021232010-2310220111320023"></a>

<a id="canonical-0100331232200311-2132020033201020-2021213103211022-3322033020333310-0121323232333130-1232033122323113-1210321313001111-2030313120203110"></a>

## link_polling_interval property — bond_interface / 312112223211 / 5

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

<a id="canonical-3123030311230312-2223203313103321-0310121332230121-2113331310013222-2030222322000320-0102100011323222-1132132130102321-1032311320030330"></a>

<a id="canonical-1122120031333332-0311331300011211-3131012323123320-3022331223210112-1121110213103211-3021300132102133-3212231231110331-1310301003202002"></a>

## link_up_delay property — bond_interface / 312112223211 / 6

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

<a id="canonical-1031033000212131-0003103011013322-3203202330303102-1322003020322232-0121121222121110-1131231010323111-0233310102212033-3010300032212233"></a>

<a id="canonical-3323103030003021-2120322310323120-3102103322201302-2301220003120023-3303321021302230-2031133031202122-2132231132332101-3133302020300311"></a>

## name property — bond_interface / 312112223211 / 7

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

<a id="canonical-0213001202023323-3313322311103011-1113022110213233-3212032113200312-1102110110021103-2330203220013233-2233201131321031-2211112221223222"></a>

## Next pages — bond_interface / 312112223211 / 8

- [openstack.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-016.md#canonical-0313313111010312-1332003030032011-3332132101103332-0030123133323310-0100012132121011-0121233023100132-1102032120312012-1021002230013121)
- [openstack.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-016.md#canonical-3102223220023321-0111220332033310-1212222201122120-0302303033312332-0302021232311303-3210110002003212-0133322031022003-0032123012223003)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0313313111010312-1332003030032011-3332132101103332-0030123133323310-0100012132121011-0121233023100132-1102032120312012-1021002230013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112220003123110-1111000223232012-2111220320320133-3123202320021321-0131301101123101-2100121031212230-2322032210113303-0230301123101200"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 323120013301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- openstack.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1110120221223333-3301101230330102-0200202221010312-0210031211200101-1322001301102122-1111213103010221-0211011032222210-2231033332322332"></a>

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

<a id="canonical-0223021301302011-1032011323031321-1110103002320110-2210033223003100-1303133233020113-2223013101013333-2132300101011320-0222021002223302"></a>

## Direct properties — active_backup / 323120013301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330123112323023-3331020332112230-2113100113121112-0023212300301103-0331200333032201-1100220002203020-3301113103000003-1202011220300300"></a>

## Next pages — active_backup / 323120013301 / 4

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3102223220023321-0111220332033310-1212222201122120-0302303033312332-0302021232311303-3210110002003212-0133322031022003-0032123012223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020012312020120-0233113300133112-1213111010210020-1120320113303103-2201303213113130-2203001033310020-1312101203201203-0101201321312302"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 101021023020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- openstack.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1013111110012220-0132231333203211-2213011220100303-2331122310223222-0010011221020333-2333111121330022-3311213322301303-2033111223133110"></a>

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

<a id="canonical-0230313321210300-3210331323013100-3001101303321201-1320101323000211-2030112230213210-0230103012010130-2032320102132122-3201033013113332"></a>

## Direct properties — lacp / 101021023020 / 3

<a id="canonical-1101300022010120-1030320023013120-0131222021330102-3322303012220131-0233231103133130-1121312332102030-2202000102200101-3210103300011033"></a>

<a id="canonical-1222112310100023-2033022202213110-0130030301230001-3200033111331230-0111303213303203-3312212032022013-1131112021022311-2023011032222021"></a>

## rate property — lacp / 101021023020 / 4

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

<a id="canonical-2310312121333301-3220201133320312-3210033100030230-1210331031330002-3110322230133322-3223021203322201-1301023100102332-3302332022221232"></a>

## Next pages — lacp / 101021023020 / 5

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3230113312212313-1330233220212100-3203311232020130-3012012231200201-3230331132330213-1320323302203122-0310130112001121-3002210101233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103120133222303-0013011323301200-3332332331230023-3232231221020332-2030132320110131-3223020013231200-3032310121222131-2212010332113100"></a>

## openstack.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 232220121300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3123020230132302-0312211133100030-3102303011030233-3330030212022012-1122302102022131-2003203012231020-2203300003120100-1132301122333331"></a>

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

<a id="canonical-2133132333033032-1031333112030201-2313030203231310-3123003013001323-1001210231211022-2230002221032303-0322300011132303-1102323122113313"></a>

## Direct properties — dhcp_client / 232220121300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232033112323231-0322102131010130-0113301302103101-3011322300223023-0311020310223233-2200121030300010-0111321321001013-1330211113202132"></a>

## Next pages — dhcp_client / 232220121300 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122130131121222-1303103012020302-3003330223303302-2032102221030330-2020333311030003-3033033003202203-3032032222223020-0000101011231311"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 110223020021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2130113031122220-2112031113002113-2312332133031122-1212112200030200-2031330012313232-2211201133333012-1330011323312031-2112330132000030"></a>

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

<a id="canonical-3201230011030022-2103101233202310-3113030332210011-2133131131203101-1230213312102210-0330113030023110-0102323100201231-1222030101220302"></a>

## Direct properties — dhcp_server / 110223020021 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-2302311110230322-2331300112102023-1310333113223023-3101113100210310-2121020330213321-2023003011112013-0022332313112100-2102032032111321): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3321233201103332-2232003203311011-3323102233123000-3133102222202121-1313333012121122-3333000030031300-3020011013332300-0332303113201033): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210): complete subsection reference.

<a id="canonical-1220101221202033-3303012120332230-2222202203120110-3112001301231311-1021003203223333-1202001230302331-1031101023301203-2011133032111313"></a>

<a id="canonical-3333133002332011-1231003312021333-3321033213102323-3031202013212332-2320002103011202-3210130233020322-2223031220303310-1130210121102011"></a>

## dhcp_option82_tag property — dhcp_server / 110223020021 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0323031113123031-1112033003333333-0200301133123221-1211121010002222-2112323010231112-2223020312001312-3013200000002022-0320120033223111"></a>

<a id="canonical-1121133200232300-3121222232131002-2330002021111303-1231303211012002-0201112331012023-2010010302032322-0213002112322210-1130333130220233"></a>

## fixed_ip_map property — dhcp_server / 110223020021 / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-2031202130120130-3313113302311303-3313032330222000-3222123212121221-1122131300323113-0111313233133023-2323232012022323-0133302123301321): complete subsection reference.

<a id="canonical-0030213322021233-2011130003103211-1020112133022232-2200111222232212-3320212131312300-3322030023233202-1010230120000123-3211001020101333"></a>

## Next pages — dhcp_server / 110223020021 / 6

- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-2302311110230322-2331300112102023-1310333113223023-3101113100210310-2121020330213321-2023003011112013-0022332313112100-2102032032111321)
- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3321233201103332-2232003203311011-3323102233123000-3133102222202121-1313333012121122-3333000030031300-3020011013332300-0332303113201033)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- [openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-2031202130120130-3313113302311303-3313032330222000-3222123212121221-1122131300323113-0111313233133023-2323232012022323-0133302123301321)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2302311110230322-2331300112102023-1310333113223023-3101113100210310-2121020330213321-2023003011112013-0022332313112100-2102032032111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021300010310212-3110110120031011-0023210310030201-0011102310333233-3313113223102203-2300123332320302-1012303000223233-0311201120231312"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 233033233021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0311003302301031-0320002032011110-0022133013201111-1310203100033332-0033312301211112-1223132323122200-2101123032113022-0210101310230310"></a>

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

<a id="canonical-0020322233123213-2330322202203312-0101110132001003-0211122220333213-1313002233120012-3201333322303220-2001300122320131-1032020010010330"></a>

## Direct properties — automatic_from_end / 233033233021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120320010211320-3012110020112211-2223033110220110-1033021201110103-0002022210210321-3111022220323011-2011313302132101-3332202123332313"></a>

## Next pages — automatic_from_end / 233033233021 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3321233201103332-2232003203311011-3323102233123000-3133102222202121-1313333012121122-3333000030031300-3020011013332300-0332303113201033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001032101233021-3332210223220001-3201011113111133-0332230232321100-3212002333230002-0033102231032300-2113122233302221-3030223203133020"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 111311230100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3231221111001022-2123030000222011-2033020311303331-3300002222303230-2210313133132012-0132212103003301-1010012203001112-0013002321010031"></a>

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

<a id="canonical-3032023311323331-3131232110123311-1010312322121032-0211112131321112-3022001332133012-0012201232120321-3002101232021103-3321223223200120"></a>

## Direct properties — automatic_from_start / 111311230100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110013002221233-3133102312330103-2121132202323112-3233331121220012-2310330220233010-3132132301222332-3233011010312211-3132123233323102"></a>

## Next pages — automatic_from_start / 111311230100 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333023113231323-0311030321011213-0231032302201321-2031031001130131-1200031011333223-1000323122202201-2200003133132102-1112000003221010"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 302020330312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1203223323032001-1011113333003031-0222031202310323-2132333332303230-3221222212223011-3233033101100210-2103200133233322-2100300002201302"></a>

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

<a id="canonical-2332030020313102-3323320312223022-0320213221323222-2331000133230302-1011320231121332-3011222200002033-1102110223110012-2301003310103133"></a>

## Direct properties — dhcp_networks / 302020330312 / 3

<a id="canonical-3002130012200132-2003023102112232-3302333222133201-2021332021133311-3230303111012110-0031323213301112-2300312302133111-2202312330332122"></a>

<a id="canonical-0113102122011032-2211332320221312-2330312232313110-2332031303231230-2023033211212100-3122132011132223-3211013020321311-3201131010333133"></a>

## dgw_address property — dhcp_networks / 302020330312 / 4

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

<a id="canonical-0311103220033302-1321303313321000-3333111320212321-1333111102013122-2130030312301132-1013331112100021-2321310233130000-1200032100002300"></a>

<a id="canonical-0302210300032300-3201220310011033-0023003010303331-3301102003112033-1201210302322330-2231311223031022-3311322102200220-1330030000203210"></a>

## dns_address property — dhcp_networks / 302020330312 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2030321121302210-1132311132101102-1332201022213230-1000023223002301-1233320130001203-0101122001212122-2012003123113200-1221130033031333): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-1120122231310133-2021022122030222-2132203033012332-3333021233011031-3012221110122212-2022113012013020-3121302100021232-1211030212021330): complete subsection reference.

<a id="canonical-2120120110021120-0031003011011221-0320302130012333-2010200223010012-2213220200130121-3022312212132231-1011212220223131-0032323020310311"></a>

<a id="canonical-2122303000101121-2332102132231211-0031101033000323-0201212232231202-2200301003321320-2311000320103202-0120223220313221-2033111020120323"></a>

## network_prefix property — dhcp_networks / 302020330312 / 6

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

<a id="canonical-1103301303031133-1323232331023211-2001010332302133-2221223112331110-3123111030312020-3131022033303121-1213021021103310-1131112323332310"></a>

<a id="canonical-1102030110230023-3213023321203010-2312031013012022-1003312132133322-1313201021231130-2210131012311220-2211211012331012-0133030203002202"></a>

## pool_settings property — dhcp_networks / 302020330312 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-016.md#canonical-2330021230203122-3212121232232300-1003313223311032-2010132213210202-1032231111212023-1111213002222030-2200023232201132-3112020313023310): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-016.md#canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023): complete subsection reference.

<a id="canonical-3001033103332231-0021332133113321-1032030112323113-3121300211130310-3023130202132321-1322330123222212-3202210322333013-2000300331203213"></a>

## Next pages — dhcp_networks / 302020330312 / 8

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2030321121302210-1132311132101102-1332201022213230-1000023223002301-1233320130001203-0101122001212122-2012003123113200-1221130033031333)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-1120122231310133-2021022122030222-2132203033012332-3333021233011031-3012221110122212-2022113012013020-3121302100021232-1211030212021330)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-016.md#canonical-2330021230203122-3212121232232300-1003313223311032-2010132213210202-1032231111212023-1111213002222030-2200023232201132-3112020313023310)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-016.md#canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2030321121302210-1132311132101102-1332201022213230-1000023223002301-1233320130001203-0101122001212122-2012003123113200-1221130033031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202010330122300-1201332331301302-2210012120010211-3102331000001222-1120121011313030-3221311331233011-3131102321302103-1323231103221310"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 031111300013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0201220113110120-2010020113311330-3033322333203210-0002122122020203-0113032002103230-0213131213210030-0223312033213232-2122102032202331"></a>

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

<a id="canonical-3300003130213000-0111120330322111-1000302213233300-2110313122201202-2022200022001201-3320331211123131-0021321312003023-2233032203033230"></a>

## Direct properties — first_address / 031111300013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332103221110232-3003022203003002-3022123012231023-3320102321302223-0121323222133220-0003203002003212-3213231313021101-0202311232330220"></a>

## Next pages — first_address / 031111300013 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120122231310133-2021022122030222-2132203033012332-3333021233011031-3012221110122212-2022113012013020-3121302100021232-1211030212021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212313120312332-3221201213220223-1233300120311022-3231133301112221-2301331011013202-1312030303030321-3030020131233221-1111231101221110"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 001113211020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2303020121330123-1001110321201201-0131200322221002-3021230222033000-1002130031300313-3303100302020113-2313223213032232-0122310131301033"></a>

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

<a id="canonical-2200302032003023-3131003123110132-0031302230030022-2311211313312200-0202100002112001-1001202103123311-0131230020333033-0301111031200202"></a>

## Direct properties — last_address / 001113211020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110231320221331-3313312212333231-2310321021012021-0312000001111020-0210303002200331-3320312332200132-2113131331333232-1303210010211330"></a>

## Next pages — last_address / 001113211020 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2330021230203122-3212121232232300-1003313223311032-2010132213210202-1032231111212023-1111213002222030-2200023232201132-3112020313023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030012111132100-1000202032302321-3221033203112003-2032230202202231-0023302011301111-0210330010213021-0211130022313220-3331103001330112"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 232301023301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2312000301221222-2312221230331123-1022331233110221-1310222323102303-2133131201130122-0321122113313220-2323330212003120-0330123012122010"></a>

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

<a id="canonical-1100120002122331-2211101300000001-0011032021301022-3123200013222330-2330232203111110-3121312002223132-2232310311122030-2023210332330123"></a>

## Direct properties — pools / 232301023301 / 3

<a id="canonical-1023021121220321-1232100213001200-3230232202330130-1331200330333000-0301220322311231-0213312012211312-3113023212220321-3120033313131230"></a>

<a id="canonical-0310320323022300-1211201031321110-0313200301000201-0012100103121130-2031010012200222-3232320311323311-1333233003221023-2213223022323022"></a>

## end_ip property — pools / 232301023301 / 4

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

<a id="canonical-3200132320001310-3022303223303322-2132012002201003-2023320330021311-0023100130330013-3332312203213001-0122222203202302-2220201221212111"></a>

<a id="canonical-0011000030110303-2313220313333110-1031222203021320-2303320322023130-0003002223232200-1101010330332212-0002213012013312-0100213313203230"></a>

## exclude property — pools / 232301023301 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1202001113232012-3112302030211222-3122230032203013-1020030021220131-2332010330132002-1031330320011033-0200302133210310-2332301130102123"></a>

<a id="canonical-2013202333311102-1002011233212200-0300322221021021-1220002310301030-3310022013123332-3321322133002310-2221101012133130-2221321312321210"></a>

## start_ip property — pools / 232301023301 / 6

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

<a id="canonical-3220131221223123-1011223020012311-2010200031331101-1211233023103310-3211330231230312-2103122011202112-3320201311212002-0030022303102330"></a>

## Next pages — pools / 232301023301 / 7

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030332212221000-2231220220100010-3332120220232301-1023311202310030-2010233221313122-2310100123312022-3032013222231222-2212123133123101"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 112203023323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2331020133330021-0331230100331001-3003331130133133-3213222013011132-1120003311312011-3013332201031313-3130022333331322-1022312002322202"></a>

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

<a id="canonical-3011303000303210-2200111011130000-0330020222332231-0122200221210323-3122213103023212-3022103110332023-0022101313223022-0213201200112321"></a>

## Direct properties — same_as_dgw / 112203023323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121213331122103-1321131320211110-1120023220001322-2200321313322333-3202223212003331-0231312222022121-2113312000011220-1100031030131112"></a>

## Next pages — same_as_dgw / 112203023323 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2031202130120130-3313113302311303-3313032330222000-3222123212121221-1122131300323113-0111313233133023-2323232012022323-0133302123301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013303100300203-2332133300120112-0330110220002133-2102112013312230-2021323031232131-2232230131130232-1112230113230110-1321313212033012"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 303013210311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2121302300323232-3033312130231012-0302301130123033-3002313201003320-0100311003321311-0220313113113133-3010231022031203-0123000310113102"></a>

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

<a id="canonical-2213120320010313-0132231012230311-2000002033313230-2301310101233202-3222002102310201-3223032310133133-1221310300123020-0121310013012221"></a>

## Direct properties — interface_ip_map / 303013210311 / 3

<a id="canonical-0303030333301303-2233003002011332-1121123311122212-0330302012033210-1231210301221123-1122123032322210-0121301101320122-0123303113311012"></a>

<a id="canonical-3012321113211032-0122202223331020-1121232320332130-2112300122110300-0002303031312221-0332031212203310-0112310003302123-2021121231233333"></a>

## interface_ip_map property — interface_ip_map / 303013210311 / 4

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

<a id="canonical-1332331030101010-2331220332110031-0002202323023130-2121111130023103-3002120211011331-0012230123103303-0222120210312131-3210311101101030"></a>

## Next pages — interface_ip_map / 303013210311 / 5

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330203230212300-2103311130233331-1103133013123230-3202300010230033-3332211103233312-2123113330001101-0223131232213221-1022112130113323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312221203131120-0210200202121222-0030103022032112-3031201321010002-1102303310023220-2102003330210021-1221103002212301-1322020210301332"></a>

## openstack.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 021132120220 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3033113113133001-0230321130230020-2202231302303211-0110130213311233-0303223313030002-2220230203330100-1110313332023103-1230231320221112"></a>

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

<a id="canonical-0230322211102303-2110133302112233-1302211100001323-1030221201321220-1331021300230121-2022122110033233-2033312130310301-2032333013232021"></a>

## Direct properties — ethernet_interface / 021132120220 / 3

<a id="canonical-1303032220333330-0302212301311110-3202110100113311-2120011233202130-0301303333320311-0133011232033033-0000202110023030-3012000102001001"></a>

<a id="canonical-2002221100010000-2121033213112323-0330002013031231-2031110231121223-1211111202033202-1032013312213000-1002002101332013-2303203030031202"></a>

## device property — ethernet_interface / 021132120220 / 4

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

<a id="canonical-2000022321113303-3202203302113123-3013320321100231-1032303101301131-3201303210223302-0001103132013110-1331133133002122-1021000223202310"></a>

<a id="canonical-2100130213010012-3211022120130020-2233333020030321-2001232000222311-3201333202212202-2030301221323333-3103120110220033-3212100020033301"></a>

## mac property — ethernet_interface / 021132120220 / 5

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

<a id="canonical-1223120213303231-3011323332023011-1332030202321212-3322102022030201-0211223310231333-0011021223102023-2212211213023131-3331110312203110"></a>

## Next pages — ethernet_interface / 021132120220 / 6

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022302021212312-1321201031100331-0033211031000112-1302202100220010-0113323203212110-0310013130130210-2313323013030002-2221000211221133"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 112331323031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3003013202210011-1321002233223101-2332231012022101-2203013101213032-2213033010223113-2202102312001033-0030002023031130-1000330332113201"></a>

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

<a id="canonical-1210132220223202-1221110203121322-2032022300131112-3021312010131202-3132230110100131-1101110221023323-0203010122013133-2113123122100011"></a>

## Direct properties — ipv6_auto_config / 112331323031 / 3

- [host](resources--securemesh_site_v2--reference--group-016.md#canonical-3101302123110213-2302210111232033-2033131103122200-2313321121312001-1300002132312021-2213200322300011-2202312130002102-0311220131010110): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001): complete subsection reference.

<a id="canonical-0231310022320010-1213010331300022-2321120020323031-2022010022003301-2002013010203103-1000203000320031-3112102111112231-2312321233001303"></a>

## Next pages — ipv6_auto_config / 112331323031 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-016.md#canonical-3101302123110213-2302210111232033-2033131103122200-2313321121312001-1300002132312021-2213200322300011-2202312130002102-0311220131010110)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3101302123110213-2302210111232033-2033131103122200-2313321121312001-1300002132312021-2213200322300011-2202312130002102-0311220131010110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123310130200020-2320323233130131-0133121222010000-2012210311300211-3023202303313301-1223322130323000-2300320010023103-0100000211120031"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 001212031113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-3210030021330003-0313101311311301-0103202331332101-3030032112213221-2332300102232213-0012322002013223-0020200112003213-3303301013111120"></a>

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

<a id="canonical-2131002132231220-1320112121220100-3223003033212132-3021332022132030-2100310302022033-1212221333122301-0002003030223220-0001210032200110"></a>

## Direct properties — host / 001212031113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322021101131021-3232121303113000-1002023230213220-2030012030320130-3331011023221231-3110122002322010-0320301330231003-2001002032003103"></a>

## Next pages — host / 001212031113 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323101120210121-3310003213101110-3333102001321302-2120331201212223-1010020010103320-0001230312233020-3032333311112121-2020112020032110"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 011113132111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-3103003233133202-0230112212301131-1200331120211210-1030300121010100-1313212020301132-1001132320211220-3200303301220311-3131100213332101"></a>

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

<a id="canonical-1002010111032013-2022312232232013-3112121211230102-1112012302211031-0213230200020230-3123111113222033-3320223133102311-3033122112303030"></a>

## Direct properties — router / 011113132111 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232): complete subsection reference.

<a id="canonical-2100321022021333-2232221230133232-1303120113100102-0310123302320332-1101301002023300-2300103011300211-1210221310311310-0213112202331103"></a>

<a id="canonical-3020302113101012-3003301221000221-0022302113211001-2132213203331012-2100120113330223-2213211222323232-3102221233312323-2311332311121220"></a>

## network_prefix property — router / 011113132111 / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201): complete subsection reference.

<a id="canonical-3311003200001133-2311303102023200-1122230202022122-3321032203210102-1033220301323331-2312222030231002-3110102203311012-0101033310023300"></a>

## Next pages — router / 011113132111 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111032013021103-3230102322013003-1003302032013133-0203101223310110-2330210123210303-3222320012101201-1111002103111112-0333222121321013"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 033310030210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1332311200303030-1121112303200321-0100222120103230-3011333032000122-0332033123113031-0020232301031112-0013230213102323-0310010200121210"></a>

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

<a id="canonical-1223001301022213-2231110310303002-2231001100022030-2223013023122310-3331201112000302-2331013323312200-3303103002121022-0321322021311313"></a>

## Direct properties — dns_config / 033310030210 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-016.md#canonical-0333223122131111-1321002232230113-2030111233101131-1013020301312013-2012011332223310-0030232101132133-1330310000103200-3312112110323330): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132): complete subsection reference.

<a id="canonical-0032202021230203-1332002300333230-3300010333131300-2032002023120203-3013100201320310-2011320101132122-0310202131013013-3333301012203102"></a>

## Next pages — dns_config / 033310030210 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-016.md#canonical-0333223122131111-1321002232230113-2030111233101131-1013020301312013-2012011332223310-0030232101132133-1330310000103200-3312112110323330)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0333223122131111-1321002232230113-2030111233101131-1013020301312013-2012011332223310-0030232101132133-1330310000103200-3312112110323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321322123221003-0312223223331323-0212013321010023-2110033300031202-0031331122231021-0211201312031320-1202021301131213-3332230013030201"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 301203130102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0023120322021322-2301300120113302-3212123002322202-1111230010113011-3333211331200212-0120223102003213-3232011232213322-1130332023222010"></a>

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

<a id="canonical-1320301212102102-3323331123310123-2303331301123212-1102330111321302-3212303033332010-1201210232322311-2001113100023020-3000320012200303"></a>

## Direct properties — configured_list / 301203130102 / 3

<a id="canonical-2003332233010223-1303333113232111-3131203133302320-1000322102100232-0032313200301313-0122030231133300-3223133101330332-3202113212132203"></a>

<a id="canonical-2330302130203011-2131022330033032-1031320023233201-2310201232223333-2330330130231131-1122033023010133-0013201231321303-3123020231200031"></a>

## dns_list property — configured_list / 301203130102 / 4

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

<a id="canonical-3311201033013133-1030111323231310-3100013000020123-2301211110023233-1301103032100110-2230110032321333-3301131013113233-3133131310303112"></a>

## Next pages — configured_list / 301203130102 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111211122312311-2003311013123232-3032210110332311-2111300232201102-1023131133000313-0003223010300221-2213131120023102-0311002231120321"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 122233332101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-3213132210132301-3030013330110303-3323300021131232-2320300321001201-3000332002000301-2001112200100313-3303100012011131-2233321311330011"></a>

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

<a id="canonical-1201102313120223-0113232301021212-2230213002301312-0213013233102030-3001221003310331-3022022122010033-2313212221002213-1032100210103000"></a>

## Direct properties — local_dns / 122233332101 / 3

<a id="canonical-1303011120221230-2112022112312102-1330302113320322-3131103110100203-1211001300130331-0021223003300012-2103220233131331-3310332000223231"></a>

<a id="canonical-2020101222032212-1230131032201031-0133302120033102-0133331100323123-3133311211301333-1310031003222121-0303002112023123-2310010232301032"></a>

## configured_address property — local_dns / 122233332101 / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3120101230221223-0101223001131310-2100002221320213-1031101332220310-2313332131011122-3332111233323300-0021100331123221-0122203023222001): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3110023203302313-2030221003021200-2202333000122302-3312000300223033-1122032220211022-1103220201110333-0223021213320223-0222011331320020): complete subsection reference.

<a id="canonical-3330131000130330-3012132220101101-3200301111223323-3212301223033320-2112310202213222-0003303321100010-1312322103101312-2130100323121101"></a>

## Next pages — local_dns / 122233332101 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3120101230221223-0101223001131310-2100002221320213-1031101332220310-2313332131011122-3332111233323300-0021100331123221-0122203023222001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3110023203302313-2030221003021200-2202333000122302-3312000300223033-1122032220211022-1103220201110333-0223021213320223-0222011331320020)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120101230221223-0101223001131310-2100002221320213-1031101332220310-2313332131011122-3332111233323300-0021100331123221-0122203023222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120101100303333-2133101132311313-2021120320003111-2031232232013003-0322030300330211-3032201310003133-1330322302323311-3330333030000103"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 111223220020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2000330102310302-1320023302312302-0002212332003300-0002332132301212-3021130223023002-0321100310323231-3222232012132210-1033023103033330"></a>

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

<a id="canonical-0100223312021033-1020001323121201-3222123323011302-0120321100123300-0131233232133031-0120202321000331-3131300330223033-3132012220102303"></a>

## Direct properties — first_address / 111223220020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332331130022333-0130221131120023-2010330022132113-0221111221001001-0032300311220011-2321323001201231-1002313302013031-3212311322133220"></a>

## Next pages — first_address / 111223220020 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3110023203302313-2030221003021200-2202333000122302-3312000300223033-1122032220211022-1103220201110333-0223021213320223-0222011331320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113101001222123-3113021221222033-2220333021131110-3101311101301013-0030003110221211-3301331102022111-3100122211030210-1033313103301101"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 313231202302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2331332001033010-0212300302211000-2100120110230030-2220033123310130-1013011323001003-1111201310021100-3323201332011122-1032002101323200"></a>

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

<a id="canonical-2331030130123221-2302210110313111-1212002323333221-0011003133130010-2321322102023332-0332032332102211-0030130120001013-1221231202112032"></a>

## Direct properties — last_address / 313231202302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332030321212310-2012121021332030-3320120210010011-2120300003112001-1310320332222100-3213133012012032-1101112032202131-2222001023001102"></a>

## Next pages — last_address / 313231202302 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113010132120131-0320133101222200-3020103112212131-2120213012121032-0002033223221000-1131111312313333-0312113020312132-3030133113213330"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 031031010310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0020332200200101-1010113111101103-2103113201300123-3310002133003023-2012321302333303-1002112223032031-1202112023210013-1002123011213011"></a>

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

<a id="canonical-0001113211022302-3230313132332002-0323022233203211-2003123011021300-1111313330011331-1022013013110313-3012123300010112-1202103102232221"></a>

## Direct properties — stateful / 031031010310 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-0101230031010110-0202202331010300-1113112001210101-1123203121202331-3201211011322323-2110320010212110-3121311012213011-2311130132001030): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3223222203121003-3022102222321010-2113202330312222-0022230323211122-0003003133231000-1222121020020323-3312000002322220-2102121211130103): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203): complete subsection reference.

<a id="canonical-1220331032200230-3130020022103033-0123122322320031-0332000311210212-0310332331021113-0101003210013123-0133330211221332-3001313010032222"></a>

<a id="canonical-3221322233103000-0313031012030021-2322123110310100-1232330022023033-0323303120032321-1010102202333211-3232212232310222-3231311210023231"></a>

## fixed_ip_map property — stateful / 031031010310 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-1233203333000302-3323203123130330-1302220232001130-0023020001111022-0120031233220223-0313230211120030-3332301111301311-1122212112311013): complete subsection reference.

<a id="canonical-3031322221112031-3111102011220023-2320223210013133-3100211132301103-1011311212331100-1001102111223031-2132233022302112-2103321310213221"></a>

## Next pages — stateful / 031031010310 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-0101230031010110-0202202331010300-1113112001210101-1123203121202331-3201211011322323-2110320010212110-3121311012213011-2311130132001030)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3223222203121003-3022102222321010-2113202330312222-0022230323211122-0003003133231000-1222121020020323-3312000002322220-2102121211130103)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-1233203333000302-3323203123130330-1302220232001130-0023020001111022-0120031233220223-0313230211120030-3332301111301311-1122212112311013)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0101230031010110-0202202331010300-1113112001210101-1123203121202331-3201211011322323-2110320010212110-3121311012213011-2311130132001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301231200012203-0133121123330313-1201310002010120-0330020221002023-3232132210032300-0200220132021132-3333012002322300-0013221231022311"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 021231130123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3010230202023302-3103322332031021-1232010102110303-2312230230323303-1220033221030000-1103133012221223-2303110132201122-2001222033232321"></a>

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

<a id="canonical-0331220131333332-3322113020011021-2033221211311222-0122300231122011-3201330120333131-2103201123331120-2202120010002123-3132200212202113"></a>

## Direct properties — automatic_from_end / 021231130123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322131202230120-0211333300100213-2012122201221200-2222033003301330-0322031223103302-2110033110013132-2322320122330111-0203230003121230"></a>

## Next pages — automatic_from_end / 021231130123 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3223222203121003-3022102222321010-2113202330312222-0022230323211122-0003003133231000-1222121020020323-3312000002322220-2102121211130103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033111122312123-1122303021011121-0010203223122123-3111110233000300-3011210013031130-1313330312301113-0212112203202201-3022222001110110"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 033020132313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3130103331300132-3001111001221033-1200101220221100-1030232201320213-3333030030011110-1201331312203022-1313220031321123-0012010300330022"></a>

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

<a id="canonical-2003231311021001-2123112112301001-0123113313303121-1011322021003101-3302132013202033-2212023231220230-1103020112031132-2121120021312223"></a>

## Direct properties — automatic_from_start / 033020132313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130213020300232-1301313201101222-2222122010303013-0213013011223130-2332201220011113-2300323312000113-0223321030111013-1231120133113100"></a>

## Next pages — automatic_from_start / 033020132313 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220010300001130-2220300121200233-0300213033231031-0011130111011311-1012233011020022-1133320232313211-1013311100302103-2132110012000301"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 011112012002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0320013031003230-1101202100321100-0332102303223003-2232102202200302-1223130223110332-1210212311132210-3210233210202232-3303013203211001"></a>

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

<a id="canonical-3313232331221303-1103230132113301-0313012300323003-0011101031312112-2122103211112303-3203030112113322-3332112010020003-3112100222311122"></a>

## Direct properties — dhcp_networks / 011112012002 / 3

<a id="canonical-0130122133331210-0130320303213123-2131301033032010-0300113313033003-3103212222130103-0310022211010002-0223302303131000-1201030220232213"></a>

<a id="canonical-1102223023022203-0010202023333003-3230011320000003-3220320111111303-2102323131020302-0011320120110211-0212322132102303-2203111320311013"></a>

## network_prefix property — dhcp_networks / 011112012002 / 4

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

<a id="canonical-2222302133330120-0021112223212012-0011123132202022-1323200300132231-3202100112323321-3312022030133100-1320022203222021-0231221201222321"></a>

<a id="canonical-2301330013000001-2301023210031300-3332021332032103-0133330222201030-0210111231320001-2020021003312321-1232023233131033-2320032200322033"></a>

## pool_settings property — dhcp_networks / 011112012002 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-016.md#canonical-1120202130022001-1221101210333031-2113300303310313-0220131201023200-1333212323233323-2103331103132212-3112303330300303-0233011131222120): complete subsection reference.

<a id="canonical-2121003322333012-1313010221131303-2010100023111330-0311231330033100-2321223002020002-1230210122220120-1230132222132313-1110011020031021"></a>

## Next pages — dhcp_networks / 011112012002 / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-016.md#canonical-1120202130022001-1221101210333031-2113300303310313-0220131201023200-1333212323233323-2103331103132212-3112303330300303-0233011131222120)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120202130022001-1221101210333031-2113300303310313-0220131201023200-1333212323233323-2103331103132212-3112303330300303-0233011131222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230121312220121-2031211200100333-2122010011021333-3122301201200203-0200333232013021-1213010003203232-1022132233231111-0320312120002003"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 202300221000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1313021222022102-3100330011020110-1223023200133020-2202021321131332-2330312322303332-0120100001122221-2032231133020012-2322030011100220"></a>

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

<a id="canonical-0231030121032303-0313330031321123-1230120020301333-0010112110013202-0322102001301022-0023021201331322-2002112003131302-0113101112013100"></a>

## Direct properties — pools / 202300221000 / 3

<a id="canonical-1322331321113333-2100030022230031-3212013330230023-0011213201121310-0132111313131213-1201203103223332-0022232132133122-3013020332022022"></a>

<a id="canonical-3002001033003213-3203313013010102-2130331013100021-2130321013000122-2330210331222332-0133122133101020-3211013001131120-3211112033223213"></a>

## end_ip property — pools / 202300221000 / 4

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

<a id="canonical-0303021133000203-3031111023133023-3230110100133212-0110110133303031-2000232301133100-0301013332113101-3210011220123131-1121022230311313"></a>

<a id="canonical-1233102233302301-1222101200132221-1201003303023203-2110020113322322-1013002302133322-3010101121303002-0030230202011222-0132123321133312"></a>

## start_ip property — pools / 202300221000 / 5

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

<a id="canonical-2313210123013323-1111312000123003-1323332331112201-0331122321222133-2223120233220201-1121102212110000-2010333222313232-3222122312311003"></a>

## Next pages — pools / 202300221000 / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1233203333000302-3323203123130330-1302220232001130-0023020001111022-0120031233220223-0313230211120030-3332301111301311-1122212112311013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112313210013032-1010131013003202-0113101122300213-2201011203211321-3220000321131213-1333223100022033-3102213220011132-2313013010122320"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 003113020301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3133120020122330-0222011130310000-3212301032022101-2121112021230000-0301230311331301-3133111101210033-0103310100301100-1022210320122330"></a>

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

<a id="canonical-0132322123331221-1102221111102223-2230212311211123-3320010212323201-0222320003230131-1331210313321012-1310023232102111-2023310201021230"></a>

## Direct properties — interface_ip_map / 003113020301 / 3

<a id="canonical-3212133303011123-0222332100010122-2002210102102313-3220231101303111-0220030012103312-3203121120112033-3310333030123121-0000331222120100"></a>

<a id="canonical-2103111200100312-1300302123300232-2112101212330013-2313330032302121-3201331323130131-0300111302212030-1233001230103010-2121122233133002"></a>

## interface_ip_map property — interface_ip_map / 003113020301 / 4

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

<a id="canonical-1201121000120000-3122123333110032-2313232131020020-3122231001022211-2201110322300012-2333233121323021-2302302133232100-1200331231110321"></a>

## Next pages — interface_ip_map / 003113020301 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3101113233000123-3211233033203031-3230112121131320-3021032013030320-2211023113300001-0002133320033300-3100233023020320-2021301303200202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022210230303023-0203321333310000-0111310130030122-1213020130203320-0302021102332023-1203103311210011-3013133122001023-3233120102222120"></a>

## openstack.not_managed.node_list.interface_list.monitor — monitor / 120023302023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.monitor

<a id="canonical-3002020310100322-0033123131303030-2321312302001100-0203032321210222-2300100110000330-3202120310033003-1012333101323122-1133322010100220"></a>

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

<a id="canonical-2200010010011023-1331301333313013-3233300022010021-1202121103303201-1002111003112223-3133011223201013-2302121101330220-3131200113312101"></a>

## Direct properties — monitor / 120023302023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323103133001012-0111331033312220-2033011221221120-3023013201020022-2010131221101102-3012301002013210-1100202313331033-1131210121312131"></a>

## Next pages — monitor / 120023302023 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3222123322300001-0110133122012110-0111313122300121-0010031303320211-3130112222100330-1130133233002013-2230032311102211-0211232000012200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102303000232121-3113321311131332-2211222011002022-1021203330200021-2110111101320021-3121231102331003-2300231002310133-2310112011020313"></a>

## openstack.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 321311322032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0102301002233113-3313121101120203-3002200232123313-1111001300200222-1210001200100320-1131220101312010-1313200230210002-3031011323211133"></a>

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

<a id="canonical-2112003222303122-1300302100212300-1212222121302301-0330321123203231-0323301230301302-2330111310330003-0210222111203010-2202032202333201"></a>

## Direct properties — monitor_disabled / 321311322032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003130023311220-3310023310110101-1110233112313110-3032311312203100-1132011030032201-1323323201310233-2022002303110113-0323221302203023"></a>

## Next pages — monitor_disabled / 321311322032 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311312311300223-0132330103031000-0013213011131023-3121131102003113-3321002200132201-1220231230311321-3220113200021132-1212230322331033"></a>

## openstack.not_managed.node_list.interface_list.network_option — network_option / 310310100211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.network_option

<a id="canonical-1221022311300321-1302203013101332-2003330120001330-2210102123031300-2033332021210133-1312230133210033-1313223301031113-2023220310201111"></a>

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

<a id="canonical-2201032332320122-2021302231103332-1130322022200312-1002230030120120-2103311330022213-2113231011000311-3022221310200212-0233022211201311"></a>

## Direct properties — network_option / 310310100211 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-016.md#canonical-0032303011130121-0221223212220232-3032001112003222-3210312230001020-0123130031231220-0301113022022211-0111203021013101-2121013332011302): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-016.md#canonical-1121221200221011-0102213201120211-3113331233332011-2121300111333010-1310331032212121-3100032230120321-2101221333230302-3200112330103012): complete subsection reference.

<a id="canonical-0310331102023232-1131232303113233-1202223033322230-0022011111021230-1303222003212222-3011121030232100-0320122010022102-2301100330333120"></a>

## Next pages — network_option / 310310100211 / 4

- [openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-016.md#canonical-0032303011130121-0221223212220232-3032001112003222-3210312230001020-0123130031231220-0301113022022211-0111203021013101-2121013332011302)
- [openstack.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-016.md#canonical-1121221200221011-0102213201120211-3113331233332011-2121300111333010-1310331032212121-3100032230120321-2101221333230302-3200112330103012)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0032303011130121-0221223212220232-3032001112003222-3210312230001020-0123130031231220-0301113022022211-0111203021013101-2121013332011302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322303113003112-1033332123212131-3100003003011013-0212212211200303-3231231023103231-1333303310133212-0020302230021112-2100333130013032"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 221231112001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0311002312212122-3111320231110010-3232231201222130-1103301232212110-2013212003032121-3320230210333001-2113301023322113-3203113021133112"></a>

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

<a id="canonical-1300323231220110-3312001111231231-0301010103211210-0222030030000000-0231020232111300-2010023021122302-3331311112332320-1101210013231230"></a>

## Direct properties — site_local_inside_network / 221231112001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330001012233131-3312131002220313-2013320310012023-0313202112100322-2012203003022023-2022223013331312-3001003112010211-0130330113332000"></a>

## Next pages — site_local_inside_network / 221231112001 / 4

- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1121221200221011-0102213201120211-3113331233332011-2121300111333010-1310331032212121-3100032230120321-2101221333230302-3200112330103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333211211011221-2313203321212110-1321122012223230-2110133331133320-3011003333202001-3233231100322111-2103333330121031-3122332033001231"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 231132130022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- openstack.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2022212123101203-1121021231131322-0122210030311202-1302000201333301-3330020300012222-0203210321012103-3203212000323323-3201023032333012"></a>

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

<a id="canonical-1012002320322023-2020020331211302-1000213021003011-1032032110332123-0021020323000130-1231321320101131-2232213131220223-2033122133011103"></a>

## Direct properties — site_local_network / 231132130022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023330001000310-2022220030030301-1313133321121232-0232011021131121-0033021000231320-1133122233023313-3002301120121103-3132002220232201"></a>

## Next pages — site_local_network / 231132130022 / 4

- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2332311111121233-1023200032232001-3122133030130130-1123313102323230-0200221122211002-3230133312201011-0122132200032030-0013213232231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231033311110103-2232121213021131-0233320311122222-3213212310213322-3332113310103101-2013033010223100-0020112113030330-3021011213010001"></a>

## openstack.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 323233332030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3301210213120202-3113030232301101-0211310101031210-2101022233211000-2312113233013010-2111223012210331-2032032013032200-0003210012200330"></a>

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

<a id="canonical-3303321310102101-2212333002231332-1233211312121202-3022210333131022-1032110103213030-0101003330102031-1010332103313222-0002212110030221"></a>

## Direct properties — no_ipv4_address / 323233332030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212331133203332-0111000003131222-0320320111230210-3011323011232102-3201011002202300-0202200233333131-1031102233201133-0331020203010231"></a>

## Next pages — no_ipv4_address / 323233332030 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2110313201023123-2310200020000232-1311232001030303-2222100103031323-0211313131023030-2300033103012320-1233022303303331-0011200211020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210320003230211-2020102023323302-2022112332300012-0133300203021201-2011030102003323-2211032333210002-2001333233313121-2010003320330023"></a>

## openstack.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 032031003101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3222332033122201-3211302030330232-3333123003212103-1021003121110323-3302101320112313-1020213222020001-0021013103130200-1210330011210311"></a>

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

<a id="canonical-0020013103330212-2102310003001112-3322311201102120-2312323210130011-2211210120333203-1330203132133232-2030133103030112-1310302233310110"></a>

## Direct properties — no_ipv6_address / 032031003101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332030302030013-0000200022323023-0130110202220002-2313103101220323-3311331213012230-1212110300210211-2322110000123322-0300130013201312"></a>

## Next pages — no_ipv6_address / 032031003101 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0210012301212222-0111212112333310-1312311313231131-0323100032010230-0030031102013120-1311123222203330-0103003211100302-2202220211100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110103023332131-3301132123331001-2330120222103031-3032230000210231-3311120110301000-2232300303003032-1020123203032212-1331220232330022"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 133321212012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1221110132120232-3011130131213201-1200322023322102-2122000120133233-3322100003323301-1232222011133203-0330221211011300-3032213220230331"></a>

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

<a id="canonical-1003113001123103-3221320333000230-3303102032001110-3023201200320123-2130030221333023-2231210030230223-0131300120302210-2011311131211010"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 133321212012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223231131232210-3110300103220333-0010313122233133-3301332312332011-1031331230031011-0333233233121111-1232301103101233-2013002123123133"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 133321212012 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2031330202202022-3331000031001100-1200313120202122-3313120233311233-3322331012221303-0023120222112321-1130321210023233-3011313203321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122213000122230-3000023311003030-3130301203323102-1020131030233001-3111300101113321-3121031110320123-3311010320132223-2000222100222332"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 313002213221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0130231101201023-0003213102020302-2220300023020211-0013323201310120-1313232332022233-3310130130001111-1102230113330032-3031221321100313"></a>

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

<a id="canonical-3223312203201122-3203023211033202-2212122203222232-0311021311221323-1213202101311021-3131023122003201-1211101013220132-3020022213331310"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 313002213221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310010212121113-1031232012102312-1022102010202131-2132101201222330-0110013323011230-1211112000001213-1321310310000010-1323302011211332"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 313002213221 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-016.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2112231302320012-3123103320233130-1022202103132302-0121321121313211-3111011202010200-2222311200010133-3021032133223331-0323330130002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133131112002020-2130102231103030-3001321121231032-2033311303012022-1220010002310213-3212333001322022-1030110133110001-0212012100333221"></a>

## openstack.not_managed.node_list.interface_list.static_ip — static_ip / 131202221311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
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
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
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
