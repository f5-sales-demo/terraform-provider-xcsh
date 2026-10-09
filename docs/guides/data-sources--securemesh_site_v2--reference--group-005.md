---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2201132232311322-2012210120031130-3130332323323102-3001201320111020-3020230320101020-0030133332132223-2212311232311013-2202012102131131"></a>

## `azure.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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

<a id="canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- baremetal

<a id="canonical-0212333130112232-3233322121301030-2112110113200202-3301102111213022-1120302223130200-1323021021133202-2301222012010332-1102111112023311"></a>

Type: `"single"`. Computed.

Baremetal Provider Type. Baremetal Provider Type.

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

<a id="canonical-0303232112112323-0333322032110313-2211300101103010-1010113110020202-0332013332210300-3211003003322030-2010103133011131-0223200310311322"></a>

### Direct properties for `baremetal`

- [not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133): complete subsection reference.

<a id="canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- baremetal.not_managed

<a id="canonical-2121222231323121-2333122033032121-0133030030030013-3323332022200102-2300002332000303-1202312020211331-3023332103332322-1032013102310201"></a>

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

<a id="canonical-1210030331201211-2131330002332101-2333031201121132-2101112223103320-3223301232102110-2331013201320331-1321002022112133-3002300230310231"></a>

### Direct properties for `baremetal.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120): complete subsection reference.

<a id="canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- baremetal.not_managed.node_list

<a id="canonical-1033033000331132-1322010230033322-3301332331211131-3101302022210132-1103200323223232-1033020120332123-1102121001012112-3120101122312212"></a>

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

<a id="canonical-0033012303231101-0000212032102212-1200013321301132-2010213320022032-2122022011023031-1221033120233010-2201100201011033-1222102210031221"></a>

### Direct properties for `baremetal.not_managed.node_list`

<a id="canonical-3232203311331011-2030300211202113-1023202301333310-0310012220312021-2122121023020301-3133111232203311-3001020113231220-0101320032311333"></a>

#### `baremetal.not_managed.node_list.hostname` property

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102): complete subsection reference.

<a id="canonical-1332303313232013-2301233010313101-3112321211321011-0112222111312232-2120131103302000-3330130201222020-1133222233010210-0233230120332231"></a>

<a id="canonical-3100300121210301-2032130003103200-2112013021202010-1021312303333302-3233130010133131-0030122322312012-1230321233302113-0331103302332112"></a>

#### `baremetal.not_managed.node_list.public_ip` property

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

<a id="canonical-2320013101330320-0130123013123213-3320132101132303-0232232013021312-0322132333000323-0113033213313033-0212303012003121-2313030020331230"></a>

<a id="canonical-2333031012102200-0300333230031132-3300220221322022-1030213010221223-3303321220031223-2020012012102222-3321010133223331-2310100313132222"></a>

#### `baremetal.not_managed.node_list.type` property

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

<a id="canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- baremetal.not_managed.node_list.interface_list

<a id="canonical-0331001132103200-1022103302123031-1122311003022011-3300032322110332-0321232132133220-0032002311012122-2113030301133232-2333132333302121"></a>

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

<a id="canonical-2002220311013321-2201102030110010-0120203000222223-3021311230103130-0222331023233122-2102330010223231-3233122102330323-0200210032100200"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2111333001001001-0112233213222031-2102031002303002-1221100232210020-3222220222000233-3123203022013003-0210320110003120-3021101320132223): complete subsection reference.

<a id="canonical-2033000023332330-1111033321121202-2013302102321311-3330003021113010-0222013211133311-3021023321301113-0010110233122223-2320211121102300"></a>

<a id="canonical-1233210112333302-3103233322323202-3110103033110122-3233120332320303-1310131203311022-2131203031113233-0333203310201010-1113221331230223"></a>

#### `baremetal.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2133211102030112-2321300132001113-0332010023002111-2201202013302201-3331301030110300-2322222220012322-3001120101023112-1020310220302021): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1201331302232102-0231221302220033-0210323223010322-3000002130231132-1132032331322033-3003330113013012-3012100201120311-2032010311231020): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133): complete subsection reference.

