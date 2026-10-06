---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2203133003210120-2231112330111110-2030222311211302-1102201321301033-0331013123002130-3332120021121130-3212320331331020-2101001020112320"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1201130233120220-3200313302313223-1232322332203121-0013221321202000-0100003220123222-2231103323221103-1031201311221212-1223202212122323"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list`

- [interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3320203210031330-3201020130111221-1210003000320233-3220230121132000-1230112313023102-2223033221002103-3011331113213111-0113010300113013): complete subsection reference.

<a id="canonical-1222010011300012-2022232310022023-0121300110011310-0310000032312333-3112023301023311-3310211020312021-1021330313232200-2330233112233103"></a>

<a id="canonical-2333320230002122-2213110320030311-2330233300022330-3022232113123323-3202102212123331-3011103330223020-3011211012001233-1320022303230010"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Computed.

Node. Node name on this site.

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

<a id="canonical-3320203210031330-3201020130111221-1210003000320233-3220230121132000-1230112313023102-2223033221002103-3011331113213111-0113010300113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2003320121033102-2031122111320223-0022232112102003-3013133031311101-3220220022203002-1023213120011103-1110321122300201-0132322321302200"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

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

<a id="canonical-1202221021200000-2013031230030121-3020310321332223-3003112132100001-1120220102203022-0312320333103001-2103330311030013-0220032213300130"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-1112310021231002-3330011330202323-3310323012021303-3122313213120033-2320303213322021-1133221112331313-0332100013033301-0023332023131122"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0222023130032112-2031033320123123-0323213130133110-1321333332033012-2133012033203313-2000131331330202-0203102130123122-2032001132112121"></a>

<a id="canonical-3331231113103320-1220003122110321-3331331023222312-1023112100311333-1103101022333122-2301220200313222-2100022330123201-1111233200300231"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0231223020031033-3223032220031031-2312120030032121-1022131210030120-1220000202010202-1301220230020312-3032132223113003-2133321033311023"></a>

<a id="canonical-0003030313021010-0003310030321313-3000231130031221-2033222101001320-2030113220200322-1113113320022231-1130220020233231-2130120300313101"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1312013100113120-0032112331103330-3213222100202313-2201022002112222-0102112313133110-2030021132222132-0330321312331121-3120333221130021"></a>

<a id="canonical-3223202302331213-3113322112333102-3021102120022301-2200010113232210-2222323120122123-0220211203311302-2031311002300232-1132012333101200"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1102312102213333-1002310122130312-0223211321110233-2210302313202222-2213132101030122-1330121123012000-3032321003121021-1032312132230220"></a>

<a id="canonical-2123102310002030-2113333300203202-3020120013130233-1232122121131130-0022220322130110-3100022211100330-0322323222311113-2001222000120012"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- local_vrf.slo_config

<a id="canonical-1013302003210320-1010321312202012-1113002230031012-0313322103110330-0130330110021332-0323032233130022-3103102233330120-2102022322220331"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

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

<a id="canonical-2010230311121000-1322011102010213-2110100303211223-0311103122212132-0102120322033331-0103103023130313-2102221123103133-1300221122230210"></a>

### Direct properties for `local_vrf.slo_config`

<a id="canonical-0003123121022021-1013130021332101-3031320132131101-3310013302032120-3000013031212022-2100203033213000-1222322300100011-0022322221221010"></a>

#### `local_vrf.slo_config.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-2100000031210301-3230203023002212-2231110221100212-0212131323310321-3003023031311232-1312133000102332-1110111333303201-2001121103321111"></a>

<a id="canonical-0202200323303030-2321233022110302-3330012300010230-1302110013222110-0011303112323211-3012121002231211-1231123033301122-1100123323010023"></a>

#### `local_vrf.slo_config.nameserver` property

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2132113300200203-0000123232033102-3232001122302031-2003222000012132-3003101002012012-0302212131100313-3321333020133012-1300133231211202): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1121322113210100-2321101323310121-2303211003010113-2020132223121213-3103101202011320-1012033032022001-3233330203020010-0220002333022120): complete subsection reference.

<a id="canonical-3121312133022332-0101220202232212-0302013101310122-0201222131010013-2133000231332121-3233123223111223-0312213222231333-1323123302111311"></a>

<a id="canonical-2331111223311323-2002031331302302-1102200313112300-2301001322221021-0200122320231231-2321220130213130-2111133203230001-3221133123013110"></a>

