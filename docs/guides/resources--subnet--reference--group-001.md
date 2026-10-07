---
page_title: "xcsh_subnet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet reference."
---

# xcsh_subnet reference

<a id="canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- Property reference

<a id="canonical-1302331033331311-1102233300233010-0023322123000202-2122201112102023-0033131200212023-0233031313311132-0302021310011133-3112331103130220"></a>

### Direct properties for `xcsh_subnet`

<a id="canonical-3321033210013131-1133031020003301-2120032123032202-3211333123222031-1303002201312200-0101001030033031-2211132132320221-3301112113331320"></a>

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

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-3211320301332132-0311221202303332-0010032231210300-2000132130330101-1303202203222131-1231323330211010-0321233002313030-1020202233221222): complete subsection reference.

- [connect_to_slo](resources--subnet--reference--group-001.md#canonical-2221112002303323-1312003022203230-1110230122111201-1302202132320223-1132333021002002-1133133132221231-0220310313110131-0031223211313300): complete subsection reference.

<a id="canonical-0222010230232023-2010201132203220-0312101310003231-3002100320110000-2313330031133331-3230101120111101-2223102000111231-0012113300212022"></a>

<a id="canonical-3132301031012132-3231032212002013-0012212312101331-0331312103200120-0013303213300022-2122310201021121-3012222311033130-3303031102320130"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2333131333102210-3121001202212213-3331232231101223-0121030102120121-1112030123032122-2232122231211020-3000330131023121-3211133131123201"></a>

<a id="canonical-0130330012011000-2122231000300312-1213323021221002-2011303210233132-0213303111101310-0112020232231012-1210030330100231-3030321130302210"></a>

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

<a id="canonical-0231233213221330-1323212323012002-3211300121321122-3110112112230323-1222113303110333-0110302103201300-3000230300102320-0330123030023100"></a>

<a id="canonical-3221310020303233-1011033202131030-1323010223110111-0213032123201213-3023132021311103-3200322333333002-0121133233123211-0302331022031033"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated_nw](resources--subnet--reference--group-001.md#canonical-0103130132011023-3120212213021202-2003313000030100-2223133212222130-0211123130000313-0300010310210232-2313122120312201-1221310100311300): complete subsection reference.

<a id="canonical-1122210130121000-3121102221103112-0101210230322110-1220222122312110-2301331030030331-3330222031221033-1200132110100202-0231313310123022"></a>

<a id="canonical-0213320222123111-1320310010131110-3002332112213213-0201301301103301-2303011230330130-0220223311211013-3232200031231112-2303101020111233"></a>

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

<a id="canonical-2132022001120311-1323202310001302-0012310223221301-3122203213133321-3210121311132023-1133331012302312-1230011303333110-2213332133331103"></a>

<a id="canonical-3303202202133002-1200310300213002-2100230033323203-2223230332121220-1020132101032121-2223032130113221-0232002003301120-0312031323311010"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Subnet. Must be unique within the namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0212222223322113-1312003133303113-0330030121013211-1202101030303002-1200203111001321-2033222023002222-0322002201311131-2320332210122103"></a>

<a id="canonical-1233001200230100-0313023213100103-0223212322220031-0031230110323013-2122300320111023-2233213300020223-1303022031310021-1112001201330112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Subnet is created.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013): complete subsection reference.

- [timeouts](resources--subnet--reference--group-001.md#canonical-2320222121331132-1330210203133001-3132321032323002-2221201331031332-2113013120310311-0231130220332233-1213301033300322-3332123321212011): complete subsection reference.

<a id="canonical-2213012221331003-0110003021000310-2211112233233320-2132233030303211-2303020211023233-1313222123010230-2122330103302103-3302230202002211"></a>

### All schema paths for `xcsh_subnet`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--subnet--reference--group-001.md#canonical-3321033210013131-1133031020003301-2120032123032202-3211333123222031-1303002201312200-0101001030033031-2211132132320221-3301112113331320) |
| `connect_to_layer2` | [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-0113133323022310-2101233320103333-0211020213331020-3332020132202132-2010133320030303-2001021032112000-1221203001101132-1132131012301110) |
| `connect_to_layer2.layer2_intf_ref` | [connect_to_layer2.layer2_intf_ref](resources--subnet--reference--group-001.md#canonical-3100000110030220-0301211302330311-2133223013201220-0111000022311033-1302203202011211-0300030230230120-1111330312013023-2222211322203010) |
| `connect_to_layer2.layer2_intf_ref.name` | [connect_to_layer2.layer2_intf_ref.name](resources--subnet--reference--group-001.md#canonical-3202032123210121-3220330331002103-1023313220201101-2130320123200022-2113232210132312-3100033223312301-2022010110312111-1030332011102302) |
| `connect_to_layer2.layer2_intf_ref.namespace` | [connect_to_layer2.layer2_intf_ref.namespace](resources--subnet--reference--group-001.md#canonical-1201203133321020-1222232232233122-0323310122302302-3222300213223311-0221103203113133-3003322221033302-0211310100213111-0222331011123032) |
| `connect_to_layer2.layer2_intf_ref.tenant` | [connect_to_layer2.layer2_intf_ref.tenant](resources--subnet--reference--group-001.md#canonical-0320102111312221-2012131233203311-3031020232133113-1323203312000022-3302312321113330-0012011320311203-2122102200102003-0122330300003210) |
| `connect_to_slo` | [connect_to_slo](resources--subnet--reference--group-001.md#canonical-1311201320301333-0330203132203013-1201121302022031-3211102222111123-1333011123030121-3202100000121130-1200033303110023-1111112330122000) |
| `description` | [description](resources--subnet--reference--group-001.md#canonical-0222010230232023-2010201132203220-0312101310003231-3002100320110000-2313330031133331-3230101120111101-2223102000111231-0012113300212022) |
| `disable` | [disable](resources--subnet--reference--group-001.md#canonical-2333131333102210-3121001202212213-3331232231101223-0121030102120121-1112030123032122-2232122231211020-3000330131023121-3211133131123201) |
| `id` | [ID](resources--subnet--reference--group-001.md#canonical-0231233213221330-1323212323012002-3211300121321122-3110112112230323-1222113303110333-0110302103201300-3000230300102320-0330123030023100) |
| `isolated_nw` | [isolated_nw](resources--subnet--reference--group-001.md#canonical-0231320233111301-0023031002133331-2110203132223130-3233210100021220-0000010101023323-0221100003220232-2101311110001102-2323012132313002) |
| `labels` | [labels](resources--subnet--reference--group-001.md#canonical-1122210130121000-3121102221103112-0101210230322110-1220222122312110-2301331030030331-3330222031221033-1200132110100202-0231313310123022) |
| `name` | [name](resources--subnet--reference--group-001.md#canonical-2132022001120311-1323202310001302-0012310223221301-3122203213133321-3210121311132023-1133331012302312-1230011303333110-2213332133331103) |
| `namespace` | [namespace](resources--subnet--reference--group-001.md#canonical-0212222223322113-1312003133303113-0330030121013211-1202101030303002-1200203111001321-2033222023002222-0322002201311131-2320332210122103) |
| `site_subnet_params` | [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0212110131321303-1220211102001101-1110302332123230-3312012121121110-2311210213303013-2200320220012102-2003311010213132-1100330211122303) |
| `site_subnet_params.dhcp` | [site_subnet_params.dhcp](resources--subnet--reference--group-001.md#canonical-0100020112122211-1303203032102011-2011101033032210-3022121220231310-3030120003113310-2222221202321001-2231322113302231-2323330002130232) |
| `site_subnet_params.site` | [site_subnet_params.site](resources--subnet--reference--group-001.md#canonical-1003321012111212-2123030231011113-0310012231200201-2023220222101232-2130303321002220-3032121332300123-2021110020202122-3010020233221200) |
| `site_subnet_params.site.name` | [site_subnet_params.site.name](resources--subnet--reference--group-001.md#canonical-3233332100303030-0131223201233321-2203210033133100-3221131322002202-3131211310332002-2120330101231002-1231233033110001-1020210030131330) |
| `site_subnet_params.site.namespace` | [site_subnet_params.site.namespace](resources--subnet--reference--group-001.md#canonical-3330312111110203-2332230333013011-0211313233020322-1203103313012011-0121331111330300-3012300323332222-3331130130303132-2133213203311022) |
| `site_subnet_params.site.tenant` | [site_subnet_params.site.tenant](resources--subnet--reference--group-001.md#canonical-3122300232303011-2111321102110133-1231201102330100-3212002213001202-3122031302220323-0233112223201313-1230102233101110-2131133313001222) |
| `site_subnet_params.static_ip` | [site_subnet_params.static_ip](resources--subnet--reference--group-001.md#canonical-3323311302002220-0023101221021113-1033133202120230-2203023321122221-2232102123203031-0333121312303221-2222203131033123-1103232101331320) |
| `site_subnet_params.subnet_dhcp_server_params` | [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-3112130113013133-1013301203221310-0031300203210200-3211101100331130-0211111303332032-2113130030001102-2121130132333310-2212122301333112) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](resources--subnet--reference--group-001.md#canonical-2233311102012112-1232202213030303-0213221103231213-3311002302322100-2321033200120201-1300222322201031-0003030232130111-1002220122012233) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix](resources--subnet--reference--group-001.md#canonical-1010230332011311-0022312102110132-2200321110310123-3122030112033231-3313010231100121-0000213302012100-3003322110302320-1212221002103230) |
| `timeouts` | [timeouts](resources--subnet--reference--group-001.md#canonical-3013201002022112-2213122313310210-1212100223031121-1132300101303220-1330332033201331-1310233001320002-0101303311310222-3221030133333221) |
| `timeouts.create` | [timeouts.create](resources--subnet--reference--group-001.md#canonical-0330323133103210-3111033011233010-0003132021212220-1121203111131302-2132310101021021-1102020203130033-1220033303301123-3203211310313020) |
| `timeouts.delete` | [timeouts.delete](resources--subnet--reference--group-001.md#canonical-0213020113200302-2211011110233201-1133303000210320-1333222110310223-1320231010331030-3220120211123132-1231132031211122-0330112032313312) |
| `timeouts.read` | [timeouts.read](resources--subnet--reference--group-001.md#canonical-1323201233131213-0123323130323020-2123303031312120-3322323021201120-0122110220230022-2212210012220111-2332012131312302-2210030332222121) |
| `timeouts.update` | [timeouts.update](resources--subnet--reference--group-001.md#canonical-1200112131021230-3220231320212211-2301213101321302-3202220030020003-1212301020121110-1220131101133023-1300030323203331-3123120311022210) |

<a id="canonical-3211320301332132-0311221202303332-0010032231210300-2000132130330101-1303202203222131-1231323330211010-0321233002313030-1020202233221222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_layer2` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- connect_to_layer2

<a id="canonical-0113133323022310-2101233320103333-0211020213331020-3332020132202132-2010133320030303-2001021032112000-1221203001101132-1132131012301110"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: connect\_to\_layer2, connect\_to\_slo, isolated\_nw\] Configuration parameter for connect
to layer2.

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

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-0113133323022310-2101233320103333-0211020213331020-3332020132202132-2010133320030303-2001021032112000-1221203001101132-1132131012301110)
- [connect_to_slo](resources--subnet--reference--group-001.md#canonical-1311201320301333-0330203132203013-1201121302022031-3211102222111123-1333011123030121-3202100000121130-1200033303110023-1111112330122000)
- [isolated_nw](resources--subnet--reference--group-001.md#canonical-0231320233111301-0023031002133331-2110203132223130-3233210100021220-0000010101023323-0221100003220232-2101311110001102-2323012132313002)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
connect_to_layer2 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210322001200231-3130330010130111-0023013102332331-3103130310102130-0123212112302322-1023233322212320-1210222232121031-3102312232230012"></a>

### Direct properties for `connect_to_layer2`

- [layer2_intf_ref](resources--subnet--reference--group-001.md#canonical-0121201100331010-2020211020103022-2212030020121300-3231222332300013-3302101232320020-3312110332310300-3221131133221210-2112311313020300): complete subsection reference.

<a id="canonical-0121201100331010-2020211020103022-2212030020121300-3231222332300013-3302101232320020-3312110332310300-3221131133221210-2112311313020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_layer2.layer2_intf_ref` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-3211320301332132-0311221202303332-0010032231210300-2000132130330101-1303202203222131-1231323330211010-0321233002313030-1020202233221222)
- connect_to_layer2.layer2_intf_ref

<a id="canonical-3100000110030220-0301211302330311-2133223013201220-0111000022311033-1302203202011211-0300030230230120-1111330312013023-2222211322203010"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
layer2_intf_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322223132323000-3023200303120220-2320031122202132-3212102313110131-0231100112213230-2201011123310111-1203302121323213-3113301033200033"></a>

### Direct properties for `connect_to_layer2.layer2_intf_ref`

<a id="canonical-3202032123210121-3220330331002103-1023313220201101-2130320123200022-2113232210132312-3100033223312301-2022010110312111-1030332011102302"></a>

#### `connect_to_layer2.layer2_intf_ref.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201203133321020-1222232232233122-0323310122302302-3222300213223311-0221103203113133-3003322221033302-0211310100213111-0222331011123032"></a>

<a id="canonical-0211300103103202-2113300221203023-3202101333102001-1011132211303230-1310203120212133-3231220011311202-0321330212113322-2202301202322300"></a>

#### `connect_to_layer2.layer2_intf_ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320102111312221-2012131233203311-3031020232133113-1323203312000022-3302312321113330-0012011320311203-2122102200102003-0122330300003210"></a>

<a id="canonical-3023312112302023-1320133133122113-3202303211233001-2130331120013330-2210203220220332-3222211312303213-2113112023311230-1332002331003000"></a>

#### `connect_to_layer2.layer2_intf_ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2221112002303323-1312003022203230-1110230122111201-1302202132320223-1132333021002002-1133133132221231-0220310313110131-0031223211313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_slo` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- connect_to_slo

<a id="canonical-1311201320301333-0330203132203013-1201121302022031-3211102222111123-1333011123030121-3202100000121130-1200033303110023-1111112330122000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for connect to slo.

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
connect_to_slo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103130132011023-3120212213021202-2003313000030100-2223133212222130-0211123130000313-0300010310210232-2313122120312201-1221310100311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `isolated_nw` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- isolated_nw

<a id="canonical-0231320233111301-0023031002133331-2110203132223130-3233210100021220-0000010101023323-0221100003220232-2101311110001102-2323012132313002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for isolated nw.

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
isolated_nw = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- site_subnet_params

<a id="canonical-0212110131321303-1220211102001101-1110302332123230-3312012121121110-2311210213303013-2200320220012102-2003311010213132-1100330211122303"></a>

Type: `"object"`. list nested block, Optional.

Site Subnet Parameters. Configure subnet parameters per site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dhcp",
    "static_ip")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_subnet_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202212210301312-3233102313010120-0202103332110302-2031203033300030-3020231300231203-1303030223311330-0120032303301132-2023303020220130"></a>

### Direct properties for `site_subnet_params`

- [dhcp](resources--subnet--reference--group-001.md#canonical-3303233001212322-1212220213212332-1300302000210202-3332110310101303-2032200312331123-2000331200303132-0103011020313002-1223100020112112): complete subsection reference.

- [site](resources--subnet--reference--group-001.md#canonical-3102020100213213-0002031032202122-0201213213220111-2202302322212011-3032200003013331-1213031301300221-3110030030231300-0220330301021013): complete subsection reference.

- [static_ip](resources--subnet--reference--group-001.md#canonical-0000120300121031-2031001210121031-2323022232221120-0013310013233000-3222123111101312-0122021331012020-3112332211122123-1012123321213200): complete subsection reference.

- [subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-3222100203231312-0002211232302021-3221300102032233-3133331312322010-0103123030010313-2202020001001100-2330211233023033-3003221313122203): complete subsection reference.

<a id="canonical-3303233001212322-1212220213212332-1300302000210202-3332110310101303-2032200312331123-2000331200303132-0103011020313002-1223100020112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.dhcp` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013)
- site_subnet_params.dhcp

<a id="canonical-0100020112122211-1303203032102011-2011101033032210-3022121220231310-3030120003113310-2222221202321001-2231322113302231-2323330002130232"></a>

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
dhcp = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102020100213213-0002031032202122-0201213213220111-2202302322212011-3032200003013331-1213031301300221-3110030030231300-0220330301021013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.site` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013)
- site_subnet_params.site

<a id="canonical-1003321012111212-2123030231011113-0310012231200201-2023220222101232-2130303321002220-3032121332300123-2021110020202122-3010020233221200"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300222212020033-2113031303130222-3030023321102321-3330101110033302-2030121301012011-1121032331302013-0033310003111123-0212111220201121"></a>

### Direct properties for `site_subnet_params.site`

<a id="canonical-3233332100303030-0131223201233321-2203210033133100-3221131322002202-3131211310332002-2120330101231002-1231233033110001-1020210030131330"></a>

#### `site_subnet_params.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3330312111110203-2332230333013011-0211313233020322-1203103313012011-0121331111330300-3012300323332222-3331130130303132-2133213203311022"></a>

<a id="canonical-3001112022323121-3121011032023102-1012010220201031-2112030103210210-1233022230001311-2100332201312220-1100122002031133-1333031122223200"></a>

#### `site_subnet_params.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3122300232303011-2111321102110133-1231201102330100-3212002213001202-3122031302220323-0233112223201313-1230102233101110-2131133313001222"></a>

<a id="canonical-0301332233020333-1233313000121001-3120102222310302-0012330212233231-1210213130133213-1302330322000200-3231310020300012-3233132303321212"></a>

#### `site_subnet_params.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0000120300121031-2031001210121031-2323022232221120-0013310013233000-3222123111101312-0122021331012020-3112332211122123-1012123321213200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.static_ip` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013)
- site_subnet_params.static_ip

<a id="canonical-3323311302002220-0023101221021113-1033133202120230-2203023321122221-2232102123203031-0333121312303221-2222203131033123-1103232101331320"></a>

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
static_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222100203231312-0002211232302021-3221300102032233-3133331312322010-0103123030010313-2202020001001100-2330211233023033-3003221313122203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.subnet_dhcp_server_params` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013)
- site_subnet_params.subnet_dhcp_server_params

<a id="canonical-3112130113013133-1013301203221310-0031300203210200-3211101100331130-0211111303332032-2113130030001102-2121130132333310-2212122301333112"></a>

Type: `"object"`. single nested block, Optional.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

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
subnet_dhcp_server_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013313132203200-1233131011320212-0313033111123131-3312110331022233-0312332321120133-0202103103310300-2003012201300033-1031131320330213"></a>

### Direct properties for `site_subnet_params.subnet_dhcp_server_params`

- [dhcp_networks](resources--subnet--reference--group-001.md#canonical-1212322310001120-0113213123323222-3120021021013102-1202231121203200-0221330101002121-3330223233322031-0210133223212210-3100033011133231): complete subsection reference.

<a id="canonical-1212322310001120-0113213123323222-3120021021013102-1202231121203200-0221330101002121-3330223233322031-0210133223212210-3100033011133231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-0311002121020233-0121010300123320-3332310122013110-1001032303211011-3120300131202223-1122230330101010-1133333103020011-0333012002330013)
- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-3222100203231312-0002211232302021-3221300102032233-3133331312322010-0103123030010313-2202020001001100-2330211233023033-3003221313122203)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="canonical-2233311102012112-1232202213030303-0213221103231213-3311002302322100-2321033200120201-1300222322201031-0003030232130111-1002220122012233"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
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

<a id="canonical-0231213001300223-0132303100130223-0220302130022231-1330030330331002-0002100222013333-3020233200123013-1021122220133212-1331123300220203"></a>

### Direct properties for `site_subnet_params.subnet_dhcp_server_params.dhcp_networks`

<a id="canonical-1010230332011311-0022312102110132-2200321110310123-3122030112033231-3313010231100121-0000213302012100-3003322110302320-1212221002103230"></a>

#### `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

Exclusive with \[\] Network prefix for subnet.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2320222121331132-1330210203133001-3132321032323002-2221201331031332-2113013120310311-0231130220332233-1213301033300322-3332123321212011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- timeouts

<a id="canonical-3013201002022112-2213122313310210-1212100223031121-1132300101303220-1330332033201331-1310233001320002-0101303311310222-3221030133333221"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122321300102120-0122222213120323-2203003322333211-2030032231300200-2123030330131222-0313022223301203-1010203303012210-1102112023230213"></a>

### Direct properties for `timeouts`

<a id="canonical-0330323133103210-3111033011233010-0003132021212220-1121203111131302-2132310101021021-1102020203130033-1220033303301123-3203211310313020"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0213020113200302-2211011110233201-1133303000210320-1333222110310223-1320231010331030-3220120211123132-1231132031211122-0330112032313312"></a>

<a id="canonical-2132121000213201-1213131120211120-0031201311303232-3212322012020322-0101210011230220-3330110130013300-0233122331211220-3012121202103100"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1323201233131213-0123323130323020-2123303031312120-3322323021201120-0122110220230022-2212210012220111-2332012131312302-2210030332222121"></a>

<a id="canonical-0300011311232332-3221112313210232-3013131130322200-2131322221201320-2010233322331121-3200213330231231-0330321103313101-1223201113231312"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1200112131021230-3220231320212211-2301213101321302-3202220030020003-1212301020121110-1220131101133023-1300030323203331-3123120311022210"></a>

<a id="canonical-2031232013321033-2322121333131213-1310213113030320-0312322101113022-2333323033310231-3030330322002113-3201003232012323-1223310031310233"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