<a id="canonical-2001131020200220-2002320323120033-0230303130310003-2023001331130122-0323113122231013-0323121032232323-0230213013121200-2330231003202200"></a>

<a id="canonical-1323210211020131-3330012222033120-3202323321011201-2213312112312011-1000233302002323-3101131133200023-0000313213321003-1112003111022110"></a>

#### `baremetal.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0032211203110313-1300231322222302-2000001102133230-0332123130200111-0001001030301301-2221302102111201-1031122133332013-0100112021321323"></a>

<a id="canonical-1010223231301310-2332220133303030-1230020221333220-0013213333212001-0310132013200000-2132221303023012-2231111332223202-0002212332001312"></a>

#### `baremetal.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3011331302022011-3000113310100233-0302130301231100-0113122333321102-3032033302000131-3322331300320300-2022101221102210-3110103012301211"></a>

<a id="canonical-2301222123100032-0112120031012013-3211222221002122-0312223111033123-1222032200333323-1120130312031133-2210332310230201-3012120120322311"></a>

#### `baremetal.not_managed.node_list.interface_list.labels` property

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

- [monitor](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2030130213100000-3232033031002323-0013002100013323-1211123323030313-1300220132321001-2201110231332002-1311121300012331-2302131200123103): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3320130321111221-1110112030313131-3310001033213133-2120000321310203-0102312030323121-0012212130320111-0221233013213101-1033231010231120): complete subsection reference.

<a id="canonical-1010012030110031-0213232020220332-0102001031303302-0000202013102021-2221001121000212-0122113112200100-2003130001113303-0000213013303100"></a>

<a id="canonical-0102113101033321-1321022322323100-1012303232120311-2223131113032311-1311131212113131-0033331322210202-3012002132002122-2312133103222023"></a>

#### `baremetal.not_managed.node_list.interface_list.mtu` property

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

<a id="canonical-1113321010102031-1001033113033113-3313212101002210-0110033032302321-0311230013030302-0103113203220013-1033331212133330-3202310100101311"></a>

<a id="canonical-2230323001011102-1013112222310232-3002123001303320-1301030303230201-1000303313320230-3310020303110331-1201323020021230-2011031102123002"></a>

#### `baremetal.not_managed.node_list.interface_list.name` property

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

- [network_option](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3012202230012330-2102202222013211-0000311222021122-2210313230023033-0311321220210013-0002323022003021-0022212133121103-0320312301110103): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0233211023221012-2110300021113311-2330020122313122-1223310323023133-2222313002330101-0032111031220301-2222132030330033-2220332332023131): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2322323312201230-1200231313313122-1122002002321231-1202000302022022-3031013310203012-0221202303003000-1120101113222030-1222130032321013): complete subsection reference.

<a id="canonical-2231301300211013-3032132320323321-0211003101232323-0031200130233110-0122201311033211-0032120100210032-1323101313201223-3023113010110223"></a>

<a id="canonical-3020101120103021-3130022010122223-3201131311320022-0310333310132322-0211310211013301-2130023312321322-3232332123103210-1030321203021002"></a>

#### `baremetal.not_managed.node_list.interface_list.priority` property

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0330300002200112-0231232111131012-2220031021102003-0321122133033011-2002132033012111-1232332332200323-0212311313220333-2230022112220322): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0010110332132313-2220123010121131-0123123122131001-1022130013233211-0133123033103130-3133023011230222-3213033323030313-3133020012312011): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3220100210122233-2232222101312223-0231022011003321-3200130303221330-0123231010311233-3332210312102123-2022120033112010-1332101001221333): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2333303002302030-0233133210011332-2301333332001111-3101012131133121-0010210120113200-2112032003123012-3313110002032311-2110111310031300): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0100332010022300-0003110001013331-1302332000010220-0320133313322201-2313103122022232-1233102002201012-2121231022011032-2032111233031310): complete subsection reference.