#### `local_vrf.slo_config.secondary_nameserver` property

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222): complete subsection reference.

<a id="canonical-3011120203301231-3323122103212010-0330131333212013-0232013330121133-0020200200013021-1213021202230313-3222101000010120-3033321121120310"></a>

<a id="canonical-2111001023000032-3131330102232223-0323332202011300-2030110002031222-2203123132100101-2210023230032023-3323210220111012-3132233203123111"></a>

#### `local_vrf.slo_config.vip` property

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-2132113300200203-0000123232033102-3232001122302031-2003222000012132-3003101002012012-0302212131100313-3321333020133012-1300133231211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.no_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.no_static_routes

<a id="canonical-0002212233003002-0322113100013000-3302133202131001-2302111033033130-2000013323302313-3031130110223001-1302333023201033-3003010011013200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-1121322113210100-2321101323310121-2303211003010113-2020132223121213-3103101202011320-1012033032022001-3233330203020010-0220002333022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.no_v6_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-3113133221211201-1030232331033303-0323111112303330-0102331120232112-2212003230131113-1102200300123002-1003312120131222-1001311300132103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.static_routes

<a id="canonical-1211031221132030-3133211012021313-1213200303110122-0102031232111012-3210020221002201-2231011023221302-0030331020230032-1330113131322112"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

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

<a id="canonical-1131331132012320-1030033130113010-2301001311011303-2113003332223133-3333311202202212-2131110301102202-0002322021110312-2022313022100023"></a>

### Direct properties for `local_vrf.slo_config.static_routes`

- [static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013): complete subsection reference.

<a id="canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-2322111233332011-2221112233033231-0000010132122210-2222010201011333-0111331220032211-3010030020221120-1213011321211302-2223200202032112"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

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

<a id="canonical-0231120003000222-1322321311101001-1100032030202033-3121323121123113-2031011223230331-2211021212112211-3202200023122222-3120100312233303"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes`

<a id="canonical-2033020313220012-0120010320332210-3321201232131230-1210113301300232-3322321000231312-0102312322303033-2112003313101100-1300330312200213"></a>

#### `local_vrf.slo_config.static_routes.static_routes.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121): complete subsection reference.

<a id="canonical-1212213113121112-0132323130032223-2321232113330030-3212300300212220-2103000132200311-0233321231331123-3102330211111232-0312323231322122"></a>

<a id="canonical-0023001333312323-0132030101321301-1312012200003031-0332102231302132-3113201033033121-1030010330033010-2103031100200133-3230211112320201"></a>

#### `local_vrf.slo_config.static_routes.static_routes.ip_address` property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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

<a id="canonical-1131131131121101-2310103200012303-1212132133031311-2223211113100021-2203011201302000-0310310133032133-3010023023233311-0003130023312030"></a>

<a id="canonical-2122300302331332-3132121322103132-3333320100211331-3212202033131021-2000121123021120-1032223120303013-2021311301021303-0323330000030221"></a>

#### `local_vrf.slo_config.static_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001): complete subsection reference.

<a id="canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-3220330001301102-0211021232103133-1220312322120331-1132120200002130-3310210310221021-1221131131120012-3032103110011130-3023012230320200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-2110101322321323-2120022210221122-1023312033200032-1022112331200222-1132103121322230-0010032302212112-0302120011020300-1120313303300022"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3000222130110313-0230210113233123-1131020331002331-3003012103003310-2302221111320113-3002100022020030-0201122212122020-2123033221232002"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface`

- [list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321): complete subsection reference.

<a id="canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-0131011133311122-3303033023132210-3312103233201200-1211120210011323-2300123130213031-2123001333300012-3302103012121122-1201123031123031"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2002123001022201-3222221213311202-3021200213020131-1213120003101221-1123320322330113-0301323311311123-0233330021112021-1301102023133331"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface.list`

- [interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3222110233202020-1302103100230010-1333230331311102-2302021032112130-0101113212111130-1232010111110101-1320033022010320-3300022031111302): complete subsection reference.

<a id="canonical-1132030223300113-2013203113301002-0331033300112132-3021032201203200-0013302231023323-1231233220202130-3133121203232001-3301013232221322"></a>

<a id="canonical-2020120210200320-3021333210211202-3210212101112201-0033022312203030-0320012013100320-2333100233332010-0021123310131310-1330121022123312"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Computed.

