---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3330120120130113-2301032103112201-0233033103033231-0200123111011121-3232213321210302-1010132101311032-2322123003200111-2012332010001310"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 312211113231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
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
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
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

<a id="canonical-0333121323323031-1331013221313002-1301311102021330-3132320312110321-3210211130022301-0311311323021303-1003310011221013-2330203003032233"></a>

## Direct properties — interface_ip_map / 231002131210 / 3

<a id="canonical-2113131002221113-1132231120233121-3223310010101012-1112113123111101-0121201121101132-1333210010113232-2011301222330312-2311330320303033"></a>

<a id="canonical-1011113233233101-0112331233113311-2332113313333212-2200012203111222-1320000000103212-1130233302121013-2220112322023330-2011001311312122"></a>

## interface_ip_map property — interface_ip_map / 231002131210 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
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

<a id="canonical-1322013302033131-3202311021120121-1112223002103122-0000221311121313-2120111033212330-0003111103030330-2221102000333120-2103012203000022"></a>

## Direct properties — monitor / 110222311232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330030330123000-2110020322201121-1312321210220203-2332113232003103-2011101211021103-1220222003010012-0030023302102211-0002101103100301"></a>

## Next pages — monitor / 110222311232 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330202020000230-2132302030123012-3332113120321203-0302120003031122-0113112001220300-1130112322022111-3000020222001033-1000201100023122"></a>

## vmware.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 312133201330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311223021320123-0103001333123132-0221212021332202-2301232012313222-2131102030330301-0113123320123230-1022312331123031-1233312133321302"></a>

## vmware.not_managed.node_list.interface_list.network_option — network_option / 222200022311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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
EnumExtractionComplete: false
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

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-019.md#canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-019.md#canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132): complete subsection reference.

<a id="canonical-0310001300313133-0313121233033103-2101100232003202-1222233133231321-0231321323300323-3132132332100121-1033211231303011-2223323122313200"></a>

## Next pages — network_option / 222200022311 / 4