<a id="canonical-2111333001001001-0112233213222031-2102031002303002-1221100232210020-3222220222000233-3123203022013003-0210320110003120-3021101320132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- baremetal.not_managed.node_list.interface_list.bond_interface

<a id="canonical-0213222333000132-3232131133333030-0231221320120211-2210000313333103-2003031131013102-0320312310023311-3131021000022313-2333301330321311"></a>

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

<a id="canonical-2102323020133022-3333320032013120-3300133203111311-3321201310211313-0013032320232222-3310321002001212-3021203010003301-0322120102123330"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0310303101332333-1010000102122310-0030132101311221-2321212310320111-1033102213201131-1022321321123312-2030111011002111-0332030300220022): complete subsection reference.

<a id="canonical-0213230102133201-3001001011011033-1102320310001303-3303012111321131-0301313233011100-1221220301232213-0222100201102013-0030103120302131"></a>

<a id="canonical-1113022003211122-0331331223002022-3031311002102332-0002200021013302-1133133012023310-1002203301123230-3021200330330000-3201331230233121"></a>

#### `baremetal.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1003113222223112-0011002020233020-3330302313213120-0222033200300030-1101003223010233-0122231100321221-2110201321303311-2110120031233022): complete subsection reference.

<a id="canonical-3311103111313000-1121221322031123-0112023301322322-1123200131310230-1022201322023232-1000133131120312-0222331230331210-2310320001203031"></a>

<a id="canonical-0031333002013013-3230311303310132-1220320111121202-3332002212002000-0201023113310100-3232100111323320-1112111003330033-1203113231131220"></a>

#### `baremetal.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-3230301033021121-3001021030101311-0102100112332220-2113032302001012-0110022311210132-0203310330321011-2211213322123211-0012033012111320"></a>

<a id="canonical-2000000333002301-1222223013013223-3100112023123210-2022333122120111-3222020211213103-0213223103232000-2103332103231000-0133320023031201"></a>

#### `baremetal.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-2313020002032120-1212311122030033-0033221103323231-0001200003110222-1133323011012201-3130323130002221-1122322100012200-2122301330100221"></a>

<a id="canonical-2223213213222202-0100300213031201-2123320322023312-2330332222021312-3212122322323100-1201020113202120-3121333200332023-3303003312032000"></a>

#### `baremetal.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-0310303101332333-1010000102122310-0030132101311221-2321212310320111-1033102213201131-1022321321123312-2030111011002111-0332030300220022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2111333001001001-0112233213222031-2102031002303002-1221100232210020-3222220222000233-3123203022013003-0210320110003120-3021101320132223)
- baremetal.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1101202032013032-0032211321230022-3021230232030011-0100311001022102-1102020011030120-1211120231010312-2203102003000132-1202100122332103"></a>

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

<a id="canonical-1003113222223112-0011002020233020-3330302313213120-0222033200300030-1101003223010233-0122231100321221-2110201321303311-2110120031233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2111333001001001-0112233213222031-2102031002303002-1221100232210020-3222220222000233-3123203022013003-0210320110003120-3021101320132223)
- baremetal.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1133100103012101-2132223223222310-1131113210322330-1213221000203102-0301132232223011-1320332023120232-3221211122300313-2320020023013301"></a>

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

<a id="canonical-0103222202220201-3230122211303311-2020322012020020-1032031120003120-0302301321303301-3031320001330303-3213312121021101-3100211231213313"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-1312212201023131-2313133330111333-3322211200331033-0231220311002222-2103021320302310-1112233031011022-1001120211332032-1303323112032131"></a>

#### `baremetal.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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

<a id="canonical-2133211102030112-2321300132001113-0332010023002111-2201202013302201-3331301030110300-2322222220012322-3001120101023112-1020310220302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- baremetal.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0100032203010220-2013300333033310-0113033313303023-2111213013132231-2003332023010301-3332132031113331-0320012103302302-3221013121312320"></a>

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

<a id="canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- baremetal.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3232021113133212-3123223311010230-1322313002023331-3311122112103010-0133223133302323-2120010112200223-0311332033320201-0323011010012213"></a>

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