Node. Node name on this site.

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

<a id="canonical-3222110233202020-1302103100230010-1333230331311102-2302021032112130-0101113212111130-1232010111110101-1320033022010320-3300022031111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3102200011002202-0023300013011023-0320001131323010-2222022321331213-3213021322013032-3332112221222122-0011032300321001-1211122332032322"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

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

<a id="canonical-3223012230000103-3121101011201320-3300033313130110-0201131100000101-1101132021332233-3203123122302232-0033213032001231-3231231102320003"></a>

### Direct properties for `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface`

<a id="canonical-3001213332310221-3003130233233210-2211033321302213-0112322130103003-1121202013212310-1333223120301021-1201210003013012-0103320001200022"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2022302232111132-0203120230303202-0230030123332322-3312312211132031-2021323101310020-1122323203012113-1103010103110321-0023330203130200"></a>

<a id="canonical-0131012003311301-1313331300133332-3010211222133231-3202012203230233-3223230102301020-0213112310311112-2110031211320020-1212001133201231"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3302323311302113-1230213130112132-2102313030313111-0312123320333023-0201013330232333-0302003223210221-1012200132220132-1332101223210212"></a>

<a id="canonical-0123323220221312-2132023122201003-2332213113320112-0022130212203013-3101303120330111-1321310023310323-0100231201103001-0121320311332232"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2021233303231212-3100332321212311-1200032321230300-3103010023311001-2111122013312110-1111130003121031-3202313202333021-3032203330300100"></a>

<a id="canonical-2132330111013332-3030113311333003-0232211001002122-2313222203223221-0112203102113013-2232220110312013-0012123331100001-1132212001102330"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1230323310020303-3313001210120131-0303203330220123-0230002000000021-1012321203030022-1101103303120022-0022201222223101-3201000330212120"></a>

<a id="canonical-0033002323022111-3123332123332220-2322210212130313-1102030331210221-1300223303130020-0102020230031223-2133233121122320-3201310011001303"></a>

#### `local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-3021000113330001-3321113322001211-1011231322212022-3113121001321003-0123212233030033-3122011301021111-1020001210322133-3210101231221100"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Additional upstream details:

List of IPv6 static routes.

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

<a id="canonical-1112323112210032-2222331332332011-3203111303010030-0112102222023011-0100232133130033-3201000003001321-2232230230312232-2230333031003323"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes`

- [static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212): complete subsection reference.

<a id="canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-3102201303300212-3120103322112201-1220112002102133-0330120333020022-3213300110031010-2200132301012001-1112022013032032-3121310322022331"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

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

<a id="canonical-0100203210113320-0210013300032101-0011022003103112-0032021112033200-2333122000333133-1110333101310032-2231012011222133-0310211233021220"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes`

<a id="canonical-0323031301222001-3333010301212322-1322012210102103-3131301132212331-3103020030321223-2321212213210030-0230000110331203-2320002330201013"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3301233302320310-3121032100211223-2201033200120110-3133303030123212-1302320331022131-0022233320230203-1112001202211222-1111312210323023): complete subsection reference.

<a id="canonical-3113333332233000-2303202102100130-1320021003201321-0220102112303110-1020321032023011-0123320330103203-2123222310222210-0002112203122323"></a>

<a id="canonical-0310231000121131-1003013221101303-0023202210021003-0002001331213103-3131013113331300-1021033311302113-3013111231310211-3101303100020113"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.ip_address` property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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

<a id="canonical-2122312133132133-3102013210022201-3200120210103013-0120031310102202-3323033011003020-3001211103032232-1001113131033331-0102212301011303"></a>

<a id="canonical-3211222112222021-1131203003132322-1022230200300013-0003023112332013-3201022010001120-0120031033130312-0323203223033312-0310202031223212"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102): complete subsection reference.

<a id="canonical-3301233302320310-3121032100211223-2201033200120110-3133303030123212-1302320331022131-0022233320230203-1112001202211222-1111312210323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-2001222310330030-1002020310301230-3121311010133001-2210211322303300-0301011113100132-0321212330231222-3230110012323020-2213323211231103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-1020022301232312-1331133130331211-0101110100032310-0100212303210122-2010113012022133-3222130203010231-2220331102113131-1012330202200131"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3233120203000033-0322023220203013-3003132332321233-3203312113023112-0322331232330312-3131012122232232-0103123032013013-1020223002101321"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface`