- [vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-019.md#canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120)
- [vmware.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-019.md#canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311122321331022-0003310002330110-3233203223211121-0311322201322023-1322121231110333-3201012200133321-0011203033011121-2320032323133312"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 330212102310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
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

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132010312010011-0222312112100311-0301113113303311-2330030310031002-1202131333320231-1222230212111313-2003212111000031-1010222203112030"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 311232232331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
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

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320220120031011-2333322012301133-1320330012300110-3311322222210032-0101331121030232-3331011200210233-1333310001220201-0222021101031020"></a>

## vmware.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 021133203101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123332120312210-3330131021230332-3332322221321001-3313101310212003-0211312113323322-0301112133012021-2033313300010023-2230112320002000"></a>

## vmware.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 310112110200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302111202333312-0030120133011203-0123311203100202-3101002220312223-3020021133122023-2212233030110123-3300132310013222-2120320101332201"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 220013223100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220101211221201-0100023030323111-3320311010300232-2001322203013011-0310112210230103-2303212012113223-1203102313333023-0313201001201202"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 100213100300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
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

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330130301101001-1303333012122231-1033312131300133-3212032313202301-0212131100323021-1013301203123121-1112002233122122-1101302011332211"></a>

## vmware.not_managed.node_list.interface_list.static_ip — static_ip / 301111131213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-0321323012033110-0130002002320113-0023221021200032-2022300001330200-2132301012303011-2120120330223323-0333002032103200-3320111300233213"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
EnumExtractionComplete: false
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
EnumExtractionComplete: false
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

<a id="canonical-3232132201313213-2012320323221102-0333013203021132-0120102120020333-2311132220310320-0010310103332203-2032123231213310-2233130102233301"></a>

## Next pages — static_ip / 301111131213 / 7

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020222103322302-1030123202301033-2101211013212202-2332003221232130-0201010020003331-1003132331322001-1131110033330330-2123233302013332"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 211003202132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-1210012201220010-1122313132100032-3202000232113200-0210133321320020-0111132302331030-0100202232322131-0201213310210013-2302130010302031"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031): complete subsection reference.

<a id="canonical-2021132302323313-0320211221212310-2210312303223101-2033013110200220-2213331312312120-0332222231012030-0232210003112020-2321202233310203"></a>

## Next pages — static_ipv6_address / 211003202132 / 4

- [vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322220031021120-2113221311332011-1201032133331333-1111333302002123-3132311111211023-2101233022102323-1131333101031201-3332331012332012"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 210133301003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3232032313302110-2021332121313203-2203311322303000-3222210122231102-2132121001301110-2330313330133002-0200311021020101-3113011312201020"></a>

## Next pages — cluster_static_ip / 210133301003 / 5

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301130303021200-2300033031120013-2001301112200123-3021022202212100-2130023221220020-3330221133203231-1330130033210121-1200301102313311"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 131100223312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2001232022213220-0212320021003220-1203212311232032-3331333302023320-3013312331222210-0230301110301220-0020102332002012-0103133331321300"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3122211310221101-1131111201020123-0133330200133222-3232221112130301-0032111101313200-3311310133032003-3300133331221220-3121210303222223"></a>

## Direct properties — node_static_ip / 131100223312 / 3

<a id="canonical-0030203311312021-0001330230101023-1202113101002013-1000213002121332-0220220310010013-2002202200132300-3330310200023330-3223132000113212"></a>

<a id="canonical-1321010312311022-2101230021231303-3030312202222330-3023000101333012-1022221130022210-3331212131033120-0321023131311303-0301122131010203"></a>

## default_gw property — node_static_ip / 131100223312 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3313302212131223-3112000233130233-0232223220221122-0100301223301231-3021130220301232-1321021120133331-0302012100000002-2213010320012123"></a>

<a id="canonical-0211323222100331-2101211201202102-3331120002121032-3302101131032303-3032321322232320-0311331202312212-3233211010101002-3120233320320102"></a>

## dns_server property — node_static_ip / 131100223312 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0301003230301122-1002301130031231-1213023323030110-1211002301122321-0302011331031202-3112213102013002-2021023230211030-1301131212213210"></a>

<a id="canonical-2213021212120130-1200033330132213-0000221100133123-2002233122032222-2212220230231133-0330133000301122-3133132133333021-3012013331300330"></a>

## ip_address property — node_static_ip / 131100223312 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0031230213110031-3310233330122303-3131222333231323-1233121222012100-0313301323011021-0033020322110232-1232303212012120-1313230322221323"></a>

## Next pages — node_static_ip / 131100223312 / 7

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100103112101210-0331313333013021-1221332111021210-0220031300331011-1023200323321021-1120122203100003-2233012233312323-0023021130220032"></a>

## vmware.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 030012111131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-018.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-018.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-018.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0033222112000013-1313333033331201-1321221321001220-1300120210320313-1102331313112030-2330301122321300-3131213121232031-2123332010230223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0210300232112030-1021023013032302-0120130210220131-0312222303012231-0230223031003000-0030000202030232-2020030001231100-0212113210233300"></a>

## Direct properties — vlan_interface / 030012111131 / 3

<a id="canonical-2303302012110113-2333301121032223-2320122233000101-1303001300222130-1313323001111112-1110010330123211-2130122000102120-2232022122331221"></a>

<a id="canonical-1103002321110102-1102133001233023-3110220231320310-1023300303222330-3111000131312331-3130321000210100-2322103300021123-0231020331113222"></a>

## device property — vlan_interface / 030012111131 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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

<a id="canonical-0011123311231110-1321322012202231-3103330130320202-0103102112022032-3211031113223300-2312213213001221-1222333000332331-0332221221001030"></a>

<a id="canonical-3021321101103202-3223023220203011-3303122020022030-0032003203123110-1021333021133333-2210323112312200-0031213033031023-3301122002100003"></a>

## vlan_id property — vlan_interface / 030012111131 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0332023330111322-3003213133310113-3120303301120120-0001223220023023-0110311123313201-0031230213023100-1322101203231113-3330033303230322"></a>

## Next pages — vlan_interface / 030012111131 / 6

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