<a id="canonical-3330313221311011-1020011230320303-0021013303320313-1333020023212300-1023310200331111-1001010113320320-2103033310102230-1311222333233222"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3110211211023202-1121321221320012-1123111310202121-3121330100113231-0200210310301010-0211130220211011-0313010331322202-3011101222100301): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1200231121232003-3110210310111130-1321031331132000-0032101133103333-2232333233023100-2221332120133302-2201111102031103-0131020333031101): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213): complete subsection reference.

<a id="canonical-3033111222312010-0201332230030211-3003123232133130-3201012230231100-0221323132322220-0111331320000103-0312112002312312-2032030030203122"></a>

<a id="canonical-2033232022011211-2020311112302210-0003131011310330-2122113103120112-2012231211113123-3212310022233203-2203101233020222-3322031230113000"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1001300131133220-2323201323311233-3123123300310200-1303132130301323-1120312033021213-0322200232300212-3022311100202220-3222120301123332"></a>

<a id="canonical-2011020123213200-1123120033011021-1023133311022302-0123102022022331-2300310203212230-3313300031130203-1100002312301112-1302120123220020"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0130230321021100-1121321022211023-3311030001031122-2303033333330031-1202202031321231-2330121102133330-2312213120031201-1322210323121333): complete subsection reference.

<a id="canonical-3110211211023202-1121321221320012-1123111310202121-3121330100113231-0200210310301010-0211130220211011-0313010331322202-3011101222100301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3110122233302123-3102300003333002-3110322133122010-1123200010010131-2222110112000030-0011110112222130-1030211220230312-1233012212233203"></a>

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

<a id="canonical-1200231121232003-3110210310111130-1321031331132000-0032101133103333-2232333233023100-2221332120133302-2201111102031103-0131020333031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3122010130321010-3232120001122303-0022302232202012-3300211103103101-2320300221322331-3200111000023321-2323321201113003-2010331011323020"></a>

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

<a id="canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1202200211323112-1333322230303323-1103312333330220-3211222331313232-3101121023302131-3230232102123201-0232100323310102-1200101333013101"></a>

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

<a id="canonical-2112131330322301-3211132230321232-0203332112330230-3312312112012003-2310030120010231-2332022310331020-1320021331221002-3023111203212213"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2100312120130100-2213031313332212-0010212103310112-2101120210212102-3232101030233310-1200001003002123-3203101112003123-0112012223102301"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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

<a id="canonical-0132033233032013-2103030020222121-0121303301220223-3200213131203202-1213323121330210-1313331213122303-2303201221032003-1303023131232133"></a>

<a id="canonical-0202033222321000-2311221023213123-1310012232302211-1221032102322120-2330021310103220-0011102003003031-0212010021102210-2211220333123000"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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

- [first_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132211221011332-2223001203013030-1032232301223132-2110001220022102-2121220232123323-3010012311031133-0123101212000213-2213333232122333): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2210211310301310-0102033023312331-3002020121131223-2133012132331331-1203103203323101-3101100233030013-0100322032100101-3301303021313001): complete subsection reference.

<a id="canonical-2321102133313101-1122131223210131-3101333331000133-1201020211021233-0330132303323121-2023220220020102-1030032100330020-0232030110130300"></a>

<a id="canonical-1113103310030222-2200330020100001-2030233103121230-2001120322211230-2301131332221200-2231112232331103-0013311012233120-3031032131202112"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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

<a id="canonical-2200313302020033-1222020301031013-1101130230311221-1300013232330132-1020323320330132-1032233320331311-1333331332023032-3013112301020320"></a>

<a id="canonical-3330330101003302-1230320312111301-0300212202101220-1213132200300222-2031232030103310-2301003322030032-1211100103123130-2103332233203300"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3111030310202231-1102030331113303-1002313122020132-2000223301022131-3233132020202320-2120012223022000-2303113313131032-3230311322201033): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3320011303321110-0022011002032200-1110313302221331-2121013023303223-3231312202200021-0111122330311220-2000330211322121-0033032113022221): complete subsection reference.