- [list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231): complete subsection reference.

<a id="canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2231002122122320-0202332302121221-1301123013101130-1111103130110021-0201130132030331-0230103120212120-1221110012021020-2022222211031321"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2302320002001031-2213302201201102-3132100111111021-2300300321123211-0013103123221102-2130130130033000-2002113313310032-0110030302111213"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list`

- [interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2130210023031001-3320220020021231-0221232120113120-0322133310310110-1332023110311001-3212033233102033-2310202323201210-1000210003232311): complete subsection reference.

<a id="canonical-3130002332320222-2210003103212003-0213210231233321-2221323001323233-0333101203221233-2330113301321133-3320200100303210-1301110120331211"></a>

<a id="canonical-0101303220202212-2030131120011131-0103312132210022-0321002031333130-1322021120320121-2021221120312010-0021201022030323-2012303331032022"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Computed.

Node. Node name on this site.

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

<a id="canonical-2130210023031001-3320220020021231-0221232120113120-0322133310310110-1332023110311001-3212033233102033-2310202323201210-1000210003232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-0333000012102302-1012302203223013-2202230332132231-0221300022112030-3322311202221210-1012000210103210-2112310321313231-3103000102031013"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

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

<a id="canonical-2120101300123211-1332003111123003-1120311010031331-2132302221031303-3333022033223133-3120013021102111-2020312323200301-2320102333002112"></a>

### Direct properties for `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-0221330031230002-0023220303213203-2202300023332232-1233112133232223-1312132033332222-3021301331223321-0321131101313331-1203033100000212"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-1032120000331223-1012020322220111-0010101302031012-0000110230201102-3003203200102220-2022102301010313-0210123012100223-3122111303231232"></a>

<a id="canonical-2130211032111221-1220222311332211-2132203313223133-2130001133311213-0002113101130230-1303231331320022-0323011112201221-1231113310203100"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2220023030220013-0220001200221232-2321323221131103-1303321000230011-0312201201200220-1213333103032011-3023031000233330-2133303221110000"></a>

<a id="canonical-0102121033300122-1011213200020321-1101020331002121-3301033032101031-0231011310122102-2231000323303313-1222202301030120-0003013213233301"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2220211021330123-2012113002021322-1331303203132210-1133011201113222-1023223320311033-0300331000303100-3032113303002110-2222223200031113"></a>

<a id="canonical-0301113212202310-3130233002102032-3312132303313023-2003220321212330-3200303222023033-2310031310332111-2023103031230002-3203101003033313"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3122311222302011-1321030011033013-1233003102311303-2100221113201232-2012212033100131-2000012021211223-1200113000203010-3102210030331320"></a>

<a id="canonical-2301333310303312-2012203021100120-0210130121210002-3012013000321011-1313312313313311-1001032303010122-1210100310211110-3200201100121101"></a>

#### `local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- log_receiver_with_net

<a id="canonical-0301211103003021-0313300321312033-2120323001233212-2030032032023221-0303303100300233-2222013302331212-1311132303231023-1123201110102230"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0301211103003021-0313300321312033-2120323001233212-2030032032023221-0303303100300233-2222013302331212-1311132303231023-1123201110102230)
- [logs_streaming_disabled](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0113001212320132-1202200001010210-0110120331313212-3133300233102210-3213232231120120-2022103212230103-2202212230013003-0132130233031300)

Select alternatives according to the provider validators above.

<a id="canonical-1032201332133312-1022021330032222-0132300011013010-1021011013102130-2103322302031031-0200103101102320-1122310231213302-1033011023101132"></a>

### Direct properties for `log_receiver_with_net`

- [log_receiver](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0211033200202122-1123101121020121-3330110012222130-0013311202020033-0110213122332301-2301011133233130-0220202020200001-2120111032000123): complete subsection reference.

- [use_management_network](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2101311113003120-2012100221130321-1033020122303322-3303312323312102-1022312221231230-2331301302230323-2100211112320301-2322112033300303): complete subsection reference.

- [use_slo_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020330001303210-1313031003103231-2011211201321223-3320330302021110-1100112022000203-2320231020030103-2023112321021011-1022131212022232): complete subsection reference.

<a id="canonical-0211033200202122-1123101121020121-3330110012222130-0013311202020033-0110213122332301-2301011133233130-0220202020200001-2120111032000123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.log_receiver` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.log_receiver