<a id="canonical-2132211221011332-2223001203013030-1032232301223132-2110001220022102-2121220232123323-3010012311031133-0123101212000213-2213333232122333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2102231033300100-3122313112010223-2012030212023120-2121013112333120-3301323212201333-2212303020230232-3030301311001021-0133031102220211"></a>

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

<a id="canonical-2210211310301310-0102033023312331-3002020121131223-2133012132331331-1203103203323101-3101100233030013-0100322032100101-3301303021313001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1022013300232330-0312212202130132-1321132222023031-3111313223220233-3322033023201130-3301002123213002-1220110020312030-3121222211213121"></a>

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

<a id="canonical-3111030310202231-1102030331113303-1002313122020132-2000223301022131-3233132020202320-2120012223022000-2303113313131032-3230311322201033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-1200000232021003-0102122213223231-1222133232011313-3203331333310222-3331230021322313-1320130011223322-2332110333031120-3012003223310211"></a>

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

<a id="canonical-1033201000300103-2131120310221312-3332002311222303-1002023021023012-2001300313101300-3233330121110102-0120331223123120-1121330003220221"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3232312112020121-0110322233120302-2013113233230333-1212331030330131-2223123023210312-0023232231112210-0331321233023213-0121003020010130"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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

<a id="canonical-2333321300133130-2201202133320032-3010002231022301-0102202011110103-1233121233012120-3301331011120321-1130102130302020-1121123031030211"></a>

<a id="canonical-3122022113302332-0221303121010220-1011013022211130-3023022123013223-3223030203311311-1232310300331022-3320001200231100-1230131231133020"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-3021212132022222-3313133200220312-0132131130222101-1210230103312230-2230120103200332-2202222012011103-1023020210330001-3131210021232222"></a>

<a id="canonical-2330133201220322-3113000112202213-0112233310232011-0030132221312321-0323211100310000-2210120213200123-3123301223101000-3210323012230131"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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

<a id="canonical-3320011303321110-0022011002032200-1110313302221331-2121013023303223-3231312202200021-0111122330311220-2000330211322121-0033032113022221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1113112323012300-2100323210010331-0230213021023033-2131313231030122-2303213300211030-3311133210210213-2313120322230003-0102032221122213)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2133022131012033-1222331230230320-1130122201313032-2212001303311313-3132101200320100-1103331012321013-3022311102303320-2100333101000000"></a>

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

<a id="canonical-0130230321021100-1121321022211023-3311030001031122-2303033333330031-1202202031321231-2330121102133330-2312213120031201-1322210323121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3032003333312220-1212223021301202-3210032033330331-3222120100021123-0032331310003302-3231122113133031-1013120210330110-2101032112221113)
- baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0202201230123231-3032103221123011-1320010111121131-3131312220032001-0030100031301322-1022233030032323-3022301032132303-1103301211230101"></a>

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

<a id="canonical-2201213123212113-2323012123330200-0310200122321113-0022013220212132-1121301010122231-3023133301200203-3013023131322123-2310302300113122"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-2113230123131202-2113023110113320-2222330330322003-3322220031023030-0110102013223012-2011122230223132-1002121020011310-3202011212131003"></a>

#### `baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1201331302232102-0231221302220033-0210323223010322-3000002130231132-1132032331322033-3003330113013012-3012100201120311-2032010311231020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- baremetal.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0223203102212303-0322030113211133-0330112130331313-0332301322210021-0033112031200311-2300210202213200-1231202213233211-0110120132322211"></a>

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

<a id="canonical-2203332231120133-3132112211222202-3022330100330133-1100132021100030-3101103032122200-1322302303331331-3212110021131132-0233010031131300"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1300112122121030-2211000312302123-3202230123133031-1313103330311300-1120133100210011-0323131223331201-0011323131333001-2313120222020210"></a>

#### `baremetal.not_managed.node_list.interface_list.ethernet_interface.device` property

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

<a id="canonical-2212331231302120-0200311013101202-3230121323300130-2003111030322212-0202013023033222-1132211221123133-3302102132101220-2322122033001333"></a>