<a id="canonical-1211223232303330-1012111302230330-0323030310202212-3000130203020131-3312303223322032-1133300011203313-0232232013011211-0302230313320332"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3003103302011233-1223221121132123-2032130301233111-3030120020320100-1200331203212322-2020220222112133-3210021110132203-1011120310112300"></a>

### Direct properties for `log_receiver_with_net.log_receiver`

<a id="canonical-2031202203012301-1223210111303231-2230131032301211-1130212032210313-2011300233001300-1321213103011232-3222012300022312-0313100021202202"></a>

#### `log_receiver_with_net.log_receiver.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1332210210032121-3013211223113130-3132310012122113-3311110013123312-0133123210132202-0313212203122322-1021102021333132-3323202120110020"></a>

<a id="canonical-2101223233023200-3200021233202122-3000222223210231-1032111211123211-3322312100112012-3332311222122230-2313101100102112-2031101300221331"></a>

#### `log_receiver_with_net.log_receiver.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2323310332031313-1031130010122301-1133113112322201-0322233311020221-0212033030003122-1113211120002032-1120112102112130-0333322102103002"></a>

<a id="canonical-0202221131002022-2300010211311110-0011233300323210-2010011132213110-2122323002210233-2303030020033133-2301320223201111-1230131021020132"></a>

#### `log_receiver_with_net.log_receiver.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2101311113003120-2012100221130321-1033020122303322-3303312323312102-1022312221231230-2331301302230323-2100211112320301-2322112033300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.use_management_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.use_management_network

<a id="canonical-2232233012321222-2213200222012010-3010230110222303-0213031230130033-2301212122113110-0322302333012233-1331103132233221-0332121100113003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use management network.

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

<a id="canonical-3020330001303210-1313031003103231-2011211201321223-3320330302021110-1100112022000203-2320231020030103-2023112321021011-1022131212022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver_with_net.use_slo_sli` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.use_slo_sli

<a id="canonical-2202032112213110-3032320112313233-2300120113123213-0200010333011013-0103100002012021-2102033230230220-2100333311331331-1123113301010100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use slo sli.

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

<a id="canonical-3312333002021230-2313121320113132-0202121101132333-1020313233323133-1212020320020130-0320131000002003-3213221321030320-2011010311010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- logs_streaming_disabled

<a id="canonical-0113001212320132-1202200001010210-0110120331313212-3133300233102210-3213232231120120-2022103212230103-2202212230013003-0132130233031300"></a>

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

<a id="canonical-0233333211333233-2312020110021301-2123021233033012-3010102332313210-3231313033300120-1121330230203121-1210102010020212-3302330001121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_forward_proxy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_forward_proxy

<a id="canonical-3131331111200213-1222033103123101-0313132300212121-0002301220000031-2200101113220100-2131220000133320-0323310300120230-2032200330321331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-2300302332233003-0012321023230001-0002203123133321-3301231303333030-2131311231231300-3321021131122301-1111003011023013-3302002202323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_network_policy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_network_policy

<a id="canonical-0020220320210233-1230003021333300-3003201101220303-1103313102222220-0120311331012121-0322133310300111-0322120121200032-2332221121003102"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-3320113330113312-0230001301120322-3002203321220201-2321010231000321-1003133122001312-0312021000200203-3112012033001130-1333100212331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_proxy_bypass` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_proxy_bypass

<a id="canonical-3233031100200303-2012031220300110-1200001122032311-3300310213323033-0220300312010320-1001231012313000-0301200031232320-1231203313021022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no proxy bypass.

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

<a id="canonical-1323111133021223-1232233001220010-1033101302021013-2102121030310221-2202010203020030-2300133202310213-2120230100101000-2011231333110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_s2s_connectivity_sli` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_s2s_connectivity_sli

<a id="canonical-0023211121101021-3112022133311112-1013000103021013-3131302032212312-3112130132131102-0300211302323112-1110300220332030-2202211101320312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no s2s connectivity sli.

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

<a id="canonical-2020332030132111-0211012220122133-0321312221332102-2200333121310232-3111223133212200-2331110022130102-3133103133003221-2231231211133133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_s2s_connectivity_slo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_s2s_connectivity_slo

<a id="canonical-0230110110333031-0311111112012121-3303032310321332-0313321023302122-1120302120112112-0333332102333221-3131131213021211-3102303131313322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no s2s connectivity slo.

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

<a id="canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- nutanix