<a id="canonical-1203103120022232-3213112200310320-0010303001310130-2312032102010213-3031001101123232-0311133321232223-0030132100100302-3021133330013301"></a>

#### `baremetal.not_managed.node_list.interface_list.ethernet_interface.mac` property

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

<a id="canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-2320312223030220-3013013212122200-3012110102113130-2200130030212032-2223100222023200-1323132300131331-2121233221001000-2022133131300121"></a>

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

<a id="canonical-1020100210013033-3112220003131300-2223301110020113-1213201102112021-0123033303311002-3102020321311011-1302020313221022-1113302212203112"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3232231022023233-1031331231211132-1230320101111321-2322121211012331-3220033020121111-2022300102132203-2313033130321322-0232123202021000): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121): complete subsection reference.

<a id="canonical-3232231022023233-1031331231211132-1230320101111321-2322121211012331-3220033020121111-2022300102132203-2313033130321322-0232123202021000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0321202323132303-3003212202102321-2102223112330133-2032221203121100-1123111022320310-1221222100123021-2100233032221211-2331222222120201"></a>

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

<a id="canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2032123222303311-0202230132002033-2300113130212000-2020231200321220-1133230302001232-3200213210230212-0131102311311311-2233322132122130"></a>

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

<a id="canonical-3222321001230001-0320123231301210-0033011312120001-1101103023213303-0031333321133201-2310302101111333-2232103031010211-1301313111310222"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330): complete subsection reference.

<a id="canonical-2302200220113313-2032200132323022-2201020111021031-1003101311110022-2001113013332113-3232013122030223-0012213310001303-0300211213311001"></a>

<a id="canonical-2300300002101021-1103131223002021-2313303311132230-2131113313001103-3103000032122110-1211000022130300-1313322010002112-2131322130202313"></a>

#### `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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

- [stateful](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0011012300112323-2211231221031132-0103223031110303-0013231002100202-2121100330230013-2311203232210223-1010121220330031-1110201313033103): complete subsection reference.

<a id="canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3133332033320321-0010033030220101-2122000213131011-0222331012103123-2320102212321102-1032302332303032-2101212310133200-2103221333222313"></a>

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

<a id="canonical-0131312113231321-2003122000222302-1202313000020330-3313032212000213-1120020010200320-1322201011031303-0303203203322213-3232101212120102"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0102011133321003-3331210302132200-0123201003003213-2210121221232130-3323013122113223-0131331133200121-3110232133320113-0330301313203101): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1121332233232110-3321320121233320-2032003132233123-2202021102212022-1002030112233122-0212032110001013-2031020031023333-2100013201212131): complete subsection reference.

<a id="canonical-0102011133321003-3331210302132200-0123201003003213-2210121221232130-3323013122113223-0131331133200121-3110232133320113-0330301313203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2311030230203100-0010313103011100-1311012112301121-2230103232002011-2121311200111332-3213202323230123-1331121123333220-3233200213212133"></a>

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

<a id="canonical-2322021123021022-1012213112230102-3213012233030022-3033321221223111-0200221003203100-2031320103113302-3013323232303323-0220100300112302"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-2110010121132302-0101000232120033-1320233311012320-0002331022020333-1311203033233202-1213112311313231-2123030202033022-3103013313310033"></a>

#### `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-1121332233232110-3321320121233320-2032003132233123-2202021102212022-1002030112233122-0212032110001013-2031020031023333-2100013201212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2113223200001200-3201213323313003-2133200312332321-3210212220220323-3202123033030223-3220113101021333-0202232120100313-2002201303312213"></a>

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

<a id="canonical-3223133132203113-1330221021322102-2222133030313200-1130021233121110-1101010311020102-1233021220331111-2310302022103201-1223221020230320"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-1131330313122002-0312122303330212-2311232303003132-3120033012230101-2201330333010210-0330000220330311-3303232331103220-3103133002312332"></a>

#### `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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