<a id="canonical-3332230322311330-3002233110322202-1013300122022103-1102232321033032-1321221122311231-0102010101321012-3100100301113010-2103000300030332"></a>

Type: `"single"`. Computed.

Nutanix Provider Type. Nutanix Provider Type.

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

<a id="canonical-0023100013302010-1203121212200322-2101303012100021-1111113232011302-2012212010203023-3012203103120031-2220023102012333-1320232311232010"></a>

### Direct properties for `nutanix`

- [not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210): complete subsection reference.

<a id="canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- nutanix.not_managed

<a id="canonical-3010312013300001-3022021300030233-2232123022231212-0121010110222011-1100210231020023-2023010202312210-3231323231312130-2313313020321121"></a>

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

<a id="canonical-1203111323221032-0003103210230301-0002010031011312-2010332322012002-3311200021311103-3000100021102121-1020110011131230-3331130112333300"></a>

### Direct properties for `nutanix.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213): complete subsection reference.

<a id="canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- nutanix.not_managed.node_list

<a id="canonical-2210330020132033-2020320230202002-3301121131030320-2001321022200022-0222021320103101-0123221233332202-3003200032311301-1311110130301131"></a>

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1103231223302221-2123321022011230-0300211113210333-1302033023231003-0310110012112023-1111012103022112-3010210002320103-0323012133101002"></a>

### Direct properties for `nutanix.not_managed.node_list`

<a id="canonical-3321130331300330-1201210123230233-1221021013220202-3130311113222210-0320003311233113-2103031303212223-2212033133121200-0312111133312222"></a>

#### `nutanix.not_managed.node_list.hostname` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232): complete subsection reference.

<a id="canonical-2323101212233122-0103202112121310-1220100201330202-2011110330331123-1333211112212213-0103303203021101-0230333011010333-1232020022313101"></a>

<a id="canonical-2133033303330313-1012112202201010-2231122321210310-3202302221333320-1312323023220203-3032110210112133-2203220112310322-1013122003210232"></a>

#### `nutanix.not_managed.node_list.public_ip` property

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0120133231203223-2210323113122301-3131011000301303-1020203101021110-2131321122211232-0130212132331330-3321331112231112-0002320223332200"></a>

<a id="canonical-3231000311013321-0121322101103102-0023022321022132-1202000102311102-1020222313231133-0222330110302130-0001010121322332-2032133031222013"></a>

#### `nutanix.not_managed.node_list.type` property

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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- nutanix.not_managed.node_list.interface_list

<a id="canonical-1001101002312322-1033111112002333-0100021120220121-2130233231303013-0203213130313122-2132011322230321-0203211000312312-0211030212313020"></a>

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0323231200110001-1023102000302231-0332300023103111-1222021211332101-1211220121230111-3322011122002122-0323003102130101-3313200120112313"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330): complete subsection reference.

<a id="canonical-0120132310102103-2001102020233013-1201011312020213-0001321232132120-2101202301030012-0000121132331333-0300120230001101-2210330112023022"></a>

<a id="canonical-2333330020033303-2122222010120332-3202033221202220-1312333322213301-3222213130110133-1021311213013201-2011222203022031-0221013210133332"></a>

#### `nutanix.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3103231010110021-2023313200103100-1120310120232031-3322130103032011-2120210133322033-0230203202133330-0003331122120101-3020301001222200): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0222200321012331-3320100132003121-1232322313113033-3320300022010221-2033103203122030-2301022032313011-1130313202022011-1320012112122330): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310): complete subsection reference.

<a id="canonical-3221303201012201-0023011222013010-3123221023013203-0300211031012230-2033120330332111-0112020322222112-0230203310333233-2023010233020023"></a>

<a id="canonical-3110321222212211-0020001211223100-3102323231220211-0332333312311012-1112030110113022-3130121121320303-2032223033110033-3333020212320121"></a>

#### `nutanix.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-1112201300330320-1110010022021031-2023012310323131-0103203103100210-3310330000323013-3322033333222211-0102311111332110-0112223022032300"></a>

<a id="canonical-1210131331210203-3000320101131111-3300211232310000-0323123332333332-0100013032103200-3100021233222100-0313111032313301-2321300210332301"></a>

#### `nutanix.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-0131030103203330-0233001110230011-2220112122013023-3221110321133032-1300003210001002-1232101000233100-2300013331010103-0321123200130310"></a>