- [first_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0111121201211112-1002022131033033-3020021201320223-1123320120130201-3123020313203123-2010113011111132-0002131210123231-2101330223202331): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0310201213023122-1223112120103212-1221223201300100-1100000200011200-2303302021333020-2322113202023201-2223002232113320-2203130023003100): complete subsection reference.

<a id="canonical-0111121201211112-1002022131033033-3020021201320223-1123320120130201-3123020313203123-2010113011111132-0002131210123231-2101330223202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1121332233232110-3321320121233320-2032003132233123-2202021102212022-1002030112233122-0212032110001013-2031020031023333-2100013201212131)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2223230321011223-1333112020121100-0310300233201323-0002111330001303-2200013000030300-3310032103220220-3001001311122301-2122001331233021"></a>

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

<a id="canonical-0310201213023122-1223112120103212-1221223201300100-1100000200011200-2303302021333020-2322113202023201-2223002232113320-2203130023003100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2132320222223133-3321033011101223-3302321113201202-3011202202002120-2320220323313333-1211121330231302-2033033120321233-3220133301111330)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1121332233232110-3321320121233320-2032003132233123-2202021102212022-1002030112233122-0212032110001013-2031020031023333-2100013201212131)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2332113333132022-0000202203312111-2122213132002323-1033100013012011-0020202333030011-3220001301130202-2231023333023122-3232302222213313"></a>

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

<a id="canonical-0011012300112323-2211231221031132-0103223031110303-0013231002100202-2121100330230013-2311203232210223-1010121220330031-1110201313033103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3232102012223122-2222012113310112-3021021032001113-0101311020111123-0110221012231323-3000001203002210-2202310201101010-3021130210113210"></a>

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

<a id="canonical-3133220323121230-3010022033212031-1003311301332301-2033103103031202-2322112021010103-1322030133310102-2020333201123223-2303230121113213"></a>

### Direct properties for `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3122103101010310-2223232322001313-2111013330221223-1013321120223011-3203222203320132-3010120233120221-2232212131323011-2122032331010311): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0233303311321102-1112021311121021-1002102012310331-2100103223300100-2331130222013022-3011001032200211-1330312200230121-2232001212112212): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0020030120313131-1121033032133002-0321301200102001-0322122122221110-1223332203220233-1133301313032103-0110312033333222-1101111320232103): complete subsection reference.

<a id="canonical-0123201021001121-2203120112331120-3202130103030232-1013111223021220-0331331333222031-3232123203302121-1000133122203131-1133213211111331"></a>

<a id="canonical-3301202130101322-2231330011122210-0032011130301113-2230323303030101-1022301331030201-0311300323223103-1220322111103000-1302222131212301"></a>

#### `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2101030003110223-3122331020132001-2122001333121022-0021333132221333-1300001030021301-0101323013232233-0200020332112311-2220220111020311): complete subsection reference.

<a id="canonical-3122103101010310-2223232322001313-2111013330221223-1013321120223011-3203222203320132-3010120233120221-2232212131323011-2122032331010311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3112130220103321-0222301102012110-3130100021102020-3131100011331012-1033223021300130-2203021223122121-1203022230220231-1222321223133133)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3203030210103330-2221020320313220-3333010131002320-2122021233123301-2212123113113323-0021210002002332-3100033101203003-2011213121111120)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2203000220210131-2022322210211232-3231132120223110-3001211000032111-1320010213302012-2012101233212111-1120100233231202-1230332003113102)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-005.md#canonical-3231132020000200-0220133211000333-2010131003033211-0221031111132110-2112103322012312-2111000310302223-2111201323031013-3131122022123133)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-005.md#canonical-2220010321311231-1212033103130301-0210013123221031-2322102311113331-1112012010122012-2032010332313033-0002310030200130-3201202333102121)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-005.md#canonical-0011012300112323-2211231221031132-0103223031110303-0013231002100202-2121100330230013-2311203232210223-1010121220330031-1110201313033103)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3122200330311223-1003313330103030-1223013212122313-1230323122132111-1122031020002102-3102232030001122-0202013311121032-0300120103011230"></a>

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
